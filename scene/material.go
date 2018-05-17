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
	"fmt"
	"strconv"

	"github.com/haakenlabs/ember/core"
	"github.com/haakenlabs/ember/gfx"
	"github.com/haakenlabs/ember/system/asset/shader"
	"github.com/haakenlabs/ember/system/asset/texture"
	"github.com/haakenlabs/ember/system/instance"
)

type MaterialTexture uint32

const (
	MaterialTextureAttachment0 MaterialTexture = iota
	MaterialTextureAttachment1
	MaterialTextureDepth
	MaterialTextureEnvironment
	MaterialTextureIrradiance
	MaterialTextureAlbedo
	MaterialTextureNormal
	MaterialTextureMetallic
)

// MaterialMaxTextures is the maximum number of textures supported
// by a material.
const MaterialMaxTextures = 16

type MaterialData struct {
	Textures         map[string]string             `json:"textures"`
	ShaderProperties map[string]gfx.ShaderProperty `json:"shader_properties"`
	Shader           string                        `json:"shader"`
}

type Material struct {
	core.BaseObject

	textures         [MaterialMaxTextures]gfx.Texture
	shaderProperties map[string]interface{}
	shader           gfx.Shader
}

func (m *Material) SetTexture(id MaterialTexture, texture gfx.Texture) {
	if id < MaterialMaxTextures {
		m.textures[id] = texture
	}
}

func (m *Material) SetShader(shader gfx.Shader) {
	m.shader = shader
}

func (m *Material) Texture(id MaterialTexture) gfx.Texture {
	if id >= MaterialMaxTextures {
		return nil
	}

	return m.textures[id]
}

func (m *Material) Shader() gfx.Shader {
	return m.shader
}

func (m *Material) Bind() {
	if m.shader == nil {
		return
	}

	m.shader.Bind()

	for i := range m.textures {
		if m.textures[i] != nil {
			m.textures[i].Activate(uint32(i))
		}
	}
	for key, value := range m.shaderProperties {
		m.shader.SetUniform(key, value)
	}
}

func (m *Material) Unbind() {
	m.shader.Unbind()
}

func (m *Material) SupportsDeferredPath() bool {
	if m.shader != nil {
		return m.shader.Deferred()
	}

	return false
}

func (m *Material) SetProperty(property string, value interface{}) {
	m.shaderProperties[property] = value
}

func NewMaterial() *Material {
	m := &Material{
		shaderProperties: make(map[string]interface{}),
	}

	m.SetName("Material")
	instance.MustAssign(m)

	return m
}

func BuildMaterial(data *MaterialData) (*Material, error) {
	var err error

	m := NewMaterial()

	/* Populate shader */
	m.shader, err = shader.Get(data.Shader)
	if err != nil {
		return nil, err
	}

	/* Populate textures. */
	for k, v := range data.Textures {
		/* Attempt to parse format: "0": "texture.png" */
		i, err := strconv.ParseUint(k, 10, 8)
		if err != nil {
			return nil, err
		}
		if i < 16 {
			t, err := texture.Get(v)
			if err != nil {
				return nil, err
			}
			m.textures[i] = t
		} else {
			return nil, fmt.Errorf("texture index out of range: %d", i)
		}
	}

	/* Populate shader properties. */
	for k, v := range data.ShaderProperties {
		m.shaderProperties[k] = v
	}

	return m, nil
}
