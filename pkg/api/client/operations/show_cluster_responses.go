// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/sapcc/kubernikus/pkg/api/models"
)

// ShowClusterReader is a Reader for the ShowCluster structure.
type ShowClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewShowClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewShowClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewShowClusterOK creates a ShowClusterOK with default headers values
func NewShowClusterOK() *ShowClusterOK {
	return &ShowClusterOK{}
}

/*ShowClusterOK handles this case with default header values.

OK
*/
type ShowClusterOK struct {
	Payload *models.Cluster
}

func (o *ShowClusterOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/clusters/{name}][%d] showClusterOK  %+v", 200, o.Payload)
}

func (o *ShowClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowClusterDefault creates a ShowClusterDefault with default headers values
func NewShowClusterDefault(code int) *ShowClusterDefault {
	return &ShowClusterDefault{
		_statusCode: code,
	}
}

/*ShowClusterDefault handles this case with default header values.

Error
*/
type ShowClusterDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the show cluster default response
func (o *ShowClusterDefault) Code() int {
	return o._statusCode
}

func (o *ShowClusterDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/clusters/{name}][%d] ShowCluster default  %+v", o._statusCode, o.Payload)
}

func (o *ShowClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}