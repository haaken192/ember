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
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/juju/errors"
	"github.com/sirupsen/logrus"

	"github.com/haakenlabs/ember/core"
	"github.com/haakenlabs/ember/gfx"
)

const (
	// maxFrameSkip is the maximum number of frames that are allowed to
	// be skipped during logic updates.
	maxFrameSkip = 5

	// builtinAssets is the path to the builtin asset bundle.
	builtinAssets = "<builtin>:builtin.json"
)

var (
	// appInst is the currently running app.
	appInst *App

	// appOnce is used to set the appInst variable only once.
	appOnce sync.Once
)

// App is the backbone of any Ember application.
type App struct {
	// Name is the name of this app.
	Name string

	// Company is the name of the company responsible for this app.
	Company string

	// PreSetupFunc is a callback invoked prior to app setup.
	PreSetupFunc func() error

	// PostSetupFunc is a callback invoked after app setup.
	PostSetupFunc func() error

	// PreTeardownFunc is a callback invoked prior to app teardown.
	PreTeardownFunc func()

	// PostTeardownFunc is a callback invoked after app teardown.
	PostTeardownFunc func()

	// systems is a list of systems used by this app.
	systems []core.System

	// running indicates that the application is running.
	running bool
}

// Setup sets up the App.
func (a *App) Setup() error {
	if appInst != nil {
		return errors.New("app already created")
	}
	setApp(a)

	if a.PreSetupFunc != nil {
		if err := a.PreSetupFunc(); err != nil {
			return err
		}
	}

	for i := range a.systems {
		logrus.Debug("Setting up system: ", a.systems[i].Name())

		if err := a.systems[i].Setup(); err != nil {
			return err
		}
	}

	//
	//if err := asset.LoadManifest(builtinAssets); err != nil {
	//	return err
	//}

	if a.PostSetupFunc != nil {
		if err := a.PostSetupFunc(); err != nil {
			return err
		}
	}

	return nil
}

// Teardown tears down the app.
func (a *App) Teardown() {
	if a.PreTeardownFunc != nil {
		a.PreTeardownFunc()
	}

	for i := len(a.systems) - 1; i >= 0; i-- {
		logrus.Debug("Tearing down system: ", a.systems[i].Name())

		a.systems[i].Teardown()
	}

	if a.PostTeardownFunc != nil {
		a.PostTeardownFunc()
	}
}

// Quit instructs the App to shutdown by setting the running variable to false.
func (a *App) Quit() {
	a.running = false
}

// RegisterSystem registers a system with the App. A system can only be added
// once, it is an error to add a system more than once. Systems are initialized
// in the order they are added and torn down in the reverse order.
func (a *App) RegisterSystem(s core.System) error {
	// Check for existing system.
	if a.SystemRegistered(s.Name()) {
		return core.ErrSystemExists(s.Name())
	}

	// Add system to the systems slice.
	a.systems = append(a.systems, s)

	logrus.Debug("Registered system: ", s.Name())

	return nil
}

// SystemRegistered returns true if the system with the given name is registered
// with this App.
func (a *App) SystemRegistered(name string) bool {
	for i := range a.systems {
		if a.systems[i].Name() == name {
			return true
		}
	}

	return false
}

// System returns a system by the given name.
func (a *App) System(name string) (core.System, error) {
	for i := range a.systems {
		if a.systems[i].Name() == name {
			return a.systems[i], nil
		}
	}

	return nil, core.ErrSystemNotFound(name)
}

// MustSystem is like System, but panics if the system cannot be found.
func (a *App) MustSystem(name string) core.System {
	s, err := a.System(name)
	if err != nil {
		panic(err)
	}

	return s
}

func (a *App) setupSignalHandler() {
	s := make(chan os.Signal)
	signal.Notify(s, os.Interrupt, syscall.SIGTERM)
	go handleSignal(s, a)
}

func handleSignal(s chan os.Signal, a *App) {
	<-s
	a.Quit()
}

// CurrentApp returns the currently running app.
func CurrentApp() *App {
	return appInst
}

// NewApp creates a new application.
func NewApp(renderer gfx.Renderer) *App {
	a := &App{}

	a.RegisterSystem(core.NewWindowSystem(a.Name, renderer))
	a.RegisterSystem(core.NewInstanceSystem())
	a.RegisterSystem(core.NewAssetSystem())
	a.RegisterSystem(core.NewTimeSystem())
	a.RegisterSystem(core.NewSceneSystem())

	//asset.RegisterHandler(texture.NewHandler())
	//asset.RegisterHandler(shader.NewHandler())
	//asset.RegisterHandler(mesh.NewHandler())
	//asset.RegisterHandler(font.NewHandler())
	//asset.RegisterHandler(skybox.NewHandler())

	return a
}

// setApp sets the App, but only once.
func setApp(a *App) {
	appOnce.Do(func() {
		appInst = a
	})
}
