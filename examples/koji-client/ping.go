package main

import (
	"fmt"
	"os"

	kojiClient "github.com/odra/kubeji/pkg/koji-client"
)

// KOJI_HUB_URL defaults to https://koji.fedoraproject.org/kojihub
// KOJI_HUB_URL=$URL go run examples/koji-client/ping.go
// go run examples/koji-client/ping.go
func main() {
	// try to get the hub url from an env var
	// falls back to https://koji.fedoraproject.org/kojihub otherwise
	url := os.Getenv("KOJI_HUB_URL")
	if url == "" {
		url = "https://koji.fedoraproject.org/kojihub"
	}

	fmt.Printf("Using %s\n", url)

	// creates a new client instance
	client, err := kojiClient.New(url)
	if err != nil {
		fmt.Printf("KojiClient.New error %s", err)
		os.Exit(1)
	}

	// pings the remote xmlrpc server
	err = client.Ping()
	if err != nil {
		fmt.Printf("KojiClient.Ping error %s", err)
		os.Exit(1)
	}

	fmt.Println("KojiClient.Ping success!")

	// closes client connection
	err = client.Defer()
	if err != nil {
		fmt.Printf("KojiClient.Close error %s", err)
		os.Exit(1)
	}

	os.Exit(0)
}
