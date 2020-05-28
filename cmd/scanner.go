package cmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

var target net.IP
var start int32
var end int32
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan for open ports",
	Long:  "Concurrently scan the provided range (by default 0 to 65535) to check if any port is open",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(start, end)
	},
}

func init() {

	rootCmd.AddCommand(scanCmd)
	scanCmd.Flags().IPVarP(&target, "target", "t", net.IPv4(127, 0, 0, 1), "IP address of the machine to be scanned")
	scanCmd.Flags().Int32VarP(&start, "start", "s", 1, "starting port number")
	scanCmd.Flags().Int32VarP(&end, "end", "e", 65535, "last port number")
}
