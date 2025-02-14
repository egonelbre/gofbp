package testrtn

import (
	//"fmt"

	"github.com/jpaulm/gofbp/core"
)

type Kick struct {
	opt core.OutputConn
}

func (kick *Kick) Setup(p *core.Process) {
	kick.opt = p.OpenOutPort("OUT")
}

func (kick *Kick) Execute(p *core.Process) {
	//fmt.Println(p.GetName() + " started")

	var pkt = p.Create("Kicker IP")
	p.Send(kick.opt, pkt)

	//fmt.Println(p.GetName() + " ended")
}
