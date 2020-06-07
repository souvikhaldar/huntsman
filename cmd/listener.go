package cmd

import (
	"fmt"
	"io/ioutil"
	"net"

	"github.com/spf13/cobra"
)

var listenPort string

func init() {
	rootCmd.AddCommand(listenCmd)
	listenCmd.Flags().StringVar(&listenPort, "port", "8192", "Port at which listener should run")
}

var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Listen to incoming TCP requests on a port",
	Run: func(cmd *cobra.Command, args []string) {
		listener, err := net.Listen("tcp", ":"+listenPort)
		if err != nil {
			fmt.Println("Unable to run the listener: ", err)
			return
		}
		fmt.Println("Listening on: ", listenPort)
		defer listener.Close()
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Unable to take requests: ", err)
				return
			}
			go func(conn net.Conn) {
				req, err := ioutil.ReadAll(conn)
				if err != nil {
					fmt.Println("Unable to read the request body: ", err)
				}
				fmt.Println("Request: ", string(req))
				conn.Close()
			}(conn)

		}
	},
}
