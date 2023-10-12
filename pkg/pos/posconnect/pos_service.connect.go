// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: pos/pos_service.proto

package posconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	common "github.com/KaguraGateway/cafelogos-grpc/pkg/common"
	pos "github.com/KaguraGateway/cafelogos-grpc/pkg/pos"
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
	// PosServiceName is the fully-qualified name of the PosService service.
	PosServiceName = "cafelogos.pos.PosService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// PosServiceGetOrdersProcedure is the fully-qualified name of the PosService's GetOrders RPC.
	PosServiceGetOrdersProcedure = "/cafelogos.pos.PosService/GetOrders"
	// PosServiceGetUnpaidOrdersBySeatIdProcedure is the fully-qualified name of the PosService's
	// GetUnpaidOrdersBySeatId RPC.
	PosServiceGetUnpaidOrdersBySeatIdProcedure = "/cafelogos.pos.PosService/GetUnpaidOrdersBySeatId"
	// PosServicePostOrderProcedure is the fully-qualified name of the PosService's PostOrder RPC.
	PosServicePostOrderProcedure = "/cafelogos.pos.PosService/PostOrder"
	// PosServiceDeleteOrderProcedure is the fully-qualified name of the PosService's DeleteOrder RPC.
	PosServiceDeleteOrderProcedure = "/cafelogos.pos.PosService/DeleteOrder"
	// PosServicePostPaymentProcedure is the fully-qualified name of the PosService's PostPayment RPC.
	PosServicePostPaymentProcedure = "/cafelogos.pos.PosService/PostPayment"
	// PosServiceUpdatePaymentProcedure is the fully-qualified name of the PosService's UpdatePayment
	// RPC.
	PosServiceUpdatePaymentProcedure = "/cafelogos.pos.PosService/UpdatePayment"
	// PosServiceGetProductsProcedure is the fully-qualified name of the PosService's GetProducts RPC.
	PosServiceGetProductsProcedure = "/cafelogos.pos.PosService/GetProducts"
	// PosServicePostNewClientProcedure is the fully-qualified name of the PosService's PostNewClient
	// RPC.
	PosServicePostNewClientProcedure = "/cafelogos.pos.PosService/PostNewClient"
	// PosServiceUpdateClientProcedure is the fully-qualified name of the PosService's UpdateClient RPC.
	PosServiceUpdateClientProcedure = "/cafelogos.pos.PosService/UpdateClient"
	// PosServiceGetProductCategoriesProcedure is the fully-qualified name of the PosService's
	// GetProductCategories RPC.
	PosServiceGetProductCategoriesProcedure = "/cafelogos.pos.PosService/GetProductCategories"
	// PosServicePostProductCategoryProcedure is the fully-qualified name of the PosService's
	// PostProductCategory RPC.
	PosServicePostProductCategoryProcedure = "/cafelogos.pos.PosService/PostProductCategory"
	// PosServicePostProductProcedure is the fully-qualified name of the PosService's PostProduct RPC.
	PosServicePostProductProcedure = "/cafelogos.pos.PosService/PostProduct"
	// PosServiceUpdateProductProcedure is the fully-qualified name of the PosService's UpdateProduct
	// RPC.
	PosServiceUpdateProductProcedure = "/cafelogos.pos.PosService/UpdateProduct"
	// PosServiceDeleteProductProcedure is the fully-qualified name of the PosService's DeleteProduct
	// RPC.
	PosServiceDeleteProductProcedure = "/cafelogos.pos.PosService/DeleteProduct"
	// PosServicePostStockProcedure is the fully-qualified name of the PosService's PostStock RPC.
	PosServicePostStockProcedure = "/cafelogos.pos.PosService/PostStock"
	// PosServiceGetStocksProcedure is the fully-qualified name of the PosService's GetStocks RPC.
	PosServiceGetStocksProcedure = "/cafelogos.pos.PosService/GetStocks"
	// PosServicePostCoffeeBeanProcedure is the fully-qualified name of the PosService's PostCoffeeBean
	// RPC.
	PosServicePostCoffeeBeanProcedure = "/cafelogos.pos.PosService/PostCoffeeBean"
	// PosServiceGetCoffeeBeansProcedure is the fully-qualified name of the PosService's GetCoffeeBeans
	// RPC.
	PosServiceGetCoffeeBeansProcedure = "/cafelogos.pos.PosService/GetCoffeeBeans"
	// PosServiceDeleteAllOrdersProcedure is the fully-qualified name of the PosService's
	// DeleteAllOrders RPC.
	PosServiceDeleteAllOrdersProcedure = "/cafelogos.pos.PosService/DeleteAllOrders"
	// PosServicePostSeatProcedure is the fully-qualified name of the PosService's PostSeat RPC.
	PosServicePostSeatProcedure = "/cafelogos.pos.PosService/PostSeat"
	// PosServiceUpdateSeatProcedure is the fully-qualified name of the PosService's UpdateSeat RPC.
	PosServiceUpdateSeatProcedure = "/cafelogos.pos.PosService/UpdateSeat"
	// PosServiceGetSeatsProcedure is the fully-qualified name of the PosService's GetSeats RPC.
	PosServiceGetSeatsProcedure = "/cafelogos.pos.PosService/GetSeats"
	// PosServiceGetDiscountsProcedure is the fully-qualified name of the PosService's GetDiscounts RPC.
	PosServiceGetDiscountsProcedure = "/cafelogos.pos.PosService/GetDiscounts"
	// PosServicePostDiscountProcedure is the fully-qualified name of the PosService's PostDiscount RPC.
	PosServicePostDiscountProcedure = "/cafelogos.pos.PosService/PostDiscount"
)

// PosServiceClient is a client for the cafelogos.pos.PosService service.
type PosServiceClient interface {
	GetOrders(context.Context, *connect.Request[pos.GetOrdersRequest]) (*connect.Response[pos.GetOrdersResponse], error)
	GetUnpaidOrdersBySeatId(context.Context, *connect.Request[pos.GetUnpaidOrdersBySeatIdRequest]) (*connect.Response[pos.GetOrdersResponse], error)
	PostOrder(context.Context, *connect.Request[pos.PostOrderRequest]) (*connect.Response[pos.PostOrderResponse], error)
	DeleteOrder(context.Context, *connect.Request[pos.DeleteOrderRequest]) (*connect.Response[common.Empty], error)
	PostPayment(context.Context, *connect.Request[pos.PostPaymentRequest]) (*connect.Response[pos.PaymentResponse], error)
	UpdatePayment(context.Context, *connect.Request[pos.UpdatePaymentRequest]) (*connect.Response[pos.PaymentResponse], error)
	GetProducts(context.Context, *connect.Request[common.Empty]) (*connect.Response[pos.GetProductsResponse], error)
	PostNewClient(context.Context, *connect.Request[pos.PostNewClientRequest]) (*connect.Response[pos.PostNewClientResponse], error)
	UpdateClient(context.Context, *connect.Request[pos.UpdateClientRequest]) (*connect.Response[common.Empty], error)
	// Only Admin
	GetProductCategories(context.Context, *connect.Request[common.Empty]) (*connect.Response[pos.GetProductCategoriesResponse], error)
	PostProductCategory(context.Context, *connect.Request[pos.PostProductCategoryRequest]) (*connect.Response[common.Empty], error)
	PostProduct(context.Context, *connect.Request[pos.PostProductRequest]) (*connect.Response[common.Empty], error)
	UpdateProduct(context.Context, *connect.Request[pos.UpdateProductRequest]) (*connect.Response[common.Empty], error)
	DeleteProduct(context.Context, *connect.Request[pos.DeleteProductRequest]) (*connect.Response[common.Empty], error)
	PostStock(context.Context, *connect.Request[pos.PostStockRequest]) (*connect.Response[common.Empty], error)
	GetStocks(context.Context, *connect.Request[common.Empty]) (*connect.Response[pos.GetStocksResponse], error)
	PostCoffeeBean(context.Context, *connect.Request[pos.PostCoffeeBeanRequest]) (*connect.Response[common.Empty], error)
	GetCoffeeBeans(context.Context, *connect.Request[common.Empty]) (*connect.Response[pos.GetCoffeeBeansResponse], error)
	DeleteAllOrders(context.Context, *connect.Request[common.Empty]) (*connect.Response[common.Empty], error)
	PostSeat(context.Context, *connect.Request[pos.PostSeatRequest]) (*connect.Response[common.Empty], error)
	UpdateSeat(context.Context, *connect.Request[pos.UpdateSeatRequest]) (*connect.Response[common.Empty], error)
	GetSeats(context.Context, *connect.Request[common.Empty]) (*connect.Response[pos.GetSeatsResponse], error)
	GetDiscounts(context.Context, *connect.Request[common.Empty]) (*connect.Response[pos.GetDiscountsResponse], error)
	PostDiscount(context.Context, *connect.Request[pos.PostDiscountRequest]) (*connect.Response[common.Empty], error)
}

// NewPosServiceClient constructs a client for the cafelogos.pos.PosService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPosServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) PosServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &posServiceClient{
		getOrders: connect.NewClient[pos.GetOrdersRequest, pos.GetOrdersResponse](
			httpClient,
			baseURL+PosServiceGetOrdersProcedure,
			opts...,
		),
		getUnpaidOrdersBySeatId: connect.NewClient[pos.GetUnpaidOrdersBySeatIdRequest, pos.GetOrdersResponse](
			httpClient,
			baseURL+PosServiceGetUnpaidOrdersBySeatIdProcedure,
			opts...,
		),
		postOrder: connect.NewClient[pos.PostOrderRequest, pos.PostOrderResponse](
			httpClient,
			baseURL+PosServicePostOrderProcedure,
			opts...,
		),
		deleteOrder: connect.NewClient[pos.DeleteOrderRequest, common.Empty](
			httpClient,
			baseURL+PosServiceDeleteOrderProcedure,
			opts...,
		),
		postPayment: connect.NewClient[pos.PostPaymentRequest, pos.PaymentResponse](
			httpClient,
			baseURL+PosServicePostPaymentProcedure,
			opts...,
		),
		updatePayment: connect.NewClient[pos.UpdatePaymentRequest, pos.PaymentResponse](
			httpClient,
			baseURL+PosServiceUpdatePaymentProcedure,
			opts...,
		),
		getProducts: connect.NewClient[common.Empty, pos.GetProductsResponse](
			httpClient,
			baseURL+PosServiceGetProductsProcedure,
			opts...,
		),
		postNewClient: connect.NewClient[pos.PostNewClientRequest, pos.PostNewClientResponse](
			httpClient,
			baseURL+PosServicePostNewClientProcedure,
			opts...,
		),
		updateClient: connect.NewClient[pos.UpdateClientRequest, common.Empty](
			httpClient,
			baseURL+PosServiceUpdateClientProcedure,
			opts...,
		),
		getProductCategories: connect.NewClient[common.Empty, pos.GetProductCategoriesResponse](
			httpClient,
			baseURL+PosServiceGetProductCategoriesProcedure,
			opts...,
		),
		postProductCategory: connect.NewClient[pos.PostProductCategoryRequest, common.Empty](
			httpClient,
			baseURL+PosServicePostProductCategoryProcedure,
			opts...,
		),
		postProduct: connect.NewClient[pos.PostProductRequest, common.Empty](
			httpClient,
			baseURL+PosServicePostProductProcedure,
			opts...,
		),
		updateProduct: connect.NewClient[pos.UpdateProductRequest, common.Empty](
			httpClient,
			baseURL+PosServiceUpdateProductProcedure,
			opts...,
		),
		deleteProduct: connect.NewClient[pos.DeleteProductRequest, common.Empty](
			httpClient,
			baseURL+PosServiceDeleteProductProcedure,
			opts...,
		),
		postStock: connect.NewClient[pos.PostStockRequest, common.Empty](
			httpClient,
			baseURL+PosServicePostStockProcedure,
			opts...,
		),
		getStocks: connect.NewClient[common.Empty, pos.GetStocksResponse](
			httpClient,
			baseURL+PosServiceGetStocksProcedure,
			opts...,
		),
		postCoffeeBean: connect.NewClient[pos.PostCoffeeBeanRequest, common.Empty](
			httpClient,
			baseURL+PosServicePostCoffeeBeanProcedure,
			opts...,
		),
		getCoffeeBeans: connect.NewClient[common.Empty, pos.GetCoffeeBeansResponse](
			httpClient,
			baseURL+PosServiceGetCoffeeBeansProcedure,
			opts...,
		),
		deleteAllOrders: connect.NewClient[common.Empty, common.Empty](
			httpClient,
			baseURL+PosServiceDeleteAllOrdersProcedure,
			opts...,
		),
		postSeat: connect.NewClient[pos.PostSeatRequest, common.Empty](
			httpClient,
			baseURL+PosServicePostSeatProcedure,
			opts...,
		),
		updateSeat: connect.NewClient[pos.UpdateSeatRequest, common.Empty](
			httpClient,
			baseURL+PosServiceUpdateSeatProcedure,
			opts...,
		),
		getSeats: connect.NewClient[common.Empty, pos.GetSeatsResponse](
			httpClient,
			baseURL+PosServiceGetSeatsProcedure,
			opts...,
		),
		getDiscounts: connect.NewClient[common.Empty, pos.GetDiscountsResponse](
			httpClient,
			baseURL+PosServiceGetDiscountsProcedure,
			opts...,
		),
		postDiscount: connect.NewClient[pos.PostDiscountRequest, common.Empty](
			httpClient,
			baseURL+PosServicePostDiscountProcedure,
			opts...,
		),
	}
}

// posServiceClient implements PosServiceClient.
type posServiceClient struct {
	getOrders               *connect.Client[pos.GetOrdersRequest, pos.GetOrdersResponse]
	getUnpaidOrdersBySeatId *connect.Client[pos.GetUnpaidOrdersBySeatIdRequest, pos.GetOrdersResponse]
	postOrder               *connect.Client[pos.PostOrderRequest, pos.PostOrderResponse]
	deleteOrder             *connect.Client[pos.DeleteOrderRequest, common.Empty]
	postPayment             *connect.Client[pos.PostPaymentRequest, pos.PaymentResponse]
	updatePayment           *connect.Client[pos.UpdatePaymentRequest, pos.PaymentResponse]
	getProducts             *connect.Client[common.Empty, pos.GetProductsResponse]
	postNewClient           *connect.Client[pos.PostNewClientRequest, pos.PostNewClientResponse]
	updateClient            *connect.Client[pos.UpdateClientRequest, common.Empty]
	getProductCategories    *connect.Client[common.Empty, pos.GetProductCategoriesResponse]
	postProductCategory     *connect.Client[pos.PostProductCategoryRequest, common.Empty]
	postProduct             *connect.Client[pos.PostProductRequest, common.Empty]
	updateProduct           *connect.Client[pos.UpdateProductRequest, common.Empty]
	deleteProduct           *connect.Client[pos.DeleteProductRequest, common.Empty]
	postStock               *connect.Client[pos.PostStockRequest, common.Empty]
	getStocks               *connect.Client[common.Empty, pos.GetStocksResponse]
	postCoffeeBean          *connect.Client[pos.PostCoffeeBeanRequest, common.Empty]
	getCoffeeBeans          *connect.Client[common.Empty, pos.GetCoffeeBeansResponse]
	deleteAllOrders         *connect.Client[common.Empty, common.Empty]
	postSeat                *connect.Client[pos.PostSeatRequest, common.Empty]
	updateSeat              *connect.Client[pos.UpdateSeatRequest, common.Empty]
	getSeats                *connect.Client[common.Empty, pos.GetSeatsResponse]
	getDiscounts            *connect.Client[common.Empty, pos.GetDiscountsResponse]
	postDiscount            *connect.Client[pos.PostDiscountRequest, common.Empty]
}

// GetOrders calls cafelogos.pos.PosService.GetOrders.
func (c *posServiceClient) GetOrders(ctx context.Context, req *connect.Request[pos.GetOrdersRequest]) (*connect.Response[pos.GetOrdersResponse], error) {
	return c.getOrders.CallUnary(ctx, req)
}

// GetUnpaidOrdersBySeatId calls cafelogos.pos.PosService.GetUnpaidOrdersBySeatId.
func (c *posServiceClient) GetUnpaidOrdersBySeatId(ctx context.Context, req *connect.Request[pos.GetUnpaidOrdersBySeatIdRequest]) (*connect.Response[pos.GetOrdersResponse], error) {
	return c.getUnpaidOrdersBySeatId.CallUnary(ctx, req)
}

// PostOrder calls cafelogos.pos.PosService.PostOrder.
func (c *posServiceClient) PostOrder(ctx context.Context, req *connect.Request[pos.PostOrderRequest]) (*connect.Response[pos.PostOrderResponse], error) {
	return c.postOrder.CallUnary(ctx, req)
}

// DeleteOrder calls cafelogos.pos.PosService.DeleteOrder.
func (c *posServiceClient) DeleteOrder(ctx context.Context, req *connect.Request[pos.DeleteOrderRequest]) (*connect.Response[common.Empty], error) {
	return c.deleteOrder.CallUnary(ctx, req)
}

// PostPayment calls cafelogos.pos.PosService.PostPayment.
func (c *posServiceClient) PostPayment(ctx context.Context, req *connect.Request[pos.PostPaymentRequest]) (*connect.Response[pos.PaymentResponse], error) {
	return c.postPayment.CallUnary(ctx, req)
}

// UpdatePayment calls cafelogos.pos.PosService.UpdatePayment.
func (c *posServiceClient) UpdatePayment(ctx context.Context, req *connect.Request[pos.UpdatePaymentRequest]) (*connect.Response[pos.PaymentResponse], error) {
	return c.updatePayment.CallUnary(ctx, req)
}

// GetProducts calls cafelogos.pos.PosService.GetProducts.
func (c *posServiceClient) GetProducts(ctx context.Context, req *connect.Request[common.Empty]) (*connect.Response[pos.GetProductsResponse], error) {
	return c.getProducts.CallUnary(ctx, req)
}

// PostNewClient calls cafelogos.pos.PosService.PostNewClient.
func (c *posServiceClient) PostNewClient(ctx context.Context, req *connect.Request[pos.PostNewClientRequest]) (*connect.Response[pos.PostNewClientResponse], error) {
	return c.postNewClient.CallUnary(ctx, req)
}

// UpdateClient calls cafelogos.pos.PosService.UpdateClient.
func (c *posServiceClient) UpdateClient(ctx context.Context, req *connect.Request[pos.UpdateClientRequest]) (*connect.Response[common.Empty], error) {
	return c.updateClient.CallUnary(ctx, req)
}

// GetProductCategories calls cafelogos.pos.PosService.GetProductCategories.
func (c *posServiceClient) GetProductCategories(ctx context.Context, req *connect.Request[common.Empty]) (*connect.Response[pos.GetProductCategoriesResponse], error) {
	return c.getProductCategories.CallUnary(ctx, req)
}

// PostProductCategory calls cafelogos.pos.PosService.PostProductCategory.
func (c *posServiceClient) PostProductCategory(ctx context.Context, req *connect.Request[pos.PostProductCategoryRequest]) (*connect.Response[common.Empty], error) {
	return c.postProductCategory.CallUnary(ctx, req)
}

// PostProduct calls cafelogos.pos.PosService.PostProduct.
func (c *posServiceClient) PostProduct(ctx context.Context, req *connect.Request[pos.PostProductRequest]) (*connect.Response[common.Empty], error) {
	return c.postProduct.CallUnary(ctx, req)
}

// UpdateProduct calls cafelogos.pos.PosService.UpdateProduct.
func (c *posServiceClient) UpdateProduct(ctx context.Context, req *connect.Request[pos.UpdateProductRequest]) (*connect.Response[common.Empty], error) {
	return c.updateProduct.CallUnary(ctx, req)
}

// DeleteProduct calls cafelogos.pos.PosService.DeleteProduct.
func (c *posServiceClient) DeleteProduct(ctx context.Context, req *connect.Request[pos.DeleteProductRequest]) (*connect.Response[common.Empty], error) {
	return c.deleteProduct.CallUnary(ctx, req)
}

// PostStock calls cafelogos.pos.PosService.PostStock.
func (c *posServiceClient) PostStock(ctx context.Context, req *connect.Request[pos.PostStockRequest]) (*connect.Response[common.Empty], error) {
	return c.postStock.CallUnary(ctx, req)
}

// GetStocks calls cafelogos.pos.PosService.GetStocks.
func (c *posServiceClient) GetStocks(ctx context.Context, req *connect.Request[common.Empty]) (*connect.Response[pos.GetStocksResponse], error) {
	return c.getStocks.CallUnary(ctx, req)
}

// PostCoffeeBean calls cafelogos.pos.PosService.PostCoffeeBean.
func (c *posServiceClient) PostCoffeeBean(ctx context.Context, req *connect.Request[pos.PostCoffeeBeanRequest]) (*connect.Response[common.Empty], error) {
	return c.postCoffeeBean.CallUnary(ctx, req)
}

// GetCoffeeBeans calls cafelogos.pos.PosService.GetCoffeeBeans.
func (c *posServiceClient) GetCoffeeBeans(ctx context.Context, req *connect.Request[common.Empty]) (*connect.Response[pos.GetCoffeeBeansResponse], error) {
	return c.getCoffeeBeans.CallUnary(ctx, req)
}

// DeleteAllOrders calls cafelogos.pos.PosService.DeleteAllOrders.
func (c *posServiceClient) DeleteAllOrders(ctx context.Context, req *connect.Request[common.Empty]) (*connect.Response[common.Empty], error) {
	return c.deleteAllOrders.CallUnary(ctx, req)
}

// PostSeat calls cafelogos.pos.PosService.PostSeat.
func (c *posServiceClient) PostSeat(ctx context.Context, req *connect.Request[pos.PostSeatRequest]) (*connect.Response[common.Empty], error) {
	return c.postSeat.CallUnary(ctx, req)
}

// UpdateSeat calls cafelogos.pos.PosService.UpdateSeat.
func (c *posServiceClient) UpdateSeat(ctx context.Context, req *connect.Request[pos.UpdateSeatRequest]) (*connect.Response[common.Empty], error) {
	return c.updateSeat.CallUnary(ctx, req)
}

// GetSeats calls cafelogos.pos.PosService.GetSeats.
func (c *posServiceClient) GetSeats(ctx context.Context, req *connect.Request[common.Empty]) (*connect.Response[pos.GetSeatsResponse], error) {
	return c.getSeats.CallUnary(ctx, req)
}

// GetDiscounts calls cafelogos.pos.PosService.GetDiscounts.
func (c *posServiceClient) GetDiscounts(ctx context.Context, req *connect.Request[common.Empty]) (*connect.Response[pos.GetDiscountsResponse], error) {
	return c.getDiscounts.CallUnary(ctx, req)
}

// PostDiscount calls cafelogos.pos.PosService.PostDiscount.
func (c *posServiceClient) PostDiscount(ctx context.Context, req *connect.Request[pos.PostDiscountRequest]) (*connect.Response[common.Empty], error) {
	return c.postDiscount.CallUnary(ctx, req)
}

// PosServiceHandler is an implementation of the cafelogos.pos.PosService service.
type PosServiceHandler interface {
	GetOrders(context.Context, *connect.Request[pos.GetOrdersRequest]) (*connect.Response[pos.GetOrdersResponse], error)
	GetUnpaidOrdersBySeatId(context.Context, *connect.Request[pos.GetUnpaidOrdersBySeatIdRequest]) (*connect.Response[pos.GetOrdersResponse], error)
	PostOrder(context.Context, *connect.Request[pos.PostOrderRequest]) (*connect.Response[pos.PostOrderResponse], error)
	DeleteOrder(context.Context, *connect.Request[pos.DeleteOrderRequest]) (*connect.Response[common.Empty], error)
	PostPayment(context.Context, *connect.Request[pos.PostPaymentRequest]) (*connect.Response[pos.PaymentResponse], error)
	UpdatePayment(context.Context, *connect.Request[pos.UpdatePaymentRequest]) (*connect.Response[pos.PaymentResponse], error)
	GetProducts(context.Context, *connect.Request[common.Empty]) (*connect.Response[pos.GetProductsResponse], error)
	PostNewClient(context.Context, *connect.Request[pos.PostNewClientRequest]) (*connect.Response[pos.PostNewClientResponse], error)
	UpdateClient(context.Context, *connect.Request[pos.UpdateClientRequest]) (*connect.Response[common.Empty], error)
	// Only Admin
	GetProductCategories(context.Context, *connect.Request[common.Empty]) (*connect.Response[pos.GetProductCategoriesResponse], error)
	PostProductCategory(context.Context, *connect.Request[pos.PostProductCategoryRequest]) (*connect.Response[common.Empty], error)
	PostProduct(context.Context, *connect.Request[pos.PostProductRequest]) (*connect.Response[common.Empty], error)
	UpdateProduct(context.Context, *connect.Request[pos.UpdateProductRequest]) (*connect.Response[common.Empty], error)
	DeleteProduct(context.Context, *connect.Request[pos.DeleteProductRequest]) (*connect.Response[common.Empty], error)
	PostStock(context.Context, *connect.Request[pos.PostStockRequest]) (*connect.Response[common.Empty], error)
	GetStocks(context.Context, *connect.Request[common.Empty]) (*connect.Response[pos.GetStocksResponse], error)
	PostCoffeeBean(context.Context, *connect.Request[pos.PostCoffeeBeanRequest]) (*connect.Response[common.Empty], error)
	GetCoffeeBeans(context.Context, *connect.Request[common.Empty]) (*connect.Response[pos.GetCoffeeBeansResponse], error)
	DeleteAllOrders(context.Context, *connect.Request[common.Empty]) (*connect.Response[common.Empty], error)
	PostSeat(context.Context, *connect.Request[pos.PostSeatRequest]) (*connect.Response[common.Empty], error)
	UpdateSeat(context.Context, *connect.Request[pos.UpdateSeatRequest]) (*connect.Response[common.Empty], error)
	GetSeats(context.Context, *connect.Request[common.Empty]) (*connect.Response[pos.GetSeatsResponse], error)
	GetDiscounts(context.Context, *connect.Request[common.Empty]) (*connect.Response[pos.GetDiscountsResponse], error)
	PostDiscount(context.Context, *connect.Request[pos.PostDiscountRequest]) (*connect.Response[common.Empty], error)
}

// NewPosServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPosServiceHandler(svc PosServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	posServiceGetOrdersHandler := connect.NewUnaryHandler(
		PosServiceGetOrdersProcedure,
		svc.GetOrders,
		opts...,
	)
	posServiceGetUnpaidOrdersBySeatIdHandler := connect.NewUnaryHandler(
		PosServiceGetUnpaidOrdersBySeatIdProcedure,
		svc.GetUnpaidOrdersBySeatId,
		opts...,
	)
	posServicePostOrderHandler := connect.NewUnaryHandler(
		PosServicePostOrderProcedure,
		svc.PostOrder,
		opts...,
	)
	posServiceDeleteOrderHandler := connect.NewUnaryHandler(
		PosServiceDeleteOrderProcedure,
		svc.DeleteOrder,
		opts...,
	)
	posServicePostPaymentHandler := connect.NewUnaryHandler(
		PosServicePostPaymentProcedure,
		svc.PostPayment,
		opts...,
	)
	posServiceUpdatePaymentHandler := connect.NewUnaryHandler(
		PosServiceUpdatePaymentProcedure,
		svc.UpdatePayment,
		opts...,
	)
	posServiceGetProductsHandler := connect.NewUnaryHandler(
		PosServiceGetProductsProcedure,
		svc.GetProducts,
		opts...,
	)
	posServicePostNewClientHandler := connect.NewUnaryHandler(
		PosServicePostNewClientProcedure,
		svc.PostNewClient,
		opts...,
	)
	posServiceUpdateClientHandler := connect.NewUnaryHandler(
		PosServiceUpdateClientProcedure,
		svc.UpdateClient,
		opts...,
	)
	posServiceGetProductCategoriesHandler := connect.NewUnaryHandler(
		PosServiceGetProductCategoriesProcedure,
		svc.GetProductCategories,
		opts...,
	)
	posServicePostProductCategoryHandler := connect.NewUnaryHandler(
		PosServicePostProductCategoryProcedure,
		svc.PostProductCategory,
		opts...,
	)
	posServicePostProductHandler := connect.NewUnaryHandler(
		PosServicePostProductProcedure,
		svc.PostProduct,
		opts...,
	)
	posServiceUpdateProductHandler := connect.NewUnaryHandler(
		PosServiceUpdateProductProcedure,
		svc.UpdateProduct,
		opts...,
	)
	posServiceDeleteProductHandler := connect.NewUnaryHandler(
		PosServiceDeleteProductProcedure,
		svc.DeleteProduct,
		opts...,
	)
	posServicePostStockHandler := connect.NewUnaryHandler(
		PosServicePostStockProcedure,
		svc.PostStock,
		opts...,
	)
	posServiceGetStocksHandler := connect.NewUnaryHandler(
		PosServiceGetStocksProcedure,
		svc.GetStocks,
		opts...,
	)
	posServicePostCoffeeBeanHandler := connect.NewUnaryHandler(
		PosServicePostCoffeeBeanProcedure,
		svc.PostCoffeeBean,
		opts...,
	)
	posServiceGetCoffeeBeansHandler := connect.NewUnaryHandler(
		PosServiceGetCoffeeBeansProcedure,
		svc.GetCoffeeBeans,
		opts...,
	)
	posServiceDeleteAllOrdersHandler := connect.NewUnaryHandler(
		PosServiceDeleteAllOrdersProcedure,
		svc.DeleteAllOrders,
		opts...,
	)
	posServicePostSeatHandler := connect.NewUnaryHandler(
		PosServicePostSeatProcedure,
		svc.PostSeat,
		opts...,
	)
	posServiceUpdateSeatHandler := connect.NewUnaryHandler(
		PosServiceUpdateSeatProcedure,
		svc.UpdateSeat,
		opts...,
	)
	posServiceGetSeatsHandler := connect.NewUnaryHandler(
		PosServiceGetSeatsProcedure,
		svc.GetSeats,
		opts...,
	)
	posServiceGetDiscountsHandler := connect.NewUnaryHandler(
		PosServiceGetDiscountsProcedure,
		svc.GetDiscounts,
		opts...,
	)
	posServicePostDiscountHandler := connect.NewUnaryHandler(
		PosServicePostDiscountProcedure,
		svc.PostDiscount,
		opts...,
	)
	return "/cafelogos.pos.PosService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case PosServiceGetOrdersProcedure:
			posServiceGetOrdersHandler.ServeHTTP(w, r)
		case PosServiceGetUnpaidOrdersBySeatIdProcedure:
			posServiceGetUnpaidOrdersBySeatIdHandler.ServeHTTP(w, r)
		case PosServicePostOrderProcedure:
			posServicePostOrderHandler.ServeHTTP(w, r)
		case PosServiceDeleteOrderProcedure:
			posServiceDeleteOrderHandler.ServeHTTP(w, r)
		case PosServicePostPaymentProcedure:
			posServicePostPaymentHandler.ServeHTTP(w, r)
		case PosServiceUpdatePaymentProcedure:
			posServiceUpdatePaymentHandler.ServeHTTP(w, r)
		case PosServiceGetProductsProcedure:
			posServiceGetProductsHandler.ServeHTTP(w, r)
		case PosServicePostNewClientProcedure:
			posServicePostNewClientHandler.ServeHTTP(w, r)
		case PosServiceUpdateClientProcedure:
			posServiceUpdateClientHandler.ServeHTTP(w, r)
		case PosServiceGetProductCategoriesProcedure:
			posServiceGetProductCategoriesHandler.ServeHTTP(w, r)
		case PosServicePostProductCategoryProcedure:
			posServicePostProductCategoryHandler.ServeHTTP(w, r)
		case PosServicePostProductProcedure:
			posServicePostProductHandler.ServeHTTP(w, r)
		case PosServiceUpdateProductProcedure:
			posServiceUpdateProductHandler.ServeHTTP(w, r)
		case PosServiceDeleteProductProcedure:
			posServiceDeleteProductHandler.ServeHTTP(w, r)
		case PosServicePostStockProcedure:
			posServicePostStockHandler.ServeHTTP(w, r)
		case PosServiceGetStocksProcedure:
			posServiceGetStocksHandler.ServeHTTP(w, r)
		case PosServicePostCoffeeBeanProcedure:
			posServicePostCoffeeBeanHandler.ServeHTTP(w, r)
		case PosServiceGetCoffeeBeansProcedure:
			posServiceGetCoffeeBeansHandler.ServeHTTP(w, r)
		case PosServiceDeleteAllOrdersProcedure:
			posServiceDeleteAllOrdersHandler.ServeHTTP(w, r)
		case PosServicePostSeatProcedure:
			posServicePostSeatHandler.ServeHTTP(w, r)
		case PosServiceUpdateSeatProcedure:
			posServiceUpdateSeatHandler.ServeHTTP(w, r)
		case PosServiceGetSeatsProcedure:
			posServiceGetSeatsHandler.ServeHTTP(w, r)
		case PosServiceGetDiscountsProcedure:
			posServiceGetDiscountsHandler.ServeHTTP(w, r)
		case PosServicePostDiscountProcedure:
			posServicePostDiscountHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedPosServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPosServiceHandler struct{}

func (UnimplementedPosServiceHandler) GetOrders(context.Context, *connect.Request[pos.GetOrdersRequest]) (*connect.Response[pos.GetOrdersResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cafelogos.pos.PosService.GetOrders is not implemented"))
}

func (UnimplementedPosServiceHandler) GetUnpaidOrdersBySeatId(context.Context, *connect.Request[pos.GetUnpaidOrdersBySeatIdRequest]) (*connect.Response[pos.GetOrdersResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cafelogos.pos.PosService.GetUnpaidOrdersBySeatId is not implemented"))
}

func (UnimplementedPosServiceHandler) PostOrder(context.Context, *connect.Request[pos.PostOrderRequest]) (*connect.Response[pos.PostOrderResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cafelogos.pos.PosService.PostOrder is not implemented"))
}

func (UnimplementedPosServiceHandler) DeleteOrder(context.Context, *connect.Request[pos.DeleteOrderRequest]) (*connect.Response[common.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cafelogos.pos.PosService.DeleteOrder is not implemented"))
}

func (UnimplementedPosServiceHandler) PostPayment(context.Context, *connect.Request[pos.PostPaymentRequest]) (*connect.Response[pos.PaymentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cafelogos.pos.PosService.PostPayment is not implemented"))
}

func (UnimplementedPosServiceHandler) UpdatePayment(context.Context, *connect.Request[pos.UpdatePaymentRequest]) (*connect.Response[pos.PaymentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cafelogos.pos.PosService.UpdatePayment is not implemented"))
}

func (UnimplementedPosServiceHandler) GetProducts(context.Context, *connect.Request[common.Empty]) (*connect.Response[pos.GetProductsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cafelogos.pos.PosService.GetProducts is not implemented"))
}

func (UnimplementedPosServiceHandler) PostNewClient(context.Context, *connect.Request[pos.PostNewClientRequest]) (*connect.Response[pos.PostNewClientResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cafelogos.pos.PosService.PostNewClient is not implemented"))
}

func (UnimplementedPosServiceHandler) UpdateClient(context.Context, *connect.Request[pos.UpdateClientRequest]) (*connect.Response[common.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cafelogos.pos.PosService.UpdateClient is not implemented"))
}

func (UnimplementedPosServiceHandler) GetProductCategories(context.Context, *connect.Request[common.Empty]) (*connect.Response[pos.GetProductCategoriesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cafelogos.pos.PosService.GetProductCategories is not implemented"))
}

func (UnimplementedPosServiceHandler) PostProductCategory(context.Context, *connect.Request[pos.PostProductCategoryRequest]) (*connect.Response[common.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cafelogos.pos.PosService.PostProductCategory is not implemented"))
}

func (UnimplementedPosServiceHandler) PostProduct(context.Context, *connect.Request[pos.PostProductRequest]) (*connect.Response[common.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cafelogos.pos.PosService.PostProduct is not implemented"))
}

func (UnimplementedPosServiceHandler) UpdateProduct(context.Context, *connect.Request[pos.UpdateProductRequest]) (*connect.Response[common.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cafelogos.pos.PosService.UpdateProduct is not implemented"))
}

func (UnimplementedPosServiceHandler) DeleteProduct(context.Context, *connect.Request[pos.DeleteProductRequest]) (*connect.Response[common.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cafelogos.pos.PosService.DeleteProduct is not implemented"))
}

func (UnimplementedPosServiceHandler) PostStock(context.Context, *connect.Request[pos.PostStockRequest]) (*connect.Response[common.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cafelogos.pos.PosService.PostStock is not implemented"))
}

func (UnimplementedPosServiceHandler) GetStocks(context.Context, *connect.Request[common.Empty]) (*connect.Response[pos.GetStocksResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cafelogos.pos.PosService.GetStocks is not implemented"))
}

func (UnimplementedPosServiceHandler) PostCoffeeBean(context.Context, *connect.Request[pos.PostCoffeeBeanRequest]) (*connect.Response[common.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cafelogos.pos.PosService.PostCoffeeBean is not implemented"))
}

func (UnimplementedPosServiceHandler) GetCoffeeBeans(context.Context, *connect.Request[common.Empty]) (*connect.Response[pos.GetCoffeeBeansResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cafelogos.pos.PosService.GetCoffeeBeans is not implemented"))
}

func (UnimplementedPosServiceHandler) DeleteAllOrders(context.Context, *connect.Request[common.Empty]) (*connect.Response[common.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cafelogos.pos.PosService.DeleteAllOrders is not implemented"))
}

func (UnimplementedPosServiceHandler) PostSeat(context.Context, *connect.Request[pos.PostSeatRequest]) (*connect.Response[common.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cafelogos.pos.PosService.PostSeat is not implemented"))
}

func (UnimplementedPosServiceHandler) UpdateSeat(context.Context, *connect.Request[pos.UpdateSeatRequest]) (*connect.Response[common.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cafelogos.pos.PosService.UpdateSeat is not implemented"))
}

func (UnimplementedPosServiceHandler) GetSeats(context.Context, *connect.Request[common.Empty]) (*connect.Response[pos.GetSeatsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cafelogos.pos.PosService.GetSeats is not implemented"))
}

func (UnimplementedPosServiceHandler) GetDiscounts(context.Context, *connect.Request[common.Empty]) (*connect.Response[pos.GetDiscountsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cafelogos.pos.PosService.GetDiscounts is not implemented"))
}

func (UnimplementedPosServiceHandler) PostDiscount(context.Context, *connect.Request[pos.PostDiscountRequest]) (*connect.Response[common.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cafelogos.pos.PosService.PostDiscount is not implemented"))
}
