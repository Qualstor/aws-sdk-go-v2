// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type TagAttendeeInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime SDK attendee ID.
	//
	// AttendeeId is a required field
	AttendeeId *string `location:"uri" locationName:"attendeeId" type:"string" required:"true"`

	// The Amazon Chime SDK meeting ID.
	//
	// MeetingId is a required field
	MeetingId *string `location:"uri" locationName:"meetingId" type:"string" required:"true"`

	// The tag key-value pairs.
	//
	// Tags is a required field
	Tags []Tag `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s TagAttendeeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TagAttendeeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TagAttendeeInput"}

	if s.AttendeeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AttendeeId"))
	}

	if s.MeetingId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeetingId"))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s TagAttendeeInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.AttendeeId != nil {
		v := *s.AttendeeId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "attendeeId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MeetingId != nil {
		v := *s.MeetingId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "meetingId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type TagAttendeeOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s TagAttendeeOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s TagAttendeeOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opTagAttendee = "TagAttendee"

// TagAttendeeRequest returns a request value for making API operation for
// Amazon Chime.
//
// Applies the specified tags to the specified Amazon Chime SDK attendee.
//
//    // Example sending a request using TagAttendeeRequest.
//    req := client.TagAttendeeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/TagAttendee
func (c *Client) TagAttendeeRequest(input *TagAttendeeInput) TagAttendeeRequest {
	op := &aws.Operation{
		Name:       opTagAttendee,
		HTTPMethod: "POST",
		HTTPPath:   "/meetings/{meetingId}/attendees/{attendeeId}/tags?operation=add",
	}

	if input == nil {
		input = &TagAttendeeInput{}
	}

	req := c.newRequest(op, input, &TagAttendeeOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return TagAttendeeRequest{Request: req, Input: input, Copy: c.TagAttendeeRequest}
}

// TagAttendeeRequest is the request type for the
// TagAttendee API operation.
type TagAttendeeRequest struct {
	*aws.Request
	Input *TagAttendeeInput
	Copy  func(*TagAttendeeInput) TagAttendeeRequest
}

// Send marshals and sends the TagAttendee API request.
func (r TagAttendeeRequest) Send(ctx context.Context) (*TagAttendeeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TagAttendeeResponse{
		TagAttendeeOutput: r.Request.Data.(*TagAttendeeOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TagAttendeeResponse is the response type for the
// TagAttendee API operation.
type TagAttendeeResponse struct {
	*TagAttendeeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TagAttendee request.
func (r *TagAttendeeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}