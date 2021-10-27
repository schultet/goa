package search

import (
	"fmt"
	"log"

	"github.com/schultet/goa/pkg/opt"
)

var (
	strategyRegistry []*StrategyInfo
)

// Status represents a search strategy state
type Status int

const (
	inProgress Status = iota // search in progress
	timeout                  // timeout occured
	failed                   // search failed, goal not reachable
	solved                   // goal state found
	idle                     // nothing to do, e.g. when openlist is empty
)

// Strategy is the basic search strategy interface
type Strategy interface {
	// Initialize is called once before the search starts and contains all the
	// required initialization logic.
	Initialize()

	// Step is continiously called in the search.Engine's main loop. It executes
	// a single search step and returns the new search status.
	Step() Status
}

// StrategyInfo holds information necessary for strategy creation based on command
// line arguments. For each search strategy selectable via cmd arguments, an
// according StrategyInfo object must exist in the (global) strategyRegistry
type StrategyInfo struct {
	Name        string
	Description string
	NewStrategy func(*Engine, *opt.OptionSet) Strategy
	Options     *opt.OptionSet
}

// RegisterStrategy a new StrategyInfo object to the strategyRegistry. This step is
// required to select the underlying strategy from the command line.
func RegisterStrategy(info *StrategyInfo) {
	strategyRegistry = append(strategyRegistry, info)
}

// GetStrategy returns the strategyinfo registered with the provided name (= name of
// the strategy, as used in the command line argument)
func GetStrategy(name string) (*StrategyInfo, error) {
	for _, item := range strategyRegistry {
		if item.Name == name {
			return item, nil
		}
	}
	return nil, fmt.Errorf("strategy <%s> not found", name)
}

// ParseStrategy retrieves the specified strategy from the strategyRegistry (if
// the respective strategy was registered before), parses its command line
// options, then builds and returns the specified Strategy
func NewStrategy(engine *Engine, args []string) Strategy {
	arguments := append([]string{}, args...) // copy argument list
	// TODO: move general search options somewhere appropriate
	strategyOpts := opt.NewOptionSet()
	strategyOpts.Add(opt.Option{
		Type:         opt.String,
		Name:         "strategy",
		Short:        's',
		DefaultValue: "dmt-bfs",
		Description:  "search strategy"})
	strategyOpts.Parse(arguments)
	strategy := strategyOpts.GetString("strategy")

	strategyInfo, err := GetStrategy(strategy)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	opts := strategyInfo.Options
	opts.Parse(arguments)
	return strategyInfo.NewStrategy(engine, opts)
}
