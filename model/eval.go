package model

import (
    "fmt"
    "strconv"
    "strings"
)

type NodeEval struct {
    Value           string
    IdType          int
    IntValue        int
    StringValue     string
    BoolValue       bool
    FloatValue      float64
    ObjectRef       string
    StepDocument    *StepDocument
}

func NewNodeEval() *NodeEval {
    v := new(NodeEval)
    
    return v
}

func (f *Frame) EvalFrameExprStmt(node *ExprStmt, indent int) *NodeEval  {
    e := NewNodeEval()
    if node != nil {
        rlog3(indent,"EvalFrameExprStmt: %v\n", node)
        val1 := f.EvalFrameExprAssign(node.ExprAssign, indent+1)
        val2 := f.EvalFrameExpr(node.Expr, indent+1)
        //rlog(indent,"EvalFrameExprStmt: ExprAssign=%v Expr=%v \n", val1, val2)
        if val1.IdType > 0 {
            e = val1
            rlog(indent,"EvalFrameExprStmt ExprAssign: %v \n", e)
        }
        if val2.IdType > 0 {
            e = val2
            rlog(indent,"EvalFrameExprStmt Expr: %v \n", e)
        }
    }
    return e
}

func (f *Frame) EvalFrameExprAssign(node *ExprAssign, indent int) *NodeEval  {
    e := NewNodeEval()
    if node != nil {
        rlog6(indent,"EvalFrameExprAssign: node=%v\n", node)
        val1 := f.EvalFrameExprCmd(node.ExprCmd, indent+1)
        val2 := f.EvalFrameExprAssign(node.ExprAssign, indent+1)
        val3 := f.EvalFrameExprUnary(node.ExprUnary, indent+1)
        val4 := f.EvalFrameExprConditional(node.ExprConditional, indent+1)

        rlog6(indent,"EvalFrameExprAssign: ExprCmd 26 %v val1.IdType=%v\n", val1, val1.IdType)
        if len(val1.Value) > 0 {
            e = val1
            rlog6(indent,"EvalFrameExprAssign: ExprCmd 27 %v\n", e)
            //return e
        }

        rlog6(indent,"EvalFrameExprAssign 28: ExprCmd=%v ExprAssign=%v ExprUnary=%v ExprConditional=%v \n", val1, val2, val3, val4)
        if val1.IdType > 0{
            e = val1
            rlog6(indent,"EvalFrameExprAssign: ExprCmd=%v val1.IdType=%v\n", val1, val1.IdType)
        }
        if val4.IdType > 0{
            e = val4
            rlog(indent,"EvalFrameExprAssign: EvalFrameExprConditional=%v \n", val4)
        }
        if val2.IdType > 0 && val3.IdType > 0 {
            localVal := f.GetLocal(val3.Value,indent+1)
            localVal.Value = val2.Value
            rlog6(indent,"EvalFrameExprAssign: ExprUnary=%v SpecType=%v \n", localVal, localVal.SpecType)
            if localVal.SpecType.Id == TYPE_STRING_ID || localVal.SpecType.Id == LITERAL_STRING {
                rlog6(indent,"EvalFrameExprAssign: LITERAL_STRING = %v StringValue=%v Value=%v \n", val2, val2.StringValue, val2.Value)
                if val2.IdType == IDENTIFIER {
                    mLocal := f.GetLocal(val2.Value,indent+1)
                    if mLocal.SpecType.Id == TYPE_STRING_ID || mLocal.SpecType.Id == LITERAL_STRING  {
                        localVal.StringValue = mLocal.StringValue
                    } else {
                        errmsg := "ERROR: Incompatible Type, expecting a string.\n"
                        rlog(indent,errmsg)
                        panic(errmsg)
                    }
                } else {
                    rlog6(indent,"EvalFrameExprAssign: LITERAL_STRING NON IDENTIFIER  1  val2.Value=%v val2.StringValue=%v\n", val2.Value, val2.StringValue)
                    localVal.StringValue = val2.StringValue
                }
                localVal.Value = localVal.StringValue
                rlog6(indent,"EvalFrameExprAssign: LITERAL_STRING = %v \n", localVal)
                f.SetLocal(localVal, indent+1)
            }
            if localVal.SpecType.Id == TYPE_INT_ID {
                rlog6(indent,"EvalFrameExprAssign: TYPE_INT_ID = %v \n", val2)
                if val2.IdType == IDENTIFIER {
                    mLocal := f.GetLocal(val2.Value,indent+1)
                    if mLocal.SpecType.Id == TYPE_INT_ID || mLocal.SpecType.Id == INTEGER_ID {
                        localVal.IntValue = mLocal.IntValue
                    } else {
                        
                        errmsg := "ERROR: Incompatible Type, expecting a int.\n"
                        fmt.Printf(errmsg)
                        panic(errmsg)
                    }
                } else {
                    localVal.IntValue, err = strconv.Atoi(val2.Value)
                }
                localVal.Value = fmt.Sprintf("%v", localVal.IntValue)
                f.SetLocal(localVal, indent+1)
            }
            if localVal.SpecType.Id == TYPE_BOOL_ID {
                newLocal := NewLocalVariable()
                //*newLocal = *localVal
                newLocal.Identifier = new(Identifier)
                newLocal.SpecType = new(SpecType)
                *newLocal.Identifier = *localVal.Identifier
                *newLocal.SpecType = *localVal.SpecType
                if val2.IdType == IDENTIFIER {
                    rlog6(indent,"EvalFrameExprAssign: TYPE_BOOL_ID IDENTIFIER = %v \n", val2)
                    mLocal := f.GetLocal(val2.Value,indent+1)
                    if mLocal.SpecType.Id == TYPE_BOOL_ID {
                        newLocal.BoolValue = mLocal.BoolValue
                    } else {
                        errmsg := "ERROR: Incompatible Type, expecting a boolean.\n"
                        rlog(indent,errmsg)
                        panic(errmsg)
                    }
                } else {
                    newLocal.BoolValue = val2.BoolValue
                }
                newLocal.Value = fmt.Sprintf("%v",newLocal.BoolValue)
                rlog(indent,"EvalFrameExprAssign: TYPE_BOOL_ID : localVal=%v val2.BoolValue=%v \n", newLocal, val2.BoolValue)
                rlog(indent,"EvalFrameExprAssign: Identifier=%v SpecType=%v \n", newLocal.Identifier , newLocal.SpecType)
                rlog(indent,"EvalFrameExprAssign: BoolValue=%v Value=%v \n", newLocal.BoolValue , newLocal.Value)
                f.SetLocal(newLocal, indent+1)
            }
            if localVal.SpecType.Id == FLOAT_ID || localVal.SpecType.Id == TYPE_FLOAT_ID {
                newLocal := NewLocalVariable()
                newLocal.Identifier = new(Identifier)
                newLocal.SpecType = new(SpecType)
                *newLocal.Identifier = *localVal.Identifier
                *newLocal.SpecType = *localVal.SpecType
                if val2.IdType == IDENTIFIER {
                    rlog6(indent,"EvalFrameExprAssign3: IDENTIFIER \n")
                    mLocal := f.GetLocal(val2.Value,indent+1)
                    newLocal.FloatValue = mLocal.FloatValue
                    newLocal.Value = fmt.Sprintf("%v",mLocal.FloatValue)
                } else {
                    if len(val2.Value) > 0 {
                        rlog6(indent,"EvalFrameExprAssign3: NON IDENTIFIER  1 Value=%v Float=%v\n", val2.Value, val2.FloatValue)
                        newLocal.FloatValue, _ = strconv.ParseFloat(val2.Value,64)
                    } else {
                        rlog6(indent,"EvalFrameExprAssign3: NON IDENTIFIER  2\n")
                        newLocal.FloatValue = val2.FloatValue
                    }
                    newLocal.FloatValue = val2.FloatValue
                    newLocal.Value = fmt.Sprintf("%v",newLocal.FloatValue)
                }
                rlog(indent,"EvalFrameExprAssign3: TYPE_BOOL_ID : localVal=%v val2.FloatValue=%v val2.Value=%v \n", newLocal, val2.FloatValue, val2.Value)
                rlog(indent,"EvalFrameExprAssign3: Identifier=%v SpecType=%v \n", newLocal.Identifier , newLocal.SpecType)
                rlog(indent,"EvalFrameExprAssign3: FloatValue=%v FloatValue=%v idtype=%v \n", newLocal.FloatValue , newLocal.FloatValue, newLocal.SpecType.Id)
                f.SetLocal(newLocal, indent+1)
            }
            localVal2 := f.GetLocal(val3.Value,indent+1)
            rlog6(indent,"EvalFrameExprAssign : localVal=%v localVal2=%v \n", localVal, localVal2)
        }
    }
    return e
}

func (f *Frame) EvalFrameExprCmd(node *ExprCmd, indent int) *NodeEval  {
    e := NewNodeEval()
    if node != nil {
        rlog6(indent,"EvalFrameExprCmd: %v\n", node)
        e.Value = "1"
        val1  := f.EvalFrameCmdListApp(node.CmdListApp, indent+1)
        val2  := f.EvalFrameCmdGenData(node.CmdGenData, indent+1)
        val3  := f.EvalFrameCmdMakeAppTid(node.CmdMakeAppTid, indent+1)
        val4  := f.EvalFrameCmdMakeAppUser(node.CmdMakeAppUser, indent+1)
        val5  := f.EvalFrameCmdTestRoute(node.CmdTestRoute, indent+1)
        val6  := f.EvalFrameCmdCreateApp(node.CmdCreateApp, indent+1)
        val7  := f.EvalFrameCmdDeleteApp(node.CmdDeleteApp, indent+1)
        val8  := f.EvalFrameCmdDeleteAppTid(node.CmdDeleteAppTid, indent+1)
        val9  := f.EvalFrameCmdRunSQL(node.CmdRunSQL, indent+1)
        val10 := f.EvalFrameCmdDeleteAppUser(node.CmdDeleteAppUser, indent+1)

        rlog6(indent,"EvalFrameExprCmd CmdListApp 22: %v\n", val1)
        rlog7(indent,"EvalFrameExprCmd CmdTestRoute 27: %v\n", val5)
        
        if val1 != nil {
            if len(val1.Value) > 0 {
                e = val1
                rlog6(indent,"EvalFrameExprCmd CmdListApp Result 23: %v\n", e)
            }
        }
        if val2 != nil {
            if len(val2.Value) > 0 {
                e = val2
            }
        }
        if val3 != nil {
            if len(val3.Value) > 0 {
                e = val3
            }
        }
        if val4 != nil {
            if len(val4.Value) > 0 {
                e = val4
            }
        }
        if val5 != nil {
            if len(val5.Value) > 0 {
                e = val5
            }
        }
        if val6 != nil {
            if len(val6.Value) > 0 {
                e = val6
            }
        }
        if val7 != nil {
            if len(val7.Value) > 0 {
                e = val7
            }
        }
        if val8 != nil {
            if len(val8.Value) > 0 {
                e = val8
            }
        }
        if val9 != nil {
            if len(val9.Value) > 0 {
                e = val9
            }
        }
        if val10 != nil {
            if len(val10.Value) > 0 {
                e = val10
            }
        }
    }
    return e
}


func (f *Frame) EvalFrameCmdDeleteAppUser(node *CmdDeleteAppUser, indent int) *NodeEval  {
    e := NewNodeEval()
    if node != nil {
        rlog7(indent,"EvalFrameCmdDeleteAppUser: %v\n", node)

        val1 := f.EvalFrameExpr(node.Expr1, indent+1)
        val2 := f.EvalFrameExpr(node.Expr2, indent+1)
        rlog7(indent,"EvalFrameCmdDeleteAppUser: val1=%v val2=%v\n", val1, val2)

        emailVal := val1.Value
        if val1.IdType == IDENTIFIER {
            localVal := f.GetLocal(val1.Value,indent+1)
            emailVal = localVal.Value
        }
        appVal := val2.Value
        if val2.IdType == IDENTIFIER {
            localVal := f.GetLocal(val2.Value,indent+1)
            rlog8(indent,"EvalFrameCmdDeleteAppUser: localVal=%v\n", localVal)
            appVal = localVal.Value
        }

        rlog7(indent,"EvalFrameCmdDeleteAppUser: sqlVal=%v appVal=%v \n", emailVal, appVal)

        rval := DeleteAppUser(emailVal, appVal)
        rlog7(indent,"EvalFrameCmdDeleteAppUser: rval=%v\n", rval)
        //e.BoolValue = rval.Result
        e.Value = rval.Message
        e.StringValue = e.Value
        e.IdType=TYPE_STRING_ID
        rlog7(indent,"EvalFrameCmdDeleteAppUser: rval Message=%v\n", rval.Message)
        if strings.Contains(rval.Message, "404") {
            localVal := f.GetLocal("stepPass",indent+1)
            localVal.BoolValue = false
            localVal.Value = fmt.Sprintf("%v", localVal.BoolValue)
            localVal.StringValue = localVal.Value
            f.SetLocal(localVal,indent+1)
        }
        return e
    }
    return e
}

func (f *Frame) EvalFrameCmdRunSQL(node *CmdRunSQL, indent int) *NodeEval  {
    e := NewNodeEval()
    if node != nil {
        rlog7(indent,"EvalFrameCmdRunSQL: %v\n", node)

        val1 := f.EvalFrameExpr(node.Expr1, indent+1)
        val2 := f.EvalFrameExpr(node.Expr2, indent+1)
        rlog7(indent,"EvalFrameCmdRunSQL: val2=%v\n", val2)

        sqlVal := val1.Value
        if val1.IdType == IDENTIFIER {
            localVal := f.GetLocal(val1.Value,indent+1)
            sqlVal = localVal.Value
        }
        appVal := val2.Value
        if val2.IdType == IDENTIFIER {
            localVal := f.GetLocal(val2.Value,indent+1)
            rlog7(indent,"EvalFrameCmdRunSQL: localVal=%v\n", localVal)
            appVal = localVal.Value
        }

        rlog7(indent,"EvalFrameCmdRunSQL: sqlVal=%v appVal=%v \n", sqlVal, appVal)

        rval := RunSQL(sqlVal, appVal)
        rlog7(indent,"EvalFrameCmdRunSQL: rval=%v\n", rval)
        //e.BoolValue = rval.Result
        e.Value = rval.Message
        e.StringValue = e.Value
        e.IdType=TYPE_STRING_ID
        //rlog7(indent,"EvalFrameCmdRunSQL: rval=%v  e.Value=%v\n", rval, e.Value)
        //fmt.Printf("EvalFrameCmdRunSQL: rval=%v  e.Value=%v\n", rval, e.Value)
        return e
    }
    return e
}


func (f *Frame) EvalFrameCmdDeleteApp(node *CmdDeleteApp, indent int) *NodeEval  {
    e := NewNodeEval()
    if node != nil {
        rlog7(indent,"EvalFrameCmdDeleteApp: %v\n", node)

        val1 := f.EvalFrameExpr(node.Expr1, indent+1)

        appName := val1.Value
        if val1.IdType == IDENTIFIER {
            localVal := f.GetLocal(val1.Value,indent+1)
            appName = localVal.Value
            rlog7(indent,"EvalFrameCmdDeleteApp appName: %v\n", appName)
        }

        rlog7(indent,"EvalFrameCmdDeleteApp: Expr2=%v\n", appName)

        rval := DeleteApp(appName)
        rlog(indent,"EvalFrameCmdDeleteApp: DeleteApp rval=%v\n", rval)

        mode := 1
        if mode == 0 {
            e.BoolValue = rval.Result
            e.Value = fmt.Sprintf("%v",rval)
            e.StringValue = e.Value
            e.IdType=TYPE_BOOL_ID
            //rlog7(indent,"EvalFrameCmdDeleteApp: Result=%v  e.Value=%v\n", result, e.Value)
        }
        if mode == 1 {
            e.Value = rval.Message
            e.StringValue = e.Value
            e.IdType=TYPE_STRING_ID
            rlog(indent,"EvalFrameCmdDeleteApp: Message= %v\n", e.Value)
            if strings.Contains(rval.Message, "200") {
                localVal := f.GetLocal("stepPass",indent+1)
                localVal.BoolValue = true
                localVal.Value = fmt.Sprintf("%v", localVal.BoolValue)
                localVal.StringValue = localVal.Value
                f.SetLocal(localVal,indent+1)
            } else {
                localVal := f.GetLocal("stepPass",indent+1)
                localVal.BoolValue = false
                localVal.Value = fmt.Sprintf("%v", localVal.BoolValue)
                localVal.StringValue = localVal.Value
                f.SetLocal(localVal,indent+1)
            }
            localVal := f.GetLocal("stepPass",indent+1)
            rlog(indent,"EvalFrameCmdDeleteApp: stepPass Value= %v\n", localVal.Value)
        }

        return e
    }
    return e
}

func (f *Frame) EvalFrameCmdCreateApp(node *CmdCreateApp, indent int) *NodeEval  {
    //var appName string
    e := NewNodeEval()
    if node != nil {
        rlog7(indent,"EvalFrameCmdCreateApp: %v\n", node)

        val1 := f.EvalFrameExpr(node.Expr1, indent+1)

        appName := val1.Value
        if val1.IdType == IDENTIFIER {
            localVal := f.GetLocal(val1.Value,indent+1)
            appName = localVal.Value
            rlog7(indent,"EvalFrameCmdCreateApp appName: %v\n", appName)
        }

        rlog7(indent,"EvalFrameCmdCreateApp: Expr2=%v\n", appName)

        rval := CreateApp(appName)
        rlog7(indent,"EvalFrameCmdCreateApp: CreateApp rval=%v\n", rval)
        mode := 1
        if mode == 0 {
            e.BoolValue = rval.Result
            e.Value = fmt.Sprintf("%v",rval)
            e.StringValue = e.Value
            e.IdType=TYPE_BOOL_ID
        }
        if mode == 1 {
            e.Value = rval.Message
            e.StringValue = e.Value
            e.IdType=TYPE_STRING_ID
            if strings.Contains(rval.Message, "201") {
                localVal := f.GetLocal("stepPass",indent+1)
                localVal.BoolValue = true
                localVal.Value = fmt.Sprintf("%v", localVal.BoolValue)
                localVal.StringValue = localVal.Value
                f.SetLocal(localVal,indent+1)
            } else {
                localVal := f.GetLocal("stepPass",indent+1)
                localVal.BoolValue = false
                localVal.Value = fmt.Sprintf("%v", localVal.BoolValue)
                localVal.StringValue = localVal.Value
                f.SetLocal(localVal,indent+1)
            }
        }
        //rlog7(indent,"EvalFrameCmdCreateApp: Result=%v  e.Value=%v\n", result, e.Value)
        return e
    }
    return e
}


func (f *Frame) EvalFrameCmdTestRoute(node *CmdTestRoute, indent int) *NodeEval  {
    //var appName string
    e := NewNodeEval()
    if node != nil {
        rlog7(indent,"EvalFrameCmdTestRoute: %v\n", node)

        val1 := f.EvalFrameExpr(node.Expr1, indent+1)
        val2 := f.EvalFrameExpr(node.Expr2, indent+1)

        appName := val2.Value
        if val2.IdType == IDENTIFIER {
            localVal := f.GetLocal(val2.Value,indent+1)
            appName = localVal.Value
            rlog7(indent,"EvalFrameCmdTestRoute appName: %v\n", appName)
        }

        param2 := val1.Value
        if val1.IdType == IDENTIFIER {
            localVal := f.GetLocal(val1.Value,indent+1)
            param2 = localVal.Value
            rlog7(indent,"EvalFrameCmdTestRoute appName: %v\n", param2)
        }

        rlog7(indent,"EvalFrameCmdTestRoute: Expr1=%v\n", param2)
        rlog7(indent,"EvalFrameCmdTestRoute: Expr2=%v\n", appName)

        result, rval := DoTestRoute1(appName, param2)
        if len(rval) == 0 {
            rval = "(empty string value)"
        }
        rlog7(indent,"EvalFrameCmdTestRoute: DoTestRoute1 rval=%v\n", rval)
        //DoTest(val2.Value)
        e.BoolValue = result
        e.Value = rval
        e.StringValue = rval
        e.IdType=TYPE_STRING_ID
        //rlog7(indent,"EvalFrameCmdTestRoute: Result=%v  e.Value=%v\n", result, e.Value)
        return e
    }
    return e
}


func (f *Frame) EvalFrameCmdMakeAppUser(node *CmdMakeAppUser, indent int) *NodeEval  {
    e := NewNodeEval()
    if node != nil {
        rlog7(indent,"EvalFrameCmdMakeAppUser: %v\n", node)

        val1 := f.EvalFrameExpr(node.Expr1, indent+1)
        val2 := f.EvalFrameExpr(node.Expr2, indent+1)

        appName := val2.Value
        if val2.IdType == IDENTIFIER {
            localVal := f.GetLocal(val2.Value,indent+1)
            appName = localVal.Value
            rlog7(indent,"EvalFrameCmdTestRoute appName: %v\n", appName)
        }

        param2 := val1.Value
        if val1.IdType == IDENTIFIER {
            localVal := f.GetLocal(val1.Value,indent+1)
            param2 = localVal.Value
            rlog7(indent,"EvalFrameCmdTestRoute appName: %v\n", param2)
        }

        rlog7(indent,"EvalFrameCmdMakeAppUser: Expr1=%v\n", param2)
        rlog7(indent,"EvalFrameCmdMakeAppUser: Expr2=%v\n", appName)
        var mode int
        mode = 1
        if mode == 0 {
            result := MakeAppuser(param2, appName)
            e.BoolValue = result
            e.Value = fmt.Sprintf("%v",result)
            e.StringValue = fmt.Sprintf("%v",result)
            e.IdType=TYPE_BOOL_ID
        }
        if mode == 1 {
            rval := MakeAppuser2(param2, appName)
            e.Value = fmt.Sprintf("%v",rval.Message)
            e.StringValue = e.Value 
            e.IdType=TYPE_STRING_ID
            if strings.Contains(rval.Message, "201") {
                localVal := f.GetLocal("stepPass",indent+1)
                localVal.BoolValue = true
                localVal.Value = fmt.Sprintf("%v", localVal.BoolValue)
                localVal.StringValue = localVal.Value
                f.SetLocal(localVal,indent+1)
            } else {
                localVal := f.GetLocal("stepPass",indent+1)
                localVal.BoolValue = false
                localVal.Value = fmt.Sprintf("%v", localVal.BoolValue)
                localVal.StringValue = localVal.Value
                f.SetLocal(localVal,indent+1)
            }
            rlog7(indent,"EvalFrameCmdMakeAppUser:  MakeAppuser2 e.Value=%v\n", e.Value)
        }

        rlog7(indent,"EvalFrameCmdMakeAppUser:  e.Value=%v\n", e.Value)
        return e
    }
    return e
}


func (f *Frame) EvalFrameCmdMakeAppTid(node *CmdMakeAppTid, indent int) *NodeEval  {
    e := NewNodeEval()
    if node != nil {
        rlog7(indent,"EvalFrameCmdMakeAppTid: %v\n", node)

    //ExprLoop    *Expr
    //ExprApp     *Expr
        val1 := f.EvalFrameExpr(node.Expr1, indent+1)
        val2 := f.EvalFrameExpr(node.Expr2, indent+1)

        rlog7(indent,"EvalFrameCmdMakeAppTid: uuid=%v\n", val1.Value)
        rlog7(indent,"EvalFrameCmdMakeAppTid: app=%v\n", val2.Value)

        appName := val2.Value
        if val2.IdType == IDENTIFIER {
            localVal := f.GetLocal(val2.Value,indent+1)
            appName = localVal.Value
            rlog7(indent,"EvalFrameCmdMakeAppTid appName: %v\n", appName)
        }

        param2 := val1.Value
        if val1.IdType == IDENTIFIER {
            localVal := f.GetLocal(val1.Value,indent+1)
            param2 = localVal.Value
            rlog7(indent,"EvalFrameCmdMakeAppTid appName: %v\n", param2)
        }

        rval := MakeAppTid(param2, appName)
        //e.BoolValue = resultTID
        //e.Value = fmt.Sprintf("%v",resultTID)
        //e.StringValue = fmt.Sprintf("%v",resultTID)
        //e.IdType=TYPE_BOOL_ID
        e.Value = rval.Message
        e.StringValue  = e.Value
        e.IdType=TYPE_STRING_ID
        rlog(indent,"EvalFrameCmdMakeAppTid Value: %v\n", e.Value)
        if strings.Contains(rval.Message, "200") || strings.Contains(rval.Message, "201") {
            localVal := f.GetLocal("stepPass",indent+1)
            localVal.BoolValue = true
            localVal.Value = fmt.Sprintf("%v", localVal.BoolValue)
            localVal.StringValue = localVal.Value
            f.SetLocal(localVal,indent+1)
        } else {
            localVal := f.GetLocal("stepPass",indent+1)
            localVal.BoolValue = false
            localVal.Value = fmt.Sprintf("%v", localVal.BoolValue)
            localVal.StringValue = localVal.Value
            f.SetLocal(localVal,indent+1)
        }
        localVal := f.GetLocal("stepPass",indent+1)
        rlog(indent,"EvalFrameCmdMakeAppTid: 123  e.Value=%v stepPass Value=%v\n", e.Value, localVal.Value)
        return e
    }
    return e
}

func (f *Frame) EvalFrameCmdDeleteAppTid(node *CmdDeleteAppTid, indent int) *NodeEval  {
    e := NewNodeEval()
    if node != nil {
        rlog7(indent,"EvalFrameCmdDeleteAppTid: %v\n", node)

        app := f.EvalFrameExpr(node.Expr1, indent+1)
        uuid := f.EvalFrameExpr(node.Expr2, indent+1)

        appName := app.Value
        if app.IdType == IDENTIFIER {
            localVal := f.GetLocal(app.Value,indent+1)
            appName = localVal.Value
            rlog7(indent,"EvalFrameCmdDeleteAppTid appName: %v\n", appName)
        }

        uuidVal := uuid.Value
        if uuid.IdType == IDENTIFIER {
            localVal := f.GetLocal(uuid.Value,indent+1)
            uuidVal = localVal.Value
            rlog7(indent,"EvalFrameCmdDeleteAppTid appName: %v\n", uuidVal)
        }

        rlog7(indent,"EvalFrameCmdDeleteAppTid: Expr2=%v\n", appName)

        rval := DeleteAppTid(appName, uuidVal)
        //e.BoolValue = rval.Result
        //e.Value = fmt.Sprintf("%v",rval)
        //e.StringValue = e.Value
        //e.IdType=TYPE_BOOL_ID
        e.Value = rval.Message
        e.StringValue = e.Value
        e.IdType=TYPE_STRING_ID
        rlog(indent,"EvalFrameCmdDeleteAppTid: 234 rval=%v\n", rval)
        if strings.Contains(rval.Message, "200") {
            localVal := f.GetLocal("stepPass",indent+1)
            localVal.BoolValue = true
            localVal.Value = fmt.Sprintf("%v", localVal.BoolValue)
            localVal.StringValue = localVal.Value
            f.SetLocal(localVal,indent+1)
        } else {
            localVal := f.GetLocal("stepPass",indent+1)
            localVal.BoolValue = false
            localVal.Value = fmt.Sprintf("%v", localVal.BoolValue)
            localVal.StringValue = localVal.Value
            f.SetLocal(localVal,indent+1)
        }

        //rlog7(indent,"EvalFrameCmdDeleteAppTid: Result=%v  e.Value=%v\n", result, e.Value)
        return e
    }
    return e
}


func (f *Frame) EvalFrameCmdGenData(node *CmdGenData, indent int) *NodeEval  {
    e := NewNodeEval()
    if node != nil {
        rlog7(indent,"EvalFrameCmdGenData: %v\n", node)

        val2 := f.EvalFrameExpr(node.ExprLoop, indent+1)
        val3 := f.EvalFrameExpr(node.ExprApp, indent+1)

        appName := val3.Value
        if val3.IdType == IDENTIFIER {
            localVal := f.GetLocal(val3.Value,indent+1)
            appName = localVal.Value
            rlog7(indent,"EvalFrameCmdTestRoute appName: %v\n", appName)
        }

        param2 := val2.Value
        if val2.IdType == IDENTIFIER {
            localVal := f.GetLocal(val2.Value,indent+1)
            param2 = localVal.Value
            rlog7(indent,"EvalFrameCmdTestRoute appName: %v\n", param2)
        }

        rlog7(indent,"EvalFrameCmdGenData: val2=%v\n", param2)
        rlog7(indent,"EvalFrameCmdGenData: val3=%v\n", appName)

        loopCount, _ := strconv.Atoi(param2)

        val1 := GenerateTestData(appName, loopCount)
        e.BoolValue = val1
        e.Value = fmt.Sprintf("%v",val1)
        e.StringValue = fmt.Sprintf("%v",val1)
        e.IdType=TYPE_BOOL_ID
        rlog7(indent,"EvalFrameCmdGenData: Result=%v  e.Value=%v\n", val1, e.Value)
        return e
    }
    return e
}

func (f *Frame) EvalFrameCmdListApp(node *CmdListApp, indent int) *NodeEval  {
    e := NewNodeEval()
    if node != nil {
        rlog6(indent,"EvalFrameCmdListApp: %v\n", node)
        val1 := DoListApps()
        e.Value = val1
        e.StringValue = val1
        e.IdType=TYPE_STRING_ID
        rlog6(indent,"EvalFrameCmdListApp: Result=%v  e.Value=%v\n", val1, e.Value)
        return e
    }
    return e
}

//func (f *Frame) EvalFrameExprUnary(node *ExprUnary, indent int) {
//    if node != nil {
//        
//        rlog(indent,"EvalFrameExprUnary: ", node)
//    }
//}

func (f *Frame) EvalFrameExprConditional(node *ExprConditional, indent int) *NodeEval {
    e := NewNodeEval()
    if node != nil {
        
        rlog3(indent,"EvalFrameExprConditional: %v\n", node)
        val1 := f.EvalFrameExprLogicOr(node.ExprLogicOr, indent+1)
        rlog(indent,"EvalFrameExprConditional: ExprLogicOr=%v \n", val1)
        e = val1
    }
    return e
}

func (f *Frame) EvalFrameExprLogicOr(node *ExprLogicOr, indent int) *NodeEval {
    e := NewNodeEval()
    if node != nil {
        
        rlog3(indent,"EvalFrameExprLogicOr: %v \n", node)
        val1 := f.EvalFrameExprLogicAnd(node.ExprLogicAnd, indent+1)
        val2 := f.EvalFrameExprLogicOr(node.ExprLogicOr, indent+1)
        
        rlog(indent,"EvalFrameExprLogicOr: ExprLogicAnd=%v ExprLogicOr=%v \n", val1, val2)
        
        rlog(indent,"asdasdasdadasd1\n")
        if val1.IdType > 0 {
            
            rlog(indent,"asdasdasdadasd2\n")
            e = val1
        }
        if val2.IdType > 0 {
            
            rlog(indent,"asdasdasdadasd3\n")
            e = val2
        }
        if val1.IdType > 0 && val2.IdType > 0 {
            
            rlog(indent,"asdasdasdadasd4 %v %v\n", val1.IdType, val2.IdType)
            evalNode1 := NewNodeEval()
            evalNode2 := NewNodeEval()
            
            if val1.IdType == IDENTIFIER {
                localVar1 := f.GetLocal(val1.Value, indent+1)
                evalNode1.Value = localVar1.Value
                evalNode1.IdType = localVar1.SpecType.Id
                evalNode1.BoolValue = localVar1.BoolValue
                evalNode1.FloatValue = localVar1.FloatValue
                evalNode1.BoolValue = localVar1.BoolValue
                evalNode1.StringValue = localVar1.StringValue
                //
                //rlog(indent,"111111", val1.Value, localVar1.SpecType)
            }
            if val2.IdType == IDENTIFIER {
                localVar2 := f.GetLocal(val2.Value, indent+1)
                evalNode2.Value = localVar2.Value
                evalNode2.IdType = localVar2.SpecType.Id
                evalNode2.BoolValue = localVar2.BoolValue
                evalNode2.FloatValue = localVar2.FloatValue
                evalNode2.BoolValue = localVar2.BoolValue
                evalNode2.StringValue = localVar2.StringValue
                //
                //rlog(indent,"2222222", evalNode2.Value, evalNode2.IdType  )
            }
            if val1.IdType == TYPE_BOOL_ID  {
                evalNode1.Value = val1.Value
                evalNode1.IdType = val1.IdType
                evalNode1.BoolValue = val1.BoolValue
                //
                //rlog(indent,"333333333", evalNode1.Value , evalNode1.BoolValue, evalNode1.IdType)
            }
            if val2.IdType == TYPE_BOOL_ID  {
                evalNode2.Value = val2.Value
                evalNode2.IdType = val2.IdType
                evalNode2.BoolValue = val2.BoolValue
                //
                //rlog(indent,"4444444444", evalNode2.Value , evalNode2.BoolValue, evalNode2.IdType)
            }
            
            rlog(indent,"Type: %v %v \n", evalNode1.IdType, evalNode2.IdType )
            if evalNode1.IdType == TYPE_BOOL_ID  && evalNode2.IdType == TYPE_BOOL_ID  {
                
                rlog(indent,"Valid: Compatible Type:\n")
            } else {
                
                rlog(indent,"Invalid: Incompatible Type:\n")
            }
            if evalNode1.IdType == TYPE_BOOL_ID  && evalNode2.IdType == TYPE_BOOL_ID  {
                evalNodeResult := NewNodeEval()
                evalNodeResult.IdType = TYPE_BOOL_ID
                evalNodeResult.BoolValue = evalNode2.BoolValue || evalNode1.BoolValue
                evalNodeResult.Value = fmt.Sprintf("%v", evalNodeResult.BoolValue)
                
                rlog(indent,"%v || %v = %v \n", evalNode2.BoolValue, evalNode1.BoolValue, evalNodeResult.BoolValue)
                return evalNodeResult
            }
        }

    }
    return e
}

func (f *Frame) EvalFrameExprLogicAnd(node *ExprLogicAnd, indent int) *NodeEval {
    e := NewNodeEval()
    if node != nil {
        
        rlog3(indent,"EvalFrameExprLogicAnd: %v \n ", node)
        val1 := f.EvalFrameExprInclOr(node.ExprInclOr, indent+1)
        val2 := f.EvalFrameExprLogicAnd(node.ExprLogicAnd, indent+1)
        
        rlog(indent,"EvalFrameExprLogicAnd: ExprInclOr=%v ExprLogicAnd=%v \n", val1, val2)
        e = val1
    }
    return e
}

func (f *Frame) EvalFrameExprInclOr(node *ExprInclOr, indent int) *NodeEval {
    e := NewNodeEval()
    if node != nil {
        
        rlog3(indent,"EvalFrameExprInclOr: %v \n", node)
        val1 := f.EvalFrameExprExclOr(node.ExprExclOr, indent+1)
        val2 := f.EvalFrameExprInclOr(node.ExprInclOr, indent+1)
        
        rlog(indent,"EvalFrameExprInclOr: ExprExclOr=%v ExprInclOr=%v \n", val1, val2)
        e = val1
    }
    return e
}

func (f *Frame) EvalFrameExprExclOr(node *ExprExclOr, indent int) *NodeEval {
    e := NewNodeEval()
    if node != nil {
        
        rlog3(indent,"EvalFrameExprExclOr: %v \n", node)
        val1 := f.EvalFrameExprAnd(node.ExprAnd, indent+1)
        val2 := f.EvalFrameExprExclOr(node.ExprExclOr, indent+1)
        
        rlog(indent,"EvalFrameExprExclOr: ExprAnd=%v ExprExclOr=%v \n", val1, val2)
        e = val1
    }
    return e
}

func (f *Frame) EvalFrameExprAnd(node *ExprAnd, indent int) *NodeEval {
    e := NewNodeEval()
    if node != nil {
        
        rlog3(indent,"EvalFrameExprAnd: %v \n", node)
        val1 := f.EvalFrameExprEquality(node.ExprEquality, indent+1)
        val2 := f.EvalFrameExprAnd(node.ExprAnd, indent+1)
        
        rlog(indent,"EvalFrameExprAnd: ExprEquality=%v ExprAnd=%v \n", val1, val2)
        e = val1
        if val1.IdType > 0 && val2.IdType > 0 {
            evalNode1 := NewNodeEval()
            evalNode2 := NewNodeEval()
            
            if val1.IdType == IDENTIFIER {
                localVar1 := f.GetLocal(val1.Value, indent+1)
                evalNode1.Value = localVar1.Value
                evalNode1.IdType = localVar1.SpecType.Id
                evalNode1.BoolValue = localVar1.BoolValue
                evalNode1.FloatValue = localVar1.FloatValue
                evalNode1.BoolValue = localVar1.BoolValue
                evalNode1.StringValue = localVar1.StringValue
                //
                //rlog(indent,"111111", val1.Value, localVar1.SpecType)
            }
            if val2.IdType == IDENTIFIER {
                localVar2 := f.GetLocal(val2.Value, indent+1)
                evalNode2.Value = localVar2.Value
                evalNode2.IdType = localVar2.SpecType.Id
                evalNode2.BoolValue = localVar2.BoolValue
                evalNode2.FloatValue = localVar2.FloatValue
                evalNode2.BoolValue = localVar2.BoolValue
                evalNode2.StringValue = localVar2.StringValue
                //
                //rlog(indent,"2222222", evalNode2.Value, evalNode2.IdType  )
            }
            if val1.IdType == TYPE_BOOL_ID  {
                evalNode1.Value = val1.Value
                evalNode1.IdType = val1.IdType
                evalNode1.BoolValue = val1.BoolValue
                //
                //rlog(indent,"333333333", evalNode1.Value , evalNode1.BoolValue, evalNode1.IdType)
            }
            if val2.IdType == TYPE_BOOL_ID  {
                evalNode2.Value = val2.Value
                evalNode2.IdType = val2.IdType
                evalNode2.BoolValue = val2.BoolValue
                //
                //rlog(indent,"4444444444", evalNode2.Value , evalNode2.BoolValue, evalNode2.IdType)
            }
            
            rlog(indent,"Type: %v %v \n", evalNode1.IdType, evalNode2.IdType )
            if evalNode1.IdType == TYPE_BOOL_ID  && evalNode2.IdType == TYPE_BOOL_ID  {
                
                rlog(indent,"Valid: Compatible Type:\n")
            } else {
                
                rlog(indent,"Invalid: Incompatible Type:\n")
            }
            if evalNode1.IdType == TYPE_BOOL_ID  && evalNode2.IdType == TYPE_BOOL_ID  {
                evalNodeResult := NewNodeEval()
                evalNodeResult.IdType = TYPE_BOOL_ID
                evalNodeResult.BoolValue = evalNode2.BoolValue && evalNode1.BoolValue
                evalNodeResult.Value = fmt.Sprintf("%v", evalNodeResult.BoolValue)
                
                rlog(indent,"%v && %v = %v \n", evalNode2.IntValue, evalNode1.IntValue, evalNodeResult.BoolValue)
                return evalNodeResult
            }
        }

    }
    return e
}

func (f *Frame) EvalFrameExprEquality(node *ExprEquality, indent int) *NodeEval {
    e := NewNodeEval()
    if node != nil {
        
        rlog3(indent,"EvalFrameExprEquality: %v \n", node)
        val1 := f.EvalFrameExprRelational(node.ExprRelational, indent+1)
        val2 := f.EvalFrameExprEquality(node.ExprEquality, indent+1)
        
        rlog(indent,"EvalFrameExprEquality: ExprRelational=%v ExprEquality=%v \n", val1, val2)
        e = val1
        if val1.IdType > 0 && val2.IdType > 0 {
            evalNode1 := NewNodeEval()
            evalNode2 := NewNodeEval()
            
            if val1.IdType == IDENTIFIER {
                localVar1 := f.GetLocal(val1.Value, indent+1)
                evalNode1.Value = localVar1.Value
                evalNode1.IdType = localVar1.SpecType.Id
                evalNode1.IntValue = localVar1.IntValue
                evalNode1.FloatValue = localVar1.FloatValue
                evalNode1.BoolValue = localVar1.BoolValue
                evalNode1.StringValue = localVar1.StringValue
                
                rlog(indent,"111111 %v %v\n", val1.Value, localVar1.SpecType)
            }
            if val2.IdType == IDENTIFIER {
                localVar2 := f.GetLocal(val2.Value, indent+1)
                evalNode2.Value = localVar2.Value
                evalNode2.IdType = localVar2.SpecType.Id
                evalNode2.IntValue = localVar2.IntValue
                evalNode2.FloatValue = localVar2.FloatValue
                evalNode2.BoolValue = localVar2.BoolValue
                evalNode2.StringValue = localVar2.StringValue
                
                rlog(indent,"2222222 %v %v \n", evalNode2.Value, evalNode2.IdType  )
            }
            if val1.IdType == INTEGER_ID || val1.IdType == TYPE_INT_ID {
                evalNode1.Value = val1.Value
                evalNode1.IdType = val1.IdType
                evalNode1.IntValue, _ = strconv.Atoi(val1.Value)
                
                rlog(indent,"333333333 %v %v %v \n", evalNode1.Value , evalNode1.IntValue, evalNode1.IdType)
            }
            if val2.IdType == INTEGER_ID || val2.IdType == TYPE_INT_ID {
                evalNode2.Value = val2.Value
                evalNode2.IdType = val2.IdType
                evalNode2.IntValue, _ = strconv.Atoi(val2.Value)
                
                rlog(indent,"4444444444 %v %v %v \n", evalNode2.Value , evalNode2.IntValue, evalNode2.IdType)
            }
            if val1.IdType == FLOAT_ID || val1.IdType == TYPE_FLOAT_ID {
                evalNode1.Value = val1.Value
                evalNode1.IdType = val1.IdType
                evalNode1.FloatValue, _ = strconv.ParseFloat(val1.Value,64)
                
                rlog(indent,"5345345345 %v %v %v\n", evalNode1.Value , evalNode1.IntValue, evalNode1.IdType)
            }
            if val2.IdType == FLOAT_ID || val2.IdType == TYPE_FLOAT_ID {
                evalNode2.Value = val2.Value
                evalNode2.IdType = val2.IdType
                evalNode2.FloatValue, _ = strconv.ParseFloat(val2.Value,64)
                
                rlog(indent,"44447897899444444 %v %v %v\n", evalNode2.Value , evalNode2.IntValue, evalNode2.IdType)
            }
            if val1.IdType == TYPE_BOOL_ID  {
                evalNode1.Value = val1.Value
                evalNode1.IdType = val1.IdType
                evalNode1.BoolValue = val1.BoolValue
                //
                //rlog(indent,"55555555555", evalNode1.Value , evalNode1.BoolValue, evalNode1.IdType)
            }
            if val2.IdType == TYPE_BOOL_ID  {
                evalNode2.Value = val2.Value
                evalNode2.IdType = val2.IdType
                evalNode2.BoolValue = val2.BoolValue
                //
                //rlog(indent,"66666666666", evalNode2.Value , evalNode2.BoolValue, evalNode2.IdType)
            }

            
            rlog(indent,"Type: %v %v \n", evalNode1.IdType, evalNode2.IdType )
            if (evalNode1.IdType == INTEGER_ID || evalNode1.IdType == TYPE_INT_ID) && (evalNode2.IdType == INTEGER_ID || evalNode2.IdType == TYPE_INT_ID) {
                
                rlog(indent,"Valid: Compatible Type: INTEGER_ID\n")
            } else if (evalNode1.IdType == FLOAT_ID || evalNode1.IdType == TYPE_FLOAT_ID) && (evalNode2.IdType == FLOAT_ID || evalNode2.IdType == TYPE_FLOAT_ID) {
                rlog(indent,"Valid: Compatible Type: FLOAT_ID\n")
            } else if (evalNode1.IdType == TYPE_BOOL_ID ) && (evalNode2.IdType == TYPE_BOOL_ID ) {
                rlog(indent,"Valid: Compatible Type: TYPE_BOOL_ID\n")
            } else {
                rlog(indent,"Invalid: Incompatible Type:\n")
                panic("EvalFrameExprEquality Invalid: Incompatible Type")
            }
            if evalNode1.IdType == TYPE_BOOL_ID  && evalNode2.IdType == TYPE_BOOL_ID  {
                evalNodeResult := NewNodeEval()
                evalNodeResult.IdType = TYPE_BOOL_ID
                if node.Sign == "EQ" {
                    rlog(indent,"EvalFrameExprEquality EQ :\n")
                    evalNodeResult.BoolValue = evalNode2.BoolValue == evalNode1.BoolValue
                }
                if node.Sign == "NE" {
                    rlog(indent,"EvalFrameExprEquality NE :\n")
                    evalNodeResult.BoolValue = evalNode2.BoolValue != evalNode1.BoolValue
                }
                evalNodeResult.Value = fmt.Sprintf("%v", evalNodeResult.BoolValue)
                
                rlog(indent,"%v %v %v = %v \n", evalNode2.BoolValue, node.Sign, evalNode1.BoolValue, evalNodeResult.BoolValue)
                return evalNodeResult
            }
            if (evalNode1.IdType == INTEGER_ID || evalNode1.IdType == TYPE_INT_ID) && (evalNode2.IdType == INTEGER_ID || evalNode2.IdType == TYPE_INT_ID) {
                evalNodeResult := NewNodeEval()
                evalNodeResult.IdType = TYPE_BOOL_ID
                if node.Sign == "EQ" {
                    evalNodeResult.BoolValue = evalNode2.IntValue == evalNode1.IntValue
                }
                if node.Sign == "NE" {
                    evalNodeResult.BoolValue = evalNode2.IntValue != evalNode1.IntValue
                }
                evalNodeResult.Value = fmt.Sprintf("%v", evalNodeResult.BoolValue)
                
                rlog(indent,"%v %v %v = %v \n", evalNode2.IntValue, node.Sign, evalNode1.IntValue, evalNodeResult.BoolValue)
                return evalNodeResult
            }
            if (evalNode1.IdType == FLOAT_ID || evalNode1.IdType == TYPE_FLOAT_ID) && (evalNode2.IdType == FLOAT_ID || evalNode2.IdType == TYPE_FLOAT_ID) {
                evalNodeResult := NewNodeEval()
                evalNodeResult.IdType = TYPE_BOOL_ID
                if node.Sign == "EQ" {
                    evalNodeResult.BoolValue = evalNode2.FloatValue == evalNode1.FloatValue
                }
                if node.Sign == "NE" {
                    evalNodeResult.BoolValue = evalNode2.FloatValue != evalNode1.FloatValue
                }
                evalNodeResult.Value = fmt.Sprintf("%v", evalNodeResult.BoolValue)
                
                rlog(indent,"FLOAT_ID: %v %v %v = %v \n", evalNode2.FloatValue, node.Sign, evalNode1.FloatValue, evalNodeResult.BoolValue)
                return evalNodeResult
            }
        }

        return e
    }
    return e
}

func (f *Frame) EvalFrameExprRelational(node *ExprRelational, indent int) *NodeEval {
    e := NewNodeEval()
    if node != nil {
        
        rlog3(indent,"EvalFrameExprRelational: %v\n", node)
        val1 := f.EvalFrameExprShift(node.ExprShift, indent+1)
        val2 := f.EvalFrameExprRelational(node.ExprRelational, indent+1)
        
        rlog(indent,"EvalFrameExprRelational: ExprShift=%v ExprRelational=%v \n", val1, val2)
        e = val1
        if val1.IdType > 0 && val2.IdType > 0 {
            evalNode1 := NewNodeEval()
            evalNode2 := NewNodeEval()
            
            if val1.IdType == IDENTIFIER {
                localVar1 := f.GetLocal(val1.Value, indent+1)
                evalNode1.Value = localVar1.Value
                evalNode1.IdType = localVar1.SpecType.Id
                evalNode1.IntValue = localVar1.IntValue
                evalNode1.FloatValue = localVar1.FloatValue
                evalNode1.BoolValue = localVar1.BoolValue
                evalNode1.StringValue = localVar1.StringValue
                //
                //rlog(indent,"212121", val1.Value, localVar1.SpecType)
            }
            if val2.IdType == IDENTIFIER {
                localVar2 := f.GetLocal(val2.Value, indent+1)
                evalNode2.Value = localVar2.Value
                evalNode2.IdType = localVar2.SpecType.Id
                evalNode2.IntValue = localVar2.IntValue
                evalNode2.FloatValue = localVar2.FloatValue
                evalNode2.BoolValue = localVar2.BoolValue
                evalNode2.StringValue = localVar2.StringValue
                //
                //rlog(indent,"222222", evalNode2.Value, evalNode2.IdType  )
            }
            if val1.IdType == FLOAT_ID || val1.IdType == TYPE_FLOAT_ID {
                evalNode1.Value = val1.Value
                evalNode1.IdType = val1.IdType
                evalNode1.FloatValue, _ = strconv.ParseFloat(val1.Value,64)
                //
                //rlog(indent,"333333333", evalNode1.Value , evalNode1.IntValue, evalNode1.IdType)
            }
            if val2.IdType == FLOAT_ID || val2.IdType == TYPE_FLOAT_ID {
                evalNode2.Value = val2.Value
                evalNode2.IdType = val2.IdType
                evalNode2.FloatValue, _ = strconv.ParseFloat(val2.Value,64)
                //
                //rlog(indent,"4444444444", evalNode2.Value , evalNode2.IntValue, evalNode2.IdType)
            }
            if val1.IdType == INTEGER_ID || val1.IdType == TYPE_INT_ID {
                evalNode1.Value = val1.Value
                evalNode1.IdType = val1.IdType
                evalNode1.IntValue, _ = strconv.Atoi(val1.Value)
                //
                //rlog(indent,"23232323", evalNode1.Value , evalNode1.IntValue, evalNode1.IdType)
            }
            if val2.IdType == INTEGER_ID || val2.IdType == TYPE_INT_ID {
                evalNode2.Value = val2.Value
                evalNode2.IdType = val2.IdType
                evalNode2.IntValue, _ = strconv.Atoi(val2.Value)
                //
                //rlog(indent,"24242424", evalNode2.Value , evalNode2.IntValue, evalNode2.IdType)
            }
            
            rlog(indent,"Type: %v %v \n", evalNode1.IdType, evalNode2.IdType )
            if (evalNode1.IdType == INTEGER_ID || evalNode1.IdType == TYPE_INT_ID) && (evalNode2.IdType == INTEGER_ID || evalNode2.IdType == TYPE_INT_ID) {
                
                rlog(indent,"Valid: Compatible Type: INTEGER_ID\n")
            } else if (evalNode1.IdType == FLOAT_ID || evalNode1.IdType == TYPE_FLOAT_ID) && (evalNode2.IdType == FLOAT_ID || evalNode2.IdType == TYPE_FLOAT_ID) {
                
                rlog(indent,"Valid: Compatible Type: FLOAT_ID\n")
            } else {
                
                rlog(indent,"Invalid: Incompatible Type:\n")
            }
            if (evalNode1.IdType == INTEGER_ID || evalNode1.IdType == TYPE_INT_ID) && (evalNode2.IdType == INTEGER_ID || evalNode2.IdType == TYPE_INT_ID) {
                evalNodeResult := NewNodeEval()
                evalNodeResult.IdType = TYPE_BOOL_ID
                if node.Sign == "<" {
                    evalNodeResult.BoolValue = evalNode2.IntValue < evalNode1.IntValue
                }
                if node.Sign == ">" {
                    evalNodeResult.BoolValue = evalNode2.IntValue > evalNode1.IntValue
                }
                if node.Sign == "LE" {
                    evalNodeResult.BoolValue = evalNode2.IntValue <= evalNode1.IntValue
                }
                if node.Sign == "GE" {
                    evalNodeResult.BoolValue = evalNode2.IntValue >= evalNode1.IntValue
                }
                evalNodeResult.Value = fmt.Sprintf("%v", evalNodeResult.BoolValue)
                
                rlog(indent,"%v %v %v = %v \n", evalNode2.IntValue, node.Sign, evalNode1.IntValue, evalNodeResult.BoolValue)
                return evalNodeResult
            }
            if (evalNode1.IdType == FLOAT_ID || evalNode1.IdType == TYPE_FLOAT_ID) && (evalNode2.IdType == FLOAT_ID || evalNode2.IdType == TYPE_FLOAT_ID) {
                evalNodeResult := NewNodeEval()
                evalNodeResult.IdType = TYPE_BOOL_ID
                if node.Sign == "<" {
                    evalNodeResult.BoolValue = evalNode2.FloatValue < evalNode1.FloatValue
                }
                if node.Sign == ">" {
                    evalNodeResult.BoolValue = evalNode2.FloatValue > evalNode1.FloatValue
                }
                if node.Sign == "LE" {
                    evalNodeResult.BoolValue = evalNode2.FloatValue <= evalNode1.FloatValue
                }
                if node.Sign == "GE" {
                    evalNodeResult.BoolValue = evalNode2.FloatValue >= evalNode1.FloatValue
                }
                evalNodeResult.Value = fmt.Sprintf("%v", evalNodeResult.BoolValue)
                
                rlog(indent,"%v %v %v = %v \n", evalNode2.FloatValue, node.Sign, evalNode1.FloatValue, evalNodeResult.BoolValue)
                return evalNodeResult
            }
        }
        return e
    }
    return e
}

func (f *Frame) EvalFrameExprShift(node *ExprShift, indent int) *NodeEval {
    e := NewNodeEval()
    if node != nil {
        
        rlog3(indent,"EvalFrameExprShift: %v\n", node)
        val1 := f.EvalFrameExprAdditive(node.ExprAdditive, indent+1)
        val2 := f.EvalFrameExprShift(node.ExprShift, indent+1)
        e = val1
        
        rlog(indent,"EvalFrameExprShift: ExprAdditive=%v ExprShift=%v \n", val1, val2)
    }
    return e
}

func (f *Frame) EvalFrameExprAdditive(node *ExprAdditive, indent int) *NodeEval {
    e := NewNodeEval()
    if node != nil {
        
        rlog3(indent,"EvalFrameExprAdditive: %v\n", node)
        val1 := f.EvalFrameExprMultiplicative(node.ExprMultiplicative, indent+1)
        val2 := f.EvalFrameExprAdditive(node.ExprAdditive, indent+1)
        
        rlog7(indent,"EvalFrameExprAdditive: ExprMultiplicative=%v ExprAdditive=%v \n", val1, val2)
        e = val1
        if len(val1.Value) > 0 && len(val2.Value) > 0 {
            evalNode1 := NewNodeEval()
            evalNode2 := NewNodeEval()


            if val1.IdType == IDENTIFIER {
                localVar1 := f.GetLocal(val1.Value, indent+1)
                evalNode1.Value = localVar1.Value
                evalNode1.IdType = localVar1.SpecType.Id
                evalNode1.IntValue = localVar1.IntValue
                evalNode1.FloatValue = localVar1.FloatValue
                evalNode1.BoolValue = localVar1.BoolValue
                evalNode1.StringValue = localVar1.StringValue
                //
                rlog7(indent,"EvalFrameExprAdditive IDENTIFIER1 ID=%v IdType=%v Value1=%v IdType1=%v \n", val1.Value, val1.IdType, evalNode1.Value, evalNode1.IdType  )
            }
            if val2.IdType == IDENTIFIER {
                localVar2 := f.GetLocal(val2.Value, indent+1)
                evalNode2.Value = localVar2.Value
                evalNode2.IdType = localVar2.SpecType.Id
                evalNode2.IntValue = localVar2.IntValue
                evalNode2.FloatValue = localVar2.FloatValue
                evalNode2.BoolValue = localVar2.BoolValue
                evalNode2.StringValue = localVar2.StringValue
                //
                rlog7(indent,"EvalFrameExprAdditive IDENTIFIER2 ID=%v IdType=%v Value2=%v IdType2=%v \n", val2.Value, val2.IdType, evalNode2.Value, evalNode2.IdType  )
            }

            if val1.IdType == TYPE_STRING_ID || val1.IdType == LITERAL_STRING  {
                evalNode1.Value = val1.Value
                evalNode1.StringValue = val1.StringValue
                evalNode1.IdType = TYPE_STRING_ID
                rlog7(indent,"EvalFrameExprAdditive TYPE_STRING_ID1 Value=%v StringValue=%v \n", val1.Value, val1.StringValue)
            } 

            if val2.IdType == TYPE_STRING_ID || val2.IdType == LITERAL_STRING  {
                evalNode2.Value = val2.Value
                evalNode2.StringValue = val2.StringValue
                evalNode2.IdType = TYPE_STRING_ID
                rlog7(indent,"EvalFrameExprAdditive TYPE_STRING_ID2 Value=%v StringValue=%v \n", val2.Value, val2.StringValue)
            } 
            rlog7(indent,"nodetype1=%v nodetype2=%v \n", evalNode1.IdType, evalNode2.IdType)
            if val1.IdType == FLOAT_ID || val1.IdType == TYPE_FLOAT_ID {
                evalNode1.Value = val1.Value
                evalNode1.IdType = val1.IdType
                evalNode1.FloatValue, _ = strconv.ParseFloat(val1.Value,64)
                //
                //rlog(indent,"333333333", evalNode1.Value , evalNode1.IntValue, evalNode1.IdType)
            }
            if val2.IdType == FLOAT_ID || val2.IdType == TYPE_FLOAT_ID {
                evalNode2.Value = val2.Value
                evalNode2.IdType = val2.IdType
                evalNode2.FloatValue, _ = strconv.ParseFloat(val2.Value,64)
                //
                //rlog(indent,"4444444444", evalNode2.Value , evalNode2.IntValue, evalNode2.IdType)
            }
            if val1.IdType == INTEGER_ID || val1.IdType == TYPE_INT_ID {
                evalNode1.Value = val1.Value
                evalNode1.IdType = val1.IdType
                evalNode1.IntValue, _ = strconv.Atoi(val1.Value)
                //
                //rlog(indent,"cccccccccccc", evalNode1.Value , evalNode1.IntValue, evalNode1.IdType)
            }
            if val2.IdType == INTEGER_ID || val2.IdType == TYPE_INT_ID {
                evalNode2.Value = val2.Value
                evalNode2.IdType = val2.IdType
                evalNode2.IntValue, _ = strconv.Atoi(val2.Value)
                //
                //rlog(indent,"ddddddddddd", evalNode2.Value , evalNode2.IntValue, evalNode2.IdType)
            }
            
            rlog(indent,"Type: %v %v \n", evalNode1.IdType, evalNode2.IdType )
            if (evalNode1.IdType == INTEGER_ID || evalNode1.IdType == TYPE_INT_ID) && (evalNode2.IdType == INTEGER_ID || evalNode2.IdType == TYPE_INT_ID) {
                rlog(indent,"Valid: Compatible Type:\n")
            } else if (evalNode1.IdType == FLOAT_ID || evalNode1.IdType == TYPE_FLOAT_ID) && (evalNode2.IdType == FLOAT_ID || evalNode2.IdType == TYPE_FLOAT_ID) {
                rlog(indent,"Valid: Compatible Type: FLOAT_ID\n")
            } else if (evalNode1.IdType == TYPE_STRING_ID || evalNode1.IdType == LITERAL_STRING) && (evalNode2.IdType == TYPE_STRING_ID || evalNode2.IdType == LITERAL_STRING) {
                rlog(indent,"Valid: Compatible Type: FLOAT_ID\n")
            } else {
                
                rlog(indent,"Invalid: Incompatible Type:\n")
            }
            if (evalNode1.IdType == INTEGER_ID || evalNode1.IdType == TYPE_INT_ID) && (evalNode2.IdType == INTEGER_ID || evalNode2.IdType == TYPE_INT_ID) {
                evalNodeResult := NewNodeEval()
                evalNodeResult.IdType = INTEGER_ID
                if node.Sign == "+" {
                    evalNodeResult.IntValue = evalNode2.IntValue + evalNode1.IntValue
                }
                if node.Sign == "-" {
                    evalNodeResult.IntValue = evalNode2.IntValue - evalNode1.IntValue
                }
                evalNodeResult.Value = fmt.Sprintf("%v", evalNodeResult.IntValue)
                
                rlog(indent,"%v %v %v = %v \n", evalNode2.IntValue, node.Sign, evalNode1.IntValue, evalNodeResult.IntValue)
                return evalNodeResult
            }
            if (evalNode1.IdType == FLOAT_ID || evalNode1.IdType == TYPE_FLOAT_ID) && (evalNode2.IdType == FLOAT_ID || evalNode2.IdType == TYPE_FLOAT_ID) {
                evalNodeResult := NewNodeEval()
                evalNodeResult.IdType = FLOAT_ID
                if node.Sign == "+" {
                    evalNodeResult.FloatValue = evalNode2.FloatValue + evalNode1.FloatValue
                }
                if node.Sign == "-" {
                    evalNodeResult.FloatValue = evalNode2.FloatValue - evalNode1.FloatValue
                }
                evalNodeResult.Value = fmt.Sprintf("%v", evalNodeResult.FloatValue)
                
                rlog(indent,"%v %v %v = %v \n", evalNode2.FloatValue, node.Sign, evalNode1.FloatValue, evalNodeResult.FloatValue)
                return evalNodeResult
            }
            if (evalNode1.IdType == TYPE_STRING_ID || evalNode1.IdType == LITERAL_STRING) && (evalNode2.IdType == TYPE_STRING_ID || evalNode2.IdType == LITERAL_STRING) {
                evalNodeResult := NewNodeEval()
                evalNodeResult.IdType = TYPE_STRING_ID
                if node.Sign == "+" {
                    evalNodeResult.StringValue = fmt.Sprintf("%v%v", evalNode2.StringValue, evalNode1.StringValue)
                    evalNodeResult.Value = evalNodeResult.StringValue
                }
                //if node.Sign == "-" {
                //}
                
                rlog7(indent,"string+ >> %v %v %v = %v \n", evalNode2.StringValue, node.Sign, evalNode1.StringValue, evalNodeResult.StringValue)
                rlog7(indent,"string+ >>>> %v %v %v = %v \n", evalNode2.Value, node.Sign, evalNode1.Value, evalNodeResult.Value)
                return evalNodeResult
            }
        }
        return e
    }
    return e
}

func (f *Frame) EvalFrameExprMultiplicative(node *ExprMultiplicative, indent int) *NodeEval {
    e := NewNodeEval()
    if node != nil {
        
        rlog3(indent,"EvalFrameExprMultiplicative: %v\n", node)
        val1 := f.EvalFrameExprCast(node.ExprCast, indent+1)
        val2 := f.EvalFrameExprMultiplicative(node.ExprMultiplicative, indent+1)
        
        rlog(indent,"EvalFrameExprMultiplicative: ExprCast=%v ExprMultiplicative=%v \n", val1, val2)
        e = val1
        if len(val1.Value) > 0 && len(val2.Value) > 0 {
            evalNode1 := NewNodeEval()
            evalNode2 := NewNodeEval()
            
            if val1.IdType == IDENTIFIER {
                localVar1 := f.GetLocal(val1.Value, indent+1)
                evalNode1.Value = localVar1.Value
                evalNode1.IdType = localVar1.SpecType.Id
                evalNode1.IntValue = localVar1.IntValue
                evalNode1.FloatValue = localVar1.FloatValue
                evalNode1.BoolValue = localVar1.BoolValue
                evalNode1.StringValue = localVar1.StringValue
                //
                //rlog(indent,"aaaaaaaaaaa", val1.Value, localVar1.SpecType)
            }
            if val2.IdType == IDENTIFIER {
                localVar2 := f.GetLocal(val2.Value, indent+1)
                evalNode2.Value = localVar2.Value
                evalNode2.IdType = localVar2.SpecType.Id
                evalNode2.IntValue = localVar2.IntValue
                evalNode2.FloatValue = localVar2.FloatValue
                evalNode2.BoolValue = localVar2.BoolValue
                evalNode2.StringValue = localVar2.StringValue
                //
                //rlog(indent,"bbbbbbbbbbbb", evalNode2.Value, evalNode2.IdType  )
            }
            if val1.IdType == FLOAT_ID || val1.IdType == TYPE_FLOAT_ID {
                evalNode1.Value = val1.Value
                evalNode1.IdType = val1.IdType
                evalNode1.FloatValue, _ = strconv.ParseFloat(val1.Value,64)
                //
                //rlog(indent,"333333333", evalNode1.Value , evalNode1.IntValue, evalNode1.IdType)
            }
            if val2.IdType == FLOAT_ID || val2.IdType == TYPE_FLOAT_ID {
                evalNode2.Value = val2.Value
                evalNode2.IdType = val2.IdType
                evalNode2.FloatValue, _ = strconv.ParseFloat(val2.Value,64)
                //
                //rlog(indent,"4444444444", evalNode2.Value , evalNode2.IntValue, evalNode2.IdType)
            }
            if val1.IdType == INTEGER_ID || val1.IdType == TYPE_INT_ID {
                evalNode1.Value = val1.Value
                evalNode1.IdType = val1.IdType
                evalNode1.IntValue, _ = strconv.Atoi(val1.Value)
                //
                //rlog(indent,"cccccccccccc", evalNode1.Value , evalNode1.IntValue, evalNode1.IdType)
            }
            if val2.IdType == INTEGER_ID || val2.IdType == TYPE_INT_ID {
                evalNode2.Value = val2.Value
                evalNode2.IdType = val2.IdType
                evalNode2.IntValue, _ = strconv.Atoi(val2.Value)
                //
                //rlog(indent,"ddddddddddd", evalNode2.Value , evalNode2.IntValue, evalNode2.IdType)
            }
            
            rlog(indent,"Type: %v %v \n", evalNode1.IdType, evalNode2.IdType )
            if (evalNode1.IdType == INTEGER_ID || evalNode1.IdType == TYPE_INT_ID) && (evalNode2.IdType == INTEGER_ID || evalNode2.IdType == TYPE_INT_ID) {
                
                rlog(indent,"Valid: Compatible Type:\n")
            } else if (evalNode1.IdType == FLOAT_ID || evalNode1.IdType == TYPE_FLOAT_ID) && (evalNode2.IdType == FLOAT_ID || evalNode2.IdType == TYPE_FLOAT_ID) {
                
                rlog(indent,"Valid: Compatible Type: FLOAT_ID\n")
            } else {
                
                rlog(indent,"Invalid: Incompatible Type:\n")
            }
            if (evalNode1.IdType == INTEGER_ID || evalNode1.IdType == TYPE_INT_ID) && (evalNode2.IdType == INTEGER_ID || evalNode2.IdType == TYPE_INT_ID) {
                evalNodeResult := NewNodeEval()
                evalNodeResult.IdType = INTEGER_ID
                if node.Sign == "*" {
                    evalNodeResult.IntValue = evalNode2.IntValue * evalNode1.IntValue
                }
                if node.Sign == "/" {
                    evalNodeResult.IntValue = evalNode2.IntValue / evalNode1.IntValue
                }
                if node.Sign == "%" {
                    evalNodeResult.IntValue = evalNode2.IntValue % evalNode1.IntValue
                }
                evalNodeResult.Value = fmt.Sprintf("%v", evalNodeResult.IntValue)
                
                rlog(indent,"%v %v %v = %v \n", evalNode2.IntValue, node.Sign, evalNode1.IntValue, evalNodeResult.IntValue)
                return evalNodeResult
            }
            if (evalNode1.IdType == FLOAT_ID || evalNode1.IdType == TYPE_FLOAT_ID) && (evalNode2.IdType == FLOAT_ID || evalNode2.IdType == TYPE_FLOAT_ID) {
                evalNodeResult := NewNodeEval()
                evalNodeResult.IdType = FLOAT_ID
                if node.Sign == "*" {
                    evalNodeResult.FloatValue = evalNode2.FloatValue * evalNode1.FloatValue
                    evalNodeResult.Value = fmt.Sprintf("%v", evalNodeResult.FloatValue)
                }
                if node.Sign == "/" {
                    evalNodeResult.FloatValue = evalNode2.FloatValue / evalNode1.FloatValue
                    evalNodeResult.Value = fmt.Sprintf("%v", evalNodeResult.FloatValue)
                }
                //if node.Sign == "%" {
                //    evalNodeResult.FloatValue = evalNode2.FloatValue % evalNode1.FloatValue
                //}
                //evalNodeResult.Value = fmt.Sprintf("%v", evalNodeResult.IntValue)
                
                rlog(indent,"EvalFrameExprMultiplicative FLOAT_ID %v %v %v = %v : %v \n", evalNode2.FloatValue, node.Sign, evalNode1.FloatValue, evalNodeResult.FloatValue, evalNodeResult.Value)
                return evalNodeResult
            }
        }
        return e
    }
    return e
}

func (f *Frame) EvalFrameExprCast(node *ExprCast, indent int) *NodeEval {
    e := NewNodeEval()
    if node != nil {
        
        rlog3(indent,"EvalFrameExprCast: %v\n", node)
        val1 := f.EvalFrameExprUnary(node.ExprUnary, indent+1)
        
        rlog(indent,"EvalFrameExprCast: ExprUnary=%v\n", val1)
        e = val1
        return e
    }
    return e
}

func (f *Frame) EvalFrameExprUnary(node *ExprUnary, indent int) *NodeEval {
    e := NewNodeEval()
    if node != nil {
        
        rlog3(indent,"EvalFrameExprUnary: %v\n", node)
        val1 := f.EvalFrameExprPostfix(node.ExprPostfix, indent+1)
        val2 := f.EvalFrameUnaryOperator(node.UnaryOperator, indent+1)
        val3 := f.EvalFrameExprCast(node.ExprCast, indent+1)
        
        rlog(indent,"EvalFrameExprUnary: ExprPostfix=%v UnaryOperator=%v ExprCast=%v \n", val1, val2, val3)
        e = val1
        if val1.IdType > 0 {
            if val1.IdType == FUNC_CALL_ID && len(val1.ObjectRef) < 1 {
                // Calling a func means:
                // 1. If there is no package spec, use translation unit of caller.
                // 2. If there is a package spec, use package name to get translation unit.
                // 3. If func instance is not found, default to main translation unit.
                mParams := f.PackageCurrentCollection(indent)
                f.ClearParamCollection() 
                if val1.Value == "instr" {
                    rv := f.HandleInstr(val1, mParams, indent+1)
                    e.IntValue = rv
                    e.Value = fmt.Sprintf("%v", rv)
                    e.StringValue = e.Value
                    e.IdType = TYPE_INT_ID
                    return e
                }
                if val1.Value == "ucase" || val1.Value == "toupper" {
                    rv := f.HandleUcase(val1, mParams, indent+1)
                    e.StringValue = rv
                    e.Value = fmt.Sprintf("%v", rv)
                    e.IdType = TYPE_STRING_ID
                    return e
                }
                if val1.Value == "lcase" || val1.Value == "tolower" {
                    rv := f.HandleLcase(val1, mParams, indent+1)
                    e.StringValue = rv
                    e.Value = fmt.Sprintf("%v", rv)
                    e.IdType = TYPE_STRING_ID
                    return e
                }
                if val1.Value == "repeat" {
                    rv := f.HandleRepeat(val1, mParams, indent+1)
                    e.StringValue = rv
                    e.Value = fmt.Sprintf("%v", rv)
                    e.IdType = TYPE_STRING_ID
                    return e
                }
                if val1.Value == "replace" {
                    rv := f.HandleReplace(val1, mParams, indent+1)
                    e.StringValue = rv
                    e.Value = fmt.Sprintf("%v", rv)
                    e.IdType = TYPE_STRING_ID
                    return e
                }
                mCurrentUnit := f.RtTranslationUnit
                mFuncDef := mCurrentUnit.GetFuncDef(val1.Value, indent+1)
                rlog5(indent,"Result: looking %v in current Translation Unit : %v\n", val1.Value,mFuncDef)
                if mFuncDef != nil {
                    rlog5(indent,"Found in current TU %v \n", val1.Value)
                    retVal := mCurrentUnit.CallFunc(val1.Value, mParams)
                    return retVal
                } else {
                    rlog5(indent,"Not found in current TU, looking in main TU, if not there the use pointer to main \n")
                    packageUnit := mRtTranslationUnitList.GetTranslationUnit("main", val1.Value, indent+1)
                    if packageUnit != nil {
                        retVal := packageUnit.CallFunc(val1.Value, mParams)
                        return retVal
                    } 
                    //if packageUnit == nil && false {
                    //    rlog5(indent,"EvalFrameExprUnary: FUNC_CALL_ID ExprPostfix=%v ParamCollection=%v mParams=%v \n", val1, f.ParamCollection, mParams)
                    //   retVal := mMainTranslation.CallFunc(val1.Value, mParams)
                    //    rlog5(indent,"retVal from %v : %v\n", val1.Value,retVal)
                    //    return retVal
                    //}
                }
            }
            if val1.IdType == FUNC_CALL_ID && len(val1.ObjectRef) > 0 {
                mParams := f.PackageCurrentCollection(indent)
                f.ClearParamCollection() 
                rlog5(indent,"EvalFrameExprUnary: OBJECT_FIELD_REF FUNC_CALL_ID ExprPostfix=%v ParamCollection=%v mParams=%v \n", val1, f.ParamCollection, mParams)
                //retVal := mRtTranslationUnit.CallFunc(val1.Value, mParams)
                rlog5(indent,"val1.ObjectRef=%v val1.Value=%v\n",val1.ObjectRef,val1.Value)
                packageUnit := mRtTranslationUnitList.GetTranslationUnit(val1.ObjectRef, val1.Value, indent+1)
                if packageUnit != nil {
                    rlog5(indent,"packageUnit: %v\n", packageUnit)
                    retVal := packageUnit.CallFunc(val1.Value, mParams)
                    return retVal
                } else {
                    rlog5(indent,"*************************** package name, func name not found.****************************************\n")
                }
            }
        }
        return e
    }
    return e
}



func (f *Frame) EvalFrameExprPostfix2(node *ExprPostfix, indent int) *NodeEval {
    e := NewNodeEval()
    if node != nil {
        
        rlog3(indent,"EvalFrameExprPostfix: %v\n", node)
        /*
        ExprPostfix         *ExprPostfix
        ExprPrimary         *ExprPrimary
        ExprListArgument    *ExprListArgument
        Expr                *Expr
        Identifier          *Identifier
        Operator            string
        Modifier            string
        */
        if node.Modifier == "()" && node.ExprListArgument == nil {
            fc := NewFunctionCall()
            fc.Add("()")
            f.EvalFrameFcExprPostfix(node.ExprPostfix, indent+1, fc)
            
            fc.Reverse()
            //rlog(indent,"Function Call", fc.List)
            if len(fc.List) == 2 {
                fname := fc.List[0]
                //rlog(indent,"fname:", fname)
                retVal := mRtTranslationUnit.CallFunc(fname, nil)
                
                rlog(indent,"EvalFrameExprPostfix: Return Value=%v \n", retVal)
                return retVal
            }
        }
        val1 := f.EvalFrameNodeValue(node.Modifier, indent)
        val2 := f.EvalFrameNodeValue(node.Operator, indent)
        f.EvalFrameIdentifier(node.Identifier, indent)
        val4 := f.EvalFrameExprPrimary(node.ExprPrimary, indent+1)
        f.EvalFrameExprListArgument(node.ExprListArgument, indent+1)
        f.EvalFrameExpr(node.Expr, indent+1)
        f.EvalFrameExprPostfix(node.ExprPostfix, indent+1)
        _ = val1
        _ = val2
        e = val4
        
        rlog(indent,"EvalFrameExprPostfix : ExprPrimary=%v Modifier=%v Operator=%v \n ", val4, val1, val2)
        return e
    }
    return e
}

func (f *Frame) EvalFrameExprPostfix(node *ExprPostfix, indent int) *NodeEval {
    e := NewNodeEval()
    if node != nil {
        
        rlog3(indent,"EvalFrameExprPostfix: %v\n", node)
        /*
        ExprPostfix         *ExprPostfix
        ExprPrimary         *ExprPrimary
        ExprListArgument    *ExprListArgument
        Expr                *Expr
        Identifier          *Identifier
        Operator            string
        Modifier            string
        */
        if node.Modifier == "()" && node.ExprListArgument == nil && false {
            fc := NewFunctionCall()
            fc.Add("()")
            f.EvalFrameFcExprPostfix(node.ExprPostfix, indent+1, fc)
            
            fc.Reverse()
            //rlog(indent,"Function Call", fc.List)
            if len(fc.List) == 2 {
                fname := fc.List[0]
                //rlog(indent,"fname:", fname)
                retVal := mRtTranslationUnit.CallFunc(fname, nil)
                
                rlog(indent,"EvalFrameExprPostfix: Return Value=%v \n", retVal)
                return retVal
            }
        }
        val1 := f.EvalFrameNodeValue(node.Modifier, indent)
        val2 := f.EvalFrameNodeValue(node.Operator, indent)
        val3 := f.EvalFrameIdentifier(node.Identifier, indent)
        val4 := f.EvalFrameExprPrimary(node.ExprPrimary, indent+1)
        val5 := f.EvalFrameExprListArgument(node.ExprListArgument, indent+1)
        val6 := f.EvalFrameExpr(node.Expr, indent+1)
        val7 := f.EvalFrameExprPostfix(node.ExprPostfix, indent+1)
        _ = val1
        _ = val2
        e = val4
        
        rlog(indent,"EvalFrameExprPostfix1 : node=%v, Modifier=%v \n ", node, val1)
        
        rlog(indent,"EvalFrameExprPostfix2 : Operator=%v Identifier=%v \n ", val2, val3)
        
        rlog(indent,"EvalFrameExprPostfix3 : ExprPostfix=%v Expr=%v \n ", val7, val6)
        
        rlog(indent,"EvalFrameExprPostfix4: ExprListArgument=%v ExprPrimary=%v \n ", val5, val4)

        if val1.Value == "()" && len(val7.Value) > 0 && len(val3.Value) == 0 && len(val2.Value) == 0 {
            if val7.IdType ==  IDENTIFIER {
                rlog5(indent,"EvalFrameExprPostfix is IDENTIFIER:\n")
                e.Value = val7.Value
                e.IdType = FUNC_CALL_ID
                return e
            }
            if val7.IdType ==  OBJECT_FIELD_REF {
                rlog5(indent,"EvalFrameExprPostfix is OBJECT_FIELD_REF:\n")
                // e.ObjectRef is package name
                // default to main for now. No package lookup coded yet.
                e.Value = val7.Value
                e.ObjectRef = val7.ObjectRef
                e.IdType = FUNC_CALL_ID
                return e
            }
        }
        if len(val1.Value) == 0 && len(val7.Value) > 0 && len(val3.Value) > 0 && val2.Value == "." {
            rlog5(indent,"EvalFrameExprPostfix OBJECT DOT FIELD access:\n")
            e.Value = val3.Value
            e.ObjectRef = val7.Value
            e.IdType = OBJECT_FIELD_REF
            return e
        }

        return e
    }
    return e
}

func (f *Frame) EvalFrameIdentifier(node *Identifier, indent int) *NodeEval {
    e := NewNodeEval()
    if node != nil {
        e.Value = node.Name
        e.IdType = IDENTIFIER
    }
    return e
}

func (f *Frame) EvalFrameExprPrimary(node *ExprPrimary, indent int) *NodeEval {
    e := NewNodeEval()
    if node != nil {
        
        rlog3(indent,"EvalFrameExprPrimary: node=>%v\n", node)
        val1 := f.EvalFrameExpr(node.Expr, indent+1)
        val2 := f.EvalFrameExprAssign(node.ExprAssign, indent+1)
        val3 := f.EvalFrameNodeValue(node.Value, indent+1)
        if len(val1.Value) > 0 {        
            e = val1
            if node.IdType == TYPE_BOOL_ID {
                e.BoolValue = val1.BoolValue
            }
            if node.IdType == TRUE_ID  {
                e.Value = "true"
                e.BoolValue = true
                e.IdType = TYPE_BOOL_ID
            }
            if node.IdType == FALSE_ID  {
                e.Value = "false"
                e.BoolValue = false
                e.IdType = TYPE_BOOL_ID
            }
            return e
        }
        if len(val2.Value) > 0 {        
            e = val2
            if node.IdType == TYPE_BOOL_ID {
                e.BoolValue = val2.BoolValue
            }
            if node.IdType == TRUE_ID  {
                e.Value = "true"
                e.BoolValue = true
                e.IdType = TYPE_BOOL_ID
            }
            if node.IdType == FALSE_ID  {
                e.Value = "false"
                e.BoolValue = false
                e.IdType = TYPE_BOOL_ID
            }
            return e
        }
        if len(val3.Value) > 0 {          
            e = val3
            e.IdType = node.IdType
            if node.IdType == TYPE_BOOL_ID  {
                e.BoolValue = val3.BoolValue
            }
            if node.IdType == TRUE_ID  {
                e.Value = "true"
                e.BoolValue = true
                e.IdType = TYPE_BOOL_ID
            }
            if node.IdType == FALSE_ID  {
                e.Value = "false"
                e.BoolValue = false
                e.IdType = TYPE_BOOL_ID
            }
            if node.IdType == FLOAT_ID  {
                e.FloatValue, _ = strconv.ParseFloat(e.Value,64)
            }
            if node.IdType == LITERAL_STRING || node.IdType == TYPE_STRING_ID  {
                e.StringValue = e.Value
            }
            return e
        }
    }
    return e
}

func (f *Frame) EvalFrameNodeValue(value string, indent int) *NodeEval {
    e := NewNodeEval()
    if len(value) > 0 {
        
        rlog3(indent,"EvalFrameNodeValue2: %v\n", value)
        //f.EvalFrameExprUnary(node.ExprUnary, indent+1)
        e.Value = value
        return e
    }
    return e
}


func (f *Frame) EvalFrameExprListArgument(node *ExprListArgument, indent int) *NodeEval {
    e := NewNodeEval()
    if node != nil {
        
        rlog3(indent,"EvalFrameExprListArgument: %v\n", node)
        val1 := f.EvalFrameExprListArgument(node.ExprListArgument, indent+1)
        val2 := f.EvalFrameExprAssign(node.ExprAssign, indent+1)
        
        rlog(indent,"EvalFrameExprListArgument Value: ExprListArgument=%v ExprAssign=%v \n", val1, val2)
        if val1 != nil {
            e = val1
        }
        if val2 != nil {
            f.AddParam(val2)
            e = val2
        }
    }
    return e
}


func (f *Frame) EvalFrameUnaryOperator(node *UnaryOperator, indent int) string {
    if node != nil {
        
        rlog3(indent,"EvalFrameUnaryOperator: %v\n", node)
        //f.EvalFrameExprUnary(node.ExprUnary, indent+1)
        return ""
    }
    return ""
}

//-----------------------------------
//
func (f *Frame) EvalFrameExpr(node *Expr, indent int) *NodeEval {
    e := NewNodeEval()
    if node != nil {
        
        rlog3(indent,"EvalFrameExpr: %v \n", node)
        val1 := f.EvalFrameExprAssign(node.ExprAssign, indent+1)
        val2 := f.EvalFrameExpr(node.Expr, indent+1)
        
        //rlog(indent,"EvalFrameExpr: ExprAssign=%v Expr=%v \n", val1, val2)
        if val1.IdType > 0 {
            e = val1
        }
        if val2.IdType > 0 {
            e = val2
        }
    }
    return e
}

func (f *Frame) EvalFrameStmtJump(node *StmtJump, indent int) *NodeEval {
    e := NewNodeEval()
    if node != nil {
        
        rlog3(indent,"EvalFrameStmtJump: %v\n", node)
        val1 := f.EvalFrameExpr(node.Expr, indent+1)
        e = val1
        if val1.IdType == IDENTIFIER {
            localVal := f.GetLocal(val1.Value,indent+1)
            e.IdType = localVal.SpecType.Id
            
            rlog(indent,"EvalFrameStmtJump IDENTIFIER: %v %v\n", localVal, localVal.SpecType.Id)
            if e.IdType == TYPE_BOOL_ID {
                e.BoolValue = localVal.BoolValue
                e.Value = fmt.Sprintf("%v", localVal.BoolValue)
            }
            if e.IdType == INTEGER_ID || e.IdType == TYPE_INT_ID {
                e.IntValue = localVal.IntValue
                e.Value = fmt.Sprintf("%v", localVal.IntValue)
            }
            if e.IdType == FLOAT_ID || e.IdType == TYPE_FLOAT_ID {
                e.FloatValue = localVal.FloatValue
                e.Value = fmt.Sprintf("%v", localVal.IntValue)
            }
            if e.IdType == LITERAL_STRING || e.IdType == TYPE_STRING_ID {
                e.StringValue = localVal.Value
                e.Value = localVal.StringValue
                
                rlog(indent,"EvalFrameStmtJump IDENTIFIER LITERAL_STRING: %v \n", e)
            }
        }
        
        rlog(indent,"EvalFrameStmtJump: Expr=%v NodeEval=%v\n", val1, e)
    }
    return e
}

func (f *Frame) EvalFrameStmtIteration(node *StmtIteration, indent int) *NodeEval {
    e := NewNodeEval()
    if node != nil {
        
        rlog3(indent,"EvalFrameStmtIteration: %v\n", node)
        val1 := f.EvalFrameStmtWhile(node.StmtWhile, indent+1)
        val2 := f.EvalFrameStmtFor(node.StmtFor, indent+1)
        val3 := f.EvalFrameStmtDo(node.StmtDo, indent+1)
        val4 := f.EvalFrameStmtStep(node.StmtStep, indent+1)
        
        rlog(indent,"EvalFrameStmtIteration: StmtWhile=%v StmtFor=%v StmtDo=%v StmtStep=%v\n", val1, val2, val3, val4)
        if val1 != nil {
            e = val1
        }
        if val2 != nil {
            e = val2
        }
        if val3 != nil {
            e = val3
        }
        if val4 != nil {
            e = val4
        }
    }
    return e
}

func (f *Frame) EvalFrameStmtWhile(node *StmtWhile, indent int) *NodeEval {
    var inLoop bool
    inLoop = true
    var val1 *NodeEval
    var val2 *NodeEval
    e := NewNodeEval()

    if node != nil {
        rlog(indent,"EvalFrameStmtWhile entry: %v\n", node)
        evalNode1 := NewNodeEval()
        for inLoop {
            val1 = f.EvalFrameExpr(node.Expr, indent+1)
            rlog(indent,"EvalFrameStmtWhile in loop: val1.IdType=%v\n", val1)
            if val1.IdType == IDENTIFIER {
                localVar1 := f.GetLocal(val1.Value, indent+1)
                rlog(indent,"EvalFrameStmtWhile in loop: localVar1=%v\n", localVar1)
                evalNode1.Value = localVar1.Value
                evalNode1.IdType = localVar1.SpecType.Id
                evalNode1.IntValue = localVar1.IntValue
                evalNode1.FloatValue = localVar1.FloatValue
                evalNode1.BoolValue = localVar1.BoolValue
                evalNode1.StringValue = localVar1.StringValue
            }

            rlog(indent,"EvalFrameStmtWhile in loop: evalNode1=%v\n", evalNode1)
            if val1.IdType == TYPE_BOOL_ID {
                evalNode1.Value = val1.Value
                evalNode1.IdType = val1.IdType
                evalNode1.IntValue = val1.IntValue
                evalNode1.FloatValue = val1.FloatValue
                evalNode1.BoolValue = val1.BoolValue
                evalNode1.StringValue = val1.StringValue
                inLoop = evalNode1.BoolValue
                if evalNode1.BoolValue {
                    val2 = f.EvalFrameStmt(node.Stmt, indent+1)
                    rlog(indent,"EvalFrameStmt in while loop: Expr=%v Stmt=%v \n", val1, val2 )
                }
            } else {
                inLoop = false
                panic("While expression is not boolean")
            }
        }
        rlog(indent,"EvalFrameStmtWhile exit: Expr=%v Stmt=%v \n", val1, val2 )
    }
    return e
}

func (f *Frame) EvalFrameStmtFor(node *StmtFor, indent int) *NodeEval {
    var inLoop bool
    inLoop = true
    var val1 *NodeEval
    var val2 *NodeEval
    var val3 *NodeEval
    var val4 *NodeEval
    e := NewNodeEval()
    if node != nil {
        rlog(indent,"EvalFrameStmtFor: %v\n", node)
        val1 = f.EvalFrameExprStmt(node.ExprStmt1, indent+1)
        val2 = f.EvalFrameExprStmt(node.ExprStmt2, indent+1)
        for inLoop {
            evalNode2 := NewNodeEval()
            if val2.IdType == IDENTIFIER {
                localVar2 := f.GetLocal(val2.Value, indent+1)
                rlog(indent,"EvalFrameStmtWhile in loop: localVar2=%v\n", localVar2)
                evalNode2.Value = localVar2.Value
                evalNode2.IdType = localVar2.SpecType.Id
                evalNode2.IntValue = localVar2.IntValue
                evalNode2.FloatValue = localVar2.FloatValue
                evalNode2.BoolValue = localVar2.BoolValue
                evalNode2.StringValue = localVar2.StringValue
            }
            if val2.IdType == TYPE_BOOL_ID {
                evalNode2.Value = val2.Value
                evalNode2.IdType = val2.IdType
                evalNode2.IntValue = val2.IntValue
                evalNode2.FloatValue = val2.FloatValue
                evalNode2.BoolValue = val2.BoolValue
                evalNode2.StringValue = val2.StringValue
                inLoop = evalNode2.BoolValue
                if evalNode2.BoolValue {
                    val3 = f.EvalFrameStmt(node.Stmt, indent+1)
                    val4 = f.EvalFrameExpr(node.Expr, indent+1)
                    val2 = f.EvalFrameExprStmt(node.ExprStmt2, indent+1)
                    rlog(indent,"EvalFrameStmt in while loop: Expr=%v Stmt=%v \n", val3, val4 )
                }
            }
        }
        rlog(indent,"EvalFrameStmtFor: Expr=%v ExprStmt1=%v ExprStmt2=%v Stmt=%v \n", val1, val2, val3, val4  )
    }
    return e
}

func (f *Frame) EvalFrameStmtDo(node *StmtDo, indent int) *NodeEval {
    var inLoop bool
    inLoop = true
    var val1 *NodeEval
    var val2 *NodeEval
    e := NewNodeEval()

    if node != nil {
        rlog(indent,"EvalFrameStmtDo entry: %v\n", node)
        evalNode1 := NewNodeEval()
        for inLoop {
            val1 = f.EvalFrameExpr(node.Expr, indent+1)
            rlog(indent,"EvalFrameStmtDo in loop: val1.IdType=%v\n", val1)
            if val1.IdType == IDENTIFIER {
                localVar1 := f.GetLocal(val1.Value, indent+1)
                rlog(indent,"EvalFrameStmtDo in loop: localVar1=%v\n", localVar1)
                evalNode1.Value = localVar1.Value
                evalNode1.IdType = localVar1.SpecType.Id
                evalNode1.IntValue = localVar1.IntValue
                evalNode1.FloatValue = localVar1.FloatValue
                evalNode1.BoolValue = localVar1.BoolValue
                evalNode1.StringValue = localVar1.StringValue
            }

            rlog(indent,"EvalFrameStmtDo : evalNode1=%v\n", evalNode1)
            if val1.IdType == TYPE_BOOL_ID {
                evalNode1.Value = val1.Value
                evalNode1.IdType = val1.IdType
                evalNode1.IntValue = val1.IntValue
                evalNode1.FloatValue = val1.FloatValue
                evalNode1.BoolValue = val1.BoolValue
                evalNode1.StringValue = val1.StringValue
                inLoop = evalNode1.BoolValue
                if evalNode1.BoolValue {
                    val2 = f.EvalFrameStmt(node.Stmt, indent+1)
                    rlog(indent,"EvalFrameStmtDo in eval stmt in loop: Expr=%v Stmt=%v \n", val1, val2 )
                }
            } else {
                inLoop = false
                panic("While expression is not boolean")
            }
        }

        rlog(indent,"EvalFrameStmtDo exit: Expr=%v Stmt=%v \n", val1, val2 )
    }
    return e
}

func (f *Frame) EvalFrameStmtSelection(node *StmtSelection, indent int) *NodeEval {
    e := NewNodeEval()
    if node != nil {
        
        rlog3(indent,"EvalFrameStmtSelection:%v \n", node)
        ifExpr := f.EvalFrameExpr(node.Expr, indent+1)
        e = ifExpr
        
        rlog(indent,"EvalFrameStmtSelection ifExpr: %v\n", ifExpr)
        if ifExpr.IdType == IDENTIFIER {
            localVal := f.GetLocal(ifExpr.Value,indent+1)
            if localVal.SpecType.Id == TYPE_BOOL_ID {
                if localVal.BoolValue {
                    rlog(indent,"TRUE SCOPE 0\n")
                    f.EvalFrameStmt(node.Stmt, indent+1)
                } else {
                    rlog(indent,"FALSE SCOPE 0\n")
                    f.EvalFrameStmt(node.Stmt2, indent+1)
                }
            }
        } else if ifExpr.IdType == TRUE_ID {
            rlog(indent,"TRUE SCOPE 1\n")
            f.EvalFrameStmt(node.Stmt, indent+1)
        }  else if ifExpr.IdType == FALSE_ID {
            rlog(indent,"FALSE SCOPE1\n")
            f.EvalFrameStmt(node.Stmt2, indent+1)
        } else if ifExpr.IdType == TYPE_BOOL_ID  {
            if ifExpr.BoolValue {
                rlog(indent,"TRUE SCOPE 2\n")
                f.EvalFrameStmt(node.Stmt, indent+1)
            } else {
                rlog(indent,"FALSE SCOPE2\n")
                f.EvalFrameStmt(node.Stmt2, indent+1)
            }
        } 
    }
    return e
}

func (f *Frame) EvalFrameStmt(node *Stmt, indent int) *NodeEval {
    e := NewNodeEval()
    if node != nil {
        
        rlog6(indent,"EvalFrameStmt 24: %v\n", node)
        val1 := f.EvalFrameExprCmd(node.ExprCmd, indent+1)
        val2 := f.EvalFrameExprStmt(node.ExprStmt, indent+1)
        val3 := f.EvalFrameStmtJump(node.StmtJump, indent+1)
        val4 := f.EvalFrameStmtIteration(node.StmtIteration, indent+1)
        val5 := f.EvalFrameStmtSelection(node.StmtSelection, indent+1)
        val6 := f.EvalFrameCmpndStmt(node.CmpndStmt, indent+1)
        val7 := f.EvalFrameStmtStep(node.StmtStep, indent+1)
        

        rlog6(indent,"EvalFrameStmt vals: ExprCmd=%v ExprStmt=%v StmtJump=%v StmtIteration=%v StmtSelection=%v CmpndStmt=%v StmtStep=%v\n", val1, val2, val3, val4, val5, val6, val7)

        rlog6(indent,"EvalFrameStmt ExprCmd 25: %v \n", val1)
        if len(val1.Value) > 0 {
            rlog6(indent,"EvalFrameStmt ExprCmd: %v \n", val1)
        }

        if val1.IdType > 0 {
            e = val1
            rlog6(indent,"EvalFrameStmt ExprCmd: \n", e)
        }
        if val2.IdType > 0 {
            e = val2
            
            rlog(indent,"EvalFrameStmt ExprStmt: \n", e)
        }
        if val3.IdType > 0 {
            e = val3
            
            rlog(indent,"EvalFrameStmt StmtJump: \n", e)
        }
        if val4.IdType > 0 {
            e = val4
            
            rlog(indent,"EvalFrameStmt StmtIteration: \n", e)
        }
        if val5.IdType > 0 {
            e = val5
            
            rlog(indent,"EvalFrameStmt StmtSelection: \n", e)
        }
        if val6.IdType > 0 {
            e = val6
            
            rlog(indent,"EvalFrameStmt CmpndStmt: \n", e)
        }
        if val7.IdType > 0 {
            e = val7
            
            rlog(indent,"EvalFrameStmt StmtStep: \n", e)
        }
    }
    return e
}

func (f *Frame) EvalFrameCmpndStmt(node *CmpndStmt, indent int) *NodeEval  {
    e := NewNodeEval()
    if node != nil {
        
        rlog3(indent,"EvalFrameCmpndStmt: %v\n", node)
        val1 := f.EvalFrameStmtList(node.StmtList, indent+1)
        val2 := f.EvalFrameDeclarationList(node.DeclarationList, indent+1)
        
        rlog(indent,"EvalFrameCmpndStmt vals: StmtList=%v DeclarationList=%v \n", val1, val2)
        if val1.IdType > 0 {
            e = val1
            
            rlog(indent,"EvalFrameCmpndStmt StmtList: \n", e)
        }
        if val2.IdType > 0 {
            e = val2
            
            rlog(indent,"EvalFrameCmpndStmt DeclarationList: \n", e)
        }
    }
    return e
}

func (f *Frame) EvalFrameDeclarationList(node *DeclarationList, indent int) *NodeEval {
    e := NewNodeEval()
    if node != nil {
        
        rlog(indent,"EvalFrameDeclarationList: %v\n", node)
    }
    return e
}

func (f *Frame) EvalFrameStmtList(node *StmtList, indent int) *NodeEval {
    e := NewNodeEval()
    if node != nil {
        
        rlog3(indent,"EvalFrameStmtList: %v\n", node)
        val1 := f.EvalFrameStmtList(node.StmtList, indent+1)
        val2 := f.EvalFrameStmt(node.Stmt, indent+1)
        
        rlog(indent,"EvalFrameStmtList vals: StmtList=%v Stmt=%v\n", val1, val2)

        if val1.IdType > 0 {
            e = val1
            rlog(indent,"EvalFrameStmtList StmtList: %v \n", e)
        }
        if val2.IdType > 0 {
            e = val2
            rlog(indent,"EvalFrameStmtList Stmt: %v \n", e)
        }
    }
    return e
}


func (f *Frame) EvalFrameStmtStep(node *StmtStep, indent int) *NodeEval {
    e := NewNodeEval()
    
    if node != nil {
        
        rlog3(indent,"EvalFrameStmtStep: %v\n", node)

        mStep := NewStep()
        mStep.Identifier = node.Identifier
        mStep.CmpndStmtStimulus = node.CmpndStmtStimulus
        mStep.CmpndStmtResponse = node.CmpndStmtResponse
        mStep.StepDocument = new(StepDocument)

        val1 := f.EvalFrameExpr(node.Expr, indent+1)
        val2 := f.EvalFrameCmpndStmt(node.CmpndStmtStimulus, indent+1)
        val3 := f.EvalFrameCmpndStmt(node.CmpndStmtResponse, indent+1)
        val4 := f.EvalFrameIdentifier(node.Identifier, indent+1)

        e.StepDocument = new(StepDocument)
        localVarText := f.GetLocal("stimulusText",indent+1)
        rlog3(indent,"stimulusText: %v\n", localVarText)
        if localVarText != nil {
            e.StepDocument.TextStimulus = localVarText.Value
            mStep.StepDocument.TextStimulus = e.StepDocument.TextStimulus 
        }
        localVarText = f.GetLocal("responseText",indent+1)
        if localVarText != nil {
            e.StepDocument.TextResponse =  localVarText.Value
            mStep.StepDocument.TextResponse = e.StepDocument.TextResponse 
            rlog5(indent,"StepDocument: TextStimulus=%v TextResponse=%v\n", e.StepDocument.TextStimulus, e.StepDocument.TextResponse)
        }
        localVarText = f.GetLocal("stepPass",indent+1)
        if localVarText != nil {
            e.StepDocument.Pass = localVarText.BoolValue
            mStep.StepDocument.Pass = e.StepDocument.Pass
        }
        rlog5(indent+1,"EvalFrameStmtStep mStep: %v\n", mStep)
        rlog5(indent+1,"EvalFrameStmtStep mCurrTestPlan: %v\n", mCurrTestPlan)
        mPlan := mTestPlanList.GetTestPlan(mCurrTestPlan.Name)
        mPlan.AddTestPlanStep(mStep, indent+1)
        rlog5(indent+1,"EvalFrameStmtStep mPlan: %v\n", mPlan)

        localVarText = f.GetLocal("testPlanText",indent+1)
        if localVarText != nil {
            mPlan.TestPlanText = localVarText.Value
            rlog6(indent,"HandleTestPlanTree testPlanText: found\n")
        }


        rlog(indent,"EvalFrameStmtStep vals: Expr=%v CmpndStmtStimulus=%v CmpndStmtResponse=%v Identifier=%v \n", val1, val2, val3, val4)
        if val1.IdType > 0 {
            e = val1
            
            rlog(indent,"EvalFrameStmtStep Expr: \n", e)
        }
        if val2.IdType > 0 {
            e = val2
            
            rlog(indent,"EvalFrameStmtStep CmpndStmtStimulus: \n", e)
        }
        if val3.IdType > 0 {
            e = val3
            
            rlog(indent,"EvalFrameStmtStep CmpndStmtResponse: \n", e)
        }
        if val4.IdType > 0 {
            e = val4
            
            rlog(indent,"EvalFrameStmtStep Identifier: \n", e)
        }
    }
    return e
}

