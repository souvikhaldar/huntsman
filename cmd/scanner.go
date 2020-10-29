package cmd

import (
	"fmt"
	"net"
	"strconv"
	"sync"

	"github.com/spf13/cobra"
)

var target string
var start int32
var end int32
var threads int32
var scanCmd = &cobra.Command{
	Use:   "portscan",
	Short: "Scan for open ports",
	Long:  "Concurrently scan the provided range (by default 0 to 65535) to check if any port is open",
	Run: func(cmd *cobra.Command, args []string) {
		threadPool := make(chan int, threads)
		openCount := make(chan int)
		var wg sync.WaitGroup
		go func(threadPool chan int, openCount chan int) {
			for i := start; i <= end; i++ {
				threadPool <- 1
				wg.Add(1)
				go func(target string, threadCount chan int, port int32) {
					isOpen(target, openCount, port)
				}(target, openCount, i)
			}
		}(threadPool, openCount)
		numOfPorts := end - start
		num := int32(0)
		for open := range openCount {
			num++
			if open != 0 {
				fmt.Println("Open port: ", open)
			}
			<-threadPool
			wg.Done()
			if num == numOfPorts {
				break
			}

		}
		wg.Wait()

	},
}

func isOpen(target string, openCount chan int, port int32) {

	address := target + ":" + strconv.Itoa(int(port))
	if _, err := net.Dial("tcp", address); err == nil {
		openCount <- int(port)
		return
	}
	openCount <- 0
}

func init() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.Flags().StringVar(&target, "target", "", "IP/URL address of the machine to be scanned")
	scanCmd.Flags().Int32VarP(&start, "start", "s", 1, "starting port number")
	scanCmd.Flags().Int32VarP(&end, "end", "e", 65535, "last port number")
	scanCmd.Flags().Int32VarP(&threads, "threads", "t", 100, "the number of goroutines to execute at a time")
}
