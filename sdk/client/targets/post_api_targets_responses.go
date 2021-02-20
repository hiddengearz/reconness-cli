// Code generated by go-swagger; DO NOT EDIT.

package targets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostAPITargetsReader is a Reader for the PostAPITargets structure.
type PostAPITargetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPITargetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAPITargetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAPITargetsOK creates a PostAPITargetsOK with default headers values
func NewPostAPITargetsOK() *PostAPITargetsOK {
	return &PostAPITargetsOK{}
}

/* PostAPITargetsOK describes a response with status code 200, with default header values.

Success
*/
type PostAPITargetsOK struct {
}

func (o *PostAPITargetsOK) Error() string {
	return fmt.Sprintf("[POST /api/Targets][%d] postApiTargetsOK ", 200)
}

func (o *PostAPITargetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}