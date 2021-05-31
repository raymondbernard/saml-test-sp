package main

import (
	"github.com/raymondbernard/saml-test-sp-go/pkg/server"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)

	server.RunServer()
}
