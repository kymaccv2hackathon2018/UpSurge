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

// AddPaymentDetailsUsingPOST1Reader is a Reader for the AddPaymentDetailsUsingPOST1 structure.
type AddPaymentDetailsUsingPOST1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddPaymentDetailsUsingPOST1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewAddPaymentDetailsUsingPOST1Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewAddPaymentDetailsUsingPOST1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewAddPaymentDetailsUsingPOST1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAddPaymentDetailsUsingPOST1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddPaymentDetailsUsingPOST1Created creates a AddPaymentDetailsUsingPOST1Created with default headers values
func NewAddPaymentDetailsUsingPOST1Created() *AddPaymentDetailsUsingPOST1Created {
	return &AddPaymentDetailsUsingPOST1Created{}
}

/*AddPaymentDetailsUsingPOST1Created handles this case with default header values.

Created
*/
type AddPaymentDetailsUsingPOST1Created struct {
	Payload *models.PaymentDetailsWsDTO
}

func (o *AddPaymentDetailsUsingPOST1Created) Error() string {
	return fmt.Sprintf("[POST /{baseSiteId}/users/{userId}/carts/{cartId}/paymentdetails][%d] addPaymentDetailsUsingPOST1Created  %+v", 201, o.Payload)
}

func (o *AddPaymentDetailsUsingPOST1Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentDetailsWsDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddPaymentDetailsUsingPOST1Unauthorized creates a AddPaymentDetailsUsingPOST1Unauthorized with default headers values
func NewAddPaymentDetailsUsingPOST1Unauthorized() *AddPaymentDetailsUsingPOST1Unauthorized {
	return &AddPaymentDetailsUsingPOST1Unauthorized{}
}

/*AddPaymentDetailsUsingPOST1Unauthorized handles this case with default header values.

Unauthorized
*/
type AddPaymentDetailsUsingPOST1Unauthorized struct {
}

func (o *AddPaymentDetailsUsingPOST1Unauthorized) Error() string {
	return fmt.Sprintf("[POST /{baseSiteId}/users/{userId}/carts/{cartId}/paymentdetails][%d] addPaymentDetailsUsingPOST1Unauthorized ", 401)
}

func (o *AddPaymentDetailsUsingPOST1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddPaymentDetailsUsingPOST1Forbidden creates a AddPaymentDetailsUsingPOST1Forbidden with default headers values
func NewAddPaymentDetailsUsingPOST1Forbidden() *AddPaymentDetailsUsingPOST1Forbidden {
	return &AddPaymentDetailsUsingPOST1Forbidden{}
}

/*AddPaymentDetailsUsingPOST1Forbidden handles this case with default header values.

Forbidden
*/
type AddPaymentDetailsUsingPOST1Forbidden struct {
}

func (o *AddPaymentDetailsUsingPOST1Forbidden) Error() string {
	return fmt.Sprintf("[POST /{baseSiteId}/users/{userId}/carts/{cartId}/paymentdetails][%d] addPaymentDetailsUsingPOST1Forbidden ", 403)
}

func (o *AddPaymentDetailsUsingPOST1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddPaymentDetailsUsingPOST1NotFound creates a AddPaymentDetailsUsingPOST1NotFound with default headers values
func NewAddPaymentDetailsUsingPOST1NotFound() *AddPaymentDetailsUsingPOST1NotFound {
	return &AddPaymentDetailsUsingPOST1NotFound{}
}

/*AddPaymentDetailsUsingPOST1NotFound handles this case with default header values.

Not Found
*/
type AddPaymentDetailsUsingPOST1NotFound struct {
}

func (o *AddPaymentDetailsUsingPOST1NotFound) Error() string {
	return fmt.Sprintf("[POST /{baseSiteId}/users/{userId}/carts/{cartId}/paymentdetails][%d] addPaymentDetailsUsingPOST1NotFound ", 404)
}

func (o *AddPaymentDetailsUsingPOST1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
