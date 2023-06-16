package transpile

import (
	"fmt"
	"go/ast"
	"strconv"
	"strings"

	"golang.org/x/tools/go/ast/astutil"
)

func ChangeErr4AssignStmt(x *ast.AssignStmt, c *astutil.Cursor) (changed bool) {
	errsArgs := make([]ast.Expr, len(x.Lhs))
	errs := []string{}
	for i, v := range x.Lhs {
		t := SkipChangeErr4AssignStmt(v)
		if t == nil {
			errsArgs[i] = ast.NewIdent("nil")
			continue
		}
		errsArgs[i] = ast.NewIdent("&" + t.Name)
		errs = append(errs, t.Name)
		t.Name = "_"
	}
	if len(errs) == 0 {
		return
	}
	changed = true
	check := "Check"
	if n := len(x.Lhs) - 1; n > 0 {
		check = check + strconv.Itoa(n)
	}
	fn := &ast.SelectorExpr{
		X:   ast.NewIdent("err4"),
		Sel: ast.NewIdent(check),
	}
	fn2 := &ast.CallExpr{
		Fun:  fn,
		Args: x.Rhs,
	}
	call := &ast.CallExpr{
		Fun:  fn2,
		Args: errsArgs,
	}
	x.Rhs = []ast.Expr{call}
	// add return stmt
	keys := map[string]int{}
	vars := []string{}
	for _, name := range errs {
		if keys[name] != 0 {
			continue
		}
		keys[name] = 1

		vars = append(vars, fmt.Sprintf("%s != nil", name))
	}
	cond := ast.NewIdent(strings.Join(vars, " || "))
	checkReturnStmt := &ast.IfStmt{
		Cond: cond,
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ReturnStmt{},
			},
		},
	}
	c.InsertAfter(checkReturnStmt)
	return
}

func SkipChangeErr4AssignStmt(v ast.Expr) *ast.Ident {
	vv, ok := v.(*ast.Ident)
	if !ok {
		return nil
	}
	if !strings.HasPrefix(vv.Name, "qT") {
		return nil
	}
	switch src := vv.Obj.Decl.(type) {
	case *ast.ValueSpec:
		t, ok := src.Type.(*ast.Ident)
		if !ok {
			return nil
		}
		if t.Name != "error" {
			return nil
		}
	case *ast.Field:
		t, ok := src.Type.(*ast.Ident)
		if !ok {
			return nil
		}
		if t.Name != "error" {
			return nil
		}
	default:
		return nil
	}

	return vv
}

func RepalceErr4BuildTag(x *ast.File) bool {
	for _, cc := range x.Comments {
		for _, c := range cc.List {
			if strings.HasPrefix(c.Text, "//go:build ") {
				c.Text = strings.ReplaceAll(c.Text, " err4", " !err4")
				return true
			}
		}
	}
	return false
}
