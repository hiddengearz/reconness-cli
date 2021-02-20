// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetAPIAccountsCurrentVersionReader is a Reader for the GetAPIAccountsCurrentVersion structure.
type GetAPIAccountsCurrentVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIAccountsCurrentVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIAccountsCurrentVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIAccountsCurrentVersionOK creates a GetAPIAccountsCurrentVersionOK with default headers values
func NewGetAPIAccountsCurrentVersionOK() *GetAPIAccountsCurrentVersionOK {
	return &GetAPIAccountsCurrentVersionOK{}
}

/* GetAPIAccountsCurrentVersionOK describes a response with status code 200, with default header values.

Success
*/
type GetAPIAccountsCurrentVersionOK struct {
}

func (o *GetAPIAccountsCurrentVersionOK) Error() string {
	return fmt.Sprintf("[GET /api/Accounts/currentVersion][%d] getApiAccountsCurrentVersionOK ", 200)
}

func (o *GetAPIAccountsCurrentVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
