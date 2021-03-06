package model

import (
    //"fmt"
    //"intel/test/verify/parser"
    //"os"
    //"flag"
    //"io"
    //"bytes"
    //_ "github.com/lib/pq"
    //"database/sql"
    //"intel/test/selectorproperty"
)



type ExtDecl struct {
    PackageDef  *PackageDef
}

type ParamDecl struct {
    Declarator      *Declarator
    SpecDecl        *SpecDecl
}

type ParamList struct {
    ParamList       *ParamList
    ParamDecl       *ParamDecl
}

type ParamTypeList struct {
    ParamList       *ParamList
}

type FuncDef struct {
    Name            string
    Arglist         []ParamDecl
    CmpndStmt       *CmpndStmt
    Declarator      *Declarator
    SpecDecl        *SpecDecl
}

type PackageDef struct {
    Name            string
    ImportDef      *ImportDef
    FuncDef        *FuncDef
}

type ImportDef struct {
    ImportList  *ImportList
}

type ImportList struct {
    List        []string
}


type ExprCmd   struct {
    CmdListApp          *CmdListApp
    CmdMakeAppUser      *CmdMakeAppUser
    CmdVerify           *CmdVerify
    CmdGenData          *CmdGenData
    CmdMakeAppTid       *CmdMakeAppTid
    CmdTestRoute        *CmdTestRoute
    CmdCreateApp        *CmdCreateApp
    CmdDeleteApp        *CmdDeleteApp
    CmdDeleteAppTid     *CmdDeleteAppTid
    CmdRunSQL           *CmdRunSQL
    CmdDeleteAppUser    *CmdDeleteAppUser
}

type ExprAssign struct {
    ExprCmd             *ExprCmd
    ExprAssign          *ExprAssign
    ExprUnary           *ExprUnary
    ExprConditional     *ExprConditional
}

type ExprStmt struct {
    ExprAssign  *ExprAssign
    Expr        *Expr
}

type Expr struct {
    ExprAssign  *ExprAssign
    Expr        *Expr
}


type CmdListApp struct {
    Name        string
}

type CmdRunSQL struct {
    Name        string
    Expr1       *Expr
    Expr2       *Expr
    SQL         string
}

type CmdMakeAppUser struct {
    Name        string
    Expr1       *Expr
    Expr2       *Expr
}

type CmdMakeAppTid struct {
    Name        string
    Expr1       *Expr
    Expr2       *Expr
}

type CmdVerify struct {
    Name        string
}

type CmdTestRoute struct {
    Name        string
    Expr1       *Expr
    Expr2       *Expr
}

type CmdCreateApp struct {
    Name        string
    Expr1       *Expr
    Expr2       *Expr
}

type CmdDeleteApp struct {
    Name        string
    Expr1       *Expr
    Expr2       *Expr
}

type CmdDeleteAppUser struct {
    Name        string
    Expr1       *Expr
    Expr2       *Expr
}

type CmdDeleteAppTid struct {
    Name        string
    Expr1       *Expr
    Expr2       *Expr
}

type CmdGenData struct {
    Name        string
    ExprLoop    *Expr
    ExprApp     *Expr
}

type Stmt struct {
    ExprCmd         *ExprCmd
    ExprStmt        *ExprStmt
    StmtJump        *StmtJump
    StmtIteration   *StmtIteration
    StmtSelection   *StmtSelection
    CmpndStmt       *CmpndStmt
    StmtStep        *StmtStep
}

type StmtList struct {
    List    []Stmt
    StmtList    *StmtList
    Stmt        *Stmt
}

type CmpndStmt struct {
    StmtList        *StmtList
    DeclarationList *DeclarationList
}

type ExprPrimary struct {
    Value       string
    IdType      int
    Expr        *Expr
    ExprAssign  *ExprAssign
}

type ExprPostfix struct {
    ExprPostfix         *ExprPostfix
    ExprPrimary         *ExprPrimary
    ExprListArgument    *ExprListArgument
    Expr                *Expr
    Identifier          *Identifier
    Operator            string
    Modifier            string
}

type ExprUnary struct {
    ExprPostfix     *ExprPostfix
    UnaryOperator   *UnaryOperator
    ExprCast        *ExprCast
}

type UnaryOperator struct {
    Sign        string
}


type Initializer struct {
    ExprAssign          *ExprAssign
    ExprUnary           *ExprUnary
    ExprCmd             *ExprCmd
    InitializerList     *InitializerList
}

type InitializerList struct {
    Initializer         *Initializer
    InitializerList     *InitializerList
}


type DeclInit   struct {
    Declarator  *Declarator
    Initializer *Initializer
}

type Declarator struct {
    DeclDirect  *DeclDirect
}

type DeclDirect struct {
    Identifier      *Identifier
    Declarator      *Declarator
    DeclDirect      *DeclDirect
    ParamTypeList   *ParamTypeList
    ExprConstant    *ExprConstant
    IdentifierList  *IdentifierList
    Modifier        string
}

type DeclInitList   struct {
    List            []DeclInit
    DeclInit        *DeclInit
    DeclInitList    *DeclInitList
}

type Identifier struct {
    Name       string
}

type IdentifierList struct {
    Identifier       *Identifier
    IdentifierList   *IdentifierList
}


type Declaration struct {
    DeclInitList    *DeclInitList
    SpecDecl        *SpecDecl
}

type DeclarationList struct {
    List            []Declaration
    DeclInitList    *DeclInitList
    DeclarationList *DeclarationList
    Declaration     *Declaration
}

type ExprAdditive struct {
    ExprAdditive            *ExprAdditive
    ExprMultiplicative      *ExprMultiplicative
    Sign                    string
}

type ExprMultiplicative struct {
    ExprCast                *ExprCast
    ExprMultiplicative      *ExprMultiplicative
    Sign                    string
}

type ExprCast struct {
    ExprUnary               *ExprUnary
}

type ExprShift struct {
    ExprAdditive            *ExprAdditive
    ExprShift               *ExprShift
}

type ExprRelational struct {
    ExprShift               *ExprShift
    ExprRelational          *ExprRelational
    Sign                  string
}

type ExprEquality struct {
    ExprRelational        *ExprRelational
    ExprEquality          *ExprEquality
    Sign                  string
}

type ExprAnd struct {
    ExprEquality          *ExprEquality
    ExprAnd               *ExprAnd
}

type ExprExclOr struct {
    ExprAnd             *ExprAnd
    ExprExclOr          *ExprExclOr
}

type ExprInclOr struct {
    ExprExclOr          *ExprExclOr
    ExprInclOr          *ExprInclOr
}

type ExprLogicAnd struct {
    ExprInclOr          *ExprInclOr
    ExprLogicAnd        *ExprLogicAnd
}

type ExprLogicOr struct {
    ExprLogicAnd        *ExprLogicAnd
    ExprLogicOr         *ExprLogicOr
}

type ExprConditional struct {
    ExprLogicOr         *ExprLogicOr
}

type SpecDecl struct {
    SpecType    *SpecType
}

type SpecType struct {
    Name    string
    Id      int
}


type TranslationUnit struct {
    TranslationUnit *TranslationUnit
    DeclExternal    *DeclExternal
}

type DeclExternal struct {
    FuncDef         *FuncDef
    PackageDef      *PackageDef
    ImportDef       *ImportDef
}


type ExprConstant struct {
    ExprConditional    *ExprConditional
}

type StmtJump struct {
    Name        string
    Expr        *Expr
}

type ExprListArgument struct {
    ExprListArgument    *ExprListArgument
    ExprAssign          *ExprAssign
}

type StmtIteration struct {
    StmtWhile           *StmtWhile
    StmtFor             *StmtFor
    StmtDo              *StmtDo
    StmtStep            *StmtStep
}

type StmtWhile struct {
    Expr                *Expr
    Stmt                *Stmt
}

type StmtFor struct {
    ExprStmt1           *ExprStmt
    ExprStmt2           *ExprStmt
    Expr                *Expr
    Stmt                *Stmt
}

type StmtDo struct {
    Expr                *Expr
    Stmt                *Stmt
}

type StmtSelection struct {
    Expr                *Expr
    Stmt                *Stmt
    Stmt2               *Stmt
}

type StmtStep struct {
    Expr                *Expr
    CmpndStmtStimulus   *CmpndStmt
    CmpndStmtResponse   *CmpndStmt
    Identifier          *Identifier
}

var mImportList             *ImportList
var mExpr                   *Expr
var mPackageDef             *PackageDef
var mExtDecl                *ExtDecl
var mParamDecl              *ParamDecl
var mParamDeclList          []ParamDecl
var mFuncDef                *FuncDef
var mExprCmd                *ExprCmd
var mCmdListApp             *CmdListApp
var mCmdMakeAppTid          *CmdMakeAppTid
var mCmpndStmt              *CmpndStmt
var mExprAssign             *ExprAssign
var mExprStmt               *ExprStmt
var mStmt                   *Stmt
var mStmtList               *StmtList
var mCmdMakeAppUser         *CmdMakeAppUser
var mCmdVerify              *CmdVerify
var mCmdGenData             *CmdGenData
var mExprPrimary            *ExprPrimary
var mExprPostfix            *ExprPostfix
var mExprUnary              *ExprUnary
var mInitializer            *Initializer
var mDeclDirect             *DeclDirect
var mIdentifier             *Identifier
var mDeclarator             *Declarator
var mDeclInit               *DeclInit
var mDeclInitList           *DeclInitList
var mDeclaration            *Declaration
var mDeclarationList        *DeclarationList
var mExprCast               *ExprCast
var mExprMultiplicative     *ExprMultiplicative
var mExprAdditive           *ExprAdditive
var mExprShift              *ExprShift
var mExprRelational         *ExprRelational
var mExprEquality           *ExprEquality
var mExprAnd                *ExprAnd
var mExprExclOr             *ExprExclOr
var mExprInclOr             *ExprInclOr
var mExprLogicAnd           *ExprLogicAnd
var mExprLogicOr            *ExprLogicOr
var mExprConditional        *ExprConditional
var mSpecDecl               *SpecDecl
var mSpecType               *SpecType
var mDeclExternal           *DeclExternal
var mTranslationUnit        *TranslationUnit
var mImportDef              *ImportDef
var mExprConstant           *ExprConstant
var mStmtJump               *StmtJump
var mExprListArgument       *ExprListArgument
var mStmtIteration          *StmtIteration
var mStmtSelection          *StmtSelection
var mStmtWhile              *StmtWhile
var mStmtDo                 *StmtDo
var mStmtFor                *StmtFor
var mStmtStep               *StmtStep
var mInitializerList        *InitializerList
var mUnaryOperator          *UnaryOperator
var mParamList              *ParamList
var mParamTypeList          *ParamTypeList
var mIdentifierList         *IdentifierList
var mCmdTestRoute           *CmdTestRoute
var mCmdCreateApp           *CmdCreateApp
var mCmdDeleteApp           *CmdDeleteApp
var mCmdDeleteAppTid        *CmdDeleteAppTid
var mCmdRunSQL              *CmdRunSQL
var mCmdDeleteAppUser       *CmdDeleteAppUser