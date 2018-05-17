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

package mock

import (
	"github.com/sirupsen/logrus"

	"github.com/haakenlabs/ember/gfx"
	"github.com/haakenlabs/ember/pkg/math"
)

var _ gfx.Texture = &Texture{}

type Texture struct {
	size        math.IVec2
	textureType gfx.TextureType
}

func (t *Texture) Bind() {}

func (t *Texture) Unbind() {}

func (t *Texture) Reference() uint32 {
	return 1
}

func (t *Texture) Alloc() error {
	logrus.Infof("alloc mock texture with type %s", t.textureType)

	return nil
}

func (t *Texture) Dealloc() {}

func (t *Texture) Type() gfx.TextureType {
	return t.textureType
}

func (t *Texture) Size() math.IVec2 {
	return t.size
}

func (t *Texture) SetSize(size math.IVec2) {
	t.size = size
}

func (t *Texture) SetResizable(bool) {}

func (t *Texture) Resize() {}

func (t *Texture) Resizable() bool {
	return true
}

func (t *Texture) SetMagFilter(int32) {}

func (t *Texture) SetMinFilter(int32) {}

func (t *Texture) FilterMag() int32 {
	return 0
}

func (t *Texture) FilterMin() int32 {
	return 0
}

func (t *Texture) WrapR() int32 {
	return 0
}

func (t *Texture) WrapS() int32 {
	return 0
}

func (t *Texture) WrapT() int32 {
	return 0
}

func (t *Texture) SetWrapR(int32) {}

func (t *Texture) SetWrapRST(int32, int32, int32) {}

func (t *Texture) SetWrapS(int32) {}

func (t *Texture) SetWrapST(int32, int32) {}

func (t *Texture) SetWrapT(int32) {}

func (t *Texture) Activate(uint32) {}

func (t *Texture) MipLevels() uint32 {
	return 0
}

func (t *Texture) GenerateMipmaps() {}

func (t *Texture) Layers() int32 {
	return 0
}

func (t *Texture) SetLayers(int32) {}

func (r *Renderer) MakeTexture(size math.IVec2, textureType gfx.TextureType) gfx.Texture {
	return &Texture{
		size:        size,
		textureType: textureType,
	}
}
