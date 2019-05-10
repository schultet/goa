package comm

import (
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net"
	"reflect"
)

const (
	// MsgBufferSize specifies how many messages will be buffered
	MsgBufferSize = 100000
)

// Message is the interface type for any message
type Message interface{}

// Server is the interface that wraps the basic Messaging methods.
type Server interface {
	// Quit quits the servers client connections
	Quit()

	// RegisterMessageChan registers a type and a channel to the server. When
	// the server receives a message of the registered type, it is forwarded to
	// the registered channel.
	RegisterMessageChan(interface{}, chan interface{})
}

// MessageHandler is called on message objects. Returns a boolean stating
// whether handling the message object was successful. The primary use case is
// to filter messages by their type and then put them on a channel. (When the
// MessageHandler is created, the channel should be in its closure).
type MessageHandler func(m interface{}) (success bool)

// ChanServer instances use channels instead of tcp/ip (or other type of
// network) connections
type ChanServer struct {
	incoming    map[int]chan Message
	agents      ConnList
	agentID     int
	lastIndex   int // for FetchNext save last position
	cases       []reflect.SelectCase
	nextAgentID map[int]int
	handlers    []MessageHandler
	//m           sync.Mutex
}

// NewChanServer returns a new ChanServer instance
func NewChanServer(agentID int, agents ConnList, in map[int]chan Message) *ChanServer {
	server := &ChanServer{
		//incoming: make(map[int]chan Message),
		incoming:    in,
		agents:      agents,
		agentID:     agentID,
		cases:       make([]reflect.SelectCase, len(agents)),
		nextAgentID: make(map[int]int, len(agents)),
	}
	for i, agent := range agents {
		//server.incoming[agent.ID] = make(chan Message, 2)
		//go MessageBuffer(in[agent.ID], server.incoming[agent.ID])
		server.cases[i] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(server.incoming[agent.AgentID]),
		}
		server.nextAgentID[agent.AgentID] = agents[(i+1)%len(agents)].AgentID
	}
	server.lastIndex = server.agents[0].AgentID
	//fmt.Printf("server.nextAgentID = %+v\n", server.nextAgentID)
	go server.applyHandlers()
	return server
}

func (s *ChanServer) applyHandlers() {
	for {
		//s.m.Lock()
		s.lastIndex = s.nextAgentID[s.lastIndex]
		select {
		case m := <-s.incoming[s.lastIndex]:
			for _, handler := range s.handlers {
				if handler(m) {
					break
				}
			}
		default:
		}
		//s.m.Unlock()
	}
}

// RegisterMessageChan registers a channel together with a type to the server.
// When the server receives a message of the registered type, it is forwarded to
// the registered channel.
func (s *ChanServer) RegisterMessageChan(t interface{}, ch chan interface{}) {
	//s.m.Lock()
	h := func(m interface{}) bool {
		if reflect.TypeOf(m) == reflect.TypeOf(t) {
			ch <- m
			return true
		}
		return false
	}
	s.handlers = append(s.handlers, h)
	//s.m.Unlock()
}

// Quit quits the servers client connections
func (s *ChanServer) Quit() { /* TODO close(s.incoming) */ }

//func (s *ChanServer) FetchNext() Message {
//	for i := 0; i < len(s.agents); i++ {
//		s.lastIndex = s.nextAgentID[s.lastIndex]
//		//j := rand.Intn(len(s.agents))
//		select {
//		case m := <-s.incoming[s.lastIndex]:
//			return m
//		default:
//		}
//	}
//	return nil
//}
//
//func (s *ChanServer) FetchNextWait() Message {
//	_, value, ok := reflect.Select(s.cases)
//	if ok {
//		return value.Interface().(Message)
//	}
//	return nil
//}

// Checks whether an error occured and prints error to os.Stderr
func checkError(err error) {
	if err != nil {
		log.Fatalf("Fatal error: %s\n", err.Error())
	}
}

// Logs an error, when err is not nil
func logIfError(err error) {
	if err != nil {
		log.Printf("%s\n", err)
	}
}

// TCPServer ...
type TCPServer struct {
	host      string
	port      int
	agentID   int
	agents    ConnList
	incoming  map[int]chan Message
	conns     map[int]net.Conn
	cases     []reflect.SelectCase
	quit      chan bool
	closed    bool
	lastIndex int
	listener  net.Listener
	handlers  []MessageHandler
}

// NewTCPServer creates, runs, and returns a new TcpServer2
func NewTCPServer(host string, port int, agents ConnList,
	agentID int) *TCPServer {
	s := &TCPServer{
		host:      host,
		port:      port,
		agents:    agents,
		incoming:  make(map[int]chan Message),
		quit:      make(chan bool, 1),
		conns:     make(map[int]net.Conn),
		cases:     make([]reflect.SelectCase, len(agents)),
		lastIndex: 0,
	}
	for i, agent := range agents {
		s.incoming[agent.AgentID] = make(chan Message, MsgBufferSize)
		s.cases[i] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(s.incoming[agent.AgentID]),
		}
	}
	go s.run()
	return s
}

func (s *TCPServer) run() {
	var tcpAddr = net.TCPAddr{IP: net.ParseIP(s.host), Port: s.port}
	var err error
	s.listener, err = net.ListenTCP("tcp", &tcpAddr)
	checkError(err)

L:
	for {
		select {
		case <-s.quit:
			break L
		default:
		}
		conn, err := s.listener.Accept()
		if err != nil {
			select {
			case <-s.quit:
				break L
			default:
			}
			continue
		}
		go s.handle(conn) //, s.incoming[int(agentID)])
	}

	// server quit cleanup
	s.listener.Close()
	fmt.Printf("QUITTING SERVER: %+v\n", tcpAddr)
	for _, c := range s.conns {
		c.Close()
	}
}

func (s *TCPServer) handle(c net.Conn) { //, ch chan Message) {
	dec := gob.NewDecoder(c)
	var agentID int
	err := dec.Decode(&agentID)
	if err != nil {
		log.Fatalf("STH WENT WRONG (registering clientlistener): %s", err)
	}

	s.conns[agentID] = c
	ch := s.incoming[agentID]

L:
	for {
		var m Message
		err := dec.Decode(&m)
		if err != nil && err != io.EOF {
			log.Printf("m = %+v\nerr = %+v\n%T\n", m, err, err)
			continue
		}
		// handle registered message types directly
		for _, handler := range s.handlers {
			if handler(m) {
				continue L
			}
		}
		ch <- m
	}
	//c.Close()
}

// Quit quits the servers client connections
func (s *TCPServer) Quit() {
	close(s.quit)
}

// RegisterMessageChan registers a channel together with a type to the server.
// When the server receives a message of the registered type, it is forwarded to
// the registered channel.
func (s *TCPServer) RegisterMessageChan(t interface{}, ch chan interface{}) {
	gob.Register(reflect.TypeOf(t))
	h := func(m interface{}) bool {
		if reflect.TypeOf(m) == reflect.TypeOf(t) {
			ch <- m
			return true
		}
		return false
	}
	s.handlers = append(s.handlers, h)
}

//func (s *tcpServer2) FetchNext() Message {
//	for i := 0; i < len(s.agents); i++ {
//		//s.lastIndex = (s.lastIndex + 1) % len(s.agents)
//		j := rand.Intn(len(s.agents))
//		select {
//		case m := <-s.incoming[s.agents[j].ID]:
//			return m
//		default:
//		}
//	}
//	return nil
//}
//
//func (s *tcpServer2) FetchNextWait() Message {
//	_, value, ok := reflect.Select(s.cases)
//	if ok {
//		return value.Interface().(Message)
//	}
//	return nil
//}
