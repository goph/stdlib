package main

import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"os"

	"github.com/goph/stdlib/strings"
)

var types = []struct {
	t   string
	def string
}{
	{"string", `""`},
	{"bool", "false"},
	{"int", "0"},
	{"int32", "0"},
	{"int64", "0"},
	{"float32", "0.0"},
	{"float64", "0.0"},
}

func main() {
	fset := token.NewFileSet()

	decl := []ast.Decl{
		&ast.GenDecl{
			Tok: token.IMPORT,
			Specs: []ast.Spec{
				&ast.ImportSpec{
					Path: &ast.BasicLit{
						Kind:  token.STRING,
						Value: `"fmt"`,
					},
				},
			},
		},
	}

	for _, t := range types {
		decl = append(decl, typeErrorGetter(t.t, t.def), typeGetter(t.t))
	}

	file := &ast.File{
		Name:  ast.NewIdent("ext"),
		Decls: decl,
	}

	f, err := os.OpenFile("args_gen.go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("%x", err)
		return
	}

	f.WriteString("//+build experimental\n\n")

	printer.Fprint(f, fset, file)

	err = f.Close()
	if err != nil {
		fmt.Printf("%x", err)
		return
	}
}

func typeErrorGetter(t string, def string) *ast.FuncDecl {
	return &ast.FuncDecl{
		Name: ast.NewIdent(fmt.Sprintf("%sE", strings.ToCamel(t))),
		Doc: &ast.CommentGroup{
			List: []*ast.Comment{
				{
					Text: fmt.Sprintf(
						"// %sE returns a(n) %s argument from the list or an error if it cannot be found or not %s.",
						strings.ToCamel(t),
						t,
						t,
					),
				},
			},
		},
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{
						ast.NewIdent("a"),
					},
					Type: ast.NewIdent("Arguments"),
				},
			},
		},
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{
							ast.NewIdent("index"),
						},
						Type: ast.NewIdent("int"),
					},
				},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					{
						Type: ast.NewIdent(t),
					},
					{
						Type: ast.NewIdent("error"),
					},
				},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.AssignStmt{
					Tok: token.DEFINE,
					Lhs: []ast.Expr{
						ast.NewIdent("arg"),
						ast.NewIdent("err"),
					},
					Rhs: []ast.Expr{
						&ast.CallExpr{
							Fun: ast.NewIdent("a.GetE"),
							Args: []ast.Expr{
								ast.NewIdent("index"),
							},
						},
					},
				},
				&ast.IfStmt{
					Cond: &ast.BinaryExpr{
						X:  ast.NewIdent("err"),
						Op: token.NEQ,
						Y:  ast.NewIdent("nil"),
					},
					Body: &ast.BlockStmt{
						List: []ast.Stmt{
							&ast.ReturnStmt{
								Results: []ast.Expr{
									&ast.BasicLit{
										Kind:  token.STRING,
										Value: def,
									},
									ast.NewIdent("err"),
								},
							},
						},
					},
				},
				&ast.AssignStmt{
					Tok: token.DEFINE,
					Lhs: []ast.Expr{
						ast.NewIdent("v"),
						ast.NewIdent("ok"),
					},
					Rhs: []ast.Expr{
						&ast.TypeAssertExpr{
							X:    ast.NewIdent("arg"),
							Type: ast.NewIdent(t),
						},
					},
				},
				&ast.IfStmt{
					Cond: &ast.UnaryExpr{
						X:  ast.NewIdent("ok"),
						Op: token.NOT,
					},
					Body: &ast.BlockStmt{
						List: []ast.Stmt{
							&ast.ReturnStmt{
								Results: []ast.Expr{
									&ast.BasicLit{
										Kind:  token.STRING,
										Value: def,
									},
									&ast.CallExpr{
										Fun: ast.NewIdent("fmt.Errorf"),
										Args: []ast.Expr{
											&ast.BasicLit{
												Kind:  token.STRING,
												Value: fmt.Sprintf(`"cannot return argument (%%d) as %s because it is of type %%T"`, t),
											},
											ast.NewIdent("index"),
											ast.NewIdent("arg"),
										},
									},
								},
							},
						},
					},
				},
				&ast.ReturnStmt{
					Results: []ast.Expr{
						ast.NewIdent("v"),
						ast.NewIdent("nil"),
					},
				},
			},
		},
	}
}

func typeGetter(t string) *ast.FuncDecl {
	return &ast.FuncDecl{
		Name: ast.NewIdent(fmt.Sprintf("%s", strings.ToCamel(t))),
		Doc: &ast.CommentGroup{
			List: []*ast.Comment{
				{
					Text: fmt.Sprintf(
						"// %s returns a(n) %s argument from the list.",
						strings.ToCamel(t),
						t,
					),
				},
				{
					Text: "//",
				},
				{
					Text: fmt.Sprintf("// It panics if the argument with such index cannot be found or it is not %s.", t),
				},
			},
		},
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{
						ast.NewIdent("a"),
					},
					Type: ast.NewIdent("Arguments"),
				},
			},
		},
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{
							ast.NewIdent("index"),
						},
						Type: ast.NewIdent("int"),
					},
				},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					{
						Type: ast.NewIdent(t),
					},
				},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.AssignStmt{
					Tok: token.DEFINE,
					Lhs: []ast.Expr{
						ast.NewIdent("arg"),
						ast.NewIdent("err"),
					},
					Rhs: []ast.Expr{
						&ast.CallExpr{
							Fun: ast.NewIdent(fmt.Sprintf("a.%sE", strings.ToCamel(t))),
							Args: []ast.Expr{
								ast.NewIdent("index"),
							},
						},
					},
				},
				&ast.IfStmt{
					Cond: &ast.BinaryExpr{
						X:  ast.NewIdent("err"),
						Op: token.NEQ,
						Y:  ast.NewIdent("nil"),
					},
					Body: &ast.BlockStmt{
						List: []ast.Stmt{
							&ast.ExprStmt{
								X: &ast.CallExpr{
									Fun: ast.NewIdent("panic"),
									Args: []ast.Expr{
										ast.NewIdent("err"),
									},
								},
							},
						},
					},
				},
				&ast.ReturnStmt{
					Results: []ast.Expr{
						ast.NewIdent("arg"),
					},
				},
			},
		},
	}
}
