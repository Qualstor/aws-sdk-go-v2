// Code generated by smithy-go-codegen DO NOT EDIT.

package connectparticipant

import (
	"context"
	"errors"
	"fmt"
	"github.com/Qualstor/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	internalendpoints "github.com/Qualstor/aws-sdk-go-v2/service/connectparticipant/internal/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"net/url"
	"strings"
)

// EndpointResolverOptions is the service endpoint resolver options
type EndpointResolverOptions = internalendpoints.Options

// EndpointResolver interface for resolving service endpoints.
type EndpointResolver interface {
	ResolveEndpoint(region string, options EndpointResolverOptions) (aws.Endpoint, error)
}

var _ EndpointResolver = &internalendpoints.Resolver{}

// NewDefaultEndpointResolver constructs a new service endpoint resolver
func NewDefaultEndpointResolver() *internalendpoints.Resolver {
	return internalendpoints.New()
}

// EndpointResolverFunc is a helper utility that wraps a function so it satisfies
// the EndpointResolver interface. This is useful when you want to add additional
// endpoint resolving logic, or stub out specific endpoints with custom values.
type EndpointResolverFunc func(region string, options EndpointResolverOptions) (aws.Endpoint, error)

func (fn EndpointResolverFunc) ResolveEndpoint(region string, options EndpointResolverOptions) (endpoint aws.Endpoint, err error) {
	return fn(region, options)
}

func resolveDefaultEndpointConfiguration(o *Options) {
	if o.EndpointResolver != nil {
		return
	}
	o.EndpointResolver = NewDefaultEndpointResolver()
}

// EndpointResolverFromURL returns an EndpointResolver configured using the
// provided endpoint url. By default, the resolved endpoint resolver uses the
// client region as signing region, and the endpoint source is set to
// EndpointSourceCustom.You can provide functional options to configure endpoint
// values for the resolved endpoint.
func EndpointResolverFromURL(url string, optFns ...func(*aws.Endpoint)) EndpointResolver {
	e := aws.Endpoint{URL: url, Source: aws.EndpointSourceCustom}
	for _, fn := range optFns {
		fn(&e)
	}

	return EndpointResolverFunc(
		func(region string, options EndpointResolverOptions) (aws.Endpoint, error) {
			if len(e.SigningRegion) == 0 {
				e.SigningRegion = region
			}
			return e, nil
		},
	)
}

type ResolveEndpoint struct {
	Resolver EndpointResolver
	Options  EndpointResolverOptions
}

func (*ResolveEndpoint) ID() string {
	return "ResolveEndpoint"
}

func (m *ResolveEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.Resolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	eo := m.Options
	eo.Logger = middleware.GetLogger(ctx)

	var endpoint aws.Endpoint
	endpoint, err = m.Resolver.ResolveEndpoint(awsmiddleware.GetRegion(ctx), eo)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL, err = url.Parse(endpoint.URL)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to parse endpoint URL: %w", err)
	}

	if len(awsmiddleware.GetSigningName(ctx)) == 0 {
		signingName := endpoint.SigningName
		if len(signingName) == 0 {
			signingName = "connect"
		}
		ctx = awsmiddleware.SetSigningName(ctx, signingName)
	}
	ctx = awsmiddleware.SetEndpointSource(ctx, endpoint.Source)
	ctx = smithyhttp.SetHostnameImmutable(ctx, endpoint.HostnameImmutable)
	ctx = awsmiddleware.SetSigningRegion(ctx, endpoint.SigningRegion)
	ctx = awsmiddleware.SetPartitionID(ctx, endpoint.PartitionID)
	return next.HandleSerialize(ctx, in)
}
func addResolveEndpointMiddleware(stack *middleware.Stack, o Options) error {
	return stack.Serialize.Insert(&ResolveEndpoint{
		Resolver: o.EndpointResolver,
		Options:  o.EndpointOptions,
	}, "OperationSerializer", middleware.Before)
}

func removeResolveEndpointMiddleware(stack *middleware.Stack) error {
	_, err := stack.Serialize.Remove((&ResolveEndpoint{}).ID())
	return err
}

type wrappedEndpointResolver struct {
	awsResolver aws.EndpointResolverWithOptions
	resolver    EndpointResolver
}

func (w *wrappedEndpointResolver) ResolveEndpoint(region string, options EndpointResolverOptions) (endpoint aws.Endpoint, err error) {
	if w.awsResolver == nil {
		goto fallback
	}
	endpoint, err = w.awsResolver.ResolveEndpoint(ServiceID, region, options)
	if err == nil {
		return endpoint, nil
	}

	if nf := (&aws.EndpointNotFoundError{}); !errors.As(err, &nf) {
		return endpoint, err
	}

fallback:
	if w.resolver == nil {
		return endpoint, fmt.Errorf("default endpoint resolver provided was nil")
	}
	return w.resolver.ResolveEndpoint(region, options)
}

type awsEndpointResolverAdaptor func(service, region string) (aws.Endpoint, error)

func (a awsEndpointResolverAdaptor) ResolveEndpoint(service, region string, options ...interface{}) (aws.Endpoint, error) {
	return a(service, region)
}

var _ aws.EndpointResolverWithOptions = awsEndpointResolverAdaptor(nil)

// withEndpointResolver returns an EndpointResolver that first delegates endpoint resolution to the awsResolver.
// If awsResolver returns aws.EndpointNotFoundError error, the resolver will use the the provided
// fallbackResolver for resolution.
//
// fallbackResolver must not be nil
func withEndpointResolver(awsResolver aws.EndpointResolver, awsResolverWithOptions aws.EndpointResolverWithOptions, fallbackResolver EndpointResolver) EndpointResolver {
	var resolver aws.EndpointResolverWithOptions

	if awsResolverWithOptions != nil {
		resolver = awsResolverWithOptions
	} else if awsResolver != nil {
		resolver = awsEndpointResolverAdaptor(awsResolver.ResolveEndpoint)
	}

	return &wrappedEndpointResolver{
		awsResolver: resolver,
		resolver:    fallbackResolver,
	}
}

func finalizeClientEndpointResolverOptions(options *Options) {
	options.EndpointOptions.LogDeprecated = options.ClientLogMode.IsDeprecatedUsage()

	if len(options.EndpointOptions.ResolvedRegion) == 0 {
		const fipsInfix = "-fips-"
		const fipsPrefix = "fips-"
		const fipsSuffix = "-fips"

		if strings.Contains(options.Region, fipsInfix) ||
			strings.Contains(options.Region, fipsPrefix) ||
			strings.Contains(options.Region, fipsSuffix) {
			options.EndpointOptions.ResolvedRegion = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(
				options.Region, fipsInfix, "-"), fipsPrefix, ""), fipsSuffix, "")
			options.EndpointOptions.UseFIPSEndpoint = aws.FIPSEndpointStateEnabled
		}
	}

}
