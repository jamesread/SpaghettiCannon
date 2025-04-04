// @generated by protoc-gen-connect-es v1.5.0 with parameter "target=ts"
// @generated from file SpaghettiCannon.proto (package SpaghettiCannon.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { AddUpdateRequest, AddUpdateResponse, GetReadyzRequest, GetReadyzResponse } from "./SpaghettiCannon_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service SpaghettiCannon.v1.SpaghettiCannonApiService
 */
export const SpaghettiCannonApiService = {
  typeName: "SpaghettiCannon.v1.SpaghettiCannonApiService",
  methods: {
    /**
     * @generated from rpc SpaghettiCannon.v1.SpaghettiCannonApiService.GetReadyz
     */
    getReadyz: {
      name: "GetReadyz",
      I: GetReadyzRequest,
      O: GetReadyzResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc SpaghettiCannon.v1.SpaghettiCannonApiService.AddUpdate
     */
    addUpdate: {
      name: "AddUpdate",
      I: AddUpdateRequest,
      O: AddUpdateResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

