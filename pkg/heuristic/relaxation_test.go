package heuristic

//import (
//	"github.com/schultet/goa/src/comm"
//)

//func setComm(agentID int, agents comm.AgentList) (comm.Server, comm.Dispatcher) {
//	agentInfo, _ := agents.Get(agentID) // TODO: remove second return value
//	otherInfo := agents.Except(agentID)
//	server := comm.NewTcpServer2(agentInfo.Host, agentInfo.Port, otherInfo, agentID)
//	dispatcher := comm.NewDispatcher(agentID, otherInfo)
//	return server, dispatcher
//}

//func TestDistributedFFHeuristicEvaluate(t *testing.T) {
//	t0 := task.NewTaskFromFile("tests/0.json")
//	t1 := task.NewTaskFromFile("tests/1.json")
//	// agents := comm.AgentList(...)
//	server0, dispatcher0 := setComm(t0.AgentID, agents)
//	server1, dispatcher1 := setComm(t1.AgentID, agents)
//
//	// Test sending of heuristic messages
//	// dispatcher0.Send(byte(t1.AgentID), comm.HeuristicReqMSG(...))
//	// m := server1.FetchNextWait()
//
//	defer server0.Quit()
//	defer server1.Quit()
//	// TODO: create heuristic objects for each agent
//	// test Evaluate, Message passing, etc.
//	dff0 := NewDistributedFFHeuristic(t0, server0, dispatcher0)
//	dff1 := NewDistributedFFHeuristic(t1, server1, dispatcher1)
//
//	// state = statespace.State(...)
//	res := dff0.Evaluate(state)
//
//	//TODO: COMPLETE
//
//}
