// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volume_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudVolumegroupsStorageDetailsGetReader is a Reader for the PcloudVolumegroupsStorageDetailsGet structure.
type PcloudVolumegroupsStorageDetailsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudVolumegroupsStorageDetailsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudVolumegroupsStorageDetailsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudVolumegroupsStorageDetailsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudVolumegroupsStorageDetailsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudVolumegroupsStorageDetailsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPcloudVolumegroupsStorageDetailsGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudVolumegroupsStorageDetailsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudVolumegroupsStorageDetailsGetOK creates a PcloudVolumegroupsStorageDetailsGetOK with default headers values
func NewPcloudVolumegroupsStorageDetailsGetOK() *PcloudVolumegroupsStorageDetailsGetOK {
	return &PcloudVolumegroupsStorageDetailsGetOK{}
}

/* PcloudVolumegroupsStorageDetailsGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudVolumegroupsStorageDetailsGetOK struct {
	Payload *models.VolumeGroupStorageDetails
}

func (o *PcloudVolumegroupsStorageDetailsGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}/storage-details][%d] pcloudVolumegroupsStorageDetailsGetOK  %+v", 200, o.Payload)
}
func (o *PcloudVolumegroupsStorageDetailsGetOK) GetPayload() *models.VolumeGroupStorageDetails {
	return o.Payload
}

func (o *PcloudVolumegroupsStorageDetailsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeGroupStorageDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsStorageDetailsGetBadRequest creates a PcloudVolumegroupsStorageDetailsGetBadRequest with default headers values
func NewPcloudVolumegroupsStorageDetailsGetBadRequest() *PcloudVolumegroupsStorageDetailsGetBadRequest {
	return &PcloudVolumegroupsStorageDetailsGetBadRequest{}
}

/* PcloudVolumegroupsStorageDetailsGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudVolumegroupsStorageDetailsGetBadRequest struct {
	Payload *models.Error
}

func (o *PcloudVolumegroupsStorageDetailsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}/storage-details][%d] pcloudVolumegroupsStorageDetailsGetBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudVolumegroupsStorageDetailsGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsStorageDetailsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsStorageDetailsGetForbidden creates a PcloudVolumegroupsStorageDetailsGetForbidden with default headers values
func NewPcloudVolumegroupsStorageDetailsGetForbidden() *PcloudVolumegroupsStorageDetailsGetForbidden {
	return &PcloudVolumegroupsStorageDetailsGetForbidden{}
}

/* PcloudVolumegroupsStorageDetailsGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudVolumegroupsStorageDetailsGetForbidden struct {
	Payload *models.Error
}

func (o *PcloudVolumegroupsStorageDetailsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}/storage-details][%d] pcloudVolumegroupsStorageDetailsGetForbidden  %+v", 403, o.Payload)
}
func (o *PcloudVolumegroupsStorageDetailsGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsStorageDetailsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsStorageDetailsGetNotFound creates a PcloudVolumegroupsStorageDetailsGetNotFound with default headers values
func NewPcloudVolumegroupsStorageDetailsGetNotFound() *PcloudVolumegroupsStorageDetailsGetNotFound {
	return &PcloudVolumegroupsStorageDetailsGetNotFound{}
}

/* PcloudVolumegroupsStorageDetailsGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudVolumegroupsStorageDetailsGetNotFound struct {
	Payload *models.Error
}

func (o *PcloudVolumegroupsStorageDetailsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}/storage-details][%d] pcloudVolumegroupsStorageDetailsGetNotFound  %+v", 404, o.Payload)
}
func (o *PcloudVolumegroupsStorageDetailsGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsStorageDetailsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsStorageDetailsGetTooManyRequests creates a PcloudVolumegroupsStorageDetailsGetTooManyRequests with default headers values
func NewPcloudVolumegroupsStorageDetailsGetTooManyRequests() *PcloudVolumegroupsStorageDetailsGetTooManyRequests {
	return &PcloudVolumegroupsStorageDetailsGetTooManyRequests{}
}

/* PcloudVolumegroupsStorageDetailsGetTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type PcloudVolumegroupsStorageDetailsGetTooManyRequests struct {
	Payload *models.Error
}

func (o *PcloudVolumegroupsStorageDetailsGetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}/storage-details][%d] pcloudVolumegroupsStorageDetailsGetTooManyRequests  %+v", 429, o.Payload)
}
func (o *PcloudVolumegroupsStorageDetailsGetTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsStorageDetailsGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsStorageDetailsGetInternalServerError creates a PcloudVolumegroupsStorageDetailsGetInternalServerError with default headers values
func NewPcloudVolumegroupsStorageDetailsGetInternalServerError() *PcloudVolumegroupsStorageDetailsGetInternalServerError {
	return &PcloudVolumegroupsStorageDetailsGetInternalServerError{}
}

/* PcloudVolumegroupsStorageDetailsGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudVolumegroupsStorageDetailsGetInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudVolumegroupsStorageDetailsGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}/storage-details][%d] pcloudVolumegroupsStorageDetailsGetInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudVolumegroupsStorageDetailsGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsStorageDetailsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
