// Code generated by go-swagger; DO NOT EDIT.

package component

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/gopdf/models"
)

// GetComponentByIDUsingGETReader is a Reader for the GetComponentByIDUsingGET structure.
type GetComponentByIDUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetComponentByIDUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetComponentByIDUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetComponentByIDUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetComponentByIDUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetComponentByIDUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetComponentByIDUsingGETOK creates a GetComponentByIDUsingGETOK with default headers values
func NewGetComponentByIDUsingGETOK() *GetComponentByIDUsingGETOK {
	return &GetComponentByIDUsingGETOK{}
}

/*GetComponentByIDUsingGETOK handles this case with default header values.

OK
*/
type GetComponentByIDUsingGETOK struct {
	Payload models.ComponentAdaptedData
}

func (o *GetComponentByIDUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/cms/components/{componentId}][%d] getComponentByIdUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetComponentByIDUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetComponentByIDUsingGETUnauthorized creates a GetComponentByIDUsingGETUnauthorized with default headers values
func NewGetComponentByIDUsingGETUnauthorized() *GetComponentByIDUsingGETUnauthorized {
	return &GetComponentByIDUsingGETUnauthorized{}
}

/*GetComponentByIDUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetComponentByIDUsingGETUnauthorized struct {
}

func (o *GetComponentByIDUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/cms/components/{componentId}][%d] getComponentByIdUsingGETUnauthorized ", 401)
}

func (o *GetComponentByIDUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetComponentByIDUsingGETForbidden creates a GetComponentByIDUsingGETForbidden with default headers values
func NewGetComponentByIDUsingGETForbidden() *GetComponentByIDUsingGETForbidden {
	return &GetComponentByIDUsingGETForbidden{}
}

/*GetComponentByIDUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetComponentByIDUsingGETForbidden struct {
}

func (o *GetComponentByIDUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/cms/components/{componentId}][%d] getComponentByIdUsingGETForbidden ", 403)
}

func (o *GetComponentByIDUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetComponentByIDUsingGETNotFound creates a GetComponentByIDUsingGETNotFound with default headers values
func NewGetComponentByIDUsingGETNotFound() *GetComponentByIDUsingGETNotFound {
	return &GetComponentByIDUsingGETNotFound{}
}

/*GetComponentByIDUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetComponentByIDUsingGETNotFound struct {
}

func (o *GetComponentByIDUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/cms/components/{componentId}][%d] getComponentByIdUsingGETNotFound ", 404)
}

func (o *GetComponentByIDUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
