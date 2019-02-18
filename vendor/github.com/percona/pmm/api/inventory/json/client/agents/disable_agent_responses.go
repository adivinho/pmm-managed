// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// DisableAgentReader is a Reader for the DisableAgent structure.
type DisableAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DisableAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDisableAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDisableAgentOK creates a DisableAgentOK with default headers values
func NewDisableAgentOK() *DisableAgentOK {
	return &DisableAgentOK{}
}

/*DisableAgentOK handles this case with default header values.

(empty)
*/
type DisableAgentOK struct {
	Payload interface{}
}

func (o *DisableAgentOK) Error() string {
	return fmt.Sprintf("[POST /v0/inventory/Agents/Disable][%d] disableAgentOK  %+v", 200, o.Payload)
}

func (o *DisableAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DisableAgentBody disable agent body
swagger:model DisableAgentBody
*/
type DisableAgentBody struct {

	// Unique Agent identifier.
	ID string `json:"id,omitempty"`
}

// Validate validates this disable agent body
func (o *DisableAgentBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DisableAgentBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DisableAgentBody) UnmarshalBinary(b []byte) error {
	var res DisableAgentBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}