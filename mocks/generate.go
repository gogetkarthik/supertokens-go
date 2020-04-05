//go:generate mockgen  -destination ./mock_hello/mock_supertokens_hello.go github.com/gogetkarthik/service-specification/supertokens/client/hello ClientService
//go:generate mockgen  -destination ./mock_handshake/mock_supertokens_handshake.go github.com/gogetkarthik/service-specification/supertokens/client/handshake ClientService

package mocks
