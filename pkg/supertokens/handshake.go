package supertokens

import (
	hClient "github.com/gogetkarthik/service-specification/supertokens/client/handshake"
	"github.com/gogetkarthik/service-specification/supertokens/models"
)

type handshake struct {
	client hClient.ClientService
}

func newHandshake(hc hClient.ClientService) handshakeService {
	return handshake{client: hc}
}

type handshakeService interface {
	Handshake(*models.DeviceDriverInfoType) (*models.HandshakeOutput, error)
}

func (h handshake) Handshake(dt *models.DeviceDriverInfoType) (*models.HandshakeOutput, error) {
	hp := hClient.NewHandshakeParams()
	hp.DeviceDriverInfoType = dt

	rsp, err := h.client.Handshake(hp)
	if err != nil {
		return nil, err
	}
	return rsp.Payload, nil
}
