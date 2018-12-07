// Code generated by go-swagger; DO NOT EDIT.

package export

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/gopdf/models"
)

// ExportProductsUsingGETReader is a Reader for the ExportProductsUsingGET structure.
type ExportProductsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportProductsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewExportProductsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewExportProductsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewExportProductsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewExportProductsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewExportProductsUsingGETOK creates a ExportProductsUsingGETOK with default headers values
func NewExportProductsUsingGETOK() *ExportProductsUsingGETOK {
	return &ExportProductsUsingGETOK{}
}

/*ExportProductsUsingGETOK handles this case with default header values.

OK
*/
type ExportProductsUsingGETOK struct {
	Payload *models.ProductListWsDTO
}

func (o *ExportProductsUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/export/products][%d] exportProductsUsingGETOK  %+v", 200, o.Payload)
}

func (o *ExportProductsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProductListWsDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportProductsUsingGETUnauthorized creates a ExportProductsUsingGETUnauthorized with default headers values
func NewExportProductsUsingGETUnauthorized() *ExportProductsUsingGETUnauthorized {
	return &ExportProductsUsingGETUnauthorized{}
}

/*ExportProductsUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type ExportProductsUsingGETUnauthorized struct {
}

func (o *ExportProductsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/export/products][%d] exportProductsUsingGETUnauthorized ", 401)
}

func (o *ExportProductsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportProductsUsingGETForbidden creates a ExportProductsUsingGETForbidden with default headers values
func NewExportProductsUsingGETForbidden() *ExportProductsUsingGETForbidden {
	return &ExportProductsUsingGETForbidden{}
}

/*ExportProductsUsingGETForbidden handles this case with default header values.

Forbidden
*/
type ExportProductsUsingGETForbidden struct {
}

func (o *ExportProductsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/export/products][%d] exportProductsUsingGETForbidden ", 403)
}

func (o *ExportProductsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportProductsUsingGETNotFound creates a ExportProductsUsingGETNotFound with default headers values
func NewExportProductsUsingGETNotFound() *ExportProductsUsingGETNotFound {
	return &ExportProductsUsingGETNotFound{}
}

/*ExportProductsUsingGETNotFound handles this case with default header values.

Not Found
*/
type ExportProductsUsingGETNotFound struct {
}

func (o *ExportProductsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/export/products][%d] exportProductsUsingGETNotFound ", 404)
}

func (o *ExportProductsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
