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

import "github.com/haakenlabs/ember/gfx"

var _ gfx.Pipeline = &PipelineDeferred{}
var _ gfx.Pipeline = &PipelineForward{}

type BasePipeline struct {
	stages []gfx.Stage
}

type PipelineDeferred struct {
	BasePipeline
}

type PipelineForward struct {
	BasePipeline
}

func (p *BasePipeline) AddStage(stages ...gfx.Stage) {
	for _, v := range stages {
		for _, x := range p.stages {
			if v.Name() == x.Name() {
				continue
			}

			p.stages = append(p.stages, v)
		}
	}
}

func (p *BasePipeline) EnableStage(name string) {
	for _, v := range p.stages {
		if v.Name() == name {
			v.SetEnabled(true)
		}
	}
}

func (p *BasePipeline) DisableStage(name string) {
	for _, v := range p.stages {
		if v.Name() == name {
			v.SetEnabled(false)
		}
	}
}

func (p *PipelineDeferred) Process(camera gfx.Camera) {
	if camera == nil {
		return
	}
}

func (p *PipelineForward) Process(camera gfx.Camera) {
	if camera == nil {
		return
	}
}
