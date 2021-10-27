package search

//func printNodes(nodes []*DmtNode) {
//var par, left, right int
//for _, n := range nodes {
//	if n.parent == nil {
//		par = -1
//	} else {
//		par = int(n.parent.id)
//	}
//	if n.children[0] == nil {
//		left = -1
//	} else {
//		left = int(n.children[0].id)
//	}
//	if n.children[1] == nil {
//		right = -1
//	} else {
//		right = int(n.children[1].id)
//	}
//	fmt.Printf("(id:%d, par:%d, left:%d, right%d)\n",
//		n.id, par, left, right)
//}
//}

//func binaryTree(ids []int) []*DmtNode {
//var insert func(x, y *DmtNode)
//insert = func(root, node *DmtNode) {
//	var n *DmtNode
//	if node.id < root.id { // insert left
//		n = root.children[0]
//		if n == nil {
//			root.children[0] = node
//			node.parent = root
//			return
//		}
//	} else {
//		n = root.children[1]
//		if n == nil {
//			root.children[1] = node
//			node.parent = root
//			return
//		}
//	}
//	insert(n, node)
//}
//nodes := make([]*DmtNode, 0, len(ids))
//for i, id := range ids {
//	n := &DmtNode{
//		id:       NodeID(id),
//		children: make([]*DmtNode, 2),
//	}
//	nodes = append(nodes, n)
//	if i > 0 {
//		insert(nodes[0], n)
//	}
//}
//return nodes
//}

//func dfsOrderedIDs(root *DmtNode) []int {
//if root == nil {
//	return []int{}
//}
//res := []int{int(root.id)}
//for i := range root.children {
//	res = append(res, dfsOrderedIDs(root.children[i])...)
//}
//return res
//}

//func TestAddChild(t *testing.T) {
// TODO: write a real testcase
//nodes := make([]*DmtNode, 0, 10)
//for i := 0; i < 10; i++ {
//	nodes = append(nodes, &DmtNode{id: NodeID(i)})
//}
//nodes[0].addChild(nodes[1])
//nodes[0].addChild(nodes[2])
//nodes[0].addChild(nodes[3])

//nodes[3].addChild(nodes[5])
//nodes[3].addChild(nodes[4])
//nodes[3].addChild(nodes[2])

//lst := dfsOrderedIDs(nodes[0])
//fmt.Printf("lst = %+v\n", lst)
//}

//func TestForbearOf(t *testing.T) {
//	nodes := binaryTree([]int{10, 5, 1, 6, 20, 15, 22, 34})
//	//........................ 0  1  2  3   4   5   6   7
//	type test struct {
//		nodeIndex1     int
//		nodeIndex2     int
//		expectedResult bool
//	}
//	tests := []test{
//		{0, 0, false}, {1, 1, false}, {2, 2, false}, {7, 7, false},
//		{0, 1, true}, {0, 2, true}, {0, 3, true}, {0, 4, true}, {0, 7, true},
//		{1, 0, false}, {2, 0, false}, {3, 0, false}, {4, 0, false},
//		{7, 0, false}, {6, 7, true}, {7, 6, false},
//		{1, 7, false},
//		{2, 3, false}, {3, 2, false}, {1, 4, false}, {4, 1, false},
//		{5, 6, false}, {6, 5, false},
//	}
//	for _, tc := range tests {
//		node1, node2 := nodes[tc.nodeIndex1], nodes[tc.nodeIndex2]
//		res := node1.forbearOf(node2)
//		if res != tc.expectedResult {
//			t.Errorf("\nexp: %v\nwas: %v\nnode1:%d forbearOf node2:%d\n",
//				tc.expectedResult, res, tc.nodeIndex1, tc.nodeIndex2)
//		}
//	}
//}

//func TestDmtNodeStateID(t *testing.T) {
//	type test struct {
//		stateID int
//	}
//	tests := []test{test{35}, test{0}, test{1}, test{100}, test{2000}}
//	for _, tc := range tests {
//		n := NewDmtNode(state.StateID(tc.stateID), 0, nil, nil, 0, 0, 0, 35)
//		if int(n.sid) != tc.stateID {
//			t.Errorf("ERROR: bla, %d %d\n", n.sid, tc.stateID)
//		}
//	}
//}

// TODO: make a real testcase of this
//func TestPerAgentCost(t *testing.T) {
//	tokens0 := state.NewTokenArray([]int32{0, 0, 0})
//	tokens1 := state.NewTokenArray([]int32{1, 2, 0})
//	tokens2 := state.NewTokenArray([]int32{35, 2, 0})
//
//	dmt := &Dmt{
//		Engine: &Engine{
//			agentID:       0,
//			tokenRegistry: state.NewTokenRegistry()},
//	}
//
//	root := &DmtNode{
//		tid: dmt.tokenRegistry.Register(tokens0),
//	}
//	n1 := &DmtNode{
//		parent: root,
//		tid:    dmt.tokenRegistry.Register(tokens1),
//		cost:   4,
//	}
//	n2 := &DmtNode{
//		parent: n1,
//		tid:    dmt.tokenRegistry.Register(tokens2),
//		cost:   2,
//	}
//	n3 := &DmtNode{
//		id:     35,
//		parent: n2,
//		cost:   3,
//		tid:    dmt.tokenRegistry.Register(tokens2),
//	}
//	interestedAgents := []int{1, 2}
//	fmt.Printf("%v", dmt.Tokens(n3))
//	res := dmt.perAgentCost(n3, interestedAgents)
//	fmt.Printf("%v", res)
//	//for agent := range res {
//	//	fmt.Printf("%d: %d\n", agent, res[agent])
//	//}
//}
