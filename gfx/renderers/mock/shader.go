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
	"github.com/sirupsen/logrus"

	"github.com/haakenlabs/ember/gfx"
)

var _ gfx.Shader = &Shader{}

type Shader struct{}

func (s *Shader) Bind() {}

func (s *Shader) Unbind() {}

func (s *Shader) Reference() uint32 {
	return 1
}

func (s *Shader) Alloc() error {
	logrus.Info("alloc mock shader")
	return nil
}

func (s *Shader) Dealloc() {}

func (s *Shader) Deferred() bool {
	return true
}

func (s *Shader) AddData([]byte) {}

func (s *Shader) ResetData() {}

func (s *Shader) Compile() error {
	return nil
}

func (s *Shader) ID() int32 {
	return 1
}

func (r *Renderer) MakeShader(bool) gfx.Shader {
	return &Shader{}
}
