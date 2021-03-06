// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new products API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for products API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CountSearchProductStockByLocationUsingHEAD gets header with a total number of product s stock levels

In the response header, the "x-total-count" indicates the total number of a product's stock levels. The following two sets of parameters are available: location (required); or longitude (required), and latitude (required).
*/
func (a *Client) CountSearchProductStockByLocationUsingHEAD(params *CountSearchProductStockByLocationUsingHEADParams, authInfo runtime.ClientAuthInfoWriter) (*CountSearchProductStockByLocationUsingHEADOK, *CountSearchProductStockByLocationUsingHEADNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCountSearchProductStockByLocationUsingHEADParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "countSearchProductStockByLocationUsingHEAD",
		Method:             "HEAD",
		PathPattern:        "/{baseSiteId}/products/{productCode}/stock",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CountSearchProductStockByLocationUsingHEADReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CountSearchProductStockByLocationUsingHEADOK:
		return value, nil, nil
	case *CountSearchProductStockByLocationUsingHEADNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
CountSearchProductsUsingHEAD gets a header with total number of products

In the response header, the "x-total-count" indicates the total number of products satisfying a query.
*/
func (a *Client) CountSearchProductsUsingHEAD(params *CountSearchProductsUsingHEADParams, authInfo runtime.ClientAuthInfoWriter) (*CountSearchProductsUsingHEADOK, *CountSearchProductsUsingHEADNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCountSearchProductsUsingHEADParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "countSearchProductsUsingHEAD",
		Method:             "HEAD",
		PathPattern:        "/{baseSiteId}/products/search",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CountSearchProductsUsingHEADReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CountSearchProductsUsingHEADOK:
		return value, nil, nil
	case *CountSearchProductsUsingHEADNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
CreateReviewUsingPOST1 creates a new customer review as an anonymous user

Creates a new customer review as an anonymous user.
*/
func (a *Client) CreateReviewUsingPOST1(params *CreateReviewUsingPOST1Params, authInfo runtime.ClientAuthInfoWriter) (*CreateReviewUsingPOST1Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateReviewUsingPOST1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createReviewUsingPOST_1",
		Method:             "POST",
		PathPattern:        "/{baseSiteId}/products/{productCode}/reviews",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateReviewUsingPOST1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateReviewUsingPOST1Created), nil

}

/*
ExportProductReferencesUsingGET gets a product reference

Returns references for a product with a given product code. Reference type specifies which references to return.
*/
func (a *Client) ExportProductReferencesUsingGET(params *ExportProductReferencesUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*ExportProductReferencesUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExportProductReferencesUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "exportProductReferencesUsingGET",
		Method:             "GET",
		PathPattern:        "/{baseSiteId}/products/{productCode}/references",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ExportProductReferencesUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ExportProductReferencesUsingGETOK), nil

}

/*
ExpressUpdateUsingGET gets products added to the express update feed

Returns products added to the express update feed. Returns only elements updated after the provided timestamp. The queue is cleared using a defined cronjob.
*/
func (a *Client) ExpressUpdateUsingGET(params *ExpressUpdateUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*ExpressUpdateUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExpressUpdateUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "expressUpdateUsingGET",
		Method:             "GET",
		PathPattern:        "/{baseSiteId}/products/expressupdate",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ExpressUpdateUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ExpressUpdateUsingGETOK), nil

}

/*
GetProductByCodeUsingGET gets product details

Returns details of a single product according to a product code.
*/
func (a *Client) GetProductByCodeUsingGET(params *GetProductByCodeUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetProductByCodeUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProductByCodeUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProductByCodeUsingGET",
		Method:             "GET",
		PathPattern:        "/{baseSiteId}/products/{productCode}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetProductByCodeUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetProductByCodeUsingGETOK), nil

}

/*
GetProductReviewsUsingGET gets reviews for a product

Returns the reviews for a product with a given product code.
*/
func (a *Client) GetProductReviewsUsingGET(params *GetProductReviewsUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetProductReviewsUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProductReviewsUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProductReviewsUsingGET",
		Method:             "GET",
		PathPattern:        "/{baseSiteId}/products/{productCode}/reviews",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetProductReviewsUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetProductReviewsUsingGETOK), nil

}

/*
GetStockDataUsingGET gets a product s stock level for a store

Returns a product's stock level for a particular store (in other words, for a particular point of sale).
*/
func (a *Client) GetStockDataUsingGET(params *GetStockDataUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetStockDataUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStockDataUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getStockDataUsingGET",
		Method:             "GET",
		PathPattern:        "/{baseSiteId}/products/{productCode}/stock/{storeName}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetStockDataUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStockDataUsingGETOK), nil

}

/*
GetSuggestionsUsingGET gets a list of available suggestions

Returns a list of all available suggestions related to a given term and limits the results to a specific value of the max parameter.
*/
func (a *Client) GetSuggestionsUsingGET(params *GetSuggestionsUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetSuggestionsUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSuggestionsUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSuggestionsUsingGET",
		Method:             "GET",
		PathPattern:        "/{baseSiteId}/products/suggestions",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSuggestionsUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSuggestionsUsingGETOK), nil

}

/*
SearchProductStockByLocationUsingGET gets a product s stock level

Returns a product's stock levels sorted by distance from the specified location, which is provided using the free-text "location" parameter, or by using the longitude and latitude parameters. The following two sets of parameters are available: location (required), currentPage (optional), pageSize (optional); or longitude (required), latitude (required), currentPage (optional), pageSize(optional).
*/
func (a *Client) SearchProductStockByLocationUsingGET(params *SearchProductStockByLocationUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*SearchProductStockByLocationUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchProductStockByLocationUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "searchProductStockByLocationUsingGET",
		Method:             "GET",
		PathPattern:        "/{baseSiteId}/products/{productCode}/stock",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SearchProductStockByLocationUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SearchProductStockByLocationUsingGETOK), nil

}

/*
SearchProductsUsingGET gets a list of products and additional data

Returns a list of products and additional data, such as available facets, available sorting, and pagination options. It can also include spelling suggestions. To make spelling suggestions work, you need to make sure that "enableSpellCheck" on the SearchQuery is set to "true" (by default, it should already be set to "true"). You also need to have indexed properties configured to be used for spellchecking.
*/
func (a *Client) SearchProductsUsingGET(params *SearchProductsUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*SearchProductsUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchProductsUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "searchProductsUsingGET",
		Method:             "GET",
		PathPattern:        "/{baseSiteId}/products/search",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SearchProductsUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SearchProductsUsingGETOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
