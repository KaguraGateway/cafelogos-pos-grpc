// @generated by protoc-gen-connect-es v1.0.0-rc1
// @generated from file pos/pos_service.proto (package cafelogos.pos, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { DeleteProductRequest, GetCoffeeBeansResponse, GetOrdersRequest, GetOrdersResponse, GetProductCategoriesResponse, GetProductsResponse, GetStocksResponse, PostCoffeeBeanRequest, PostOrderRequest, PostProductCategoryRequest, PostProductRequest, PostStockRequest, UpdateProductRequest } from "./pos_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";
import { Empty } from "../common/common_pb.js";

/**
 * @generated from service cafelogos.pos.PosService
 */
export const PosService = {
  typeName: "cafelogos.pos.PosService",
  methods: {
    /**
     * @generated from rpc cafelogos.pos.PosService.GetOrders
     */
    getOrders: {
      name: "GetOrders",
      I: GetOrdersRequest,
      O: GetOrdersResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc cafelogos.pos.PosService.PostOrder
     */
    postOrder: {
      name: "PostOrder",
      I: PostOrderRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc cafelogos.pos.PosService.GetProducts
     */
    getProducts: {
      name: "GetProducts",
      I: Empty,
      O: GetProductsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Only Admin
     *
     * @generated from rpc cafelogos.pos.PosService.GetProductCategories
     */
    getProductCategories: {
      name: "GetProductCategories",
      I: Empty,
      O: GetProductCategoriesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc cafelogos.pos.PosService.PostProductCategory
     */
    postProductCategory: {
      name: "PostProductCategory",
      I: PostProductCategoryRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc cafelogos.pos.PosService.PostProduct
     */
    postProduct: {
      name: "PostProduct",
      I: PostProductRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc cafelogos.pos.PosService.UpdateProduct
     */
    updateProduct: {
      name: "UpdateProduct",
      I: UpdateProductRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc cafelogos.pos.PosService.DeleteProduct
     */
    deleteProduct: {
      name: "DeleteProduct",
      I: DeleteProductRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc cafelogos.pos.PosService.PostStock
     */
    postStock: {
      name: "PostStock",
      I: PostStockRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc cafelogos.pos.PosService.GetStocks
     */
    getStocks: {
      name: "GetStocks",
      I: Empty,
      O: GetStocksResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc cafelogos.pos.PosService.PostCoffeeBean
     */
    postCoffeeBean: {
      name: "PostCoffeeBean",
      I: PostCoffeeBeanRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc cafelogos.pos.PosService.GetCoffeeBeans
     */
    getCoffeeBeans: {
      name: "GetCoffeeBeans",
      I: Empty,
      O: GetCoffeeBeansResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc cafelogos.pos.PosService.DeleteAllOrders
     */
    deleteAllOrders: {
      name: "DeleteAllOrders",
      I: Empty,
      O: Empty,
      kind: MethodKind.Unary,
    },
  }
};

