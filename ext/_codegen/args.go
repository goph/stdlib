package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/printer"
	"go/token"
	"os"

	"github.com/goph/stdlib/strings"
)

var types = []struct {
	typ string
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

	decl := []ast.Decl{}

	for _, t := range types {
		decl = append(
			decl,
			typeLookup(t.typ, t.def),
			typeGet(t.typ, t.def),
			typeDefault(t.typ),
		)
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

	defer f.Close()

	var buf bytes.Buffer
	printer.Fprint(&buf, fset, file)

	formatted, err := format.Source(buf.Bytes())

	if err != nil {
		fmt.Printf("%x", err)
		return
	}

	f.WriteString("//+build experimental\n\n")
	f.Write(formatted)
}

func typeLookup(t string, def string) *ast.FuncDecl {
	return &ast.FuncDecl{
		Name: ast.NewIdent(fmt.Sprintf("Lookup%s", strings.ToCamel(t))),
		Doc: &ast.CommentGroup{
			List: []*ast.Comment{
				{
					Text: fmt.Sprintf(
						"// Lookup%s retrieves an argument of type %s from the list stored under the specified the index.",
						strings.ToCamel(t),
						t,
					),
				},
				{
					Text: "//",
				},
				{
					Text: fmt.Sprintf(
						"// If the index is present in the list and it is of type %s the value is returned and the boolean is true.",
						t,
					),
				},
				{
					Text: "//",
				},
				{
					Text: "// Otherwise the type's zero value and false are returned.",
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
						Type: ast.NewIdent("bool"),
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
						ast.NewIdent("ok"),
					},
					Rhs: []ast.Expr{
						&ast.CallExpr{
							Fun: ast.NewIdent("a.Lookup"),
							Args: []ast.Expr{
								ast.NewIdent("index"),
							},
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
									ast.NewIdent("false"),
								},
							},
						},
					},
				},
				&ast.IfStmt{
					Init: &ast.AssignStmt{
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
					Cond: ast.NewIdent("ok"),
					Body: &ast.BlockStmt{
						List: []ast.Stmt{
							&ast.ReturnStmt{
								Results: []ast.Expr{
									ast.NewIdent("v"),
									ast.NewIdent("true"),
								},
							},
						},
					},
				},
				&ast.ReturnStmt{
					Results: []ast.Expr{
						&ast.BasicLit{
							Kind:  token.STRING,
							Value: def,
						},
						ast.NewIdent("false"),
					},
				},
			},
		},
	}
}

func typeGet(t string, def string) *ast.FuncDecl {
	return &ast.FuncDecl{
		Name: ast.NewIdent(fmt.Sprintf("%s", strings.ToCamel(t))),
		Doc: &ast.CommentGroup{
			List: []*ast.Comment{
				{
					Text: fmt.Sprintf(
						"// %s retrieves an argument of type %s from the list stored under the specified the index.",
						strings.ToCamel(t),
						t,
					),
				},
				{
					Text: "//",
				},
				{
					Text: fmt.Sprintf(
						"// If the index is present in the list and it is of type %s the value is returned.",
						t,
					),
				},
				{
					Text: "//",
				},
				{
					Text: fmt.Sprintf(
						"// Otherwise the type's zero value is returned. To distinguish between an empty value and an unset value, use Lookup%s.",
						strings.ToCamel(t),
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
				},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.IfStmt{
					Init: &ast.AssignStmt{
						Tok: token.DEFINE,
						Lhs: []ast.Expr{
							ast.NewIdent("arg"),
							ast.NewIdent("ok"),
						},
						Rhs: []ast.Expr{
							&ast.CallExpr{
								Fun: ast.NewIdent(fmt.Sprintf("a.Lookup%s", strings.ToCamel(t))),
								Args: []ast.Expr{
									ast.NewIdent("index"),
								},
							},
						},
					},
					Cond: ast.NewIdent("ok"),
					Body: &ast.BlockStmt{
						List: []ast.Stmt{
							&ast.ReturnStmt{
								Results: []ast.Expr{
									ast.NewIdent("arg"),
								},
							},
						},
					},
				},
				&ast.ReturnStmt{
					Results: []ast.Expr{
						&ast.BasicLit{
							Kind:  token.STRING,
							Value: def,
						},
					},
				},
			},
		},
	}
}

func typeDefault(t string) *ast.FuncDecl {
	return &ast.FuncDecl{
		Name: ast.NewIdent(fmt.Sprintf("Default%s", strings.ToCamel(t))),
		Doc: &ast.CommentGroup{
			List: []*ast.Comment{
				{
					Text: fmt.Sprintf(
						"// Default%s retrieves an argument of type %s from the list stored under the specified the index.",
						strings.ToCamel(t),
						t,
					),
				},
				{
					Text: "//",
				},
				{
					Text: fmt.Sprintf(
						"// If the index is present in the list and it is of type %s the value is returned.",
						t,
					),
				},
				{
					Text: "//",
				},
				{
					Text: "// Otherwise the specified default value is returned.",
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
					{
						Names: []*ast.Ident{
							ast.NewIdent("def"),
						},
						Type: ast.NewIdent(t),
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
				&ast.IfStmt{
					Init: &ast.AssignStmt{
						Tok: token.DEFINE,
						Lhs: []ast.Expr{
							ast.NewIdent("arg"),
							ast.NewIdent("ok"),
						},
						Rhs: []ast.Expr{
							&ast.CallExpr{
								Fun: ast.NewIdent(fmt.Sprintf("a.Lookup%s", strings.ToCamel(t))),
								Args: []ast.Expr{
									ast.NewIdent("index"),
								},
							},
						},
					},
					Cond: ast.NewIdent("ok"),
					Body: &ast.BlockStmt{
						List: []ast.Stmt{
							&ast.ReturnStmt{
								Results: []ast.Expr{
									ast.NewIdent("arg"),
								},
							},
						},
					},
				},
				&ast.ReturnStmt{
					Results: []ast.Expr{
						ast.NewIdent("def"),
					},
				},
			},
		},
	}
}
