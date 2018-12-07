// Code generated by go-swagger; DO NOT EDIT.

package stores

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new stores API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for stores API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CountLocationSearchUsingHEAD gets a header with the number of store locations

In the response header, the "x-total-count" indicates the number of all store locations that are near the location specified in a query, or based on latitude and longitude.
*/
func (a *Client) CountLocationSearchUsingHEAD(params *CountLocationSearchUsingHEADParams, authInfo runtime.ClientAuthInfoWriter) (*CountLocationSearchUsingHEADOK, *CountLocationSearchUsingHEADNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCountLocationSearchUsingHEADParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "countLocationSearchUsingHEAD",
		Method:             "HEAD",
		PathPattern:        "/{baseSiteId}/stores",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CountLocationSearchUsingHEADReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CountLocationSearchUsingHEADOK:
		return value, nil, nil
	case *CountLocationSearchUsingHEADNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
LocationDetailsUsingGET gets a store location

Returns store location based on its unique name.
*/
func (a *Client) LocationDetailsUsingGET(params *LocationDetailsUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*LocationDetailsUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLocationDetailsUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "locationDetailsUsingGET",
		Method:             "GET",
		PathPattern:        "/{baseSiteId}/stores/{storeId}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LocationDetailsUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LocationDetailsUsingGETOK), nil

}

/*
LocationSearchUsingGET gets a list of store locations

Lists all store locations that are near the location specified in a query or based on latitude and longitude.
*/
func (a *Client) LocationSearchUsingGET(params *LocationSearchUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*LocationSearchUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLocationSearchUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "locationSearchUsingGET",
		Method:             "GET",
		PathPattern:        "/{baseSiteId}/stores",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LocationSearchUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LocationSearchUsingGETOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
