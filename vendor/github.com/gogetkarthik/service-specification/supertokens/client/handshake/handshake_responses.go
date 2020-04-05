// Code generated by go-swagger; DO NOT EDIT.

package handshake

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/gogetkarthik/service-specification/supertokens/models"
)

// HandshakeReader is a Reader for the Handshake structure.
type HandshakeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HandshakeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHandshakeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHandshakeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewHandshakeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewHandshakeMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHandshakeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHandshakeOK creates a HandshakeOK with default headers values
func NewHandshakeOK() *HandshakeOK {
	return &HandshakeOK{}
}

/*HandshakeOK handles this case with default header values.

Handshake destails
*/
type HandshakeOK struct {
	Payload *models.HandshakeOutput
}

func (o *HandshakeOK) Error() string {
	return fmt.Sprintf("[POST /handshake][%d] handshakeOK  %+v", 200, o.Payload)
}

func (o *HandshakeOK) GetPayload() *models.HandshakeOutput {
	return o.Payload
}

func (o *HandshakeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HandshakeOutput)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHandshakeBadRequest creates a HandshakeBadRequest with default headers values
func NewHandshakeBadRequest() *HandshakeBadRequest {
	return &HandshakeBadRequest{}
}

/*HandshakeBadRequest handles this case with default header values.

Invalid input
*/
type HandshakeBadRequest struct {
	Payload models.Error
}

func (o *HandshakeBadRequest) Error() string {
	return fmt.Sprintf("[POST /handshake][%d] handshakeBadRequest  %+v", 400, o.Payload)
}

func (o *HandshakeBadRequest) GetPayload() models.Error {
	return o.Payload
}

func (o *HandshakeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHandshakeNotFound creates a HandshakeNotFound with default headers values
func NewHandshakeNotFound() *HandshakeNotFound {
	return &HandshakeNotFound{}
}

/*HandshakeNotFound handles this case with default header values.

The specified resource was not found
*/
type HandshakeNotFound struct {
	Payload models.Error
}

func (o *HandshakeNotFound) Error() string {
	return fmt.Sprintf("[POST /handshake][%d] handshakeNotFound  %+v", 404, o.Payload)
}

func (o *HandshakeNotFound) GetPayload() models.Error {
	return o.Payload
}

func (o *HandshakeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHandshakeMethodNotAllowed creates a HandshakeMethodNotAllowed with default headers values
func NewHandshakeMethodNotAllowed() *HandshakeMethodNotAllowed {
	return &HandshakeMethodNotAllowed{}
}

/*HandshakeMethodNotAllowed handles this case with default header values.

Method not souported
*/
type HandshakeMethodNotAllowed struct {
	Payload models.Error
}

func (o *HandshakeMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /handshake][%d] handshakeMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *HandshakeMethodNotAllowed) GetPayload() models.Error {
	return o.Payload
}

func (o *HandshakeMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHandshakeInternalServerError creates a HandshakeInternalServerError with default headers values
func NewHandshakeInternalServerError() *HandshakeInternalServerError {
	return &HandshakeInternalServerError{}
}

/*HandshakeInternalServerError handles this case with default header values.

Internal server error
*/
type HandshakeInternalServerError struct {
	Payload models.Error
}

func (o *HandshakeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /handshake][%d] handshakeInternalServerError  %+v", 500, o.Payload)
}

func (o *HandshakeInternalServerError) GetPayload() models.Error {
	return o.Payload
}

func (o *HandshakeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}