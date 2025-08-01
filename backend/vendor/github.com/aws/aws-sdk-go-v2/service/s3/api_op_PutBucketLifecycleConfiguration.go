// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalChecksum "github.com/aws/aws-sdk-go-v2/service/internal/checksum"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new lifecycle configuration for the bucket or replaces an existing
// lifecycle configuration. Keep in mind that this will overwrite an existing
// lifecycle configuration, so if you want to retain any configuration details,
// they must be included in the new lifecycle configuration. For information about
// lifecycle configuration, see [Managing your storage lifecycle].
//
// Bucket lifecycle configuration now supports specifying a lifecycle rule using
// an object key name prefix, one or more object tags, object size, or any
// combination of these. Accordingly, this section describes the latest API. The
// previous version of the API supported filtering based only on an object key name
// prefix, which is supported for backward compatibility. For the related API
// description, see [PutBucketLifecycle].
//
// Rules Permissions HTTP Host header syntax You specify the lifecycle
// configuration in your request body. The lifecycle configuration is specified as
// XML consisting of one or more rules. An Amazon S3 Lifecycle configuration can
// have up to 1,000 rules. This limit is not adjustable.
//
// Bucket lifecycle configuration supports specifying a lifecycle rule using an
// object key name prefix, one or more object tags, object size, or any combination
// of these. Accordingly, this section describes the latest API. The previous
// version of the API supported filtering based only on an object key name prefix,
// which is supported for backward compatibility for general purpose buckets. For
// the related API description, see [PutBucketLifecycle].
//
// Lifecyle configurations for directory buckets only support expiring objects and
// cancelling multipart uploads. Expiring of versioned objects,transitions and tag
// filters are not supported.
//
// A lifecycle rule consists of the following:
//
//   - A filter identifying a subset of objects to which the rule applies. The
//     filter can be based on a key name prefix, object tags, object size, or any
//     combination of these.
//
//   - A status indicating whether the rule is in effect.
//
//   - One or more lifecycle transition and expiration actions that you want
//     Amazon S3 to perform on the objects identified by the filter. If the state of
//     your bucket is versioning-enabled or versioning-suspended, you can have many
//     versions of the same object (one current version and zero or more noncurrent
//     versions). Amazon S3 provides predefined actions that you can specify for
//     current and noncurrent object versions.
//
// For more information, see [Object Lifecycle Management] and [Lifecycle Configuration Elements].
//
//   - General purpose bucket permissions - By default, all Amazon S3 resources
//     are private, including buckets, objects, and related subresources (for example,
//     lifecycle configuration and website configuration). Only the resource owner
//     (that is, the Amazon Web Services account that created it) can access the
//     resource. The resource owner can optionally grant access permissions to others
//     by writing an access policy. For this operation, a user must have the
//     s3:PutLifecycleConfiguration permission.
//
// You can also explicitly deny permissions. An explicit deny also supersedes any
//
//	other permissions. If you want to block users or accounts from removing or
//	deleting objects from your bucket, you must deny them permissions for the
//	following actions:
//
//	- s3:DeleteObject
//
//	- s3:DeleteObjectVersion
//
//	- s3:PutLifecycleConfiguration
//
// For more information about permissions, see [Managing Access Permissions to Your Amazon S3 Resources].
//
//   - Directory bucket permissions - You must have the
//     s3express:PutLifecycleConfiguration permission in an IAM identity-based policy
//     to use this operation. Cross-account access to this API operation isn't
//     supported. The resource owner can optionally grant access permissions to others
//     by creating a role or user for them as long as they are within the same account
//     as the owner and resource.
//
// For more information about directory bucket policies and permissions, see [Authorizing Regional endpoint APIs with IAM]in
//
//	the Amazon S3 User Guide.
//
// Directory buckets - For directory buckets, you must make requests for this API
//
//	operation to the Regional endpoint. These endpoints support path-style requests
//	in the format https://s3express-control.region-code.amazonaws.com/bucket-name
//	. Virtual-hosted-style requests aren't supported. For more information about
//	endpoints in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more
//	information about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// Directory buckets - The HTTP Host header syntax is
// s3express-control.region.amazonaws.com .
//
// The following operations are related to PutBucketLifecycleConfiguration :
//
// [GetBucketLifecycleConfiguration]
//
// [DeleteBucketLifecycle]
//
// [Object Lifecycle Management]: https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html
// [Lifecycle Configuration Elements]: https://docs.aws.amazon.com/AmazonS3/latest/dev/intro-lifecycle-rules.html
// [GetBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLifecycleConfiguration.html
// [Authorizing Regional endpoint APIs with IAM]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam.html
// [PutBucketLifecycle]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketLifecycle.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [DeleteBucketLifecycle]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketLifecycle.html
// [Managing your storage lifecycle]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lifecycle-mgmt.html
//
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
func (c *Client) PutBucketLifecycleConfiguration(ctx context.Context, params *PutBucketLifecycleConfigurationInput, optFns ...func(*Options)) (*PutBucketLifecycleConfigurationOutput, error) {
	if params == nil {
		params = &PutBucketLifecycleConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBucketLifecycleConfiguration", params, optFns, c.addOperationPutBucketLifecycleConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBucketLifecycleConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketLifecycleConfigurationInput struct {

	// The name of the bucket for which to set the configuration.
	//
	// This member is required.
	Bucket *string

	// Indicates the algorithm used to create the checksum for the request when you
	// use the SDK. This header will not provide any additional functionality if you
	// don't use the SDK. When you send this header, there must be a corresponding
	// x-amz-checksum or x-amz-trailer header sent. Otherwise, Amazon S3 fails the
	// request with the HTTP status code 400 Bad Request . For more information, see [Checking object integrity]
	// in the Amazon S3 User Guide.
	//
	// If you provide an individual checksum, Amazon S3 ignores any provided
	// ChecksumAlgorithm parameter.
	//
	// [Checking object integrity]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/checking-object-integrity.html
	ChecksumAlgorithm types.ChecksumAlgorithm

	// The account ID of the expected bucket owner. If the account ID that you provide
	// does not match the actual owner of the bucket, the request fails with the HTTP
	// status code 403 Forbidden (access denied).
	//
	// This parameter applies to general purpose buckets only. It is not supported for
	// directory bucket lifecycle configurations.
	ExpectedBucketOwner *string

	// Container for lifecycle rules. You can add as many as 1,000 rules.
	LifecycleConfiguration *types.BucketLifecycleConfiguration

	// Indicates which default minimum object size behavior is applied to the
	// lifecycle configuration.
	//
	// This parameter applies to general purpose buckets only. It is not supported for
	// directory bucket lifecycle configurations.
	//
	//   - all_storage_classes_128K - Objects smaller than 128 KB will not transition
	//   to any storage class by default.
	//
	//   - varies_by_storage_class - Objects smaller than 128 KB will transition to
	//   Glacier Flexible Retrieval or Glacier Deep Archive storage classes. By default,
	//   all other storage classes will prevent transitions smaller than 128 KB.
	//
	// To customize the minimum object size for any transition you can add a filter
	// that specifies a custom ObjectSizeGreaterThan or ObjectSizeLessThan in the body
	// of your transition rule. Custom filters always take precedence over the default
	// transition behavior.
	TransitionDefaultMinimumObjectSize types.TransitionDefaultMinimumObjectSize

	noSmithyDocumentSerde
}

func (in *PutBucketLifecycleConfigurationInput) bindEndpointParams(p *EndpointParameters) {

	p.Bucket = in.Bucket
	p.UseS3ExpressControlEndpoint = ptr.Bool(true)
}

type PutBucketLifecycleConfigurationOutput struct {

	// Indicates which default minimum object size behavior is applied to the
	// lifecycle configuration.
	//
	// This parameter applies to general purpose buckets only. It is not supported for
	// directory bucket lifecycle configurations.
	//
	//   - all_storage_classes_128K - Objects smaller than 128 KB will not transition
	//   to any storage class by default.
	//
	//   - varies_by_storage_class - Objects smaller than 128 KB will transition to
	//   Glacier Flexible Retrieval or Glacier Deep Archive storage classes. By default,
	//   all other storage classes will prevent transitions smaller than 128 KB.
	//
	// To customize the minimum object size for any transition you can add a filter
	// that specifies a custom ObjectSizeGreaterThan or ObjectSizeLessThan in the body
	// of your transition rule. Custom filters always take precedence over the default
	// transition behavior.
	TransitionDefaultMinimumObjectSize types.TransitionDefaultMinimumObjectSize

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutBucketLifecycleConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutBucketLifecycleConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketLifecycleConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutBucketLifecycleConfiguration"); err != nil {
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
	if err = addRequestChecksumMetricsTracking(stack, options); err != nil {
		return err
	}
	if err = addCredentialSource(stack, options); err != nil {
		return err
	}
	if err = addOpPutBucketLifecycleConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketLifecycleConfiguration(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addPutBucketLifecycleConfigurationInputChecksumMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addPutBucketLifecycleConfigurationUpdateEndpoint(stack, options); err != nil {
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
	if err = s3cust.AddExpressDefaultChecksumMiddleware(stack); err != nil {
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

func (v *PutBucketLifecycleConfigurationInput) bucket() (string, bool) {
	if v.Bucket == nil {
		return "", false
	}
	return *v.Bucket, true
}

func newServiceMetadataMiddleware_opPutBucketLifecycleConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutBucketLifecycleConfiguration",
	}
}

// getPutBucketLifecycleConfigurationRequestAlgorithmMember gets the request
// checksum algorithm value provided as input.
func getPutBucketLifecycleConfigurationRequestAlgorithmMember(input interface{}) (string, bool) {
	in := input.(*PutBucketLifecycleConfigurationInput)
	if len(in.ChecksumAlgorithm) == 0 {
		return "", false
	}
	return string(in.ChecksumAlgorithm), true
}

func addPutBucketLifecycleConfigurationInputChecksumMiddlewares(stack *middleware.Stack, options Options) error {
	return addInputChecksumMiddleware(stack, internalChecksum.InputMiddlewareOptions{
		GetAlgorithm:                     getPutBucketLifecycleConfigurationRequestAlgorithmMember,
		RequireChecksum:                  true,
		RequestChecksumCalculation:       options.RequestChecksumCalculation,
		EnableTrailingChecksum:           false,
		EnableComputeSHA256PayloadHash:   true,
		EnableDecodedContentLengthHeader: true,
	})
}

// getPutBucketLifecycleConfigurationBucketMember returns a pointer to string
// denoting a provided bucket member valueand a boolean indicating if the input has
// a modeled bucket name,
func getPutBucketLifecycleConfigurationBucketMember(input interface{}) (*string, bool) {
	in := input.(*PutBucketLifecycleConfigurationInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addPutBucketLifecycleConfigurationUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getPutBucketLifecycleConfigurationBucketMember,
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
