package comm

import "testing"

func TestConnListGet(t *testing.T) {
	al := ConnList{
		Conn{AgentID: 0, Port: 3035, Host: "127.0.0.1"},
		Conn{AgentID: 35, Port: 3036, Host: "127.0.0.1"},
		Conn{AgentID: 2, Port: 3037, Host: "127.0.0.1"},
		Conn{AgentID: 3, Port: 3038, Host: "127.0.0.1"},
	}
	queries := []int{0, 35, 2, 3}
	for i, q := range queries {
		agent, _ := al.Get(q)
		if al[i] != agent {
			t.Errorf("AgentList.Get returns wrong agent!\nExp:%v\nWas:%v\n", al[i], agent)
		}
	}
	agent, err := al.Get(70)

	if err == nil || (agent != Conn{}) {
		t.Errorf("AgentList.Get returns wrong AgentInfo! Error expected!\n")
	}
}

func TestConnListExcept(t *testing.T) {
	al1 := ConnList{}
	al2 := ConnList{Conn{35, 0, "hstnm"}}
	al3 := ConnList{Conn{35, 0, "hstnm"}, Conn{70, 3, ""}}
	al4 := ConnList{Conn{35, 0, "hstnm"}, Conn{70, 3, ""}, Conn{101, 45, "test"}}
	all := []ConnList{al1, al2, al3, al4}

	// test to remove element not in list
	for _, al := range all {
		result := al.Except(200)
		if !al.Equals(&result) {
			t.Errorf("AgentList.Except returns wrong list!")
		}
	}

	res := al2.Except(35)
	exp := ConnList{}
	if !res.Equals(&exp) {
		t.Errorf("sth went wrong")
	}

	res = al4.Except(70)
	exp = ConnList{Conn{35, 0, "hstnm"}, Conn{101, 45, "test"}}
	if !res.Equals(&exp) {
		t.Errorf("sth went wrong")
	}
}
