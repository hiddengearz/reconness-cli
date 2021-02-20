// Code generated by go-swagger; DO NOT EDIT.

package subdomains

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutAPISubdomainsLabelIDReader is a Reader for the PutAPISubdomainsLabelID structure.
type PutAPISubdomainsLabelIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAPISubdomainsLabelIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutAPISubdomainsLabelIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutAPISubdomainsLabelIDOK creates a PutAPISubdomainsLabelIDOK with default headers values
func NewPutAPISubdomainsLabelIDOK() *PutAPISubdomainsLabelIDOK {
	return &PutAPISubdomainsLabelIDOK{}
}

/* PutAPISubdomainsLabelIDOK describes a response with status code 200, with default header values.

Success
*/
type PutAPISubdomainsLabelIDOK struct {
}

func (o *PutAPISubdomainsLabelIDOK) Error() string {
	return fmt.Sprintf("[PUT /api/Subdomains/label/{id}][%d] putApiSubdomainsLabelIdOK ", 200)
}

func (o *PutAPISubdomainsLabelIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
