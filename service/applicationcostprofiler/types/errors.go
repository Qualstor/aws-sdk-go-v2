// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// You do not have permission to perform this action.
type AccessDeniedException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *AccessDeniedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessDeniedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessDeniedException) ErrorCode() string             { return "AccessDeniedException" }
func (e *AccessDeniedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An internal server error occurred. Retry your request.
type InternalServerException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InternalServerException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServerException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServerException) ErrorCode() string             { return "InternalServerException" }
func (e *InternalServerException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Your request exceeds one or more of the service quotas.
type ServiceQuotaExceededException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ServiceQuotaExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceQuotaExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceQuotaExceededException) ErrorCode() string             { return "ServiceQuotaExceededException" }
func (e *ServiceQuotaExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The calls to AWS Application Cost Profiler API are throttled. The request was
// denied.
type ThrottlingException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ThrottlingException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ThrottlingException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ThrottlingException) ErrorCode() string             { return "ThrottlingException" }
func (e *ThrottlingException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The input fails to satisfy the constraints for the API.
type ValidationException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ValidationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ValidationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ValidationException) ErrorCode() string             { return "ValidationException" }
func (e *ValidationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
