// @generated by protoc-gen-connect-es v1.0.0-rc1 with parameter "target=ts"
// @generated from file orderlink/orderlink_service.proto (package cafelogos.orderlink, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { PostOrderInput } from "./orderlink_service_pb.js";
import { Empty } from "../common/common_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service cafelogos.orderlink.OrderLinkService
 */
export const OrderLinkService = {
  typeName: "cafelogos.orderlink.OrderLinkService",
  methods: {
    /**
     * @generated from rpc cafelogos.orderlink.OrderLinkService.PostOrder
     */
    postOrder: {
      name: "PostOrder",
      I: PostOrderInput,
      O: Empty,
      kind: MethodKind.Unary,
    },
  }
} as const;

