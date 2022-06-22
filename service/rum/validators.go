// Code generated by smithy-go-codegen DO NOT EDIT.

package rum

import (
	"context"
	"fmt"
	"github.com/Qualstor/aws-sdk-go-v2/service/rum/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateAppMonitor struct {
}

func (*validateOpCreateAppMonitor) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateAppMonitor) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateAppMonitorInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateAppMonitorInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteAppMonitor struct {
}

func (*validateOpDeleteAppMonitor) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteAppMonitor) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteAppMonitorInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteAppMonitorInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetAppMonitorData struct {
}

func (*validateOpGetAppMonitorData) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetAppMonitorData) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetAppMonitorDataInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetAppMonitorDataInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetAppMonitor struct {
}

func (*validateOpGetAppMonitor) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetAppMonitor) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetAppMonitorInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetAppMonitorInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutRumEvents struct {
}

func (*validateOpPutRumEvents) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutRumEvents) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutRumEventsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutRumEventsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateAppMonitor struct {
}

func (*validateOpUpdateAppMonitor) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateAppMonitor) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateAppMonitorInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateAppMonitorInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateAppMonitorValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateAppMonitor{}, middleware.After)
}

func addOpDeleteAppMonitorValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteAppMonitor{}, middleware.After)
}

func addOpGetAppMonitorDataValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetAppMonitorData{}, middleware.After)
}

func addOpGetAppMonitorValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetAppMonitor{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpPutRumEventsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutRumEvents{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateAppMonitorValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateAppMonitor{}, middleware.After)
}

func validateRumEvent(v *types.RumEvent) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RumEvent"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if v.Timestamp == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Timestamp"))
	}
	if v.Type == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if v.Details == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Details"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRumEventList(v []types.RumEvent) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RumEventList"}
	for i := range v {
		if err := validateRumEvent(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTimeRange(v *types.TimeRange) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TimeRange"}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateAppMonitorInput(v *CreateAppMonitorInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateAppMonitorInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Domain == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Domain"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteAppMonitorInput(v *DeleteAppMonitorInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteAppMonitorInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetAppMonitorDataInput(v *GetAppMonitorDataInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetAppMonitorDataInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.TimeRange == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TimeRange"))
	} else if v.TimeRange != nil {
		if err := validateTimeRange(v.TimeRange); err != nil {
			invalidParams.AddNested("TimeRange", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetAppMonitorInput(v *GetAppMonitorInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetAppMonitorInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutRumEventsInput(v *PutRumEventsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutRumEventsInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if v.BatchId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BatchId"))
	}
	if v.AppMonitorDetails == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppMonitorDetails"))
	}
	if v.UserDetails == nil {
		invalidParams.Add(smithy.NewErrParamRequired("UserDetails"))
	}
	if v.RumEvents == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RumEvents"))
	} else if v.RumEvents != nil {
		if err := validateRumEventList(v.RumEvents); err != nil {
			invalidParams.AddNested("RumEvents", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateAppMonitorInput(v *UpdateAppMonitorInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateAppMonitorInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
