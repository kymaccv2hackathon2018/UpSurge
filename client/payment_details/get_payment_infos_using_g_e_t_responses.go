// Code generated by go-swagger; DO NOT EDIT.

package payment_details

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/gopdf/models"
)

// GetPaymentInfosUsingGETReader is a Reader for the GetPaymentInfosUsingGET structure.
type GetPaymentInfosUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentInfosUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentInfosUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetPaymentInfosUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetPaymentInfosUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetPaymentInfosUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentInfosUsingGETOK creates a GetPaymentInfosUsingGETOK with default headers values
func NewGetPaymentInfosUsingGETOK() *GetPaymentInfosUsingGETOK {
	return &GetPaymentInfosUsingGETOK{}
}

/*GetPaymentInfosUsingGETOK handles this case with default header values.

OK
*/
type GetPaymentInfosUsingGETOK struct {
	Payload *models.PaymentDetailsListWsDTO
}

func (o *GetPaymentInfosUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/users/{userId}/paymentdetails][%d] getPaymentInfosUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetPaymentInfosUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentDetailsListWsDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentInfosUsingGETUnauthorized creates a GetPaymentInfosUsingGETUnauthorized with default headers values
func NewGetPaymentInfosUsingGETUnauthorized() *GetPaymentInfosUsingGETUnauthorized {
	return &GetPaymentInfosUsingGETUnauthorized{}
}

/*GetPaymentInfosUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetPaymentInfosUsingGETUnauthorized struct {
}

func (o *GetPaymentInfosUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/users/{userId}/paymentdetails][%d] getPaymentInfosUsingGETUnauthorized ", 401)
}

func (o *GetPaymentInfosUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPaymentInfosUsingGETForbidden creates a GetPaymentInfosUsingGETForbidden with default headers values
func NewGetPaymentInfosUsingGETForbidden() *GetPaymentInfosUsingGETForbidden {
	return &GetPaymentInfosUsingGETForbidden{}
}

/*GetPaymentInfosUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetPaymentInfosUsingGETForbidden struct {
}

func (o *GetPaymentInfosUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/users/{userId}/paymentdetails][%d] getPaymentInfosUsingGETForbidden ", 403)
}

func (o *GetPaymentInfosUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPaymentInfosUsingGETNotFound creates a GetPaymentInfosUsingGETNotFound with default headers values
func NewGetPaymentInfosUsingGETNotFound() *GetPaymentInfosUsingGETNotFound {
	return &GetPaymentInfosUsingGETNotFound{}
}

/*GetPaymentInfosUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetPaymentInfosUsingGETNotFound struct {
}

func (o *GetPaymentInfosUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/users/{userId}/paymentdetails][%d] getPaymentInfosUsingGETNotFound ", 404)
}

func (o *GetPaymentInfosUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
