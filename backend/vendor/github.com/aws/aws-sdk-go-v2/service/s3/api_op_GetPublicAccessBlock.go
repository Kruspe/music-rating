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

// This operation is not supported for directory buckets.
//
// Retrieves the PublicAccessBlock configuration for an Amazon S3 bucket. To use
// this operation, you must have the s3:GetBucketPublicAccessBlock permission. For
// more information about Amazon S3 permissions, see [Specifying Permissions in a Policy].
//
// When Amazon S3 evaluates the PublicAccessBlock configuration for a bucket or an
// object, it checks the PublicAccessBlock configuration for both the bucket (or
// the bucket that contains the object) and the bucket owner's account. If the
// PublicAccessBlock settings are different between the bucket and the account,
// Amazon S3 uses the most restrictive combination of the bucket-level and
// account-level settings.
//
// For more information about when Amazon S3 considers a bucket or an object
// public, see [The Meaning of "Public"].
//
// The following operations are related to GetPublicAccessBlock :
//
// [Using Amazon S3 Block Public Access]
//
// [PutPublicAccessBlock]
//
// [GetPublicAccessBlock]
//
// [DeletePublicAccessBlock]
//
// [GetPublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetPublicAccessBlock.html
// [PutPublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutPublicAccessBlock.html
// [DeletePublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeletePublicAccessBlock.html
// [Using Amazon S3 Block Public Access]: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html
// [Specifying Permissions in a Policy]: https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html
// [The Meaning of "Public"]: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status
func (c *Client) GetPublicAccessBlock(ctx context.Context, params *GetPublicAccessBlockInput, optFns ...func(*Options)) (*GetPublicAccessBlockOutput, error) {
	if params == nil {
		params = &GetPublicAccessBlockInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPublicAccessBlock", params, optFns, c.addOperationGetPublicAccessBlockMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPublicAccessBlockOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPublicAccessBlockInput struct {

	// The name of the Amazon S3 bucket whose PublicAccessBlock configuration you want
	// to retrieve.
	//
	// This member is required.
	Bucket *string

	// The account ID of the expected bucket owner. If the account ID that you provide
	// does not match the actual owner of the bucket, the request fails with the HTTP
	// status code 403 Forbidden (access denied).
	ExpectedBucketOwner *string

	noSmithyDocumentSerde
}

func (in *GetPublicAccessBlockInput) bindEndpointParams(p *EndpointParameters) {

	p.Bucket = in.Bucket
	p.UseS3ExpressControlEndpoint = ptr.Bool(true)
}

type GetPublicAccessBlockOutput struct {

	// The PublicAccessBlock configuration currently in effect for this Amazon S3
	// bucket.
	PublicAccessBlockConfiguration *types.PublicAccessBlockConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPublicAccessBlockMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetPublicAccessBlock{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetPublicAccessBlock{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetPublicAccessBlock"); err != nil {
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
	if err = addOpGetPublicAccessBlockValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPublicAccessBlock(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addGetPublicAccessBlockUpdateEndpoint(stack, options); err != nil {
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

func (v *GetPublicAccessBlockInput) bucket() (string, bool) {
	if v.Bucket == nil {
		return "", false
	}
	return *v.Bucket, true
}

func newServiceMetadataMiddleware_opGetPublicAccessBlock(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetPublicAccessBlock",
	}
}

// getGetPublicAccessBlockBucketMember returns a pointer to string denoting a
// provided bucket member valueand a boolean indicating if the input has a modeled
// bucket name,
func getGetPublicAccessBlockBucketMember(input interface{}) (*string, bool) {
	in := input.(*GetPublicAccessBlockInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addGetPublicAccessBlockUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getGetPublicAccessBlockBucketMember,
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
