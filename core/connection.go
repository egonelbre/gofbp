package core

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Based on https://stackoverflow.com/questions/36857167/how-to-correctly-use-sync-cond

type Connection struct {
	network     *Network
	pktArray    []*Packet
	is, ir      int // send index and receive index
	mtx         sync.Mutex
	condNE      sync.Cond
	condNF      sync.Cond
	closed      bool
	upStrmCnt   int
	portName    string
	fullName    string
	array       []*Connection
	downStrProc *Process
}

func (c *Connection) send(p *Process, pkt *Packet) bool {
	if pkt.owner != p {
		panic("Sending packet not owned by this process")
	}
	c.condNF.L.Lock()
	defer c.condNF.L.Unlock()
	fmt.Println(p.name, "Sending", pkt.Contents)
	c.downStrProc.ensureRunning()
	c.condNE.Broadcast()
	for c.nolockIsFull() { // connection is full
		atomic.StoreInt32(&p.status, SuspSend)
		c.condNF.Wait()
		atomic.StoreInt32(&p.status, Active)
	}
	fmt.Println(p.name, "Sent", pkt.Contents)
	c.pktArray[c.is] = pkt
	c.is = (c.is + 1) % len(c.pktArray)
	pkt.owner = nil
	p.ownedPkts--
	return true
}

func (c *Connection) receive(p *Process) *Packet {
	c.condNE.L.Lock()
	defer c.condNE.L.Unlock()

	fmt.Println(p.name, "Receiving")
	for c.nolockIsEmpty() { // connection is empty
		if c.closed {
			c.condNF.Broadcast()
			return nil
		}
		atomic.StoreInt32(&p.status, SuspRecv)
		c.condNE.Wait()
		atomic.StoreInt32(&p.status, Active)

	}
	pkt := c.pktArray[c.ir]
	c.pktArray[c.ir] = nil
	fmt.Println(p.name, "Received", pkt.Contents)
	c.ir = (c.ir + 1) % len(c.pktArray)
	pkt.owner = p
	p.ownedPkts++
	c.condNF.Broadcast()

	return pkt
}

func (c *Connection) incUpstream() {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.upStrmCnt++
}

func (c *Connection) decUpstream() {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.upStrmCnt--
	if c.upStrmCnt == 0 {
		c.closed = true
		c.condNE.Broadcast()
		c.downStrProc.ensureRunning()

	}
}

func (c *Connection) Close() {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.closed = true
	c.condNE.Broadcast()
	c.downStrProc.ensureRunning()
}

func (c *Connection) isDrained() bool {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	return c.nolockIsEmpty() && c.closed
}

func (c *Connection) IsEmpty() bool {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	return c.nolockIsEmpty()
}

func (c *Connection) nolockIsEmpty() bool {
	return c.ir == c.is && c.pktArray[c.is] == nil
}

func (c *Connection) IsClosed() bool {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	return c.closed
}

func (c *Connection) nolockIsFull() bool {
	return c.ir == c.is && c.pktArray[c.is] != nil
}

func (c *Connection) resetForNextExecution() {}

func (c *Connection) GetType() string {
	return "Connection"
}

func (c *Connection) GetArrayItem(i int) *Connection {
	return nil
}

func (c *Connection) SetArrayItem(c2 *Connection, i int) {}

func (c *Connection) ArrayLength() int {
	return 0
}
