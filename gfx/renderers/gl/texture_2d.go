/*
Copyright (c) 2018 HaakenLabs

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package gl

import (
	"unsafe"

	"github.com/go-gl/gl/v4.3-core/gl"

	"github.com/haakenlabs/ember/gfx"
	"github.com/haakenlabs/ember/system/instance"
)

var _ gfx.Texture = &Texture2D{}

type Texture2D struct {
	BaseTexture

	data    []uint8
	hdrData []float32
}

func NewTexture2D(cfg *gfx.TextureConfig) *Texture2D {
	t := &Texture2D{}

	t.textureType = gl.TEXTURE_2D

	t.SetName("Texture2D")
	instance.MustAssign(t)

	t.size = cfg.Size
	t.uploadFunc = t.Upload

	t.internalFormat = TextureFormatToInternal(cfg.Format)
	t.glFormat = TextureFormatToFormat(cfg.Format)
	t.storageFormat = TextureFormatToStorage(cfg.Format)

	return t
}

func (t *Texture2D) Type() gfx.TextureType {
	return gfx.Texture2D
}

func (t *Texture2D) Upload() {
	t.Bind()

	var ptr unsafe.Pointer

	if t.hdrData != nil && len(t.hdrData) > 0 {
		ptr = gl.Ptr(t.hdrData)
	} else if len(t.data) > 0 {
		ptr = gl.Ptr(t.data)
	}

	gl.TexImage2D(gl.TEXTURE_2D, 0, t.internalFormat, t.size.X(), t.size.Y(), 0, t.glFormat, t.storageFormat, ptr)
}

func (t *Texture2D) SetData(data []uint8) {
	t.data = data
}

func (t *Texture2D) SetHDRData(data []float32) {
	t.hdrData = data
}
