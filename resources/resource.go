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
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

type ImageResource struct {
	*ebiten.Image
}

func (res *ImageResource) BoundX() int {
	if res == nil {
		panic("invalid resource reference when retrieving X bound")
	}

	return res.Bounds().Dx()
}

func (res *ImageResource) BoundY() int {
	if res == nil {
		panic("invalid resource reference when retrieving Y bound")
	}

	return res.Bounds().Dy()
}

type FontResource struct {
	font.Face
	Size int
}
