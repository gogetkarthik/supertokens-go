package main

import (
	"fmt"
	"github.com/gogetkarthik/service-specification/supertokens/models"
	"log"

	"github.com/gogetkarthik/supertokens-go/pkg/supertokens"
)

func main() {
	s := supertokens.New(supertokens.WithHost("localhost:3567"))
	rsp, err := s.HelloGet()
	if err != nil {
		log.Fatal("hello get err: %v", err)
	}
	fmt.Printf("hello get resp: %s", *rsp)

	rsp, err = s.HelloPost()
	if err != nil {
		log.Fatal("hello get err: %v", err)
	}
	fmt.Printf("hello post resp: %s", *rsp)

	rsp, err = s.HelloPut()
	if err != nil {
		log.Fatal("hello get err: %v", err)
	}
	fmt.Printf("hello put resp: %s", *rsp)

	rsp, err = s.HelloDelete()
	if err != nil {
		log.Fatal("hello get err: %v", err)
	}
	fmt.Printf("hello delete resp: %s", *rsp)

	handRsp, err := s.Handshake(&models.DeviceDriverInfoType{
		Driver: &models.Driver{
			Name:    "supertokens-go",
			Version: "0.0",
		},
		FrontendSDK: []*models.FrontendSDK{
			{
				Name:    "vuejs",
				Version: "1.1",
			},
			{
				Name:    "ReachJS",
				Version: "2",
			},
		},
	})
	if err != nil {
		log.Fatal("handshake err: %v", err)
	}
	fmt.Printf("handshake resp %#v", handRsp)

}
