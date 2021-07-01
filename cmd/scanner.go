package cmd

import (
	"fmt"
	"net"
	"strconv"
	"time"

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
		openCount := make(chan int, end-1)

		go func(threadPool, openCount chan int) {
			for i := start; i < end; i++ {
				threadPool <- 1
				go func(target string, openCount chan int, port int32) {
					defer func() {
						<-threadPool
						if i == end-1 {
							time.Sleep(1 * time.Second)
							close(openCount)
						}
					}()

					if isOpen(target, port) {
						openCount <- int(port)
					}
					return
				}(target, openCount, i)
			}
		}(threadPool, openCount)

		for open := range openCount {
			if open != 0 {
				fmt.Println("Open port: ", open)
			}

		}

	},
}

func isOpen(target string, port int32) bool {

	address := target + ":" + strconv.Itoa(int(port))
	if _, err := net.Dial("tcp", address); err == nil {
		return true
	}
	return false
}

func init() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.Flags().StringVar(&target, "target", "", "IP/URL address of the machine to be scanned")
	scanCmd.Flags().Int32VarP(&start, "start", "s", 1, "starting port number")
	scanCmd.Flags().Int32VarP(&end, "end", "e", 65535, "last port number")
	scanCmd.Flags().Int32VarP(&threads, "threads", "t", 100, "the number of goroutines to execute at a time")
}
