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
	"bytes"
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sirupsen/logrus"

	"github.com/haakenlabs/ember/gfx"
)

var _ gfx.Shader = &Shader{}

type Shader struct {
	reference  uint32
	components map[gfx.ShaderComponent]uint32
	data       []byte
	deferred   bool
}

// Alloc allocates resources for the shader.
func (s *Shader) Alloc() error {
	return s.Compile()
}

// Dealloc releases any previously allocated resources for the shader.
func (s *Shader) Dealloc() {
	if s.reference != 0 {
		for k := range s.components {
			destroyComponent(s.components[k], s.reference)
			delete(s.components, k)
		}

		gl.DeleteProgram(s.reference)

		s.reference = 0
	}
}

// Reference gets the reference ID of the shader.
func (s *Shader) Reference() uint32 {
	return s.reference
}

// Bind will activate this shader.
func (s *Shader) Bind() {
	gl.UseProgram(s.reference)
}

// Bind will deactivate this shader.
func (s *Shader) Unbind() {
	gl.UseProgram(0)
}

// Deferred is true when this shader can be used for deferred rendering.
func (s *Shader) Deferred() bool {
	return s.deferred
}

func (s *Shader) Compile() error {
	s.reference = gl.CreateProgram()

	if containsShaderType(gfx.ShaderComponentVertex, s.data) {
		componentId, err := loadComponent(s.reference, gfx.ShaderComponentVertex, s.data)
		if err != nil {
			return err
		}
		s.components[gfx.ShaderComponentVertex] = componentId
	}
	if containsShaderType(gfx.ShaderComponentGeometry, s.data) {
		componentId, err := loadComponent(s.reference, gfx.ShaderComponentGeometry, s.data)
		if err != nil {
			return err
		}
		s.components[gfx.ShaderComponentGeometry] = componentId
	}
	if containsShaderType(gfx.ShaderComponentFragment, s.data) {
		componentId, err := loadComponent(s.reference, gfx.ShaderComponentFragment, s.data)
		if err != nil {
			return err
		}
		s.components[gfx.ShaderComponentFragment] = componentId
	}
	if containsShaderType(gfx.ShaderComponentCompute, s.data) {
		componentId, err := loadComponent(s.reference, gfx.ShaderComponentCompute, s.data)
		if err != nil {
			return err
		}
		s.components[gfx.ShaderComponentCompute] = componentId
	}
	if containsShaderType(gfx.ShaderComponentTessControl, s.data) {
		componentId, err := loadComponent(s.reference, gfx.ShaderComponentTessControl, s.data)
		if err != nil {
			return err
		}
		s.components[gfx.ShaderComponentTessControl] = componentId
	}
	if containsShaderType(gfx.ShaderComponentTessEvaluation, s.data) {
		componentId, err := loadComponent(s.reference, gfx.ShaderComponentTessEvaluation, s.data)
		if err != nil {
			return err
		}
		s.components[gfx.ShaderComponentTessEvaluation] = componentId
	}

	// Set transform feedback varyings
	// TODO: Implement this

	// Validate and link
	return link(s.reference)
}

func (s *Shader) SetSubroutine(componentType gfx.ShaderComponent, subroutineName string) {
	idx := gl.GetSubroutineIndex(s.reference, uint32(componentType), gl.Str(subroutineName+"\x00"))
	gl.UniformSubroutinesuiv(uint32(componentType), 1, &idx)
}

func (s *Shader) SetUniform(uniformName string, value interface{}) {
	switch v := value.(type) {
	case bool:
		var val int32
		if v {
			val = 1
		}
		gl.Uniform1i(gl.GetUniformLocation(s.reference, gl.Str(uniformName+"\x00")), val)
	case int32:
		gl.Uniform1i(gl.GetUniformLocation(s.reference, gl.Str(uniformName+"\x00")), v)
	case float32:
		gl.Uniform1f(gl.GetUniformLocation(s.reference, gl.Str(uniformName+"\x00")), v)
	case uint32:
		gl.Uniform1ui(gl.GetUniformLocation(s.reference, gl.Str(uniformName+"\x00")), v)
	case mgl32.Vec2:
		gl.Uniform2fv(gl.GetUniformLocation(s.reference, gl.Str(uniformName+"\x00")), 1, &v[0])
	case mgl32.Vec3:
		gl.Uniform3fv(gl.GetUniformLocation(s.reference, gl.Str(uniformName+"\x00")), 1, &v[0])
	case mgl32.Vec4:
		gl.Uniform4fv(gl.GetUniformLocation(s.reference, gl.Str(uniformName+"\x00")), 1, &v[0])
	case mgl32.Mat2:
		gl.UniformMatrix2fv(gl.GetUniformLocation(s.reference, gl.Str(uniformName+"\x00")), 1, false, &v[0])
	case mgl32.Mat3:
		gl.UniformMatrix3fv(gl.GetUniformLocation(s.reference, gl.Str(uniformName+"\x00")), 1, false, &v[0])
	case mgl32.Mat4:
		gl.UniformMatrix4fv(gl.GetUniformLocation(s.reference, gl.Str(uniformName+"\x00")), 1, false, &v[0])
	}
}

func link(program uint32) error {
	gl.LinkProgram(program)

	return validateProgram(program)
}

func validateComponent(component uint32) error {
	var status int32
	gl.GetShaderiv(component, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(component, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(component, logLength, nil, gl.Str(log))

		return fmt.Errorf("shader %d compilation failed: %v", component, log)
	}

	return nil
}

func validateProgram(program uint32) error {
	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return fmt.Errorf("program %d link failed: %v", program, log)
	}

	return nil
}

func destroyComponent(component uint32, program uint32) {
	gl.DetachShader(program, component)
	gl.DeleteShader(component)
}

func containsShaderType(shaderType gfx.ShaderComponent, data []byte) bool {
	switch shaderType {
	case gfx.ShaderComponentVertex:
		return bytes.Contains(data, []byte("#ifdef _VERTEX_"))
	case gfx.ShaderComponentGeometry:
		return bytes.Contains(data, []byte("#ifdef _GEOMETRY_"))
	case gfx.ShaderComponentFragment:
		return bytes.Contains(data, []byte("#ifdef _FRAGMENT_"))
	case gfx.ShaderComponentCompute:
		return bytes.Contains(data, []byte("#ifdef _COMPUTE_"))
	case gfx.ShaderComponentTessControl:
		return bytes.Contains(data, []byte("#ifdef _TESSCONTROL_"))
	case gfx.ShaderComponentTessEvaluation:
		return bytes.Contains(data, []byte("#ifdef _TESSEVAL_"))
	}

	return false
}

func loadComponent(program uint32, componentType gfx.ShaderComponent, data []byte) (uint32, error) {
	header := []byte("#version 430\n")

	switch componentType {
	case gfx.ShaderComponentVertex:
		header = append(header, []byte("#define _VERTEX_\n")...)
	case gfx.ShaderComponentGeometry:
		header = append(header, []byte("#define _GEOMETRY_\n")...)
	case gfx.ShaderComponentFragment:
		header = append(header, []byte("#define _FRAGMENT_\n")...)
	case gfx.ShaderComponentCompute:
		header = append(header, []byte("#define _COMPUTE_\n")...)
	case gfx.ShaderComponentTessControl:
		header = append(header, []byte("#define _TESSCONTROL_\n")...)
	case gfx.ShaderComponentTessEvaluation:
		header = append(header, []byte("#define _TESSEVAL_\n")...)
	default:
		return 0, fmt.Errorf("loadComponent failed: unknown component type: %d", componentType)
	}

	data = append(header, data...)

	component := gl.CreateShader(uint32(componentType))

	csrc, free := gl.Strs(string(data))
	srcLength := int32(len(data))
	gl.ShaderSource(component, 1, csrc, &srcLength)
	free()
	gl.CompileShader(component)

	err := validateComponent(component)
	if err != nil {
		fmt.Println(string(data))
		return 0, err
	}

	gl.AttachShader(program, component)

	logrus.Debugf("Loaded component(%s) %d for program %d", ShaderComponentToString(componentType), component, program)

	return component, nil
}

// ShaderComponentToString returns the string representation of a core.ShaderComponent.
func ShaderComponentToString(component gfx.ShaderComponent) string {
	switch component {
	case gfx.ShaderComponentVertex:
		return "VERTEX"
	case gfx.ShaderComponentGeometry:
		return "GEOMETRY"
	case gfx.ShaderComponentFragment:
		return "FRAGMENT"
	case gfx.ShaderComponentCompute:
		return "COMPUTE"
	case gfx.ShaderComponentTessControl:
		return "TESSCONTROL"
	case gfx.ShaderComponentTessEvaluation:
		return "TESSEVAL"
	}

	return "INVALID"
}

func (r *Renderer) MakeShader(deferred bool) gfx.Shader {
	s := &Shader{
		components: make(map[gfx.ShaderComponent]uint32),
		deferred:   deferred,
	}

	return s
}
