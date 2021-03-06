package model

import (
    "fmt"
    _ "github.com/lib/pq"
    "database/sql"
    "intel/test/selectorproperty"
    "os"
    "io"
    "bytes"

)

type App struct {
    App_id  string
    Appname string
}


var Db                          *sql.DB
var DbHub                       *sql.DB

var DbConnectionStringHub       string
var ClientDbConnectionString    string
var err                         error
var currentApp              string
var currentAppId            string

var selectorList            []string
var arrApps                 []App

var host                    string = "http://localhost:8080/"
var rootPath                string
var lineCounter             int
var executionMode           byte
var symObjectListId         int
var symObjectId             int
var mapPropertyId           int
var selectorRefList *selectorproperty.SelectorRefList 

func StartIt(fname string, indent int) {
    mRtTranslationUnitList = NewRtTranslationUnitList()

    s := LoadFile(fname)
    mRtTranslationUnit = Parse(s)
    mMainTranslation = mRtTranslationUnit
    mRtTranslationUnitList.AddRtTranslationUnit(mRtTranslationUnit)
    HandleImport(mRtTranslationUnit,indent+1)
    ValidateImports(indent+1)
    rlog(indent,"--------------------------------- E X E C U T E ---------------------\n")
    RtExecute(mMainTranslation)
    rlog5(indent,"--------------------------------- TEST PLAN---------------------\n")
    mList := GetRuntimeTestPlanList()
    sOut := mList.EmitTestPlanResultToHtml2()
    //rlog5(indent,"EmitTestPlan: \n%v\n", sOut)
    f, _ := os.Create(fmt.Sprintf("%vreports/test1.html",GetRootPath()))
    defer f.Close()
    f.WriteString(sOut)
    Unload(mRtTranslationUnitList)
}



func LoadFile(filepath string) string {
    rootPath = GetRootPath()
    loadpath := filepath
    fullpath := fmt.Sprintf("%s%s",rootPath, loadpath)
    if FileExists(fullpath) {
        return fullpath
    }
    return ""
}

func Parse(s string) *RtTranslationUnit {
    indent := 1
    if len(s) > 0 {
        prvLexer = new(PrvLexer)
        buf := bytes.NewBuffer(nil)
        f, _ := os.Open(s) 
        io.Copy(buf, f) 
        f.Close()
        sbuf := string(buf.Bytes())
        rlog(indent,"BUFFER:%v\n\n\n", sbuf)
        prvLexer.s = sbuf
        prvLexer.pos = 0
        currState = 0
        commandState = 0
        lineCounter = 0
        yyParse(prvLexer)
        rlog(indent,"Parse completed.\n")
        rlog(indent,"--------------------------------- D U M P   A S T ---------------------\n")
        mTU := DoDumpTree(mTranslationUnit, 1)
        rlog(indent,"--------------------------------- D U M P   R T     A S T ---------------------\n")
        mRtu := RtWalkTranslationUnit(mTU, 1)
        rlog(indent,"--------------------------------- D U M P   R T    T R E E ---------------------\n")
        RtDump(mRtu,1)
        return mRtu
    }
    return nil
}


func FileExists(name string) bool {
    _, err := os.Stat(name)
    return !os.IsNotExist(err)
}

// ---------------------------------------------------------------------

func RtDump(node *RtTranslationUnit, indent int) {
    if node != nil {
        rlog(indent,"RtTranslationUnit: %v \n", node)
        DumpRtPackageDef(node.RtPackageDef, indent+1)
        DumpRtImportList(node.RtImportList, indent+1)
        DumpRtFuncDefList(node.RtFuncDefList, indent+1)
    }    
}

func DumpRtPackageDef(node *RtPackageDef, indent int) {
    if node != nil {
        
        rlog(indent,"RtPackageDef: %v\n", node)
    }    
}

func DumpRtImportList(node *RtImportList, indent int) {
    if node != nil {
        
        rlog(indent,"RtImportList: %v\n", node)
    }    
}

func DumpRtFuncDefList(nodeList []RtFuncDef, indent int) {
    if nodeList != nil {
        for i := range nodeList {
            node := nodeList[i]
            DumpRtFuncDef(&node, indent+1)
        }
    }    
}

func DumpParamDeclList(list []ParamDecl , indent int) {
    if len(list) > 0  {
        
        rlog(indent,"ParamDeclList: %v \n", list)
    }
}

func DumpRtFuncDef(node *RtFuncDef, indent int) {
    if node != nil {
        
        rlog(indent,"RtFuncDef: node=%v, list=%v \n", node, node.ParamDeclList)
        rlog(indent,"Name: %v \n", node.Identifier.Name)
        DumpRtCmpndStmt(node.RtCmpndStmt, indent+1)
        DumpIdentifier(node.Identifier, indent+1)
        DumpSpecType(node.SpecType, indent+1)
        DumpParamDeclList(node.ParamDeclList, indent+1)
    }    
}

func DumpRtCmpndStmt(node *RtCmpndStmt, indent int) {
    if node != nil {
        
        rlog(indent,"RtCmpndStmt: %v \n", node)
        DumpRtDeclarationList(node.List, indent+1)
        for i := range node.StmtList {
            mStmt := node.StmtList[i]
            DumpStmt(&mStmt, indent+1)
        }
    } else {
        
        rlog(indent,"RtCmpndStmt: nil\n")
    }   
}

func DumpRtDeclarationList(nodeList []RtDeclaration, indent int) {
    if nodeList != nil {
        
        rlog(indent,"RtDeclarationList: %v \n", nodeList)
        for i := range nodeList {
            node := nodeList[i]
            DumpRtDeclaration(&node, indent+1)
        }
    }    
}

func DumpRtDeclaration(node *RtDeclaration, indent int) {
    if node != nil {
        
        rlog(indent,"RtDeclaration: %v \n", node)
        DumpSpecType(node.SpecType, indent+1)
        for i := range node.List {
            mRtDeclInit := node.List[i]
            //
            //rlog(indent,"mRtDeclInit[%v] %v \n", i, mRtDeclInit)
            DumpRtDeclInit(&mRtDeclInit, indent+1)
        }
    }    
}

func DumpRtDeclInit (node *RtDeclInit , indent int) {
    if node != nil {
        
        rlog(indent,"RtDeclInit : %v \n", node)
        DumpIdentifier(node.Identifier, indent+1)
        //DumpInitializer(node.Initializer, indent+2)
    }    
}

// ---------------------------------------------------
//
//
type StackRuntime struct {
    Frames      []Frame
}

func NewStackRuntime() *StackRuntime {
    v := new(StackRuntime)
    v.Frames = []Frame{}
    return v
}

func (s *StackRuntime) Push(f *Frame) {
    s.Frames = append(s.Frames, *f)
    rlog(1,"Push: %v\n", f.Name)
}

func (s *StackRuntime) Pop() {
    f := s.Frames[len(s.Frames)-1]
    rlog(1,"Popping: %v \n", f.Name)
    if len(s.Frames) > 0 {
        f.UnloadFrame(&s.Frames[len(s.Frames)-1])
    }
    s.Frames = s.Frames[:len(s.Frames)-1]
    if len(s.Frames) > 0 {
        f = s.Frames[len(s.Frames)-1]
        rlog(1,"Current: %v  \n", f.Name)
    } else {
        rlog(1,"Current: empty\n")
    }
}

// ---------------------------------------------------
//
//

type LocalVariable struct {
    Identifier  *Identifier
    SpecType    *SpecType
    Value       string
    IntValue    int
    FloatValue  float64
    StringValue string
    BoolValue   bool
}

func NewLocalVariable() *LocalVariable {
    v := new(LocalVariable)
    return v
}


// ---------------------------------------------------
//
//
type Frame struct {
    Name                    string
    LocalVariableList       []LocalVariable
    IP                      int
    RtFuncDef               *RtFuncDef
    StmtList                []Stmt
    ParamCollection         []NodeEval
    RtTranslationUnit       *RtTranslationUnit
}

func NewFrame(name string, pFuncDef *RtFuncDef) *Frame {
    v := new(Frame)
    v.Name = name
    v.RtFuncDef = pFuncDef
    v.StmtList = pFuncDef.RtCmpndStmt.StmtList
    v.LocalVariableList = []LocalVariable{}
    v.ParamCollection = []NodeEval{}
    return v
}

func (f *Frame) UnloadFrame(v *Frame) {
    v.RtFuncDef = nil
    v.StmtList = nil
    //for i := range v.LocalVariableList {
    //    itm := v.LocalVariableList[i]
    //    itm.Value = ""
    //}
    v.LocalVariableList = []LocalVariable{}
    v.ParamCollection = []NodeEval{}
}

func (f *Frame) AddParam(node *NodeEval) {
    f.ParamCollection = append(f.ParamCollection, *node)
}

func (f *Frame) ClearParamCollection() {
    f.ParamCollection = []NodeEval{}
}

func (f *Frame) PackageCurrentCollection(indent int) []NodeEval {
    mLocalVariableList := []NodeEval{}
    for i := range f.ParamCollection {
        mNodeEval := f.ParamCollection[i]
        if mNodeEval.IdType == IDENTIFIER {
            mLocalVar := f.GetLocal(mNodeEval.Value,indent)
            mNewNode := NewNodeEval()
            mNewNode.Value = mLocalVar.Identifier.Name
            mNewNode.IdType = mLocalVar.SpecType.Id
            mNewNode.Value = mLocalVar.Value
            mLocalVariableList = append(mLocalVariableList, *mNewNode)
            rlog(indent,"PackageCurrentCollection: IDENTIFIER mLocalVar=%v mNewNode=%v \n", mLocalVar, mNewNode)
        } else {
            mLocalVariableList = append(mLocalVariableList, mNodeEval)
            rlog(indent,"PackageCurrentCollection: NON IDENTIFIER mNodeEval=%v  \n", mNodeEval)
        }
    }
    rlog(indent,"PackageCurrentCollection: mLocalVariableList=%v  \n", mLocalVariableList)
    return mLocalVariableList
}


func (f *Frame) GetLocal(lname string, indent int) *LocalVariable {
    rlog(indent, "GetLocal: %v \n", lname)
    for i := range f.LocalVariableList {
        local := f.LocalVariableList[i]
        if lname == local.Identifier.Name {
            rlog3(indent,"GetLocal Identifier=%v\n", local.Identifier.Name)
            rlog3(indent+2,"GetLocal Value=%v\n", local.Value)
            return &local
        }
    }    
    rlog(indent,".... not found \n")
    return nil
}

func (f *Frame) SetLocal(localVar *LocalVariable, indent int)  {
    rlog(indent, "SetLocal: \n")
    for i := range f.LocalVariableList {
        local := f.LocalVariableList[i]
        if local.Identifier.Name  == localVar.Identifier.Name {
            rlog(indent, " found %v %v %v %v %v \n", localVar.Identifier, localVar.BoolValue, localVar.Value, local.SpecType, local )
            f.LocalVariableList[i] = *localVar
            rlog3(indent,"SetLocal Identifier=%v\n", localVar.Identifier.Name)
            rlog3(indent+2,"SetLocal Value=%v\n", localVar.Value)
            return 
        } 
    }    
    rlog(indent, ".... not found %v = %v \n", localVar.Identifier.Name, localVar.Value)
}

func (f *Frame) SetupParams(params []NodeEval, indent int) {
    if len(params) == 0 || len(f.RtFuncDef.ParamDeclList) == 0 {
        return
    }
    if len(f.RtFuncDef.ParamDeclList) != len(params) {
        panic("Arguments do not match parameters")
    }
    for i := range f.RtFuncDef.ParamDeclList {
        j := len(f.RtFuncDef.ParamDeclList) - i - 1
        mParamDecl := f.RtFuncDef.ParamDeclList[j]
        mIdentifier := mParamDecl.Declarator.DeclDirect.Identifier
        mSpecType := mParamDecl.SpecDecl.SpecType
        paramArg := params[i]
        mLocal := f.GetLocal(mIdentifier.Name, indent)
        if mLocal.SpecType.Id == LITERAL_STRING || mLocal.SpecType.Id == TYPE_STRING_ID {
            if paramArg.IdType == LITERAL_STRING || paramArg.IdType == TYPE_STRING_ID {
                mLocal.Value = paramArg.Value
                mLocal.StringValue = paramArg.StringValue
                f.SetLocal(mLocal, indent)
            } else {
                panic("Parameter type not matching.")
            }
        }
        
        rlog(indent,"SetupParams:  mSpecType=%v mIdentifier=%v paramArg=%v mLocal=%v \n", mSpecType, mIdentifier,paramArg, mLocal)
    }
}


func (f *Frame) PopulateFunctionArguments() {
    for i := range f.RtFuncDef.ParamDeclList {
        mParamDecl := f.RtFuncDef.ParamDeclList[i]
        mIdentifier = mParamDecl.Declarator.DeclDirect.Identifier
    }
}

func (f *Frame) InitLocals() {
    indent := 1
    rlog3(indent,"-------------------------------- INIT LOCALS BEGIN---------------------------------\n")
    for i := range f.RtFuncDef.RtCmpndStmt.List {
        mRtDeclaration := f.RtFuncDef.RtCmpndStmt.List[i]
        mSpecType = mRtDeclaration.SpecType
        rlog3(indent,"mSpecType=%v \n", mSpecType.Id)
        //rlog3(indent,"mRtDeclaration.List[%v]:=%v \n", i,  mRtDeclaration)
        for j := range mRtDeclaration.List {
            mRtDeclInit := mRtDeclaration.List[j]
            mIdentifier := mRtDeclInit.Identifier
            //mInitializer := mRtDeclInit.Initializer
            mLocalVariable := NewLocalVariable()
            mLocalVariable.Identifier = mIdentifier
            mLocalVariable.SpecType = mSpecType
            rlog3(indent+1,"Adding mIdentifier=%v \n", mIdentifier.Name)
            f.LocalVariableList = append(f.LocalVariableList, *mLocalVariable)
        }
    }
    rlog3(indent,"-------------------------------- PARAM DECL---------------------------------\n")
    for i := range f.RtFuncDef.ParamDeclList {
        mParamDecl := f.RtFuncDef.ParamDeclList[i]
        mLocalVariable := NewLocalVariable()
        mLocalVariable.Identifier = mParamDecl.Declarator.DeclDirect.Identifier
        mLocalVariable.SpecType = mParamDecl.SpecDecl.SpecType
        f.LocalVariableList = append(f.LocalVariableList, *mLocalVariable)
        rlog3(indent,"Identifier=%v \n", mLocalVariable.Identifier.Name)
        rlog3(indent+1,"SpecType=%v \n", mLocalVariable.SpecType.Id)
    }

    rlog3(indent,"-------------------------------- LOCAL VARS---------------------------------\n")
    for i := range f.LocalVariableList {
        local := f.LocalVariableList[i]
        rlog3(indent,"Identifier=%v \n", local.Identifier.Name)
        rlog3(indent+1,"SpecType=%v \n", local.SpecType.Id)
    }
    rlog3(indent,"-------------------------------- INIT LOCALS END---------------------------------\n")
}

func (f *Frame) HandleTestPlanTree(rfd *RtFuncDef, indent int) {
    rlog5(indent,"HandleTestPlanTree: %v\n", rfd)
    if rfd.SpecType != nil {
        if rfd.SpecType.Id == TYPE_TESTPLAN_ID {
            rlog5(indent,"HandleTestPlanTree Add Test Plan: %v\n", rfd.Identifier)
            rtp := GetRuntimeTestPlanList()
            rlog5(indent,"HandleTestPlanTree rtp: %v\n", rtp)
            tp := NewTestPlan()
            tp.Name = rfd.Identifier.Name
            localVarText := f.GetLocal("testPlanText",indent+1)
            if localVarText != nil {
                tp.TestPlanText = localVarText.Value
                rlog6(indent,"HandleTestPlanTree testPlanText: %v\n", localVarText.Value)
            }
            rtp.AddTestPlan(tp)
            rlog5(indent,"HandleTestPlanTree tp: %v\n", tp)
        }
    }
}

func (f *Frame) Run(params []NodeEval, rfd *RtFuncDef, indent int) *NodeEval {
    e := NewNodeEval()
    f.InitLocals()
    f.SetupParams(params,1)
    f.HandleTestPlanTree(rfd, indent+1)
    for i := range f.StmtList {
        node := f.StmtList[len(f.StmtList)-1-i]
        indent := 1
        val1 := f.EvalFrameExprCmd(node.ExprCmd, indent+1)
        val2 := f.EvalFrameExprStmt(node.ExprStmt, indent+1)
        val3 := f.EvalFrameStmtJump(node.StmtJump, indent+1)
        val4 := f.EvalFrameStmtIteration(node.StmtIteration, indent+1)
        val5 := f.EvalFrameStmtSelection(node.StmtSelection, indent+1)
        val6 := f.EvalFrameCmpndStmt(node.CmpndStmt, indent+1)
        val7 := f.EvalFrameStmtStep(node.StmtStep, indent+1)
        
        rlog(indent,"RunFrame vals: ExprCmd=%v ExprStmt=%v StmtJump=%v StmtIteration=%v StmtSelection=%v CmpndStmt=%v StmtStep=%v\n", val1, val2, val3, val4, val5, val6, val7)
        if val1.IdType > 0 {
            e = val1
            rlog(indent,"RunFrame ExprCmd: %v\n", e)
        }
        if val2.IdType > 0 {
            e = val2
            rlog(indent,"RunFrame ExprStmt: %v\n", e)
        }
        if val3.IdType > 0 {
            e = val3
            rlog(indent,"RunFrame StmtJump: %v\n", e)
            return e
        }
        if val4.IdType > 0 {
            e = val4
            rlog(indent,"RunFrame StmtIteration: %v\n", e)
        }
        if val5.IdType > 0 {
            e = val5
            rlog(indent,"RunFrame StmtSelection: %v\n", e)
        }
        if val6.IdType > 0 {
            e = val6
            rlog(indent,"RunFrame CmpndStmt: %v\n", e)
        }
        if val7.IdType > 0 {
            e = val7
            rlog(indent,"RunFrame StmtStep: %v\n", e)
        }
    }
    return e
}

func (f *Frame) AddVariable(v *LocalVariable) {
    f.LocalVariableList = append(f.LocalVariableList, *v)
}


// ---------------------------------------------------
//
//
var mRuntimeStack *StackRuntime

func RtExecute(rtu *RtTranslationUnit) {
    mRuntimeStack = NewStackRuntime()
    rtu.CallFunc("main", nil)
}

func (rtu *RtTranslationUnit) CallFunc(funcName string, params []NodeEval) *NodeEval {
    if len(mRuntimeStack.Frames) > 12 * 1024 {
        panic("ERROR: Runtime Stack Overflow.")
    }
    e := NewNodeEval()
    for i := range rtu.RtFuncDefList {
        mFuncDef := rtu.RtFuncDefList[i]
        //rlog5(1,"name:  %v \n", mFuncDef.Identifier.Name)
        if funcName == mFuncDef.Identifier.Name {
            rlog(1,"\nFunc Name %v() found. Loading Function \n", funcName)
            frame := NewFrame(funcName, &mFuncDef)
            frame.RtTranslationUnit = rtu
            mRuntimeStack.Push(frame)
            rlog3(1,"==========================================================================================\n")
            rlog3(1,"F U N C T I O N  :   %v  SpecType=%v \n", funcName, mFuncDef.SpecType)
            e = frame.Run(params, &mFuncDef, 1)
            rlog3(1,"\nUnloading Function %v \n", funcName)
            mRuntimeStack.Pop()
            return e
        }
    }
    rlog5(1,"Did not find funcName  %v \n", funcName)
    return e
}

func rlog(indent int, frmt string, v...interface{}  ) {
    return
    DoInd(indent)
    if v != nil {
        fmt.Printf(frmt, v...)
    } else {
        fmt.Printf(frmt)
    }
}

func rlog2(frmt string, v ... interface{}  ) {
    return
    fmt.Printf(frmt, v...)
}
func rlog3(indent int, frmt string, v...interface{}  ) {
    return
    DoInd(indent)
    if v != nil {
        fmt.Printf(frmt, v...)
    } else {
        fmt.Printf(frmt)
    }
}
func rlog4(indent int, frmt string, v...interface{}  ) {
    return
    DoInd(indent)
    if v != nil {
        fmt.Printf(frmt, v...)
    } else {
        fmt.Printf(frmt)
    }
}
func rlog5(indent int, frmt string, v...interface{}  ) {
    return
    DoInd(indent)
    if v != nil {
        fmt.Printf(frmt, v...)
    } else {
        fmt.Printf(frmt)
    }
}
func rlog6(indent int, frmt string, v...interface{}  ) {
    return
    DoInd(indent)
    if v != nil {
        fmt.Printf(frmt, v...)
    } else {
        fmt.Printf(frmt)
    }
}
func rlog7(indent int, frmt string, v...interface{}  ) {
    return
    DoInd(indent)
    if v != nil {
        fmt.Printf(frmt, v...)
    } else {
        fmt.Printf(frmt)
    }
}
func rlog8(indent int, frmt string, v...interface{}  ) {
    return
    DoInd(indent)
    if v != nil {
        fmt.Printf(frmt, v...)
    } else {
        fmt.Printf(frmt)
    }
}
func rlog9(indent int, frmt string, v...interface{}  ) {
    //return
    DoInd(indent)
    if v != nil {
        fmt.Printf(frmt, v...)
    } else {
        fmt.Printf(frmt)
    }
}
