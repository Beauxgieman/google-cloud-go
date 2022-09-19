// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package apikeys

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	emptypb "github.com/golang/protobuf/ptypes/empty"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	apikeyspb "google.golang.org/genproto/googleapis/api/apikeys/v2"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	CreateKey    []gax.CallOption
	ListKeys     []gax.CallOption
	GetKey       []gax.CallOption
	GetKeyString []gax.CallOption
	UpdateKey    []gax.CallOption
	DeleteKey    []gax.CallOption
	UndeleteKey  []gax.CallOption
	LookupKey    []gax.CallOption
	GetOperation []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("apikeys.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("apikeys.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://apikeys.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		CreateKey:    []gax.CallOption{},
		ListKeys:     []gax.CallOption{},
		GetKey:       []gax.CallOption{},
		GetKeyString: []gax.CallOption{},
		UpdateKey:    []gax.CallOption{},
		DeleteKey:    []gax.CallOption{},
		UndeleteKey:  []gax.CallOption{},
		LookupKey:    []gax.CallOption{},
		GetOperation: []gax.CallOption{},
	}
}

// internalClient is an interface that defines the methods available from API Keys API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateKey(context.Context, *apikeyspb.CreateKeyRequest, ...gax.CallOption) (*CreateKeyOperation, error)
	CreateKeyOperation(name string) *CreateKeyOperation
	ListKeys(context.Context, *apikeyspb.ListKeysRequest, ...gax.CallOption) *KeyIterator
	GetKey(context.Context, *apikeyspb.GetKeyRequest, ...gax.CallOption) (*apikeyspb.Key, error)
	GetKeyString(context.Context, *apikeyspb.GetKeyStringRequest, ...gax.CallOption) (*apikeyspb.GetKeyStringResponse, error)
	UpdateKey(context.Context, *apikeyspb.UpdateKeyRequest, ...gax.CallOption) (*UpdateKeyOperation, error)
	UpdateKeyOperation(name string) *UpdateKeyOperation
	DeleteKey(context.Context, *apikeyspb.DeleteKeyRequest, ...gax.CallOption) (*DeleteKeyOperation, error)
	DeleteKeyOperation(name string) *DeleteKeyOperation
	UndeleteKey(context.Context, *apikeyspb.UndeleteKeyRequest, ...gax.CallOption) (*UndeleteKeyOperation, error)
	UndeleteKeyOperation(name string) *UndeleteKeyOperation
	LookupKey(context.Context, *apikeyspb.LookupKeyRequest, ...gax.CallOption) (*apikeyspb.LookupKeyResponse, error)
	GetOperation(context.Context, *longrunningpb.GetOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
}

// Client is a client for interacting with API Keys API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Manages the API keys associated with projects.
type Client struct {
	// The internal transport-dependent client.
	internalClient internalClient

	// The call options for this service.
	CallOptions *CallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateKey creates a new API key.
//
// NOTE: Key is a global resource; hence the only supported value for
// location is global.
func (c *Client) CreateKey(ctx context.Context, req *apikeyspb.CreateKeyRequest, opts ...gax.CallOption) (*CreateKeyOperation, error) {
	return c.internalClient.CreateKey(ctx, req, opts...)
}

// CreateKeyOperation returns a new CreateKeyOperation from a given name.
// The name must be that of a previously created CreateKeyOperation, possibly from a different process.
func (c *Client) CreateKeyOperation(name string) *CreateKeyOperation {
	return c.internalClient.CreateKeyOperation(name)
}

// ListKeys lists the API keys owned by a project. The key string of the API key
// isn’t included in the response.
//
// NOTE: Key is a global resource; hence the only supported value for
// location is global.
func (c *Client) ListKeys(ctx context.Context, req *apikeyspb.ListKeysRequest, opts ...gax.CallOption) *KeyIterator {
	return c.internalClient.ListKeys(ctx, req, opts...)
}

// GetKey gets the metadata for an API key. The key string of the API key
// isn’t included in the response.
//
// NOTE: Key is a global resource; hence the only supported value for
// location is global.
func (c *Client) GetKey(ctx context.Context, req *apikeyspb.GetKeyRequest, opts ...gax.CallOption) (*apikeyspb.Key, error) {
	return c.internalClient.GetKey(ctx, req, opts...)
}

// GetKeyString get the key string for an API key.
//
// NOTE: Key is a global resource; hence the only supported value for
// location is global.
func (c *Client) GetKeyString(ctx context.Context, req *apikeyspb.GetKeyStringRequest, opts ...gax.CallOption) (*apikeyspb.GetKeyStringResponse, error) {
	return c.internalClient.GetKeyString(ctx, req, opts...)
}

// UpdateKey patches the modifiable fields of an API key.
// The key string of the API key isn’t included in the response.
//
// NOTE: Key is a global resource; hence the only supported value for
// location is global.
func (c *Client) UpdateKey(ctx context.Context, req *apikeyspb.UpdateKeyRequest, opts ...gax.CallOption) (*UpdateKeyOperation, error) {
	return c.internalClient.UpdateKey(ctx, req, opts...)
}

// UpdateKeyOperation returns a new UpdateKeyOperation from a given name.
// The name must be that of a previously created UpdateKeyOperation, possibly from a different process.
func (c *Client) UpdateKeyOperation(name string) *UpdateKeyOperation {
	return c.internalClient.UpdateKeyOperation(name)
}

// DeleteKey deletes an API key. Deleted key can be retrieved within 30 days of
// deletion. Afterward, key will be purged from the project.
//
// NOTE: Key is a global resource; hence the only supported value for
// location is global.
func (c *Client) DeleteKey(ctx context.Context, req *apikeyspb.DeleteKeyRequest, opts ...gax.CallOption) (*DeleteKeyOperation, error) {
	return c.internalClient.DeleteKey(ctx, req, opts...)
}

// DeleteKeyOperation returns a new DeleteKeyOperation from a given name.
// The name must be that of a previously created DeleteKeyOperation, possibly from a different process.
func (c *Client) DeleteKeyOperation(name string) *DeleteKeyOperation {
	return c.internalClient.DeleteKeyOperation(name)
}

// UndeleteKey undeletes an API key which was deleted within 30 days.
//
// NOTE: Key is a global resource; hence the only supported value for
// location is global.
func (c *Client) UndeleteKey(ctx context.Context, req *apikeyspb.UndeleteKeyRequest, opts ...gax.CallOption) (*UndeleteKeyOperation, error) {
	return c.internalClient.UndeleteKey(ctx, req, opts...)
}

// UndeleteKeyOperation returns a new UndeleteKeyOperation from a given name.
// The name must be that of a previously created UndeleteKeyOperation, possibly from a different process.
func (c *Client) UndeleteKeyOperation(name string) *UndeleteKeyOperation {
	return c.internalClient.UndeleteKeyOperation(name)
}

// LookupKey find the parent project and resource name of the API
// key that matches the key string in the request. If the API key has been
// purged, resource name will not be set.
// The service account must have the apikeys.keys.lookup permission
// on the parent project.
func (c *Client) LookupKey(ctx context.Context, req *apikeyspb.LookupKeyRequest, opts ...gax.CallOption) (*apikeyspb.LookupKeyResponse, error) {
	return c.internalClient.LookupKey(ctx, req, opts...)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *Client) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.GetOperation(ctx, req, opts...)
}

// gRPCClient is a client for interacting with API Keys API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type gRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions

	// The gRPC API client.
	client apikeyspb.ApiKeysClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	operationsClient longrunningpb.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewClient creates a new api keys client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Manages the API keys associated with projects.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := defaultGRPCClientOptions()
	if newClientHook != nil {
		hookOpts, err := newClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := Client{CallOptions: defaultCallOptions()}

	c := &gRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		client:           apikeyspb.NewApiKeysClient(connPool),
		CallOptions:      &client.CallOptions,
		operationsClient: longrunningpb.NewOperationsClient(connPool),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *gRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *gRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *gRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *gRPCClient) CreateKey(ctx context.Context, req *apikeyspb.CreateKeyRequest, opts ...gax.CallOption) (*CreateKeyOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateKey[0:len((*c.CallOptions).CreateKey):len((*c.CallOptions).CreateKey)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateKey(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateKeyOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *gRPCClient) ListKeys(ctx context.Context, req *apikeyspb.ListKeysRequest, opts ...gax.CallOption) *KeyIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListKeys[0:len((*c.CallOptions).ListKeys):len((*c.CallOptions).ListKeys)], opts...)
	it := &KeyIterator{}
	req = proto.Clone(req).(*apikeyspb.ListKeysRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*apikeyspb.Key, string, error) {
		resp := &apikeyspb.ListKeysResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.ListKeys(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetKeys(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *gRPCClient) GetKey(ctx context.Context, req *apikeyspb.GetKeyRequest, opts ...gax.CallOption) (*apikeyspb.Key, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetKey[0:len((*c.CallOptions).GetKey):len((*c.CallOptions).GetKey)], opts...)
	var resp *apikeyspb.Key
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetKey(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) GetKeyString(ctx context.Context, req *apikeyspb.GetKeyStringRequest, opts ...gax.CallOption) (*apikeyspb.GetKeyStringResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetKeyString[0:len((*c.CallOptions).GetKeyString):len((*c.CallOptions).GetKeyString)], opts...)
	var resp *apikeyspb.GetKeyStringResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetKeyString(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) UpdateKey(ctx context.Context, req *apikeyspb.UpdateKeyRequest, opts ...gax.CallOption) (*UpdateKeyOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "key.name", url.QueryEscape(req.GetKey().GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateKey[0:len((*c.CallOptions).UpdateKey):len((*c.CallOptions).UpdateKey)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.UpdateKey(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &UpdateKeyOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *gRPCClient) DeleteKey(ctx context.Context, req *apikeyspb.DeleteKeyRequest, opts ...gax.CallOption) (*DeleteKeyOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteKey[0:len((*c.CallOptions).DeleteKey):len((*c.CallOptions).DeleteKey)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.DeleteKey(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeleteKeyOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *gRPCClient) UndeleteKey(ctx context.Context, req *apikeyspb.UndeleteKeyRequest, opts ...gax.CallOption) (*UndeleteKeyOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UndeleteKey[0:len((*c.CallOptions).UndeleteKey):len((*c.CallOptions).UndeleteKey)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.UndeleteKey(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &UndeleteKeyOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *gRPCClient) LookupKey(ctx context.Context, req *apikeyspb.LookupKeyRequest, opts ...gax.CallOption) (*apikeyspb.LookupKeyResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).LookupKey[0:len((*c.CallOptions).LookupKey):len((*c.CallOptions).LookupKey)], opts...)
	var resp *apikeyspb.LookupKeyResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.LookupKey(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.operationsClient.GetOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateKeyOperation manages a long-running operation from CreateKey.
type CreateKeyOperation struct {
	lro *longrunning.Operation
}

// CreateKeyOperation returns a new CreateKeyOperation from a given name.
// The name must be that of a previously created CreateKeyOperation, possibly from a different process.
func (c *gRPCClient) CreateKeyOperation(name string) *CreateKeyOperation {
	return &CreateKeyOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreateKeyOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*apikeyspb.Key, error) {
	var resp apikeyspb.Key
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *CreateKeyOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*apikeyspb.Key, error) {
	var resp apikeyspb.Key
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *CreateKeyOperation) Metadata() (*emptypb.Empty, error) {
	var meta emptypb.Empty
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreateKeyOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreateKeyOperation) Name() string {
	return op.lro.Name()
}

// DeleteKeyOperation manages a long-running operation from DeleteKey.
type DeleteKeyOperation struct {
	lro *longrunning.Operation
}

// DeleteKeyOperation returns a new DeleteKeyOperation from a given name.
// The name must be that of a previously created DeleteKeyOperation, possibly from a different process.
func (c *gRPCClient) DeleteKeyOperation(name string) *DeleteKeyOperation {
	return &DeleteKeyOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *DeleteKeyOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*apikeyspb.Key, error) {
	var resp apikeyspb.Key
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *DeleteKeyOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*apikeyspb.Key, error) {
	var resp apikeyspb.Key
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *DeleteKeyOperation) Metadata() (*emptypb.Empty, error) {
	var meta emptypb.Empty
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *DeleteKeyOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *DeleteKeyOperation) Name() string {
	return op.lro.Name()
}

// UndeleteKeyOperation manages a long-running operation from UndeleteKey.
type UndeleteKeyOperation struct {
	lro *longrunning.Operation
}

// UndeleteKeyOperation returns a new UndeleteKeyOperation from a given name.
// The name must be that of a previously created UndeleteKeyOperation, possibly from a different process.
func (c *gRPCClient) UndeleteKeyOperation(name string) *UndeleteKeyOperation {
	return &UndeleteKeyOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *UndeleteKeyOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*apikeyspb.Key, error) {
	var resp apikeyspb.Key
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *UndeleteKeyOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*apikeyspb.Key, error) {
	var resp apikeyspb.Key
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *UndeleteKeyOperation) Metadata() (*emptypb.Empty, error) {
	var meta emptypb.Empty
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *UndeleteKeyOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *UndeleteKeyOperation) Name() string {
	return op.lro.Name()
}

// UpdateKeyOperation manages a long-running operation from UpdateKey.
type UpdateKeyOperation struct {
	lro *longrunning.Operation
}

// UpdateKeyOperation returns a new UpdateKeyOperation from a given name.
// The name must be that of a previously created UpdateKeyOperation, possibly from a different process.
func (c *gRPCClient) UpdateKeyOperation(name string) *UpdateKeyOperation {
	return &UpdateKeyOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *UpdateKeyOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*apikeyspb.Key, error) {
	var resp apikeyspb.Key
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *UpdateKeyOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*apikeyspb.Key, error) {
	var resp apikeyspb.Key
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *UpdateKeyOperation) Metadata() (*emptypb.Empty, error) {
	var meta emptypb.Empty
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *UpdateKeyOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *UpdateKeyOperation) Name() string {
	return op.lro.Name()
}

// KeyIterator manages a stream of *apikeyspb.Key.
type KeyIterator struct {
	items    []*apikeyspb.Key
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*apikeyspb.Key, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *KeyIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *KeyIterator) Next() (*apikeyspb.Key, error) {
	var item *apikeyspb.Key
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *KeyIterator) bufLen() int {
	return len(it.items)
}

func (it *KeyIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}