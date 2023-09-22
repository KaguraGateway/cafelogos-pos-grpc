// @generated by protoc-gen-es v1.3.1
// @generated from file orderlink/orderlink_service.proto (package cafelogos.orderlink, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message cafelogos.orderlink.PostOrderItemInput
 */
export declare class PostOrderItemInput extends Message<PostOrderItemInput> {
  /**
   * @generated from field: string product_id = 1;
   */
  productId: string;

  /**
   * @generated from field: string coffee_brew_id = 2;
   */
  coffeeBrewId: string;

  /**
   * @generated from field: uint32 quantity = 3;
   */
  quantity: number;

  constructor(data?: PartialMessage<PostOrderItemInput>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "cafelogos.orderlink.PostOrderItemInput";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PostOrderItemInput;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PostOrderItemInput;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PostOrderItemInput;

  static equals(a: PostOrderItemInput | PlainMessage<PostOrderItemInput> | undefined, b: PostOrderItemInput | PlainMessage<PostOrderItemInput> | undefined): boolean;
}

/**
 * @generated from message cafelogos.orderlink.PostOrderInput
 */
export declare class PostOrderInput extends Message<PostOrderInput> {
  /**
   * @generated from field: string order_id = 1;
   */
  orderId: string;

  /**
   * @generated from field: string order_at = 2;
   */
  orderAt: string;

  /**
   * @generated from field: repeated cafelogos.orderlink.PostOrderItemInput items = 3;
   */
  items: PostOrderItemInput[];

  /**
   * @generated from field: cafelogos.orderlink.PostOrderInput.OrderType type = 4;
   */
  type: PostOrderInput_OrderType;

  /**
   * @generated from field: string ticket_id = 5;
   */
  ticketId: string;

  /**
   * @generated from field: string ticket_addr = 6;
   */
  ticketAddr: string;

  /**
   * @generated from field: string seat_name = 7;
   */
  seatName: string;

  constructor(data?: PartialMessage<PostOrderInput>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "cafelogos.orderlink.PostOrderInput";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PostOrderInput;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PostOrderInput;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PostOrderInput;

  static equals(a: PostOrderInput | PlainMessage<PostOrderInput> | undefined, b: PostOrderInput | PlainMessage<PostOrderInput> | undefined): boolean;
}

/**
 * @generated from enum cafelogos.orderlink.PostOrderInput.OrderType
 */
export declare enum PostOrderInput_OrderType {
  /**
   * @generated from enum value: EAT_IN = 0;
   */
  EAT_IN = 0,

  /**
   * @generated from enum value: TAKE_OUT = 1;
   */
  TAKE_OUT = 1,
}

