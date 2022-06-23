// Code generated by smithy-go-codegen DO NOT EDIT.

package mediastoredata

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Qualstor/aws-sdk-go-v2/aws/protocol/restjson"
	"github.com/Qualstor/aws-sdk-go-v2/service/mediastoredata/types"
	smithy "github.com/aws/smithy-go"
	smithyio "github.com/aws/smithy-go/io"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
	"strconv"
	"strings"
)

type awsRestjson1_deserializeOpDeleteObject struct {
}

func (*awsRestjson1_deserializeOpDeleteObject) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpDeleteObject) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestjson1_deserializeOpErrorDeleteObject(response, &metadata)
	}
	output := &DeleteObjectOutput{}
	out.Result = output

	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorDeleteObject(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	code := response.Header.Get("X-Amzn-ErrorType")
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	code, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("ContainerNotFoundException", errorCode):
		return awsRestjson1_deserializeErrorContainerNotFoundException(response, errorBody)

	case strings.EqualFold("InternalServerError", errorCode):
		return awsRestjson1_deserializeErrorInternalServerError(response, errorBody)

	case strings.EqualFold("ObjectNotFoundException", errorCode):
		return awsRestjson1_deserializeErrorObjectNotFoundException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

type awsRestjson1_deserializeOpDescribeObject struct {
}

func (*awsRestjson1_deserializeOpDescribeObject) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpDescribeObject) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestjson1_deserializeOpErrorDescribeObject(response, &metadata)
	}
	output := &DescribeObjectOutput{}
	out.Result = output

	err = awsRestjson1_deserializeOpHttpBindingsDescribeObjectOutput(output, response)
	if err != nil {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
	}

	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorDescribeObject(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	code := response.Header.Get("X-Amzn-ErrorType")
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	code, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("ContainerNotFoundException", errorCode):
		return awsRestjson1_deserializeErrorContainerNotFoundException(response, errorBody)

	case strings.EqualFold("InternalServerError", errorCode):
		return awsRestjson1_deserializeErrorInternalServerError(response, errorBody)

	case strings.EqualFold("ObjectNotFoundException", errorCode):
		return awsRestjson1_deserializeErrorObjectNotFoundException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsRestjson1_deserializeOpHttpBindingsDescribeObjectOutput(v *DescribeObjectOutput, response *smithyhttp.Response) error {
	if v == nil {
		return fmt.Errorf("unsupported deserialization for nil %T", v)
	}

	if headerValues := response.Header.Values("Cache-Control"); len(headerValues) != 0 {
		headerValues[0] = strings.TrimSpace(headerValues[0])
		v.CacheControl = ptr.String(headerValues[0])
	}

	if headerValues := response.Header.Values("Content-Length"); len(headerValues) != 0 {
		headerValues[0] = strings.TrimSpace(headerValues[0])
		vv, err := strconv.ParseInt(headerValues[0], 0, 64)
		if err != nil {
			return err
		}
		v.ContentLength = ptr.Int64(vv)
	}

	if headerValues := response.Header.Values("Content-Type"); len(headerValues) != 0 {
		headerValues[0] = strings.TrimSpace(headerValues[0])
		v.ContentType = ptr.String(headerValues[0])
	}

	if headerValues := response.Header.Values("ETag"); len(headerValues) != 0 {
		headerValues[0] = strings.TrimSpace(headerValues[0])
		v.ETag = ptr.String(headerValues[0])
	}

	if headerValues := response.Header.Values("Last-Modified"); len(headerValues) != 0 {
		headerValues[0] = strings.TrimSpace(headerValues[0])
		t, err := smithytime.ParseHTTPDate(headerValues[0])
		if err != nil {
			return err
		}
		v.LastModified = ptr.Time(t)
	}

	return nil
}

type awsRestjson1_deserializeOpGetObject struct {
}

func (*awsRestjson1_deserializeOpGetObject) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpGetObject) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestjson1_deserializeOpErrorGetObject(response, &metadata)
	}
	output := &GetObjectOutput{}
	out.Result = output

	err = awsRestjson1_deserializeOpHttpBindingsGetObjectOutput(output, response)
	if err != nil {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
	}

	err = awsRestjson1_deserializeOpDocumentGetObjectOutput(output, response.Body)
	if err != nil {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
	}

	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorGetObject(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	code := response.Header.Get("X-Amzn-ErrorType")
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	code, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("ContainerNotFoundException", errorCode):
		return awsRestjson1_deserializeErrorContainerNotFoundException(response, errorBody)

	case strings.EqualFold("InternalServerError", errorCode):
		return awsRestjson1_deserializeErrorInternalServerError(response, errorBody)

	case strings.EqualFold("ObjectNotFoundException", errorCode):
		return awsRestjson1_deserializeErrorObjectNotFoundException(response, errorBody)

	case strings.EqualFold("RequestedRangeNotSatisfiableException", errorCode):
		return awsRestjson1_deserializeErrorRequestedRangeNotSatisfiableException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsRestjson1_deserializeOpHttpBindingsGetObjectOutput(v *GetObjectOutput, response *smithyhttp.Response) error {
	if v == nil {
		return fmt.Errorf("unsupported deserialization for nil %T", v)
	}

	if headerValues := response.Header.Values("Cache-Control"); len(headerValues) != 0 {
		headerValues[0] = strings.TrimSpace(headerValues[0])
		v.CacheControl = ptr.String(headerValues[0])
	}

	if headerValues := response.Header.Values("Content-Length"); len(headerValues) != 0 {
		headerValues[0] = strings.TrimSpace(headerValues[0])
		vv, err := strconv.ParseInt(headerValues[0], 0, 64)
		if err != nil {
			return err
		}
		v.ContentLength = ptr.Int64(vv)
	}

	if headerValues := response.Header.Values("Content-Range"); len(headerValues) != 0 {
		headerValues[0] = strings.TrimSpace(headerValues[0])
		v.ContentRange = ptr.String(headerValues[0])
	}

	if headerValues := response.Header.Values("Content-Type"); len(headerValues) != 0 {
		headerValues[0] = strings.TrimSpace(headerValues[0])
		v.ContentType = ptr.String(headerValues[0])
	}

	if headerValues := response.Header.Values("ETag"); len(headerValues) != 0 {
		headerValues[0] = strings.TrimSpace(headerValues[0])
		v.ETag = ptr.String(headerValues[0])
	}

	if headerValues := response.Header.Values("Last-Modified"); len(headerValues) != 0 {
		headerValues[0] = strings.TrimSpace(headerValues[0])
		t, err := smithytime.ParseHTTPDate(headerValues[0])
		if err != nil {
			return err
		}
		v.LastModified = ptr.Time(t)
	}

	v.StatusCode = int32(response.StatusCode)

	return nil
}
func awsRestjson1_deserializeOpDocumentGetObjectOutput(v *GetObjectOutput, body io.ReadCloser) error {
	if v == nil {
		return fmt.Errorf("unsupported deserialization of nil %T", v)
	}

	v.Body = body
	return nil
}

type awsRestjson1_deserializeOpListItems struct {
}

func (*awsRestjson1_deserializeOpListItems) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpListItems) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestjson1_deserializeOpErrorListItems(response, &metadata)
	}
	output := &ListItemsOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)

	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	err = awsRestjson1_deserializeOpDocumentListItemsOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body with invalid JSON, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorListItems(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	code := response.Header.Get("X-Amzn-ErrorType")
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	code, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("ContainerNotFoundException", errorCode):
		return awsRestjson1_deserializeErrorContainerNotFoundException(response, errorBody)

	case strings.EqualFold("InternalServerError", errorCode):
		return awsRestjson1_deserializeErrorInternalServerError(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsRestjson1_deserializeOpDocumentListItemsOutput(v **ListItemsOutput, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *ListItemsOutput
	if *v == nil {
		sv = &ListItemsOutput{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Items":
			if err := awsRestjson1_deserializeDocumentItemList(&sv.Items, value); err != nil {
				return err
			}

		case "NextToken":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected PaginationToken to be of type string, got %T instead", value)
				}
				sv.NextToken = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

type awsRestjson1_deserializeOpPutObject struct {
}

func (*awsRestjson1_deserializeOpPutObject) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpPutObject) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestjson1_deserializeOpErrorPutObject(response, &metadata)
	}
	output := &PutObjectOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)

	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	err = awsRestjson1_deserializeOpDocumentPutObjectOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body with invalid JSON, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorPutObject(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	code := response.Header.Get("X-Amzn-ErrorType")
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	code, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("ContainerNotFoundException", errorCode):
		return awsRestjson1_deserializeErrorContainerNotFoundException(response, errorBody)

	case strings.EqualFold("InternalServerError", errorCode):
		return awsRestjson1_deserializeErrorInternalServerError(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsRestjson1_deserializeOpDocumentPutObjectOutput(v **PutObjectOutput, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *PutObjectOutput
	if *v == nil {
		sv = &PutObjectOutput{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "ContentSHA256":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected SHA256Hash to be of type string, got %T instead", value)
				}
				sv.ContentSHA256 = ptr.String(jtv)
			}

		case "ETag":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ETag to be of type string, got %T instead", value)
				}
				sv.ETag = ptr.String(jtv)
			}

		case "StorageClass":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected StorageClass to be of type string, got %T instead", value)
				}
				sv.StorageClass = types.StorageClass(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeErrorContainerNotFoundException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.ContainerNotFoundException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentContainerNotFoundException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeErrorInternalServerError(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.InternalServerError{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentInternalServerError(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeErrorObjectNotFoundException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.ObjectNotFoundException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentObjectNotFoundException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeErrorRequestedRangeNotSatisfiableException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.RequestedRangeNotSatisfiableException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentRequestedRangeNotSatisfiableException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeDocumentContainerNotFoundException(v **types.ContainerNotFoundException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.ContainerNotFoundException
	if *v == nil {
		sv = &types.ContainerNotFoundException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ErrorMessage to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentInternalServerError(v **types.InternalServerError, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.InternalServerError
	if *v == nil {
		sv = &types.InternalServerError{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ErrorMessage to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentItem(v **types.Item, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.Item
	if *v == nil {
		sv = &types.Item{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "ContentLength":
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected NonNegativeLong to be json.Number, got %T instead", value)
				}
				i64, err := jtv.Int64()
				if err != nil {
					return err
				}
				sv.ContentLength = ptr.Int64(i64)
			}

		case "ContentType":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ContentType to be of type string, got %T instead", value)
				}
				sv.ContentType = ptr.String(jtv)
			}

		case "ETag":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ETag to be of type string, got %T instead", value)
				}
				sv.ETag = ptr.String(jtv)
			}

		case "LastModified":
			if value != nil {
				switch jtv := value.(type) {
				case json.Number:
					f64, err := jtv.Float64()
					if err != nil {
						return err
					}
					sv.LastModified = ptr.Time(smithytime.ParseEpochSeconds(f64))

				default:
					return fmt.Errorf("expected TimeStamp to be a JSON Number, got %T instead", value)

				}
			}

		case "Name":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ItemName to be of type string, got %T instead", value)
				}
				sv.Name = ptr.String(jtv)
			}

		case "Type":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ItemType to be of type string, got %T instead", value)
				}
				sv.Type = types.ItemType(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentItemList(v *[]types.Item, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var cv []types.Item
	if *v == nil {
		cv = []types.Item{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col types.Item
		destAddr := &col
		if err := awsRestjson1_deserializeDocumentItem(&destAddr, value); err != nil {
			return err
		}
		col = *destAddr
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsRestjson1_deserializeDocumentObjectNotFoundException(v **types.ObjectNotFoundException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.ObjectNotFoundException
	if *v == nil {
		sv = &types.ObjectNotFoundException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ErrorMessage to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentRequestedRangeNotSatisfiableException(v **types.RequestedRangeNotSatisfiableException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.RequestedRangeNotSatisfiableException
	if *v == nil {
		sv = &types.RequestedRangeNotSatisfiableException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ErrorMessage to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}
