package model

import (
 	"strings"
 	"fmt"
)


func (f *Frame) HandleUcase(val1 *NodeEval, mParams []NodeEval, indent int) string {
    evalNode1 := NewNodeEval()
    rlog8(indent,"HandleUcase IdType=%v IdType=%v \n",mParams[0].IdType,mParams[1].IdType)

    if mParams[0].IdType == TYPE_STRING_ID || mParams[0].IdType == LITERAL_STRING  {
        evalNode1.Value = mParams[0].Value
        evalNode1.StringValue = mParams[0].StringValue
        evalNode1.IdType = TYPE_STRING_ID
        //rlog7(indent,"HandleUcase TYPE_STRING_ID1 Value=%v StringValue=%v \n", val1.Value, val1.StringValue)
    } 

    rlog8(indent,"HandleUcase evalNode1= %v \n", evalNode1)

    if evalNode1.IdType == TYPE_STRING_ID || evalNode1.IdType == LITERAL_STRING {
        p1 := evalNode1.Value
        rv := strings.ToLower(p1)
        rlog8(indent,"HandleUcase ucase %v \n", rv)
        return rv
    }
    return ""

}

func (f *Frame) HandleLcase(val1 *NodeEval, mParams []NodeEval, indent int) string {
    evalNode1 := NewNodeEval()
    rlog8(indent,"HandleLcase IdType=%v IdType=%v \n",mParams[0].IdType,mParams[1].IdType)

    if mParams[0].IdType == TYPE_STRING_ID || mParams[0].IdType == LITERAL_STRING  {
        evalNode1.Value = mParams[0].Value
        evalNode1.StringValue = mParams[0].StringValue
        evalNode1.IdType = TYPE_STRING_ID
        //rlog7(indent,"HandleLcase TYPE_STRING_ID1 Value=%v StringValue=%v \n", val1.Value, val1.StringValue)
    } 

    rlog8(indent,"HandleLcase evalNode1= %v \n", evalNode1)

    if evalNode1.IdType == TYPE_STRING_ID || evalNode1.IdType == LITERAL_STRING {
        p1 := evalNode1.Value
        rv := strings.ToUpper(p1)
        rlog8(indent,"HandleLcase lcase %v \n", rv)
        return rv
    }
    return ""

}


func (f *Frame) HandleReplace(val1 *NodeEval, mParams []NodeEval, indent int) string {
    evalNode1 := NewNodeEval()
    evalNode2 := NewNodeEval()
    evalNode3 := NewNodeEval()
    evalNode4 := NewNodeEval()
    rlog8(indent,"HandleRepeat IdType=%v IdType=%v \n",mParams[0].IdType,mParams[1].IdType)

    if mParams[0].IdType == TYPE_STRING_ID || mParams[0].IdType == LITERAL_STRING  {
        evalNode1.Value = mParams[0].Value
        evalNode1.StringValue = mParams[0].StringValue
        evalNode1.IdType = TYPE_STRING_ID
        //rlog7(indent,"HandleRepeat TYPE_STRING_ID1 Value=%v StringValue=%v \n", val1.Value, val1.StringValue)
    } 
    if mParams[1].IdType == TYPE_STRING_ID || mParams[1].IdType == LITERAL_STRING  {
        evalNode2.Value = mParams[1].Value
        evalNode2.StringValue = mParams[1].StringValue
        evalNode2.IdType = TYPE_STRING_ID
        //rlog7(indent,"HandleRepeat TYPE_STRING_ID1 Value=%v StringValue=%v \n", val1.Value, val1.StringValue)
    } 
    if mParams[2].IdType == TYPE_STRING_ID || mParams[2].IdType == LITERAL_STRING  {
        evalNode3.Value = mParams[2].Value
        evalNode3.StringValue = mParams[2].StringValue
        evalNode3.IdType = TYPE_STRING_ID
        //rlog7(indent,"HandleRepeat TYPE_STRING_ID1 Value=%v StringValue=%v \n", val1.Value, val1.StringValue)
    } 
    if mParams[3].IdType == TYPE_INT_ID || mParams[3].IdType == INTEGER_ID  {
        evalNode4.Value = mParams[3].Value
        evalNode4.StringValue = fmt.Sprintf("%v",evalNode1.Value)
        evalNode4.IntValue = mParams[3].IntValue
        evalNode4.IdType = TYPE_INT_ID
        //rlog7(indent,"HandleRepeat TYPE_STRING_ID1 Value=%v StringValue=%v \n", val1.Value, val1.StringValue)
    } 

    rlog8(indent,"HandleRepeat evalNode1= %v \n", evalNode1)

    if evalNode1.IdType == TYPE_STRING_ID || evalNode1.IdType == LITERAL_STRING {
	    if evalNode2.IdType == TYPE_STRING_ID || evalNode2.IdType == LITERAL_STRING {
		    if evalNode3.IdType == TYPE_STRING_ID || evalNode3.IdType == LITERAL_STRING {
			    if evalNode4.IdType == TYPE_INT_ID || evalNode4.IdType == INTEGER_ID {
			        p1 := evalNode1.Value
			        p2 := evalNode2.Value
			        p3 := evalNode2.Value
			        p4 := evalNode2.IntValue
			        rv := strings.Replace(p1, p2, p3, p4)
			        rlog8(indent,"HandleRepeat repeat %v \n", rv)
			        return rv
			    }
			}
	    }
    }
    return ""
}

func (f *Frame) HandleRepeat(val1 *NodeEval, mParams []NodeEval, indent int) string {
    evalNode1 := NewNodeEval()
    evalNode2 := NewNodeEval()
    rlog8(indent,"HandleRepeat IdType=%v IdType=%v \n",mParams[0].IdType,mParams[1].IdType)

    if mParams[0].IdType == TYPE_STRING_ID || mParams[0].IdType == LITERAL_STRING  {
        evalNode1.Value = mParams[0].Value
        evalNode1.StringValue = mParams[0].StringValue
        evalNode1.IdType = TYPE_STRING_ID
        //rlog7(indent,"HandleRepeat TYPE_STRING_ID1 Value=%v StringValue=%v \n", val1.Value, val1.StringValue)
    } 
    if mParams[1].IdType == TYPE_INT_ID || mParams[1].IdType == INTEGER_ID  {
        evalNode2.Value = mParams[1].Value
        evalNode2.StringValue = fmt.Sprintf("%v",evalNode1.Value)
        evalNode2.IntValue = mParams[1].IntValue
        evalNode2.IdType = TYPE_INT_ID
        //rlog7(indent,"HandleRepeat TYPE_STRING_ID1 Value=%v StringValue=%v \n", val1.Value, val1.StringValue)
    } 

    rlog8(indent,"HandleRepeat evalNode1= %v \n", evalNode1)

    if evalNode1.IdType == TYPE_STRING_ID || evalNode1.IdType == LITERAL_STRING {
	    if evalNode2.IdType == TYPE_INT_ID || evalNode2.IdType == INTEGER_ID {
	        p1 := evalNode1.Value
	        p2 := evalNode2.IntValue
	        rv := strings.Repeat(p1, p2)
	        rlog8(indent,"HandleRepeat repeat %v \n", rv)
	        return rv
	    }
    }
    return ""
}


func (f *Frame) HandleInstr(val1 *NodeEval, mParams []NodeEval, indent int) int {
    evalNode1 := NewNodeEval()
    evalNode2 := NewNodeEval()
    rlog8(indent,"HandleInstr IdType=%v IdType=%v \n",mParams[0].IdType,mParams[1].IdType)
    if mParams[0].IdType == IDENTIFIER {
        localVar1 := f.GetLocal(mParams[0].Value, indent+1)
        evalNode1.Value = localVar1.Value
        evalNode1.IdType = localVar1.SpecType.Id
        evalNode1.IntValue = localVar1.IntValue
        evalNode1.FloatValue = localVar1.FloatValue
        evalNode1.BoolValue = localVar1.BoolValue
        evalNode1.StringValue = evalNode1.Value
        //
        rlog8(indent,"HandleInstr IDENTIFIER1 StringValue=%v Value1=%v IdType1=%v \n",  evalNode1.StringValue, evalNode1.Value, evalNode1.IdType  )
    }
    if mParams[1].IdType == IDENTIFIER {
        localVar2 := f.GetLocal(mParams[1].Value, indent+1)
        evalNode2.Value = localVar2.Value
        evalNode2.IdType = localVar2.SpecType.Id
        evalNode2.IntValue = localVar2.IntValue
        evalNode2.FloatValue = localVar2.FloatValue
        evalNode2.BoolValue = localVar2.BoolValue
        evalNode2.StringValue = evalNode2.Value
        //
        rlog8(indent,"HandleInstr IDENTIFIER2 StringValue=%v Value2=%v IdType2=%v \n",  evalNode2.StringValue, evalNode2.Value, evalNode2.IdType  )
    }

    if mParams[0].IdType == TYPE_STRING_ID || mParams[0].IdType == LITERAL_STRING  {
        evalNode1.Value = mParams[0].Value
        evalNode1.StringValue = mParams[0].StringValue
        evalNode1.IdType = TYPE_STRING_ID
        //rlog7(indent,"HandleInstr TYPE_STRING_ID1 Value=%v StringValue=%v \n", val1.Value, val1.StringValue)
    } 

    if mParams[1].IdType == TYPE_STRING_ID || mParams[1].IdType == LITERAL_STRING  {
        evalNode2.Value = mParams[1].Value
        evalNode2.StringValue = mParams[1].StringValue
        evalNode2.IdType = TYPE_STRING_ID
        //rlog7(indent,"HandleInstr TYPE_STRING_ID2 Value=%v StringValue=%v \n", val2.Value, val2.StringValue)
    } 

    rlog8(indent,"HandleInstr evalNode1= %v \n", evalNode1)
    rlog8(indent,"HandleInstr evalNode2= %v \n", evalNode2)

    if evalNode1.IdType == TYPE_STRING_ID || evalNode1.IdType == LITERAL_STRING {
        if evalNode2.IdType == TYPE_STRING_ID || evalNode2.IdType == LITERAL_STRING {
            p1 := evalNode1.Value
            p2 := evalNode2.Value
            rv := strings.Index(p1, p2)
            rlog8(indent,"HandleInstr instr %v \n", rv)
            return rv
        }
    }
    return -1
}