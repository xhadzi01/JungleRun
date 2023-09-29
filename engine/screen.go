package engine

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
)

type Screen struct {
	screenWidth  int
	screenHeight int
}

func NewScreen(screenWidth, screenHeight int) (screen *Screen) {
	return &Screen{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
	}
}

func (screen *Screen) Initialize(title string) (err error) {
	if screen == nil {
		err = errors.New("screen reference is invalid")
		return
	}

	ebiten.SetWindowSize(screen.screenWidth, screen.screenHeight)
	ebiten.SetWindowTitle(title)
	return
}

func (screen *Screen) Width() int {
	if screen == nil {
		panic("screen reference is invalid")
	}
	return screen.screenWidth
}

func (screen *Screen) Height() int {
	if screen == nil {
		panic("screen reference is invalid")
	}
	return screen.screenHeight
}
