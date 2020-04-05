package supertokens

import (
	"errors"
	"testing"

	hClient "github.com/gogetkarthik/service-specification/supertokens/client/hello"
	"github.com/gogetkarthik/service-specification/supertokens/models"
	"github.com/gogetkarthik/supertokens-go/mocks/mock_hello"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_hello_get(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mcs := mock_hello.NewMockClientService(mc)
	helloResp := "hello"
	pl := models.HelloOutput(helloResp)
	hp := hClient.NewHelloGetParams()

	mcs.EXPECT().HelloGet(gomock.Eq(hp)).Return(&hClient.HelloGetOK{pl}, nil)
	//Success
	hs := newHello(mcs)
	resp, err := hs.get()
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "hello", string(*resp))

	//Failure
	mcs.EXPECT().HelloGet(gomock.Eq(hp)).Return(nil, errors.New("Bad request"))
	resp, err = hs.get()
	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.Equal(t, "Bad request", err.Error())
}

func Test_hello_post(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mcs := mock_hello.NewMockClientService(mc)
	helloResp := "hello"
	pl := models.HelloOutput(helloResp)
	hp := hClient.NewHelloPostParams()

	mcs.EXPECT().HelloPost(gomock.Eq(hp)).Return(&hClient.HelloPostOK{pl}, nil)

	hs := newHello(mcs)
	resp, err := hs.post()
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "hello", string(*resp))

	mcs.EXPECT().HelloPost(gomock.Eq(hp)).Return(nil, errors.New("Bad request"))
	resp, err = hs.post()
	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.Equal(t, "Bad request", err.Error())
}

func Test_hello_put(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mcs := mock_hello.NewMockClientService(mc)
	helloResp := "hello"
	pl := models.HelloOutput(helloResp)
	hp := hClient.NewHelloPutParams()

	mcs.EXPECT().HelloPut(gomock.Eq(hp)).Return(&hClient.HelloPutOK{pl}, nil)

	hs := newHello(mcs)
	resp, err := hs.put()
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "hello", string(*resp))

	mcs.EXPECT().HelloPut(gomock.Eq(hp)).Return(nil, errors.New("Bad request"))
	resp, err = hs.put()
	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.Equal(t, "Bad request", err.Error())
}

func Test_hello_delete(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mcs := mock_hello.NewMockClientService(mc)
	helloResp := "hello"
	pl := models.HelloOutput(helloResp)
	hp := hClient.NewHelloDeleteParams()

	mcs.EXPECT().HelloDelete(gomock.Eq(hp)).Return(&hClient.HelloDeleteOK{pl}, nil)

	hs := newHello(mcs)
	resp, err := hs.delete()
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "hello", string(*resp))

	mcs.EXPECT().HelloDelete(gomock.Eq(hp)).Return(nil, errors.New("Bad request"))
	resp, err = hs.delete()
	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.Equal(t, "Bad request", err.Error())
}
