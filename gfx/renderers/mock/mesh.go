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
	"github.com/go-gl/mathgl/mgl32"
	"github.com/haakenlabs/ember/gfx"
	"github.com/sirupsen/logrus"
)

var _ gfx.Mesh = &Mesh{}

type Mesh struct{}

func (m *Mesh) Bind() {}

func (m *Mesh) Unbind() {}

func (m *Mesh) Reference() uint32 {
	return 1
}

func (m *Mesh) Alloc() error {
	logrus.Info("alloc mock mesh")
	return nil
}

func (m *Mesh) Dealloc() {}

func (m *Mesh) ID() int32 {}

func (m *Mesh) Deferred() bool {
	return true
}

func (m *Mesh) Draw() {}

func (m *Mesh) Clear() {}

func (m *Mesh) Upload() error {
	return nil
}

func (m *Mesh) Vertices() []mgl32.Vec3 {
	return nil
}

func (m *Mesh) Normals() []mgl32.Vec3 {
	return nil
}

func (m *Mesh) UVs() []mgl32.Vec2 {
	return nil
}

func (m *Mesh) Triangles() []uint32 {
	return nil
}

func (m *Mesh) Indexed() bool {
	return false
}
func (m *Mesh) ReversedWinding() bool {
	return false
}

func (m *Mesh) SetVertices(vertices []mgl32.Vec3) {}

func (m *Mesh) SetNormals(normals []mgl32.Vec3) {}

func (m *Mesh) SetUVs(uvs []mgl32.Vec2) {}

func (m *Mesh) SetReversedWinding(reverse bool) {}
