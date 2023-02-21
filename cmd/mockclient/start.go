package main

import (
	"github.com/IBM/gedsmds/internal/mockclient"
	//_ "google.golang.org/grpc/encoding/gzip"
)

func main() {
	ex := mockclient.NewExecutor()
	ex.SendSubscription()
}
