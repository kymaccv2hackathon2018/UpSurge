// Code generated by go-swagger; DO NOT EDIT.

package carts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

// GuestLoginUsingPUTReader is a Reader for the GuestLoginUsingPUT structure.
type GuestLoginUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GuestLoginUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGuestLoginUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewGuestLoginUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGuestLoginUsingPUTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGuestLoginUsingPUTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGuestLoginUsingPUTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGuestLoginUsingPUTOK creates a GuestLoginUsingPUTOK with default headers values
func NewGuestLoginUsingPUTOK() *GuestLoginUsingPUTOK {
	return &GuestLoginUsingPUTOK{}
}

/*GuestLoginUsingPUTOK handles this case with default header values.

OK
*/
type GuestLoginUsingPUTOK struct {
}

func (o *GuestLoginUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /{baseSiteId}/users/{userId}/carts/{cartId}/email][%d] guestLoginUsingPUTOK ", 200)
}

func (o *GuestLoginUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGuestLoginUsingPUTCreated creates a GuestLoginUsingPUTCreated with default headers values
func NewGuestLoginUsingPUTCreated() *GuestLoginUsingPUTCreated {
	return &GuestLoginUsingPUTCreated{}
}

/*GuestLoginUsingPUTCreated handles this case with default header values.

Created
*/
type GuestLoginUsingPUTCreated struct {
}

func (o *GuestLoginUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /{baseSiteId}/users/{userId}/carts/{cartId}/email][%d] guestLoginUsingPUTCreated ", 201)
}

func (o *GuestLoginUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGuestLoginUsingPUTUnauthorized creates a GuestLoginUsingPUTUnauthorized with default headers values
func NewGuestLoginUsingPUTUnauthorized() *GuestLoginUsingPUTUnauthorized {
	return &GuestLoginUsingPUTUnauthorized{}
}

/*GuestLoginUsingPUTUnauthorized handles this case with default header values.

Unauthorized
*/
type GuestLoginUsingPUTUnauthorized struct {
}

func (o *GuestLoginUsingPUTUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /{baseSiteId}/users/{userId}/carts/{cartId}/email][%d] guestLoginUsingPUTUnauthorized ", 401)
}

func (o *GuestLoginUsingPUTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGuestLoginUsingPUTForbidden creates a GuestLoginUsingPUTForbidden with default headers values
func NewGuestLoginUsingPUTForbidden() *GuestLoginUsingPUTForbidden {
	return &GuestLoginUsingPUTForbidden{}
}

/*GuestLoginUsingPUTForbidden handles this case with default header values.

Forbidden
*/
type GuestLoginUsingPUTForbidden struct {
}

func (o *GuestLoginUsingPUTForbidden) Error() string {
	return fmt.Sprintf("[PUT /{baseSiteId}/users/{userId}/carts/{cartId}/email][%d] guestLoginUsingPUTForbidden ", 403)
}

func (o *GuestLoginUsingPUTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGuestLoginUsingPUTNotFound creates a GuestLoginUsingPUTNotFound with default headers values
func NewGuestLoginUsingPUTNotFound() *GuestLoginUsingPUTNotFound {
	return &GuestLoginUsingPUTNotFound{}
}

/*GuestLoginUsingPUTNotFound handles this case with default header values.

Not Found
*/
type GuestLoginUsingPUTNotFound struct {
}

func (o *GuestLoginUsingPUTNotFound) Error() string {
	return fmt.Sprintf("[PUT /{baseSiteId}/users/{userId}/carts/{cartId}/email][%d] guestLoginUsingPUTNotFound ", 404)
}

func (o *GuestLoginUsingPUTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
