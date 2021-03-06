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

// GetClusterInfoReader is a Reader for the GetClusterInfo structure.
type GetClusterInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClusterInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetClusterInfoDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterInfoOK creates a GetClusterInfoOK with default headers values
func NewGetClusterInfoOK() *GetClusterInfoOK {
	return &GetClusterInfoOK{}
}

/*GetClusterInfoOK handles this case with default header values.

OK
*/
type GetClusterInfoOK struct {
	Payload *models.KlusterInfo
}

func (o *GetClusterInfoOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/clusters/{name}/info][%d] getClusterInfoOK  %+v", 200, o.Payload)
}

func (o *GetClusterInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KlusterInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterInfoDefault creates a GetClusterInfoDefault with default headers values
func NewGetClusterInfoDefault(code int) *GetClusterInfoDefault {
	return &GetClusterInfoDefault{
		_statusCode: code,
	}
}

/*GetClusterInfoDefault handles this case with default header values.

Error
*/
type GetClusterInfoDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get cluster info default response
func (o *GetClusterInfoDefault) Code() int {
	return o._statusCode
}

func (o *GetClusterInfoDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/clusters/{name}/info][%d] GetClusterInfo default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterInfoDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
