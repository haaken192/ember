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

	"github.com/haakenlabs/ember/gfx"
	"github.com/haakenlabs/ember/pkg/math"
)

var _ gfx.Attachment = &AttachmentRenderbuffer{}
var _ gfx.Attachment = &AttachmentTexture2D{}

type AttachmentRenderbuffer struct {
	attachment *RenderBuffer
}

type AttachmentTexture2D struct {
	attachment *Texture2D
	mipLevel   int32
}

func NewAttachmentRenderBuffer(cfg *gfx.TextureConfig) *AttachmentRenderbuffer {
	rbuffer := NewRenderBuffer(cfg.Size, cfg.Format)

	return NewAttachmentRenderBufferFrom(rbuffer)
}

func NewAttachmentRenderBufferFrom(buffer *RenderBuffer) *AttachmentRenderbuffer {
	a := &AttachmentRenderbuffer{
		attachment: buffer,
	}

	return a
}

func (a *AttachmentRenderbuffer) Attach(location uint32) {
	gl.FramebufferRenderbuffer(gl.FRAMEBUFFER, location, gl.RENDERBUFFER, a.attachment.Reference())
}

func (a *AttachmentRenderbuffer) SetSize(size math.IVec2) {
	a.attachment.SetSize(size)
}

func (a *AttachmentRenderbuffer) AttachmentObject() *RenderBuffer {
	return a.attachment
}

func (a *AttachmentRenderbuffer) Type() gfx.AttachmentType {
	return gfx.AttachmentRenderbuffer
}

func NewAttachmentTexture2D(cfg *gfx.TextureConfig) *AttachmentTexture2D {
	t := NewTexture2D(cfg)
	t.Alloc()

	a := &AttachmentTexture2D{
		attachment: t,
	}

	return a
}

func NewAttachmentTexture2DFrom(texture gfx.Texture) *AttachmentTexture2D {
	t, ok := texture.(*Texture2D)
	if !ok {
		return nil
	}

	a := &AttachmentTexture2D{
		attachment: t,
	}

	return a
}

func (a *AttachmentTexture2D) Attach(location uint32) {
	gl.FramebufferTexture2D(gl.FRAMEBUFFER, location, gl.TEXTURE_2D, a.attachment.Reference(), a.mipLevel)
}

func (a *AttachmentTexture2D) SetSize(size math.IVec2) {
	a.attachment.SetSize(size)
}

func (a *AttachmentTexture2D) MipLevel() int32 {
	return a.mipLevel
}

func (a *AttachmentTexture2D) SetMipLevel(mipLevel int32) {
	a.mipLevel = mipLevel
}

func (a *AttachmentTexture2D) AttachmentObject() *Texture2D {
	return a.attachment
}

func (a *AttachmentTexture2D) Type() gfx.AttachmentType {
	return gfx.AttachmentTexture2D
}

func (r *Renderer) MakeAttachment(cfg *gfx.AttachmentConfig) gfx.Attachment {
	switch cfg.Type {
	case gfx.AttachmentTexture2D:
		if cfg.Tex != nil {
			return NewAttachmentTexture2DFrom(cfg.Tex)
		} else if cfg.TexCfg != nil {
			return NewAttachmentTexture2D(cfg.TexCfg)
		}
	case gfx.AttachmentRenderbuffer:
		if cfg.TexCfg != nil {
			return NewAttachmentRenderBuffer(cfg.TexCfg)
		}
	}

	return nil
}
