package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GameControlls struct {
}

type PressedKey ebiten.Key
type PressedKeys []PressedKey

type PressedMouseKey ebiten.MouseButton
type PressedMouseKeys []PressedMouseKey

func NewGameControlls() (controlls GameControlls) {
	controlls = GameControlls{}
	return
}

func (gc *GameControlls) GetPressedKeys() PressedKeys {
	allPressedKeys := inpututil.PressedKeys()
	filteredKeys := make(PressedKeys, 0)

	maxKeyIdx := int(ebiten.KeyMax)
	for _, key := range allPressedKeys {
		if int(key) < maxKeyIdx {
			filteredKeys = append(filteredKeys, PressedKey(key))
		}
	}

	return filteredKeys
}

func (k PressedKey) String() string {
	retStr := ebiten.Key(k).String()
	if retStr == "" {
		retStr = "?????"
	}

	return retStr
}

func (gc *GameControlls) GetPressedMouseKeys() PressedMouseKeys {
	mouseKeys := make(PressedMouseKeys, 0)

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mouseKeys = append(mouseKeys, PressedMouseKey(ebiten.MouseButtonLeft))
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonMiddle) {
		mouseKeys = append(mouseKeys, PressedMouseKey(ebiten.MouseButtonMiddle))
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		mouseKeys = append(mouseKeys, PressedMouseKey(ebiten.MouseButtonRight))
	}

	return mouseKeys
}

func (k PressedMouseKey) String() string {
	switch ebiten.MouseButton(k) {
	case ebiten.MouseButtonLeft:
		return "MouseButtonLeft"
	case ebiten.MouseButtonMiddle:
		return "MouseButtonMiddle"
	case ebiten.MouseButtonRight:
		return "MouseButtonRight"
	default:
		return "?????"
	}
}
