// @generated by protoc-gen-es v2.2.0 with parameter "target=ts"
// @generated from file ticket/ticket_service.proto (package cafelogos.ticket, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file ticket/ticket_service.proto.
 */
export const file_ticket_ticket_service: GenFile = /*@__PURE__*/
  fileDesc("Cht0aWNrZXQvdGlja2V0X3NlcnZpY2UucHJvdG8SEGNhZmVsb2dvcy50aWNrZXQiJAoSUmVxdWVzdElzc3VlVGlja2V0Eg4KBnByZWZpeBgBIAEoCSI/ChNSZXNwb25zZUlzc3VlVGlja2V0EigKBnRpY2tldBgBIAEoCzIYLmNhZmVsb2dvcy50aWNrZXQuVGlja2V0IiEKE1JlcXVlc3RSZXZva2VUaWNrZXQSCgoCaWQYASABKAkiFgoUUmVzcG9uc2VSZXZva2VUaWNrZXQiPQoGVGlja2V0EgoKAmlkGAEgASgJEhMKC3RpY2tldF9hZGRyGAIgASgJEhIKCmNyZWF0ZWRfYXQYAyABKAkyzgEKDVRpY2tldFNlcnZpY2USXAoLSXNzdWVUaWNrZXQSJC5jYWZlbG9nb3MudGlja2V0LlJlcXVlc3RJc3N1ZVRpY2tldBolLmNhZmVsb2dvcy50aWNrZXQuUmVzcG9uc2VJc3N1ZVRpY2tldCIAEl8KDFJldm9rZVRpY2tldBIlLmNhZmVsb2dvcy50aWNrZXQuUmVxdWVzdFJldm9rZVRpY2tldBomLmNhZmVsb2dvcy50aWNrZXQuUmVzcG9uc2VSZXZva2VUaWNrZXQiAEI7WjlnaXRodWIuY29tL0thZ3VyYUdhdGV3YXkvY2FmZWxvZ29zLWdycGMvcGtnL3RpY2tldDt0aWNrZXRiBnByb3RvMw");

/**
 * @generated from message cafelogos.ticket.RequestIssueTicket
 */
export type RequestIssueTicket = Message<"cafelogos.ticket.RequestIssueTicket"> & {
  /**
   * @generated from field: string prefix = 1;
   */
  prefix: string;
};

/**
 * Describes the message cafelogos.ticket.RequestIssueTicket.
 * Use `create(RequestIssueTicketSchema)` to create a new message.
 */
export const RequestIssueTicketSchema: GenMessage<RequestIssueTicket> = /*@__PURE__*/
  messageDesc(file_ticket_ticket_service, 0);

/**
 * @generated from message cafelogos.ticket.ResponseIssueTicket
 */
export type ResponseIssueTicket = Message<"cafelogos.ticket.ResponseIssueTicket"> & {
  /**
   * @generated from field: cafelogos.ticket.Ticket ticket = 1;
   */
  ticket?: Ticket;
};

/**
 * Describes the message cafelogos.ticket.ResponseIssueTicket.
 * Use `create(ResponseIssueTicketSchema)` to create a new message.
 */
export const ResponseIssueTicketSchema: GenMessage<ResponseIssueTicket> = /*@__PURE__*/
  messageDesc(file_ticket_ticket_service, 1);

/**
 * @generated from message cafelogos.ticket.RequestRevokeTicket
 */
export type RequestRevokeTicket = Message<"cafelogos.ticket.RequestRevokeTicket"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;
};

/**
 * Describes the message cafelogos.ticket.RequestRevokeTicket.
 * Use `create(RequestRevokeTicketSchema)` to create a new message.
 */
export const RequestRevokeTicketSchema: GenMessage<RequestRevokeTicket> = /*@__PURE__*/
  messageDesc(file_ticket_ticket_service, 2);

/**
 * @generated from message cafelogos.ticket.ResponseRevokeTicket
 */
export type ResponseRevokeTicket = Message<"cafelogos.ticket.ResponseRevokeTicket"> & {
};

/**
 * Describes the message cafelogos.ticket.ResponseRevokeTicket.
 * Use `create(ResponseRevokeTicketSchema)` to create a new message.
 */
export const ResponseRevokeTicketSchema: GenMessage<ResponseRevokeTicket> = /*@__PURE__*/
  messageDesc(file_ticket_ticket_service, 3);

/**
 * @generated from message cafelogos.ticket.Ticket
 */
export type Ticket = Message<"cafelogos.ticket.Ticket"> & {
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
};

/**
 * Describes the message cafelogos.ticket.Ticket.
 * Use `create(TicketSchema)` to create a new message.
 */
export const TicketSchema: GenMessage<Ticket> = /*@__PURE__*/
  messageDesc(file_ticket_ticket_service, 4);

/**
 * @generated from service cafelogos.ticket.TicketService
 */
export const TicketService: GenService<{
  /**
   * @generated from rpc cafelogos.ticket.TicketService.IssueTicket
   */
  issueTicket: {
    methodKind: "unary";
    input: typeof RequestIssueTicketSchema;
    output: typeof ResponseIssueTicketSchema;
  },
  /**
   * @generated from rpc cafelogos.ticket.TicketService.RevokeTicket
   */
  revokeTicket: {
    methodKind: "unary";
    input: typeof RequestRevokeTicketSchema;
    output: typeof ResponseRevokeTicketSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_ticket_ticket_service, 0);

