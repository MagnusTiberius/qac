package model

import (
    "fmt"
)



func DoDumpTree(node *TranslationUnit, indent int) *TranslationUnit {
    if node == nil {
        return node
    }
    rlog(indent,"mTranslationUnit:%v\n", node)
    DumpTranslationUnit(node, 1)
    return node
}

func DoInd(n int) {
    for i:=0; i<n; i++ {
        fmt.Printf(" ")
    }
}

func DumpTranslationUnit(node *TranslationUnit, indent int) {
    if node != nil {
        
        rlog(indent,"TranslationUnit: %v \n", node)
        DumpTranslationUnit(node.TranslationUnit, indent+1)
        DumpDeclExternal(node.DeclExternal, indent+1)
    }    
}

func DumpDeclExternal(node *DeclExternal, indent int) {
    if node != nil {
        
        rlog(indent,"DeclExternal: %v\n", node)
        DumpPackageDef(node.PackageDef, indent+1)
        DumpImportDef(node.ImportDef, indent+1)
        DumpFuncDef(node.FuncDef, indent+1)
    }    
}

func DumpFuncDef(node *FuncDef, indent int) {
    if node != nil {
        
        rlog(indent,"FuncDef: %v\n", node)
        DumpSpecDecl(node.SpecDecl, indent+1)
        DumpDeclarator(node.Declarator, indent+1)
        DumpCmpndStmt(node.CmpndStmt, indent+1)
    }    
}

func DumpImportDef(node *ImportDef, indent int) {
    if node != nil {
        
        rlog(indent,"ImportDef: %v\n", node)
        DumpImportList(node.ImportList, indent+1)
    }    
}

func DumpImportList(node *ImportList, indent int) {
    if node != nil {
        
        rlog(indent,"ImportList: %v\n", node.List )
    }    
}

func DumpPackageDef(node *PackageDef, indent int) {
    if node != nil {
        
        rlog(indent,"PackageDef: %v\n", node)
        DumpImportDef(node.ImportDef, indent+1)
        DumpFuncDef(node.FuncDef, indent+1)
    }    
}

func DumpExprListArgument(node *ExprListArgument, indent int) {
    if node != nil {
        
        rlog(indent,"ExprListArgument: %v\n", node)
        DumpExprListArgument(node.ExprListArgument, indent+1)
        DumpExprAssign(node.ExprAssign, indent+1)
    }    
}

func DumpCmpndStmt(node *CmpndStmt, indent int) {
    if node != nil {
        mCmpndStmt = node
        
        rlog(indent,"CmpndStmt: %v %v %v\n", node, mCmpndStmt.StmtList, mCmpndStmt.DeclarationList )
        DumpDeclarationList(node.DeclarationList,indent+1)
        DumpStmtList(node.StmtList ,indent+1)
    }    
}

func DumpDeclarationList(node *DeclarationList, indent int) {
    if node != nil {
        
        rlog(indent,"DeclarationList: %v\n", node)
        DumpDeclInitList(node.DeclInitList,indent+1 )
        DumpDeclarationList(node.DeclarationList, indent+1 )
        DumpDeclaration(node.Declaration, indent+1 )

    }
}

func DumpDeclaration(node *Declaration, indent int) {
    if node != nil {
        
        rlog(indent,"Declaration: %v\n", node)
        DumpDeclInitList(node.DeclInitList,indent+1)
        DumpSpecDecl(node.SpecDecl,indent+1)
    }
}

func DumpSpecDecl(node *SpecDecl,indent int) {
    if node != nil {
        
        rlog(indent,"SpecDecl: %v\n", node)
        DumpSpecType(node.SpecType,indent+1)
    }
}

func DumpSpecType(node *SpecType, indent int) {
    if node != nil {
        
        rlog(indent,"SpecType: %v\n", node)
    }
}

func DumpDeclInitList(node *DeclInitList, indent int) {
    if node != nil {
        
        rlog(indent,"DeclInitList: %v\n", node)
        DumpDeclInit(node.DeclInit, indent+1)
        DumpDeclInitList(node.DeclInitList, indent+1)
    }    
}

func DumpDeclInit(node *DeclInit, indent int) {
    if node != nil {
        
        rlog(indent,"DeclInit: %v\n", node)
        DumpDeclarator(node.Declarator, indent+1)
        DumpInitializer(node.Initializer, indent+1)
    }
}

func DumpStmtList(node *StmtList, indent int) {
    if node != nil {
        
        rlog(indent,"StmtList: %v\n", node)
        DumpStmtList(node.StmtList, indent+1)
        DumpStmt(node.Stmt, indent+1)
    }
}

func DumpStmt(node *Stmt, indent int) {
    if node != nil {
        
        rlog(indent,"Stmt: %v\n", node)
        DumpExprCmd(node.ExprCmd,indent+1)
        DumpExprStmt(node.ExprStmt, indent+1)
        DumpStmtJump(node.StmtJump, indent+1)
        DumpStmtIteration(node.StmtIteration, indent+1)
        DumpStmtSelection(node.StmtSelection, indent+1)
        DumpCmpndStmt(node.CmpndStmt, indent+1)
        DumpStmtStep(node.StmtStep, indent+1)
    }
}

func DumpStmtStep(node *StmtStep, indent int) {
    if node != nil {
        
        rlog(indent,"StmtStep: %v\n", node)
        DumpIdentifier(node.Identifier, indent+1)
        DumpExpr(node.Expr, indent+1)
        DumpCmpndStmt(node.CmpndStmtStimulus, indent+1)
        DumpCmpndStmt(node.CmpndStmtResponse, indent+1)
    }
}


func DumpStmtIteration(node *StmtIteration, indent int) {
    if node != nil {
        
        rlog(indent,"StmtIteration: %v\n", node)
        DumpStmtWhile(node.StmtWhile, indent+1)
        DumpStmtFor(node.StmtFor, indent+1)
        DumpStmtDo(node.StmtDo, indent+1)
        DumpStmtStep(node.StmtStep, indent+1)
    }
}

func DumpStmtWhile(node *StmtWhile, indent int) {
    if node != nil {
        
        rlog(indent,"StmtWhile: %v\n", node)
        DumpExpr(node.Expr, indent+1)
        DumpStmt(node.Stmt, indent+1)
    }
}
func DumpStmtFor(node *StmtFor, indent int) {
    if node != nil {
        
        rlog(indent,"StmtFor: %v\n", node)
        DumpExpr(node.Expr, indent+1)
        DumpExprStmt(node.ExprStmt1, indent+1)
        DumpExprStmt(node.ExprStmt2, indent+1)
        DumpStmt(node.Stmt, indent+1)
    }
}
func DumpStmtDo(node *StmtDo, indent int) {
    if node != nil {
        
        rlog(indent,"StmtDo: %v\n", node)
        DumpExpr(node.Expr, indent+1)
        DumpStmt(node.Stmt, indent+1)
    }
}

func DumpStmtSelection(node *StmtSelection, indent int) {
    if node != nil {
        
        rlog(indent,"StmtSelection: %vn", node)
        DumpExpr(node.Expr, indent+1)
        DumpStmt(node.Stmt, indent+1)
    }
}


func DumpStmtJump(node *StmtJump, indent int) {
    if node != nil {
        
        rlog(indent,"StmtJump: %v\n", node)
        DumpExpr(node.Expr, indent+1)
    }
}

func DumpExprStmt(node *ExprStmt, indent int) {
    if node != nil {
        
        rlog(indent,"ExprStmt: %v %v\n", node.ExprAssign, node.Expr )
        mExprStmt := node
        DumpExprAssign(mExprStmt.ExprAssign, indent+1)
        DumpExpr(node.Expr, indent+1)
    }
}

func DumpExpr(node *Expr, indent int) {
    if node != nil {
        DumpExprAssign(node.ExprAssign, indent+1)
        DumpExpr(node.Expr, indent+1)
    }
}

func DumpExprAssign(node *ExprAssign, indent int) {
    if node != nil {
        
        rlog(indent,"mExprAssign: %v\n", node)
        DumpExprUnary(node.ExprUnary, indent+1)
        DumpExprCmd(node.ExprCmd,indent+1)
        DumpExprAssign(node.ExprAssign, indent+1)
        DumpExprConditional(node.ExprConditional,indent+1)
    }
}

func DumpExprConditional(node *ExprConditional, indent int) {
    if node != nil {
        
        rlog(indent,"ExprConditional: %v\n", node)
        DumpExprLogicOr(node.ExprLogicOr,indent+1)
    }
}

func DumpExprLogicOr(node *ExprLogicOr, indent int) {
    if node != nil {
        
        rlog(indent,"ExprLogicOr: %v\n", node)
        DumpExprLogicAnd(node.ExprLogicAnd,indent+1)
        DumpExprLogicOr(node.ExprLogicOr,indent+1)
    }    
}

func DumpExprLogicAnd(node *ExprLogicAnd, indent int) {
    if node != nil {
        
        rlog(indent,"ExprLogicAnd: %v\n", node)
        DumpExprLogicAnd(node.ExprLogicAnd,indent+1)
        DumpExprInclOr(node.ExprInclOr,indent+1)
    }
}

func DumpExprInclOr(node *ExprInclOr, indent int) {
    if node != nil {
        
        rlog(indent,"ExprInclOr: %v\n", node)
        DumpExprInclOr(node.ExprInclOr,indent+1)
        DumpExprExclOr(node.ExprExclOr,indent+1)
    }
}

func DumpExprExclOr(node *ExprExclOr, indent int) {
    if node != nil {
        
        rlog(indent,"ExprExclOr: %v\n", node)
        DumpExprExclOr(node.ExprExclOr,indent+1)
        DumpExprAnd(node.ExprAnd,indent+1)
    }
}

func DumpExprAnd(node *ExprAnd, indent int) {
    if node != nil {
        
        rlog(indent,"ExprAnd: %v\n", node)
        DumpExprAnd(node.ExprAnd,indent+1)
        DumpExprEquality(node.ExprEquality,indent+1)
    }
}

func DumpExprEquality(node *ExprEquality, indent int) {
    if node != nil {
        
        rlog(indent,"ExprEquality: %v %v\n", node, node.Sign )
        DumpExprEquality(node.ExprEquality,indent+1)
        DumpExprRelational(node.ExprRelational,indent+1)
    }
}

func DumpExprRelational(node *ExprRelational, indent int) {
    if node != nil {
        
        rlog(indent,"ExprRelational: %v %v\n", node, node.Sign )
        DumpExprRelational(node.ExprRelational,indent+1)
        DumpExprShift(node.ExprShift,indent+1)
    }
}

func DumpExprShift(node *ExprShift, indent int) {
    if node != nil {
        
        rlog(indent,"ExprShift: %v\n", node)
        DumpExprShift(node.ExprShift,indent+1)
        DumpExprAdditive(node.ExprAdditive,indent+1)
    }
}

func DumpExprAdditive(node *ExprAdditive, indent int) {
    if node != nil {
        
        rlog(indent,"ExprAdditive: %v %v\n", node, node.Sign )
        DumpExprAdditive(node.ExprAdditive,indent+1)
        DumpExprMultiplicative(node.ExprMultiplicative,indent+1)
    }
}

func DumpExprMultiplicative(node *ExprMultiplicative, indent int) {
    if node != nil {
        
        rlog(indent,"ExprMultiplicative: %v %v\n", node, node.Sign )
        DumpExprMultiplicative(node.ExprMultiplicative,indent+1)
        DumpExprCast(node.ExprCast,indent+1)
    }
}

func DumpExprCast(node *ExprCast, indent int) {
    if node != nil {
        
        rlog(indent,"ExprCast: %v\n", node)
        DumpExprUnary(node.ExprUnary,indent+1)
    }
}

func DumpExprUnary(node *ExprUnary, indent int) {
    if node != nil {
        
        rlog(indent,"ExprUnary: %v\n", node)
        DumpExprPostfix(node.ExprPostfix,indent+1)
        DumpExprCast(node.ExprCast,indent+1)
        DumpUnaryOperator(node.UnaryOperator,indent+1)
    }    
}

func DumpUnaryOperator(node *UnaryOperator, indent int) {
    if node != nil {
        
        rlog(indent,"UnaryOperator: %v\n", node)
    }
}

func DumpExprPostfix(node *ExprPostfix, indent int) {
    if node != nil {
        
        rlog(indent,"ExprPostfix: %v\n", node)
        if len(node.Modifier) > 0  {
            
            rlog(indent,"Modifier: %v\n", node.Modifier )
        }
        DumpExprPostfix(node.ExprPostfix,indent+1)
        DumpExprPrimary(node.ExprPrimary,indent+1)
        DumpExpr(node.Expr,indent+1)
        DumpExprListArgument(node.ExprListArgument,indent+1)
        if len(node.Operator) > 0  {
            
            rlog(indent,"Operator: %v\n", node.Operator )
            DumpIdentifier(node.Identifier,indent+1)
        }
    }    
}

func DumpDeclarator(node *Declarator, indent int) {
    if node != nil {
        mDeclarator := node
        
        rlog(indent,"Declarator: %v\n", node)
        DumpDeclDirect(mDeclarator.DeclDirect, indent+1)
    }
}

func DumpDeclDirect(node *DeclDirect, indent int) {
    if node != nil {
        
        rlog(indent,"DeclDirect: %v\n", node)
        DumpIdentifier(node.Identifier, indent+1)
        DumpIdentifierList(node.IdentifierList, indent+1)
        DumpParamTypeList(node.ParamTypeList, indent+1)
        DumpDeclarator(node.Declarator, indent+1)
        DumpDeclDirect(node.DeclDirect, indent+1)
        DumpExprConstant(node.ExprConstant, indent+1)

    }    
}

func DumpIdentifier(node *Identifier, indent int) {
    if node != nil {
        
        rlog(indent,"Identifier: %v\n", node)
        //DumpDeclDirect(mDeclarator.DeclDirect, indent+1)
    }
}

func DumpInitializer(node *Initializer, indent int) {
    if node != nil {
        
        rlog(indent,"Initializer: %v\n", node)
        DumpExprUnary(node.ExprUnary, indent+1)
        DumpExprAssign(node.ExprAssign, indent+1)
        DumpExprCmd(node.ExprCmd,indent+1)
        DumpInitializerList(node.InitializerList, indent+1)
    }
}


func DumpInitializerList(node *InitializerList, indent int) {
    if node != nil {
        
        rlog(indent,"InitializerList: %v\n", node)
        DumpInitializer(node.Initializer, indent+1)
        DumpInitializerList(node.InitializerList, indent+1)
    }
}

func DumpExprPrimary(node *ExprPrimary, indent int) {
    if node != nil {
        mExprPrimary = node
        
        rlog(indent,"mExprPrimary value: %v\n", mExprPrimary )
        if node.IdType == EXPR_ID {
            DumpExprAssign(node.ExprAssign,indent+1)
        }
        if node.IdType == IDENTIFIER {
            
            rlog(indent,"mExprPrimary IDENTIFIER: %v\n", mExprPrimary.Value )
        }
        if node.IdType == INTEGER_ID {
            
            rlog(indent,"mExprPrimary INTEGER_ID: %v\n", mExprPrimary.Value )
        }
        if node.IdType == EMAIL_ID {
            
            rlog(indent,"mExprPrimary EMAIL_ID: %v\n", mExprPrimary.Value )
        }
        if node.IdType == FLOAT_ID {
            
            rlog(indent,"mExprPrimary FLOAT_ID: %v\n", mExprPrimary.Value )
        }
        if node.IdType == TRUE_ID {
            
            rlog(indent,"mExprPrimary TRUE_ID: %v\n", mExprPrimary.Value )
        }
        if node.IdType == FALSE_ID {
            
            rlog(indent,"mExprPrimary FALSE_ID: %v\n", mExprPrimary.Value )
        }
    }
}

func DumpExprCmd(node *ExprCmd, indent int) {
    if node != nil {
        mExprCmd := node
        
        rlog(indent,"mExprCmd: %v\n", mExprCmd )
        if mExprCmd.CmdListApp != nil {
            mCmdListApp = mExprCmd.CmdListApp
            
            rlog(indent,"mCmdListApp: %v\n", mCmdListApp )
        }
        if mExprCmd.CmdMakeAppUser != nil {
            mCmdMakeAppUser = mExprCmd.CmdMakeAppUser
            
            rlog(indent,"mCmdMakeAppUser: %v\n", mCmdMakeAppUser )
        }
        if mExprCmd.CmdVerify != nil {
            mCmdVerify = mExprCmd.CmdVerify
            
            rlog(indent,"mCmdVerify: %v\n", mCmdVerify )
        }
        if mExprCmd.CmdGenData != nil {
            mCmdGenData = mExprCmd.CmdGenData
            
            rlog(indent,"mCmdGenData:%v \n", mCmdGenData )
        }
    }
}


func DumpParamList(node *ParamList, indent int) {
    if node != nil {
        
        rlog(indent,"ParamList: %v\n", node)
        DumpParamDecl(node.ParamDecl, indent+1)
        DumpParamList(node.ParamList, indent+1)
    }
}

func DumpParamTypeList(node *ParamTypeList, indent int) {
    if node != nil {
        
        rlog(indent,"ParamTypeList: %v\n", node)
        DumpParamList(node.ParamList, indent+1)
    }
}

func DumpIdentifierList(node *IdentifierList, indent int) {
    if node != nil {
        
        rlog(indent,"IdentifierList: %v\n", node)
        DumpIdentifier(node.Identifier, indent+1)
        DumpIdentifierList(node.IdentifierList, indent+1)
    }
}


func DumpParamDecl(node *ParamDecl, indent int) {
    if node != nil {
        
        rlog(indent,"ParamDecl: %v\n", node)
        DumpDeclarator(node.Declarator, indent+1)
        DumpSpecDecl(node.SpecDecl, indent+1)
    }
}


func DumpExprConstant(node *ExprConstant, indent int) {
    if node != nil {
        
        rlog(indent,"ExprConstant: %v\n", node)
        DumpExprConditional(node.ExprConditional, indent+1)
    }
}
