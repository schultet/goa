package search

import (
	"fmt"
	"log"
	"time"

	"github.com/schultet/goa/src/comm"
	"github.com/schultet/goa/src/util/ints"
)

// ProgressInfo collects information on search progress
type ProgressInfo struct {
	agentID      int
	plan         string
	steps        string
	smallestH    int
	currentH     int
	expansions   int
	messagesOut  int
	messagesIn   int
	t            time.Time
	lastPrint    time.Time
	planCost     int
	planLength   int
	planMakespan int
	agentCount   int
	logger       *log.Logger
}

// NewProgressInfo creates a new ProgressInfo object
func NewProgressInfo(agentID int, agents comm.ConnList, l *log.Logger) *ProgressInfo {
	return &ProgressInfo{
		agentID:      agentID,
		agentCount:   len(agents),
		planCost:     -1,
		planLength:   -1,
		planMakespan: -1,
		smallestH:    ints.MaxValue,
		t:            time.Now(),
		logger:       l,
	}
}

// Print prints the current progress to stdout
func (i *ProgressInfo) Print() {
	i.logger.Printf("[t:%5.2f, h:%3d, exp:%d, m->:%d, m<-:%d, ratio:%5.2f]\n",
		time.Since(i.t).Minutes(), i.smallestH, i.expansions,
		i.messagesOut, i.messagesIn,
		float64(i.messagesOut)/float64(i.messagesIn))
	i.lastPrint = time.Now()
}

// PrintCurrent prints the current progress to stdout (same as Print) but is
// used to report progress after a given duration
func (i *ProgressInfo) PrintCurrent() {
	i.logger.Printf("(t:%5.2f, cur_h:%3d, exp:%d, m->:%d, m<-:%d, ratio:%5.2f)\n",
		time.Since(i.t).Minutes(), i.currentH, i.expansions,
		i.messagesOut, i.messagesIn,
		float64(i.messagesOut)/float64(i.messagesIn))
	i.lastPrint = time.Now()
}

// PrintTime prints the time (in minutes) since the object was created
func (i *ProgressInfo) PrintTime() {
	i.logger.Printf("(T:%5.2f\n", time.Since(i.t).Minutes())
}

// SetPlan that this progress information is about.
func (i *ProgressInfo) SetPlan(plan *PlanExtractor) {
	i.plan = plan.descriptor
	i.planCost = plan.totalCost
	i.planMakespan = plan.totalMakespan
	i.steps = "{\n"
	for s, step := range plan.plan {
		if step.Action != nil {
			// TODO: Do proper marshaling.
			i.steps += fmt.Sprintf("    \"%v\": \"%v\"", step.T, step.Action)
			if s < len(plan.plan)-1 {
				i.steps += ",\n"
			}
		}
	}
	i.steps += "\n  }"
}

// SummaryJSON summarizes the progress as JSON string.
func (i *ProgressInfo) SummaryJSON() string {
	summary := fmt.Sprintln("{")
	summary += fmt.Sprintf("  \"AgentID\": %v,\n", i.agentID)
	summary += fmt.Sprintf("  \"Plan\": \"%s\",\n", i.plan)
	summary += fmt.Sprintf("  \"Steps\": %v,\n", i.steps)
	summary += fmt.Sprintf("  \"Seconds\": %v,\n", time.Since(i.t).Seconds())
	summary += fmt.Sprintf("  \"Makespan\": %v,\n", i.planMakespan)
	summary += fmt.Sprintf("  \"Cost\": %v,\n", i.planCost)
	summary += fmt.Sprintf("  \"Expansions\": %v,\n", i.expansions)
	summary += fmt.Sprintf("  \"MessagesIn\": %v,\n", i.messagesIn)
	summary += fmt.Sprintf("  \"MessagesOut\": %v,\n", i.messagesOut)
	summary += fmt.Sprintf("  \"HeuristicValue\": %v,\n", i.currentH)
	summary += fmt.Sprintf("  \"AgentCount\": %v\n", i.agentCount)
	return fmt.Sprintf("%s}\n", summary)
}

func (i *ProgressInfo) incExpansions(d int)  { i.expansions += d }
func (i *ProgressInfo) incMessagesOut(d int) { i.messagesOut += d }
func (i *ProgressInfo) incMessagesIn(d int)  { i.messagesIn += d }

func (i *ProgressInfo) setHeuristicValue(h int) {
	if h < i.smallestH {
		i.smallestH = h
		i.Print()
	}
	i.currentH = h
}
