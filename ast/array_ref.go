package ast

import "strings"

type Array_ref struct {
	Type string
	Op0  string
	Op1  string
}

func (a Array_ref) GenNodeName() string {
	return "Array_ref"
}

func parse_array_ref(line string) (n Node) {
	groups := groupsFromRegex(
		`
	type:(?P<type>.*) +
	op 0:(?P<op0>.*) +
	op 1:(?P<op1>.*) +
	`,
		line,
	)
	return Array_ref{
		Type: strings.TrimSpace(groups["type"]),
		Op0:  strings.TrimSpace(groups["op0"]),
		Op1:  strings.TrimSpace(groups["op1"]),
	}
}
