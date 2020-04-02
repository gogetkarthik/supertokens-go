package supertokens

import (
	"context"
	"github.com/gogetkarthik/service-specification/supertokens/client/hello"
	"github.com/gogetkarthik/service-specification/supertokens/models"
)

type helloPKG struct {
	client hello.ClientService
}

type helloService interface {
	get() (*models.HelloOutput, error)
	post() (*models.HelloOutput, error)
	put() (*models.HelloOutput, error)
	delete() (*models.HelloOutput, error)
}

func NewHelloServ(c hello.ClientService) helloService {
	return helloPKG{client: c}
}

func (h helloPKG) get() (*models.HelloOutput, error) {
	hp := &hello.HelloGetParams{
		Context: context.Background(),
	}

	rsp, err := h.client.HelloGet(hp)
	if err != nil {
		return nil, err
	}
	return &rsp.Payload, nil
}

func (h helloPKG) post() (*models.HelloOutput, error) {
	hp := &hello.HelloPostParams{}

	rsp, err := h.client.HelloPost(hp)
	if err != nil {
		return nil, err
	}
	return &rsp.Payload, nil
}

func (h helloPKG) put() (*models.HelloOutput, error) {
	hp := &hello.HelloPutParams{}

	rsp, err := h.client.HelloPut(hp)
	if err != nil {
		return nil, err
	}
	return &rsp.Payload, nil
}

func (h helloPKG) delete() (*models.HelloOutput, error) {
	hp := &hello.HelloDeleteParams{}

	rsp, err := h.client.HelloDelete(hp)
	if err != nil {
		return nil, err
	}
	return &rsp.Payload, nil
}
