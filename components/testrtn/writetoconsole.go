package testrtn

import (
	"fmt"

	"github.com/jpaulm/gofbp/core"
)

type WriteToConsole struct {
	ipt *core.InPort
	out *core.OutPort
}

func (writeToConsole *WriteToConsole) Setup(p *core.Process) {
	writeToConsole.ipt = p.OpenInPort("IN")
	writeToConsole.out = p.OpenOutPort("OUT", "opt")
}

func (WriteToConsole) MustRun() {}

func (writeToConsole *WriteToConsole) Execute(p *core.Process) {
	//fmt.Println(p.GetName() + " started")

	for {
		var pkt = p.Receive(writeToConsole.ipt)
		if pkt == nil {
			break
		}
		fmt.Println(pkt.Contents)
		//if writeToConsole.out.GetType() == "OutPort" {
		if writeToConsole.out.IsConnected() {
			p.Send(writeToConsole.out, pkt)
		} else {
			p.Discard(pkt)
		}
	}

	//fmt.Println(p.GetName() + " ended")
}
