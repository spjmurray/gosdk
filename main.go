package main

import (
	"flag"
	"fmt"

	"github.com/couchbase/gocb"
)

func main() {
	connection := flag.String("connection", "", "Hostname")
	username := flag.String("username", "", "Username")
	password := flag.String("password", "", "Password")
	cafile := flag.String("cafile", "", "CA filename")

	flag.Parse()

	connstr := fmt.Sprintf("%s?cacertpath=%s", *connection, *cafile)
	fmt.Println("using connection string", connstr)

	cluster, err := gocb.Connect(connstr)
	if err != nil {
		fmt.Println("failed to open connection:", err)
		return
	}

	authenticator := gocb.PasswordAuthenticator{
		Username: *username,
		Password: *password,
	}

	if err := cluster.Authenticate(authenticator); err != nil {
		fmt.Println("failed to authenticate:", err)
		return
	}

	if _, err := cluster.OpenBucket("default", ""); err != nil {
		fmt.Println("failed to use bucket:", err)
		return
	}
}
