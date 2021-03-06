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

// GetSupportedDeliveryModesUsingGETReader is a Reader for the GetSupportedDeliveryModesUsingGET structure.
type GetSupportedDeliveryModesUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSupportedDeliveryModesUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSupportedDeliveryModesUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetSupportedDeliveryModesUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetSupportedDeliveryModesUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetSupportedDeliveryModesUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSupportedDeliveryModesUsingGETOK creates a GetSupportedDeliveryModesUsingGETOK with default headers values
func NewGetSupportedDeliveryModesUsingGETOK() *GetSupportedDeliveryModesUsingGETOK {
	return &GetSupportedDeliveryModesUsingGETOK{}
}

/*GetSupportedDeliveryModesUsingGETOK handles this case with default header values.

OK
*/
type GetSupportedDeliveryModesUsingGETOK struct {
	Payload *models.DeliveryModeListWsDTO
}

func (o *GetSupportedDeliveryModesUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/users/{userId}/carts/{cartId}/deliverymodes][%d] getSupportedDeliveryModesUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetSupportedDeliveryModesUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeliveryModeListWsDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSupportedDeliveryModesUsingGETUnauthorized creates a GetSupportedDeliveryModesUsingGETUnauthorized with default headers values
func NewGetSupportedDeliveryModesUsingGETUnauthorized() *GetSupportedDeliveryModesUsingGETUnauthorized {
	return &GetSupportedDeliveryModesUsingGETUnauthorized{}
}

/*GetSupportedDeliveryModesUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetSupportedDeliveryModesUsingGETUnauthorized struct {
}

func (o *GetSupportedDeliveryModesUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/users/{userId}/carts/{cartId}/deliverymodes][%d] getSupportedDeliveryModesUsingGETUnauthorized ", 401)
}

func (o *GetSupportedDeliveryModesUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSupportedDeliveryModesUsingGETForbidden creates a GetSupportedDeliveryModesUsingGETForbidden with default headers values
func NewGetSupportedDeliveryModesUsingGETForbidden() *GetSupportedDeliveryModesUsingGETForbidden {
	return &GetSupportedDeliveryModesUsingGETForbidden{}
}

/*GetSupportedDeliveryModesUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetSupportedDeliveryModesUsingGETForbidden struct {
}

func (o *GetSupportedDeliveryModesUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/users/{userId}/carts/{cartId}/deliverymodes][%d] getSupportedDeliveryModesUsingGETForbidden ", 403)
}

func (o *GetSupportedDeliveryModesUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSupportedDeliveryModesUsingGETNotFound creates a GetSupportedDeliveryModesUsingGETNotFound with default headers values
func NewGetSupportedDeliveryModesUsingGETNotFound() *GetSupportedDeliveryModesUsingGETNotFound {
	return &GetSupportedDeliveryModesUsingGETNotFound{}
}

/*GetSupportedDeliveryModesUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetSupportedDeliveryModesUsingGETNotFound struct {
}

func (o *GetSupportedDeliveryModesUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/users/{userId}/carts/{cartId}/deliverymodes][%d] getSupportedDeliveryModesUsingGETNotFound ", 404)
}

func (o *GetSupportedDeliveryModesUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
