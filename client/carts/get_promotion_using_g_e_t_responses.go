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

// GetPromotionUsingGETReader is a Reader for the GetPromotionUsingGET structure.
type GetPromotionUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPromotionUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPromotionUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetPromotionUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetPromotionUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetPromotionUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPromotionUsingGETOK creates a GetPromotionUsingGETOK with default headers values
func NewGetPromotionUsingGETOK() *GetPromotionUsingGETOK {
	return &GetPromotionUsingGETOK{}
}

/*GetPromotionUsingGETOK handles this case with default header values.

OK
*/
type GetPromotionUsingGETOK struct {
	Payload *models.PromotionResultListWsDTO
}

func (o *GetPromotionUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/users/{userId}/carts/{cartId}/promotions/{promotionId}][%d] getPromotionUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetPromotionUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PromotionResultListWsDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPromotionUsingGETUnauthorized creates a GetPromotionUsingGETUnauthorized with default headers values
func NewGetPromotionUsingGETUnauthorized() *GetPromotionUsingGETUnauthorized {
	return &GetPromotionUsingGETUnauthorized{}
}

/*GetPromotionUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetPromotionUsingGETUnauthorized struct {
}

func (o *GetPromotionUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/users/{userId}/carts/{cartId}/promotions/{promotionId}][%d] getPromotionUsingGETUnauthorized ", 401)
}

func (o *GetPromotionUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPromotionUsingGETForbidden creates a GetPromotionUsingGETForbidden with default headers values
func NewGetPromotionUsingGETForbidden() *GetPromotionUsingGETForbidden {
	return &GetPromotionUsingGETForbidden{}
}

/*GetPromotionUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetPromotionUsingGETForbidden struct {
}

func (o *GetPromotionUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/users/{userId}/carts/{cartId}/promotions/{promotionId}][%d] getPromotionUsingGETForbidden ", 403)
}

func (o *GetPromotionUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPromotionUsingGETNotFound creates a GetPromotionUsingGETNotFound with default headers values
func NewGetPromotionUsingGETNotFound() *GetPromotionUsingGETNotFound {
	return &GetPromotionUsingGETNotFound{}
}

/*GetPromotionUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetPromotionUsingGETNotFound struct {
}

func (o *GetPromotionUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/users/{userId}/carts/{cartId}/promotions/{promotionId}][%d] getPromotionUsingGETNotFound ", 404)
}

func (o *GetPromotionUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}