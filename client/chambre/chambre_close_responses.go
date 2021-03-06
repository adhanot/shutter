// Code generated by go-swagger; DO NOT EDIT.

package chambre

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"shutter/models"
)

// ChambreCloseReader is a Reader for the ChambreClose structure.
type ChambreCloseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChambreCloseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChambreCloseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewChambreCloseBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewChambreCloseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewChambreCloseOK creates a ChambreCloseOK with default headers values
func NewChambreCloseOK() *ChambreCloseOK {
	return &ChambreCloseOK{}
}

/* ChambreCloseOK describes a response with status code 200, with default header values.

Successful Close
*/
type ChambreCloseOK struct {
	Payload *models.Success
}

func (o *ChambreCloseOK) Error() string {
	return fmt.Sprintf("[POST /chambre/close][%d] chambreCloseOK  %+v", 200, o.Payload)
}
func (o *ChambreCloseOK) GetPayload() *models.Success {
	return o.Payload
}

func (o *ChambreCloseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Success)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChambreCloseBadRequest creates a ChambreCloseBadRequest with default headers values
func NewChambreCloseBadRequest() *ChambreCloseBadRequest {
	return &ChambreCloseBadRequest{}
}

/* ChambreCloseBadRequest describes a response with status code 400, with default header values.

Error
*/
type ChambreCloseBadRequest struct {
	Payload *models.Error
}

func (o *ChambreCloseBadRequest) Error() string {
	return fmt.Sprintf("[POST /chambre/close][%d] chambreCloseBadRequest  %+v", 400, o.Payload)
}
func (o *ChambreCloseBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ChambreCloseBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChambreCloseInternalServerError creates a ChambreCloseInternalServerError with default headers values
func NewChambreCloseInternalServerError() *ChambreCloseInternalServerError {
	return &ChambreCloseInternalServerError{}
}

/* ChambreCloseInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type ChambreCloseInternalServerError struct {
	Payload *models.ServerError
}

func (o *ChambreCloseInternalServerError) Error() string {
	return fmt.Sprintf("[POST /chambre/close][%d] chambreCloseInternalServerError  %+v", 500, o.Payload)
}
func (o *ChambreCloseInternalServerError) GetPayload() *models.ServerError {
	return o.Payload
}

func (o *ChambreCloseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
