// Code generated by go-swagger; DO NOT EDIT.

package carts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

// RemoveDeliveryModeUsingDELETEReader is a Reader for the RemoveDeliveryModeUsingDELETE structure.
type RemoveDeliveryModeUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveDeliveryModeUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRemoveDeliveryModeUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewRemoveDeliveryModeUsingDELETENoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewRemoveDeliveryModeUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewRemoveDeliveryModeUsingDELETEForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveDeliveryModeUsingDELETEOK creates a RemoveDeliveryModeUsingDELETEOK with default headers values
func NewRemoveDeliveryModeUsingDELETEOK() *RemoveDeliveryModeUsingDELETEOK {
	return &RemoveDeliveryModeUsingDELETEOK{}
}

/*RemoveDeliveryModeUsingDELETEOK handles this case with default header values.

OK
*/
type RemoveDeliveryModeUsingDELETEOK struct {
}

func (o *RemoveDeliveryModeUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /{baseSiteId}/users/{userId}/carts/{cartId}/deliverymode][%d] removeDeliveryModeUsingDELETEOK ", 200)
}

func (o *RemoveDeliveryModeUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveDeliveryModeUsingDELETENoContent creates a RemoveDeliveryModeUsingDELETENoContent with default headers values
func NewRemoveDeliveryModeUsingDELETENoContent() *RemoveDeliveryModeUsingDELETENoContent {
	return &RemoveDeliveryModeUsingDELETENoContent{}
}

/*RemoveDeliveryModeUsingDELETENoContent handles this case with default header values.

No Content
*/
type RemoveDeliveryModeUsingDELETENoContent struct {
}

func (o *RemoveDeliveryModeUsingDELETENoContent) Error() string {
	return fmt.Sprintf("[DELETE /{baseSiteId}/users/{userId}/carts/{cartId}/deliverymode][%d] removeDeliveryModeUsingDELETENoContent ", 204)
}

func (o *RemoveDeliveryModeUsingDELETENoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveDeliveryModeUsingDELETEUnauthorized creates a RemoveDeliveryModeUsingDELETEUnauthorized with default headers values
func NewRemoveDeliveryModeUsingDELETEUnauthorized() *RemoveDeliveryModeUsingDELETEUnauthorized {
	return &RemoveDeliveryModeUsingDELETEUnauthorized{}
}

/*RemoveDeliveryModeUsingDELETEUnauthorized handles this case with default header values.

Unauthorized
*/
type RemoveDeliveryModeUsingDELETEUnauthorized struct {
}

func (o *RemoveDeliveryModeUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /{baseSiteId}/users/{userId}/carts/{cartId}/deliverymode][%d] removeDeliveryModeUsingDELETEUnauthorized ", 401)
}

func (o *RemoveDeliveryModeUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveDeliveryModeUsingDELETEForbidden creates a RemoveDeliveryModeUsingDELETEForbidden with default headers values
func NewRemoveDeliveryModeUsingDELETEForbidden() *RemoveDeliveryModeUsingDELETEForbidden {
	return &RemoveDeliveryModeUsingDELETEForbidden{}
}

/*RemoveDeliveryModeUsingDELETEForbidden handles this case with default header values.

Forbidden
*/
type RemoveDeliveryModeUsingDELETEForbidden struct {
}

func (o *RemoveDeliveryModeUsingDELETEForbidden) Error() string {
	return fmt.Sprintf("[DELETE /{baseSiteId}/users/{userId}/carts/{cartId}/deliverymode][%d] removeDeliveryModeUsingDELETEForbidden ", 403)
}

func (o *RemoveDeliveryModeUsingDELETEForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
