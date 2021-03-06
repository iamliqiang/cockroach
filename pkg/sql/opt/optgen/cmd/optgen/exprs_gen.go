// Copyright 2018 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package main

import (
	"fmt"
	"io"

	"github.com/cockroachdb/cockroach/pkg/sql/opt/optgen/lang"
)

// exprsGen generates the memo expression structs used by the optimizer, as
// well as lookup tables used to implement the ExprView methods.
type exprsGen struct {
	compiled *lang.CompiledExpr
	w        io.Writer
}

func (g *exprsGen) generate(compiled *lang.CompiledExpr, w io.Writer) {
	g.compiled = compiled
	g.w = w

	fmt.Fprintf(g.w, "package xform\n\n")

	fmt.Fprintf(g.w, "import (\n")
	fmt.Fprintf(g.w, "  \"github.com/cockroachdb/cockroach/pkg/sql/opt\"\n")
	fmt.Fprintf(g.w, ")\n\n")

	g.genLayoutTable()
	g.genTagLookup()
	g.genIsTag()

	for _, define := range g.compiled.Defines {
		// Skip enforcers, since they are not memoized.
		if define.Tags.Contains("Enforcer") {
			continue
		}

		g.genExprType(define)
		g.genExprFuncs(define)
		g.genMemoFuncs(define)
	}
}

// genLayoutTable generates the layout table; see opLayout.
func (g *exprsGen) genLayoutTable() {
	fmt.Fprintf(g.w, "var opLayoutTable = [...]opLayout{\n")
	fmt.Fprintf(g.w, "  opt.UnknownOp: 0xFF, // will cause a crash if used\n")
	for _, define := range g.compiled.Defines {
		var count, listVal, privVal, enfVal int

		count = len(define.Fields)
		if privateField(define) != nil {
			privVal = count
			count--
		}
		list := listField(define)
		if list != nil {
			listVal = count
			if privVal != 0 {
				// The list takes two slots; adjust the private position.
				privVal++
			}
			count--
		}
		if define.Tags.Contains("Enforcer") {
			enfVal = 1
		}
		fmt.Fprintf(
			g.w, "  opt.%sOp: makeOpLayout(%d /*base*/, %d /*list*/, %d /*priv*/, %d /*enforcer*/),\n",
			define.Name, count, listVal, privVal, enfVal,
		)
	}
	fmt.Fprintf(g.w, "}\n\n")
}

// genTagLookup generates a lookup table used to implement the ExprView IsXXX
// methods for each different define tag. These methods indicate whether the
// expression is associated with that particular tag.
func (g *exprsGen) genTagLookup() {
	for _, tag := range g.compiled.DefineTags {
		if tag == "Custom" {
			// Don't create method, since this is compiler directive.
			continue
		}

		fmt.Fprintf(g.w, "var is%sLookup = [...]bool{\n", tag)
		fmt.Fprintf(g.w, "  false, // UnknownOp\n\n")

		for _, define := range g.compiled.Defines {
			fmt.Fprintf(g.w, "  %v, // %sOp\n", define.Tags.Contains(tag), define.Name)
		}

		fmt.Fprintf(g.w, "}\n\n")
	}
}

// genIsTag generates IsXXX tag methods on ExprView and memoExpr for every
// unique tag.
func (g *exprsGen) genIsTag() {
	for _, tag := range g.compiled.DefineTags {
		fmt.Fprintf(g.w, "func (ev ExprView) Is%s() bool {\n", tag)
		fmt.Fprintf(g.w, "  return is%sLookup[ev.op]\n", tag)
		fmt.Fprintf(g.w, "}\n\n")
	}

	for _, tag := range g.compiled.DefineTags {
		fmt.Fprintf(g.w, "func (me *memoExpr) is%s() bool {\n", tag)
		fmt.Fprintf(g.w, "  return is%sLookup[me.op]\n", tag)
		fmt.Fprintf(g.w, "}\n\n")
	}
}

// genExprType generates the type definition for the expression, as well as a
// constructor function.
func (g *exprsGen) genExprType(define *lang.DefineExpr) {
	opType := fmt.Sprintf("%sOp", define.Name)
	exprType := fmt.Sprintf("%sExpr", unTitle(string(define.Name)))

	// Generate comment for the expression type.
	generateDefineComments(g.w, define, exprType)

	// Generate the expression type.
	fmt.Fprintf(g.w, "type %s memoExpr\n\n", exprType)

	// Generate a strongly-typed constructor function for the type.
	fmt.Fprintf(g.w, "func make%sExpr(", define.Name)
	for i, field := range define.Fields {
		if i != 0 {
			fmt.Fprint(g.w, ", ")
		}
		fmt.Fprintf(g.w, "%s opt.%s", unTitle(string(field.Name)), mapType(string(field.Type)))
	}
	fmt.Fprintf(g.w, ") %s {\n", exprType)
	fmt.Fprintf(g.w, "  return %s{op: opt.%s, state: exprState{", exprType, opType)

	for i, field := range define.Fields {
		fieldName := unTitle(string(field.Name))

		if i != 0 {
			fmt.Fprintf(g.w, ", ")
		}

		if isListType(string(field.Type)) {
			fmt.Fprintf(g.w, "%s.Offset, %s.Length", fieldName, fieldName)
		} else {
			fmt.Fprintf(g.w, "uint32(%s)", fieldName)
		}
	}

	fmt.Fprint(g.w, "}}\n")
	fmt.Fprint(g.w, "}\n\n")
}

// genExprFuncs generates the expression's accessor functions, one for each
// field in the type.
func (g *exprsGen) genExprFuncs(define *lang.DefineExpr) {
	exprType := fmt.Sprintf("%sExpr", unTitle(string(define.Name)))

	// Generate the strongly-typed accessor methods.
	stateIndex := 0
	for _, field := range define.Fields {
		fieldName := unTitle(string(field.Name))
		fieldType := mapType(string(field.Type))

		fmt.Fprintf(g.w, "func (e *%s) %s() opt.%s {\n", exprType, fieldName, fieldType)
		if isListType(string(field.Type)) {
			format := "  return opt.ListID{Offset: e.state[%d], Length: e.state[%d]}\n"
			fmt.Fprintf(g.w, format, stateIndex, stateIndex+1)
			stateIndex += 2
		} else if isPrivateType(string(field.Type)) {
			fmt.Fprintf(g.w, "  return opt.PrivateID(e.state[%d])\n", stateIndex)
			stateIndex++
		} else {
			fmt.Fprintf(g.w, "  return opt.GroupID(e.state[%d])\n", stateIndex)
			stateIndex++
		}
		fmt.Fprintf(g.w, "}\n\n")
	}

	// Generate the fingerprint method.
	fmt.Fprintf(g.w, "func (e *%s) fingerprint() fingerprint {\n", exprType)
	fmt.Fprintf(g.w, "  return fingerprint(*e)\n")
	fmt.Fprintf(g.w, "}\n\n")
}

// genMemoFuncs generates conversion methods on the memo expression, one for
// each more specialized expression type.
func (g *exprsGen) genMemoFuncs(define *lang.DefineExpr) {
	opType := fmt.Sprintf("%sOp", define.Name)
	exprType := fmt.Sprintf("%sExpr", unTitle(string(define.Name)))

	// Generate a conversion method from memoExpr to the more specialized
	// expression type.
	fmt.Fprintf(g.w, "func (m *memoExpr) as%s() *%s {\n", define.Name, exprType)
	fmt.Fprintf(g.w, "  if m.op != opt.%s {\n", opType)
	fmt.Fprintf(g.w, "    return nil\n")
	fmt.Fprintf(g.w, "  }\n")

	fmt.Fprintf(g.w, "  return (*%s)(m)\n", exprType)
	fmt.Fprintf(g.w, "}\n\n")
}
