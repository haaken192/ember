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

package gfx

import "github.com/haakenlabs/ember/pkg/math"

type TextureType uint8

const (
	Texture2D TextureType = iota
	Texture3D
	TextureCubemap
	TextureFont
	TextureColor
)

func (t TextureType) String() string {
	switch t {
	case Texture2D:
		return "Texture2D"
	case Texture3D:
		return "Texture3D"
	case TextureCubemap:
		return "TextureCubemap"
	case TextureFont:
		return "TextureFont"
	case TextureColor:
		return "TextureColor"
	default:
		return "Unknown Texture Type"
	}
}

type TextureFormat uint32

const (
	TextureFormatDefaultColor TextureFormat = iota
	TextureFormatDefaultHDRColor
	TextureFormatDefaultDepth
	TextureFormatR8
	TextureFormatRG8
	TextureFormatRGB8
	TextureFormatRGBA8
	TextureFormatR16
	TextureFormatRG16
	TextureFormatRGB16
	TextureFormatRGBA16
	TextureFormatRGBA16UI
	TextureFormatR32
	TextureFormatRG32
	TextureFormatRGB32
	TextureFormatRGBA32
	TextureFormatRGB32UI
	TextureFormatRGBA32UI
	TextureFormatDepth16
	TextureFormatDepth24
	TextureFormatDepth24Stencil8
	TextureFormatStencil8
)

type Texture interface {
	Allocater
	Binder
	Sizable

	Type() TextureType
	SetMagFilter(int32)
	SetMinFilter(int32)
	FilterMag() int32
	FilterMin() int32
	WrapR() int32
	WrapS() int32
	WrapT() int32
	SetWrapR(int32)
	SetWrapRST(int32, int32, int32)
	SetWrapS(int32)
	SetWrapST(int32, int32)
	SetWrapT(int32)
	Activate(uint32)
	MipLevels() uint32
	GenerateMipmaps()
	Layers() int32
	SetLayers(int32)
}

type TextureConfig struct {
	Type   TextureType
	Format TextureFormat
	Layers int32
	Size   math.IVec2
}
