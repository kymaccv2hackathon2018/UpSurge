// Code generated by go-swagger; DO NOT EDIT.

package carts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/gopdf/models"
)

// UpdateCartEntryUsingPATCH1Reader is a Reader for the UpdateCartEntryUsingPATCH1 structure.
type UpdateCartEntryUsingPATCH1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCartEntryUsingPATCH1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateCartEntryUsingPATCH1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewUpdateCartEntryUsingPATCH1NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewUpdateCartEntryUsingPATCH1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUpdateCartEntryUsingPATCH1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateCartEntryUsingPATCH1OK creates a UpdateCartEntryUsingPATCH1OK with default headers values
func NewUpdateCartEntryUsingPATCH1OK() *UpdateCartEntryUsingPATCH1OK {
	return &UpdateCartEntryUsingPATCH1OK{}
}

/*UpdateCartEntryUsingPATCH1OK handles this case with default header values.

OK
*/
type UpdateCartEntryUsingPATCH1OK struct {
	Payload *models.CartModificationWsDTO
}

func (o *UpdateCartEntryUsingPATCH1OK) Error() string {
	return fmt.Sprintf("[PATCH /{baseSiteId}/users/{userId}/carts/{cartId}/entries/{entryNumber}][%d] updateCartEntryUsingPATCH1OK  %+v", 200, o.Payload)
}

func (o *UpdateCartEntryUsingPATCH1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CartModificationWsDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCartEntryUsingPATCH1NoContent creates a UpdateCartEntryUsingPATCH1NoContent with default headers values
func NewUpdateCartEntryUsingPATCH1NoContent() *UpdateCartEntryUsingPATCH1NoContent {
	return &UpdateCartEntryUsingPATCH1NoContent{}
}

/*UpdateCartEntryUsingPATCH1NoContent handles this case with default header values.

No Content
*/
type UpdateCartEntryUsingPATCH1NoContent struct {
}

func (o *UpdateCartEntryUsingPATCH1NoContent) Error() string {
	return fmt.Sprintf("[PATCH /{baseSiteId}/users/{userId}/carts/{cartId}/entries/{entryNumber}][%d] updateCartEntryUsingPATCH1NoContent ", 204)
}

func (o *UpdateCartEntryUsingPATCH1NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateCartEntryUsingPATCH1Unauthorized creates a UpdateCartEntryUsingPATCH1Unauthorized with default headers values
func NewUpdateCartEntryUsingPATCH1Unauthorized() *UpdateCartEntryUsingPATCH1Unauthorized {
	return &UpdateCartEntryUsingPATCH1Unauthorized{}
}

/*UpdateCartEntryUsingPATCH1Unauthorized handles this case with default header values.

Unauthorized
*/
type UpdateCartEntryUsingPATCH1Unauthorized struct {
}

func (o *UpdateCartEntryUsingPATCH1Unauthorized) Error() string {
	return fmt.Sprintf("[PATCH /{baseSiteId}/users/{userId}/carts/{cartId}/entries/{entryNumber}][%d] updateCartEntryUsingPATCH1Unauthorized ", 401)
}

func (o *UpdateCartEntryUsingPATCH1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateCartEntryUsingPATCH1Forbidden creates a UpdateCartEntryUsingPATCH1Forbidden with default headers values
func NewUpdateCartEntryUsingPATCH1Forbidden() *UpdateCartEntryUsingPATCH1Forbidden {
	return &UpdateCartEntryUsingPATCH1Forbidden{}
}

/*UpdateCartEntryUsingPATCH1Forbidden handles this case with default header values.

Forbidden
*/
type UpdateCartEntryUsingPATCH1Forbidden struct {
}

func (o *UpdateCartEntryUsingPATCH1Forbidden) Error() string {
	return fmt.Sprintf("[PATCH /{baseSiteId}/users/{userId}/carts/{cartId}/entries/{entryNumber}][%d] updateCartEntryUsingPATCH1Forbidden ", 403)
}

func (o *UpdateCartEntryUsingPATCH1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
