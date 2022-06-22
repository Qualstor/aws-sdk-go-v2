// Code generated by smithy-go-codegen DO NOT EDIT.

package apprunner

import (
	"context"
	"fmt"
	"github.com/Qualstor/aws-sdk-go-v2/service/apprunner/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpAssociateCustomDomain struct {
}

func (*validateOpAssociateCustomDomain) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAssociateCustomDomain) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AssociateCustomDomainInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAssociateCustomDomainInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateAutoScalingConfiguration struct {
}

func (*validateOpCreateAutoScalingConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateAutoScalingConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateAutoScalingConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateAutoScalingConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateConnection struct {
}

func (*validateOpCreateConnection) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateConnection) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateConnectionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateConnectionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateObservabilityConfiguration struct {
}

func (*validateOpCreateObservabilityConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateObservabilityConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateObservabilityConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateObservabilityConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateService struct {
}

func (*validateOpCreateService) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateService) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateServiceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateServiceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateVpcConnector struct {
}

func (*validateOpCreateVpcConnector) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateVpcConnector) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateVpcConnectorInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateVpcConnectorInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteAutoScalingConfiguration struct {
}

func (*validateOpDeleteAutoScalingConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteAutoScalingConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteAutoScalingConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteAutoScalingConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteConnection struct {
}

func (*validateOpDeleteConnection) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteConnection) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteConnectionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteConnectionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteObservabilityConfiguration struct {
}

func (*validateOpDeleteObservabilityConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteObservabilityConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteObservabilityConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteObservabilityConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteService struct {
}

func (*validateOpDeleteService) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteService) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteServiceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteServiceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteVpcConnector struct {
}

func (*validateOpDeleteVpcConnector) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteVpcConnector) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteVpcConnectorInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteVpcConnectorInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeAutoScalingConfiguration struct {
}

func (*validateOpDescribeAutoScalingConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeAutoScalingConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeAutoScalingConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeAutoScalingConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeCustomDomains struct {
}

func (*validateOpDescribeCustomDomains) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeCustomDomains) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeCustomDomainsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeCustomDomainsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeObservabilityConfiguration struct {
}

func (*validateOpDescribeObservabilityConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeObservabilityConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeObservabilityConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeObservabilityConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeService struct {
}

func (*validateOpDescribeService) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeService) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeServiceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeServiceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeVpcConnector struct {
}

func (*validateOpDescribeVpcConnector) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeVpcConnector) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeVpcConnectorInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeVpcConnectorInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDisassociateCustomDomain struct {
}

func (*validateOpDisassociateCustomDomain) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDisassociateCustomDomain) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DisassociateCustomDomainInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDisassociateCustomDomainInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListOperations struct {
}

func (*validateOpListOperations) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListOperations) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListOperationsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListOperationsInput(input); err != nil {
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

type validateOpPauseService struct {
}

func (*validateOpPauseService) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPauseService) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PauseServiceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPauseServiceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpResumeService struct {
}

func (*validateOpResumeService) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpResumeService) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ResumeServiceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpResumeServiceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartDeployment struct {
}

func (*validateOpStartDeployment) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartDeployment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartDeploymentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartDeploymentInput(input); err != nil {
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

type validateOpUpdateService struct {
}

func (*validateOpUpdateService) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateService) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateServiceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateServiceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAssociateCustomDomainValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAssociateCustomDomain{}, middleware.After)
}

func addOpCreateAutoScalingConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateAutoScalingConfiguration{}, middleware.After)
}

func addOpCreateConnectionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateConnection{}, middleware.After)
}

func addOpCreateObservabilityConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateObservabilityConfiguration{}, middleware.After)
}

func addOpCreateServiceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateService{}, middleware.After)
}

func addOpCreateVpcConnectorValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateVpcConnector{}, middleware.After)
}

func addOpDeleteAutoScalingConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteAutoScalingConfiguration{}, middleware.After)
}

func addOpDeleteConnectionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteConnection{}, middleware.After)
}

func addOpDeleteObservabilityConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteObservabilityConfiguration{}, middleware.After)
}

func addOpDeleteServiceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteService{}, middleware.After)
}

func addOpDeleteVpcConnectorValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteVpcConnector{}, middleware.After)
}

func addOpDescribeAutoScalingConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeAutoScalingConfiguration{}, middleware.After)
}

func addOpDescribeCustomDomainsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeCustomDomains{}, middleware.After)
}

func addOpDescribeObservabilityConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeObservabilityConfiguration{}, middleware.After)
}

func addOpDescribeServiceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeService{}, middleware.After)
}

func addOpDescribeVpcConnectorValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeVpcConnector{}, middleware.After)
}

func addOpDisassociateCustomDomainValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDisassociateCustomDomain{}, middleware.After)
}

func addOpListOperationsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListOperations{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpPauseServiceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPauseService{}, middleware.After)
}

func addOpResumeServiceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpResumeService{}, middleware.After)
}

func addOpStartDeploymentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartDeployment{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateServiceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateService{}, middleware.After)
}

func validateCodeConfiguration(v *types.CodeConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CodeConfiguration"}
	if len(v.ConfigurationSource) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ConfigurationSource"))
	}
	if v.CodeConfigurationValues != nil {
		if err := validateCodeConfigurationValues(v.CodeConfigurationValues); err != nil {
			invalidParams.AddNested("CodeConfigurationValues", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCodeConfigurationValues(v *types.CodeConfigurationValues) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CodeConfigurationValues"}
	if len(v.Runtime) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Runtime"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCodeRepository(v *types.CodeRepository) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CodeRepository"}
	if v.RepositoryUrl == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RepositoryUrl"))
	}
	if v.SourceCodeVersion == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceCodeVersion"))
	} else if v.SourceCodeVersion != nil {
		if err := validateSourceCodeVersion(v.SourceCodeVersion); err != nil {
			invalidParams.AddNested("SourceCodeVersion", err.(smithy.InvalidParamsError))
		}
	}
	if v.CodeConfiguration != nil {
		if err := validateCodeConfiguration(v.CodeConfiguration); err != nil {
			invalidParams.AddNested("CodeConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateEncryptionConfiguration(v *types.EncryptionConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EncryptionConfiguration"}
	if v.KmsKey == nil {
		invalidParams.Add(smithy.NewErrParamRequired("KmsKey"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateImageRepository(v *types.ImageRepository) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ImageRepository"}
	if v.ImageIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ImageIdentifier"))
	}
	if len(v.ImageRepositoryType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ImageRepositoryType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateServiceObservabilityConfiguration(v *types.ServiceObservabilityConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ServiceObservabilityConfiguration"}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSourceCodeVersion(v *types.SourceCodeVersion) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SourceCodeVersion"}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSourceConfiguration(v *types.SourceConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SourceConfiguration"}
	if v.CodeRepository != nil {
		if err := validateCodeRepository(v.CodeRepository); err != nil {
			invalidParams.AddNested("CodeRepository", err.(smithy.InvalidParamsError))
		}
	}
	if v.ImageRepository != nil {
		if err := validateImageRepository(v.ImageRepository); err != nil {
			invalidParams.AddNested("ImageRepository", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTraceConfiguration(v *types.TraceConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TraceConfiguration"}
	if len(v.Vendor) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Vendor"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAssociateCustomDomainInput(v *AssociateCustomDomainInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AssociateCustomDomainInput"}
	if v.ServiceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceArn"))
	}
	if v.DomainName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateAutoScalingConfigurationInput(v *CreateAutoScalingConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateAutoScalingConfigurationInput"}
	if v.AutoScalingConfigurationName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AutoScalingConfigurationName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateConnectionInput(v *CreateConnectionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateConnectionInput"}
	if v.ConnectionName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConnectionName"))
	}
	if len(v.ProviderType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ProviderType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateObservabilityConfigurationInput(v *CreateObservabilityConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateObservabilityConfigurationInput"}
	if v.ObservabilityConfigurationName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ObservabilityConfigurationName"))
	}
	if v.TraceConfiguration != nil {
		if err := validateTraceConfiguration(v.TraceConfiguration); err != nil {
			invalidParams.AddNested("TraceConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateServiceInput(v *CreateServiceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateServiceInput"}
	if v.ServiceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceName"))
	}
	if v.SourceConfiguration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceConfiguration"))
	} else if v.SourceConfiguration != nil {
		if err := validateSourceConfiguration(v.SourceConfiguration); err != nil {
			invalidParams.AddNested("SourceConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.EncryptionConfiguration != nil {
		if err := validateEncryptionConfiguration(v.EncryptionConfiguration); err != nil {
			invalidParams.AddNested("EncryptionConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.ObservabilityConfiguration != nil {
		if err := validateServiceObservabilityConfiguration(v.ObservabilityConfiguration); err != nil {
			invalidParams.AddNested("ObservabilityConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateVpcConnectorInput(v *CreateVpcConnectorInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateVpcConnectorInput"}
	if v.VpcConnectorName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VpcConnectorName"))
	}
	if v.Subnets == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Subnets"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteAutoScalingConfigurationInput(v *DeleteAutoScalingConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteAutoScalingConfigurationInput"}
	if v.AutoScalingConfigurationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AutoScalingConfigurationArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteConnectionInput(v *DeleteConnectionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteConnectionInput"}
	if v.ConnectionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConnectionArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteObservabilityConfigurationInput(v *DeleteObservabilityConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteObservabilityConfigurationInput"}
	if v.ObservabilityConfigurationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ObservabilityConfigurationArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteServiceInput(v *DeleteServiceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteServiceInput"}
	if v.ServiceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteVpcConnectorInput(v *DeleteVpcConnectorInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteVpcConnectorInput"}
	if v.VpcConnectorArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VpcConnectorArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeAutoScalingConfigurationInput(v *DescribeAutoScalingConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeAutoScalingConfigurationInput"}
	if v.AutoScalingConfigurationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AutoScalingConfigurationArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeCustomDomainsInput(v *DescribeCustomDomainsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeCustomDomainsInput"}
	if v.ServiceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeObservabilityConfigurationInput(v *DescribeObservabilityConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeObservabilityConfigurationInput"}
	if v.ObservabilityConfigurationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ObservabilityConfigurationArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeServiceInput(v *DescribeServiceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeServiceInput"}
	if v.ServiceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeVpcConnectorInput(v *DescribeVpcConnectorInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeVpcConnectorInput"}
	if v.VpcConnectorArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VpcConnectorArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDisassociateCustomDomainInput(v *DisassociateCustomDomainInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisassociateCustomDomainInput"}
	if v.ServiceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceArn"))
	}
	if v.DomainName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListOperationsInput(v *ListOperationsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListOperationsInput"}
	if v.ServiceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceArn"))
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

func validateOpPauseServiceInput(v *PauseServiceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PauseServiceInput"}
	if v.ServiceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpResumeServiceInput(v *ResumeServiceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ResumeServiceInput"}
	if v.ServiceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartDeploymentInput(v *StartDeploymentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartDeploymentInput"}
	if v.ServiceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceArn"))
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

func validateOpUpdateServiceInput(v *UpdateServiceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateServiceInput"}
	if v.ServiceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceArn"))
	}
	if v.SourceConfiguration != nil {
		if err := validateSourceConfiguration(v.SourceConfiguration); err != nil {
			invalidParams.AddNested("SourceConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.ObservabilityConfiguration != nil {
		if err := validateServiceObservabilityConfiguration(v.ObservabilityConfiguration); err != nil {
			invalidParams.AddNested("ObservabilityConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
