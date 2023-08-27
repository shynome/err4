package transpile

import (
	"fmt"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/ast/astutil"
)

func ChangeErr4AssignStmt(x *ast.AssignStmt, c *astutil.Cursor) (changed bool) {
	if _, ok := c.Parent().(*ast.BlockStmt); !ok {
		return
	}
	errs := []string{}
	for _, v := range x.Lhs {
		t := SkipChangeErr4AssignStmt(v)
		if t == nil {
			continue
		}
		errs = append(errs, t.Name)
	}
	if len(errs) == 0 {
		return
	}
	changed = true
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
	if !strings.HasPrefix(vv.Name, "i") {
		return nil
	}
	if vv.Obj == nil || vv.Obj.Decl == nil {
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
			if strings.HasPrefix(c.Text, "//go:build ierr") {
				c.Text = strings.Replace(c.Text, "//go:build ierr", "//go:build !ierr", 1)
				return true
			}
		}
	}
	return false
}
