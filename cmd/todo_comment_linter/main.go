package main

import (
	"github.com/mrymam/todo_comment_linter"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(todo_comment_linter.Analyzer) }
