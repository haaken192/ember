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
	"github.com/go-gl/mathgl/mgl32"

	"github.com/haakenlabs/ember/core"
	"github.com/haakenlabs/ember/gfx"
	"github.com/haakenlabs/ember/system/instance"
)

var _ gfx.Mesh = &Mesh{}

type Mesh struct {
	core.BaseObject

	vertices       []mgl32.Vec3
	normals        []mgl32.Vec3
	uvs            []mgl32.Vec2
	triangles      []uint32
	vao            uint32
	vbo            uint32
	ibo            uint32
	reverseWinding bool
}

func (m *Mesh) Bind() {
	gl.BindVertexArray(m.vao)
}

func (m *Mesh) Unbind() {
	gl.BindVertexArray(0)
}

func (m *Mesh) Reference() uint32 {
	return 1
}

func (m *Mesh) Alloc() error {
	gl.GenVertexArrays(1, &m.vao)
	gl.BindVertexArray(m.vao)

	gl.GenBuffers(1, &m.vbo)
	gl.GenBuffers(1, &m.ibo)
	gl.BindBuffer(gl.ARRAY_BUFFER, m.vbo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, m.ibo)

	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 32, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 32, gl.PtrOffset(12))
	gl.EnableVertexAttribArray(2)
	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 32, gl.PtrOffset(24))

	return m.Upload()
}

func (m *Mesh) Dealloc() {
	gl.DeleteBuffers(1, &m.vbo)
	gl.DeleteBuffers(1, &m.ibo)
	gl.DeleteVertexArrays(1, &m.vao)
}

func (m *Mesh) Deferred() bool {
	return true
}

func (m *Mesh) Draw() {
	if len(m.vertices) == 0 {
		return
	}

	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(m.vertices)))
}

func (m *Mesh) Clear() {
	m.vertices = m.vertices[:0]
	m.normals = m.normals[:0]
	m.uvs = m.uvs[:0]
	m.triangles = m.triangles[:0]
}

func (m *Mesh) Upload() error {
	if len(m.vertices) == 0 || len(m.normals) == 0 || len(m.uvs) == 0 {
		return fmt.Errorf("mesh upload failed: vao %d has invalid geometry definition: empty data", m.vao)
	}

	if len(m.vertices) != len(m.normals) || len(m.normals) != len(m.uvs) {
		return fmt.Errorf("mesh upload failed: vao %d has invalid geometry definition: asymmetric data", m.vao)
	}

	data := make([]gfx.Vertex, len(m.vertices))
	for idx := range m.vertices {
		data[idx] = gfx.Vertex{
			V: m.vertices[idx],
			N: m.normals[idx],
			U: m.uvs[idx],
		}
	}

	m.Bind()
	gl.BindBuffer(gl.ARRAY_BUFFER, m.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(data)*32, gl.Ptr(data), gl.STATIC_DRAW)
	m.Unbind()

	return nil
}

func (m *Mesh) Vertices() []mgl32.Vec3 {
	return m.vertices
}

func (m *Mesh) Normals() []mgl32.Vec3 {
	return m.normals
}

func (m *Mesh) UVs() []mgl32.Vec2 {
	return m.uvs
}

func (m *Mesh) Triangles() []uint32 {
	return m.triangles
}

func (m *Mesh) Indexed() bool {
	return len(m.triangles) != 0
}
func (m *Mesh) ReversedWinding() bool {
	return m.reverseWinding
}

func (m *Mesh) SetVertices(vertices []mgl32.Vec3) {
	m.vertices = vertices
}

func (m *Mesh) SetNormals(normals []mgl32.Vec3) {
	m.normals = normals
}

func (m *Mesh) SetUVs(uvs []mgl32.Vec2) {
	m.uvs = uvs
}

func (m *Mesh) SetReversedWinding(reverse bool) {
	m.reverseWinding = reverse
}

func NewMesh() *Mesh {
	m := &Mesh{}

	m.SetName("Mesh")
	instance.MustAssign(m)

	return m
}

func (r *Renderer) MakeMesh() gfx.Mesh {
	return NewMesh()
}
