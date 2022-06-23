// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/Qualstor/aws-sdk-go-v2/service/appconfigdata/types"
)

func ExampleBadRequestDetails_outputUsage() {
	var union types.BadRequestDetails
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.BadRequestDetailsMemberInvalidParameters:
		_ = v.Value // Value is map[string]types.InvalidParameterDetail

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ map[string]types.InvalidParameterDetail
