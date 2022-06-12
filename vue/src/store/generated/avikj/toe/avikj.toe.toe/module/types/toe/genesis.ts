/* eslint-disable */
import { Params } from "../toe/params";
import { NextGameId } from "../toe/next_game_id";
import { GameData } from "../toe/game_data";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "avikj.toe.toe";

/** GenesisState defines the toe module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  nextGameId: NextGameId | undefined;
  /** this line is used by starport scaffolding # genesis/proto/state */
  gameDataList: GameData[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    if (message.nextGameId !== undefined) {
      NextGameId.encode(message.nextGameId, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.gameDataList) {
      GameData.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.gameDataList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.nextGameId = NextGameId.decode(reader, reader.uint32());
          break;
        case 3:
          message.gameDataList.push(GameData.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.gameDataList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.nextGameId !== undefined && object.nextGameId !== null) {
      message.nextGameId = NextGameId.fromJSON(object.nextGameId);
    } else {
      message.nextGameId = undefined;
    }
    if (object.gameDataList !== undefined && object.gameDataList !== null) {
      for (const e of object.gameDataList) {
        message.gameDataList.push(GameData.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    message.nextGameId !== undefined &&
      (obj.nextGameId = message.nextGameId
        ? NextGameId.toJSON(message.nextGameId)
        : undefined);
    if (message.gameDataList) {
      obj.gameDataList = message.gameDataList.map((e) =>
        e ? GameData.toJSON(e) : undefined
      );
    } else {
      obj.gameDataList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.gameDataList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.nextGameId !== undefined && object.nextGameId !== null) {
      message.nextGameId = NextGameId.fromPartial(object.nextGameId);
    } else {
      message.nextGameId = undefined;
    }
    if (object.gameDataList !== undefined && object.gameDataList !== null) {
      for (const e of object.gameDataList) {
        message.gameDataList.push(GameData.fromPartial(e));
      }
    }
    return message;
  },
};

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
