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

var _ gfx.Framebuffer = &Framebuffer{}
var _ gfx.GBuffer = &GBuffer{}

type Framebuffer struct {
	size math.IVec2
}

type GBuffer struct {
	Framebuffer

	hdr bool
}

func (f *Framebuffer) Bind() {}

func (f *Framebuffer) Unbind() {}

func (f *Framebuffer) Reference() uint32 {
	return 1
}

func (f *Framebuffer) Alloc() error {
	return nil
}

func (f *Framebuffer) Dealloc() {}

func (f *Framebuffer) Validate() error {
	return nil
}

func (f *Framebuffer) Clear() {}

func (f *Framebuffer) ClearFlags(uint32) {}

func (f *Framebuffer) ID() int32 {
	return 1
}

func (f *Framebuffer) Size() math.IVec2 {
	return f.size
}

func (f *Framebuffer) SetSize(size math.IVec2) {
	f.size = size
}

func (f *Framebuffer) SetResizable(bool) {}

func (f *Framebuffer) Resize() {}

func (f *Framebuffer) Resizable() bool {
	return true
}

func (g *GBuffer) HDR() bool {
	return g.hdr
}

func (g *GBuffer) SetHDR(value bool) {
	g.hdr = value
}

func (r *Renderer) MakeFramebuffer(size math.IVec2) gfx.Framebuffer {
	return &Framebuffer{
		size: size,
	}
}

func (r *Renderer) MakeGBuffer(size math.IVec2, depth gfx.Attachment, hdr bool) gfx.GBuffer {
	return &GBuffer{
		hdr: hdr,
	}
}
