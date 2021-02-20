// Code generated by go-swagger; DO NOT EDIT.

package references

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteAPIReferencesIDReader is a Reader for the DeleteAPIReferencesID structure.
type DeleteAPIReferencesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPIReferencesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAPIReferencesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAPIReferencesIDOK creates a DeleteAPIReferencesIDOK with default headers values
func NewDeleteAPIReferencesIDOK() *DeleteAPIReferencesIDOK {
	return &DeleteAPIReferencesIDOK{}
}

/* DeleteAPIReferencesIDOK describes a response with status code 200, with default header values.

Success
*/
type DeleteAPIReferencesIDOK struct {
}

func (o *DeleteAPIReferencesIDOK) Error() string {
	return fmt.Sprintf("[DELETE /api/References/{id}][%d] deleteApiReferencesIdOK ", 200)
}

func (o *DeleteAPIReferencesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
