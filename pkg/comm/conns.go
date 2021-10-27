package comm

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/schultet/goa/pkg/util/fileio"
)

// Conn stores agent network information
type Conn struct {
	AgentID int
	Port    int
	Host    string
}

// Service returns service string for a connection, i.e. "host:port"
func (c Conn) Service() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

// ConnList is a list of Conn elements
type ConnList []Conn

// CreateConnList returns a ConnList of n agent connections, starting from a
// given port and host. Port and ID are incremented by one for each connection.
func CreateConnList(n, port int, host string) ConnList {
	var conns ConnList
	for i := 0; i < n; i++ {
		conns = append(conns, Conn{AgentID: i, Port: port, Host: host})
		port++
	}
	return conns
}

// Get returns a Conn for a specified agentID, or error if no such AgentID is in the list
func (cl ConnList) Get(agentID int) (Conn, error) {
	for _, conn := range cl {
		if conn.AgentID == agentID {
			return conn, nil
		}
	}
	return Conn{}, errors.New("no conn with specified agentID in ConnList")
}

// GetMinID returns the smallest AgentID of all Conns in the list
func (cl ConnList) GetMinID() int {
	min := int(^uint(0) >> 1) // max int value
	for _, conn := range cl {
		if conn.AgentID < min {
			min = conn.AgentID
		}
	}
	return min
}

// NextID returns the next Conns ID round robin style
func (cl ConnList) NextID(agentID int) int {
	for i, conn := range cl {
		if conn.AgentID == agentID {
			return cl[(i+1)%len(cl)].AgentID
		}
	}
	return -1
}

// Except returns the ConnList except for the entry of the specified ID
func (cl ConnList) Except(agentID int) ConnList {
	result := make(ConnList, 0, len(cl))
	for _, conn := range cl {
		if conn.AgentID != agentID {
			result = append(result, conn)
		}
	}
	return result
}

// Equals returns whether two ConnLists are identical
func (cl ConnList) Equals(other *ConnList) bool {
	if len(cl) != len(*other) {
		return false
	}
	for i, conn := range cl {
		if conn != (*other)[i] {
			return false
		}
	}
	return true
}

// ParseConnFile parses a file of agent descriptions with the following format:
//
// <n>
// <host_1> <port_1> <id_1>
// ...
// <host_n> <port_n> <id_n>
//
// where n specifies the number of agents,
// host_i is the i-th agents' hostname
// port_i its port and id_i its ID (an uint8 number)
// returns a list of AgentInfo structs
func ParseConnFile(filename string) (ConnList, error) {
	lines, err := fileio.ReadFileToLines(filename, true)
	if err != nil {
		return nil, err
	}

	numAgents, err := strconv.ParseInt(lines[0], 10, 0)
	if err != nil {
		fmt.Printf("Error parsing number of agents: %s\nError:%s\n", lines[0], err)
		return nil, err
	}

	var agents ConnList
	for i := 1; i <= int(numAgents); i++ {
		info := strings.Fields(lines[i])
		port, err := strconv.ParseInt(info[1], 10, 0)
		if err != nil {
			fmt.Printf("Error parsing agent port: %s\nError:%s\n", info[1], err)
			return nil, err
		}
		ID, err := strconv.ParseInt(info[2], 10, 0)
		if err != nil {
			fmt.Printf("Error parsing agent ID: %s\nError:%s\n", info[2], err)
			return nil, err
		}
		agents = append(agents, Conn{int(ID), int(port), info[0]})
	}
	return agents, err
}

// ParseAgents creates a ConnList from the provided --agent arguments
func NewConnList(conns []string) (ConnList, error) {
	var res ConnList
	for _, c := range conns {
		info := strings.Fields(c)
		if len(info) != 3 {
			return nil, fmt.Errorf("ERR: wrong agent definition: %v\n", info)
		}
		id, _ := strconv.ParseInt(info[0], 10, 32)
		port, _ := strconv.ParseInt(info[2], 10, 32)
		res = append(res, Conn{AgentID: int(id), Host: info[1], Port: int(port)})
	}
	return res, nil
}

func CreateTCPComm(agentID int, conns ConnList) (Server, Dispatcher) {
	agentInfo, _ := conns.Get(agentID) // TODO: remove second return value
	otherInfo := conns.Except(agentID)
	server := NewTCPServer(agentInfo.Host, agentInfo.Port, otherInfo, agentID)
	dispatcher := NewDispatcher(agentID, otherInfo)
	return server, dispatcher
}

func CreateChanComm(agentID int, conns ConnList, inChans, outChans map[int]map[int]chan Message) (Server, Dispatcher) {
	server := NewChanServer(agentID, conns.Except(agentID), inChans[agentID])
	dispatcher := NewChanDispatcher(agentID, conns.Except(agentID), outChans[agentID])
	return server, dispatcher
}
