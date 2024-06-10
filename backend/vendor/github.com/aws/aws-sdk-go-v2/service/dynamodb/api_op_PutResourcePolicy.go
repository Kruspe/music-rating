// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	internalEndpointDiscovery "github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Attaches a resource-based policy document to the resource, which can be a table
// or stream. When you attach a resource-based policy using this API, the policy
// application is [eventually consistent].
//
// PutResourcePolicy is an idempotent operation; running it multiple times on the
// same resource using the same policy document will return the same revision ID.
// If you specify an ExpectedRevisionId that doesn't match the current policy's
// RevisionId , the PolicyNotFoundException will be returned.
//
// PutResourcePolicy is an asynchronous operation. If you issue a GetResourcePolicy
// request immediately after a PutResourcePolicy request, DynamoDB might return
// your previous policy, if there was one, or return the PolicyNotFoundException .
// This is because GetResourcePolicy uses an eventually consistent query, and the
// metadata for your policy or table might not be available at that moment. Wait
// for a few seconds, and then try the GetResourcePolicy request again.
//
// [eventually consistent]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadConsistency.html
func (c *Client) PutResourcePolicy(ctx context.Context, params *PutResourcePolicyInput, optFns ...func(*Options)) (*PutResourcePolicyOutput, error) {
	if params == nil {
		params = &PutResourcePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutResourcePolicy", params, optFns, c.addOperationPutResourcePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutResourcePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutResourcePolicyInput struct {

	// An Amazon Web Services resource-based policy document in JSON format.
	//
	//   - The maximum size supported for a resource-based policy document is 20 KB.
	//   DynamoDB counts whitespaces when calculating the size of a policy against this
	//   limit.
	//
	//   - Within a resource-based policy, if the action for a DynamoDB service-linked
	//   role (SLR) to replicate data for a global table is denied, adding or deleting a
	//   replica will fail with an error.
	//
	// For a full list of all considerations that apply while attaching a
	// resource-based policy, see [Resource-based policy considerations].
	//
	// [Resource-based policy considerations]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/rbac-considerations.html
	//
	// This member is required.
	Policy *string

	// The Amazon Resource Name (ARN) of the DynamoDB resource to which the policy
	// will be attached. The resources you can specify include tables and streams.
	//
	// You can control index permissions using the base table's policy. To specify the
	// same permission level for your table and its indexes, you can provide both the
	// table and index Amazon Resource Name (ARN)s in the Resource field of a given
	// Statement in your policy document. Alternatively, to specify different
	// permissions for your table, indexes, or both, you can define multiple Statement
	// fields in your policy document.
	//
	// This member is required.
	ResourceArn *string

	// Set this parameter to true to confirm that you want to remove your permissions
	// to change the policy of this resource in the future.
	ConfirmRemoveSelfResourceAccess bool

	// A string value that you can use to conditionally update your policy. You can
	// provide the revision ID of your existing policy to make mutating requests
	// against that policy.
	//
	// When you provide an expected revision ID, if the revision ID of the existing
	// policy on the resource doesn't match or if there's no policy attached to the
	// resource, your request will be rejected with a PolicyNotFoundException .
	//
	// To conditionally attach a policy when no policy exists for the resource,
	// specify NO_POLICY for the revision ID.
	ExpectedRevisionId *string

	noSmithyDocumentSerde
}

type PutResourcePolicyOutput struct {

	// A unique string that represents the revision ID of the policy. If you're
	// comparing revision IDs, make sure to always use string comparison logic.
	RevisionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutResourcePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpPutResourcePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpPutResourcePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutResourcePolicy"); err != nil {
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
	if err = addOpPutResourcePolicyDiscoverEndpointMiddleware(stack, options, c); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addOpPutResourcePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutResourcePolicy(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addValidateResponseChecksum(stack, options); err != nil {
		return err
	}
	if err = addAcceptEncodingGzip(stack, options); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func addOpPutResourcePolicyDiscoverEndpointMiddleware(stack *middleware.Stack, o Options, c *Client) error {
	return stack.Finalize.Insert(&internalEndpointDiscovery.DiscoverEndpoint{
		Options: []func(*internalEndpointDiscovery.DiscoverEndpointOptions){
			func(opt *internalEndpointDiscovery.DiscoverEndpointOptions) {
				opt.DisableHTTPS = o.EndpointOptions.DisableHTTPS
				opt.Logger = o.Logger
			},
		},
		DiscoverOperation:            c.fetchOpPutResourcePolicyDiscoverEndpoint,
		EndpointDiscoveryEnableState: o.EndpointDiscovery.EnableEndpointDiscovery,
		EndpointDiscoveryRequired:    false,
		Region:                       o.Region,
	}, "ResolveEndpointV2", middleware.After)
}

func (c *Client) fetchOpPutResourcePolicyDiscoverEndpoint(ctx context.Context, region string, optFns ...func(*internalEndpointDiscovery.DiscoverEndpointOptions)) (internalEndpointDiscovery.WeightedAddress, error) {
	input := getOperationInput(ctx)
	in, ok := input.(*PutResourcePolicyInput)
	if !ok {
		return internalEndpointDiscovery.WeightedAddress{}, fmt.Errorf("unknown input type %T", input)
	}
	_ = in

	identifierMap := make(map[string]string, 0)
	identifierMap["sdk#Region"] = region

	key := fmt.Sprintf("DynamoDB.%v", identifierMap)

	if v, ok := c.endpointCache.Get(key); ok {
		return v, nil
	}

	discoveryOperationInput := &DescribeEndpointsInput{}

	opt := internalEndpointDiscovery.DiscoverEndpointOptions{}
	for _, fn := range optFns {
		fn(&opt)
	}

	go c.handleEndpointDiscoveryFromService(ctx, discoveryOperationInput, region, key, opt)
	return internalEndpointDiscovery.WeightedAddress{}, nil
}

func newServiceMetadataMiddleware_opPutResourcePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutResourcePolicy",
	}
}
