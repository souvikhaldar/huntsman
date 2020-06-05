package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var ipCmd = &cobra.Command{
	Use:   "iploc",
	Short: "Fetch the location information of the IP",
	Run: func(cmd *cobra.Command, args []string) {
		var ip string
		var err error
		//		if len(args) == 0 {
		//			fmt.Println("Provide the input")
		//			return
		//		}
		//		fmt.Println("input is: ", args[0])
		if tcpDump {
			if file == "" {
				fmt.Println("tcpdump log file source not provided")
				return
			}
			f, err := os.Open(file)
			if err != nil {
				fmt.Println("Unable to open log file", err)
				return
			}
			fileContent, err := ioutil.ReadAll(f)
			if err != nil {
				fmt.Println("Unable to read log file", err)
				return
			}
			ip, err = ParseIPFromTcpDump(string(fileContent))
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			ip = args[0]
		}
		resp, err := http.Get(fmt.Sprintf("http://ip-api.com/json/%s", ip))
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Body: ", string(body))
		return

	},
}
var tcpDump bool
var file string

func init() {
	rootCmd.AddCommand(ipCmd)
	ipCmd.Flags().BoolVar(&tcpDump, "tcpdump", false, "Is the source from tcpdump")
	ipCmd.Flags().StringVarP(&file, "file", "f", "", "source file of tcpdump")
}

func ParseIPFromTcpDump(tcpDump string) (string, error) {
	split := strings.Split(tcpDump, "\n")
	if len(split) == 0 {
		return "", fmt.Errorf("Error in parsing tcpdump output: %s", "split")
	}
	newSplit := strings.SplitAfter(split[0], ">")[0]
	sliceSplit := strings.Fields(newSplit)
	if len(sliceSplit) < 3 {
		return "", fmt.Errorf("Error in parsing tcpdump output: %s", "sliceSplit")
	}
	lastDot := strings.LastIndex(sliceSplit[2], ".")
	if lastDot == -1 {
		return "", fmt.Errorf("Error in parsing tcpdump output: %s", "lastDot")
	}
	if len(sliceSplit[2]) < lastDot {
		return "", fmt.Errorf("Error in parsing tcpdump output: %s", "lastDot length")
	}
	return sliceSplit[2][:lastDot], nil
}
