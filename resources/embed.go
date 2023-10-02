// Copyright 2023 by xhadzi
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
	_ "embed"

	raudio "github.com/hajimehoshi/ebiten/v2/examples/resources/audio"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font/opentype"
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

	JumpAudio AudioResource
	HitAudio  AudioResource
)

func LoadResources() (err error) {
	// images
	playerImageInst, playerErr := NewImageResource(player_res)
	if playerErr != nil {
		err = playerErr
		return
	}

	tilesImageInst, tilesErr := NewImageResource(tiles_res)
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

	titleArcadeFontInst, titleArcadeFontErr := NewFontResource(fontType, titleFontSize)
	if titleArcadeFontErr != nil {
		err = titleArcadeFontErr
		return
	}

	arcadeFontInst, arcadeFontErr := NewFontResource(fontType, fontSize)
	if arcadeFontErr != nil {
		err = arcadeFontErr
		return
	}

	smallFontSizeInst, smallFontSizeErr := NewFontResource(fontType, smallFontSize)
	if smallFontSizeErr != nil {
		err = smallFontSizeErr
		return
	}

	jumpAudioInst, jumpAudioErr := NewAudioResource(raudio.Jump_ogg, decoder_oog)
	if jumpAudioErr != nil {
		err = jumpAudioErr
		return
	}

	hitAudioInst, hitAudioErr := NewAudioResource(raudio.Jab_wav, decoder_wav)
	if hitAudioErr != nil {
		err = hitAudioErr
		return
	}

	// success
	PlayerImage = playerImageInst
	TilesImage = tilesImageInst
	TitleArcadeFont = titleArcadeFontInst
	ArcadeFont = arcadeFontInst
	SmallArcadeFont = smallFontSizeInst
	JumpAudio = jumpAudioInst
	HitAudio = hitAudioInst
	return
}
