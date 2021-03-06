// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

// CountSearchProductStockByLocationUsingHEADReader is a Reader for the CountSearchProductStockByLocationUsingHEAD structure.
type CountSearchProductStockByLocationUsingHEADReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CountSearchProductStockByLocationUsingHEADReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCountSearchProductStockByLocationUsingHEADOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewCountSearchProductStockByLocationUsingHEADNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewCountSearchProductStockByLocationUsingHEADUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCountSearchProductStockByLocationUsingHEADForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCountSearchProductStockByLocationUsingHEADOK creates a CountSearchProductStockByLocationUsingHEADOK with default headers values
func NewCountSearchProductStockByLocationUsingHEADOK() *CountSearchProductStockByLocationUsingHEADOK {
	return &CountSearchProductStockByLocationUsingHEADOK{}
}

/*CountSearchProductStockByLocationUsingHEADOK handles this case with default header values.

OK
*/
type CountSearchProductStockByLocationUsingHEADOK struct {
}

func (o *CountSearchProductStockByLocationUsingHEADOK) Error() string {
	return fmt.Sprintf("[HEAD /{baseSiteId}/products/{productCode}/stock][%d] countSearchProductStockByLocationUsingHEADOK ", 200)
}

func (o *CountSearchProductStockByLocationUsingHEADOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCountSearchProductStockByLocationUsingHEADNoContent creates a CountSearchProductStockByLocationUsingHEADNoContent with default headers values
func NewCountSearchProductStockByLocationUsingHEADNoContent() *CountSearchProductStockByLocationUsingHEADNoContent {
	return &CountSearchProductStockByLocationUsingHEADNoContent{}
}

/*CountSearchProductStockByLocationUsingHEADNoContent handles this case with default header values.

No Content
*/
type CountSearchProductStockByLocationUsingHEADNoContent struct {
}

func (o *CountSearchProductStockByLocationUsingHEADNoContent) Error() string {
	return fmt.Sprintf("[HEAD /{baseSiteId}/products/{productCode}/stock][%d] countSearchProductStockByLocationUsingHEADNoContent ", 204)
}

func (o *CountSearchProductStockByLocationUsingHEADNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCountSearchProductStockByLocationUsingHEADUnauthorized creates a CountSearchProductStockByLocationUsingHEADUnauthorized with default headers values
func NewCountSearchProductStockByLocationUsingHEADUnauthorized() *CountSearchProductStockByLocationUsingHEADUnauthorized {
	return &CountSearchProductStockByLocationUsingHEADUnauthorized{}
}

/*CountSearchProductStockByLocationUsingHEADUnauthorized handles this case with default header values.

Unauthorized
*/
type CountSearchProductStockByLocationUsingHEADUnauthorized struct {
}

func (o *CountSearchProductStockByLocationUsingHEADUnauthorized) Error() string {
	return fmt.Sprintf("[HEAD /{baseSiteId}/products/{productCode}/stock][%d] countSearchProductStockByLocationUsingHEADUnauthorized ", 401)
}

func (o *CountSearchProductStockByLocationUsingHEADUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCountSearchProductStockByLocationUsingHEADForbidden creates a CountSearchProductStockByLocationUsingHEADForbidden with default headers values
func NewCountSearchProductStockByLocationUsingHEADForbidden() *CountSearchProductStockByLocationUsingHEADForbidden {
	return &CountSearchProductStockByLocationUsingHEADForbidden{}
}

/*CountSearchProductStockByLocationUsingHEADForbidden handles this case with default header values.

Forbidden
*/
type CountSearchProductStockByLocationUsingHEADForbidden struct {
}

func (o *CountSearchProductStockByLocationUsingHEADForbidden) Error() string {
	return fmt.Sprintf("[HEAD /{baseSiteId}/products/{productCode}/stock][%d] countSearchProductStockByLocationUsingHEADForbidden ", 403)
}

func (o *CountSearchProductStockByLocationUsingHEADForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
