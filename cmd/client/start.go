package main

import (
	"github.com/IBM/gedsmds/internal/mockclient"
)

func main() {
	ex := mockclient.NewExecutor()
	ex.SendSubscription()
}
