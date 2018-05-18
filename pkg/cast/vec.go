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

package cast

import (
	"strconv"
	"strings"

	"fmt"
	"reflect"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/juju/errors"
	"github.com/spf13/cast"

	"github.com/haakenlabs/ember/pkg/math"
)

func ParseFloat32Slice(str string, sep string, length int) ([]float32, error) {
	var values []float32

	parts := strings.Split(str, sep)
	if length > 0 && len(parts) != length {
		return nil, errors.New("slice parse error, invalid number of elements")
	}

	for _, v := range parts {
		x, err := strconv.ParseFloat(v, 32)
		if err != nil {
			return nil, errors.Annotate(err, "slice parse error")
		}

		values = append(values, float32(x))
	}

	return values, nil
}

func ParseInt32Slice(str string, sep string, length int) ([]int32, error) {
	var values []int32

	parts := strings.Split(str, sep)
	if length > 0 && len(parts) != length {
		return nil, errors.New("slice parse error, invalid number of elements")
	}

	for _, v := range parts {
		x, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return nil, errors.Annotate(err, "slice parse error")
		}

		values = append(values, int32(x))
	}

	return values, nil
}

func ParseUint32Slice(str string, sep string, length int) ([]uint32, error) {
	var values []uint32

	parts := strings.Split(str, sep)
	if length > 0 && len(parts) != length {
		return nil, errors.New("slice parse error, invalid number of elements")
	}

	for _, v := range parts {
		x, err := strconv.ParseFloat(v, 32)
		if err != nil {
			return nil, errors.Annotate(err, "slice parse error")
		}

		values = append(values, uint32(x))
	}

	return values, nil
}

func ParseVec2(str string) (mgl32.Vec2, error) {
	v, err := ParseFloat32Slice(str, ",", 2)
	if err != nil {
		return mgl32.Vec2{}, err
	}

	return mgl32.Vec2{v[0], v[1]}, nil
}

func ParseVec3(str string) (mgl32.Vec3, error) {
	v, err := ParseFloat32Slice(str, ",", 3)
	if err != nil {
		return mgl32.Vec3{}, err
	}

	return mgl32.Vec3{v[0], v[1], v[2]}, nil
}

func ParseVec4(str string) (mgl32.Vec4, error) {
	v, err := ParseFloat32Slice(str, ",", 4)
	if err != nil {
		return mgl32.Vec4{}, err
	}

	return mgl32.Vec4{v[0], v[1], v[2], v[3]}, nil
}

func ToVec2(i interface{}) mgl32.Vec2 {
	v, _ := ToVec2E(i)

	return v
}

func ToIVec2(i interface{}) math.IVec2 {
	v, e := ToIVec2E(i)

	if e != nil {
		panic(e)
	}

	return v
}

func ToVec2E(i interface{}) (mgl32.Vec2, error) {
	if i == nil {
		return mgl32.Vec2{}, fmt.Errorf("unable to cast %#v to mgl32.Vec2", i)
	}

	switch v := i.(type) {
	case mgl32.Vec2:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		if s.Len() != 2 {
			return mgl32.Vec2{}, fmt.Errorf("unable to cast %#v to mgl32.Vec2", i)
		}
		a := mgl32.Vec2{}
		for j := 0; j < s.Len(); j++ {
			val, err := cast.ToIntE(s.Index(j).Interface())
			if err != nil {
				return mgl32.Vec2{}, fmt.Errorf("unable to cast %#v to mgl32.Vec2", i)
			}
			a[j] = float32(val)
		}

		return a, nil

	default:
		return mgl32.Vec2{}, fmt.Errorf("unable to cast %#v to mgl32.Vec2", i)
	}
}

func ToIVec2E(i interface{}) (math.IVec2, error) {
	if i == nil {
		return math.IVec2{}, fmt.Errorf("unable to cast %#v to IVec2", i)
	}

	switch v := i.(type) {
	case math.IVec2:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		if s.Len() != 2 {
			return math.IVec2{}, fmt.Errorf("unable to cast %#v to IVec2", i)
		}
		a := math.IVec2{}
		for j := 0; j < s.Len(); j++ {
			val, err := cast.ToIntE(s.Index(j).Interface())
			if err != nil {
				return math.IVec2{}, fmt.Errorf("unable to cast %#v to IVec2", i)
			}
			a[j] = int32(val)
		}

		return a, nil

	default:
		return math.IVec2{}, fmt.Errorf("unable to cast %#v to IVec2", i)
	}
}
