package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	var certPath string
	var certKey string
	var port int

	flag.StringVar(&certPath, "tls-cert", "/opt/simple-https-listener/tls.crt", "path to the TLS certificate")
	flag.StringVar(&certKey, "tls-key", "/opt/simple-https-listener/tls.key", "path to the TLS certificate key")
	flag.IntVar(&port, "listening-port", 8443, "port accepting HTTPS connection")

	flag.Parse()

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		h, _ := os.Hostname()
		_, _ = writer.Write([]byte(h))
	})

	if err := http.ListenAndServeTLS(fmt.Sprintf("0.0.0.0:%d", port), certPath, certKey, nil); err != nil {
		panic(err)
	}
}
