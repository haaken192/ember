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
	"fmt"

	"github.com/go-gl/gl/v4.3-core/gl"
	"github.com/sirupsen/logrus"

	"github.com/haakenlabs/ember/core"
	"github.com/haakenlabs/ember/gfx"
	"github.com/haakenlabs/ember/pkg/math"
	"github.com/haakenlabs/ember/system/instance"
)

var (
	framebufferStack []*Framebuffer
)

var _ gfx.Framebuffer = &Framebuffer{}

type Framebuffer struct {
	core.BaseObject

	size        math.IVec2
	bound       bool
	attachments map[uint32]gfx.Attachment
	drawBuffers []uint32
	reference   uint32
}

func NewFramebuffer(size math.IVec2) *Framebuffer {
	f := &Framebuffer{
		size:        size,
		attachments: make(map[uint32]gfx.Attachment),
		drawBuffers: []uint32{},
	}

	f.SetName("Framebuffer")
	instance.MustAssign(f)

	gl.GenFramebuffers(1, &f.reference)

	return f
}

func popFramebuffer() {
	if len(framebufferStack) != 0 {
		framebufferStack[len(framebufferStack)-1].RawUnbind()
		framebufferStack[len(framebufferStack)-1].bound = false
		framebufferStack = framebufferStack[:len(framebufferStack)-1]

		if len(framebufferStack) != 0 {
			framebufferStack[len(framebufferStack)-1].RawBind()
			framebufferStack[len(framebufferStack)-1].bound = true
		} else {
			gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
			gl.Viewport(
				0, 0,
				core.GetWindowSystem().Resolution().X(),
				core.GetWindowSystem().Resolution().Y())
		}
	}
}

func pushFramebuffer(framebuffer *Framebuffer) {
	if len(framebufferStack) != 0 {
		framebufferStack[len(framebufferStack)-1].RawUnbind()
		framebufferStack[len(framebufferStack)-1].bound = false
	}

	framebufferStack = append(framebufferStack, framebuffer)
	framebufferStack[len(framebufferStack)-1].RawBind()
	framebufferStack[len(framebufferStack)-1].bound = true
}

func BindCurrentFramebuffer() {
	if current := CurrentFramebuffer(); current != nil {
		current.RawBind()
	} else {
		gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	}
}

func UnbindCurrentFramebuffer() {
	popFramebuffer()
}

func CurrentFramebuffer() *Framebuffer {
	if len(framebufferStack) == 0 {
		return nil
	}

	return framebufferStack[len(framebufferStack)-1]
}

func BlitFramebuffers(in *Framebuffer, out *Framebuffer, location uint32) {
	src := in.Reference()
	dst := uint32(0)

	srcSize := in.Size()
	dstSize := core.GetWindowSystem().Resolution()

	if out != nil {
		dst = out.Reference()
		dstSize = out.Size()
	}

	gl.BindFramebuffer(gl.READ_FRAMEBUFFER, src)
	gl.BindFramebuffer(gl.DRAW_FRAMEBUFFER, dst)
	gl.ReadBuffer(location)
	gl.BlitFramebuffer(0, 0, srcSize.X(), srcSize.Y(), 0, 0, dstSize.X(), dstSize.Y(), gl.COLOR_BUFFER_BIT, gl.LINEAR)

	if err := gl.GetError(); err != gl.NO_ERROR {
		panic(err)
	}

	BindCurrentFramebuffer()
}

func (f *Framebuffer) Dealloc() {
	if f.reference != 0 {
		gl.DeleteFramebuffers(1, &f.reference)
		f.reference = 0
	}
}

func (f *Framebuffer) Alloc() error {
	f.RawBind()

	for idx := range f.attachments {
		f.attachments[idx].SetSize(f.size)
		f.attachments[idx].Attach(idx)
	}

	if len(f.drawBuffers) != 0 {
		gl.DrawBuffers(int32(len(f.drawBuffers)), &f.drawBuffers[0])
	}

	if err := f.Validate(); err != nil {
		return err
	}

	f.RawUnbind()

	return nil
}

func (f *Framebuffer) Bind() {
	pushFramebuffer(f)
}

func (f *Framebuffer) Unbind() {
	if f.bound {
		popFramebuffer()
	} else {
		f.RawUnbind()
		gl.Viewport(
			0, 0,
			core.GetWindowSystem().Resolution().X(),
			core.GetWindowSystem().Resolution().Y())
	}
}

func (f *Framebuffer) RawBind() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, f.reference)
	gl.Viewport(0, 0, f.size.X(), f.size.Y())
}

func (f *Framebuffer) Validate() error {
	if f.size.X() <= 0 || f.size.Y() <= 0 {
		return fmt.Errorf("validate: framebuffer %d has invalid size: %s", f.reference, f.size)
	}

	status := gl.CheckFramebufferStatus(gl.FRAMEBUFFER)

	if status != gl.FRAMEBUFFER_COMPLETE {
		switch status {
		case gl.FRAMEBUFFER_UNSUPPORTED:
			return fmt.Errorf("validate: framebuffer %d: unsupported framebuffer format", f.reference)
		case gl.FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT:
			return fmt.Errorf("validate: framebuffer %d: missing attachment", f.reference)
		case gl.FRAMEBUFFER_INCOMPLETE_ATTACHMENT:
			return fmt.Errorf("validate: framebuffer %d: incomplete attachment", f.reference)
		case gl.FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER:
			return fmt.Errorf("validate: framebuffer %d: missing draw buffer", f.reference)
		case gl.FRAMEBUFFER_INCOMPLETE_READ_BUFFER:
			return fmt.Errorf("validate: framebuffer %d: missing read buffer", f.reference)
		default:
			return fmt.Errorf("validate: framebuffer %d: unknown framebuffer error: %d", f.reference, status)
		}
	}

	logrus.Debugf("validate: framebuffer %d complete", f.reference)

	return nil
}

func (f *Framebuffer) RawUnbind() {
	BindCurrentFramebuffer()
}

func (f *Framebuffer) SetSize(size math.IVec2) {
	if size.X() <= 0 || size.Y() <= 0 {
		return
	}

	f.size = size
	f.Alloc()
}

func (f *Framebuffer) SetAttachment(location uint32, attachment gfx.Attachment) {
	f.attachments[location] = attachment
}

func (f *Framebuffer) SetDrawBuffers(buffers []uint32) {
	f.drawBuffers = buffers
}

func (f *Framebuffer) ApplyDrawBuffers(buffers []uint32) {
	f.SetDrawBuffers(buffers)

	if len(f.drawBuffers) != 0 {
		gl.DrawBuffers(int32(len(f.drawBuffers)), &f.drawBuffers[0])
	} else {
		gl.DrawBuffers(1, nil)
	}
}

func (f *Framebuffer) RemoveAttachment(location uint32) {
	if f.HasAttachment(location) {
		delete(f.attachments, location)
	}
}

func (f *Framebuffer) RemoveAllAttachments() {
	for idx := range f.attachments {
		delete(f.attachments, idx)
	}
}

func (f *Framebuffer) GetAttachment(location uint32) gfx.Attachment {
	a, ok := f.attachments[location]

	if !ok {
		return nil
	}

	return a
}

func (f *Framebuffer) HasAttachment(location uint32) bool {
	_, ok := f.attachments[location]

	return ok
}

func (f *Framebuffer) Size() math.IVec2 {
	return f.size
}

func (f *Framebuffer) Reference() uint32 {
	return f.reference
}

func (f *Framebuffer) Bound() bool {
	return f.bound
}

func (f *Framebuffer) Clear() {
	f.ClearBufferFlags(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT | gl.STENCIL_BUFFER_BIT)
}

func (f *Framebuffer) ClearFlags(flags uint32) {
	var glFlags uint32

	if flags&gfx.ClearColorBuffer != 0 {
		glFlags |= gl.COLOR_BUFFER_BIT
	}
	if flags&gfx.ClearDepthBuffer != 0 {
		glFlags |= gl.DEPTH_BUFFER_BIT
	}
	if flags&gfx.ClearStencilBuffer != 0 {
		glFlags |= gl.STENCIL_BUFFER_BIT
	}

	f.ClearBufferFlags(glFlags)
}

func (f *Framebuffer) ClearBufferFlags(flags uint32) {
	gl.Clear(flags)
}

func (f *Framebuffer) SetResizable(bool) {

}

func (f *Framebuffer) Resize() {}

func (f *Framebuffer) Resizable() bool {
	return true
}

func (r *Renderer) MakeFramebuffer(size math.IVec2) gfx.Framebuffer {
	return NewFramebuffer(size)
}
