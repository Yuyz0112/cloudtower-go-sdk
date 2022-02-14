// Code generated by go-swagger; DO NOT EDIT.

package backup_store_repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Yuyz0112/cloudtower-go-sdk/models"
)

// GetBackupStoreRepositoriesConnectionReader is a Reader for the GetBackupStoreRepositoriesConnection structure.
type GetBackupStoreRepositoriesConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBackupStoreRepositoriesConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBackupStoreRepositoriesConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetBackupStoreRepositoriesConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBackupStoreRepositoriesConnectionOK creates a GetBackupStoreRepositoriesConnectionOK with default headers values
func NewGetBackupStoreRepositoriesConnectionOK() *GetBackupStoreRepositoriesConnectionOK {
	return &GetBackupStoreRepositoriesConnectionOK{}
}

/* GetBackupStoreRepositoriesConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetBackupStoreRepositoriesConnectionOK struct {
	Payload *models.BackupStoreRepositoryConnection
}

func (o *GetBackupStoreRepositoriesConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-backup-store-repositories-connection][%d] getBackupStoreRepositoriesConnectionOK  %+v", 200, o.Payload)
}
func (o *GetBackupStoreRepositoriesConnectionOK) GetPayload() *models.BackupStoreRepositoryConnection {
	return o.Payload
}

func (o *GetBackupStoreRepositoriesConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BackupStoreRepositoryConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackupStoreRepositoriesConnectionBadRequest creates a GetBackupStoreRepositoriesConnectionBadRequest with default headers values
func NewGetBackupStoreRepositoriesConnectionBadRequest() *GetBackupStoreRepositoriesConnectionBadRequest {
	return &GetBackupStoreRepositoriesConnectionBadRequest{}
}

/* GetBackupStoreRepositoriesConnectionBadRequest describes a response with status code 400, with default header values.

GetBackupStoreRepositoriesConnectionBadRequest get backup store repositories connection bad request
*/
type GetBackupStoreRepositoriesConnectionBadRequest struct {
	Payload string
}

func (o *GetBackupStoreRepositoriesConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-backup-store-repositories-connection][%d] getBackupStoreRepositoriesConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetBackupStoreRepositoriesConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetBackupStoreRepositoriesConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
