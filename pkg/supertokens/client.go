package supertokens

import (
	"context"
	"net/http"
	"time"

	rtClient "github.com/go-openapi/runtime/client"
	"github.com/gogetkarthik/service-specification/supertokens/client"
	"github.com/gogetkarthik/service-specification/supertokens/models"
)

type (
	superTokens struct {
		basePath  string
		host      string
		scheme    string
		hello     helloService
		handshake handshakeService
	}

	Client interface {
		//hello end point
		HelloGet() (*models.HelloOutput, error)
		HelloPost() (*models.HelloOutput, error)
		HelloPut() (*models.HelloOutput, error)
		HelloDelete() (*models.HelloOutput, error)

		//Handshake
		Handshake(*models.DeviceDriverInfoType) (*models.HandshakeOutput, error)
	}
	Options func(st superTokens)
)

func (s superTokens) HelloGet() (*models.HelloOutput, error) {
	return s.hello.get()
}

func (s superTokens) HelloPost() (*models.HelloOutput, error) {
	return s.hello.post()
}

func (s superTokens) HelloPut() (*models.HelloOutput, error) {
	return s.hello.put()
}

func (s superTokens) HelloDelete() (*models.HelloOutput, error) {
	return s.hello.delete()
}

func (s superTokens) Handshake(dt *models.DeviceDriverInfoType) (*models.HandshakeOutput, error) {
	return s.handshake.Handshake(dt)
}

func New(options ...Options) Client {

	st := superTokens{
		host:     client.DefaultHost,
		scheme:   client.DefaultSchemes[0],
		basePath: client.DefaultBasePath,
	}

	for _, option := range options {
		option(st)
	}

	httpClient := http.Client{
		Transport: &http.Transport{},
		Timeout:   60 * time.Second,
	}
	rt := rtClient.NewWithClient("localhost:3567", "/", []string{"http"}, &httpClient)
	rt.Context = context.Background()
	stService := client.New(rt, nil)
	st.hello = newHello(stService.Hello)
	st.handshake = newHandshake(stService.Handshake)

	return st
}

func WithHost(host string) Options {
	return func(st superTokens) {
		st.host = host
	}
}

func WithScheme(scheme string) Options {
	return func(st superTokens) {
		st.scheme = scheme
	}
}

func WithBasePath(basePath string) Options {
	return func(st superTokens) {
		st.basePath = basePath
	}
}
