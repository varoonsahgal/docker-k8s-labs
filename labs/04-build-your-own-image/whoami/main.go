// whoami: a tiny web server that tells you which container answered.
// Run several copies and watch the hostname change. That's the whole trick.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// version is stamped in at build time:
//
//	go build -ldflags "-X main.version=v2"
var version = "dev"

const whale = `
        ##         .
  ## ## ##        ==
 ## ## ## ##     ===
/"""""""""""""""\___/ ===
{                      /  ===-
\______ o           __/
  \    \         __/
   \____\_______/
`

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func main() {
	health := flag.Bool("health", false, "probe the running server and exit 0 if healthy")
	flag.Parse()

	port := getenv("PORT", "8080")

	// Used by the Dockerfile HEALTHCHECK. The final image has no shell and no curl,
	// so the binary checks itself.
	if *health {
		client := http.Client{Timeout: 2 * time.Second}
		resp, err := client.Get("http://127.0.0.1:" + port + "/healthz")
		if err != nil || resp.StatusCode != http.StatusOK {
			os.Exit(1)
		}
		os.Exit(0)
	}

	host, _ := os.Hostname()
	message := getenv("MESSAGE", "Hello from inside a container!")
	started := time.Now()

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, whale)
		fmt.Fprintf(w, "\n  %s\n\n", message)
		fmt.Fprintf(w, "  hostname : %s\n", host)
		fmt.Fprintf(w, "  version  : %s\n", version)
		fmt.Fprintf(w, "  uptime   : %s\n", time.Since(started).Round(time.Second))
	})

	log.Printf("whoami %s listening on :%s", version, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
