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

type AttachmentType uint32

const (
	AttachmentTexture2D AttachmentType = iota
	AttachmentRenderbuffer
)

type ClearBufferFlags uint32

const (
	ClearColorBuffer ClearBufferFlags = iota
	ClearDepthBuffer
	ClearStencilBuffer
)

type AttachmentLocation uint32

const (
	AttachmentColor0 AttachmentLocation = iota
	AttachmentColor1
	AttachmentColor2
	AttachmentColor3
	AttachmentDepth
	AttachmentStencil
)

type Attachment interface {
	Attach(uint32)
	SetSize(math.IVec2)
	Type() AttachmentType
}

type Framebuffer interface {
	Allocater
	Binder

	Resize()
	Validate() error
	Clear()
	ClearFlags(uint32)
}

type GBuffer interface {
	Framebuffer

	HDR() bool
	SetHDR(bool)
}
