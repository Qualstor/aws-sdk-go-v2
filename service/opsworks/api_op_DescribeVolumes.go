// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DescribeVolumesRequest
type DescribeVolumesInput struct {
	_ struct{} `type:"structure"`

	// The instance ID. If you use this parameter, DescribeVolumes returns descriptions
	// of the volumes associated with the specified instance.
	InstanceId *string `type:"string"`

	// The RAID array ID. If you use this parameter, DescribeVolumes returns descriptions
	// of the volumes associated with the specified RAID array.
	RaidArrayId *string `type:"string"`

	// A stack ID. The action describes the stack's registered Amazon EBS volumes.
	StackId *string `type:"string"`

	// Am array of volume IDs. If you use this parameter, DescribeVolumes returns
	// descriptions of the specified volumes. Otherwise, it returns a description
	// of every volume.
	VolumeIds []string `type:"list"`
}

// String returns the string representation
func (s DescribeVolumesInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the response to a DescribeVolumes request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DescribeVolumesResult
type DescribeVolumesOutput struct {
	_ struct{} `type:"structure"`

	// An array of volume IDs.
	Volumes []Volume `type:"list"`
}

// String returns the string representation
func (s DescribeVolumesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeVolumes = "DescribeVolumes"

// DescribeVolumesRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Describes an instance's Amazon EBS volumes.
//
// This call accepts only one resource-identifying parameter.
//
// Required Permissions: To use this action, an IAM user must have a Show, Deploy,
// or Manage permissions level for the stack, or an attached policy that explicitly
// grants permissions. For more information about user permissions, see Managing
// User Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using DescribeVolumesRequest.
//    req := client.DescribeVolumesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DescribeVolumes
func (c *Client) DescribeVolumesRequest(input *DescribeVolumesInput) DescribeVolumesRequest {
	op := &aws.Operation{
		Name:       opDescribeVolumes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeVolumesInput{}
	}

	req := c.newRequest(op, input, &DescribeVolumesOutput{})
	return DescribeVolumesRequest{Request: req, Input: input, Copy: c.DescribeVolumesRequest}
}

// DescribeVolumesRequest is the request type for the
// DescribeVolumes API operation.
type DescribeVolumesRequest struct {
	*aws.Request
	Input *DescribeVolumesInput
	Copy  func(*DescribeVolumesInput) DescribeVolumesRequest
}

// Send marshals and sends the DescribeVolumes API request.
func (r DescribeVolumesRequest) Send(ctx context.Context) (*DescribeVolumesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeVolumesResponse{
		DescribeVolumesOutput: r.Request.Data.(*DescribeVolumesOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeVolumesResponse is the response type for the
// DescribeVolumes API operation.
type DescribeVolumesResponse struct {
	*DescribeVolumesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeVolumes request.
func (r *DescribeVolumesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}