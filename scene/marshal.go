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
	"encoding/json"
	"fmt"
)

type UnmarshalComponentFunc func([]byte) (Component, error)
type MarshalComponentFunc func(Component) ([]byte, error)

type JSONComponent struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type JSONGameObject struct {
	Name       string            `json:"name"`
	Components []*JSONComponent  `json:"components"`
	Children   []*JSONGameObject `json:"children"`
}

type JSONScene struct {
	Name        string            `json:"name"`
	GameObjects []*JSONGameObject `json:"gameobjects"`
}

var (
	componentMarshaller   = make(map[string]MarshalComponentFunc)
	componentUnmarshaller = make(map[string]UnmarshalComponentFunc)
)

func AddComponentMarshaller(t string, c MarshalComponentFunc) error {
	if _, dup := componentMarshaller[t]; dup {
		return fmt.Errorf("duplicate component marshaller for type %s", t)
	}

	componentMarshaller[t] = c

	return nil
}

func AddComponentUnmarshaller(t string, c UnmarshalComponentFunc) error {
	if _, dup := componentUnmarshaller[t]; dup {
		return fmt.Errorf("duplicate component unmarshaller for type %s", t)
	}

	componentUnmarshaller[t] = c

	return nil
}

func ComponentMarshaller(t string) (MarshalComponentFunc, error) {
	if c, ok := componentMarshaller[t]; ok {
		return c, nil
	}

	return nil, fmt.Errorf("no such component marshaller for type %s", t)
}

func ComponentUnmarshaller(t string) (UnmarshalComponentFunc, error) {
	if c, ok := componentUnmarshaller[t]; ok {
		return c, nil
	}

	return nil, fmt.Errorf("no such component marshaller for type %s", t)
}

func BuildScene(data *JSONScene) (*Scene, error) {
	s := NewScene(data.Name)

	for _, oj := range data.GameObjects {
		o, err := BuildGameObject(oj, nil)
		if err != nil {
			return nil, err
		}

		if err := s.AddObject(o, nil); err != nil {
			return nil, err
		}
	}

	return s, nil
}

func BuildGameObject(j *JSONGameObject, parent *GameObject) (*GameObject, error) {
	o := NewGameObject(j.Name)
	o.parent = parent

	for _, oj := range j.Children {
		co, err := BuildGameObject(oj, o)
		if err != nil {
			return nil, err
		}

		o.children = append(o.children, co)
	}

	for _, cj := range j.Components {
		u, err := ComponentUnmarshaller(cj.Type)
		if err != nil {
			return nil, err
		}

		c, err := u(cj.Data)
		if err != nil {
			return nil, err
		}

		o.components = append(o.components, c)
	}

	return o, nil
}
