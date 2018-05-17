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

package core

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"

	"github.com/haakenlabs/ember/gfx"
	"github.com/haakenlabs/ember/pkg/math"
	"github.com/sirupsen/logrus"
)

var _ System = &WindowSystem{}

var windowInst *WindowSystem

const SysNameWindow = "window"

type DisplayMode int
type MouseMode int

const (
	DisplayModeWindow DisplayMode = iota
	DisplayModeWindowedFullscreen
	DisplayModeFullscreen
)

type EventKey struct {
	key      glfw.Key
	scancode int
	action   glfw.Action
	mods     glfw.ModifierKey
}

type EventMouseButton struct {
	button glfw.MouseButton
	action glfw.Action
	mod    glfw.ModifierKey
}

type EventJoy struct {
	joystick int
	event    int
}

type DisplayProperties struct {
	Resolution math.IVec2
	Mode       DisplayMode
	Vsync      bool
}

// WindowSystem implements a GLFW-based window system.
type WindowSystem struct {
	window            *glfw.Window
	renderer          gfx.Renderer
	ortho             mgl32.Mat4
	resolution        math.IVec2
	mousePos          math.DVec2
	scrollAxis        math.DVec2
	cursorPosition    mgl32.Vec2
	mouseMode         MouseMode
	displayMode       DisplayMode
	mouseButtonEvents []EventMouseButton
	keyEvents         []EventKey
	joystickEvents    []EventJoy
	aspectRatio       float32
	title             string
	vsync             bool
	focus             bool
	cursorEnter       bool
	cursorMoved       bool
	scrollMoved       bool
	windowResized     bool
	shouldClose       bool
	hasEvents         bool
}

func (w *WindowSystem) Setup() (err error) {
	if windowInst != nil {
		return ErrSystemInit(SysNameWindow)
	}
	windowInst = w

	var monitor *glfw.Monitor

	if err := glfw.Init(); err != nil {
		return err
	}

	logrus.Debug("[GLFW] Library initialized")

	// Register input callbacks
	w.window.SetCharCallback(w.onChar)
	w.window.SetCursorEnterCallback(w.onCursorEnter)
	w.window.SetCursorPosCallback(w.onCursorMove)
	w.window.SetDropCallback(w.onDrop)
	w.window.SetKeyCallback(w.onKey)
	w.window.SetMouseButtonCallback(w.onMouseButton)
	w.window.SetScrollCallback(w.onScroll)
	w.window.SetCloseCallback(w.onClose)
	w.window.SetSizeCallback(w.onWindowResize)
	glfw.SetJoystickCallback(w.onJoystick)

	logrus.Debug("[GLFW] Ready")

}
