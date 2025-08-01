// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	internalEndpointDiscovery "github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all global tables that have a replica in the specified Region.
//
// This documentation is for version 2017.11.29 (Legacy) of global tables, which
// should be avoided for new global tables. Customers should use [Global Tables version 2019.11.21 (Current)]when possible,
// because it provides greater flexibility, higher efficiency, and consumes less
// write capacity than 2017.11.29 (Legacy).
//
// To determine which version you're using, see [Determining the global table version you are using]. To update existing global tables
// from version 2017.11.29 (Legacy) to version 2019.11.21 (Current), see [Upgrading global tables].
//
// [Global Tables version 2019.11.21 (Current)]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GlobalTables.html
// [Upgrading global tables]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/V2globaltables_upgrade.html
// [Determining the global table version you are using]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.DetermineVersion.html
func (c *Client) ListGlobalTables(ctx context.Context, params *ListGlobalTablesInput, optFns ...func(*Options)) (*ListGlobalTablesOutput, error) {
	if params == nil {
		params = &ListGlobalTablesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGlobalTables", params, optFns, c.addOperationListGlobalTablesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGlobalTablesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGlobalTablesInput struct {

	// The first global table name that this operation will evaluate.
	ExclusiveStartGlobalTableName *string

	// The maximum number of table names to return, if the parameter is not specified
	// DynamoDB defaults to 100.
	//
	// If the number of global tables DynamoDB finds reaches this limit, it stops the
	// operation and returns the table names collected up to that point, with a table
	// name in the LastEvaluatedGlobalTableName to apply in a subsequent operation to
	// the ExclusiveStartGlobalTableName parameter.
	Limit *int32

	// Lists the global tables in a specific Region.
	RegionName *string

	noSmithyDocumentSerde
}

type ListGlobalTablesOutput struct {

	// List of global table names.
	GlobalTables []types.GlobalTable

	// Last evaluated global table name.
	LastEvaluatedGlobalTableName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListGlobalTablesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListGlobalTables{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListGlobalTables{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListGlobalTables"); err != nil {
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
	if err = addOpListGlobalTablesDiscoverEndpointMiddleware(stack, options, c); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addUserAgentAccountIDEndpointMode(stack, options); err != nil {
		return err
	}
	if err = addCredentialSource(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListGlobalTables(options.Region), middleware.Before); err != nil {
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

func addOpListGlobalTablesDiscoverEndpointMiddleware(stack *middleware.Stack, o Options, c *Client) error {
	return stack.Finalize.Insert(&internalEndpointDiscovery.DiscoverEndpoint{
		Options: []func(*internalEndpointDiscovery.DiscoverEndpointOptions){
			func(opt *internalEndpointDiscovery.DiscoverEndpointOptions) {
				opt.DisableHTTPS = o.EndpointOptions.DisableHTTPS
				opt.Logger = o.Logger
			},
		},
		DiscoverOperation:            c.fetchOpListGlobalTablesDiscoverEndpoint,
		EndpointDiscoveryEnableState: o.EndpointDiscovery.EnableEndpointDiscovery,
		EndpointDiscoveryRequired:    false,
		Region:                       o.Region,
	}, "ResolveEndpointV2", middleware.After)
}

func (c *Client) fetchOpListGlobalTablesDiscoverEndpoint(ctx context.Context, region string, optFns ...func(*internalEndpointDiscovery.DiscoverEndpointOptions)) (internalEndpointDiscovery.WeightedAddress, error) {
	input := getOperationInput(ctx)
	in, ok := input.(*ListGlobalTablesInput)
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

func newServiceMetadataMiddleware_opListGlobalTables(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListGlobalTables",
	}
}
