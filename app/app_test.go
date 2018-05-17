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

package app

import (
	"testing"

	"github.com/juju/errors"

	"github.com/haakenlabs/ember/core"
)

type goodSystem struct {}
type badSystem struct {}
type otherSystem struct {}

func (s *goodSystem) Name() string { return "goodSystem" }
func (s *goodSystem) Setup() error { return nil }
func (s *goodSystem) Teardown() {}

func (s *badSystem) Name() string { return "badSystem" }
func (s *badSystem) Setup() error { return errors.New("expected to fail") }
func (s *badSystem) Teardown() {}

func (s *otherSystem) Name() string { return "otherSystem" }
func (s *otherSystem) Setup() error { return nil }
func (s *otherSystem) Teardown() {}

func TestApp_RegisterSystem(t *testing.T) {
	var tests = []struct{
		in core.System
		want error
	}{
		{in: &goodSystem{}, want: nil},
		{in: &badSystem{}, want: nil},
		{in: &goodSystem{}, want: core.ErrSystemExists("goodSystem")},
	}

	app := NewApp()

	for i, v := range tests {
		err := app.RegisterSystem(v.in)
		if err != v.want {
			t.Errorf(
				"%s failed on case %d. want: %v got: %v",
				t.Name(), i, v.want, err,
			)
		}
	}
}

func TestApp_System(t *testing.T) {
	var tests = []struct{
		in core.System
		skipRegister bool
		want error
	}{
		{in: &goodSystem{}, want: nil},
		{in: &badSystem{}, want: nil},
		{in: &otherSystem{}, skipRegister: true, want: core.ErrSystemNotFound("otherSystem")},
	}

	app := NewApp()

	for i, v := range tests {
		if !v.skipRegister {
			app.RegisterSystem(v.in)
		}

		_, err := app.System(v.in.Name())
		if err != v.want {
			t.Errorf(
				"%s failed on case %d. want: %v got: %v",
				t.Name(), i, v.want, err,
			)
		}
	}
}

func TestApp_Setup(t *testing.T) {
	var tests = []struct{
		in []core.System
		wantErr bool
	}{
		{in: []core.System{&goodSystem{}}},
		{in: []core.System{&goodSystem{}, &otherSystem{}}},
		{in: []core.System{&goodSystem{}, &badSystem{}}, wantErr: true},
	}

	app := NewApp()
	app.PreSetupFunc = func() error { return nil }
	app.PostSetupFunc = func() error { return nil }
	app.PreTeardownFunc = func() { }
	app.PostTeardownFunc = func() { }

	for i, v := range tests {

		for j := range v.in {
			app.RegisterSystem(v.in[j])
		}

		err := app.Setup()
		if (err != nil) != v.wantErr {
			t.Errorf("%s failed on case %d. wantErr: %v gotErr: %v err: %v", t.Name(), i, v.wantErr, err != nil, err)
		}

		app.Quit()
		app.Teardown()

		app.systems = nil
		appInst = nil
	}
}