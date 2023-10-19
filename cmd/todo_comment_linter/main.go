package main

import (
	"github.com/mrymam/todo_linter"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(todo_linter.Analyzer) }
