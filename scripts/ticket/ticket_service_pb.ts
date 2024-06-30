// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file ticket/ticket_service.proto (package cafelogos.ticket, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message cafelogos.ticket.RequestIssueTicket
 */
export class RequestIssueTicket extends Message<RequestIssueTicket> {
  /**
   * @generated from field: string prefix = 1;
   */
  prefix = "";

  constructor(data?: PartialMessage<RequestIssueTicket>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "cafelogos.ticket.RequestIssueTicket";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "prefix", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RequestIssueTicket {
    return new RequestIssueTicket().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RequestIssueTicket {
    return new RequestIssueTicket().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RequestIssueTicket {
    return new RequestIssueTicket().fromJsonString(jsonString, options);
  }

  static equals(a: RequestIssueTicket | PlainMessage<RequestIssueTicket> | undefined, b: RequestIssueTicket | PlainMessage<RequestIssueTicket> | undefined): boolean {
    return proto3.util.equals(RequestIssueTicket, a, b);
  }
}

/**
 * @generated from message cafelogos.ticket.ResponseIssueTicket
 */
export class ResponseIssueTicket extends Message<ResponseIssueTicket> {
  /**
   * @generated from field: cafelogos.ticket.Ticket ticket = 1;
   */
  ticket?: Ticket;

  constructor(data?: PartialMessage<ResponseIssueTicket>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "cafelogos.ticket.ResponseIssueTicket";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "ticket", kind: "message", T: Ticket },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ResponseIssueTicket {
    return new ResponseIssueTicket().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ResponseIssueTicket {
    return new ResponseIssueTicket().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ResponseIssueTicket {
    return new ResponseIssueTicket().fromJsonString(jsonString, options);
  }

  static equals(a: ResponseIssueTicket | PlainMessage<ResponseIssueTicket> | undefined, b: ResponseIssueTicket | PlainMessage<ResponseIssueTicket> | undefined): boolean {
    return proto3.util.equals(ResponseIssueTicket, a, b);
  }
}

/**
 * @generated from message cafelogos.ticket.RequestRevokeTicket
 */
export class RequestRevokeTicket extends Message<RequestRevokeTicket> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<RequestRevokeTicket>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "cafelogos.ticket.RequestRevokeTicket";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RequestRevokeTicket {
    return new RequestRevokeTicket().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RequestRevokeTicket {
    return new RequestRevokeTicket().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RequestRevokeTicket {
    return new RequestRevokeTicket().fromJsonString(jsonString, options);
  }

  static equals(a: RequestRevokeTicket | PlainMessage<RequestRevokeTicket> | undefined, b: RequestRevokeTicket | PlainMessage<RequestRevokeTicket> | undefined): boolean {
    return proto3.util.equals(RequestRevokeTicket, a, b);
  }
}

/**
 * @generated from message cafelogos.ticket.ResponseRevokeTicket
 */
export class ResponseRevokeTicket extends Message<ResponseRevokeTicket> {
  constructor(data?: PartialMessage<ResponseRevokeTicket>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "cafelogos.ticket.ResponseRevokeTicket";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ResponseRevokeTicket {
    return new ResponseRevokeTicket().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ResponseRevokeTicket {
    return new ResponseRevokeTicket().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ResponseRevokeTicket {
    return new ResponseRevokeTicket().fromJsonString(jsonString, options);
  }

  static equals(a: ResponseRevokeTicket | PlainMessage<ResponseRevokeTicket> | undefined, b: ResponseRevokeTicket | PlainMessage<ResponseRevokeTicket> | undefined): boolean {
    return proto3.util.equals(ResponseRevokeTicket, a, b);
  }
}

/**
 * @generated from message cafelogos.ticket.Ticket
 */
export class Ticket extends Message<Ticket> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * PrefixとNumberを結合したもの
   *
   * @generated from field: string ticket_addr = 2;
   */
  ticketAddr = "";

  /**
   * @generated from field: string created_at = 3;
   */
  createdAt = "";

  constructor(data?: PartialMessage<Ticket>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "cafelogos.ticket.Ticket";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "ticket_addr", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "created_at", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Ticket {
    return new Ticket().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Ticket {
    return new Ticket().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Ticket {
    return new Ticket().fromJsonString(jsonString, options);
  }

  static equals(a: Ticket | PlainMessage<Ticket> | undefined, b: Ticket | PlainMessage<Ticket> | undefined): boolean {
    return proto3.util.equals(Ticket, a, b);
  }
}

