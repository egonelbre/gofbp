package testrtn

import (
	"fmt"

	"github.com/jpaulm/gofbp/core"
)

type WriteToConsole struct {
	ipt core.InputConn
	//opt core.OutputConn
}

func (writeToConsole *WriteToConsole) Setup(p *core.Process) {
	writeToConsole.ipt = p.OpenInPort("IN")
	//writeToConsole.opt = p.OpenOutPort("OUT")
	//writeToConsole.opt.SetOptional(true)
}

func (WriteToConsole) MustRun() {}

func (writeToConsole *WriteToConsole) Execute(p *core.Process) {
	fmt.Println(p.Name + " started")

	for {
		var pkt = p.Receive(writeToConsole.ipt)
		if pkt == nil {
			break
		}
		fmt.Println(pkt.Contents)

		//p.Send(writeToConsole.opt.(*core.OutPort), pkt)
		p.Discard(pkt)
	}

	fmt.Println(p.Name + " ended")
}
