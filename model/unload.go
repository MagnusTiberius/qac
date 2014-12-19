package model

import (
    //"fmt"
	"runtime"
)

func Unload(rtu *RtTranslationUnitList) {
	for i := range rtu.RtTranslationUnits {
		itm := rtu.RtTranslationUnits[i]
	    UnloadRtPackageDef(rtu.RtTranslationUnits[i].RtPackageDef)
	    UnloadRtFuncDefList(rtu.RtTranslationUnits[i].RtFuncDefList)
	    UnloadRtImportList(rtu.RtTranslationUnits[i].RtImportList)
	    itm.RtPackageDef = nil
	    itm.RtFuncDefList = nil
	    itm.RtImportList = nil
	    //rtu.RtTranslationUnits = rtu.RtTranslationUnits[:len(rtu.RtTranslationUnits)-1]
	}
	rtu.Clear()
	//rtu.RtTranslationUnits = []RtTranslationUnit{}
	runtime.GC()
	UnloadTestPlan()
}

func UnloadRtPackageDef(node *RtPackageDef) {
	//node.Name = ""
}

func UnloadRtFuncDefList(list []RtFuncDef) {
	for i := range list {
		itm := list[i]
	    itm.Name            = ""
	    UnloadRtCmpndStmt(itm.RtCmpndStmt )
	    itm.RtCmpndStmt     = nil
	    itm.Identifier      = nil
	    itm.SpecType        = nil
	    UnloadParamDeclList(itm.ParamDeclList)
	    itm.ParamDeclList   = []ParamDecl{}
	}
}

func UnloadRtCmpndStmt(node *RtCmpndStmt) {
	if node != nil {
		UnloadRtDeclaration(node.List)
		UnloadRtStmtList(node.StmtList)
	    node.List      = nil
	    node.StmtList  = nil
	}
}


func UnloadRtDeclaration(list []RtDeclaration) {
	for i := range list {
		itm := list[i]
	    itm.List        = nil //[]RtDeclInit
	    itm.SpecType    = nil //*SpecType
	}
}

func UnloadRtStmtList(list []Stmt) {
	for i := range list {
		itm := list[i]
		UnloadExprCmd(itm.ExprCmd )
		UnloadExprStmt(itm.ExprStmt)
	    itm.ExprCmd         = nil //*ExprCmd
	    itm.ExprStmt        = nil //*ExprStmt
	    itm.StmtJump        = nil //*StmtJump
	    itm.StmtIteration   = nil //*StmtIteration
	    itm.StmtSelection   = nil //*StmtSelection
	    itm.CmpndStmt       = nil //*CmpndStmt
	    itm.StmtStep        = nil //*StmtStep

	}
}



func UnloadExprStmt(node *ExprStmt) {
	if node != nil {
		UnloadExpr(node.Expr)
	    node.ExprAssign  = nil //**ExprAssign
	    node.Expr        = nil //**Expr
	}
}

func UnloadExpr(node *Expr) {
	if node != nil {
		UnloadExpr(node.Expr)
		UnloadExprAssign(node.ExprAssign)
	    node.ExprAssign  = nil //**ExprAssign
	    node.Expr        = nil //**Expr
	}
}

func UnloadExprAssign(node *ExprAssign) {
	if node != nil {
		UnloadExprCmd(node.ExprCmd )
		UnloadExprAssign(node.ExprAssign)
		UnloadExprUnary(node.ExprUnary)
		UnloadExprConditional(node.ExprConditional)
	    node.ExprCmd             = nil //*ExprCmd
	    node.ExprAssign          = nil //*ExprAssign
	    node.ExprUnary           = nil //*ExprUnary
	    node.ExprConditional     = nil //*ExprConditional
	}
}

func UnloadExprUnary(node *ExprUnary) {
	if node != nil {
		UnloadExprPostfix(node.ExprPostfix )
		UnloadUnaryOperator(node.UnaryOperator )
		UnloadExprCast(node.ExprCast )
	    node.ExprPostfix     = nil //*ExprPostfix
	    node.UnaryOperator   = nil //*UnaryOperator
	    node.ExprCast        = nil //*ExprCast
	}
}

func UnloadExprPostfix(node *ExprPostfix) {
	if node != nil {
		UnloadExpr(node.Expr)
		UnloadIdentifier(node.Identifier)
		UnloadExprPostfix(node.ExprPostfix )
	    node.ExprPostfix         = nil //*ExprPostfix
	    node.ExprPrimary         = nil //*ExprPrimary
	    node.ExprListArgument    = nil //*ExprListArgument
	    node.Expr                = nil //*Expr
	    node.Identifier          = nil //*Identifier
	    node.Operator            = "" //string
	    node.Modifier            = "" //string
	}
}

func UnloadExprPrimary(node *ExprPrimary) {
	if node != nil {
		UnloadExpr(node.Expr)
		UnloadExprAssign(node.ExprAssign)
	    node.Value       = "" //string
	    node.IdType      = 0 //int
	    node.Expr        = nil //*Expr
	    node.ExprAssign  = nil //*ExprAssign
	}
}


func UnloadUnaryOperator(node *UnaryOperator) {
	if node != nil {
		node.Sign = ""
	}
}

func UnloadExprCast(node *ExprCast) {
	if node != nil {
		node.ExprUnary = nil
	}
}


func UnloadExprConditional(node *ExprConditional) {
	if node != nil {
		UnloadExprLogicOr(node.ExprLogicOr)
		node.ExprLogicOr = nil
	}
}

func UnloadExprLogicOr(node *ExprLogicOr) {
	if node != nil {
		UnloadExprLogicOr(node.ExprLogicOr)
	    node.ExprLogicAnd        = nil //*ExprLogicAnd
	    node.ExprLogicOr         = nil //*ExprLogicOr
	}
}


func UnloadExprCmd(node *ExprCmd) {
	if node != nil {
	    node.CmdListApp          = nil //*CmdListApp
	    node.CmdMakeAppUser      = nil //*CmdMakeAppUser
	    node.CmdVerify           = nil //*CmdVerify
	    node.CmdGenData          = nil //*CmdGenData
	    node.CmdMakeAppTid       = nil //*CmdMakeAppTid
	    node.CmdTestRoute        = nil //*CmdTestRoute
	    node.CmdCreateApp        = nil //*CmdCreateApp
	    node.CmdDeleteApp        = nil //*CmdDeleteApp
	    node.CmdDeleteAppTid     = nil //*CmdDeleteAppTid
	    node.CmdRunSQL           = nil //*CmdRunSQL
	    node.CmdDeleteAppUser    = nil //*CmdDeleteAppUser
	}
}


func UnloadParamDeclList(list []ParamDecl) {
	for i := range list {
		itm := list[i]
		UnloadDeclarator(itm.Declarator)
	    itm.Declarator = nil
	    UnloadSpecDecl(itm.SpecDecl)
	    itm.SpecDecl = nil
	}
}

func UnloadDeclarator(node *Declarator) {
	if node != nil {
		if node.DeclDirect != nil {
			UnloadDeclDirect(node.DeclDirect )
		}
	}
}

func UnloadDeclDirect(node *DeclDirect) {
	if node != nil {
	    //node.Identifier     = nil
	    UnloadIdentifierList(node.IdentifierList)
	    UnloadDeclarator(node.Declarator)
	    UnloadDeclDirect(node.DeclDirect )
	    node.Declarator      = nil
	    node.DeclDirect      = nil
	    node.ParamTypeList   = nil
	    node.ExprConstant    = nil
	    node.IdentifierList  = nil
	    node.Modifier        = ""
	}
}

func UnloadParamTypeList(node *ParamTypeList) {
	if node != nil {
		UnloadParamList(node.ParamList )
	    node.ParamList       = nil
	}
}

func UnloadParamList(node *ParamList) {
	if node != nil {
		UnloadParamList(node.ParamList )
		UnloadParamDecl(node.ParamDecl )
	    node.ParamList       = nil
	    node.ParamDecl       = nil
	}
}

func UnloadParamDecl(node *ParamDecl) {
	if node != nil {
		UnloadDeclarator(node.Declarator )
		UnloadSpecDecl(node.SpecDecl )
	    node.Declarator       = nil
	    node.SpecDecl       = nil

	}
}

func UnloadSpecDecl(node *SpecDecl) {
	if node != nil {
	    //node.SpecType       = nil
	}
}


func UnloadIdentifierList(node *IdentifierList) {
	if node != nil {
		UnloadIdentifier(node.Identifier )
		UnloadIdentifierList(node.IdentifierList)
	    node.Identifier       = nil
	    node.IdentifierList   = nil
	}
}

func UnloadIdentifier(node *Identifier) {
	//node.SpecType = nil
}




func UnloadRtImportList(node *RtImportList) {
	node.List = nil
}

func UnloadCmpndStmt(node *CmpndStmt) {
	if node != nil {
		UnloadStmtList(node.StmtList )
		UnloadDeclarationList(node.DeclarationList )
	    node.StmtList       	= nil
	    node.DeclarationList    = nil
	}
}

func UnloadStmtList(node *StmtList) {
	if node != nil {
		UnloadStmtList(node.StmtList )
		UnloadStmt(node.Stmt )
	    node.StmtList       	= nil
	    node.Stmt    = nil
	}
}


func UnloadStmt(node *Stmt) {
	if node != nil {
		UnloadExprCmd(node.ExprCmd )
		UnloadExprStmt(node.ExprStmt )
		UnloadStmtJump(node.StmtJump )
		UnloadStmtIteration(node.StmtIteration )
		UnloadStmtSelection(node.StmtSelection )
		UnloadCmpndStmt(node.CmpndStmt )
		UnloadStmtStep(node.StmtStep )
	    node.ExprCmd       	= nil
	    node.ExprStmt    	= nil
	    node.StmtJump       = nil
	    node.StmtIteration  = nil
	    node.StmtSelection  = nil
	    node.CmpndStmt    	= nil
	    node.StmtStep       = nil
	}
}

func UnloadStmtJump(node *StmtJump) {
	if node != nil {
	    //Name        string
		UnloadExpr(node.Expr )
	    node.Expr        = nil
	}
}

func UnloadStmtIteration(node *StmtIteration) {
	if node != nil {
		UnloadStmtWhile(node.StmtWhile )
		UnloadStmtFor(node.StmtFor )
		UnloadStmtDo(node.StmtDo )
		UnloadStmtStep(node.StmtStep )
	    node.StmtWhile           = nil
	    node.StmtFor             = nil
	    node.StmtDo              = nil
	    node.StmtStep            = nil
	}
}

func UnloadStmtWhile(node *StmtWhile) {
	if node != nil {
		UnloadExpr(node.Expr )
		UnloadStmt(node.Stmt )
	    node.Expr                = nil
	    node.Stmt                = nil
	}
}

func UnloadStmtFor(node *StmtFor) {
	if node != nil {
		UnloadExpr(node.Expr )
		UnloadExprStmt(node.ExprStmt1 )
		UnloadExprStmt(node.ExprStmt2 )
		UnloadStmt(node.Stmt )
	    node.Expr                = nil
	    node.ExprStmt1           = nil
	    node.ExprStmt2           = nil
	    node.Stmt                = nil
	}
}

func UnloadStmtDo(node *StmtDo) {
	if node != nil {
		UnloadExpr(node.Expr )
		UnloadStmt(node.Stmt )
	    node.Expr                = nil
	    node.Stmt                = nil
	}
}

func UnloadStmtSelection(node *StmtSelection) {
	if node != nil {
		UnloadExpr(node.Expr )
		UnloadStmt(node.Stmt )
		UnloadStmt(node.Stmt2 )
	    node.Expr                = nil
	    node.Stmt                = nil
	    node.Stmt2               = nil
	}
}


func UnloadStmtStep(node *StmtStep) {
	if node != nil {
		UnloadExpr(node.Expr )
		UnloadCmpndStmt(node.CmpndStmtStimulus )
		UnloadCmpndStmt(node.CmpndStmtResponse )
		UnloadIdentifier(node.Identifier )
	    node.Expr                = nil
	    node.CmpndStmtStimulus   = nil
	    node.CmpndStmtResponse   = nil
	    node.Identifier          = nil
	}
}

func UnloadDeclarationList(node *DeclarationList) {
	if node != nil {
		UnloadDeclInitList(node.DeclInitList )
		UnloadDeclarationList(node.DeclarationList )
		UnloadDeclaration(node.Declaration )
	    node.DeclInitList       = nil
	    node.DeclarationList    = nil
	    node.Declaration    	= nil
	}
}

func UnloadDeclInitList(node *DeclInitList) {
	if node != nil {
		UnloadDeclInit(node.DeclInit )
		UnloadDeclInitList(node.DeclInitList )
	    node.DeclInit        = nil
	    node.DeclInitList    = nil
	}
}

func UnloadDeclInit(node *DeclInit) {
	if node != nil {
		UnloadDeclarator(node.Declarator )
		UnloadInitializer(node.Initializer )
	    node.Declarator       = nil
	    node.Initializer      = nil
	}
}

func UnloadInitializer(node *Initializer) {
	if node != nil {
		UnloadExprAssign(node.ExprAssign )
		UnloadExprUnary(node.ExprUnary )
		UnloadExprCmd(node.ExprCmd )
		UnloadInitializerList(node.InitializerList )
	    node.ExprAssign        	= nil
	    node.ExprUnary    		= nil
	    node.ExprCmd        	= nil
	    node.InitializerList    = nil
	}
}

func UnloadInitializerList(node *InitializerList) {
	if node != nil {
		UnloadInitializer(node.Initializer )
		UnloadInitializerList(node.InitializerList )
	    node.Initializer    = nil
	    node.InitializerList        = nil
	}
}


func UnloadDeclaration(node *Declaration) {
	if node != nil {
		UnloadDeclInitList(node.DeclInitList )
		UnloadSpecDecl(node.SpecDecl )
	    node.DeclInitList    = nil
	    node.SpecDecl        = nil
	}
}




