package receiver

import (
	"fmt"
	"reflect"

	"github.com/jpaulm/gofbp/core"
)

var Name string = "Receiver"

var ipt *core.InPort

func OpenPorts(p *core.Process) {
	ipt = p.OpenInPort("IN")
}

func Execute(p *core.Process) {
	fmt.Println(p.Name + " started")

	for {

		//var pkt = p.Receive(p.InConn)
		var pkt = p.Receive(ipt.Conn)
		if pkt == nil {
			break
		}
		v := reflect.ValueOf(pkt.Contents) // display contents - assume string
		s := v.String()
		fmt.Println("Output: " + s)
		p.Discard(pkt)
	}
	fmt.Println(p.Name + " ended")
}