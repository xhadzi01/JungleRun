package engine

import (
	"JungleRun/resources"
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type ObjectsPosition struct {
	x16  int
	y16  int
	vy16 int
}

type CameraPosition struct {
	cameraX int
	cameraY int
}

type GameContent struct {
	screen *Screen
	mode   Mode

	titleTextFont   resources.FontResource
	textFont        resources.FontResource
	smallTextFont   resources.FontResource
	textColor       color.Color
	backgroundColor color.RGBA

	// The players's position
	x16  int
	y16  int
	vy16 int

	// Camera
	cameraX int
	cameraY int

	// Pipes
	pipeTileYs []int

	gameoverCount int

	touchIDs   []ebiten.TouchID
	gamepadIDs []ebiten.GamepadID

	audioContext *audio.Context
	jumpPlayer   *audio.Player
	hitPlayer    *audio.Player
}

func NewGameContent(screen *Screen) (gc *GameContent) {

	gc = &GameContent{
		screen:          screen,
		titleTextFont:   resources.TitleArcadeFont,
		textFont:        resources.ArcadeFont,
		smallTextFont:   resources.SmallArcadeFont,
		textColor:       color.White,
		backgroundColor: color.RGBA{0x80, 0xa0, 0xc0, 0xff},
	}
	return
}

func (gc *GameContent) DrawBackgroundColor(screenImage *ebiten.Image) {
	if gc == nil {
		panic("screen reference is invalid")
	}
	screenImage.Fill(gc.backgroundColor)
}

func (gc *GameContent) DrawObject(screenImage *ebiten.Image, objectImage resources.ImageResource, objPos ObjectsPosition, camPos CameraPosition) {
	if gc == nil {
		panic("screen reference is invalid")
	}

	var (
		imageWidth  int = objectImage.BoundX()
		imageHeight int = objectImage.BoundY()
	)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(imageWidth)/2.0, -float64(imageHeight)/2.0)
	op.GeoM.Rotate(float64(objPos.vy16) / 96.0 * math.Pi / 6)
	op.GeoM.Translate(float64(imageWidth)/2.0, float64(imageHeight)/2.0)
	op.GeoM.Translate(float64(objPos.x16/16.0)-float64(camPos.cameraX), float64(objPos.y16/16.0)-float64(camPos.cameraY))
	op.Filter = ebiten.FilterLinear
	screenImage.DrawImage(objectImage.Image, op)
}

func (gc *GameContent) DrawCenterText(screenImage *ebiten.Image, titleText string, additionalText []string) {
	if gc == nil {
		panic("screen reference is invalid")
	}

	// title
	// title is offset from the top of the screen by 4 title text sizes
	titleScreenXOffset := (gc.screen.Width() - len(titleText)*gc.titleTextFont.Size) / 2
	titleScreenYOffset := 4 * gc.titleTextFont.Size
	if titleText != "" {
		text.Draw(screenImage, titleText, gc.titleTextFont, titleScreenXOffset, titleScreenYOffset, gc.textColor)
	}

	// additional text
	// remaining text is filled right under the title block
	addTextScreenYAdditiveOffsetFromTitle := titleScreenYOffset + gc.titleTextFont.Size*2
	for textIdx, addText := range additionalText {
		addTextScreenXOffset := (gc.screen.Width() - len(addText)*gc.textFont.Size) / 2
		addTextScreenYOffset := addTextScreenYAdditiveOffsetFromTitle + textIdx*gc.textFont.Size*2
		if addText == "" {
			continue
		}

		text.Draw(screenImage, addText, gc.textFont, addTextScreenXOffset, addTextScreenYOffset, gc.textColor)
	}
}

func (gc *GameContent) DrawNameAndCopyright(screenImage *ebiten.Image, gameName, copyright string) {
	if gc == nil {
		panic("screen reference is invalid")
	}

	// game name
	gameNameScreenXOffset := (gc.screen.Width() - len(gameName)*gc.smallTextFont.Size) / 2
	gameNameScreenYOffset := gc.screen.Height() - 6 - gc.smallTextFont.Size
	text.Draw(screenImage, gameName, gc.smallTextFont, gameNameScreenXOffset, gameNameScreenYOffset, gc.textColor)

	// game copyright
	gameCopyrightScreenXOffset := (gc.screen.Width() - len(copyright)*gc.smallTextFont.Size) / 2
	gameCopyrightScreenYOffset := gc.screen.Height() - 4
	text.Draw(screenImage, copyright, gc.smallTextFont, gameCopyrightScreenXOffset, gameCopyrightScreenYOffset, gc.textColor)
}

func (gc *GameContent) DrawScore(screenImage *ebiten.Image, score int) {
	if gc == nil {
		panic("screen reference is invalid")
	}

	scoreStr := fmt.Sprintf("%04d", score)
	scoreStrLen := len(scoreStr)
	// upper right corner
	scoreStrScreenOffset := gc.screen.Width() - scoreStrLen*gc.textFont.Size

	text.Draw(screenImage, scoreStr, gc.textFont, scoreStrScreenOffset, gc.textFont.Size, gc.textColor)
}

func (gc *GameContent) DrawTPS(screenImage *ebiten.Image) {
	if gc == nil {
		panic("screen reference is invalid")
	}

	ebitenutil.DebugPrint(screenImage, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}
