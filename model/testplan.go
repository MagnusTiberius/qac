package model

import (
    "fmt"
    "time"
)

type StepDocument struct {
    TextStimulus string
    TextResponse string
    Pass         bool
}

// ====================================================
//
//
type Step struct {
    Sequence            int
    Stimulus            string
    Response            string
    Expr                *Expr
    CmpndStmtStimulus   *CmpndStmt
    CmpndStmtResponse   *CmpndStmt
    Identifier          *Identifier
    StepDocument        *StepDocument
}

func NewStep() *Step {
    v := new(Step)
    return v
}
// ====================================================
//
//

type TestPlan struct {
    Name            string
    Steps           []Step
    TestPlanText    string
}

var mCurrTestPlan *TestPlan

func NewTestPlan() *TestPlan {
    v := new(TestPlan)
    v.Steps = []Step{}
    mCurrTestPlan = v
    return v
}
func (tp *TestPlan) AddTestPlanStep(s *Step, indent int) {
    tp.Steps = append(tp.Steps, *s)
    rlog4(indent,"AddTestPlanStep: %v %v\n", s, s.Identifier)
}

func GetCurrentTestPlan() *TestPlan  {
    return mCurrTestPlan
}
// ====================================================
//
//
type TestPlanList struct {
    TestPlans    []TestPlan
}

func (tpl *TestPlanList) AddTestPlan(tp *TestPlan) {
    tpl.TestPlans = append(tpl.TestPlans, *tp)
    //mCurrTestPlan = tp
}

func (tpl *TestPlanList) GetTestPlan(name string) *TestPlan {
    for i := range mTestPlanList.TestPlans {
        mCurrTestPlan := mTestPlanList.TestPlans[i]
        if mCurrTestPlan.Name == name {
            return &mTestPlanList.TestPlans[i]
        }
    }
    return nil
}

var mTestPlanList *TestPlanList

func NewTestPlanList() *TestPlanList {
    v := new(TestPlanList)
    v.TestPlans = []TestPlan{}
    return v
}

func InitRuntimeTestPlanList() *TestPlanList {
    if mTestPlanList == nil {
        mTestPlanList = NewTestPlanList()
    }
    return mTestPlanList
}

func GetRuntimeTestPlanList() *TestPlanList {
    mTestPlanList = InitRuntimeTestPlanList()
    return mTestPlanList
}


func (tpl *TestPlanList) EmitTestPlanResultToHtml() string {
    var mOut string
    mOut = "<html>\n"
    mOut = fmt.Sprintf("%v<h1>IAS Testplans</h1>\n",mOut)
    for i := range mTestPlanList.TestPlans {
        mCurrTestPlan := mTestPlanList.TestPlans[i]
        rlog5(1, "mCurrTestPlan[%v]: %v \n", i,mCurrTestPlan)
        mOut = fmt.Sprintf("%v<h3>\n",mOut)
        mOut = fmt.Sprintf("%v%v\n", mOut, mCurrTestPlan.Name)
        mOut = fmt.Sprintf("%v</h3>\n",mOut)
        mOut = fmt.Sprintf("%v<table border=1>\n",mOut)
        mOut = fmt.Sprintf("%v<tr>\n",mOut)
        mOut = fmt.Sprintf("%v<td width='50'>Step</td><td>Stimulus</td><td>Response</td><td>Pass</td>\n",mOut)
        mOut = fmt.Sprintf("%v</tr>\n",mOut)
        for j := range mCurrTestPlan.Steps {
            mOut = fmt.Sprintf("%v<tr>\n",mOut)
            mCurrStep := mCurrTestPlan.Steps[j]
            mOut = fmt.Sprintf("%v<td width='100'>%v</td>\n",mOut,j+1)
            mOut = fmt.Sprintf("%v<td width='400'>\n",mOut)
            mOut = fmt.Sprintf("%v%v\n",mOut,mCurrStep.StepDocument.TextStimulus)
            mOut = fmt.Sprintf("%v</td>\n",mOut)
            mOut = fmt.Sprintf("%v<td width='400'>\n",mOut)
            mOut = fmt.Sprintf("%v%v\n",mOut,mCurrStep.StepDocument.TextResponse)
            mOut = fmt.Sprintf("%v</td>\n",mOut)
            mOut = fmt.Sprintf("%v<td width='100'>\n",mOut)
            mOut = fmt.Sprintf("%v%v\n",mOut,mCurrStep.StepDocument.Pass)
            mOut = fmt.Sprintf("%v</td>\n",mOut)
            mOut = fmt.Sprintf("%v</tr>\n",mOut)
        }
        mOut = fmt.Sprintf("%v</table><br><br>\n",mOut)
    }
    mOut = fmt.Sprintf("%v</html>\n",mOut)
    return mOut
}

func (tpl *TestPlanList) EmitTestPlanResultToHtml2() string {
    var mOut string
    var mTRow, mTRow2 string
    var mStyle string
    mStyle = `
        <style>
            .tpHdr {}
            .seq, .stim,.resp, .reslt {float:left}
            .seq, .reslt {width:100px;}
            .stim, .resp {width:300px;}
            .mg1 {padding:12px 10px; }
            .mg2 {padding:12px 0px}
            .bgc1 {background-color:#CCC}
            .fs1 {font-size:5em; color:#CCC}
            .bgc2 {color:#FFF; background-color:#444}
            .center1 {margin-left: auto; margin-right: auto; width:880px;}
            .ovflw {overflow: auto;}
        </style>
    `
    _ = mStyle
    mTRow = `
        %v
        <div class='seq mg1 fs1'>
            %v
        </div>
        <div class='stim mg1'>
            %v
        </div>
        <div class='resp mg1 ovflw'>
            %v
        </div>
        <div class='reslt mg1 '>
            %v
        </div>
        <div style='clear:both'></div>
    `
    mTRow2 = `
        %v
        <div class='seq mg1 bgc2'>
            %v
        </div>
        <div class='stim mg1 bgc2'>
            %v
        </div>
        <div class='resp mg1 bgc2'>
            %v
        </div>
        <div class='reslt mg1 bgc2'>
            %v
        </div>
        <div style='clear:both'></div>
    `
    _ = mTRow2
    mOut = "<html>\n"
    mOut = fmt.Sprintf("%v%v\n",mOut,mStyle)
    mOut = fmt.Sprintf("%v<body>\n",mOut)
    mOut = fmt.Sprintf("%v<div class='center1'>\n",mOut)
    mOut = fmt.Sprintf("%v<h1>Testcase Report</h1>\n",mOut)
    t := time.Now()
    mOut = fmt.Sprintf("%v<span style='font-size:0.9em;'>%v</span>\n",mOut,t)
    for i := range mTestPlanList.TestPlans {
        mCurrTestPlan := mTestPlanList.TestPlans[i]
        rlog5(1, "mCurrTestPlan[%v]: %v \n", i,mCurrTestPlan)
        mOut = fmt.Sprintf("%v<br><br><h3>\n",mOut)
        mOut = fmt.Sprintf("%v%v\n", mOut, mCurrTestPlan.Name)
        mOut = fmt.Sprintf("%v</h3>\n",mOut)
        mOut = fmt.Sprintf("%v<div class='mg2'>%v</div>\n",mOut,mCurrTestPlan.TestPlanText)
        mOut = fmt.Sprintf(mTRow2,mOut, "Step", "Stimulus", "Response", "Pass")
        for j := range mCurrTestPlan.Steps {
            mCurrStep := mCurrTestPlan.Steps[j]
            sFormat := "<div style='height:3px; background-color:%v'></div>%v"
            sPass := ""
            if mCurrStep.StepDocument.Pass {
                sPass = fmt.Sprintf(sFormat, "green", true)                
            } else {
                sPass = fmt.Sprintf(sFormat, "red", false)                
            }
            mOut = fmt.Sprintf(mTRow,mOut, j+1, mCurrStep.StepDocument.TextStimulus, mCurrStep.StepDocument.TextResponse, sPass)
        }
    }
    mOut = fmt.Sprintf("%v</div>\n",mOut)
    mOut = fmt.Sprintf("%v</body>\n",mOut)
    mOut = fmt.Sprintf("%v</html>\n",mOut)
    return mOut
}

// ---------------------------------------------------------------------------

func UnloadTestPlan() {
    if mTestPlanList == nil {
        return
    }
    if len(mTestPlanList.TestPlans) == 0 {
        return
    }
    for i := range mTestPlanList.TestPlans {
        tp := mTestPlanList.TestPlans[i]
        if len(tp.Steps) > 0 {
            for j := range tp.Steps {
                step := tp.Steps[j]
                UnloadExpr(step.Expr)
                UnloadCmpndStmt(step.CmpndStmtStimulus)
                UnloadCmpndStmt(step.CmpndStmtResponse)
                step.CmpndStmtStimulus   = nil
                step.CmpndStmtResponse   = nil
                step.Identifier          = nil
                step.StepDocument        = nil
            }
        }
    }
}

