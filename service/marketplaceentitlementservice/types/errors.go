// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/ptr"
)

// An internal error has occurred. Retry your request. If the problem persists,
// post a message with details on the AWS forums.
type InternalServiceErrorException struct {
	Message *string
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
func (e *InternalServiceErrorException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InternalServiceErrorException) HasMessage() bool {
	return e.Message != nil
}

// One or more parameters in your request was invalid.
type InvalidParameterException struct {
	Message *string
}

func (e *InvalidParameterException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterException) ErrorCode() string             { return "InvalidParameterException" }
func (e *InvalidParameterException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidParameterException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidParameterException) HasMessage() bool {
	return e.Message != nil
}

// The calls to the GetEntitlements API are throttled.
type ThrottlingException struct {
	Message *string
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
func (e *ThrottlingException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ThrottlingException) HasMessage() bool {
	return e.Message != nil
}