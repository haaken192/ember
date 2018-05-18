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
	"github.com/haakenlabs/ember/gfx"
	"github.com/haakenlabs/ember/system/instance"
)

type Skybox struct {
	core.BaseObject

	radiance   gfx.Texture
	specular   gfx.Texture
	irradiance gfx.Texture
}

func NewSkybox(radiance, specular, irradiance gfx.Texture) *Skybox {
	if radiance.Type() != gfx.TextureCubemap {
		return nil
	}
	if specular.Type() != gfx.TextureCubemap {
		return nil
	}
	if irradiance.Type() != gfx.TextureCubemap {
		return nil
	}

	s := &Skybox{
		radiance:   radiance,
		specular:   specular,
		irradiance: irradiance,
	}

	s.SetName("Skybox")
	instance.MustAssign(s)

	return s
}

func (s *Skybox) Radiance() gfx.Texture {
	return s.radiance
}

func (s *Skybox) Specular() gfx.Texture {
	return s.specular
}

func (s *Skybox) Irradiance() gfx.Texture {
	return s.irradiance
}
