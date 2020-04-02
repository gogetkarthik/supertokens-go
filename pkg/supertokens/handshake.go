package supertokens

import (
	hClient "github.com/gogetkarthik/service-specification/supertokens/client/handshake"
	"github.com/gogetkarthik/service-specification/supertokens/models"
)

type handshake struct {
	client hClient.Client
}

type handshakeService interface {
	Handshake() (*models.HandshakeOutput, error)
}

func (h handshake) Handshake() (*models.HandshakeOutput, error) {
	hp := &hClient.HandshakeParams{}

	rsp, err := h.client.Handshake(hp)
	if err != nil {
		return nil, err
	}
	return rsp.Payload, nil
}
