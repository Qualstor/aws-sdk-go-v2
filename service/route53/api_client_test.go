// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	"github.com/Qualstor/aws-sdk-go-v2/aws"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestClient_resolveRetryOptions(t *testing.T) {
	nopClient := smithyhttp.ClientDoFunc(func(_ *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Header:     http.Header{},
			Body:       ioutil.NopCloser(strings.NewReader("")),
		}, nil
	})

	cases := map[string]struct {
		defaultsMode            aws.DefaultsMode
		retryer                 aws.Retryer
		retryMaxAttempts        int
		opRetryMaxAttempts      *int
		retryMode               aws.RetryMode
		expectClientRetryMode   aws.RetryMode
		expectClientMaxAttempts int
		expectOpMaxAttempts     int
	}{
		"defaults": {
			defaultsMode:            aws.DefaultsModeStandard,
			expectClientRetryMode:   aws.RetryModeStandard,
			expectClientMaxAttempts: 3,
			expectOpMaxAttempts:     3,
		},
		"custom default retry": {
			retryMode:               aws.RetryModeAdaptive,
			retryMaxAttempts:        10,
			expectClientRetryMode:   aws.RetryModeAdaptive,
			expectClientMaxAttempts: 10,
			expectOpMaxAttempts:     10,
		},
		"custom op max attempts": {
			retryMode:               aws.RetryModeAdaptive,
			retryMaxAttempts:        10,
			opRetryMaxAttempts:      aws.Int(2),
			expectClientRetryMode:   aws.RetryModeAdaptive,
			expectClientMaxAttempts: 10,
			expectOpMaxAttempts:     2,
		},
		"custom op no change max attempts": {
			retryMode:               aws.RetryModeAdaptive,
			retryMaxAttempts:        10,
			opRetryMaxAttempts:      aws.Int(10),
			expectClientRetryMode:   aws.RetryModeAdaptive,
			expectClientMaxAttempts: 10,
			expectOpMaxAttempts:     10,
		},
		"custom op 0 max attempts": {
			retryMode:               aws.RetryModeAdaptive,
			retryMaxAttempts:        10,
			opRetryMaxAttempts:      aws.Int(0),
			expectClientRetryMode:   aws.RetryModeAdaptive,
			expectClientMaxAttempts: 10,
			expectOpMaxAttempts:     10,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			client := NewFromConfig(aws.Config{
				DefaultsMode: c.defaultsMode,
				Retryer: func() func() aws.Retryer {
					if c.retryer == nil {
						return nil
					}

					return func() aws.Retryer { return c.retryer }
				}(),
				HTTPClient:       nopClient,
				RetryMaxAttempts: c.retryMaxAttempts,
				RetryMode:        c.retryMode,
			})

			if e, a := c.expectClientRetryMode, client.options.RetryMode; e != a {
				t.Errorf("expect %v retry mode, got %v", e, a)
			}
			if e, a := c.expectClientMaxAttempts, client.options.Retryer.MaxAttempts(); e != a {
				t.Errorf("expect %v max attempts, got %v", e, a)
			}

			_, _, err := client.invokeOperation(context.Background(), "mockOperation", struct{}{},
				[]func(*Options){
					func(o *Options) {
						if c.opRetryMaxAttempts == nil {
							return
						}
						o.RetryMaxAttempts = *c.opRetryMaxAttempts
					},
				},
				func(s *middleware.Stack, o Options) error {
					s.Initialize.Clear()
					s.Serialize.Clear()
					s.Build.Clear()
					s.Finalize.Clear()
					s.Deserialize.Clear()

					if e, a := c.expectOpMaxAttempts, o.Retryer.MaxAttempts(); e != a {
						t.Errorf("expect %v op max attempts, got %v", e, a)
					}
					return nil
				})
			if err != nil {
				t.Fatalf("expect no operation error, got %v", err)
			}
		})
	}
}
