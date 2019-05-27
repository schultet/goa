package comm

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

const (
	// ReconnectAttempts is the number of attempts to reconnect
	ReconnectAttempts = 1000
	// ReconnectCooldownMS is the number of milliseconds to wait after each
	// reconnection attempt
	ReconnectCooldownMS = 100
)

// A Dispatcher transmits messages to clients
type Dispatcher interface {
	Send(agent int, m Message)
	Quit()
	//Notify(agent byte, m Message)
}

// ChanDispatcher is a message dispatcher based on go-channel communication
type ChanDispatcher struct {
	agentID  int
	agents   ConnList
	outgoing map[int]chan Message
}

// NewChanDispatcher returns a new message dispatcher based on go channels
func NewChanDispatcher(id int, al ConnList, chs map[int]chan Message) Dispatcher {
	cd := &ChanDispatcher{
		agentID:  id,
		agents:   al,
		outgoing: make(map[int]chan Message),
	}
	for _, agent := range al {
		cd.outgoing[agent.AgentID] = chs[agent.AgentID]
	}
	return cd
}

// Quit savely terminates the connection. For we are using channels here,
// nothing is to be done.
func (d *ChanDispatcher) Quit() { /* TODO */ }

// Send sends a message to the agent with the specified agentID
func (d *ChanDispatcher) Send(agentID int, m Message) {
	d.outgoing[int(agentID)] <- m
}

// Broadcast sends a message to all agents in an ConnList using the specified
// dispatcher.
func Broadcast(d Dispatcher, m Message, conns ConnList) {
	for _, conn := range conns {
		d.Send(conn.AgentID, m)
	}
}

// TCPDispatcher is a message dispatcher based on a tcp/ip connection
type TCPDispatcher struct {
	agentID int
	al      ConnList
	clients map[int]*Client
}

// NewDispatcher returns a new TCPDispatcher
func NewDispatcher(agentID int, al ConnList) Dispatcher {
	d := &TCPDispatcher{
		agentID: agentID,
		al:      al,
		clients: make(map[int]*Client),
	}
	var wg sync.WaitGroup
	for _, agent := range al {
		wg.Add(1)
		c := NewClientWait(agentID, agent.Service(), &wg)
		d.clients[agent.AgentID] = c
	}
	wg.Wait()
	return d
}

// Send sends a message to the specified agent(ID)
func (d *TCPDispatcher) Send(agent int, m Message) {
	d.clients[agent].outgoing <- m
}

// Quit terminates the connection to all clients
func (d *TCPDispatcher) Quit() {
	for i := range d.clients {
		d.clients[i].disconnect()
	}
}

// Client objects send messages to connected tcp-servers
type Client struct {
	agentID  int
	outgoing chan Message
	conn     net.Conn
	quit     chan bool
}

// NewClientWait creates a new Client and calls connect and run in a
// go-routine. When successfully connected WaitGroup.Done is called.
func NewClientWait(agentID int, service string, w *sync.WaitGroup) *Client {
	c := &Client{
		agentID:  agentID,
		outgoing: make(chan Message, MsgBufferSize),
		quit:     make(chan bool, 1),
	}
	go func(w *sync.WaitGroup) {
		c.connect(service, ReconnectCooldownMS, ReconnectAttempts)
		w.Done()
		c.run()
	}(w)
	log.Printf("CLIENT: %v, SERVICE: %s\n", c, service)
	return c
}

// Establishes a tcp connection between a client and a server/clientListener
func (c *Client) connect(service string, waitMs, tries int) {
	var conn net.Conn
	var err error
	var reconnects int
	for ; tries > 0; tries-- {
		conn, err = net.Dial("tcp", service)
		if err == nil {
			break
		}
		logIfError(err)
		reconnects++
		time.Sleep(time.Duration(waitMs) * time.Millisecond)
	}
	c.conn = conn
	log.Printf("successfully connected, after: %d reconnects. ID: %d\n",
		reconnects, c.agentID)
}

func interfaceEncode(enc *gob.Encoder, m Message) {
	err := enc.Encode(&m)
	if err != nil {
		log.Fatal("encode:", err)
	}
}

// TODO: make this (sending) safer, adjust select cases for reconnection
// run transmits messages received on the outgoing chan via c.conn
func (c *Client) run() {
	enc := gob.NewEncoder(c.conn)
	err := enc.Encode(&c.agentID)
	if err != nil {
		log.Fatalf("fooo failed for %v: %s", err, err)
	}
L:
	for {
		select {
		case <-c.quit:
			c.conn.Close()
			break L
		case msg := <-c.outgoing:
			interfaceEncode(enc, msg)
		default:
		}
	}
	fmt.Println("QUITTING CLIENT")
	c.conn.Close()
}

func (c *Client) disconnect() {
	close(c.quit)
}

// TODO:
// add logic for handling errors:
// - connection termination, etc.
