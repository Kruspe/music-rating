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

// The BatchWriteItem operation puts or deletes multiple items in one or more
// tables. A single call to BatchWriteItem can transmit up to 16MB of data over
// the network, consisting of up to 25 item put or delete operations. While
// individual items can be up to 400 KB once stored, it's important to note that an
// item's representation might be greater than 400KB while being sent in DynamoDB's
// JSON format for the API call. For more details on this distinction, see [Naming Rules and Data Types].
//
// BatchWriteItem cannot update items. If you perform a BatchWriteItem operation
// on an existing item, that item's values will be overwritten by the operation and
// it will appear like it was updated. To update items, we recommend you use the
// UpdateItem action.
//
// The individual PutItem and DeleteItem operations specified in BatchWriteItem
// are atomic; however BatchWriteItem as a whole is not. If any requested
// operations fail because the table's provisioned throughput is exceeded or an
// internal processing failure occurs, the failed operations are returned in the
// UnprocessedItems response parameter. You can investigate and optionally resend
// the requests. Typically, you would call BatchWriteItem in a loop. Each
// iteration would check for unprocessed items and submit a new BatchWriteItem
// request with those unprocessed items until all items have been processed.
//
// For tables and indexes with provisioned capacity, if none of the items can be
// processed due to insufficient provisioned throughput on all of the tables in the
// request, then BatchWriteItem returns a ProvisionedThroughputExceededException .
// For all tables and indexes, if none of the items can be processed due to other
// throttling scenarios (such as exceeding partition level limits), then
// BatchWriteItem returns a ThrottlingException .
//
// If DynamoDB returns any unprocessed items, you should retry the batch operation
// on those items. However, we strongly recommend that you use an exponential
// backoff algorithm. If you retry the batch operation immediately, the underlying
// read or write requests can still fail due to throttling on the individual
// tables. If you delay the batch operation using exponential backoff, the
// individual requests in the batch are much more likely to succeed.
//
// For more information, see [Batch Operations and Error Handling] in the Amazon DynamoDB Developer Guide.
//
// With BatchWriteItem , you can efficiently write or delete large amounts of data,
// such as from Amazon EMR, or copy data from another database into DynamoDB. In
// order to improve performance with these large-scale operations, BatchWriteItem
// does not behave in the same way as individual PutItem and DeleteItem calls
// would. For example, you cannot specify conditions on individual put and delete
// requests, and BatchWriteItem does not return deleted items in the response.
//
// If you use a programming language that supports concurrency, you can use
// threads to write items in parallel. Your application must include the necessary
// logic to manage the threads. With languages that don't support threading, you
// must update or delete the specified items one at a time. In both situations,
// BatchWriteItem performs the specified put and delete operations in parallel,
// giving you the power of the thread pool approach without having to introduce
// complexity into your application.
//
// Parallel processing reduces latency, but each specified put and delete request
// consumes the same number of write capacity units whether it is processed in
// parallel or not. Delete operations on nonexistent items consume one write
// capacity unit.
//
// If one or more of the following is true, DynamoDB rejects the entire batch
// write operation:
//
//   - One or more tables specified in the BatchWriteItem request does not exist.
//
//   - Primary key attributes specified on an item in the request do not match
//     those in the corresponding table's primary key schema.
//
//   - You try to perform multiple operations on the same item in the same
//     BatchWriteItem request. For example, you cannot put and delete the same item
//     in the same BatchWriteItem request.
//
//   - Your request contains at least two items with identical hash and range keys
//     (which essentially is two put operations).
//
//   - There are more than 25 requests in the batch.
//
//   - Any individual item in a batch exceeds 400 KB.
//
//   - The total request size exceeds 16 MB.
//
//   - Any individual items with keys exceeding the key length limits. For a
//     partition key, the limit is 2048 bytes and for a sort key, the limit is 1024
//     bytes.
//
// [Batch Operations and Error Handling]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ErrorHandling.html#Programming.Errors.BatchOperations
// [Naming Rules and Data Types]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.NamingRulesDataTypes.html
func (c *Client) BatchWriteItem(ctx context.Context, params *BatchWriteItemInput, optFns ...func(*Options)) (*BatchWriteItemOutput, error) {
	if params == nil {
		params = &BatchWriteItemInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchWriteItem", params, optFns, c.addOperationBatchWriteItemMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchWriteItemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a BatchWriteItem operation.
type BatchWriteItemInput struct {

	// A map of one or more table names or table ARNs and, for each table, a list of
	// operations to be performed ( DeleteRequest or PutRequest ). Each element in the
	// map consists of the following:
	//
	//   - DeleteRequest - Perform a DeleteItem operation on the specified item. The
	//   item to be deleted is identified by a Key subelement:
	//
	//   - Key - A map of primary key attribute values that uniquely identify the item.
	//   Each entry in this map consists of an attribute name and an attribute value. For
	//   each primary key, you must provide all of the key attributes. For example, with
	//   a simple primary key, you only need to provide a value for the partition key.
	//   For a composite primary key, you must provide values for both the partition key
	//   and the sort key.
	//
	//   - PutRequest - Perform a PutItem operation on the specified item. The item to
	//   be put is identified by an Item subelement:
	//
	//   - Item - A map of attributes and their values. Each entry in this map consists
	//   of an attribute name and an attribute value. Attribute values must not be null;
	//   string and binary type attributes must have lengths greater than zero; and set
	//   type attributes must not be empty. Requests that contain empty values are
	//   rejected with a ValidationException exception.
	//
	// If you specify any attributes that are part of an index key, then the data
	//   types for those attributes must match those of the schema in the table's
	//   attribute definition.
	//
	// This member is required.
	RequestItems map[string][]types.WriteRequest

	// Determines the level of detail about either provisioned or on-demand throughput
	// consumption that is returned in the response:
	//
	//   - INDEXES - The response includes the aggregate ConsumedCapacity for the
	//   operation, together with ConsumedCapacity for each table and secondary index
	//   that was accessed.
	//
	// Note that some operations, such as GetItem and BatchGetItem , do not access any
	//   indexes at all. In these cases, specifying INDEXES will only return
	//   ConsumedCapacity information for table(s).
	//
	//   - TOTAL - The response includes only the aggregate ConsumedCapacity for the
	//   operation.
	//
	//   - NONE - No ConsumedCapacity details are included in the response.
	ReturnConsumedCapacity types.ReturnConsumedCapacity

	// Determines whether item collection metrics are returned. If set to SIZE , the
	// response includes statistics about item collections, if any, that were modified
	// during the operation are returned in the response. If set to NONE (the
	// default), no statistics are returned.
	ReturnItemCollectionMetrics types.ReturnItemCollectionMetrics

	noSmithyDocumentSerde
}

func (in *BatchWriteItemInput) bindEndpointParams(p *EndpointParameters) {
	func() {
		v1 := in.RequestItems
		var v2 []string
		for k := range v1 {
			v2 = append(v2, k)
		}
		p.ResourceArnList = v2
	}()

}

// Represents the output of a BatchWriteItem operation.
type BatchWriteItemOutput struct {

	// The capacity units consumed by the entire BatchWriteItem operation.
	//
	// Each element consists of:
	//
	//   - TableName - The table that consumed the provisioned throughput.
	//
	//   - CapacityUnits - The total number of capacity units consumed.
	ConsumedCapacity []types.ConsumedCapacity

	// A list of tables that were processed by BatchWriteItem and, for each table,
	// information about any item collections that were affected by individual
	// DeleteItem or PutItem operations.
	//
	// Each entry consists of the following subelements:
	//
	//   - ItemCollectionKey - The partition key value of the item collection. This is
	//   the same as the partition key value of the item.
	//
	//   - SizeEstimateRangeGB - An estimate of item collection size, expressed in GB.
	//   This is a two-element array containing a lower bound and an upper bound for the
	//   estimate. The estimate includes the size of all the items in the table, plus the
	//   size of all attributes projected into all of the local secondary indexes on the
	//   table. Use this estimate to measure whether a local secondary index is
	//   approaching its size limit.
	//
	// The estimate is subject to change over time; therefore, do not rely on the
	//   precision or accuracy of the estimate.
	ItemCollectionMetrics map[string][]types.ItemCollectionMetrics

	// A map of tables and requests against those tables that were not processed. The
	// UnprocessedItems value is in the same form as RequestItems , so you can provide
	// this value directly to a subsequent BatchWriteItem operation. For more
	// information, see RequestItems in the Request Parameters section.
	//
	// Each UnprocessedItems entry consists of a table name or table ARN and, for that
	// table, a list of operations to perform ( DeleteRequest or PutRequest ).
	//
	//   - DeleteRequest - Perform a DeleteItem operation on the specified item. The
	//   item to be deleted is identified by a Key subelement:
	//
	//   - Key - A map of primary key attribute values that uniquely identify the item.
	//   Each entry in this map consists of an attribute name and an attribute value.
	//
	//   - PutRequest - Perform a PutItem operation on the specified item. The item to
	//   be put is identified by an Item subelement:
	//
	//   - Item - A map of attributes and their values. Each entry in this map consists
	//   of an attribute name and an attribute value. Attribute values must not be null;
	//   string and binary type attributes must have lengths greater than zero; and set
	//   type attributes must not be empty. Requests that contain empty values will be
	//   rejected with a ValidationException exception.
	//
	// If you specify any attributes that are part of an index key, then the data
	//   types for those attributes must match those of the schema in the table's
	//   attribute definition.
	//
	// If there are no unprocessed items remaining, the response contains an empty
	// UnprocessedItems map.
	UnprocessedItems map[string][]types.WriteRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchWriteItemMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpBatchWriteItem{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpBatchWriteItem{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchWriteItem"); err != nil {
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
	if err = addOpBatchWriteItemDiscoverEndpointMiddleware(stack, options, c); err != nil {
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
	if err = addOpBatchWriteItemValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchWriteItem(options.Region), middleware.Before); err != nil {
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

func addOpBatchWriteItemDiscoverEndpointMiddleware(stack *middleware.Stack, o Options, c *Client) error {
	return stack.Finalize.Insert(&internalEndpointDiscovery.DiscoverEndpoint{
		Options: []func(*internalEndpointDiscovery.DiscoverEndpointOptions){
			func(opt *internalEndpointDiscovery.DiscoverEndpointOptions) {
				opt.DisableHTTPS = o.EndpointOptions.DisableHTTPS
				opt.Logger = o.Logger
			},
		},
		DiscoverOperation:            c.fetchOpBatchWriteItemDiscoverEndpoint,
		EndpointDiscoveryEnableState: o.EndpointDiscovery.EnableEndpointDiscovery,
		EndpointDiscoveryRequired:    false,
		Region:                       o.Region,
	}, "ResolveEndpointV2", middleware.After)
}

func (c *Client) fetchOpBatchWriteItemDiscoverEndpoint(ctx context.Context, region string, optFns ...func(*internalEndpointDiscovery.DiscoverEndpointOptions)) (internalEndpointDiscovery.WeightedAddress, error) {
	input := getOperationInput(ctx)
	in, ok := input.(*BatchWriteItemInput)
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

func newServiceMetadataMiddleware_opBatchWriteItem(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchWriteItem",
	}
}
