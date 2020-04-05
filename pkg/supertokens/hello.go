package supertokens

import (
	hClient "github.com/gogetkarthik/service-specification/supertokens/client/hello"
	"github.com/gogetkarthik/service-specification/supertokens/models"
)

type hello struct {
	client hClient.ClientService
}

type helloService interface {
	get() (*models.HelloOutput, error)
	post() (*models.HelloOutput, error)
	put() (*models.HelloOutput, error)
	delete() (*models.HelloOutput, error)
}

func newHello(c hClient.ClientService) helloService {
	return hello{client: c}
}

func (h hello) get() (*models.HelloOutput, error) {
	hp := hClient.NewHelloGetParams()

	rsp, err := h.client.HelloGet(hp)
	if err != nil {
		return nil, err
	}
	return &rsp.Payload, nil
}

func (h hello) post() (*models.HelloOutput, error) {
	hp := hClient.NewHelloPostParams()

	rsp, err := h.client.HelloPost(hp)
	if err != nil {
		return nil, err
	}
	return &rsp.Payload, nil
}

func (h hello) put() (*models.HelloOutput, error) {
	hp := hClient.NewHelloPutParams()

	rsp, err := h.client.HelloPut(hp)
	if err != nil {
		return nil, err
	}
	return &rsp.Payload, nil
}

func (h hello) delete() (*models.HelloOutput, error) {
	hp := hClient.NewHelloDeleteParams()

	rsp, err := h.client.HelloDelete(hp)
	if err != nil {
		return nil, err
	}
	return &rsp.Payload, nil
}
