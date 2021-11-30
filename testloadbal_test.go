package main

import (
	"testing"

	"github.com/jpaulm/gofbp/components/testrtn"
	"github.com/jpaulm/gofbp/core"
)

func TestLoadBal(t *testing.T) {
	net := core.NewNetwork("TestLoadBal")

	proc1 := net.NewProc("Sender", &testrtn.Sender{})

	proc2 := net.NewProc("LoadBalance", &testrtn.LoadBalance{})

	proc3a := net.NewProc("Receiver0", &testrtn.DelayedReceiver{})
	proc3b := net.NewProc("Receiver1", &testrtn.Receiver{})
	proc3c := net.NewProc("Receiver2", &testrtn.Receiver{})

	net.Initialize("40", proc1, "COUNT")
	net.Connect(proc1, "OUT", proc2, "IN", 6)
	net.Connect(proc2, "OUT[0]", proc3a, "IN", 6)
	net.Connect(proc2, "OUT[1]", proc3b, "IN", 6)
	net.Connect(proc2, "OUT[2]", proc3c, "IN", 6)

	net.Run()
}