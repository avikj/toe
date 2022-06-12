/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../toe/params";
import { NextGameId } from "../toe/next_game_id";
import { GameData } from "../toe/game_data";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "avikj.toe.toe";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetNextGameIdRequest {}

export interface QueryGetNextGameIdResponse {
  NextGameId: NextGameId | undefined;
}

export interface QueryGetGameDataRequest {
  index: string;
}

export interface QueryGetGameDataResponse {
  gameData: GameData | undefined;
}

export interface QueryAllGameDataRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllGameDataResponse {
  gameData: GameData[];
  pagination: PageResponse | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetNextGameIdRequest: object = {};

export const QueryGetNextGameIdRequest = {
  encode(
    _: QueryGetNextGameIdRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetNextGameIdRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetNextGameIdRequest,
    } as QueryGetNextGameIdRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryGetNextGameIdRequest {
    const message = {
      ...baseQueryGetNextGameIdRequest,
    } as QueryGetNextGameIdRequest;
    return message;
  },

  toJSON(_: QueryGetNextGameIdRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetNextGameIdRequest>
  ): QueryGetNextGameIdRequest {
    const message = {
      ...baseQueryGetNextGameIdRequest,
    } as QueryGetNextGameIdRequest;
    return message;
  },
};

const baseQueryGetNextGameIdResponse: object = {};

export const QueryGetNextGameIdResponse = {
  encode(
    message: QueryGetNextGameIdResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.NextGameId !== undefined) {
      NextGameId.encode(message.NextGameId, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetNextGameIdResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetNextGameIdResponse,
    } as QueryGetNextGameIdResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.NextGameId = NextGameId.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNextGameIdResponse {
    const message = {
      ...baseQueryGetNextGameIdResponse,
    } as QueryGetNextGameIdResponse;
    if (object.NextGameId !== undefined && object.NextGameId !== null) {
      message.NextGameId = NextGameId.fromJSON(object.NextGameId);
    } else {
      message.NextGameId = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetNextGameIdResponse): unknown {
    const obj: any = {};
    message.NextGameId !== undefined &&
      (obj.NextGameId = message.NextGameId
        ? NextGameId.toJSON(message.NextGameId)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetNextGameIdResponse>
  ): QueryGetNextGameIdResponse {
    const message = {
      ...baseQueryGetNextGameIdResponse,
    } as QueryGetNextGameIdResponse;
    if (object.NextGameId !== undefined && object.NextGameId !== null) {
      message.NextGameId = NextGameId.fromPartial(object.NextGameId);
    } else {
      message.NextGameId = undefined;
    }
    return message;
  },
};

const baseQueryGetGameDataRequest: object = { index: "" };

export const QueryGetGameDataRequest = {
  encode(
    message: QueryGetGameDataRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetGameDataRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetGameDataRequest,
    } as QueryGetGameDataRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetGameDataRequest {
    const message = {
      ...baseQueryGetGameDataRequest,
    } as QueryGetGameDataRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetGameDataRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetGameDataRequest>
  ): QueryGetGameDataRequest {
    const message = {
      ...baseQueryGetGameDataRequest,
    } as QueryGetGameDataRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetGameDataResponse: object = {};

export const QueryGetGameDataResponse = {
  encode(
    message: QueryGetGameDataResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.gameData !== undefined) {
      GameData.encode(message.gameData, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetGameDataResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetGameDataResponse,
    } as QueryGetGameDataResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.gameData = GameData.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetGameDataResponse {
    const message = {
      ...baseQueryGetGameDataResponse,
    } as QueryGetGameDataResponse;
    if (object.gameData !== undefined && object.gameData !== null) {
      message.gameData = GameData.fromJSON(object.gameData);
    } else {
      message.gameData = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetGameDataResponse): unknown {
    const obj: any = {};
    message.gameData !== undefined &&
      (obj.gameData = message.gameData
        ? GameData.toJSON(message.gameData)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetGameDataResponse>
  ): QueryGetGameDataResponse {
    const message = {
      ...baseQueryGetGameDataResponse,
    } as QueryGetGameDataResponse;
    if (object.gameData !== undefined && object.gameData !== null) {
      message.gameData = GameData.fromPartial(object.gameData);
    } else {
      message.gameData = undefined;
    }
    return message;
  },
};

const baseQueryAllGameDataRequest: object = {};

export const QueryAllGameDataRequest = {
  encode(
    message: QueryAllGameDataRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllGameDataRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllGameDataRequest,
    } as QueryAllGameDataRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllGameDataRequest {
    const message = {
      ...baseQueryAllGameDataRequest,
    } as QueryAllGameDataRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllGameDataRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllGameDataRequest>
  ): QueryAllGameDataRequest {
    const message = {
      ...baseQueryAllGameDataRequest,
    } as QueryAllGameDataRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllGameDataResponse: object = {};

export const QueryAllGameDataResponse = {
  encode(
    message: QueryAllGameDataResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.gameData) {
      GameData.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllGameDataResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllGameDataResponse,
    } as QueryAllGameDataResponse;
    message.gameData = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.gameData.push(GameData.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllGameDataResponse {
    const message = {
      ...baseQueryAllGameDataResponse,
    } as QueryAllGameDataResponse;
    message.gameData = [];
    if (object.gameData !== undefined && object.gameData !== null) {
      for (const e of object.gameData) {
        message.gameData.push(GameData.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllGameDataResponse): unknown {
    const obj: any = {};
    if (message.gameData) {
      obj.gameData = message.gameData.map((e) =>
        e ? GameData.toJSON(e) : undefined
      );
    } else {
      obj.gameData = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllGameDataResponse>
  ): QueryAllGameDataResponse {
    const message = {
      ...baseQueryAllGameDataResponse,
    } as QueryAllGameDataResponse;
    message.gameData = [];
    if (object.gameData !== undefined && object.gameData !== null) {
      for (const e of object.gameData) {
        message.gameData.push(GameData.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a NextGameId by index. */
  NextGameId(
    request: QueryGetNextGameIdRequest
  ): Promise<QueryGetNextGameIdResponse>;
  /** Queries a GameData by index. */
  GameData(request: QueryGetGameDataRequest): Promise<QueryGetGameDataResponse>;
  /** Queries a list of GameData items. */
  GameDataAll(
    request: QueryAllGameDataRequest
  ): Promise<QueryAllGameDataResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("avikj.toe.toe.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  NextGameId(
    request: QueryGetNextGameIdRequest
  ): Promise<QueryGetNextGameIdResponse> {
    const data = QueryGetNextGameIdRequest.encode(request).finish();
    const promise = this.rpc.request("avikj.toe.toe.Query", "NextGameId", data);
    return promise.then((data) =>
      QueryGetNextGameIdResponse.decode(new Reader(data))
    );
  }

  GameData(
    request: QueryGetGameDataRequest
  ): Promise<QueryGetGameDataResponse> {
    const data = QueryGetGameDataRequest.encode(request).finish();
    const promise = this.rpc.request("avikj.toe.toe.Query", "GameData", data);
    return promise.then((data) =>
      QueryGetGameDataResponse.decode(new Reader(data))
    );
  }

  GameDataAll(
    request: QueryAllGameDataRequest
  ): Promise<QueryAllGameDataResponse> {
    const data = QueryAllGameDataRequest.encode(request).finish();
    const promise = this.rpc.request(
      "avikj.toe.toe.Query",
      "GameDataAll",
      data
    );
    return promise.then((data) =>
      QueryAllGameDataResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
