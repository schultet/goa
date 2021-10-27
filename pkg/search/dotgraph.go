package search

import (
	"fmt"
	"os"
)

var (
	colors = []string{
		"chartreuse",
		"goldenrod",
		"aquamarine",
		"cornflowerblue",
		"crimson",
		"blueviolet",
	}
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func DotNodeLabel(n *DmtMakespanNode, m int) string {
	l := "\""
	for i, action := range n.actions {
		l += fmt.Sprintf("%v: %s %v\\l", action.Owner, action, n.costs[i])
	}
	l += fmt.Sprintf("H: %v\\l", n.h)
	l += fmt.Sprintf("M: %v\\l", m)

	return l + "\""
}

func DotNodeString(n *DmtMakespanNode, m int) string {
	// Attributes:
	//stateID  state.StateID
	//parent   *DmtMakespanNode
	//children []*DmtMakespanNode
	//actions  []*task.Action
	//costs    []int
	//h        int
	//visits   int
	//closed   bool

	if n != nil {
		agentID := 0
		penwidth := 1
		if n.Action() != nil {
			agentID = n.Action().Owner
			if !n.Action().Private {
				penwidth = 10
			}
		}
		var color string
		if n.closed || n.parent == nil {
			color = "red"
		} else {
			color = colors[agentID]
		}
		return fmt.Sprintf("n_%v [ "+
			"style=filled, "+
			"fillcolor=%s, "+
			"penwidth=%d, "+
			"label=%s"+
			" ];",
			int(n.stateID),
			color,
			penwidth,
			DotNodeLabel(n, m))
	}
	return ""
}

func DotEdgeString(n1, n2 *DmtMakespanNode) string {
	return fmt.Sprintf("n_%v->n_%v;", int(n1.stateID), int(n2.stateID))
}

func DotGraph(root *DmtMakespanNode, file string) {
	f, err := os.Create("./" + file)
	check(err)
	defer f.Close()

	f.WriteString("digraph G {\n")

	nodes := []*DmtMakespanNode{root}
	var n *DmtMakespanNode
	for len(nodes) > 0 {
		n, nodes = nodes[0], nodes[1:]
		f.WriteString(DotNodeString(n, n.m))
		for _, child := range n.children {
			f.WriteString(DotEdgeString(n, child))
			nodes = append(nodes, child)
		}
	}

	f.WriteString("}")
	f.Sync()
}

func DotPlan(goal *DmtMakespanNode, file string) {
	f, err := os.Create("./" + file)
	check(err)
	defer f.Close()

	f.WriteString("digraph G {\n")

	n := goal
	n2 := n.parent
	f.WriteString(DotNodeString(n, makespanApprox(n.pathToRoot())))
	for n2 != nil {
		f.WriteString(DotNodeString(n2, makespanApprox(n2.pathToRoot())))
		f.WriteString(DotEdgeString(n2, n))
		n = n2
		n2 = n2.parent
	}

	f.WriteString("}")
	f.Sync()
}
