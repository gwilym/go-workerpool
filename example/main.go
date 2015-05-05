package main

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
	"log"
	"math/rand"
	"time"

	"github.com/gwilym/go-workerpool"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	concurrency := int32(4)

	log.Println("Starting up with concurrency", concurrency, ". Press CTRL-C to quit.")

	done := make(chan bool, 1)
	w := workerpool.NewFunctionWorkerpool(concurrency, printSleeper)
	w.Start(done)
	<-done
}

// I sleep for a random time then randomly return true or false
func printSleeper() bool {
	d := time.Duration(rand.Int63n(5)) * time.Second
	log.Println("Worker sleeping for", d, "...")
	time.Sleep(d)
	log.Println("Worker finishing")
	return rand.Intn(4) < 3
}
