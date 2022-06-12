/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "avikj.toe.toe";

export interface MsgNewGame {
  creator: string;
}

export interface MsgNewGameResponse {
  gameId: number;
}

export interface MsgJoinGame {
  creator: string;
  gameId: number;
}

export interface MsgJoinGameResponse {
  playerX: string;
  playerO: string;
}

const baseMsgNewGame: object = { creator: "" };

export const MsgNewGame = {
  encode(message: MsgNewGame, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgNewGame {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgNewGame } as MsgNewGame;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgNewGame {
    const message = { ...baseMsgNewGame } as MsgNewGame;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: MsgNewGame): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgNewGame>): MsgNewGame {
    const message = { ...baseMsgNewGame } as MsgNewGame;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseMsgNewGameResponse: object = { gameId: 0 };

export const MsgNewGameResponse = {
  encode(
    message: MsgNewGameResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.gameId !== 0) {
      writer.uint32(8).uint64(message.gameId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgNewGameResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgNewGameResponse } as MsgNewGameResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.gameId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgNewGameResponse {
    const message = { ...baseMsgNewGameResponse } as MsgNewGameResponse;
    if (object.gameId !== undefined && object.gameId !== null) {
      message.gameId = Number(object.gameId);
    } else {
      message.gameId = 0;
    }
    return message;
  },

  toJSON(message: MsgNewGameResponse): unknown {
    const obj: any = {};
    message.gameId !== undefined && (obj.gameId = message.gameId);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgNewGameResponse>): MsgNewGameResponse {
    const message = { ...baseMsgNewGameResponse } as MsgNewGameResponse;
    if (object.gameId !== undefined && object.gameId !== null) {
      message.gameId = object.gameId;
    } else {
      message.gameId = 0;
    }
    return message;
  },
};

const baseMsgJoinGame: object = { creator: "", gameId: 0 };

export const MsgJoinGame = {
  encode(message: MsgJoinGame, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.gameId !== 0) {
      writer.uint32(16).uint64(message.gameId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgJoinGame {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgJoinGame } as MsgJoinGame;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.gameId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgJoinGame {
    const message = { ...baseMsgJoinGame } as MsgJoinGame;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.gameId !== undefined && object.gameId !== null) {
      message.gameId = Number(object.gameId);
    } else {
      message.gameId = 0;
    }
    return message;
  },

  toJSON(message: MsgJoinGame): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.gameId !== undefined && (obj.gameId = message.gameId);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgJoinGame>): MsgJoinGame {
    const message = { ...baseMsgJoinGame } as MsgJoinGame;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.gameId !== undefined && object.gameId !== null) {
      message.gameId = object.gameId;
    } else {
      message.gameId = 0;
    }
    return message;
  },
};

const baseMsgJoinGameResponse: object = { playerX: "", playerO: "" };

export const MsgJoinGameResponse = {
  encode(
    message: MsgJoinGameResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.playerX !== "") {
      writer.uint32(10).string(message.playerX);
    }
    if (message.playerO !== "") {
      writer.uint32(18).string(message.playerO);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgJoinGameResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgJoinGameResponse } as MsgJoinGameResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.playerX = reader.string();
          break;
        case 2:
          message.playerO = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgJoinGameResponse {
    const message = { ...baseMsgJoinGameResponse } as MsgJoinGameResponse;
    if (object.playerX !== undefined && object.playerX !== null) {
      message.playerX = String(object.playerX);
    } else {
      message.playerX = "";
    }
    if (object.playerO !== undefined && object.playerO !== null) {
      message.playerO = String(object.playerO);
    } else {
      message.playerO = "";
    }
    return message;
  },

  toJSON(message: MsgJoinGameResponse): unknown {
    const obj: any = {};
    message.playerX !== undefined && (obj.playerX = message.playerX);
    message.playerO !== undefined && (obj.playerO = message.playerO);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgJoinGameResponse>): MsgJoinGameResponse {
    const message = { ...baseMsgJoinGameResponse } as MsgJoinGameResponse;
    if (object.playerX !== undefined && object.playerX !== null) {
      message.playerX = object.playerX;
    } else {
      message.playerX = "";
    }
    if (object.playerO !== undefined && object.playerO !== null) {
      message.playerO = object.playerO;
    } else {
      message.playerO = "";
    }
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  NewGame(request: MsgNewGame): Promise<MsgNewGameResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  JoinGame(request: MsgJoinGame): Promise<MsgJoinGameResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  NewGame(request: MsgNewGame): Promise<MsgNewGameResponse> {
    const data = MsgNewGame.encode(request).finish();
    const promise = this.rpc.request("avikj.toe.toe.Msg", "NewGame", data);
    return promise.then((data) => MsgNewGameResponse.decode(new Reader(data)));
  }

  JoinGame(request: MsgJoinGame): Promise<MsgJoinGameResponse> {
    const data = MsgJoinGame.encode(request).finish();
    const promise = this.rpc.request("avikj.toe.toe.Msg", "JoinGame", data);
    return promise.then((data) => MsgJoinGameResponse.decode(new Reader(data)));
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
