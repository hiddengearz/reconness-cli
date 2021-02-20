// Code generated by go-swagger; DO NOT EDIT.

package targets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new targets API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for targets API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteAPITargetsTargetName(params *DeleteAPITargetsTargetNameParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteAPITargetsTargetNameOK, error)

	GetAPITargets(params *GetAPITargetsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAPITargetsOK, error)

	GetAPITargetsTargetName(params *GetAPITargetsTargetNameParams, authInfo runtime.ClientAuthInfoWriter) (*GetAPITargetsTargetNameOK, error)

	PostAPITargets(params *PostAPITargetsParams, authInfo runtime.ClientAuthInfoWriter) (*PostAPITargetsOK, error)

	PostAPITargetsImportRootDomainTargetName(params *PostAPITargetsImportRootDomainTargetNameParams, authInfo runtime.ClientAuthInfoWriter) (*PostAPITargetsImportRootDomainTargetNameOK, error)

	PutAPITargetsID(params *PutAPITargetsIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutAPITargetsIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteAPITargetsTargetName delete API targets target name API
*/
func (a *Client) DeleteAPITargetsTargetName(params *DeleteAPITargetsTargetNameParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteAPITargetsTargetNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAPITargetsTargetNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteAPITargetsTargetName",
		Method:             "DELETE",
		PathPattern:        "/api/Targets/{targetName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteAPITargetsTargetNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteAPITargetsTargetNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteAPITargetsTargetName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPITargets get API targets API
*/
func (a *Client) GetAPITargets(params *GetAPITargetsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAPITargetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPITargetsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAPITargets",
		Method:             "GET",
		PathPattern:        "/api/Targets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPITargetsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAPITargetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPITargets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPITargetsTargetName get API targets target name API
*/
func (a *Client) GetAPITargetsTargetName(params *GetAPITargetsTargetNameParams, authInfo runtime.ClientAuthInfoWriter) (*GetAPITargetsTargetNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPITargetsTargetNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAPITargetsTargetName",
		Method:             "GET",
		PathPattern:        "/api/Targets/{targetName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPITargetsTargetNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAPITargetsTargetNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPITargetsTargetName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostAPITargets post API targets API
*/
func (a *Client) PostAPITargets(params *PostAPITargetsParams, authInfo runtime.ClientAuthInfoWriter) (*PostAPITargetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPITargetsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostAPITargets",
		Method:             "POST",
		PathPattern:        "/api/Targets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostAPITargetsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostAPITargetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAPITargets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostAPITargetsImportRootDomainTargetName post API targets import root domain target name API
*/
func (a *Client) PostAPITargetsImportRootDomainTargetName(params *PostAPITargetsImportRootDomainTargetNameParams, authInfo runtime.ClientAuthInfoWriter) (*PostAPITargetsImportRootDomainTargetNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPITargetsImportRootDomainTargetNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostAPITargetsImportRootDomainTargetName",
		Method:             "POST",
		PathPattern:        "/api/Targets/importRootDomain/{targetName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostAPITargetsImportRootDomainTargetNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostAPITargetsImportRootDomainTargetNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAPITargetsImportRootDomainTargetName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutAPITargetsID put API targets ID API
*/
func (a *Client) PutAPITargetsID(params *PutAPITargetsIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutAPITargetsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutAPITargetsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutAPITargetsID",
		Method:             "PUT",
		PathPattern:        "/api/Targets/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutAPITargetsIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutAPITargetsIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutAPITargetsID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
