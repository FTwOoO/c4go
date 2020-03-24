package transpiler

import "go/ast"

func cleanImplicitTypeConversion(in []ast.Decl) (ret []ast.Decl) {
	ret = make([]ast.Decl, len(in))
	for i, x := range in {
		ret[i] = _cleanImplicitTypeConversionForDecl(x)
	}
	return ret
}

func _cleanImplicitTypeConversionForDecl(in ast.Decl) (ret ast.Decl) {
	switch in.(type) {
	case *ast.FuncDecl:
		funcDecl := in.(*ast.FuncDecl)
		newStmts := []ast.Stmt{}
		for _, st := range funcDecl.Body.List {
			newSt := _cleanImplicitTypeConversionForStmt(st)
			newStmts = append(newStmts, newSt)
		}

		funcDecl.Body = &ast.BlockStmt{List: newStmts}
		return funcDecl

	default:
		return in
	}

	return in
}

func _cleanImplicitTypeConversionForStmt(in ast.Stmt) (ret ast.Stmt) {
	switch in.(type) {
	case *ast.IfStmt:
		ifStmt := in.(*ast.IfStmt)
		ifStmt.Cond = _cleanImplicitTypeConversionForExpr(ifStmt.Cond)

		newStmts := []ast.Stmt{}
		for _, st := range ifStmt.Body.List {
			newSt := _cleanImplicitTypeConversionForStmt(st)
			newStmts = append(newStmts, newSt)
		}
		ifStmt.Body = &ast.BlockStmt{List: newStmts}
		return ifStmt
		//TODO: else stmt
	default:
		return in
	}

	return in
}

func _cleanImplicitTypeConversionForExpr(in ast.Expr) (ret ast.Expr) {
	switch in.(type) {
	case *ast.BinaryExpr:
		expr := in.(*ast.BinaryExpr)
		expr.X = _cleanImplicitTypeConversionForExpr(expr.X)
		expr.Y = _cleanImplicitTypeConversionForExpr(expr.Y)
		return expr
	case *ast.CallExpr:

		f := in.(*ast.CallExpr)
		if _, ok := f.Fun.(*ast.Ident); ok {
			return _cleanCallExprChain(f)
		}

		return in
	default:
		return in
	}
}

//主要用于整数的互相转换
func _cleanCallExprChain(in *ast.CallExpr) (ret ast.Expr) {

	current := in
	for {
		if _, ok := current.Fun.(*ast.Ident); ok {
			if len(current.Args) == 1 {
				if nextF, ok := current.Args[0].(*ast.CallExpr); ok {
					current = nextF
					continue
				} else if p, ok := current.Args[0].(*ast.ParenExpr); ok {
					if nextF, ok := p.X.(*ast.CallExpr); ok {
						current = nextF
						continue
					} else {
						in.Args[0] = p.X
						return in
					}
				} else {
					in.Args[0] = current.Args[0]
					return in
				}
			}
		}

		break
	}

	return in
}
