package engine

import (
	"JungleRun/resources"
	"image"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth      = 640
	screenHeight     = 480
	tileSize         = 32
	pipeWidth        = tileSize * 2
	pipeStartOffsetX = 8
	pipeIntervalX    = 8
	pipeGapY         = 5
)

type Mode int

const (
	ModeTitle Mode = iota
	ModeGame
	ModeGameOver
)

type Game struct {
	screen    *Screen
	controlls GameControlls
	content   *GameContent
	mode      Mode

	// The players's position
	objPos ObjectPosition

	// Camera
	cameraX int
	cameraY int

	// Pipes
	pipeTileYs []int

	gameoverCount int
}

func NewGame(screen *Screen) ebiten.Game {
	g := &Game{
		screen:    screen,
		controlls: NewGameControlls(),
		content:   NewGameContent(screen),
	}
	g.init()
	return g
}

func (g *Game) init() {
	// initial position
	g.objPos.xPos = 0
	g.objPos.yPos.Add(float64(g.screen.screenHeight) / 4.0)
	g.objPos.gravity = 0

	g.cameraX = -240
	g.cameraY = 0
	g.pipeTileYs = make([]int, 256)
	for i := range g.pipeTileYs {
		g.pipeTileYs[i] = rand.Intn(6) + 2
	}
}

func (g *Game) isKeyJustPressed() bool {
	if len(g.controlls.GetPressedKeys()) > 0 {
		return true
	} else if len(g.controlls.GetPressedMouseKeys()) > 0 {
		return true
	}
	return false
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screen.Width(), g.screen.Height()
}

func (g *Game) UpdateTitle() (err error) {
	if g.isKeyJustPressed() {
		g.mode = ModeGame
	}
	return
}

func (g *Game) UpdateGame() (err error) {
	g.objPos.xPos.Add(2)

	g.cameraX += 2
	if g.isKeyJustPressed() {
		g.objPos.gravity.Reset(-6.0)
		if err = resources.JumpAudio.PlayFromStart(); err != nil {
			return
		}
	}
	g.objPos.yPos.Add(g.objPos.gravity.Value())

	// Gravity
	g.objPos.gravity.Add(0.25)
	if g.objPos.gravity.Value() > 6 {
		g.objPos.gravity.Reset(6)
	}

	if g.hit() {
		if err = resources.HitAudio.PlayFromStart(); err != nil {
			return
		}
		g.mode = ModeGameOver
		g.gameoverCount = 30
	}
	return
}

func (g *Game) UpdateOnGameOver() (err error) {
	// delay initiation of the new game for a while so that it does not register
	// last held key during cras as a signal for a new game
	if g.gameoverCount > 0 {
		g.gameoverCount--
	}

	if g.gameoverCount == 0 && g.isKeyJustPressed() {
		g.init()
		g.mode = ModeTitle
	}
	return
}

func (g *Game) Update() error {
	switch g.mode {
	case ModeTitle:
		return g.UpdateTitle()
	case ModeGame:
		return g.UpdateGame()
	case ModeGameOver:
		return g.UpdateOnGameOver()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.content.DrawBackgroundColor(screen)
	g.drawTiles(screen)

	if g.mode != ModeTitle {
		g.drawPlayer(screen)
	}

	if g.mode == ModeTitle {
		g.content.DrawCenterText(screen, "Jungle run", []string{"PRESS SPACE KEY", "OR A/B BUTTON", "OR TOUCH SCREEN"})
		g.content.DrawNameAndCopyright(screen, "Go Jungle run", "licenced under CC BY 3.0.")
	} else if g.mode == ModeGameOver {
		g.content.DrawCenterText(screen, "", []string{"GAME OVER!"})
	}

	g.content.DrawScore(screen, g.score())
	g.content.DrawTPS(screen)
}

func (g *Game) pipeAt(tileX int) (tileY int, ok bool) {
	if (tileX - pipeStartOffsetX) <= 0 {
		return 0, false
	}
	if floorMod(tileX-pipeStartOffsetX, pipeIntervalX) != 0 {
		return 0, false
	}
	idx := floorDiv(tileX-pipeStartOffsetX, pipeIntervalX)
	return g.pipeTileYs[idx%len(g.pipeTileYs)], true
}

func (g *Game) score() int {
	x := g.objPos.xPos.ValueFloored() / tileSize
	if (x - pipeStartOffsetX) <= 0 {
		return 0
	}
	return floorDiv(x-pipeStartOffsetX, pipeIntervalX)
}

func (g *Game) hit() bool {
	if g.mode != ModeGame {
		return false
	}
	const (
		PlayerImageWidth  = 30
		playerImageHeight = 60
	)
	w, h := resources.PlayerImage.BoundX(), resources.PlayerImage.BoundY()
	x0 := g.objPos.xPos.ValueFloored() + (w-PlayerImageWidth)/2
	y0 := g.objPos.yPos.ValueFloored() + (h-playerImageHeight)/2
	x1 := x0 + PlayerImageWidth
	y1 := y0 + playerImageHeight
	if y0 < -tileSize*4 {
		return true
	}
	if y1 >= g.screen.Height()-tileSize {
		return true
	}
	xMin := floorDiv(x0-pipeWidth, tileSize)
	xMax := floorDiv(x0+PlayerImageWidth, tileSize)
	for x := xMin; x <= xMax; x++ {
		y, ok := g.pipeAt(x)
		if !ok {
			continue
		}
		if x0 >= x*tileSize+pipeWidth {
			continue
		}
		if x1 < x*tileSize {
			continue
		}
		if y0 < y*tileSize {
			return true
		}
		if y1 >= (y+pipeGapY)*tileSize {
			return true
		}
	}
	return false
}

func (g *Game) drawTiles(screen *ebiten.Image) {
	var (
		nx           = g.screen.Width() / tileSize
		ny           = g.screen.Height() / tileSize
		pipeTileSrcX = 128
		pipeTileSrcY = 192
	)

	op := &ebiten.DrawImageOptions{}
	for i := -2; i < nx+1; i++ {
		// ground
		op.GeoM.Reset()
		op.GeoM.Translate(float64(i*tileSize-floorMod(g.cameraX, tileSize)),
			float64((ny-1)*tileSize-floorMod(g.cameraY, tileSize)))
		screen.DrawImage(resources.TilesImage.SubImage(image.Rect(0, 0, tileSize, tileSize)).(*ebiten.Image), op)

		// pipe
		if tileY, ok := g.pipeAt(floorDiv(g.cameraX, tileSize) + i); ok {
			for j := 0; j < tileY; j++ {
				op.GeoM.Reset()
				op.GeoM.Scale(1, -1)
				op.GeoM.Translate(float64(i*tileSize-floorMod(g.cameraX, tileSize)),
					float64(j*tileSize-floorMod(g.cameraY, tileSize)))
				op.GeoM.Translate(0, tileSize)
				var r image.Rectangle
				if j == tileY-1 {
					r = image.Rect(pipeTileSrcX, pipeTileSrcY, pipeTileSrcX+tileSize*2, pipeTileSrcY+tileSize)
				} else {
					r = image.Rect(pipeTileSrcX, pipeTileSrcY+tileSize, pipeTileSrcX+tileSize*2, pipeTileSrcY+tileSize*2)
				}
				screen.DrawImage(resources.TilesImage.SubImage(r).(*ebiten.Image), op)
			}
			for j := tileY + pipeGapY; j < g.screen.Height()/tileSize-1; j++ {
				op.GeoM.Reset()
				op.GeoM.Translate(float64(i*tileSize-floorMod(g.cameraX, tileSize)),
					float64(j*tileSize-floorMod(g.cameraY, tileSize)))
				var r image.Rectangle
				if j == tileY+pipeGapY {
					r = image.Rect(pipeTileSrcX, pipeTileSrcY, pipeTileSrcX+pipeWidth, pipeTileSrcY+tileSize)
				} else {
					r = image.Rect(pipeTileSrcX, pipeTileSrcY+tileSize, pipeTileSrcX+pipeWidth, pipeTileSrcY+tileSize+tileSize)
				}
				screen.DrawImage(resources.TilesImage.SubImage(r).(*ebiten.Image), op)
			}
		}
	}
}

func (g *Game) drawPlayer(screen *ebiten.Image) {
	g.content.DrawObject(screen, resources.PlayerImage, ObjectPosition{xPos: g.objPos.xPos, yPos: g.objPos.yPos, gravity: g.objPos.gravity}, CameraPosition{cameraX: g.cameraX, cameraY: g.cameraY})
}
