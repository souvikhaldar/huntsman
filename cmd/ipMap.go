package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var ipCmd = &cobra.Command{
	Use:   "ipinfo",
	Short: "Fetch the location information of the IP",
	Run: func(cmd *cobra.Command, args []string) {

		if tcpDump != "" {
			f, err := os.Open(tcpDump)
			if err != nil {
				fmt.Println("Unable to open log file", err)
				return
			}

			scanner := bufio.NewScanner(f)
			for {
				for scanner.Scan() {
					if strings.Contains(scanner.Text(), "IP") {
						ip, err := ParseIPFromTcpDump(scanner.Text())
						if err != nil {
							fmt.Println(err)
							return
						}
						fmt.Println("Request came from: ", ip)
						body, err := getIPInfo(ip)
						if err != nil {
							fmt.Println(err)
							return
						}

						fmt.Println(
							"Details of the requester: ",
							string(body))
					}
				}
				if err := scanner.Err(); err != nil {
					fmt.Println(err)
					return
				}
			}
		} else {
			body, err := getIPInfo(ipAdd)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Details of the requester: ", string(body))
		}

	},
}

func getIPInfo(ip string) ([]byte, error) {

	resp, err := http.Get(fmt.Sprintf("http://ip-api.com/json/%s", ip))
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return body, err
	}
	return body, err
}

var tcpDump string
var ipAdd string

func init() {
	rootCmd.AddCommand(ipCmd)
	ipCmd.Flags().StringVarP(&tcpDump, "tcp-dump", "t", "", "source file of tcpdump")
	ipCmd.Flags().StringVar(&ipAdd, "ip", "", "IP address of the target")
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
