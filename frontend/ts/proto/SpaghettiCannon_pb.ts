// @generated by protoc-gen-es v2.2.5 with parameter "target=ts"
// @generated from file SpaghettiCannon.proto (package SpaghettiCannon.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file SpaghettiCannon.proto.
 */
export const file_SpaghettiCannon: GenFile = /*@__PURE__*/
  fileDesc("ChVTcGFnaGV0dGlDYW5ub24ucHJvdG8SElNwYWdoZXR0aUNhbm5vbi52MSISChBHZXRSZWFkeXpSZXF1ZXN0IiUKEUdldFJlYWR5elJlc3BvbnNlEhAKCGlzX3JlYWR5GAEgASgIIiMKEEFkZFVwZGF0ZVJlcXVlc3QSDwoHY29udGVudBgBIAEoCSITChFBZGRVcGRhdGVSZXNwb25zZTLTAQoZU3BhZ2hldHRpQ2Fubm9uQXBpU2VydmljZRJaCglHZXRSZWFkeXoSJC5TcGFnaGV0dGlDYW5ub24udjEuR2V0UmVhZHl6UmVxdWVzdBolLlNwYWdoZXR0aUNhbm5vbi52MS5HZXRSZWFkeXpSZXNwb25zZSIAEloKCUFkZFVwZGF0ZRIkLlNwYWdoZXR0aUNhbm5vbi52MS5BZGRVcGRhdGVSZXF1ZXN0GiUuU3BhZ2hldHRpQ2Fubm9uLnYxLkFkZFVwZGF0ZVJlc3BvbnNlIgBCN1o1Z2l0aHViLmNvbS9qYW1lc3JlYWQvU3BhZ2hldHRpQ2Fubm9uL2dlbi9wcm90bztzcGFndjFiBnByb3RvMw");

/**
 * @generated from message SpaghettiCannon.v1.GetReadyzRequest
 */
export type GetReadyzRequest = Message<"SpaghettiCannon.v1.GetReadyzRequest"> & {
};

/**
 * Describes the message SpaghettiCannon.v1.GetReadyzRequest.
 * Use `create(GetReadyzRequestSchema)` to create a new message.
 */
export const GetReadyzRequestSchema: GenMessage<GetReadyzRequest> = /*@__PURE__*/
  messageDesc(file_SpaghettiCannon, 0);

/**
 * @generated from message SpaghettiCannon.v1.GetReadyzResponse
 */
export type GetReadyzResponse = Message<"SpaghettiCannon.v1.GetReadyzResponse"> & {
  /**
   * @generated from field: bool is_ready = 1;
   */
  isReady: boolean;
};

/**
 * Describes the message SpaghettiCannon.v1.GetReadyzResponse.
 * Use `create(GetReadyzResponseSchema)` to create a new message.
 */
export const GetReadyzResponseSchema: GenMessage<GetReadyzResponse> = /*@__PURE__*/
  messageDesc(file_SpaghettiCannon, 1);

/**
 * @generated from message SpaghettiCannon.v1.AddUpdateRequest
 */
export type AddUpdateRequest = Message<"SpaghettiCannon.v1.AddUpdateRequest"> & {
  /**
   * @generated from field: string content = 1;
   */
  content: string;
};

/**
 * Describes the message SpaghettiCannon.v1.AddUpdateRequest.
 * Use `create(AddUpdateRequestSchema)` to create a new message.
 */
export const AddUpdateRequestSchema: GenMessage<AddUpdateRequest> = /*@__PURE__*/
  messageDesc(file_SpaghettiCannon, 2);

/**
 * @generated from message SpaghettiCannon.v1.AddUpdateResponse
 */
export type AddUpdateResponse = Message<"SpaghettiCannon.v1.AddUpdateResponse"> & {
};

/**
 * Describes the message SpaghettiCannon.v1.AddUpdateResponse.
 * Use `create(AddUpdateResponseSchema)` to create a new message.
 */
export const AddUpdateResponseSchema: GenMessage<AddUpdateResponse> = /*@__PURE__*/
  messageDesc(file_SpaghettiCannon, 3);

/**
 * @generated from service SpaghettiCannon.v1.SpaghettiCannonApiService
 */
export const SpaghettiCannonApiService: GenService<{
  /**
   * @generated from rpc SpaghettiCannon.v1.SpaghettiCannonApiService.GetReadyz
   */
  getReadyz: {
    methodKind: "unary";
    input: typeof GetReadyzRequestSchema;
    output: typeof GetReadyzResponseSchema;
  },
  /**
   * @generated from rpc SpaghettiCannon.v1.SpaghettiCannonApiService.AddUpdate
   */
  addUpdate: {
    methodKind: "unary";
    input: typeof AddUpdateRequestSchema;
    output: typeof AddUpdateResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_SpaghettiCannon, 0);

