// Copyright(c) 2022 individual contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// <https://www.apache.org/licenses/LICENSE-2.0>
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package main

import (
	"bytes"
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"go/format"
	"io"
	"os"
	"strings"
	"text/template"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type Trait struct{ Name, Interface string }

type Type struct {
	Name       string
	Atom       int
	Parents    []string
	Interfaces []string
}

func (t *Type) HasParents() bool {
	return len(t.Parents) > 1
}

type Source struct {
	PackageName string
	TypeName    string
	TypeNameL   string

	Traits []Trait

	Types       map[string]*Type
	OrderByAtom []*Type
	OrderByName []*Type
}

// fillTypes creates all type permutations of the available traits.
func (src *Source) fillTypes() {
	slices.SortFunc(src.Traits, func(l, r Trait) bool {
		return l.Name < r.Name
	})

	src.Types = make(map[string]*Type)
	for i := 1; i < 1<<len(src.Traits); i++ {
		st := Type{Atom: i}

		for j, t := range src.Traits {
			if (i & (1 << j)) == 0 {
				continue
			}
			st.Name += t.Name
			st.Parents = append(st.Parents, t.Name)
			st.Interfaces = append(st.Interfaces, t.Interface)
		}

		// Compress structures parents.
		if len(st.Parents) == 1 && st.Parents[0] == st.Name {
			st.Parents = nil
		} else if len(st.Parents) > 2 {
			compParent := strings.Join(st.Parents[:len(st.Parents)-1], "")
			if _, ok := src.Types[compParent]; ok {
				st.Parents = []string{compParent, st.Parents[len(st.Parents)-1]}
			}
		}

		src.Types[st.Name] = &st
	}

	src.OrderByName = maps.Values(src.Types)
	slices.SortFunc(src.OrderByName, func(l, r *Type) bool {
		if l1, l2 := len(l.Name), len(r.Name); l1 != l2 {
			return l1 < l2
		}
		return l.Name < r.Name
	})

	src.OrderByAtom = maps.Values(src.Types)
	slices.SortFunc(src.OrderByAtom, func(l, r *Type) bool {
		return l.Atom < r.Atom
	})

	return
}

func write(w io.Writer, src Source) error {
	funcMap := template.FuncMap{}
	t := template.Must(template.New("main").Funcs(funcMap).Parse(srcTemplate))
	if err := t.Execute(w, src); err != nil {
		return errors.New("Failed to execute template: " + err.Error())
	}
	return nil
}

func usage() {
	_, _ = fmt.Fprintf(os.Stderr, "usage: mkwinsyscall [flags] [path ...]\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func mainErr() error {
	var (
		filename = flag.String("output", "", "output file name (standard output if omitted)")
	)

	flag.Usage = usage
	flag.Parse()

	src := Source{
		PackageName: os.Getenv("GOPACKAGE"),
		TypeName:    "ResponseProxy",
		TypeNameL:   "responseProxy",
		Traits: []Trait{
			{"Rf", "io.ReaderFrom"},
			{"Sw", "io.StringWriter"},
			{"Ph", "http.Pusher"},
			{"Fl", "http.Flusher"},
			{"Hj", "http.Hijacker"},
		},
	}
	src.fillTypes()

	var buf bytes.Buffer
	if err := write(&buf, src); err != nil {
		return err
	}

	out, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	if *filename == "" {
		_, err = os.Stdout.Write(out)
		return err
	} else {
		return os.WriteFile(*filename, out, 0644)
	}
}

func main() {
	if err := mainErr(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "error:", err)
	}
}

//go:embed gen.go.tmpl
var srcTemplate string
