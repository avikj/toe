package gamelogic

import (
	"errors"
	"fmt"
)

// The entire game state can be represented by a 9-character string;
//		The active turn and game completion can be inferred from this string.

type GameStatus int

const (
	NoStatus GameStatus = iota
	XTurn
	OTurn
	XWin
	OWin
	Tie
)

func (gs GameStatus) String() string {
	switch gs {
	case NoStatus:
		return "NoStatus"
	case XTurn:
		return "XTurn"
	case OTurn:
		return "OTurn"
	case XWin:
		return "XWin"
	case OWin:
		return "OWin"
	case Tie:
		return "Tie"
	}
	return "BAD GAMESTATUS"
}

var GameOver = map[GameStatus]bool{
	XTurn: false,
	OTurn: false,
	XWin:  true,
	OWin:  true,
	Tie:   true,
}

const XMarker = byte('x')
const OMarker = byte('o')
const EmptyMarker = byte('_')

var Turn = map[byte]GameStatus{XMarker: XTurn, OMarker: OTurn}

func PlaceMarker(boardState string, marker byte, pos int) (string, error) {
	if pos < 0 || pos >= 9 {
		return boardState, errors.New(fmt.Sprintf("Position out of bounds: %v", pos))
	}

	var status, err = getStatus(boardState)

	if err != nil {
		return boardState, err
	}
	if status != Turn[marker] {
		return boardState, errors.New(fmt.Sprintf("Not %v's turn. Game status: %v", marker, status))
	}
	if boardState[pos] != EmptyMarker {
		return boardState, errors.New(fmt.Sprintf("Position %v already occupied by %v", pos, boardState[pos]))
	}

	var newState = boardState[:pos] + string(marker) + boardState[pos+1:]

	return newState, nil
}

func PlaceO(boardState string, pos int) (string, error) {
	return PlaceMarker(boardState, OMarker, pos)
}

func PlaceX(boardState string, pos int) (newState string, err error) {
	return PlaceMarker(boardState, XMarker, pos)
}

// TODO throw error if string contains unicode characters (generally deal with string/rune/bytes abstraction)
func getStatus(boardState string) (status GameStatus, err error) {
	// returned error is either nil or InvalidBoardState
	if len(boardState) != 9 {
		return NoStatus, errors.New(fmt.Sprintf("Board state must have 9 characters. Invalid board state: %v", boardState))
	}

	var xCount = 0
	var oCount = 0
	var emptyCount = 0

	for _, c := range boardState {
		switch byte(c) {
		case XMarker:
			xCount += 1
		case OMarker:
			oCount += 1
		case EmptyMarker:
			emptyCount += 1
		default:
			return NoStatus, errors.New(fmt.Sprintf("Invalid board state - contains unrecognized character: %v.", string(c)))
		}
	}

	// X goes first and players alternate
	// so number of X's must be the same as or one greater than number of O's
	if !(xCount == oCount || xCount == (oCount+1)) {
		return NoStatus, errors.New("Invalid board state; numbers of X's and O's don't match.")
	} else {
		if didWin(boardState, XMarker) {
			return XWin, nil
		}
		if didWin(boardState, OMarker) {
			return OWin, nil
		}
		if emptyCount == 0 {
			return Tie, nil
		}
		if xCount == oCount {
			return XTurn, nil
		} else { // xCount == oCount + 1
			return OTurn, nil
		}
	}
}

func didWin(validBoardState string, marker byte) bool {
	// checks if marker (either x or o) appears 3x in a row, column, or diagonal
	return (
	//rows
	(validBoardState[0] == marker && validBoardState[1] == marker && validBoardState[2] == marker) ||
		(validBoardState[3] == marker && validBoardState[4] == marker && validBoardState[5] == marker) ||
		(validBoardState[6] == marker && validBoardState[7] == marker && validBoardState[8] == marker) ||
		// cols
		(validBoardState[0] == marker && validBoardState[3] == marker && validBoardState[6] == marker) ||
		(validBoardState[1] == marker && validBoardState[4] == marker && validBoardState[7] == marker) ||
		(validBoardState[2] == marker && validBoardState[5] == marker && validBoardState[8] == marker) ||
		// diags
		(validBoardState[0] == marker && validBoardState[4] == marker && validBoardState[8] == marker) ||
		(validBoardState[2] == marker && validBoardState[4] == marker && validBoardState[6] == marker))
}
