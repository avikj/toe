package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/toe module sentinel errors
var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrGameNotOpen = sdkerrors.Register(ModuleName, 1101, "game not open")
	ErrGameNotFound = sdkerrors.Register(ModuleName, 1102, "game not found")
	ErrCantPlaySelf = sdkerrors.Register(ModuleName, 1103, "can't play against self")
	ErrPosOutOfBounds = sdkerrors.Register(ModuleName, 1104, "position out of bounds")
	ErrTurn = sdkerrors.Register(ModuleName, 1105, "not your turn")
	ErrPosOccupied = sdkerrors.Register(ModuleName, 1106, "position occupied")
	ErrWaitingForOpponent = sdkerrors.Register(ModuleName, 1107, "waiting for opponent")
	ErrNotInGame = sdkerrors.Register(ModuleName, 1108, "not in this game")
)
