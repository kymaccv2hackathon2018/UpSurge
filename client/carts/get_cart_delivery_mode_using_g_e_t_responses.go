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

// GetCartDeliveryModeUsingGETReader is a Reader for the GetCartDeliveryModeUsingGET structure.
type GetCartDeliveryModeUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCartDeliveryModeUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCartDeliveryModeUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetCartDeliveryModeUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetCartDeliveryModeUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetCartDeliveryModeUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCartDeliveryModeUsingGETOK creates a GetCartDeliveryModeUsingGETOK with default headers values
func NewGetCartDeliveryModeUsingGETOK() *GetCartDeliveryModeUsingGETOK {
	return &GetCartDeliveryModeUsingGETOK{}
}

/*GetCartDeliveryModeUsingGETOK handles this case with default header values.

OK
*/
type GetCartDeliveryModeUsingGETOK struct {
	Payload *models.DeliveryModeWsDTO
}

func (o *GetCartDeliveryModeUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/users/{userId}/carts/{cartId}/deliverymode][%d] getCartDeliveryModeUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetCartDeliveryModeUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeliveryModeWsDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCartDeliveryModeUsingGETUnauthorized creates a GetCartDeliveryModeUsingGETUnauthorized with default headers values
func NewGetCartDeliveryModeUsingGETUnauthorized() *GetCartDeliveryModeUsingGETUnauthorized {
	return &GetCartDeliveryModeUsingGETUnauthorized{}
}

/*GetCartDeliveryModeUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCartDeliveryModeUsingGETUnauthorized struct {
}

func (o *GetCartDeliveryModeUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/users/{userId}/carts/{cartId}/deliverymode][%d] getCartDeliveryModeUsingGETUnauthorized ", 401)
}

func (o *GetCartDeliveryModeUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCartDeliveryModeUsingGETForbidden creates a GetCartDeliveryModeUsingGETForbidden with default headers values
func NewGetCartDeliveryModeUsingGETForbidden() *GetCartDeliveryModeUsingGETForbidden {
	return &GetCartDeliveryModeUsingGETForbidden{}
}

/*GetCartDeliveryModeUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetCartDeliveryModeUsingGETForbidden struct {
}

func (o *GetCartDeliveryModeUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/users/{userId}/carts/{cartId}/deliverymode][%d] getCartDeliveryModeUsingGETForbidden ", 403)
}

func (o *GetCartDeliveryModeUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCartDeliveryModeUsingGETNotFound creates a GetCartDeliveryModeUsingGETNotFound with default headers values
func NewGetCartDeliveryModeUsingGETNotFound() *GetCartDeliveryModeUsingGETNotFound {
	return &GetCartDeliveryModeUsingGETNotFound{}
}

/*GetCartDeliveryModeUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetCartDeliveryModeUsingGETNotFound struct {
}

func (o *GetCartDeliveryModeUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/users/{userId}/carts/{cartId}/deliverymode][%d] getCartDeliveryModeUsingGETNotFound ", 404)
}

func (o *GetCartDeliveryModeUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
