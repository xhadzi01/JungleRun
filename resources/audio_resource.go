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
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

type decoderIdx uint
type decorerFunctorType func(src io.Reader) (*audio.Player, error)

const (
	decoder_wav decoderIdx = iota
	decoder_oog
)

var decoders = map[decoderIdx]decorerFunctorType{
	decoder_wav: func(src io.Reader) (player *audio.Player, err error) {

		if stream, err := wav.DecodeWithoutResampling(src); err == nil {
			player, err = audioContext.NewPlayer(stream)
		}
		return
	},
	decoder_oog: func(src io.Reader) (player *audio.Player, err error) {

		if stream, err := vorbis.DecodeWithoutResampling(src); err == nil {
			player, err = audioContext.NewPlayer(stream)
		}
		return
	},
}

var audioContext *audio.Context

type AudioResource struct {
	*audio.Player
}

func NewAudioResource(b []byte, decoder decoderIdx) (res AudioResource, err error) {
	if audioContext == nil {
		audioContext = audio.NewContext(48000)
		panic_handle := recover()
		if panic_handle != nil {
			err = errors.New("could not create audio context")
			return
		}
	}

	decoderFunctor, found := decoders[decoder]
	if !found {
		err = fmt.Errorf("could not find decoder with index %d", uint(decoder))
		return
	}

	resInst, err := decoderFunctor(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
		return
	}

	res = AudioResource{
		Player: resInst,
	}
	return
}

func (res *AudioResource) PlayFromStart() (err error) {
	if res == nil {
		panic("invalid resource reference when playing audio")
	}

	if err = res.Player.Rewind(); err != nil {
		return
	}
	res.Player.Play()
	return
}
