package transmission

import (
	"fmt"
	"github.com/hekmon/transmissionrpc/v3"
	"net/url"
	"os"
)

const ConnectionUrl = "http://%s:%s@%s:%s/transmission/rpc"

var clientInstance *transmissionrpc.Client

func Client() *transmissionrpc.Client {
	if clientInstance == nil {
		clientInstance = setupClient()
	}

	return clientInstance
}

func setupClient() *transmissionrpc.Client {
	host := os.Getenv("TRANSMISSION_HOST")
	port := os.Getenv("TRANSMISSION_PORT")
	user := os.Getenv("TRANSMISSION_USER")
	password := os.Getenv("TRANSMISSION_PASSWORD")

	endpoint, err := url.Parse(fmt.Sprintf(ConnectionUrl, user, password, host, port))
	if err != nil {
		panic(err)
	}

	client, err := transmissionrpc.New(endpoint, nil)
	if err != nil {
		panic(err)
	}

	return client
}
