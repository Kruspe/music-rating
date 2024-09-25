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

// This operation is not supported by directory buckets.
//
// Lists the analytics configurations for the bucket. You can have up to 1,000
// analytics configurations per bucket.
//
// This action supports list pagination and does not return more than 100
// configurations at a time. You should always check the IsTruncated element in
// the response. If there are no more configurations to list, IsTruncated is set
// to false. If there are more configurations to list, IsTruncated is set to true,
// and there will be a value in NextContinuationToken . You use the
// NextContinuationToken value to continue the pagination of the list by passing
// the value in continuation-token in the request to GET the next page.
//
// To use this operation, you must have permissions to perform the
// s3:GetAnalyticsConfiguration action. The bucket owner has this permission by
// default. The bucket owner can grant this permission to others. For more
// information about permissions, see [Permissions Related to Bucket Subresource Operations]and [Managing Access Permissions to Your Amazon S3 Resources].
//
// For information about Amazon S3 analytics feature, see [Amazon S3 Analytics – Storage Class Analysis].
//
// The following operations are related to ListBucketAnalyticsConfigurations :
//
// [GetBucketAnalyticsConfiguration]
//
// [DeleteBucketAnalyticsConfiguration]
//
// [PutBucketAnalyticsConfiguration]
//
// [Amazon S3 Analytics – Storage Class Analysis]: https://docs.aws.amazon.com/AmazonS3/latest/dev/analytics-storage-class.html
// [DeleteBucketAnalyticsConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketAnalyticsConfiguration.html
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [GetBucketAnalyticsConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketAnalyticsConfiguration.html
// [PutBucketAnalyticsConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketAnalyticsConfiguration.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
func (c *Client) ListBucketAnalyticsConfigurations(ctx context.Context, params *ListBucketAnalyticsConfigurationsInput, optFns ...func(*Options)) (*ListBucketAnalyticsConfigurationsOutput, error) {
	if params == nil {
		params = &ListBucketAnalyticsConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBucketAnalyticsConfigurations", params, optFns, c.addOperationListBucketAnalyticsConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBucketAnalyticsConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBucketAnalyticsConfigurationsInput struct {

	// The name of the bucket from which analytics configurations are retrieved.
	//
	// This member is required.
	Bucket *string

	// The ContinuationToken that represents a placeholder from where this request
	// should begin.
	ContinuationToken *string

	// The account ID of the expected bucket owner. If the account ID that you provide
	// does not match the actual owner of the bucket, the request fails with the HTTP
	// status code 403 Forbidden (access denied).
	ExpectedBucketOwner *string

	noSmithyDocumentSerde
}

func (in *ListBucketAnalyticsConfigurationsInput) bindEndpointParams(p *EndpointParameters) {

	p.Bucket = in.Bucket
	p.UseS3ExpressControlEndpoint = ptr.Bool(true)
}

type ListBucketAnalyticsConfigurationsOutput struct {

	// The list of analytics configurations for a bucket.
	AnalyticsConfigurationList []types.AnalyticsConfiguration

	// The marker that is used as a starting point for this analytics configuration
	// list response. This value is present if it was sent in the request.
	ContinuationToken *string

	// Indicates whether the returned list of analytics configurations is complete. A
	// value of true indicates that the list is not complete and the
	// NextContinuationToken will be provided for a subsequent request.
	IsTruncated *bool

	// NextContinuationToken is sent when isTruncated is true, which indicates that
	// there are more analytics configurations to list. The next request must include
	// this NextContinuationToken . The token is obfuscated and is not a usable value.
	NextContinuationToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBucketAnalyticsConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpListBucketAnalyticsConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListBucketAnalyticsConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListBucketAnalyticsConfigurations"); err != nil {
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
	if err = addOpListBucketAnalyticsConfigurationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBucketAnalyticsConfigurations(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addListBucketAnalyticsConfigurationsUpdateEndpoint(stack, options); err != nil {
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

func (v *ListBucketAnalyticsConfigurationsInput) bucket() (string, bool) {
	if v.Bucket == nil {
		return "", false
	}
	return *v.Bucket, true
}

func newServiceMetadataMiddleware_opListBucketAnalyticsConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListBucketAnalyticsConfigurations",
	}
}

// getListBucketAnalyticsConfigurationsBucketMember returns a pointer to string
// denoting a provided bucket member valueand a boolean indicating if the input has
// a modeled bucket name,
func getListBucketAnalyticsConfigurationsBucketMember(input interface{}) (*string, bool) {
	in := input.(*ListBucketAnalyticsConfigurationsInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addListBucketAnalyticsConfigurationsUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getListBucketAnalyticsConfigurationsBucketMember,
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
