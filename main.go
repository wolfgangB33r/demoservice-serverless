package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type callee struct {
	Adr   string // URL address to call
	Count int    // number of calls per request
}

type config struct {
	Callees []callee
}

func handleIcon(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	// read config from environment variable
	configStr := os.Getenv("DEMO_SERVICE_CALLEES")

	// Check if the environment variable is set
	if configStr == "" {
		fmt.Fprintf(&buf, "Configuration environment variable DEMO_SERVICE_CONFIG is not set.\n")
		w.Write(buf.Bytes())
		return
	}
	var conf config
	// Unmarshal the config
	err := json.Unmarshal([]byte(configStr), &conf)
	if err != nil {
		fmt.Fprintf(&buf, "Configuration JSON is wrong.\n")
		fmt.Fprintf(&buf, configStr)
	} else {
		// Configuration is available
		for _, element := range conf.Callees {
			for i := 0; i < element.Count; i++ {
				// Send GET request
				response, err := http.Get(element.Adr)
				if err != nil {
					fmt.Fprintf(&buf, "Error sending GET request: %s\n", err)
				} else {
					fmt.Fprintf(&buf, "Called address (%s) returned with http %d\n", element.Adr, response.StatusCode)

				}
				defer response.Body.Close()
			}
		}
	}
	w.Write(buf.Bytes())
	defer r.Body.Close()
}

func main() {
	port := 8080
	http.HandleFunc("/", handleRequest)
	http.HandleFunc("/favicon.ico", handleIcon)
	if err := http.ListenAndServe(":"+strconv.Itoa(port), nil); err != nil {
		panic(err)
	}
}
