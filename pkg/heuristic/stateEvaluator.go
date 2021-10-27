package heuristic

import (
	"errors"
	"log"
	"strconv"
	"sync"

	"github.com/schultet/goa/pkg/state"
	"github.com/schultet/goa/pkg/task"
)

const (
	MaxCost = 100000000
	DeadEnd = -1
)

var (
	evalRegistry []*EvaluatorInfo
	evaluators   map[string]StateEvaluator = make(map[string]StateEvaluator, 0)
	mutex        *sync.Mutex               = &sync.Mutex{}
)

type StateEvaluator interface {
	Evaluate(s state.State) int
}

type EvaluatorInfo struct {
	Name         string
	Description  string
	NewEvaluator func(task *task.Task) StateEvaluator
}

func Register(info *EvaluatorInfo) {
	evalRegistry = append(evalRegistry, info)
}

func Find(name string) (*EvaluatorInfo, error) {
	for _, item := range evalRegistry {
		if item.Name == name {
			return item, nil
		}
	}
	return nil, errors.New("evaluator" + name + "not found!\n")
}

func Get(name string, agentID int) (StateEvaluator, error) {
	eval, ok := evaluators[name+strconv.Itoa(agentID)]
	if !ok {
		return nil, errors.New("state evaluator not found!\n")
	}
	return eval, nil
}

func Initialize(name string, t *task.Task) {
	info, err := Find(name)
	if err != nil {
		log.Fatalf("[%v] ERR: failed to initialize heuristic. Err: %+v\n%+v\n",
			t.AgentID, err, evaluators)
	}
	mutex.Lock()
	evaluators[name+strconv.Itoa(t.AgentID)] = info.NewEvaluator(t)
	mutex.Unlock()
	log.Printf("HEUR: initialized %s heuristic: %+v\n",
		name+strconv.Itoa(t.AgentID), evaluators[name+strconv.Itoa(t.AgentID)])
}
