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

package scene

import (
	"github.com/haakenlabs/ember/core"
)

func Register(scene core.Scene) error {
	return core.GetSceneSystem().Register(scene)
}

func Load(name string) error {
	return core.GetSceneSystem().Load(name)
}

func PurgePush(name string) error {
	return core.GetSceneSystem().PurgePush(name)
}

func Replace(name string) error {
	return core.GetSceneSystem().Replace(name)
}

func Push(name string) error {
	return core.GetSceneSystem().Push(name)
}

func Pop() string {
	return core.GetSceneSystem().Pop()
}

func RemoveAll() {
	core.GetSceneSystem().RemoveAll()
}

func Unregister(name string) error {
	return core.GetSceneSystem().Unregister(name)
}

func Registered(name string) bool {
	return core.GetSceneSystem().Registered(name)
}

func Active() core.Scene {
	return core.GetSceneSystem().Active()
}

func ActiveName() string {
	return core.GetSceneSystem().ActiveName()
}

func Count() int {
	return core.GetSceneSystem().Count()
}

func ActiveCount() int {
	return core.GetSceneSystem().ActiveCount()
}
