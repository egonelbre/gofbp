package main

import (
	"github.com/jpaulm/gofbp/components/testrtn"
	"github.com/jpaulm/gofbp/core"
	// "runtime"
)

// Merge application

func main10() {

	// runtime.GOMAXPROCS(16)

	var net *core.Network = core.NewNetwork("MergeToConsUsingNonLooper")

	proc1 := net.NewProc("Sender1", &testrtn.Sender{})
	proc2 := net.NewProc("Sender2", &testrtn.Sender{})

	proc3 := net.NewProc("WriteToConsNL", &testrtn.WriteToConsNL{})

	net.Initialize("15", proc1, "COUNT")
	net.Initialize("15", proc2, "COUNT")
	net.Connect(proc1, "OUT", proc3, "IN", 6)
	net.Connect(proc2, "OUT", proc3, "IN", 6)

	net.Run()
}