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

import "github.com/haakenlabs/ember/gfx"
import (
	"github.com/go-gl/gl/v4.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/sirupsen/logrus"
)

var _ gfx.Renderer = &Renderer{}

type Renderer struct{}

func (r *Renderer) Bind(gfx.Binder) {

}

func (r *Renderer) Unbind(gfx.Binder) {

}

func (r *Renderer) Draw(gfx.Drawer) {

}

func (r *Renderer) Begin() {

}

func (r *Renderer) End() {

}

func (r *Renderer) Alloc(gfx.Allocater) {

}

func (r *Renderer) Dealloc(gfx.Allocater) {

}

func (r *Renderer) Init(window *glfw.Window) error {
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		return err
	}

	logrus.Debug("[OpenGL] Version: ", gl.GoStr(gl.GetString(gl.VERSION)))

	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.TEXTURE_CUBE_MAP_SEAMLESS)
	gl.DepthFunc(gl.LEQUAL)
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)

	logrus.Debug("[OpenGL] Ready")

	return nil
}

func (r *Renderer) Destroy() {

}

func NewRenderer() (*Renderer, error) {
	r := &Renderer{}

	return r, nil
}
