// Copyright 2022 by xhadzi
// This game demo uses Floppy game example from ebiten repository.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resources

import (
	"bytes"
	_ "embed"
	"errors"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
)

var (
	//go:embed player.png
	player_res []byte

	//go:embed tiles.png
	tiles_res []byte
)

const (
	fontSize      = 24
	titleFontSize = fontSize * 1.5
	smallFontSize = fontSize / 2
)

var (
	PlayerImage ImageResource
	TilesImage  ImageResource

	TitleArcadeFont FontResource
	ArcadeFont      FontResource
	SmallArcadeFont FontResource
)

func loadImageResource(b []byte) (res ImageResource, err error) {
	if img, _, errDecode := image.Decode(bytes.NewReader(b)); errDecode != nil {
		err = errors.New("could not load image resource, reason: " + errDecode.Error())
	} else {
		res = ImageResource{
			Image: ebiten.NewImageFromImage(img),
		}
	}
	return
}

func loadFontResource(fontType *sfnt.Font, fontSize float64) (res FontResource, err error) {
	const dpi = 72

	faceOptions := &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	}

	if titleFont, errFace := opentype.NewFace(fontType, faceOptions); errFace != nil {
		err = errors.New("could not load font resource, reason: " + errFace.Error())
	} else {
		res = FontResource{
			Face: titleFont,
			Size: int(fontSize),
		}
	}
	return
}

func LoadResources() (err error) {
	// images
	playerImageInst, playerErr := loadImageResource(player_res)
	if playerErr != nil {
		err = playerErr
		return
	}

	tilesImageInst, tilesErr := loadImageResource(tiles_res)
	if tilesErr != nil {
		err = tilesErr
		return
	}

	// fonts
	fontType, fontParseErr := opentype.Parse(fonts.PressStart2P_ttf)
	if fontParseErr != nil {
		err = fontParseErr
		return
	}

	titleArcadeFontInst, titleArcadeFontErr := loadFontResource(fontType, titleFontSize)
	if titleArcadeFontErr != nil {
		err = titleArcadeFontErr
		return
	}

	arcadeFontInst, arcadeFontErr := loadFontResource(fontType, fontSize)
	if arcadeFontErr != nil {
		err = arcadeFontErr
		return
	}

	smallFontSizeInst, smallFontSizeErr := loadFontResource(fontType, smallFontSize)
	if smallFontSizeErr != nil {
		err = smallFontSizeErr
		return
	}

	// success
	PlayerImage = playerImageInst
	TilesImage = tilesImageInst
	TitleArcadeFont = titleArcadeFontInst
	ArcadeFont = arcadeFontInst
	SmallArcadeFont = smallFontSizeInst
	return
}
