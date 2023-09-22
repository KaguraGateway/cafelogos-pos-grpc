// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: orderlink/orderlink_service.proto

package orderlinkconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	common "github.com/KaguraGateway/cafelogos-grpc/pkg/common"
	orderlink "github.com/KaguraGateway/cafelogos-grpc/pkg/orderlink"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// OrderLinkServiceName is the fully-qualified name of the OrderLinkService service.
	OrderLinkServiceName = "cafelogos.orderlink.OrderLinkService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// OrderLinkServicePostOrderProcedure is the fully-qualified name of the OrderLinkService's
	// PostOrder RPC.
	OrderLinkServicePostOrderProcedure = "/cafelogos.orderlink.OrderLinkService/PostOrder"
)

// OrderLinkServiceClient is a client for the cafelogos.orderlink.OrderLinkService service.
type OrderLinkServiceClient interface {
	PostOrder(context.Context, *connect.Request[orderlink.PostOrderInput]) (*connect.Response[common.Empty], error)
}

// NewOrderLinkServiceClient constructs a client for the cafelogos.orderlink.OrderLinkService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewOrderLinkServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) OrderLinkServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &orderLinkServiceClient{
		postOrder: connect.NewClient[orderlink.PostOrderInput, common.Empty](
			httpClient,
			baseURL+OrderLinkServicePostOrderProcedure,
			opts...,
		),
	}
}

// orderLinkServiceClient implements OrderLinkServiceClient.
type orderLinkServiceClient struct {
	postOrder *connect.Client[orderlink.PostOrderInput, common.Empty]
}

// PostOrder calls cafelogos.orderlink.OrderLinkService.PostOrder.
func (c *orderLinkServiceClient) PostOrder(ctx context.Context, req *connect.Request[orderlink.PostOrderInput]) (*connect.Response[common.Empty], error) {
	return c.postOrder.CallUnary(ctx, req)
}

// OrderLinkServiceHandler is an implementation of the cafelogos.orderlink.OrderLinkService service.
type OrderLinkServiceHandler interface {
	PostOrder(context.Context, *connect.Request[orderlink.PostOrderInput]) (*connect.Response[common.Empty], error)
}

// NewOrderLinkServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewOrderLinkServiceHandler(svc OrderLinkServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	orderLinkServicePostOrderHandler := connect.NewUnaryHandler(
		OrderLinkServicePostOrderProcedure,
		svc.PostOrder,
		opts...,
	)
	return "/cafelogos.orderlink.OrderLinkService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case OrderLinkServicePostOrderProcedure:
			orderLinkServicePostOrderHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedOrderLinkServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedOrderLinkServiceHandler struct{}

func (UnimplementedOrderLinkServiceHandler) PostOrder(context.Context, *connect.Request[orderlink.PostOrderInput]) (*connect.Response[common.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cafelogos.orderlink.OrderLinkService.PostOrder is not implemented"))
}
