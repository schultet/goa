package pddl

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type TokenList []string

var (
	requirements = []string{
		":strips", ":adl", ":typing", ":negation",
		":equality", ":negative-preconditions", ":disjunctive-preconditions",
		":existential-preconditions", ":universal-preconditions",
		":quantified-preconditions", ":conditional-effects",
		":derived-predicates", ":action-costs",
		":factored-privacy", ":open-world",
	}
	supportedRequirements = []string{
		":strips", ":typing", ":factored-privacy",
	}
)

func Open(domainFile, problemFile string) (*Domain, *Problem) {
	log.Printf("Parsing Domain (%s)...\n", domainFile)
	tokens, err := readTokens(domainFile)
	if err != nil {
		log.Fatalf("readTokens: %s (%s)", err, domainFile)
	}
	if tokens[0] != "(" || tokens[1] != "define" {
		log.Fatalf("PddlError in domain define line %v: %q\n", 0, tokens[0])
	}
	dom := ParseDomain(tokens[3:])
	log.Printf("Done Parsing Domain (%s) ...\n", domainFile)

	log.Printf("Parsing Problem (%s)...\n", problemFile)
	tokens, err = readTokens(problemFile)
	if err != nil {
		log.Fatalf("readTokens: %s (%s)", err, problemFile)
	}
	if tokens[0] != "(" || tokens[1] != "define" {
		log.Fatalf("PddlError in problem define line %v: %q\n", 0, tokens[0])
	}
	prob := ParseProblem(tokens[3:], dom)
	log.Printf("Done Parsing Problem (%s)...\n", problemFile)
	return &dom, &prob
}

// Domain represents a PDDL Domain
type Domain struct {
	name              string
	requirements      []string
	types             TypeMap
	constants         ObjectList
	privateConstants  ObjectList
	predicates        PredicateList
	privatePredicates PredicateList
	actions           []Action
	functions         []Function
}

func (d Domain) Actions() []Action            { return d.actions }
func (d Domain) Constants() []Object          { return []Object(d.constants) }
func (d Domain) PrivateConstants() ObjectList { return d.privateConstants }
func (d Domain) Types() TypeMap               { return d.types }

// Prints all data of a Domain.
func (d Domain) Dump() {
	fmt.Printf("Domain: %s\n", d.name)
	fmt.Printf("Requirements: %v\n", d.requirements)
	fmt.Print("Types: [")
	for _, t := range d.types {
		fmt.Printf("%s ", t.name)
	}
	fmt.Print("]\n")
	fmt.Print("Constants: [")
	for _, c := range d.constants {
		fmt.Printf("%s(%s) ", c.name, c.kind)
	}
	fmt.Print("]\n")
	fmt.Print("Private Constants: [")
	for _, c := range d.privateConstants {
		fmt.Printf("%s ", c.name)
	}
	fmt.Print("]\n")
	fmt.Print("Predicates: [")
	for _, p := range d.predicates {
		fmt.Printf("%s ", p.name)
	}
	fmt.Print("]\n")
	fmt.Print("Private predicates: [")
	for _, p := range d.privatePredicates {
		fmt.Printf("%s ", p.name)
	}
	fmt.Println("]")
	fmt.Println("Actions: ----------------")
	for _, a := range d.actions {
		fmt.Println(a)
	}
	fmt.Println("-------------------------")
}

// Reads in a PDDL _domain_ file and creates all data structures needed for
// representing the corresponding PDDL Domain.
func ParseDomain(tokens TokenList) Domain {
	domain := Domain{}
	for len(tokens) > 1 {
		switch tokens[1] {
		case "domain":
			domain.name, tokens = tokens[2], tokens[3:]
		case ":requirements":
			domain.requirements, tokens = parseRequirements(tokens[2:])
			if !requirementsSupported(domain.requirements, supportedRequirements) {
				log.Fatalf("ALARM: parser does not meet requirements")
			}
		case ":types":
			domain.types, tokens = ParseTypes(tokens[1:])
		case ":constants":
			domain.constants, domain.privateConstants, tokens =
				ParseConstants(tokens)
			tokens = tokens[1:]
		case ":predicates":
			var predicates, privatePredicates PredicateList
			predicates, privatePredicates, tokens = ParsePredicates(tokens[2:])
			domain.predicates = predicates
			domain.privatePredicates = privatePredicates
		case ":functions":
			var ftokens TokenList
			ftokens, tokens = extractParanthesis(tokens)
			domain.functions = ParseFunctions(ftokens)
		case ":action":
			var action Action
			action, tokens = ParseAction(tokens)
			domain.actions = append(domain.actions, action)
		default:
			tokens = tokens[1:]
		}
	}
	return domain
}

// Problem represents a PDDL Problem
type Problem struct {
	name                     string
	domain                   string
	objects                  ObjectList
	privateObjects           ObjectList
	initialPredicates        LiteralList     // TODO: rename to init.Literals?
	privateInitialPredicates LiteralList     // or Atoms/Fluents?
	numericFluents           []NumericFluent // TODO
	Goal                     FolFormula
}

func (p Problem) PrivateInitialPredicates() LiteralList { return p.privateInitialPredicates }
func (p Problem) InitialPredicates() LiteralList        { return p.initialPredicates }
func (p Problem) PrivateObjects() ObjectList            { return p.privateObjects }
func (p Problem) Objects() []Object                     { return []Object(p.objects) }

// Prints all data of a Problem.
func (p Problem) Dump() {
	fmt.Printf("Problem: %s\n", p.name)
	fmt.Printf("Domain: %s\n", p.domain)
	fmt.Print("Objects: [")
	for _, t := range p.objects {
		fmt.Printf("%s(%s) ", t.name, t.kind)
	}
	fmt.Println("]")
	fmt.Print("Private Objects: [")
	for _, t := range p.privateObjects {
		fmt.Printf("%s(%s) ", t.name, t.kind)
	}
	fmt.Println("]")
	fmt.Printf("Initial Predicates: [\n%s]\n", p.initialPredicates)
	fmt.Printf("Private Initial Predicates: [\n%s]\n", p.privateInitialPredicates)
	fmt.Printf("Goal: %s\n", p.Goal)
}

// Reads in a PDDL _problem_ file and creates all data structures needed for
// representing the corresponding PDDL Problem.
func ParseProblem(tokens TokenList, domain Domain) Problem {
	problem := Problem{}
	i := 0
	for i < len(tokens) {
		switch tokens[i] {
		case "problem":
			problem.name, tokens = tokens[i+1], tokens[i+2:]
			i = 0
		case ":domain":
			problem.domain, tokens = tokens[i+1], tokens[i+2:]
			i = 0
		case ":objects":
			var objectTokens TokenList
			objectTokens, tokens = extractParanthesis(tokens[i-1:])
			problem.objects, problem.privateObjects, _ =
				ParseConstants(objectTokens)
			i = 0
		case ":init":
			var initPreds, privInitPreds PredicateList
			initPreds, privInitPreds, tokens = ParsePredicates(tokens[i+1:])
			problem.initialPredicates = *new(LiteralList).initP(initPreds)
			problem.privateInitialPredicates = *new(LiteralList).initP(privInitPreds)
			i = 0
		case ":goal":
			var fol TokenList
			fol, tokens = extractParanthesis(tokens)
			problem.Goal = createFolFormula(fol[2:])
			i = 0
		default:
			i += 1
		}
	}
	return problem

}

// TODO: testing
// TODO: the object type must be determined
// find unambigious objects in list of grounded predicates (i.e. :init) and add
// them to objects
func extractObjects(predicates []*Predicate, objects *ObjectList,
	privateObjects, constants, privateConstants ObjectList) {
	for _, p := range predicates {
		if p == nil {
			log.Fatalf("predicate nil in Predicate list")
		}
		for _, v := range p.parameters {
			if isVar(v.name) {
				log.Fatalf("there should be no variable in initial predicates\n"+
					"Predicate: %s", p.Pddl())
			}
			if !objects.contains(v.name) && !constants.contains(v.name) &&
				!privateObjects.contains(v.name) && !privateConstants.contains(v.name) {
				*objects = append(*objects, *NewTypedObject(v.name, v.kind))
			}
		}
	}
}

// TODO: testing
// returns true when all requirements reqs of PDDL instance are supported by the
// planner (in supp)
func requirementsSupported(reqs, supp []string) bool {
	for _, req := range reqs {
		supported := false
		for _, sup := range supp {
			if req == sup {
				supported = true
				break
			}
		}
		if !supported {
			return false
		}
	}
	return true
}

// Splits a token list at a specified token and returns two token lists.
// The first one contains all tokens up to sep (not included) and the second one
// contains the rest including sep
func extractListUntil(tl TokenList, sep string) (TokenList, TokenList) {
	pos := tl.pos(sep)
	if pos == -1 {
		return tl, TokenList{}
	} else {
		return tl[0:pos], tl[pos:]
	}
}

// Extracts all tokens from a token list that are guarded by a pair of regular
// braces. Returns two token lists, the first one containing all tokens inside
// the braces (including braces), the second one containing the rest of the
// input list.
func extractParanthesis(tl TokenList) (TokenList, TokenList) {
	if len(tl) == 0 {
		return TokenList{}, TokenList{}
	}
	if tl[0] != "(" {
		log.Fatal(folError{"error parsing paranthesis term: not starting with '('"})
	}
	bracesCount := 1
	i := 1
	for ; i < len(tl); i++ {
		if tl[i] == ")" {
			bracesCount -= 1
		}
		if tl[i] == "(" {
			bracesCount += 1
		}
		if bracesCount == 0 {
			return tl[:i+1], tl[i+1:]
		}
	}
	return tl, TokenList{}
}

// Extracts PDDL Requirements from a token list. Returns the list of
// requirements and the remaining tokens of tl.
func parseRequirements(tl TokenList) ([]string, TokenList) {
	return extractListUntil(tl, ")")
}

// Extracts PDDL Types from a token list. Returns a TypeMap which maps typenames
// to *Types, and the remaining tokens of tl.
func ParseTypes(tl TokenList) (TypeMap, TokenList) {
	for i := 0; i < len(tl); i++ {
		if tl[i] == ":types" {
			tl = tl[i+1:]
			break
		}
	}
	typelst, tl := extractListUntil(tl, ")")
	typemap := TypeMap{"object": &Type{"object", &Type{}, []*Type{}}}

	subtypes := []*Type{}
	for j := 0; j < len(typelst); j++ {
		if typelst[j] == "-" {
			j += 1
			basetype, ok := typemap[typelst[j]]
			if !ok {
				basetype = &Type{typelst[j], &Type{}, []*Type{}}
				typemap[basetype.name] = basetype
			}
			for _, t := range subtypes {
				t.parent = basetype
				basetype.children = append(basetype.children, t)
				typemap[t.name] = t
			}
			subtypes = nil
		} else {
			subtype, ok := typemap[typelst[j]]
			if !ok {
				subtype = &Type{typelst[j], typemap["object"], []*Type{}}
			}
			subtypes = append(subtypes, subtype)
		}
	}
	for _, t := range subtypes {
		typemap[t.name] = t
		t.parent.children = append(t.parent.children, t)
	}
	return typemap, tl[1:]
}

// parses a typed list as defined by PDDL (in tokenized form) and returns a map
// from typename to array of objects/variables of that type
func parseTypedList(tl TokenList) map[string][]string {
	p := 0
	result := map[string][]string{}
	for {
		p = tl.pos("-")
		if p == -1 {
			if lst, ok := result["object"]; ok {
				result["object"] = append(lst, tl...)
			} else {
				result["object"] = append([]string{}, tl...)
			}
			break
		}
		if ofType, ok := result[tl[p+1]]; ok {
			result[tl[p+1]] = append(ofType, tl[:p]...)
		} else {
			result[tl[p+1]] = append([]string{}, tl[:p]...)
		}
		tl = tl[p+2:]
	}
	return result
}

func ParseFunctions(tl TokenList) (fs []Function) {
	if tl[0] != "(" || tl[1] != ":functions" {
		log.Fatalf("Error parsing functions\nTokenList: %s\n", tl)
	}
	tl = tl[2:]
	for len(tl) > 1 {
		var f TokenList
		f, tl = extractParanthesis(tl)
		f, tl = append(f, tl[:2]...), tl[2:]
		fs = append(fs, *new(Function).Init(f))
	}
	return fs
}

// parses PDDL constants
func ParseConstants(tl TokenList) (ObjectList, ObjectList, TokenList) {
	var tokens TokenList
	tokens, tl = extractListUntil(tl[2:], ")") // skip ':constants' and '('

	constants := ObjectList{}
	privateConstants := ObjectList{}

	extractTypedObjects := func(tm map[string][]string) ObjectList {
		constants := ObjectList{}
		for k, v := range tm {
			for _, c := range v {
				constants = append(constants, *NewTypedObject(c, k))
			}
		}
		return constants
	}

	p := tokens.pos(":private")
	if p == -1 {
		constants = extractTypedObjects(parseTypedList(tokens))
	} else {
		constants = extractTypedObjects(parseTypedList(tokens[:p-1]))
		privateConstants = extractTypedObjects(parseTypedList(tokens[p+1:]))
		for i := range privateConstants {
			privateConstants[i].private = true
		}
	}

	return constants, privateConstants, tl
}

// from a token list typed/untyped variables are extracted and returned as array
func extractParameters(tl TokenList) []Variable {
	i := 0
	if tl[i] == "(" {
		i += 1
	}

	parameters := []string{}
	result := []Variable{}

	for ; i < len(tl); i++ {
		if token := tl[i]; token == ")" {
			break
		} else if token == "-" {
			i += 1
			for _, p := range parameters {
				variable := Variable{name: p, kind: tl[i]}
				result = append(result, variable)
			}
			parameters = nil
		} else {
			parameters = append(parameters, tl[i])
		}
	}
	for _, p := range parameters {
		result = append(result, Variable{name: p})
	}
	return result
}

// parses PDDL predicates from a token list. Returns array of public predicates,
// array of private predicates, and the remaining elements of the token list
func ParsePredicates(tl TokenList) (PredicateList, PredicateList, TokenList) {
	predicates := make(PredicateList, 0)
	privatePredicates := make(PredicateList, 0)
	var predicate_lst []string
	for len(tl) > 0 { // handle non-private predicates
		if tl[0] == "(" {
			if tl[1] == ":private" {
				tl = tl[2:] // skip ':private' and '('
				break       // TODO: error checking
			}
			if tl[1] == "=" { // handle numeric fluents
				_, tl = extractParanthesis(tl) // TODO: currently ignores nf's
				continue
			}
			predicate_lst, tl = extractListUntil(tl[1:], ")")
			tl = tl[1:]
			predicates = append(predicates, *NewPredicate(predicate_lst))
		} else if tl[0] == ")" {
			return predicates, privatePredicates, tl[1:]
		}
	}
	for len(tl) > 0 { // handle _private_ predicates
		if tl[0] == "(" {
			predicate_lst, tl = extractListUntil(tl[1:], ")")
			tl = tl[1:]
			pp := NewPredicate(predicate_lst)
			pp.private = true
			privatePredicates = append(privatePredicates, *pp)
		} else if tl[0] == ")" {
			return predicates, privatePredicates, tl[1:]
		}
	}
	return predicates, privatePredicates, tl[1:]
}

// Creates a list of tokens from a string
func Tokenize(s string) TokenList {
	s = strings.Replace(s, "(", " ( ", -1)
	s = strings.Replace(s, ")", " ) ", -1)
	s = strings.Replace(s, "?", " ?", -1)
	return strings.Fields(s)
}

// Reads in a (PDDL-)File and returns a list of tokens (strings).
func readTokens(filename string) (TokenList, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tokens TokenList
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		tokens = append(tokens, Tokenize(text)...)
	}
	return tokens, scanner.Err()
}

// Reads an arbitrary file and returns a []byte
func readFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Bytes()...)
	}
	return result, scanner.Err()
}

// Returns the index of a token in the TokenList, or -1 if the token is not in
// the list.
func (tl TokenList) pos(v string) int {
	for i, t := range tl {
		if t == v {
			return i
		}
	}
	return -1
}

// Returns a string representation of a TokenList.
func (tl TokenList) String() string { return strings.Join(tl, " ") }

// DumpPDDL prints out domain and problem attributes
func DumpPDDL(dom Domain, prob Problem) {
	dom.Dump()
	prob.Dump()
	fmt.Println("\ndone.\n")
}
