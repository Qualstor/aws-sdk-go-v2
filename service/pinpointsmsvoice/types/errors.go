// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// The resource specified in your request already exists.
type AlreadyExistsException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *AlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AlreadyExistsException) ErrorCode() string             { return "AlreadyExistsException" }
func (e *AlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The input you provided is invalid.
type BadRequestException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *BadRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BadRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BadRequestException) ErrorCode() string             { return "BadRequestException" }
func (e *BadRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The API encountered an unexpected error and couldn't complete the request. You
// might be able to successfully issue the request again in the future.
type InternalServiceErrorException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InternalServiceErrorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServiceErrorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServiceErrorException) ErrorCode() string             { return "InternalServiceErrorException" }
func (e *InternalServiceErrorException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// There are too many instances of the specified resource type.
type LimitExceededException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string             { return "LimitExceededException" }
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource you attempted to access doesn't exist.
type NotFoundException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *NotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotFoundException) ErrorCode() string             { return "NotFoundException" }
func (e *NotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You've issued too many requests to the resource. Wait a few minutes, and then
// try again.
type TooManyRequestsException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *TooManyRequestsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyRequestsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyRequestsException) ErrorCode() string             { return "TooManyRequestsException" }
func (e *TooManyRequestsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
