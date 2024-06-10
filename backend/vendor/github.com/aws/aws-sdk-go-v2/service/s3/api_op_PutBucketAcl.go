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

// This operation is not supported by directory buckets.
//
// Sets the permissions on an existing bucket using access control lists (ACL).
// For more information, see [Using ACLs]. To set the ACL of a bucket, you must have the
// WRITE_ACP permission.
//
// You can use one of the following two ways to set a bucket's permissions:
//
//   - Specify the ACL in the request body
//
//   - Specify permissions using request headers
//
// You cannot specify access permission using both the body and the request
// headers.
//
// Depending on your application needs, you may choose to set the ACL on a bucket
// using either the request body or the headers. For example, if you have an
// existing application that updates a bucket ACL using the request body, then you
// can continue to use that approach.
//
// If your bucket uses the bucket owner enforced setting for S3 Object Ownership,
// ACLs are disabled and no longer affect permissions. You must use policies to
// grant access to your bucket and the objects in it. Requests to set ACLs or
// update ACLs fail and return the AccessControlListNotSupported error code.
// Requests to read ACLs are still supported. For more information, see [Controlling object ownership]in the
// Amazon S3 User Guide.
//
// Permissions You can set access permissions by using one of the following
// methods:
//
//   - Specify a canned ACL with the x-amz-acl request header. Amazon S3 supports a
//     set of predefined ACLs, known as canned ACLs. Each canned ACL has a predefined
//     set of grantees and permissions. Specify the canned ACL name as the value of
//     x-amz-acl . If you use this header, you cannot use other access
//     control-specific headers in your request. For more information, see [Canned ACL].
//
//   - Specify access permissions explicitly with the x-amz-grant-read ,
//     x-amz-grant-read-acp , x-amz-grant-write-acp , and x-amz-grant-full-control
//     headers. When using these headers, you specify explicit access permissions and
//     grantees (Amazon Web Services accounts or Amazon S3 groups) who will receive the
//     permission. If you use these ACL-specific headers, you cannot use the
//     x-amz-acl header to set a canned ACL. These parameters map to the set of
//     permissions that Amazon S3 supports in an ACL. For more information, see [Access Control List (ACL) Overview].
//
// You specify each grantee as a type=value pair, where the type is one of the
//
//	following:
//
//	- id – if the value specified is the canonical user ID of an Amazon Web
//	Services account
//
//	- uri – if you are granting permissions to a predefined group
//
//	- emailAddress – if the value specified is the email address of an Amazon Web
//	Services account
//
// Using email addresses to specify a grantee is only supported in the following
//
//	Amazon Web Services Regions:
//
//	- US East (N. Virginia)
//
//	- US West (N. California)
//
//	- US West (Oregon)
//
//	- Asia Pacific (Singapore)
//
//	- Asia Pacific (Sydney)
//
//	- Asia Pacific (Tokyo)
//
//	- Europe (Ireland)
//
//	- South America (São Paulo)
//
// For a list of all the Amazon S3 supported Regions and endpoints, see [Regions and Endpoints]in the
//
//	Amazon Web Services General Reference.
//
// For example, the following x-amz-grant-write header grants create, overwrite,
//
//	and delete objects permission to LogDelivery group predefined by Amazon S3 and
//	two Amazon Web Services accounts identified by their email addresses.
//
// x-amz-grant-write: uri="http://acs.amazonaws.com/groups/s3/LogDelivery",
//
//	id="111122223333", id="555566667777"
//
// You can use either a canned ACL or specify access permissions explicitly. You
// cannot do both.
//
// Grantee Values You can specify the person (grantee) to whom you're assigning
// access rights (using request elements) in the following ways:
//
//   - By the person's ID:
//
// <>ID<><>GranteesEmail<>
//
// DisplayName is optional and ignored in the request
//
//   - By URI:
//
// <>http://acs.amazonaws.com/groups/global/AuthenticatedUsers<>
//
//   - By Email address:
//
// <>Grantees@email.com<>&
//
// The grantee is resolved to the CanonicalUser and, in a response to a GET Object
//
//	acl request, appears as the CanonicalUser.
//
// Using email addresses to specify a grantee is only supported in the following
//
//	Amazon Web Services Regions:
//
//	- US East (N. Virginia)
//
//	- US West (N. California)
//
//	- US West (Oregon)
//
//	- Asia Pacific (Singapore)
//
//	- Asia Pacific (Sydney)
//
//	- Asia Pacific (Tokyo)
//
//	- Europe (Ireland)
//
//	- South America (São Paulo)
//
// For a list of all the Amazon S3 supported Regions and endpoints, see [Regions and Endpoints]in the
//
//	Amazon Web Services General Reference.
//
// The following operations are related to PutBucketAcl :
//
// [CreateBucket]
//
// [DeleteBucket]
//
// [GetObjectAcl]
//
// [Regions and Endpoints]: https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region
// [Access Control List (ACL) Overview]: https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html
// [Controlling object ownership]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/about-object-ownership.html
// [DeleteBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucket.html
// [Using ACLs]: https://docs.aws.amazon.com/AmazonS3/latest/dev/S3_ACLs_UsingACLs.html
// [Canned ACL]: https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#CannedACL
// [GetObjectAcl]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAcl.html
// [CreateBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html
func (c *Client) PutBucketAcl(ctx context.Context, params *PutBucketAclInput, optFns ...func(*Options)) (*PutBucketAclOutput, error) {
	if params == nil {
		params = &PutBucketAclInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBucketAcl", params, optFns, c.addOperationPutBucketAclMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBucketAclOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketAclInput struct {

	// The bucket to which to apply the ACL.
	//
	// This member is required.
	Bucket *string

	// The canned ACL to apply to the bucket.
	ACL types.BucketCannedACL

	// Contains the elements that set the ACL permissions for an object per grantee.
	AccessControlPolicy *types.AccessControlPolicy

	// Indicates the algorithm used to create the checksum for the object when you use
	// the SDK. This header will not provide any additional functionality if you don't
	// use the SDK. When you send this header, there must be a corresponding
	// x-amz-checksum or x-amz-trailer header sent. Otherwise, Amazon S3 fails the
	// request with the HTTP status code 400 Bad Request . For more information, see [Checking object integrity]
	// in the Amazon S3 User Guide.
	//
	// If you provide an individual checksum, Amazon S3 ignores any provided
	// ChecksumAlgorithm parameter.
	//
	// [Checking object integrity]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/checking-object-integrity.html
	ChecksumAlgorithm types.ChecksumAlgorithm

	// The base64-encoded 128-bit MD5 digest of the data. This header must be used as
	// a message integrity check to verify that the request body was not corrupted in
	// transit. For more information, go to [RFC 1864.]
	//
	// For requests made using the Amazon Web Services Command Line Interface (CLI) or
	// Amazon Web Services SDKs, this field is calculated automatically.
	//
	// [RFC 1864.]: http://www.ietf.org/rfc/rfc1864.txt
	ContentMD5 *string

	// The account ID of the expected bucket owner. If the account ID that you provide
	// does not match the actual owner of the bucket, the request fails with the HTTP
	// status code 403 Forbidden (access denied).
	ExpectedBucketOwner *string

	// Allows grantee the read, write, read ACP, and write ACP permissions on the
	// bucket.
	GrantFullControl *string

	// Allows grantee to list the objects in the bucket.
	GrantRead *string

	// Allows grantee to read the bucket ACL.
	GrantReadACP *string

	// Allows grantee to create new objects in the bucket.
	//
	// For the bucket and object owners of existing objects, also allows deletions and
	// overwrites of those objects.
	GrantWrite *string

	// Allows grantee to write the ACL for the applicable bucket.
	GrantWriteACP *string

	noSmithyDocumentSerde
}

func (in *PutBucketAclInput) bindEndpointParams(p *EndpointParameters) {
	p.Bucket = in.Bucket
	p.UseS3ExpressControlEndpoint = ptr.Bool(true)
}

type PutBucketAclOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutBucketAclMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutBucketAcl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketAcl{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutBucketAcl"); err != nil {
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
	if err = addOpPutBucketAclValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketAcl(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addPutBucketAclInputChecksumMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addPutBucketAclUpdateEndpoint(stack, options); err != nil {
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
	return nil
}

func (v *PutBucketAclInput) bucket() (string, bool) {
	if v.Bucket == nil {
		return "", false
	}
	return *v.Bucket, true
}

func newServiceMetadataMiddleware_opPutBucketAcl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutBucketAcl",
	}
}

// getPutBucketAclRequestAlgorithmMember gets the request checksum algorithm value
// provided as input.
func getPutBucketAclRequestAlgorithmMember(input interface{}) (string, bool) {
	in := input.(*PutBucketAclInput)
	if len(in.ChecksumAlgorithm) == 0 {
		return "", false
	}
	return string(in.ChecksumAlgorithm), true
}

func addPutBucketAclInputChecksumMiddlewares(stack *middleware.Stack, options Options) error {
	return internalChecksum.AddInputMiddleware(stack, internalChecksum.InputMiddlewareOptions{
		GetAlgorithm:                     getPutBucketAclRequestAlgorithmMember,
		RequireChecksum:                  true,
		EnableTrailingChecksum:           false,
		EnableComputeSHA256PayloadHash:   true,
		EnableDecodedContentLengthHeader: true,
	})
}

// getPutBucketAclBucketMember returns a pointer to string denoting a provided
// bucket member valueand a boolean indicating if the input has a modeled bucket
// name,
func getPutBucketAclBucketMember(input interface{}) (*string, bool) {
	in := input.(*PutBucketAclInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addPutBucketAclUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getPutBucketAclBucketMember,
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
