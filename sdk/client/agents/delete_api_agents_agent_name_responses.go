// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteAPIAgentsAgentNameReader is a Reader for the DeleteAPIAgentsAgentName structure.
type DeleteAPIAgentsAgentNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPIAgentsAgentNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAPIAgentsAgentNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAPIAgentsAgentNameOK creates a DeleteAPIAgentsAgentNameOK with default headers values
func NewDeleteAPIAgentsAgentNameOK() *DeleteAPIAgentsAgentNameOK {
	return &DeleteAPIAgentsAgentNameOK{}
}

/* DeleteAPIAgentsAgentNameOK describes a response with status code 200, with default header values.

Success
*/
type DeleteAPIAgentsAgentNameOK struct {
}

func (o *DeleteAPIAgentsAgentNameOK) Error() string {
	return fmt.Sprintf("[DELETE /api/Agents/{agentName}][%d] deleteApiAgentsAgentNameOK ", 200)
}

func (o *DeleteAPIAgentsAgentNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
