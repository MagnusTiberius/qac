package model

import (
    "fmt"
    "io/ioutil"
    "path"
)

type ImportCheckList struct {
    Imports     []string
}
var mImportCheckList *ImportCheckList

func NewImportCheckList() *ImportCheckList {
    if mImportCheckList == nil {
        mImportCheckList = new(ImportCheckList)
        mImportCheckList.Imports = []string{}
    }
    return mImportCheckList
}

func (icl *ImportCheckList) Add(name string) {
    mImportCheckList.Imports = append(mImportCheckList.Imports, name)
}

func (icl *ImportCheckList) Exists(name string) bool {
    for i := range mImportCheckList.Imports {
        itm := mImportCheckList.Imports[i]
        if itm == name {
            return true
        }
    }
    return false
}


func HandleImport(rtu *RtTranslationUnit, indent int) {
    var mRtImportList *RtImportList
    mImportCheckList = NewImportCheckList() 
    mRtImportList = rtu.RtImportList
    if mRtImportList != nil {
        for i := range mRtImportList.List {
            mItem := mRtImportList.List[i]
            rlog5(indent, "Import: %v \n", mItem)
            if !mImportCheckList.Exists(mItem) {
                ProcessImportItem(mItem, indent)
            }
        }
    }
}

func ProcessImportItem(imp string, indent int) {
    mImportCheckList.Add(imp)
    pathRoot := GetRootPath()
    importPath := fmt.Sprintf("%v%v", pathRoot, imp)
    rlog5(indent, "importPath: %v \n", importPath)
    mList, _ := ioutil.ReadDir(importPath)
    rlog5(indent, "mList: %v \n", mList)
    for i := range mList {
        mFileInfo := mList[i]
        //rlog5(indent, "mFileInfo: %v \n", mFileInfo)
        mFile := mFileInfo.Name()
        mExt := path.Ext(mFile)
        //rlog5(indent, "mExt: %v \n", mExt)
        if mExt == ".ias" {
            rlog5(indent, "mExt: %s \n", mFile)
            DoImport(importPath, mFile, indent)
        }
    }
}

func DoImport(pImportPath string, pFile string, indent int) {
    mFile := fmt.Sprintf("%v/%v", pImportPath, pFile)
    rlog5(indent, "mFile: %s \n", mFile)
    //mProgramUnit := LoadFile(mFile)
    //rlog5(indent, "mProgramUnit: %s \n", mProgramUnit)
    mRtTranslationUnit = Parse(mFile)
    rlog5(indent, "mRtTranslationUnit: %v \n", mRtTranslationUnit)
    mRtTranslationUnitList.AddRtTranslationUnit(mRtTranslationUnit)
    HandleImport(mRtTranslationUnit,indent+1)
}


func ValidateImports(indent int) {
    mUnits := mRtTranslationUnitList.RtTranslationUnits
    for i := range mUnits {
        mUnit := mUnits[i]
        rlog5(indent, "TranslationUnit[%v]: %v \n",i+1, mUnit)
    }
}
