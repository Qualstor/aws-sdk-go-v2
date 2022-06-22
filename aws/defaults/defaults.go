// Code generated by github.com/Qualstor/aws-sdk-go-v2/internal/codegen/cmd/defaultsconfig. DO NOT EDIT.

package defaults

import (
	"fmt"
	"github.com/Qualstor/aws-sdk-go-v2/aws"
	"time"
)

// GetModeConfiguration returns the default Configuration descriptor for the given mode.
//
// Supports the following modes: cross-region, in-region, mobile, standard
func GetModeConfiguration(mode aws.DefaultsMode) (Configuration, error) {
	var mv aws.DefaultsMode
	mv.SetFromString(string(mode))

	switch mv {
	case aws.DefaultsModeCrossRegion:
		settings := Configuration{
			ConnectTimeout:        aws.Duration(3100 * time.Millisecond),
			RetryMode:             aws.RetryMode("standard"),
			TLSNegotiationTimeout: aws.Duration(3100 * time.Millisecond),
		}
		return settings, nil
	case aws.DefaultsModeInRegion:
		settings := Configuration{
			ConnectTimeout:        aws.Duration(1100 * time.Millisecond),
			RetryMode:             aws.RetryMode("standard"),
			TLSNegotiationTimeout: aws.Duration(1100 * time.Millisecond),
		}
		return settings, nil
	case aws.DefaultsModeMobile:
		settings := Configuration{
			ConnectTimeout:        aws.Duration(30000 * time.Millisecond),
			RetryMode:             aws.RetryMode("standard"),
			TLSNegotiationTimeout: aws.Duration(30000 * time.Millisecond),
		}
		return settings, nil
	case aws.DefaultsModeStandard:
		settings := Configuration{
			ConnectTimeout:        aws.Duration(3100 * time.Millisecond),
			RetryMode:             aws.RetryMode("standard"),
			TLSNegotiationTimeout: aws.Duration(3100 * time.Millisecond),
		}
		return settings, nil
	default:
		return Configuration{}, fmt.Errorf("unsupported defaults mode: %v", mode)
	}
}
