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

package math

import (
	"math"

	"github.com/go-gl/mathgl/mgl64"
)

type Rect64 struct {
	Min, Max mgl64.Vec2
}

func R64(minX, minY, maxX, maxY float64) Rect64 {
	return Rect64{
		Min: mgl64.Vec2{minX, minY},
		Max: mgl64.Vec2{maxX, maxY},
	}
}

func (r Rect64) Norm() Rect64 {
	return Rect64{
		Min: mgl64.Vec2{
			math.Min(r.Min.X(), r.Max.X()),
			math.Min(r.Min.Y(), r.Max.Y()),
		},
		Max: mgl64.Vec2{
			math.Max(r.Min.X(), r.Max.X()),
			math.Max(r.Min.Y(), r.Max.Y()),
		},
	}
}

func (r Rect64) Moved(delta mgl64.Vec2) Rect64 {
	return Rect64{
		Min: r.Min.Add(delta),
		Max: r.Max.Add(delta),
	}
}

func (r Rect64) Union(s Rect64) Rect64 {
	return R64(
		math.Min(r.Min.X(), s.Min.X()),
		math.Min(r.Min.Y(), s.Min.Y()),
		math.Max(r.Max.X(), s.Max.X()),
		math.Max(r.Max.Y(), s.Max.Y()),
	)
}

func (r Rect64) W() float64 {
	return r.Max.X() - r.Min.X()
}

func (r Rect64) H() float64 {
	return r.Max.Y() - r.Min.Y()
}
