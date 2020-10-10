package cmd

import (
	"fmt"
	"io"
	"net"

	"github.com/spf13/cobra"
)

var destAdd, port, servPort string

func init() {
	rootCmd.AddCommand(proxyCmd)
	proxyCmd.Flags().StringVarP(&destAdd, "target-address", "t", "", "Destination to forward traffic to")
	proxyCmd.Flags().StringVarP(&port, "target-port", "p", "80", "The port of the destination, eg 8192,80")
	proxyCmd.Flags().StringVarP(&servPort, "port", "s", "8192", "The port at which this proxy should run")
}

var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Relay traffic from source to destination via this proxy",
	Run: func(cmd *cobra.Command, args []string) {
		listener, err := net.Listen("tcp", ":"+servPort)
		if err != nil {
			fmt.Println("Unable to run the proxy: ", err)
			return
		}
		fmt.Println("Proxy listening on: ", servPort)
		defer listener.Close()
		for {
			src, err := listener.Accept()
			if err != nil {
				fmt.Println("Unable to listen to requests: ", err)
				return
			}
			go proxy(src)
		}

	},
}

func proxy(src net.Conn) {

	dest := destAdd + ":" + port
	dst, err := net.Dial("tcp", dest)
	if err != nil {
		fmt.Println("Error in connecting to the destination: ", err)
		return
	}
	fmt.Println("Connected to target: ", dest)
	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			fmt.Println("Error in proxying data: ", err)
		}
		//// echo back the response
		//if _, err := io.Copy(src, dst); err != nil {
		//	fmt.Println("Reply the response: ", err)
		//}
		dst.Close()

	}()
}
