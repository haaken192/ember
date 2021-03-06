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
	"github.com/haakenlabs/ember/gfx"
	"github.com/haakenlabs/ember/pkg/math"
)

var _ gfx.Attachment = &AttachmentTexture2D{}
var _ gfx.Attachment = &AttachmentRenderbuffer{}

type AttachmentTexture2D struct {
}

type AttachmentRenderbuffer struct {
}

func (a *AttachmentTexture2D) Attach(uint32) {}

func (a *AttachmentTexture2D) SetSize(math.IVec2) {}

func (a *AttachmentTexture2D) Type() gfx.AttachmentType {
	return gfx.AttachmentTexture2D
}

func (a *AttachmentRenderbuffer) Attach(uint32) {}

func (a *AttachmentRenderbuffer) SetSize(math.IVec2) {}

func (a *AttachmentRenderbuffer) Type() gfx.AttachmentType {
	return gfx.AttachmentRenderbuffer
}

func (r *Renderer) MakeAttachment(cfg *gfx.AttachmentConfig) gfx.Attachment {
	switch cfg.Type {
	case gfx.AttachmentTexture2D:
		return nil
	case gfx.AttachmentRenderbuffer:
		return nil
	default:
		return nil
	}
}
