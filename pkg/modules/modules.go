// Copyright Safing ICS Technologies GmbH. Use of this source code is governed by the AGPL license that can be found in the LICENSE file.

package modules

import (
	"container/list"
	"echo-basic/pkg/atomicbool"
	"os"
	"time"
)

var modules *list.List
var addModule chan *Module
var GlobalShutdown chan bool

type Module struct {
	Name  string
	Order uint8

	Start         chan bool
	Active        *atomicbool.AtomicBool
	startComplete chan bool

	Stop         chan bool
	Stopped      *atomicbool.AtomicBool
	stopComplete chan bool
}

func Register(name string, order uint8) *Module {
	newModule := &Module{
		Name:          name,
		Order:         order,
		Start:         make(chan bool),
		Active:        atomicbool.NewBool(true),
		startComplete: make(chan bool),

		Stop:         make(chan bool),
		Stopped:      atomicbool.NewBool(false),
		stopComplete: make(chan bool),
	}
	addModule <- newModule
	return newModule
}

func (module *Module) addToList() {
	for e := modules.Back(); e != nil; e = e.Prev() {
		if module.Order > e.Value.(*Module).Order {
			modules.InsertAfter(module, e)
			return
		}
	}
	modules.PushFront(module)
}

func (module *Module) stop() {
	module.Active.Set()
	defer module.Stopped.Set()
	for {
		select {
		case module.Stop <- true:
		case <-module.stopComplete:
			return
		case <-time.After(1 * time.Second):
		}
	}
}

func (module *Module) StopComplete() {
	module.stopComplete <- true
}

func (module *Module) start() {
	module.Stopped.UnSet()
	defer module.Active.Set()
	for {
		select {
		case module.Start <- true:
		case <-module.startComplete:
			return
		}
	}
}

func (module *Module) StartComplete() {
	module.startComplete <- true
}

func InitiateFullShutdown() {
	close(GlobalShutdown)
}

func fullStop() {
	for e := modules.Back(); e != nil; e = e.Prev() {
		if e.Value.(*Module).Active.IsSet() {
			e.Value.(*Module).stop()
		}
	}
}

func run() {
	for {
		select {
		case <-GlobalShutdown:
			fullStop()
			os.Exit(0)
		case m := <-addModule:
			m.addToList()
		}
	}
}

func init() {

	modules = list.New()
	addModule = make(chan *Module, 10)
	GlobalShutdown = make(chan bool)

	go run()

}
