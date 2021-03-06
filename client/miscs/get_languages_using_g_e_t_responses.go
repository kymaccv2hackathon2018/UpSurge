// Code generated by go-swagger; DO NOT EDIT.

package miscs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/gopdf/models"
)

// GetLanguagesUsingGETReader is a Reader for the GetLanguagesUsingGET structure.
type GetLanguagesUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLanguagesUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLanguagesUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetLanguagesUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetLanguagesUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetLanguagesUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLanguagesUsingGETOK creates a GetLanguagesUsingGETOK with default headers values
func NewGetLanguagesUsingGETOK() *GetLanguagesUsingGETOK {
	return &GetLanguagesUsingGETOK{}
}

/*GetLanguagesUsingGETOK handles this case with default header values.

OK
*/
type GetLanguagesUsingGETOK struct {
	Payload *models.LanguageListWsDTO
}

func (o *GetLanguagesUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/languages][%d] getLanguagesUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetLanguagesUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LanguageListWsDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesUsingGETUnauthorized creates a GetLanguagesUsingGETUnauthorized with default headers values
func NewGetLanguagesUsingGETUnauthorized() *GetLanguagesUsingGETUnauthorized {
	return &GetLanguagesUsingGETUnauthorized{}
}

/*GetLanguagesUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetLanguagesUsingGETUnauthorized struct {
}

func (o *GetLanguagesUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/languages][%d] getLanguagesUsingGETUnauthorized ", 401)
}

func (o *GetLanguagesUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLanguagesUsingGETForbidden creates a GetLanguagesUsingGETForbidden with default headers values
func NewGetLanguagesUsingGETForbidden() *GetLanguagesUsingGETForbidden {
	return &GetLanguagesUsingGETForbidden{}
}

/*GetLanguagesUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetLanguagesUsingGETForbidden struct {
}

func (o *GetLanguagesUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/languages][%d] getLanguagesUsingGETForbidden ", 403)
}

func (o *GetLanguagesUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLanguagesUsingGETNotFound creates a GetLanguagesUsingGETNotFound with default headers values
func NewGetLanguagesUsingGETNotFound() *GetLanguagesUsingGETNotFound {
	return &GetLanguagesUsingGETNotFound{}
}

/*GetLanguagesUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetLanguagesUsingGETNotFound struct {
}

func (o *GetLanguagesUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/languages][%d] getLanguagesUsingGETNotFound ", 404)
}

func (o *GetLanguagesUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
