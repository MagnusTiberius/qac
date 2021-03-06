package model

import (
    //"fmt"
)

type FunctionCall struct {
    List []string
}
func NewFunctionCall() *FunctionCall {
    v := new(FunctionCall)
    v.List = []string{}
    return v
}
func (fc *FunctionCall) Add(p string) {
    fc.List = append(fc.List, p)
}
func (fc *FunctionCall) Reverse() {
    mList := []string{}
    for i := range fc.List {
        itm := fc.List[len(fc.List)-i-1]
        _ = itm
        mList = append(mList, itm)
    }
    fc.List = mList
}

func (f *Frame) EvalFrameFcExprPostfix(node *ExprPostfix, indent int, fc *FunctionCall) {
    if node != nil {
        
        rlog(indent,"EvalFrameFcExprPostfix: \n", node)
        f.EvalFrameFcIdentifier(node.Identifier, indent+1, fc)
        f.EvalFrameFcExprPostfix(node.ExprPostfix, indent+1, fc)
        f.EvalFrameFcExprPrimary(node.ExprPrimary, indent+1, fc)
    }
}

func (f *Frame) EvalFrameFcIdentifier(node *Identifier, indent int, fc *FunctionCall) {
    if node != nil {
        
        rlog(indent,"EvalFrameFcIdentifier: \n", node)
        fc.Add(node.Name)
    }
}

func (f *Frame) EvalFrameFcExprPrimary(node *ExprPrimary, indent int, fc *FunctionCall) {
    if node != nil {
        
        rlog(indent,"EvalFrameFcExprPrimary: \n", node)

        if len(node.Value) > 0 {
            fc.Add(node.Value)
        }

        //f.EvalFrameExprUnary(node.ExprUnary, indent+1)
    /*
    Value       string
    IdType      int
    Expr        *Expr
    ExprAssign  *ExprAssign
    */
        f.EvalFrameExpr(node.Expr, indent+1)
        f.EvalFrameExprAssign(node.ExprAssign, indent+1)
        f.EvalFrameNodeValue(node.Value, indent+1)
    }
}
