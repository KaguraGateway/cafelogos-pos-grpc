// @generated by protoc-gen-connect-es v1.6.1 with parameter "target=ts+js+dts"
// @generated from file ticket/ticket_service.proto (package cafelogos.ticket, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { RequestIssueTicket, RequestRevokeTicket, ResponseIssueTicket, ResponseRevokeTicket } from "./ticket_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service cafelogos.ticket.TicketService
 */
export const TicketService = {
  typeName: "cafelogos.ticket.TicketService",
  methods: {
    /**
     * @generated from rpc cafelogos.ticket.TicketService.IssueTicket
     */
    issueTicket: {
      name: "IssueTicket",
      I: RequestIssueTicket,
      O: ResponseIssueTicket,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc cafelogos.ticket.TicketService.RevokeTicket
     */
    revokeTicket: {
      name: "RevokeTicket",
      I: RequestRevokeTicket,
      O: ResponseRevokeTicket,
      kind: MethodKind.Unary,
    },
  }
} as const;

