// Package opt is an adapter to the getopt-package `github.com/pborman/getopt`
// providing `getopt` like processing of command line arguments.
package opt

import (
	"os"
	"strconv"
	"strings"

	"github.com/pborman/getopt"
)

// OptionType allows to distinguish between different cmd option types
type OptionType byte

const (
	// Bool is an option type for cmd arguments that can be either true or false
	Bool OptionType = iota
	// String is the option type for cmd arguments holding arbitrary strings
	String
	// Int32 is the option type for cmd arguments holding int32 values
	Int32
	// List is the option type for cmd arguments that are lists of strings
	List
	// Float64 is the option type for cmd arguments holding float64 values
	Float64
)

// Option specifies a command line option (argument). Each option has a type, a
// name, a short representation of the name, a default value, and a description.
type Option struct {
	Type         OptionType
	Name         string
	Short        rune
	DefaultValue interface{}
	Description  string
}

// NewOption creates a new Option
func NewOption(t OptionType, name string, short rune, defval interface{},
	desc string) Option {
	return Option{
		Type:         t,
		Name:         name,
		Short:        short,
		DefaultValue: defval,
		Description:  desc,
	}
}

// OptionSet is a container type holding multiple Options. It provides methods
// to parse cmd arguments based on the contained options, and methods to access
// the respective values afterwards.
type OptionSet getopt.Set

// NewOptionSet creates a new OptionSet
func NewOptionSet() *OptionSet {
	optionset := OptionSet(*getopt.New())
	return &optionset
}

// Parse parses a list of command line arguments according to an OptionSet
func (s *OptionSet) Parse(args []string) {
	(*getopt.Set)(s).Getopt(args, nil)
	for (*getopt.Set)(s).NArgs() > 0 {
		args = args[1:]
		(*getopt.Set)(s).Getopt(args, nil)
	}
}

// ParseUnforgiving parses a list of command line arguments according to an
// OptionSet. When an error occurs during option parsing, os.Exit(1) is called.
func (s *OptionSet) ParseUnforgiving(args []string) {
	err := (*getopt.Set)(s).Getopt(args, nil)
	if err != nil {
		(*getopt.Set)(s).PrintUsage(os.Stderr)
		os.Exit(1)
		//fmt.Fprintln(os.Stderr, err)
	}
}

// GetBool returns the (bool) value of an option given its `name`
func (s *OptionSet) GetBool(name string) bool {
	b, _ := strconv.ParseBool((*getopt.Set)(s).Lookup(name).String())
	return b
}

// GetInt32 returns the (int32) value of an option given its `name`
func (s *OptionSet) GetInt32(name string) int32 {
	i64, _ := strconv.ParseInt((*getopt.Set)(s).Lookup(name).String(), 10, 32)
	return int32(i64)
}

// GetFloat64 returns the (float64) value of an option given its `name`
func (s *OptionSet) GetFloat64(name string) float64 {
	c, _ := strconv.ParseFloat((*getopt.Set)(s).Lookup(name).String(), 64)
	return c
}

// GetString returns the (string) value of an option given its `name`
func (s *OptionSet) GetString(name string) string {
	return (*getopt.Set)(s).Lookup(name).String()
}

// GetList returns the ([]string) value of an option given its `name`
func (s *OptionSet) GetList(name string) []string {
	o := (*getopt.Set)(s).Lookup(name).Value().String()
	return strings.FieldsFunc(o, func(r rune) bool { return r == ',' })
}

// Add adds option `o` to OptionSet `s`. Panics if OptionType is not supported.
func (s *OptionSet) Add(o Option) {
	switch optiontype := o.Type; optiontype {
	case Bool:
		(*getopt.Set)(s).BoolLong(o.Name, o.Short, o.Description)
	case String:
		(*getopt.Set)(s).StringLong(o.Name, o.Short, o.DefaultValue.(string), o.Description)
	case Int32:
		(*getopt.Set)(s).Int32Long(o.Name, o.Short, int32(o.DefaultValue.(int)), o.Description)
	case List:
		(*getopt.Set)(s).ListLong(o.Name, o.Short, o.DefaultValue.(string), o.Description)
	case Float64:
		(*getopt.Set)(s).StringLong(o.Name, o.Short, o.DefaultValue.(string), o.Description)
	default:
		// TODO: this should never happen, raise an error instead
		panic("waaah")
	}
}
