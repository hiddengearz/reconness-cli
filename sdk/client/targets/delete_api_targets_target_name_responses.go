// Code generated by go-swagger; DO NOT EDIT.

package targets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteAPITargetsTargetNameReader is a Reader for the DeleteAPITargetsTargetName structure.
type DeleteAPITargetsTargetNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPITargetsTargetNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAPITargetsTargetNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAPITargetsTargetNameOK creates a DeleteAPITargetsTargetNameOK with default headers values
func NewDeleteAPITargetsTargetNameOK() *DeleteAPITargetsTargetNameOK {
	return &DeleteAPITargetsTargetNameOK{}
}

/* DeleteAPITargetsTargetNameOK describes a response with status code 200, with default header values.

Success
*/
type DeleteAPITargetsTargetNameOK struct {
}

func (o *DeleteAPITargetsTargetNameOK) Error() string {
	return fmt.Sprintf("[DELETE /api/Targets/{targetName}][%d] deleteApiTargetsTargetNameOK ", 200)
}

func (o *DeleteAPITargetsTargetNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
