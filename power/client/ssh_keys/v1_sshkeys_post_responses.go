// Code generated by go-swagger; DO NOT EDIT.

package ssh_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// V1SshkeysPostReader is a Reader for the V1SshkeysPost structure.
type V1SshkeysPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SshkeysPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewV1SshkeysPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewV1SshkeysPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewV1SshkeysPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV1SshkeysPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV1SshkeysPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewV1SshkeysPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewV1SshkeysPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV1SshkeysPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/ssh-keys] v1.sshkeys.post", response, response.Code())
	}
}

// NewV1SshkeysPostCreated creates a V1SshkeysPostCreated with default headers values
func NewV1SshkeysPostCreated() *V1SshkeysPostCreated {
	return &V1SshkeysPostCreated{}
}

/*
V1SshkeysPostCreated describes a response with status code 201, with default header values.

Created
*/
type V1SshkeysPostCreated struct {
	Payload *models.WorkspaceSSHKey
}

// IsSuccess returns true when this v1 sshkeys post created response has a 2xx status code
func (o *V1SshkeysPostCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v1 sshkeys post created response has a 3xx status code
func (o *V1SshkeysPostCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 sshkeys post created response has a 4xx status code
func (o *V1SshkeysPostCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 sshkeys post created response has a 5xx status code
func (o *V1SshkeysPostCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 sshkeys post created response a status code equal to that given
func (o *V1SshkeysPostCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the v1 sshkeys post created response
func (o *V1SshkeysPostCreated) Code() int {
	return 201
}

func (o *V1SshkeysPostCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/ssh-keys][%d] v1SshkeysPostCreated %s", 201, payload)
}

func (o *V1SshkeysPostCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/ssh-keys][%d] v1SshkeysPostCreated %s", 201, payload)
}

func (o *V1SshkeysPostCreated) GetPayload() *models.WorkspaceSSHKey {
	return o.Payload
}

func (o *V1SshkeysPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkspaceSSHKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1SshkeysPostBadRequest creates a V1SshkeysPostBadRequest with default headers values
func NewV1SshkeysPostBadRequest() *V1SshkeysPostBadRequest {
	return &V1SshkeysPostBadRequest{}
}

/*
V1SshkeysPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type V1SshkeysPostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 sshkeys post bad request response has a 2xx status code
func (o *V1SshkeysPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 sshkeys post bad request response has a 3xx status code
func (o *V1SshkeysPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 sshkeys post bad request response has a 4xx status code
func (o *V1SshkeysPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 sshkeys post bad request response has a 5xx status code
func (o *V1SshkeysPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 sshkeys post bad request response a status code equal to that given
func (o *V1SshkeysPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the v1 sshkeys post bad request response
func (o *V1SshkeysPostBadRequest) Code() int {
	return 400
}

func (o *V1SshkeysPostBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/ssh-keys][%d] v1SshkeysPostBadRequest %s", 400, payload)
}

func (o *V1SshkeysPostBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/ssh-keys][%d] v1SshkeysPostBadRequest %s", 400, payload)
}

func (o *V1SshkeysPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1SshkeysPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1SshkeysPostUnauthorized creates a V1SshkeysPostUnauthorized with default headers values
func NewV1SshkeysPostUnauthorized() *V1SshkeysPostUnauthorized {
	return &V1SshkeysPostUnauthorized{}
}

/*
V1SshkeysPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type V1SshkeysPostUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 sshkeys post unauthorized response has a 2xx status code
func (o *V1SshkeysPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 sshkeys post unauthorized response has a 3xx status code
func (o *V1SshkeysPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 sshkeys post unauthorized response has a 4xx status code
func (o *V1SshkeysPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 sshkeys post unauthorized response has a 5xx status code
func (o *V1SshkeysPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 sshkeys post unauthorized response a status code equal to that given
func (o *V1SshkeysPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the v1 sshkeys post unauthorized response
func (o *V1SshkeysPostUnauthorized) Code() int {
	return 401
}

func (o *V1SshkeysPostUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/ssh-keys][%d] v1SshkeysPostUnauthorized %s", 401, payload)
}

func (o *V1SshkeysPostUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/ssh-keys][%d] v1SshkeysPostUnauthorized %s", 401, payload)
}

func (o *V1SshkeysPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1SshkeysPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1SshkeysPostForbidden creates a V1SshkeysPostForbidden with default headers values
func NewV1SshkeysPostForbidden() *V1SshkeysPostForbidden {
	return &V1SshkeysPostForbidden{}
}

/*
V1SshkeysPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type V1SshkeysPostForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 sshkeys post forbidden response has a 2xx status code
func (o *V1SshkeysPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 sshkeys post forbidden response has a 3xx status code
func (o *V1SshkeysPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 sshkeys post forbidden response has a 4xx status code
func (o *V1SshkeysPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 sshkeys post forbidden response has a 5xx status code
func (o *V1SshkeysPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 sshkeys post forbidden response a status code equal to that given
func (o *V1SshkeysPostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the v1 sshkeys post forbidden response
func (o *V1SshkeysPostForbidden) Code() int {
	return 403
}

func (o *V1SshkeysPostForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/ssh-keys][%d] v1SshkeysPostForbidden %s", 403, payload)
}

func (o *V1SshkeysPostForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/ssh-keys][%d] v1SshkeysPostForbidden %s", 403, payload)
}

func (o *V1SshkeysPostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1SshkeysPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1SshkeysPostNotFound creates a V1SshkeysPostNotFound with default headers values
func NewV1SshkeysPostNotFound() *V1SshkeysPostNotFound {
	return &V1SshkeysPostNotFound{}
}

/*
V1SshkeysPostNotFound describes a response with status code 404, with default header values.

Not Found
*/
type V1SshkeysPostNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 sshkeys post not found response has a 2xx status code
func (o *V1SshkeysPostNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 sshkeys post not found response has a 3xx status code
func (o *V1SshkeysPostNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 sshkeys post not found response has a 4xx status code
func (o *V1SshkeysPostNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 sshkeys post not found response has a 5xx status code
func (o *V1SshkeysPostNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 sshkeys post not found response a status code equal to that given
func (o *V1SshkeysPostNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the v1 sshkeys post not found response
func (o *V1SshkeysPostNotFound) Code() int {
	return 404
}

func (o *V1SshkeysPostNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/ssh-keys][%d] v1SshkeysPostNotFound %s", 404, payload)
}

func (o *V1SshkeysPostNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/ssh-keys][%d] v1SshkeysPostNotFound %s", 404, payload)
}

func (o *V1SshkeysPostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1SshkeysPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1SshkeysPostConflict creates a V1SshkeysPostConflict with default headers values
func NewV1SshkeysPostConflict() *V1SshkeysPostConflict {
	return &V1SshkeysPostConflict{}
}

/*
V1SshkeysPostConflict describes a response with status code 409, with default header values.

Conflict
*/
type V1SshkeysPostConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 sshkeys post conflict response has a 2xx status code
func (o *V1SshkeysPostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 sshkeys post conflict response has a 3xx status code
func (o *V1SshkeysPostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 sshkeys post conflict response has a 4xx status code
func (o *V1SshkeysPostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 sshkeys post conflict response has a 5xx status code
func (o *V1SshkeysPostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 sshkeys post conflict response a status code equal to that given
func (o *V1SshkeysPostConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the v1 sshkeys post conflict response
func (o *V1SshkeysPostConflict) Code() int {
	return 409
}

func (o *V1SshkeysPostConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/ssh-keys][%d] v1SshkeysPostConflict %s", 409, payload)
}

func (o *V1SshkeysPostConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/ssh-keys][%d] v1SshkeysPostConflict %s", 409, payload)
}

func (o *V1SshkeysPostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1SshkeysPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1SshkeysPostUnprocessableEntity creates a V1SshkeysPostUnprocessableEntity with default headers values
func NewV1SshkeysPostUnprocessableEntity() *V1SshkeysPostUnprocessableEntity {
	return &V1SshkeysPostUnprocessableEntity{}
}

/*
V1SshkeysPostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type V1SshkeysPostUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 sshkeys post unprocessable entity response has a 2xx status code
func (o *V1SshkeysPostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 sshkeys post unprocessable entity response has a 3xx status code
func (o *V1SshkeysPostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 sshkeys post unprocessable entity response has a 4xx status code
func (o *V1SshkeysPostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 sshkeys post unprocessable entity response has a 5xx status code
func (o *V1SshkeysPostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 sshkeys post unprocessable entity response a status code equal to that given
func (o *V1SshkeysPostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the v1 sshkeys post unprocessable entity response
func (o *V1SshkeysPostUnprocessableEntity) Code() int {
	return 422
}

func (o *V1SshkeysPostUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/ssh-keys][%d] v1SshkeysPostUnprocessableEntity %s", 422, payload)
}

func (o *V1SshkeysPostUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/ssh-keys][%d] v1SshkeysPostUnprocessableEntity %s", 422, payload)
}

func (o *V1SshkeysPostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1SshkeysPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1SshkeysPostInternalServerError creates a V1SshkeysPostInternalServerError with default headers values
func NewV1SshkeysPostInternalServerError() *V1SshkeysPostInternalServerError {
	return &V1SshkeysPostInternalServerError{}
}

/*
V1SshkeysPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type V1SshkeysPostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 sshkeys post internal server error response has a 2xx status code
func (o *V1SshkeysPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 sshkeys post internal server error response has a 3xx status code
func (o *V1SshkeysPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 sshkeys post internal server error response has a 4xx status code
func (o *V1SshkeysPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 sshkeys post internal server error response has a 5xx status code
func (o *V1SshkeysPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v1 sshkeys post internal server error response a status code equal to that given
func (o *V1SshkeysPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the v1 sshkeys post internal server error response
func (o *V1SshkeysPostInternalServerError) Code() int {
	return 500
}

func (o *V1SshkeysPostInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/ssh-keys][%d] v1SshkeysPostInternalServerError %s", 500, payload)
}

func (o *V1SshkeysPostInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/ssh-keys][%d] v1SshkeysPostInternalServerError %s", 500, payload)
}

func (o *V1SshkeysPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1SshkeysPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
