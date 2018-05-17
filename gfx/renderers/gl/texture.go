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
	"github.com/go-gl/gl/v4.3-core/gl"

	"github.com/haakenlabs/ember/core"
	"github.com/haakenlabs/ember/gfx"
	"github.com/haakenlabs/ember/pkg/math"
)

type BaseTexture struct {
	core.BaseObject

	uploadFunc     func()
	internalFormat int32
	storageFormat  uint32
	glFormat       uint32
	filterMag      int32
	filterMin      int32
	wrapR          int32
	wrapS          int32
	wrapT          int32
	layers         int32
	reference      uint32
	textureFormat  gfx.TextureFormat
	size           math.IVec2
	resizable      bool
	textureType    uint32
}

func (t *BaseTexture) Bind() {
	gl.BindTexture(t.textureType, t.reference)
}

func (t *BaseTexture) Unbind() {}

func (t *BaseTexture) Reference() uint32 {
	return t.reference
}

func (t *BaseTexture) Alloc() error {
	if t.reference != 0 {
		return nil
	}

	gl.GenTextures(1, &t.reference)

	t.filterMag = gl.LINEAR
	t.filterMin = gl.LINEAR
	t.wrapR = gl.CLAMP_TO_EDGE
	t.wrapS = gl.CLAMP_TO_EDGE
	t.wrapT = gl.CLAMP_TO_EDGE
	t.resizable = true
	t.layers = 1

	t.uploadFunc()

	t.SetFilter(t.filterMag, t.filterMin)
	t.SetWrapRST(t.wrapR, t.wrapS, t.wrapT)

	return nil
}

func (t *BaseTexture) Dealloc() {
	if t.reference != 0 {
		gl.DeleteTextures(1, &t.reference)
		t.reference = 0
	}
}

func (t *BaseTexture) Activate(unit uint32) {
	gl.ActiveTexture(gl.TEXTURE0 + unit)
	t.Bind()
}

func (t *BaseTexture) Size() math.IVec2 {
	return t.size
}

func (t *BaseTexture) SetSize(size math.IVec2) {
	if !t.resizable {
		return
	}
	if size.X() <= 0 || size.Y() <= 0 {
		return
	}

	t.size = size
	t.uploadFunc()
}

func (t *BaseTexture) SetResizable(resizable bool) {
	t.resizable = resizable
}

func (t *BaseTexture) Resize() {}

func (t *BaseTexture) Resizable() bool {
	return t.resizable
}

// FilterMag
func (t *BaseTexture) FilterMag() int32 {
	return t.filterMag
}

// FilterMin
func (t *BaseTexture) FilterMin() int32 {
	return t.filterMin
}

// GenerateMipmaps
func (t *BaseTexture) GenerateMipmaps() {

}

// Height
func (t *BaseTexture) Height() int32 {
	return t.size.Y()
}

// Layers
func (t *BaseTexture) Layers() int32 {
	return t.layers
}

func (t *BaseTexture) SetLayers(layers int32) {
	t.layers = layers
}

// MipLevels
func (t *BaseTexture) MipLevels() uint32 {
	return 1
}

// SetFilter
func (t *BaseTexture) SetFilter(magFilter, minFilter int32) {
	t.SetMagFilter(magFilter)
	t.SetMinFilter(minFilter)
}

// SetMagFilter
func (t *BaseTexture) SetMagFilter(magFilter int32) {
	t.filterMag = magFilter
	gl.TexParameteri(t.textureType, gl.TEXTURE_MAG_FILTER, t.filterMag)
}

// SetMinFilter
func (t *BaseTexture) SetMinFilter(minFilter int32) {
	t.filterMin = minFilter
	gl.TexParameteri(t.textureType, gl.TEXTURE_MIN_FILTER, t.filterMin)
}

// SetGLFormats
func (t *BaseTexture) SetGLFormats(internalFormat int32, format uint32, storageFormat uint32) {
	t.internalFormat = internalFormat
	t.glFormat = format
	t.storageFormat = storageFormat

	//t.Upload()
	t.uploadFunc()
}

// SetTexFormat
func (t *BaseTexture) SetFormat(format gfx.TextureFormat) {
	t.SetGLFormats(TextureFormatToInternal(format), TextureFormatToFormat(format), TextureFormatToStorage(format))
}

// SetWrapR
func (t *BaseTexture) SetWrapR(wrapR int32) {
	t.wrapR = wrapR
	gl.TexParameteri(t.textureType, gl.TEXTURE_WRAP_R, t.wrapR)
	if t.wrapR == gl.CLAMP_TO_BORDER {
		color := [4]float32{}
		gl.TexParameterfv(t.textureType, gl.TEXTURE_BORDER_COLOR, &color[0])
	}
}

// SetWrapRST
func (t *BaseTexture) SetWrapRST(wrapR, wrapS, wrapT int32) {
	t.SetWrapR(wrapR)
	t.SetWrapS(wrapS)
	t.SetWrapT(wrapT)
}

// SetWrapS
func (t *BaseTexture) SetWrapS(wrapS int32) {
	t.wrapS = wrapS
	gl.TexParameteri(t.textureType, gl.TEXTURE_WRAP_S, t.wrapS)
	if t.wrapS == gl.CLAMP_TO_BORDER {
		color := [4]float32{}
		gl.TexParameterfv(t.textureType, gl.TEXTURE_BORDER_COLOR, &color[0])
	}
}

// SetWrapST
func (t *BaseTexture) SetWrapST(wrapS, wrapT int32) {
	t.SetWrapS(wrapS)
	t.SetWrapT(wrapT)
}

// SetWrapT
func (t *BaseTexture) SetWrapT(wrapT int32) {
	t.wrapT = wrapT
	gl.TexParameteri(t.textureType, gl.TEXTURE_WRAP_T, t.wrapT)
	if t.wrapT == gl.CLAMP_TO_BORDER {
		color := [4]float32{}
		gl.TexParameterfv(t.textureType, gl.TEXTURE_BORDER_COLOR, &color[0])
	}
}

// TexFormat
func (t *BaseTexture) Format() gfx.TextureFormat {
	return t.textureFormat
}

// Upload
func (t *BaseTexture) Upload() {
	panic("Unimplemented!")
}

// Width
func (t *BaseTexture) Width() int32 {
	return t.size.X()
}

// WrapR
func (t *BaseTexture) WrapR() int32 {
	return t.wrapR
}

// WrapS
func (t *BaseTexture) WrapS() int32 {
	return t.wrapS
}

// WrapT
func (t *BaseTexture) WrapT() int32 {
	return t.wrapT
}

func (t *BaseTexture) SetData([]uint8) {}

func (t *BaseTexture) SetLayerData([]uint8, int32) {}

func (t *BaseTexture) SetHDRData([]float32) {}

func (t *BaseTexture) SetHDRLayerData([]float32, int32) {}

func TextureFormatToInternal(format gfx.TextureFormat) int32 {
	switch format {
	case gfx.TextureFormatR8:
		return gl.R8
	case gfx.TextureFormatRG8:
		return gl.RG8
	case gfx.TextureFormatRGB8:
		return gl.RGB8
	case gfx.TextureFormatDefaultColor:
		fallthrough
	case gfx.TextureFormatRGBA8:
		return gl.RGBA8
	case gfx.TextureFormatR16:
		return gl.R16F
	case gfx.TextureFormatRG16:
		return gl.RG16F
	case gfx.TextureFormatRGB16:
		return gl.RGB16F
	case gfx.TextureFormatDefaultHDRColor:
		fallthrough
	case gfx.TextureFormatRGBA16:
		return gl.RGBA16F
	case gfx.TextureFormatR32:
		return gl.R32F
	case gfx.TextureFormatRG32:
		return gl.RG32F
	case gfx.TextureFormatRGB32:
		return gl.RGB32F
	case gfx.TextureFormatRGBA32:
		return gl.RGBA32F
	case gfx.TextureFormatRGB32UI:
		return gl.RGB32UI
	case gfx.TextureFormatRGBA32UI:
		return gl.RGBA32UI
	case gfx.TextureFormatDepth16:
		return gl.DEPTH_COMPONENT16
	case gfx.TextureFormatDefaultDepth:
		fallthrough
	case gfx.TextureFormatDepth24:
		return gl.DEPTH_COMPONENT24
	case gfx.TextureFormatDepth24Stencil8:
		return gl.DEPTH24_STENCIL8
	case gfx.TextureFormatStencil8:
		return gl.STENCIL_INDEX8
	case gfx.TextureFormatRGBA16UI:
		return gl.RGBA16UI
	}

	return 0
}

func TextureFormatToFormat(format gfx.TextureFormat) uint32 {
	switch format {
	case gfx.TextureFormatR8:
		fallthrough
	case gfx.TextureFormatR16:
		fallthrough
	case gfx.TextureFormatR32:
		return gl.RED
	case gfx.TextureFormatRG8:
		fallthrough
	case gfx.TextureFormatRG16:
		fallthrough
	case gfx.TextureFormatRG32:
		return gl.RG
	case gfx.TextureFormatRGB8:
		fallthrough
	case gfx.TextureFormatRGB16:
		fallthrough
	case gfx.TextureFormatRGB32:
		return gl.RGB
	case gfx.TextureFormatRGB32UI:
		return gl.RGB_INTEGER
	case gfx.TextureFormatDefaultColor:
		fallthrough
	case gfx.TextureFormatRGBA8:
		fallthrough
	case gfx.TextureFormatDefaultHDRColor:
		fallthrough
	case gfx.TextureFormatRGBA16:
		fallthrough
	case gfx.TextureFormatRGBA16UI:
		fallthrough
	case gfx.TextureFormatRGBA32:
		return gl.RGBA
	case gfx.TextureFormatRGBA32UI:
		return gl.RGBA_INTEGER
	case gfx.TextureFormatDefaultDepth:
		fallthrough
	case gfx.TextureFormatDepth16:
		fallthrough
	case gfx.TextureFormatDepth24:
		return gl.DEPTH_COMPONENT
	case gfx.TextureFormatDepth24Stencil8:
		fallthrough
	case gfx.TextureFormatStencil8:
		return 0
	}

	return 0
}

func TextureFormatToStorage(format gfx.TextureFormat) uint32 {
	switch format {
	case gfx.TextureFormatDefaultColor:
		fallthrough
	case gfx.TextureFormatR8:
		fallthrough
	case gfx.TextureFormatRG8:
		fallthrough
	case gfx.TextureFormatRGB8:
		fallthrough
	case gfx.TextureFormatRGBA8:
		fallthrough
	case gfx.TextureFormatStencil8:
		return gl.UNSIGNED_BYTE
	case gfx.TextureFormatR16:
		fallthrough
	case gfx.TextureFormatRG16:
		fallthrough
	case gfx.TextureFormatRGB16:
		fallthrough
	case gfx.TextureFormatDefaultHDRColor:
		fallthrough
	case gfx.TextureFormatRGBA16:
		return gl.HALF_FLOAT
	case gfx.TextureFormatRGBA16UI:
		return gl.UNSIGNED_SHORT
	case gfx.TextureFormatR32:
		fallthrough
	case gfx.TextureFormatRG32:
		fallthrough
	case gfx.TextureFormatRGB32:
		fallthrough
	case gfx.TextureFormatRGBA32:
		return gl.FLOAT
	case gfx.TextureFormatRGB32UI:
		fallthrough
	case gfx.TextureFormatRGBA32UI:
		return gl.UNSIGNED_INT
	case gfx.TextureFormatDefaultDepth:
		fallthrough
	case gfx.TextureFormatDepth16:
		fallthrough
	case gfx.TextureFormatDepth24:
		return gl.FLOAT
	case gfx.TextureFormatDepth24Stencil8:
		return gl.UNSIGNED_INT_24_8
	}

	return 0
}

func (r *Renderer) MakeTexture(cfg *gfx.TextureConfig) gfx.Texture {
	switch cfg.Type {
	case gfx.Texture2D:
		return NewTexture2D(cfg)
	case gfx.Texture3D:
		return NewTexture3D(cfg)
	case gfx.TextureCubemap:
		return NewTextureCubemap(cfg)
	case gfx.TextureFont:
		return NewTextureFont(cfg)
	case gfx.TextureColor:
		return NewTextureColor()
	default:
		return nil
	}
}
