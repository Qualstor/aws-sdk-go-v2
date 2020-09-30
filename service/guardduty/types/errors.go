// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/ptr"
)

// A bad request exception object.
type BadRequestException struct {
	Message *string

	Type *string
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
func (e *BadRequestException) GetType() string {
	return ptr.ToString(e.Type)
}
func (e *BadRequestException) HasType() bool {
	return e.Type != nil
}
func (e *BadRequestException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *BadRequestException) HasMessage() bool {
	return e.Message != nil
}

// An internal server error exception object.
type InternalServerErrorException struct {
	Message *string

	Type *string
}

func (e *InternalServerErrorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServerErrorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServerErrorException) ErrorCode() string             { return "InternalServerErrorException" }
func (e *InternalServerErrorException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }
func (e *InternalServerErrorException) GetType() string {
	return ptr.ToString(e.Type)
}
func (e *InternalServerErrorException) HasType() bool {
	return e.Type != nil
}
func (e *InternalServerErrorException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InternalServerErrorException) HasMessage() bool {
	return e.Message != nil
}