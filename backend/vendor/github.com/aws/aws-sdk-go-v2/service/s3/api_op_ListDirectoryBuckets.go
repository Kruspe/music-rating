// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of all Amazon S3 directory buckets owned by the authenticated
// sender of the request. For more information about directory buckets, see [Directory buckets]in the
// Amazon S3 User Guide.
//
// Directory buckets - For directory buckets, you must make requests for this API
// operation to the Regional endpoint. These endpoints support path-style requests
// in the format https://s3express-control.region-code.amazonaws.com/bucket-name .
// Virtual-hosted-style requests aren't supported. For more information about
// endpoints in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more
// information about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// Permissions You must have the s3express:ListAllMyDirectoryBuckets permission in
// an IAM identity-based policy instead of a bucket policy. Cross-account access to
// this API operation isn't supported. This operation can only be performed by the
// Amazon Web Services account that owns the resource. For more information about
// directory bucket policies and permissions, see [Amazon Web Services Identity and Access Management (IAM) for S3 Express One Zone]in the Amazon S3 User Guide.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// s3express-control.region.amazonaws.com .
//
// The BucketRegion response element is not part of the ListDirectoryBuckets
// Response Syntax.
//
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [Directory buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/directory-buckets-overview.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
// [Amazon Web Services Identity and Access Management (IAM) for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam.html
func (c *Client) ListDirectoryBuckets(ctx context.Context, params *ListDirectoryBucketsInput, optFns ...func(*Options)) (*ListDirectoryBucketsOutput, error) {
	if params == nil {
		params = &ListDirectoryBucketsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDirectoryBuckets", params, optFns, c.addOperationListDirectoryBucketsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDirectoryBucketsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDirectoryBucketsInput struct {

	// ContinuationToken indicates to Amazon S3 that the list is being continued on
	// buckets in this account with a token. ContinuationToken is obfuscated and is
	// not a real bucket name. You can use this ContinuationToken for the pagination
	// of the list results.
	ContinuationToken *string

	// Maximum number of buckets to be returned in response. When the number is more
	// than the count of buckets that are owned by an Amazon Web Services account,
	// return all the buckets in response.
	MaxDirectoryBuckets *int32

	noSmithyDocumentSerde
}

func (in *ListDirectoryBucketsInput) bindEndpointParams(p *EndpointParameters) {

	p.UseS3ExpressControlEndpoint = ptr.Bool(true)
}

type ListDirectoryBucketsOutput struct {

	// The list of buckets owned by the requester.
	Buckets []types.Bucket

	// If ContinuationToken was sent with the request, it is included in the response.
	// You can use the returned ContinuationToken for pagination of the list response.
	ContinuationToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDirectoryBucketsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpListDirectoryBuckets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListDirectoryBuckets{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDirectoryBuckets"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addPutBucketContextMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addIsExpressUserAgent(stack); err != nil {
		return err
	}
	if err = addCredentialSource(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDirectoryBuckets(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addListDirectoryBucketsUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSerializeImmutableHostnameBucketMiddleware(stack, options); err != nil {
		return err
	}
	if err = addInterceptBeforeRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addInterceptAttempt(stack, options); err != nil {
		return err
	}
	if err = addInterceptExecution(stack, options); err != nil {
		return err
	}
	if err = addInterceptBeforeSerialization(stack, options); err != nil {
		return err
	}
	if err = addInterceptAfterSerialization(stack, options); err != nil {
		return err
	}
	if err = addInterceptBeforeSigning(stack, options); err != nil {
		return err
	}
	if err = addInterceptAfterSigning(stack, options); err != nil {
		return err
	}
	if err = addInterceptTransmit(stack, options); err != nil {
		return err
	}
	if err = addInterceptBeforeDeserialization(stack, options); err != nil {
		return err
	}
	if err = addInterceptAfterDeserialization(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

// ListDirectoryBucketsPaginatorOptions is the paginator options for
// ListDirectoryBuckets
type ListDirectoryBucketsPaginatorOptions struct {
	// Maximum number of buckets to be returned in response. When the number is more
	// than the count of buckets that are owned by an Amazon Web Services account,
	// return all the buckets in response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDirectoryBucketsPaginator is a paginator for ListDirectoryBuckets
type ListDirectoryBucketsPaginator struct {
	options   ListDirectoryBucketsPaginatorOptions
	client    ListDirectoryBucketsAPIClient
	params    *ListDirectoryBucketsInput
	nextToken *string
	firstPage bool
}

// NewListDirectoryBucketsPaginator returns a new ListDirectoryBucketsPaginator
func NewListDirectoryBucketsPaginator(client ListDirectoryBucketsAPIClient, params *ListDirectoryBucketsInput, optFns ...func(*ListDirectoryBucketsPaginatorOptions)) *ListDirectoryBucketsPaginator {
	if params == nil {
		params = &ListDirectoryBucketsInput{}
	}

	options := ListDirectoryBucketsPaginatorOptions{}
	if params.MaxDirectoryBuckets != nil {
		options.Limit = *params.MaxDirectoryBuckets
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDirectoryBucketsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.ContinuationToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDirectoryBucketsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDirectoryBuckets page.
func (p *ListDirectoryBucketsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDirectoryBucketsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.ContinuationToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxDirectoryBuckets = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListDirectoryBuckets(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.ContinuationToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// ListDirectoryBucketsAPIClient is a client that implements the
// ListDirectoryBuckets operation.
type ListDirectoryBucketsAPIClient interface {
	ListDirectoryBuckets(context.Context, *ListDirectoryBucketsInput, ...func(*Options)) (*ListDirectoryBucketsOutput, error)
}

var _ ListDirectoryBucketsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListDirectoryBuckets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDirectoryBuckets",
	}
}

func addListDirectoryBucketsUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: nopGetBucketAccessor,
		},
		UsePathStyle:                   options.UsePathStyle,
		UseAccelerate:                  options.UseAccelerate,
		SupportsAccelerate:             true,
		TargetS3ObjectLambda:           false,
		EndpointResolver:               options.EndpointResolver,
		EndpointResolverOptions:        options.EndpointOptions,
		UseARNRegion:                   options.UseARNRegion,
		DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
	})
}
