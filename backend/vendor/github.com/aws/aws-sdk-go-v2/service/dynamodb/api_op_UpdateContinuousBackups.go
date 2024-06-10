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

// UpdateContinuousBackups enables or disables point in time recovery for the
// specified table. A successful UpdateContinuousBackups call returns the current
// ContinuousBackupsDescription . Continuous backups are ENABLED on all tables at
// table creation. If point in time recovery is enabled, PointInTimeRecoveryStatus
// will be set to ENABLED.
//
// Once continuous backups and point in time recovery are enabled, you can restore
// to any point in time within EarliestRestorableDateTime and
// LatestRestorableDateTime .
//
// LatestRestorableDateTime is typically 5 minutes before the current time. You
// can restore your table to any point in time during the last 35 days.
func (c *Client) UpdateContinuousBackups(ctx context.Context, params *UpdateContinuousBackupsInput, optFns ...func(*Options)) (*UpdateContinuousBackupsOutput, error) {
	if params == nil {
		params = &UpdateContinuousBackupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateContinuousBackups", params, optFns, c.addOperationUpdateContinuousBackupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateContinuousBackupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateContinuousBackupsInput struct {

	// Represents the settings used to enable point in time recovery.
	//
	// This member is required.
	PointInTimeRecoverySpecification *types.PointInTimeRecoverySpecification

	// The name of the table. You can also provide the Amazon Resource Name (ARN) of
	// the table in this parameter.
	//
	// This member is required.
	TableName *string

	noSmithyDocumentSerde
}

type UpdateContinuousBackupsOutput struct {

	// Represents the continuous backups and point in time recovery settings on the
	// table.
	ContinuousBackupsDescription *types.ContinuousBackupsDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateContinuousBackupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateContinuousBackups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateContinuousBackups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateContinuousBackups"); err != nil {
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
	if err = addOpUpdateContinuousBackupsDiscoverEndpointMiddleware(stack, options, c); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addOpUpdateContinuousBackupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateContinuousBackups(options.Region), middleware.Before); err != nil {
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

func addOpUpdateContinuousBackupsDiscoverEndpointMiddleware(stack *middleware.Stack, o Options, c *Client) error {
	return stack.Finalize.Insert(&internalEndpointDiscovery.DiscoverEndpoint{
		Options: []func(*internalEndpointDiscovery.DiscoverEndpointOptions){
			func(opt *internalEndpointDiscovery.DiscoverEndpointOptions) {
				opt.DisableHTTPS = o.EndpointOptions.DisableHTTPS
				opt.Logger = o.Logger
			},
		},
		DiscoverOperation:            c.fetchOpUpdateContinuousBackupsDiscoverEndpoint,
		EndpointDiscoveryEnableState: o.EndpointDiscovery.EnableEndpointDiscovery,
		EndpointDiscoveryRequired:    false,
		Region:                       o.Region,
	}, "ResolveEndpointV2", middleware.After)
}

func (c *Client) fetchOpUpdateContinuousBackupsDiscoverEndpoint(ctx context.Context, region string, optFns ...func(*internalEndpointDiscovery.DiscoverEndpointOptions)) (internalEndpointDiscovery.WeightedAddress, error) {
	input := getOperationInput(ctx)
	in, ok := input.(*UpdateContinuousBackupsInput)
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

func newServiceMetadataMiddleware_opUpdateContinuousBackups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateContinuousBackups",
	}
}
