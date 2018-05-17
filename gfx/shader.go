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
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/juju/errors"

	"github.com/haakenlabs/ember/pkg/cast"
)

type ShaderComponent uint32

const (
	ShaderComponentVertex ShaderComponent = iota
	ShaderComponentGeometry
	ShaderComponentFragment
	ShaderComponentCompute
	ShaderComponentTessControl
	ShaderComponentTessEvaluation
)

const (
	PropertyTypeBool  = "bool"
	PropertyTypeInt   = "int"
	PropertyTypeUint  = "uint"
	PropertyTypeFloat = "float"
	PropertyTypeVec2  = "vec2"
	PropertyTypeVec3  = "vec3"
	PropertyTypeVec4  = "vec4"
	PropertyTypeMat2  = "mat2"
	PropertyTypeMat3  = "mat3"
	PropertyTypeMat4  = "mat4"
)

type Shader interface {
	Allocater
	Binder

	// Deferred is true if this shader can be used for deferred rendering.
	Deferred() bool

	// AddData adds shader code to this shader.
	AddData([]byte)

	// ResetData clears the shader code data for this shader.
	ResetData()

	// Compile compiles and links the shader code data for this shader.
	Compile() error
}

type ShaderProperty struct {
	Value interface{}
}

func (s *ShaderProperty) UnmarshalJSON(data []byte) error {
	var v interface{}
	var err error

	raw := strings.Trim(string(data), "\"")
	parts := strings.SplitN(raw, ":", 2)
	if len(parts) != 2 {
		return errors.New("parse error, invalid format")
	}

	switch parts[0] {
	case PropertyTypeBool:
		v, err = strconv.ParseBool(parts[1])
	case PropertyTypeFloat:
		var f float64
		f, err = strconv.ParseFloat(parts[1], 32)
		v = float32(f)
	case PropertyTypeInt:
		var i int64
		i, err = strconv.ParseInt(parts[1], 10, 32)
		v = int32(i)
	case PropertyTypeUint:
		var i uint64
		i, err = strconv.ParseUint(parts[1], 10, 32)
		v = uint32(i)
	case PropertyTypeVec2:
		v, err = cast.ParseVec2(parts[1])
	case PropertyTypeVec3:
		v, err = cast.ParseVec3(parts[1])
	case PropertyTypeVec4:
		v, err = cast.ParseVec4(parts[1])
	case PropertyTypeMat2:
		v, err = cast.ParseMat2(parts[1])
	case PropertyTypeMat3:
		v, err = cast.ParseMat3(parts[1])
	case PropertyTypeMat4:
		v, err = cast.ParseMat4(parts[1])
	default:
		return fmt.Errorf("unknown data type: %s", parts[0])
	}

	if err != nil {
		return err
	}

	s.Value = v

	return nil
}

func (s *ShaderProperty) MarshalJSON() ([]byte, error) {
	var data []byte

	switch v := s.Value.(type) {
	case bool:
		data = []byte(fmt.Sprintf("%s:%v", PropertyTypeBool, v))
	case float32:
		data = []byte(fmt.Sprintf("%s:%v", PropertyTypeFloat, v))
	case int32:
		data = []byte(fmt.Sprintf("%s:%v", PropertyTypeInt, v))
	case uint32:
		data = []byte(fmt.Sprintf("%s:%v", PropertyTypeUint, v))
	case mgl32.Vec2:
		data = []byte(fmt.Sprintf("%s:%f,%f", PropertyTypeVec2, v[0], v[1]))
	case mgl32.Vec3:
		data = []byte(fmt.Sprintf("%s:%f,%f,%f", PropertyTypeVec3, v[0], v[1], v[2]))
	case mgl32.Vec4:
		data = []byte(fmt.Sprintf("%s:%f,%f,%f,%f", PropertyTypeVec4, v[0], v[1], v[2], v[3]))
	case mgl32.Mat2:
		data = []byte(fmt.Sprintf("%s:%f,%f,%f,%f", PropertyTypeMat2, v[0], v[1], v[2], v[3]))
	case mgl32.Mat3:
		data = []byte(PropertyTypeMat3 + ":" + joinMat3(v))
	case mgl32.Mat4:
		data = []byte(PropertyTypeMat4 + ":" + joinMat4(v))
	default:
		return nil, fmt.Errorf("unknown type: %s", reflect.TypeOf(s.Value).Name())
	}

	return data, nil
}

func joinMat3(m mgl32.Mat3) string {
	var str string

	for i, v := range m {
		str += fmt.Sprintf("%s%f", str, v)
		if i < len(m)-1 {
			str += ","
		}
	}

	return str
}

func joinMat4(m mgl32.Mat4) string {
	var str string

	for i, v := range m {
		str += fmt.Sprintf("%s%f", str, v)
		if i < len(m)-1 {
			str += ","
		}
	}

	return str
}
