// Copyright 2022 by xhadzi
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
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed player.png
	player_res []byte

	//go:embed tiles.png
	tiles_res []byte
)

var (
	PlayerImage *ebiten.Image
	TilesImage  *ebiten.Image
)

func LoadResources() (err error) {
	img, _, err := image.Decode(bytes.NewReader(player_res))
	if err != nil {
		return
	}
	PlayerImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(tiles_res))
	if err != nil {
		return
	}
	TilesImage = ebiten.NewImageFromImage(img)

	return
}
