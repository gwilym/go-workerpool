package workerpool

/*
Copyright (C) 2015  Gwilym Evans

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

import (
	"sync"
	"sync/atomic"
)

type Workerpool interface {
	Start()
	Stop()
	Wait()
	CountWorkers() int32
}

type WorkerFunction func() bool

func NewFunctionWorkerpool(concurrency int32, function WorkerFunction) *FunctionWorkerpool {
	return &FunctionWorkerpool{
		concurrency: concurrency,
		function:    function,
		workers:     int32(0),
		workerGroup: &sync.WaitGroup{},
	}
}

type FunctionWorkerpool struct {
	concurrency int32
	workers     int32
	workerGroup *sync.WaitGroup
	function    WorkerFunction
	running     bool
	stopping    bool
}

// Start begins the worker routines, if not already started.
func (f *FunctionWorkerpool) Start() {
	if !f.running {
		f.running = true
		f.stopping = false
		for i := int32(0); i < f.concurrency; i++ {
			f.workerGroup.Add(1)
			atomic.AddInt32(&f.workers, 1)
			go f.work()
		}
	}
}

// Stop signals workers to stop, but returns immediately. Wait() should be used to wait for workers to stop.
func (f *FunctionWorkerpool) Stop() {
	f.stopping = true
}

// Wait blocks until all workers finish. Depending on the workers, this may block indefinitely unless Stop is called first.
func (f *FunctionWorkerpool) Wait() {
	f.workerGroup.Wait()
	f.running = false
}

// CountWorkers returns a count of the currently active workers.
func (f *FunctionWorkerpool) CountWorkers() int32 {
	return atomic.LoadInt32(&f.workers)
}

func (f *FunctionWorkerpool) work() {
	for f.running && !f.stopping && f.function() {
	}
	atomic.AddInt32(&f.workers, -1)
	f.workerGroup.Done()
}
