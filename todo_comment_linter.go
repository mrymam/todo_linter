package todo_comment_linter

import (
	"errors"
	"fmt"
	"go/ast"
	"strings"
	"time"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "todo_comment_linter is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "todo_comment_linter",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.Ident)(nil),
		(*ast.Comment)(nil),
		(*ast.CommentGroup)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		// fmt.Printf("n=%+v\n", n)
		// if v, ok := n.(*ast.Ident); ok {
		// 	// fmt.Printf("v=%+v\n", v)
		// 	if v.Obj != nil {
		// 		fmt.Printf("v=%+v\n", v.Obj.Kind.String())
		// 	}
		// 	// pass.Reportf(n.Pos(), "hogehoge")
		// }
		if v, ok := n.(*ast.Comment); ok {
			fmt.Printf("v=%+v\n", v)
		}
		// if v, ok := n.(*ast.CommentGroup); ok {
		// 	fmt.Printf("v=%+v\n", v)
		// }
		switch n := n.(type) {
		case *ast.Comment:
			fmt.Println(n.Text)
			if isTodoComment(n.Text) {
				err := checkFormat(n.Text)
				if err != nil {
					pass.Reportf(n.Pos(), err.Error())
				}
			}
		}
	})

	return nil, nil
}

func isTodoComment(comment string) bool {
	return strings.Contains(comment, "TODO")
}

func checkFormat(comment string) error {
	if !strings.HasPrefix(comment, "// TODO") {
		return errors.New("TODOから始めてください")
	}
	body := strings.TrimLeft(comment, "// TODO ")
	fields := strings.Split(body, " ")
	if len(fields) < 3 {
		return errors.New("fieldsが足りない")
	}
	if !strings.HasPrefix(fields[0], "@") {
		return errors.New("ユーザーを記入してください")
	}
	until := fields[1]
	if !strings.HasPrefix(until, "until:") {
		return errors.New("期限の記入フォーマットが異なります")
	}
	untilDateStr := strings.TrimLeft(until, "until:")
	untilDate, err := time.Parse("2006-01-02", untilDateStr)
	if err != nil {
		return errors.New("期限を正しく入力してください")
	}
	now := time.Now()
	if untilDate.Before(now) {
		return errors.New("期限切れです")
	}
	return nil
}
