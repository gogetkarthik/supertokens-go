package supertokens

import (
	"errors"
	"reflect"
	"testing"

	hClient "github.com/gogetkarthik/service-specification/supertokens/client/handshake"
	"github.com/gogetkarthik/service-specification/supertokens/models"
	"github.com/gogetkarthik/supertokens-go/mocks/mock_handshake"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_handshake_Handshake(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mcs := mock_handshake.NewMockClientService(mc)

	hi := &models.DeviceDriverInfoType{
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
	}
	hp := hClient.NewHandshakeParams()
	hp.DeviceDriverInfoType = hi

	handshakeOutput := &models.HandshakeOutput{
		AccessTokenBlacklistingEnabled: false,
		AccessTokenPath:                "/",
		CookieDomain:                   "supertokens.io",
		CookieSameSite:                 "",
		CookieSecure:                   false,
		EnableAntiCsrf:                 true,
		IDRefreshTokenPath:             "",
		JwtSigningPublicKey:            "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqWM2PuNk+iLfdKUhpHAEhiqX29MVN1K07d0ifmXYdj8YLWDrklBJfrxUY8QvM/bJvzjPmeXl3CbWCetL0hz5ZCaN8zTwrosv9qziBIyTJZjTx3I/TlyVokrzXP7368dxtXBPQxDlRUDtB7ltzaxHFx46x7tYIpmPVPFsHC4J/O/4RlQY0DCbjok5ScVjtdZ0TKqHsCMPTGtLXFjzTpF1TtxLViLGRB1uiGVuXBN530lKgxF6FkJv85kmS1z1dlUZD7o6wO6XYGoc/3CnQZ94b0AdfoR+lHm8UzOsyECiMDTPUV9Br7w+ykGobxvcc55EzA59iOboixFg9a6lRb/qtQIDAQAB",
		JwtSigningPublicKeyExpiryTime:  9223372036854775807,
		RefreshTokenPath:               "/refresh",
		Status:                         "OK",
	}
	mcs.EXPECT().Handshake(gomock.Eq(hp)).Return(&hClient.HandshakeOK{Payload: handshakeOutput}, nil)
	//Success
	hs := newHandshake(mcs)
	resp, err := hs.Handshake(hi)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.True(t, reflect.DeepEqual(handshakeOutput, resp))

	//Failure
	mcs.EXPECT().Handshake(gomock.Eq(hp)).Return(nil, errors.New("Bad request"))
	resp, err = hs.Handshake(hi)
	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.Equal(t, "Bad request", err.Error())
}
