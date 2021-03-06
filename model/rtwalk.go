package model

import (
    //"fmt"
)

// ---------------------------------------------------
//
//
type RtDeclInit struct {
    Identifier      *Identifier
    Initializer     *Initializer
}
func NewRtDeclInit() *RtDeclInit {
    v := new(RtDeclInit)
    return v
}


// ---------------------------------------------------
//
//
type RtDeclaration struct {
    List        []RtDeclInit
    SpecType    *SpecType
}
func NewRtDeclaration() *RtDeclaration {
    v := new(RtDeclaration)
    v.List = []RtDeclInit{}
    v.SpecType = new(SpecType)
    return v
}


// ---------------------------------------------------
//
//
type RtCmpndStmt struct {
    List        []RtDeclaration
    StmtList    []Stmt
}
func NewRtCmpndStmt() *RtCmpndStmt {
    v := new(RtCmpndStmt)
    v.List = []RtDeclaration{}
    v.StmtList = []Stmt{}
    return v
}
// ---------------------------------------------------
//
//
type RtImportList struct {
    List        []string
}

func NewRtImportList() *RtImportList {
    v := new(RtImportList)
    v.List = []string{}
    return v
}

// ---------------------------------------------------
//
//
type RtFuncDef struct {
    Name            string
    RtCmpndStmt     *RtCmpndStmt
    Identifier      *Identifier
    SpecType        *SpecType
    ParamDeclList   []ParamDecl       
}
func NewRtFuncDef() *RtFuncDef {
    v := new(RtFuncDef)
    return v
}
// ---------------------------------------------------
//
//
type RtPackageDef struct {
    Name    string
}

func NewRtPackageDef() *RtPackageDef {
    v := new(RtPackageDef)
    return v
}

// ---------------------------------------------------
//
//
type RtTranslationUnit struct {
    RtPackageDef    *RtPackageDef
    RtFuncDefList    []RtFuncDef
    RtImportList    *RtImportList
}

func NewRtTranslationUnit() *RtTranslationUnit {
    v := new(RtTranslationUnit)
    v.RtPackageDef = NewRtPackageDef()
    v.RtFuncDefList = []RtFuncDef{}
    v.RtImportList = NewRtImportList()
    return v
}

func (rtu *RtTranslationUnit) GetFuncDef(funcDefName string, indent int) *RtFuncDef {
    mRtFuncDefList := rtu.RtFuncDefList
    for i := range mRtFuncDefList {
        mRtFuncDef := mRtFuncDefList[i]
        if mRtFuncDef.Identifier.Name == funcDefName {
            rlog5(indent, "Found mRtFuncDef: %v \n", mRtFuncDef.Identifier.Name)
            return &mRtFuncDef
        }
    }
    return nil
}

var mRtTranslationUnit  *RtTranslationUnit
var mMainTranslation    *RtTranslationUnit

// ---------------------------------------------------
//
//
type RtTranslationUnitList struct {
    RtTranslationUnits  []RtTranslationUnit
}

var mRtTranslationUnitList *RtTranslationUnitList

func NewRtTranslationUnitList() *RtTranslationUnitList {
    if mRtTranslationUnitList == nil {
        v := new(RtTranslationUnitList)
        v.RtTranslationUnits = []RtTranslationUnit{}
        mRtTranslationUnitList = v
    }
    return mRtTranslationUnitList
}

func (tul *RtTranslationUnitList) Clear() {
    tul.RtTranslationUnits = nil
    mRtTranslationUnitList = nil
}


func (tul *RtTranslationUnitList) AddRtTranslationUnit(tu *RtTranslationUnit) {
    tul.RtTranslationUnits = append(tul.RtTranslationUnits, *tu)
}

func (tul *RtTranslationUnitList) GetTranslationUnit(packageName string, funcDefName string, indent int) *RtTranslationUnit {
    rlog5(indent,"RtTranslationUnits count=%v\n", len(tul.RtTranslationUnits))
    for i := range tul.RtTranslationUnits {
        mUnit := tul.RtTranslationUnits[i]
        mPackageDef := mUnit.RtPackageDef
        rlog5(indent,"mPackageDef Name=%v packageName=%v\n",mPackageDef.Name ,packageName)
        if mPackageDef.Name == packageName {
            for j := range mUnit.RtFuncDefList {
                mRtFuncDef := mUnit.RtFuncDefList[j]
                rlog5(indent, "mRtFuncDef: %v \n", mRtFuncDef.Identifier.Name)
                if mRtFuncDef.Identifier.Name == funcDefName {
                    return &mUnit
                }
            }
        }
    }
    return nil
}

// ---------------------------------------------------
//
//
/*
func ReParseTree() {
    RtWalkTranslationUnit(mTranslationUnit, 1)
}
*/

func RtWalkTranslationUnit(node *TranslationUnit, indent int) *RtTranslationUnit {
    if node != nil {
        mRtTranslationUnit = NewRtTranslationUnit()
        RtWalkTranslationUnit(node.TranslationUnit, indent+1)
        RtWalkDeclExternal(node.DeclExternal, indent+1, mRtTranslationUnit)
        return mRtTranslationUnit
    }    
    return nil
}

func RtWalkDeclExternal(node *DeclExternal, indent int, tu *RtTranslationUnit) {
    if node != nil {
        
        rlog(indent,"DeclExternal: %v \n", node )
        
        rlog(indent,"TranslationUnit: %v \n", tu )
        RtWalkPackageDef(node.PackageDef, indent+1)
        RtWalkImportDef(node.ImportDef, indent+1)
        mRtFuncDef := RtWalkFuncDef(node.FuncDef, indent+1)
        if mRtFuncDef != nil {
            
            rlog(indent,"RtWalkDeclExternal>>>>+++>>>mRtFuncDef: %v \n", mRtFuncDef )
            tu.RtFuncDefList = append(tu.RtFuncDefList, *mRtFuncDef)
        }
    }    
}

func RtWalkPackageDef(node *PackageDef, indent int) {
    if node != nil {
        
        rlog(indent,"PackageDef: %v \n", node )
        mRtTranslationUnit.RtPackageDef.Name = node.Name
        RtWalkImportDef(node.ImportDef, indent+1)
        RtWalkFuncDef(node.FuncDef, indent+1)
        
        rlog(indent,"RtWalkPackageDef>>..>>FuncDef: %v \n", node.FuncDef )
    }    
}

func RtWalkImportDef(node *ImportDef, indent int) {
    if node != nil {
        
        rlog(indent,"ImportDef: %v \n", node )
        RtWalkImportList(node.ImportList, indent+1)
    }    
}

func RtWalkImportList(node *ImportList, indent int) {
    if node != nil {
        
        rlog(indent,"ImportList: %v \n", node.List )
        mRtTranslationUnit.RtImportList.List = node.List
    }    
}

func RtWalkFuncDef(node *FuncDef, indent int) *RtFuncDef {
    if node != nil {
        
        rlog(indent,"FuncDef 2: %v \n", node )
        mRtFuncDef := new(RtFuncDef)
        mRtFuncDef.ParamDeclList = []ParamDecl{}
        mRtFuncDef.Name = node.Name
        mRtFuncDef.SpecType = RtWalkSpecDecl(node.SpecDecl, indent+1)
        RtWalkFuncDef_Declarator(node.Declarator, indent+1, mRtFuncDef, 0)
        mRtFuncDef.RtCmpndStmt = RtWalkCmpndStmt(node.CmpndStmt, indent+1)
        
        rlog(indent,"mRtFuncDef>>>>>>>ParamDeclList: %v \n", mRtFuncDef.ParamDeclList )
        
        rlog(indent,"mRtFuncDef.RtCmpndStmt: %v\n", mRtFuncDef.RtCmpndStmt )
        return mRtFuncDef
    }    
    return nil
}

func RtWalkFuncDef_Declarator(node *Declarator, indent int, mRtFuncDef *RtFuncDef, disp int) {
    if node != nil {
        
        rlog(indent,"Declarator 2: %v %v %v\n", node, indent, disp )
        RtWalkFuncDef_DeclDirect(node.DeclDirect, indent+1, mRtFuncDef, disp+1)
    }    
}
func RtWalkFuncDef_DeclDirect(node *DeclDirect, indent int, mRtFuncDef *RtFuncDef, disp int) {
    if node != nil {
        
        rlog(indent,"DeclDirect 2: %v %v %v\n", node, indent, disp )
        RtWalkFuncDef_DeclDirect(node.DeclDirect, indent+1, mRtFuncDef, disp+1)
        RtWalkFuncDef_Identifier(node.Identifier, indent+1, mRtFuncDef, disp+1)
        RtWalkFuncDef_ParamTypeList(node.ParamTypeList, indent+1, mRtFuncDef, disp+1)
    }    
}
func RtWalkFuncDef_ParamTypeList(node *ParamTypeList, indent int, mRtFuncDef *RtFuncDef, disp int) {
    if node != nil {
        
        rlog(indent,"ParamTypeList 2: %v %v %v\n", node, indent, disp )
        RtWalkFuncDef_ParamList(node.ParamList, indent+1, mRtFuncDef, disp+1)
    }    
}
func RtWalkFuncDef_Identifier(node *Identifier, indent int, mRtFuncDef *RtFuncDef, disp int) {
    if node != nil {
        
        rlog(indent,"Identifier 2: %v %v %v\n", node, indent, disp )
        if disp==3 {
            mRtFuncDef.Identifier = node
        }
    }    
}
func RtWalkFuncDef_ParamList(node *ParamList, indent int, mRtFuncDef *RtFuncDef, disp int) {
    if node != nil {
        
        rlog(indent,"ParamList 2: %v %v %v\n", node, indent, disp )
        RtWalkFuncDef_ParamDecl(node.ParamDecl, indent+1, mRtFuncDef, disp+1)
        RtWalkFuncDef_ParamList(node.ParamList, indent+1, mRtFuncDef, disp+1)
    }    
}
func RtWalkFuncDef_ParamDecl(node *ParamDecl, indent int, mRtFuncDef *RtFuncDef, disp int) {
    if node != nil {
        
        rlog(indent,"ParamDecl 2: %v %v %v\n", node, indent, disp )
        mRtFuncDef.ParamDeclList = append(mRtFuncDef.ParamDeclList, *node)
        
        rlog(indent,"ParamDeclList 2: %v %v %v %v\n", mRtFuncDef.ParamDeclList, indent, disp, node.Declarator.DeclDirect.Identifier )
        //RtWalkFuncDef_Declarator(node.Declarator, indent+1, mRtFuncDef, disp+1)
        //RtWalkFuncDef_SpecDecl(node.SpecDecl, indent+1, mRtFuncDef, disp+1)
    }    
}
func RtWalkFuncDef_SpecDecl(node *SpecDecl, indent int, mRtFuncDef *RtFuncDef, disp int) {
    if node != nil {
        
        rlog(indent,"SpecDecl 2: %v %v %v\n", node, indent, disp )
    }    
}

//
// -----------------------------------------------------------------
//

func RtWalkSpecDecl(node *SpecDecl,indent int) *SpecType {
    if node != nil {
        
        rlog(indent,"SpecDecl1: %v\n", node )
        DumpSpecType(node.SpecType,indent+1)
        return node.SpecType
    }
    return nil
}

func RtWalkDeclarator(node *Declarator, indent int) *Identifier {
    if node != nil {
        mDeclarator := node
        
        rlog(indent,"Declarator1: %v\n", node )
        mIdentifier := RtWalkDeclDirect(mDeclarator.DeclDirect, indent+1)
        return mIdentifier
    }
    return nil
}

func RtWalkDeclDirect(node *DeclDirect, indent int) *Identifier {
    if node != nil {
        
        rlog(indent,"DeclDirect1: %v\n", node )
        mIdentifier = RtWalkIdentifier(node.Identifier, indent+1)
        if node.DeclDirect != nil {
            mIdentifier = RtWalkDeclDirect(node.DeclDirect , indent+1)
        }
        return mIdentifier
    }    
    return nil
}

func RtWalkIdentifier(node *Identifier, indent int) *Identifier {
    if node != nil {
        
        rlog(indent,"Identifier1: %v\n", node )
        return node
    }    
    return nil
}


func RtWalkCmpndStmt(node *CmpndStmt, indent int) *RtCmpndStmt {
    if node != nil {
        mCmpndStmt = node
        
        rlog(indent,"CmpndStmt: %v %v %v\n", node, mCmpndStmt.StmtList, mCmpndStmt.DeclarationList )
        mRtCmpndStmt := NewRtCmpndStmt()
        mRtCmpndStmt.List = []RtDeclaration{}
        RtWalkDeclarationList(node.DeclarationList,indent+1, mRtCmpndStmt)
        DumpRtStmtList(node.StmtList ,indent+1, mRtCmpndStmt)
        //mRtCmpndStmt.StmtList = node.StmtList
        
        rlog(indent,"RtWalkCmpndStmt>>mRtCmpndStmt.List: %v\n", mRtCmpndStmt)
        return mRtCmpndStmt
    }    
    return nil
}

func DumpRtStmtList(node *StmtList, indent int, mRtCmpndStmt *RtCmpndStmt) {
    if node != nil {
        
        rlog(indent,"StmtList: %v\n", node)
        DumpRtStmt(node.Stmt, indent+1, mRtCmpndStmt)
        DumpRtStmtList(node.StmtList, indent+1, mRtCmpndStmt)
    }    
}

func DumpRtStmt(node *Stmt, indent int, mRtCmpndStmt *RtCmpndStmt) {
    if node != nil {
        
        rlog(indent,"Stmt: %v\n", node)
        mRtCmpndStmt.StmtList = append(mRtCmpndStmt.StmtList, *node)
    }    
}


func RtWalkDeclarationList(node *DeclarationList, indent int, mRtCmpndStmt *RtCmpndStmt) {
    if node != nil {
        
        rlog(indent,"DeclarationList: %v\n", node )
        RtWalkDeclarationList(node.DeclarationList, indent+1, mRtCmpndStmt)
        mDeclaration := RtWalkDeclaration(node.Declaration, indent+1 )
        if mDeclaration != nil {
            mRtCmpndStmt.List = append(mRtCmpndStmt.List, *mDeclaration)
        }
    }
}

func RtWalkDeclaration(node *Declaration, indent int) *RtDeclaration {
    if node != nil {
        
        rlog(indent,"Declaration: %v\n", node )
        mRtDeclaration := NewRtDeclaration()
        mRtDeclaration.SpecType = RtWalkSpecDecl(node.SpecDecl,indent+1)
        RtWalkDeclInitList(node.DeclInitList,indent+1,mRtDeclaration)
        
        rlog(indent,"mRtDeclaration.List: %v\n", mRtDeclaration.List)
        return mRtDeclaration
    }
    return nil
}

func RtWalkDeclInitList(node *DeclInitList, indent int, rtd *RtDeclaration) {
    if node != nil {
        
        rlog(indent,"DeclInitList: %v\n", node )
        mRtDeclInit := RtWalkDeclInit(node.DeclInit, indent+1)
        if mRtDeclInit != nil {
            rtd.List = append(rtd.List, *mRtDeclInit)
        }
        RtWalkDeclInitList(node.DeclInitList, indent+1, rtd)
        
        rlog(indent,"RtDeclarationList: %v\n", rtd.List )
    }    
}

func RtWalkDeclInit(node *DeclInit, indent int) *RtDeclInit {
    if node != nil {
        
        rlog(indent,"DeclInit: %v\n", node )
        mRtDeclInit := new(RtDeclInit)
        mRtDeclInit.Identifier = RtWalkDeclDirect(node.Declarator.DeclDirect, indent+1)
        mRtDeclInit.Initializer = node.Initializer
        
        rlog(indent,"DeclInit.Initializer: %v\n", node.Initializer )
        return mRtDeclInit
    }
    return nil
}

// ---------------------------------------------------
//
//




