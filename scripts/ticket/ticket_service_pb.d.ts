// @generated by protoc-gen-es v1.10.0 with parameter "target=ts+js+dts"
// @generated from file ticket/ticket_service.proto (package cafelogos.ticket, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message cafelogos.ticket.RequestIssueTicket
 */
export declare class RequestIssueTicket extends Message<RequestIssueTicket> {
  /**
   * @generated from field: string prefix = 1;
   */
  prefix: string;

  constructor(data?: PartialMessage<RequestIssueTicket>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "cafelogos.ticket.RequestIssueTicket";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RequestIssueTicket;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RequestIssueTicket;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RequestIssueTicket;

  static equals(a: RequestIssueTicket | PlainMessage<RequestIssueTicket> | undefined, b: RequestIssueTicket | PlainMessage<RequestIssueTicket> | undefined): boolean;
}

/**
 * @generated from message cafelogos.ticket.ResponseIssueTicket
 */
export declare class ResponseIssueTicket extends Message<ResponseIssueTicket> {
  /**
   * @generated from field: cafelogos.ticket.Ticket ticket = 1;
   */
  ticket?: Ticket;

  constructor(data?: PartialMessage<ResponseIssueTicket>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "cafelogos.ticket.ResponseIssueTicket";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ResponseIssueTicket;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ResponseIssueTicket;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ResponseIssueTicket;

  static equals(a: ResponseIssueTicket | PlainMessage<ResponseIssueTicket> | undefined, b: ResponseIssueTicket | PlainMessage<ResponseIssueTicket> | undefined): boolean;
}

/**
 * @generated from message cafelogos.ticket.RequestRevokeTicket
 */
export declare class RequestRevokeTicket extends Message<RequestRevokeTicket> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  constructor(data?: PartialMessage<RequestRevokeTicket>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "cafelogos.ticket.RequestRevokeTicket";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RequestRevokeTicket;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RequestRevokeTicket;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RequestRevokeTicket;

  static equals(a: RequestRevokeTicket | PlainMessage<RequestRevokeTicket> | undefined, b: RequestRevokeTicket | PlainMessage<RequestRevokeTicket> | undefined): boolean;
}

/**
 * @generated from message cafelogos.ticket.ResponseRevokeTicket
 */
export declare class ResponseRevokeTicket extends Message<ResponseRevokeTicket> {
  constructor(data?: PartialMessage<ResponseRevokeTicket>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "cafelogos.ticket.ResponseRevokeTicket";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ResponseRevokeTicket;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ResponseRevokeTicket;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ResponseRevokeTicket;

  static equals(a: ResponseRevokeTicket | PlainMessage<ResponseRevokeTicket> | undefined, b: ResponseRevokeTicket | PlainMessage<ResponseRevokeTicket> | undefined): boolean;
}

/**
 * @generated from message cafelogos.ticket.Ticket
 */
export declare class Ticket extends Message<Ticket> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * PrefixとNumberを結合したもの
   *
   * @generated from field: string ticket_addr = 2;
   */
  ticketAddr: string;

  /**
   * @generated from field: string created_at = 3;
   */
  createdAt: string;

  constructor(data?: PartialMessage<Ticket>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "cafelogos.ticket.Ticket";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Ticket;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Ticket;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Ticket;

  static equals(a: Ticket | PlainMessage<Ticket> | undefined, b: Ticket | PlainMessage<Ticket> | undefined): boolean;
}

