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
	"errors"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
)

type FontResource struct {
	font.Face
	Size int
}

func NewFontResource(fontType *sfnt.Font, fontSize float64) (res FontResource, err error) {
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
