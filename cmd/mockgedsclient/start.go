package main

import (
	"github.com/IBM/gedsmds/internal/mockgedsclient"
	//_ "google.golang.org/grpc/encoding/gzip"
)

func main() {
	ex := mockgedsclient.NewExecutor()
	ex.SendSubscription()
}
