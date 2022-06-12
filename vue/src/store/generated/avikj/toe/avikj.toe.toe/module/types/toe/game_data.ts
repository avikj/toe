/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "avikj.toe.toe";

export interface GameData {
  index: string;
  playerX: string;
  playerO: string;
  boardState: string;
  gameId: number;
}

const baseGameData: object = {
  index: "",
  playerX: "",
  playerO: "",
  boardState: "",
  gameId: 0,
};

export const GameData = {
  encode(message: GameData, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.playerX !== "") {
      writer.uint32(18).string(message.playerX);
    }
    if (message.playerO !== "") {
      writer.uint32(26).string(message.playerO);
    }
    if (message.boardState !== "") {
      writer.uint32(34).string(message.boardState);
    }
    if (message.gameId !== 0) {
      writer.uint32(40).uint64(message.gameId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GameData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGameData } as GameData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.playerX = reader.string();
          break;
        case 3:
          message.playerO = reader.string();
          break;
        case 4:
          message.boardState = reader.string();
          break;
        case 5:
          message.gameId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GameData {
    const message = { ...baseGameData } as GameData;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
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
    if (object.boardState !== undefined && object.boardState !== null) {
      message.boardState = String(object.boardState);
    } else {
      message.boardState = "";
    }
    if (object.gameId !== undefined && object.gameId !== null) {
      message.gameId = Number(object.gameId);
    } else {
      message.gameId = 0;
    }
    return message;
  },

  toJSON(message: GameData): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.playerX !== undefined && (obj.playerX = message.playerX);
    message.playerO !== undefined && (obj.playerO = message.playerO);
    message.boardState !== undefined && (obj.boardState = message.boardState);
    message.gameId !== undefined && (obj.gameId = message.gameId);
    return obj;
  },

  fromPartial(object: DeepPartial<GameData>): GameData {
    const message = { ...baseGameData } as GameData;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
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
    if (object.boardState !== undefined && object.boardState !== null) {
      message.boardState = object.boardState;
    } else {
      message.boardState = "";
    }
    if (object.gameId !== undefined && object.gameId !== null) {
      message.gameId = object.gameId;
    } else {
      message.gameId = 0;
    }
    return message;
  },
};

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
