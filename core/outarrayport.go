package core

//var _ Conn = (*InArrayPort)(nil)

type OutArrayPort struct {
	network *Network

	portName string
	fullName string
	array    []*OutPort
	closed   bool
}

func (o *OutArrayPort) send(p *Process, pkt *Packet) bool { panic("send on array port") }

func (o *OutArrayPort) SetOptional(b bool) {}

func (o *OutArrayPort) GetType() string {
	return "OutArrayPort"
}

func (o *OutArrayPort) GetArrayItem(i int) *OutPort {
	if i >= len(o.array) {
		return nil
	}
	return o.array[i]
}

func (o *OutArrayPort) SetArrayItem(o2 *OutPort, i int) {
	if i >= len(o.array) {
		// add to .array to fit c2
		increaseBy := make([]*OutPort, i-len(o.array)+1)
		o.array = append(o.array, increaseBy...)
	}
	o.array[i] = o2
}

func (o *OutArrayPort) ArrayLength() int {
	return len(o.array)
}

func (o *OutArrayPort) Close() {
	for _, v := range o.array {
		v.Close()
	}
}
