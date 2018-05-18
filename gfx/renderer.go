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

import (
	"github.com/go-gl/glfw/v3.2/glfw"

	"github.com/haakenlabs/ember/pkg/math"
)

// Binder is an object which can bind or unbind itself in the renderer.
type Binder interface {
	// Bind will make itself the active object for that type.
	Bind()

	// Unbind will release itself as the active object for that type.
	Unbind()

	// Reference gets the reference ID of the object.
	Reference() uint32
}

// Allocater is an object which can allocate or deallocate resources for itself.
type Allocater interface {
	// Alloc allocates resources for the object.
	Alloc() error

	// Dealloc releases any previously allocated resources for the object.
	Dealloc()

	// ID returns the instance ID of this object.
	ID() int32
}

type Sizable interface {
	Size() math.IVec2
	SetSize(vec2 math.IVec2)
	SetResizable(bool)
	Resize()
	Resizable() bool
}

// Drawer is an object which can perform drawing operations.
type Drawer interface {
	// Draw the object.
	Draw()
}

// Renderer implements a graphics rendering plane. It translates generic calls
// to actual rendering commands used by a rendering API, such as OpenGL
// or Vulkan.
type Renderer interface {
	Factory

	// Bind an object.
	Bind(Binder)

	// Unbind an object.
	Unbind(Binder)

	// Draw an object
	Draw(Drawer)

	// Begin drawing operations.
	Begin()

	// End drawing operations.
	End()

	// Alloc allocates an object.
	Alloc(Allocater)

	// Dealloc allocates an object.
	Dealloc(Allocater)

	// Init initializes the renderer.
	Init(*glfw.Window) error

	// Destroy tears down the renderer.
	Destroy()
}

type Factory interface {
	MakeShader(bool) Shader
	MakeTexture(*TextureConfig) Texture
	MakeAttachment(*AttachmentConfig) Attachment
	MakeFramebuffer(math.IVec2) Framebuffer
	MakeGBuffer(math.IVec2, Attachment, bool) GBuffer
	MakeMesh() Mesh
	MakeCamera() Camera
}
