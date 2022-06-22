// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	"fmt"
	"github.com/Qualstor/aws-sdk-go-v2/service/elasticbeanstalk/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpApplyEnvironmentManagedAction struct {
}

func (*validateOpApplyEnvironmentManagedAction) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpApplyEnvironmentManagedAction) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ApplyEnvironmentManagedActionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpApplyEnvironmentManagedActionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpAssociateEnvironmentOperationsRole struct {
}

func (*validateOpAssociateEnvironmentOperationsRole) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAssociateEnvironmentOperationsRole) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AssociateEnvironmentOperationsRoleInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAssociateEnvironmentOperationsRoleInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCheckDNSAvailability struct {
}

func (*validateOpCheckDNSAvailability) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCheckDNSAvailability) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CheckDNSAvailabilityInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCheckDNSAvailabilityInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateApplication struct {
}

func (*validateOpCreateApplication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateApplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateApplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateApplicationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateApplicationVersion struct {
}

func (*validateOpCreateApplicationVersion) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateApplicationVersion) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateApplicationVersionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateApplicationVersionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateConfigurationTemplate struct {
}

func (*validateOpCreateConfigurationTemplate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateConfigurationTemplate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateConfigurationTemplateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateConfigurationTemplateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateEnvironment struct {
}

func (*validateOpCreateEnvironment) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateEnvironment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateEnvironmentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateEnvironmentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreatePlatformVersion struct {
}

func (*validateOpCreatePlatformVersion) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreatePlatformVersion) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreatePlatformVersionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreatePlatformVersionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteApplication struct {
}

func (*validateOpDeleteApplication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteApplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteApplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteApplicationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteApplicationVersion struct {
}

func (*validateOpDeleteApplicationVersion) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteApplicationVersion) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteApplicationVersionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteApplicationVersionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteConfigurationTemplate struct {
}

func (*validateOpDeleteConfigurationTemplate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteConfigurationTemplate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteConfigurationTemplateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteConfigurationTemplateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteEnvironmentConfiguration struct {
}

func (*validateOpDeleteEnvironmentConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteEnvironmentConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteEnvironmentConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteEnvironmentConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeConfigurationSettings struct {
}

func (*validateOpDescribeConfigurationSettings) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeConfigurationSettings) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeConfigurationSettingsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeConfigurationSettingsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDisassociateEnvironmentOperationsRole struct {
}

func (*validateOpDisassociateEnvironmentOperationsRole) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDisassociateEnvironmentOperationsRole) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DisassociateEnvironmentOperationsRoleInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDisassociateEnvironmentOperationsRoleInput(input); err != nil {
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

type validateOpRequestEnvironmentInfo struct {
}

func (*validateOpRequestEnvironmentInfo) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRequestEnvironmentInfo) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RequestEnvironmentInfoInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRequestEnvironmentInfoInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRetrieveEnvironmentInfo struct {
}

func (*validateOpRetrieveEnvironmentInfo) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRetrieveEnvironmentInfo) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RetrieveEnvironmentInfoInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRetrieveEnvironmentInfoInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateApplication struct {
}

func (*validateOpUpdateApplication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateApplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateApplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateApplicationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateApplicationResourceLifecycle struct {
}

func (*validateOpUpdateApplicationResourceLifecycle) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateApplicationResourceLifecycle) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateApplicationResourceLifecycleInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateApplicationResourceLifecycleInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateApplicationVersion struct {
}

func (*validateOpUpdateApplicationVersion) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateApplicationVersion) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateApplicationVersionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateApplicationVersionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateConfigurationTemplate struct {
}

func (*validateOpUpdateConfigurationTemplate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateConfigurationTemplate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateConfigurationTemplateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateConfigurationTemplateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateTagsForResource struct {
}

func (*validateOpUpdateTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpValidateConfigurationSettings struct {
}

func (*validateOpValidateConfigurationSettings) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpValidateConfigurationSettings) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ValidateConfigurationSettingsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpValidateConfigurationSettingsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpApplyEnvironmentManagedActionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpApplyEnvironmentManagedAction{}, middleware.After)
}

func addOpAssociateEnvironmentOperationsRoleValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAssociateEnvironmentOperationsRole{}, middleware.After)
}

func addOpCheckDNSAvailabilityValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCheckDNSAvailability{}, middleware.After)
}

func addOpCreateApplicationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateApplication{}, middleware.After)
}

func addOpCreateApplicationVersionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateApplicationVersion{}, middleware.After)
}

func addOpCreateConfigurationTemplateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateConfigurationTemplate{}, middleware.After)
}

func addOpCreateEnvironmentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateEnvironment{}, middleware.After)
}

func addOpCreatePlatformVersionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreatePlatformVersion{}, middleware.After)
}

func addOpDeleteApplicationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteApplication{}, middleware.After)
}

func addOpDeleteApplicationVersionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteApplicationVersion{}, middleware.After)
}

func addOpDeleteConfigurationTemplateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteConfigurationTemplate{}, middleware.After)
}

func addOpDeleteEnvironmentConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteEnvironmentConfiguration{}, middleware.After)
}

func addOpDescribeConfigurationSettingsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeConfigurationSettings{}, middleware.After)
}

func addOpDisassociateEnvironmentOperationsRoleValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDisassociateEnvironmentOperationsRole{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpRequestEnvironmentInfoValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRequestEnvironmentInfo{}, middleware.After)
}

func addOpRetrieveEnvironmentInfoValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRetrieveEnvironmentInfo{}, middleware.After)
}

func addOpUpdateApplicationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateApplication{}, middleware.After)
}

func addOpUpdateApplicationResourceLifecycleValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateApplicationResourceLifecycle{}, middleware.After)
}

func addOpUpdateApplicationVersionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateApplicationVersion{}, middleware.After)
}

func addOpUpdateConfigurationTemplateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateConfigurationTemplate{}, middleware.After)
}

func addOpUpdateTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateTagsForResource{}, middleware.After)
}

func addOpValidateConfigurationSettingsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpValidateConfigurationSettings{}, middleware.After)
}

func validateApplicationResourceLifecycleConfig(v *types.ApplicationResourceLifecycleConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ApplicationResourceLifecycleConfig"}
	if v.VersionLifecycleConfig != nil {
		if err := validateApplicationVersionLifecycleConfig(v.VersionLifecycleConfig); err != nil {
			invalidParams.AddNested("VersionLifecycleConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateApplicationVersionLifecycleConfig(v *types.ApplicationVersionLifecycleConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ApplicationVersionLifecycleConfig"}
	if v.MaxCountRule != nil {
		if err := validateMaxCountRule(v.MaxCountRule); err != nil {
			invalidParams.AddNested("MaxCountRule", err.(smithy.InvalidParamsError))
		}
	}
	if v.MaxAgeRule != nil {
		if err := validateMaxAgeRule(v.MaxAgeRule); err != nil {
			invalidParams.AddNested("MaxAgeRule", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateBuildConfiguration(v *types.BuildConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BuildConfiguration"}
	if v.CodeBuildServiceRole == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CodeBuildServiceRole"))
	}
	if v.Image == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Image"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMaxAgeRule(v *types.MaxAgeRule) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MaxAgeRule"}
	if v.Enabled == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Enabled"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMaxCountRule(v *types.MaxCountRule) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MaxCountRule"}
	if v.Enabled == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Enabled"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSourceBuildInformation(v *types.SourceBuildInformation) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SourceBuildInformation"}
	if len(v.SourceType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("SourceType"))
	}
	if len(v.SourceRepository) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("SourceRepository"))
	}
	if v.SourceLocation == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceLocation"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpApplyEnvironmentManagedActionInput(v *ApplyEnvironmentManagedActionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ApplyEnvironmentManagedActionInput"}
	if v.ActionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ActionId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAssociateEnvironmentOperationsRoleInput(v *AssociateEnvironmentOperationsRoleInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AssociateEnvironmentOperationsRoleInput"}
	if v.EnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentName"))
	}
	if v.OperationsRole == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OperationsRole"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCheckDNSAvailabilityInput(v *CheckDNSAvailabilityInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CheckDNSAvailabilityInput"}
	if v.CNAMEPrefix == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CNAMEPrefix"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateApplicationInput(v *CreateApplicationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateApplicationInput"}
	if v.ApplicationName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationName"))
	}
	if v.ResourceLifecycleConfig != nil {
		if err := validateApplicationResourceLifecycleConfig(v.ResourceLifecycleConfig); err != nil {
			invalidParams.AddNested("ResourceLifecycleConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateApplicationVersionInput(v *CreateApplicationVersionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateApplicationVersionInput"}
	if v.ApplicationName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationName"))
	}
	if v.VersionLabel == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VersionLabel"))
	}
	if v.SourceBuildInformation != nil {
		if err := validateSourceBuildInformation(v.SourceBuildInformation); err != nil {
			invalidParams.AddNested("SourceBuildInformation", err.(smithy.InvalidParamsError))
		}
	}
	if v.BuildConfiguration != nil {
		if err := validateBuildConfiguration(v.BuildConfiguration); err != nil {
			invalidParams.AddNested("BuildConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateConfigurationTemplateInput(v *CreateConfigurationTemplateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateConfigurationTemplateInput"}
	if v.ApplicationName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationName"))
	}
	if v.TemplateName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TemplateName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateEnvironmentInput(v *CreateEnvironmentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateEnvironmentInput"}
	if v.ApplicationName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreatePlatformVersionInput(v *CreatePlatformVersionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreatePlatformVersionInput"}
	if v.PlatformName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PlatformName"))
	}
	if v.PlatformVersion == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PlatformVersion"))
	}
	if v.PlatformDefinitionBundle == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PlatformDefinitionBundle"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteApplicationInput(v *DeleteApplicationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteApplicationInput"}
	if v.ApplicationName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteApplicationVersionInput(v *DeleteApplicationVersionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteApplicationVersionInput"}
	if v.ApplicationName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationName"))
	}
	if v.VersionLabel == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VersionLabel"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteConfigurationTemplateInput(v *DeleteConfigurationTemplateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteConfigurationTemplateInput"}
	if v.ApplicationName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationName"))
	}
	if v.TemplateName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TemplateName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteEnvironmentConfigurationInput(v *DeleteEnvironmentConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteEnvironmentConfigurationInput"}
	if v.ApplicationName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationName"))
	}
	if v.EnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeConfigurationSettingsInput(v *DescribeConfigurationSettingsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeConfigurationSettingsInput"}
	if v.ApplicationName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDisassociateEnvironmentOperationsRoleInput(v *DisassociateEnvironmentOperationsRoleInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisassociateEnvironmentOperationsRoleInput"}
	if v.EnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentName"))
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

func validateOpRequestEnvironmentInfoInput(v *RequestEnvironmentInfoInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RequestEnvironmentInfoInput"}
	if len(v.InfoType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("InfoType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRetrieveEnvironmentInfoInput(v *RetrieveEnvironmentInfoInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RetrieveEnvironmentInfoInput"}
	if len(v.InfoType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("InfoType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateApplicationInput(v *UpdateApplicationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateApplicationInput"}
	if v.ApplicationName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateApplicationResourceLifecycleInput(v *UpdateApplicationResourceLifecycleInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateApplicationResourceLifecycleInput"}
	if v.ApplicationName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationName"))
	}
	if v.ResourceLifecycleConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceLifecycleConfig"))
	} else if v.ResourceLifecycleConfig != nil {
		if err := validateApplicationResourceLifecycleConfig(v.ResourceLifecycleConfig); err != nil {
			invalidParams.AddNested("ResourceLifecycleConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateApplicationVersionInput(v *UpdateApplicationVersionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateApplicationVersionInput"}
	if v.ApplicationName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationName"))
	}
	if v.VersionLabel == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VersionLabel"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateConfigurationTemplateInput(v *UpdateConfigurationTemplateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateConfigurationTemplateInput"}
	if v.ApplicationName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationName"))
	}
	if v.TemplateName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TemplateName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateTagsForResourceInput(v *UpdateTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateTagsForResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpValidateConfigurationSettingsInput(v *ValidateConfigurationSettingsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ValidateConfigurationSettingsInput"}
	if v.ApplicationName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationName"))
	}
	if v.OptionSettings == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OptionSettings"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
