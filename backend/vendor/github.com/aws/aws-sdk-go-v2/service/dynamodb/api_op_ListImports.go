// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists completed imports within the past 90 days.
func (c *Client) ListImports(ctx context.Context, params *ListImportsInput, optFns ...func(*Options)) (*ListImportsOutput, error) {
	if params == nil {
		params = &ListImportsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListImports", params, optFns, c.addOperationListImportsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListImportsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListImportsInput struct {

	// An optional string that, if supplied, must be copied from the output of a
	// previous call to ListImports . When provided in this manner, the API fetches the
	// next page of results.
	NextToken *string

	// The number of ImportSummary objects returned in a single page.
	PageSize *int32

	// The Amazon Resource Name (ARN) associated with the table that was imported to.
	TableArn *string

	noSmithyDocumentSerde
}

type ListImportsOutput struct {

	// A list of ImportSummary objects.
	ImportSummaryList []types.ImportSummary

	// If this value is returned, there are additional results to be displayed. To
	// retrieve them, call ListImports again, with NextToken set to this value.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListImportsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListImports{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListImports{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListImports(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	return nil
}

// ListImportsAPIClient is a client that implements the ListImports operation.
type ListImportsAPIClient interface {
	ListImports(context.Context, *ListImportsInput, ...func(*Options)) (*ListImportsOutput, error)
}

var _ ListImportsAPIClient = (*Client)(nil)

// ListImportsPaginatorOptions is the paginator options for ListImports
type ListImportsPaginatorOptions struct {
	// The number of ImportSummary objects returned in a single page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListImportsPaginator is a paginator for ListImports
type ListImportsPaginator struct {
	options   ListImportsPaginatorOptions
	client    ListImportsAPIClient
	params    *ListImportsInput
	nextToken *string
	firstPage bool
}

// NewListImportsPaginator returns a new ListImportsPaginator
func NewListImportsPaginator(client ListImportsAPIClient, params *ListImportsInput, optFns ...func(*ListImportsPaginatorOptions)) *ListImportsPaginator {
	if params == nil {
		params = &ListImportsInput{}
	}

	options := ListImportsPaginatorOptions{}
	if params.PageSize != nil {
		options.Limit = *params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListImportsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListImportsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListImports page.
func (p *ListImportsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListImportsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.PageSize = limit

	result, err := p.client.ListImports(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListImports(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "ListImports",
	}
}
