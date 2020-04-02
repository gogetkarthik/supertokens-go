package main

import (
	"fmt"
	"log"

	"github.com/supertokens-go/pkg/supertokens"
)

func main() {
	s := supertokens.New(supertokens.WithHost("localhost:3567"))
	rsp, err := s.HelloGet()
	if err != nil {
		log.Fatal("not able to talk to server %s", err)
	}

	fmt.Printf("server response: %s", *rsp)
}
