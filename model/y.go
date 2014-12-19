//line grammar.y:73
package model

import __yyfmt__ "fmt"

//line grammar.y:74
import (
// "fmt"
)

//line grammar.y:86
type yySymType struct {
	yys                 int
	val                 int
	id                  string
	sel                 []string
	tExprAdditive       *ExprAdditive
	tExprMultiplicative *ExprMultiplicative
	tExprCast           *ExprCast
	tExprUnary          *ExprUnary
	tExprPrimary        *ExprPrimary
	tExprCmd            *ExprCmd
	tExprPostfix        *ExprPostfix
	tExprAssign         *ExprAssign
	tExpr               *Expr
	tExprStmt           *ExprStmt
	tStmt               *Stmt
	tStmtList           *StmtList
	tCmpndStmt          *CmpndStmt
	tDeclarationList    *DeclarationList
	tDeclaration        *Declaration
	tExprShift          *ExprShift
	tExprRelational     *ExprRelational
	tExprEquality       *ExprEquality
	tExprAnd            *ExprAnd
	tExprExclOr         *ExprExclOr
	tExprInclOr         *ExprInclOr
	tExprLogicAnd       *ExprLogicAnd
	tExprLogicOr        *ExprLogicOr
	tExprConditional    *ExprConditional
	tSpecDecl           *SpecDecl
	tSpecType           *SpecType
	tDeclInitList       *DeclInitList
	tDeclInit           *DeclInit
	tDeclarator         *Declarator
	tInitializer        *Initializer
	tDeclDirect         *DeclDirect
	tFuncDef            *FuncDef
	tDeclExternal       *DeclExternal
	tTranslationUnit    *TranslationUnit
	tPackageDef         *PackageDef
	tImportDef          *ImportDef
	tImportList         *ImportList
	tExprConstant       *ExprConstant
	tCmdMakeAppUser     *CmdMakeAppUser
	tCmdGenData         *CmdGenData
	tCmdVerify          *CmdVerify
	tCmdListApp         *CmdListApp
	tCmdMakeAppTid      *CmdMakeAppTid
	tStmtJump           *StmtJump
	tExprListArgument   *ExprListArgument
	tStmtIteration      *StmtIteration
	tStmtWhile          *StmtWhile
	tStmtFor            *StmtFor
	tStmtDo             *StmtDo
	tStmtSelection      *StmtSelection
	tInitializerList    *InitializerList
	tUnaryOperator      *UnaryOperator
	tParamDecl          *ParamDecl
	tParamList          *ParamList
	tParamTypeList      *ParamTypeList
	tIdentifierList     *IdentifierList
	tCmdTestRoute       *CmdTestRoute
	tCmdCreateApp       *CmdCreateApp
	tCmdDeleteApp       *CmdDeleteApp
	tCmdDeleteAppTid    *CmdDeleteAppTid
	tCmdRunSQL          *CmdRunSQL
	tCmdDeleteAppUser   *CmdDeleteAppUser
}

const IDENTIFIER = 57346
const CREATE_ID = 57347
const HUB_ID = 57348
const APP_ID = 57349
const NAME_ID = 57350
const EQUAL_ID = 57351
const DROP_ID = 57352
const STYLE_ID = 57353
const LIST_ID = 57354
const APPS_ID = 57355
const RUN_ID = 57356
const FILE_SPEC_ID = 57357
const USE_ID = 57358
const TEST_ID = 57359
const MAKEAPPUSER_ID = 57360
const MAKEAPPTID_ID = 57361
const IN_ID = 57362
const EMAIL_ID = 57363
const UUID_ID = 57364
const COMMENT_SECTION_ID = 57365
const STRING_CONST_ID = 57366
const GENDATA_ID = 57367
const LOOP_ID = 57368
const INTEGER_ID = 57369
const FLOAT_ID = 57370
const POPULATE_ID = 57371
const FROM_ID = 57372
const VERIFY_ID = 57373
const URLID_ID = 57374
const WITH_ID = 57375
const IMPORT_ID = 57376
const PACKAGE_ID = 57377
const FUNC_ID = 57378
const TYPE_STRING_ID = 57379
const TYPE_INT_ID = 57380
const TYPE_FLOAT_ID = 57381
const TYPE_BOOL_ID = 57382
const TYPE_TESTPLAN_ID = 57383
const TRUE_ID = 57384
const FALSE_ID = 57385
const BAD_ID = 57386
const LITERAL_STRING = 57387
const LOAD_ID = 57388
const NEWLINE = 57389
const OR_OP = 57390
const AND_OP = 57391
const EQ_OP = 57392
const NE_OP = 57393
const LE_OP = 57394
const GE_OP = 57395
const LEFT_OP = 57396
const RIGHT_OP = 57397
const INC_OP = 57398
const DEC_OP = 57399
const PTR_OP = 57400
const ADD_ASSIGN = 57401
const SUB_ASSIGN = 57402
const MUL_ASSIGN = 57403
const DIV_ASSIGN = 57404
const MOD_ASSIGN = 57405
const AND_ASSIGN = 57406
const XOR_ASSIGN = 57407
const OR_ASSIGN = 57408
const IF_ID = 57409
const ELSE_ID = 57410
const ENDIF_ID = 57411
const NO_ELSE_ID = 57412
const RETURN_ID = 57413
const EXPR_ID = 57414
const WHILE_ID = 57415
const FOR_ID = 57416
const DO_ID = 57417
const STEP_ID = 57418
const STIMULUS_ID = 57419
const RESPONSE_ID = 57420
const END_STEP_ID = 57421
const FUNC_CALL_ID = 57422
const OBJECT_FIELD_REF = 57423
const TESTROUTE_ID = 57424
const DELETE_ID = 57425
const TID_ID = 57426
const RUNSQL_ID = 57427
const DELETE_APPUSER_ID = 57428
const UMINUS = 57429

var yyToknames = []string{
	"IDENTIFIER",
	"CREATE_ID",
	"HUB_ID",
	"APP_ID",
	"NAME_ID",
	"EQUAL_ID",
	"DROP_ID",
	"STYLE_ID",
	"LIST_ID",
	"APPS_ID",
	"RUN_ID",
	"FILE_SPEC_ID",
	"USE_ID",
	"TEST_ID",
	"MAKEAPPUSER_ID",
	"MAKEAPPTID_ID",
	"IN_ID",
	"EMAIL_ID",
	"UUID_ID",
	"COMMENT_SECTION_ID",
	"STRING_CONST_ID",
	"GENDATA_ID",
	"LOOP_ID",
	"INTEGER_ID",
	"FLOAT_ID",
	"POPULATE_ID",
	"FROM_ID",
	"VERIFY_ID",
	"URLID_ID",
	"WITH_ID",
	"IMPORT_ID",
	"PACKAGE_ID",
	"FUNC_ID",
	"TYPE_STRING_ID",
	"TYPE_INT_ID",
	"TYPE_FLOAT_ID",
	"TYPE_BOOL_ID",
	"TYPE_TESTPLAN_ID",
	"TRUE_ID",
	"FALSE_ID",
	"BAD_ID",
	"LITERAL_STRING",
	"LOAD_ID",
	"NEWLINE",
	"OR_OP",
	"AND_OP",
	"EQ_OP",
	"NE_OP",
	"LE_OP",
	"GE_OP",
	"LEFT_OP",
	"RIGHT_OP",
	"INC_OP",
	"DEC_OP",
	"PTR_OP",
	"ADD_ASSIGN",
	"SUB_ASSIGN",
	"MUL_ASSIGN",
	"DIV_ASSIGN",
	"MOD_ASSIGN",
	"AND_ASSIGN",
	"XOR_ASSIGN",
	"OR_ASSIGN",
	"IF_ID",
	"ELSE_ID",
	"ENDIF_ID",
	"NO_ELSE_ID",
	"RETURN_ID",
	"EXPR_ID",
	"WHILE_ID",
	"FOR_ID",
	"DO_ID",
	"STEP_ID",
	"STIMULUS_ID",
	"RESPONSE_ID",
	"END_STEP_ID",
	"FUNC_CALL_ID",
	"OBJECT_FIELD_REF",
	"TESTROUTE_ID",
	"DELETE_ID",
	"TID_ID",
	"RUNSQL_ID",
	"DELETE_APPUSER_ID",
	"'.'",
	"'|'",
	"'&'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"UMINUS",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line grammar.y:1755

/*  start  of  programs  */

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 167
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1467

var yyAct = []int{

	43, 35, 246, 52, 128, 36, 7, 33, 105, 22,
	50, 8, 88, 118, 104, 24, 101, 70, 103, 172,
	198, 29, 100, 178, 179, 13, 106, 131, 99, 307,
	201, 112, 98, 89, 102, 41, 23, 217, 197, 122,
	119, 196, 131, 285, 41, 130, 134, 289, 129, 288,
	96, 97, 255, 131, 216, 136, 95, 93, 193, 189,
	156, 157, 112, 319, 190, 318, 131, 158, 131, 206,
	120, 91, 92, 13, 94, 173, 317, 176, 177, 131,
	161, 162, 163, 164, 165, 56, 57, 316, 205, 303,
	131, 171, 9, 220, 144, 145, 141, 142, 143, 146,
	147, 148, 300, 10, 6, 131, 16, 15, 17, 18,
	19, 131, 302, 301, 154, 155, 153, 14, 72, 74,
	75, 73, 131, 131, 126, 90, 194, 299, 110, 296,
	131, 298, 122, 200, 131, 223, 202, 76, 77, 140,
	219, 297, 199, 204, 131, 152, 293, 222, 278, 131,
	207, 209, 291, 112, 151, 131, 150, 284, 283, 218,
	131, 131, 212, 112, 28, 14, 27, 208, 275, 224,
	225, 131, 215, 131, 112, 13, 112, 112, 112, 112,
	112, 112, 112, 112, 112, 112, 112, 112, 112, 112,
	112, 230, 231, 240, 241, 242, 229, 238, 239, 236,
	237, 228, 245, 226, 129, 119, 253, 244, 254, 247,
	252, 232, 233, 234, 235, 131, 267, 258, 266, 131,
	131, 131, 263, 264, 265, 256, 250, 249, 257, 131,
	131, 227, 191, 188, 131, 192, 25, 117, 262, 131,
	184, 185, 186, 261, 170, 260, 259, 251, 221, 169,
	168, 269, 272, 274, 137, 135, 132, 270, 271, 247,
	279, 280, 281, 282, 182, 183, 321, 14, 277, 310,
	16, 15, 17, 18, 19, 315, 276, 294, 295, 203,
	290, 292, 180, 181, 159, 187, 108, 26, 174, 175,
	149, 21, 305, 51, 34, 287, 309, 311, 306, 286,
	247, 308, 312, 313, 314, 89, 84, 16, 15, 17,
	18, 19, 31, 78, 167, 160, 166, 320, 243, 81,
	82, 111, 96, 97, 214, 213, 80, 125, 95, 93,
	115, 138, 79, 2, 139, 20, 69, 68, 16, 15,
	17, 18, 19, 91, 92, 124, 94, 67, 66, 65,
	64, 114, 113, 116, 58, 268, 38, 56, 57, 37,
	40, 211, 39, 63, 59, 60, 61, 62, 44, 109,
	107, 5, 45, 3, 46, 48, 47, 49, 1, 4,
	11, 127, 12, 83, 85, 54, 86, 87, 32, 55,
	72, 74, 75, 73, 53, 71, 0, 90, 89, 84,
	0, 0, 23, 123, 42, 0, 78, 0, 0, 76,
	77, 0, 81, 82, 0, 96, 97, 0, 0, 80,
	0, 95, 93, 0, 0, 79, 0, 0, 0, 0,
	0, 16, 15, 17, 18, 19, 91, 92, 0, 94,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	56, 57, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 44, 0, 0, 0, 45, 0, 46, 48, 47,
	49, 0, 0, 0, 0, 0, 83, 85, 0, 86,
	87, 0, 0, 72, 74, 75, 73, 0, 0, 0,
	90, 89, 84, 0, 0, 23, 30, 42, 0, 78,
	0, 0, 76, 77, 0, 81, 82, 0, 96, 97,
	0, 0, 80, 0, 95, 93, 0, 0, 79, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 91,
	92, 0, 94, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 56, 57, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 44, 0, 0, 0, 45, 0,
	46, 48, 47, 49, 0, 0, 0, 0, 0, 83,
	85, 0, 86, 87, 0, 0, 72, 74, 75, 73,
	0, 0, 0, 90, 89, 84, 0, 0, 23, 195,
	42, 0, 78, 0, 0, 76, 77, 0, 81, 82,
	0, 96, 97, 0, 0, 80, 0, 95, 93, 0,
	0, 79, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 91, 92, 0, 94, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 56, 57, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 44, 0, 0,
	0, 45, 0, 46, 48, 47, 49, 0, 0, 0,
	0, 0, 83, 85, 0, 86, 87, 0, 0, 72,
	74, 75, 73, 0, 0, 0, 90, 89, 84, 0,
	0, 23, 121, 42, 0, 78, 0, 0, 76, 77,
	0, 81, 82, 0, 96, 97, 0, 0, 80, 0,
	95, 93, 0, 0, 79, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 91, 92, 0, 94, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 56,
	57, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	44, 0, 0, 0, 45, 0, 46, 48, 47, 49,
	0, 0, 0, 0, 0, 83, 85, 0, 86, 87,
	0, 0, 72, 74, 75, 73, 0, 0, 0, 90,
	89, 84, 0, 0, 23, 0, 42, 0, 78, 0,
	0, 76, 77, 0, 81, 82, 0, 96, 97, 0,
	0, 80, 0, 95, 93, 0, 0, 79, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 91, 92,
	0, 94, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 56, 57, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 83, 85,
	0, 86, 87, 0, 0, 72, 74, 75, 73, 0,
	0, 0, 90, 89, 84, 0, 0, 248, 304, 0,
	0, 78, 0, 0, 76, 77, 0, 81, 82, 0,
	96, 97, 0, 0, 80, 0, 95, 93, 0, 0,
	79, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 91, 92, 0, 94, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 56, 57, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 83, 85, 0, 86, 87, 0, 0, 72, 74,
	75, 73, 0, 0, 0, 90, 273, 89, 84, 0,
	0, 0, 0, 0, 0, 78, 0, 76, 77, 0,
	0, 81, 82, 0, 96, 97, 0, 0, 80, 0,
	95, 93, 0, 0, 79, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 91, 92, 0, 94, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 56,
	57, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 83, 85, 0, 86, 87,
	0, 0, 72, 74, 75, 73, 0, 0, 0, 90,
	89, 84, 0, 0, 248, 0, 0, 0, 78, 0,
	0, 76, 77, 0, 81, 82, 0, 96, 97, 0,
	0, 80, 0, 95, 93, 0, 0, 79, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 91, 92,
	0, 94, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 56, 57, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 83, 85,
	0, 86, 87, 0, 0, 72, 74, 75, 73, 0,
	0, 0, 90, 89, 84, 0, 0, 0, 0, 42,
	0, 78, 0, 0, 76, 77, 0, 81, 82, 0,
	96, 97, 0, 0, 80, 0, 95, 93, 0, 0,
	79, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 91, 92, 0, 94, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 56, 57, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 83, 85, 0, 86, 87, 0, 0, 72, 74,
	75, 73, 0, 0, 0, 90, 210, 89, 84, 0,
	0, 0, 0, 0, 0, 78, 0, 76, 77, 0,
	0, 81, 82, 0, 96, 97, 0, 0, 80, 0,
	95, 93, 0, 0, 79, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 91, 92, 0, 94, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 56,
	57, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 83, 85, 0, 86, 87,
	0, 0, 72, 74, 75, 73, 0, 0, 0, 90,
	89, 84, 0, 0, 0, 0, 133, 0, 78, 0,
	0, 76, 77, 0, 81, 82, 0, 96, 97, 0,
	0, 80, 0, 95, 93, 0, 0, 79, 0, 0,
	0, 89, 0, 0, 0, 0, 0, 0, 91, 92,
	0, 94, 0, 0, 0, 0, 0, 0, 96, 97,
	0, 0, 56, 57, 95, 93, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 91,
	92, 0, 94, 0, 0, 0, 0, 0, 83, 85,
	0, 86, 87, 56, 57, 72, 74, 75, 73, 0,
	0, 0, 90, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 76, 77, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 72, 74, 75, 73,
	0, 0, 0, 90, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 76, 77,
}
var yyPact = []int{

	69, 69, -1000, -1000, -1000, -1000, 246, -65, 171, -1000,
	191, 68, -1000, -1000, 171, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 394, -65, -1000, 241, 29, 233, -27,
	-1000, 580, 301, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 21, -1000, -58, 160, 1233, 159, 673, 158, 327,
	-1000, -1000, 35, -1000, 242, 58, 1357, 1357, 1357, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	235, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 302, 1326,
	1326, 1326, 1326, 1326, 309, 307, 154, 153, 156, -1000,
	1326, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -86, -14,
	238, -29, 228, 174, 148, -1000, -1000, 188, -1000, -40,
	-1000, -1000, -1000, -33, 135, -1000, -42, -1000, -1000, 171,
	-1000, -1000, -1000, -1000, 487, -1000, -1000, -62, -1000, -84,
	-1000, 1326, 1326, -1000, -73, 1326, 206, 1046, -8, 1326,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 1357,
	1326, 1139, 321, 320, -1000, -1000, -1000, -1000, -1000, 1357,
	-1000, 22, 11, 139, 120, 73, 152, 51, 1326, 1326,
	1357, 134, 1357, 1357, 1357, 1357, 1357, 1357, 1357, 1357,
	1357, 1357, 1357, 1357, 1357, 1357, 1357, -1000, -1000, -1000,
	-1000, -1000, 314, 270, -1000, -1000, -1000, 171, 953, -1000,
	130, -1000, 129, 151, 1046, 1326, -65, -1000, 235, -47,
	-1000, 128, -1000, -1000, -1000, 156, 1326, 150, 149, 147,
	142, 1326, 1326, 1326, 121, 119, -86, -1000, -14, 238,
	-29, -29, 228, 228, 228, 228, 174, 174, 148, 148,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 953, 673,
	673, 1326, 859, 71, 198, -1000, -1000, 1326, 115, 1326,
	1326, 1326, 1326, 61, 60, 23, 279, 275, -53, -1000,
	212, -1000, 55, 673, 49, 200, -65, -1000, 33, 44,
	34, 30, 5, -1000, -1000, 17, 16, -7, -1000, 766,
	673, -74, -1000, 673, -65, 190, 1326, -1000, -1000, -1000,
	-1000, 1326, 1326, 1326, -1000, -1000, -1000, -1000, -1000, 197,
	-1000, -10, -21, -32, -34, -65, -1000, -1000, -1000, -1000,
	187, -1000,
}
var yyPgo = []int{

	0, 18, 14, 8, 3, 395, 394, 389, 10, 0,
	5, 7, 312, 1, 388, 294, 34, 16, 22, 28,
	32, 12, 17, 385, 293, 11, 382, 381, 4, 6,
	2, 380, 379, 333, 378, 373, 371, 370, 369, 367,
	366, 365, 364, 363, 362, 361, 360, 359, 356, 355,
	354, 13, 353, 352, 351, 350, 349, 348, 347, 337,
	336, 334,
}
var yyR1 = []int{

	0, 34, 34, 33, 33, 33, 35, 36, 36, 37,
	37, 32, 32, 32, 29, 31, 31, 31, 31, 31,
	31, 31, 38, 54, 54, 53, 52, 52, 51, 13,
	13, 13, 13, 12, 12, 11, 11, 11, 11, 11,
	11, 44, 44, 47, 48, 46, 46, 46, 46, 46,
	46, 14, 14, 15, 15, 25, 26, 26, 26, 26,
	26, 27, 27, 28, 28, 30, 30, 30, 49, 49,
	10, 10, 9, 9, 8, 8, 8, 61, 61, 61,
	61, 61, 61, 61, 61, 61, 24, 23, 23, 22,
	22, 21, 21, 20, 20, 19, 19, 18, 18, 18,
	17, 17, 17, 17, 17, 16, 16, 16, 1, 1,
	1, 2, 2, 2, 2, 3, 50, 50, 50, 50,
	50, 50, 4, 4, 4, 4, 5, 5, 5, 5,
	5, 5, 5, 5, 5, 7, 7, 7, 7, 7,
	7, 7, 7, 45, 45, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 39, 40, 41, 42,
	43, 55, 56, 57, 58, 59, 60,
}
var yyR2 = []int{

	0, 1, 2, 1, 1, 1, 2, 2, 4, 1,
	2, 2, 3, 1, 1, 1, 3, 4, 3, 4,
	4, 3, 1, 1, 3, 1, 1, 3, 2, 2,
	3, 3, 4, 1, 2, 1, 1, 1, 1, 1,
	1, 2, 3, 5, 7, 5, 7, 6, 7, 10,
	7, 1, 2, 2, 3, 1, 1, 1, 1, 1,
	1, 1, 3, 1, 3, 1, 3, 4, 1, 3,
	1, 2, 1, 3, 1, 3, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 3, 1,
	3, 1, 3, 1, 3, 1, 3, 1, 3, 3,
	1, 3, 3, 3, 3, 1, 3, 3, 1, 3,
	3, 1, 3, 3, 3, 1, 1, 1, 1, 1,
	1, 1, 1, 2, 2, 2, 1, 3, 1, 1,
	1, 1, 1, 1, 1, 1, 4, 3, 4, 3,
	3, 2, 2, 1, 3, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 6, 6, 8, 2,
	6, 6, 5, 5, 8, 8, 8,
}
var yyChk = []int{

	-1000, -34, -33, -35, -32, -36, 35, -29, -25, 23,
	34, -31, -26, 4, 96, 38, 37, 39, 40, 41,
	-33, 45, -13, 101, -29, 45, 96, 98, 96, -29,
	102, -12, -14, -11, -15, -13, -10, -47, -48, -44,
	-46, -25, 103, -9, 67, 71, 73, 75, 74, 76,
	-8, -24, -4, -6, -23, -7, 56, 57, -50, -42,
	-41, -40, -39, -43, -55, -56, -57, -58, -59, -60,
	-22, -5, 89, 92, 90, 91, 108, 109, 12, 31,
	25, 18, 19, 82, 5, 83, 85, 86, -21, 4,
	96, 42, 43, 28, 45, 27, 21, 22, -20, -19,
	-18, -17, -16, -1, -2, -3, -13, -37, 45, -38,
	99, -24, -4, -53, -54, 97, -52, 4, -51, -25,
	97, 102, -11, 102, -12, -15, 103, -27, -28, -29,
	103, 100, 96, 103, -9, 96, -11, 96, 4, -61,
	104, 61, 62, 63, 59, 60, 64, 65, 66, 48,
	98, 96, 87, 58, 56, 57, -4, -4, -3, 49,
	13, -9, -9, -9, -9, -9, 7, 7, 96, 96,
	88, -9, 105, 89, 50, 51, 106, 107, 52, 53,
	54, 55, 90, 91, 92, 93, 94, 97, 45, 99,
	97, 97, 100, 100, -29, 102, 103, 100, 104, -8,
	-9, 103, -9, 73, -10, 96, 77, -8, -22, -9,
	97, -45, -8, 4, 4, -21, 32, 26, 20, 20,
	20, 96, 96, 84, -9, -9, -20, 97, -19, -18,
	-17, -17, -16, -16, -16, -16, -1, -1, -2, -2,
	-3, -3, -3, 4, -51, -28, -30, -8, 101, 97,
	97, 96, -10, -9, -13, 99, 97, 100, -9, 96,
	96, 96, 96, -9, -9, -9, 97, 97, -49, -30,
	-11, -11, -9, 97, -9, 97, 78, -8, 33, -9,
	-9, -9, -9, 97, 97, 20, 20, 20, 102, 100,
	68, 97, -11, 97, 77, -13, 96, 97, 97, 97,
	97, 96, 96, 96, 102, -30, -11, 103, -11, -13,
	79, -9, -9, -9, -9, 78, 97, 97, 97, 97,
	-13, 79,
}
var yyDef = []int{

	0, -2, 1, 3, 4, 5, 0, 0, 0, 13,
	0, 14, 55, 15, 0, 56, 57, 58, 59, 60,
	2, 6, 11, 0, 0, 7, 0, 0, 0, 0,
	29, 0, 0, 33, 51, 35, 36, 37, 38, 39,
	40, 0, 70, 0, 0, 0, 0, 0, 0, 0,
	72, 74, 115, 76, 86, 122, 0, 0, 0, 145,
	146, 147, 148, 149, 150, 151, 152, 153, 154, 155,
	87, 135, 116, 117, 118, 119, 120, 121, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 89, 126,
	0, 128, 129, 130, 131, 132, 133, 134, 91, 93,
	95, 97, 100, 105, 108, 111, 12, 0, 9, 0,
	18, 22, 115, 0, 0, 21, 25, 23, 26, 0,
	16, 30, 34, 31, 0, 52, 53, 0, 61, 63,
	71, 0, 0, 41, 0, 0, 0, 0, 0, 0,
	77, 78, 79, 80, 81, 82, 83, 84, 85, 0,
	0, 0, 0, 0, 141, 142, 123, 124, 125, 0,
	159, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 8, 10, 17,
	19, 20, 0, 0, 28, 32, 54, 0, 0, 73,
	0, 42, 0, 0, 0, 0, 0, 75, 88, 0,
	137, 0, 143, 139, 140, 90, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 92, 127, 94, 96,
	98, 99, 101, 102, 103, 104, 106, 107, 109, 110,
	112, 113, 114, 24, 27, 62, 64, 65, 0, 0,
	0, 0, 0, 0, 0, 136, 138, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 68,
	43, 45, 0, 0, 0, 0, 0, 144, 0, 0,
	0, 0, 0, 162, 163, 0, 0, 0, 66, 0,
	0, 0, 47, 0, 0, 0, 0, 157, 156, 160,
	161, 0, 0, 0, 67, 69, 44, 46, 48, 0,
	50, 0, 0, 0, 0, 0, 158, 164, 165, 166,
	0, 49,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 109, 3, 3, 3, 94, 89, 3,
	96, 97, 92, 90, 100, 91, 87, 93, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 103,
	106, 104, 107, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 98, 3, 99, 105, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 101, 88, 102, 108,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73, 74, 75, 76, 77, 78, 79, 80, 81,
	82, 83, 84, 85, 86, 95,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line grammar.y:281
		{
			mTranslationUnit = new(TranslationUnit)
			mTranslationUnit.DeclExternal = yyS[yypt-0].tDeclExternal
			yyVAL.tTranslationUnit = mTranslationUnit
			rlog2("external_declaration ET=%v \n", yyS[yypt-0].tDeclExternal)
		}
	case 2:
		//line grammar.y:288
		{
			mTranslationUnit = new(TranslationUnit)
			mTranslationUnit.TranslationUnit = yyS[yypt-1].tTranslationUnit
			mTranslationUnit.DeclExternal = yyS[yypt-0].tDeclExternal
			yyVAL.tTranslationUnit = mTranslationUnit
			rlog2("translation_unit external_declaration $1=%v $1.ET=%v $2=%v \n", yyS[yypt-1].tTranslationUnit, yyS[yypt-1].tTranslationUnit.DeclExternal, yyS[yypt-0].tDeclExternal)
		}
	case 3:
		//line grammar.y:299
		{
			mDeclExternal = new(DeclExternal)
			mDeclExternal.PackageDef = yyS[yypt-0].tPackageDef
			yyVAL.tDeclExternal = mDeclExternal
			rlog2("package_definition\n", yyVAL.tDeclExternal, yyS[yypt-0].tPackageDef)
		}
	case 4:
		//line grammar.y:306
		{
			mDeclExternal = new(DeclExternal)
			mDeclExternal.FuncDef = yyS[yypt-0].tFuncDef
			yyVAL.tDeclExternal = mDeclExternal
			rlog2("function_definition\n", yyVAL.tDeclExternal, yyS[yypt-0].tFuncDef)
		}
	case 5:
		//line grammar.y:313
		{
			mDeclExternal = new(DeclExternal)
			mDeclExternal.ImportDef = yyS[yypt-0].tImportDef
			yyVAL.tDeclExternal = mDeclExternal
			rlog2("import_definition\n", yyVAL.tDeclExternal, yyS[yypt-0].tImportDef)
		}
	case 6:
		//line grammar.y:323
		{
			mPackageDef = new(PackageDef)
			mPackageDef.Name = yyS[yypt-0].id
			rlog2("package_definition > PACKAGE_ID LITERAL_STRING\n", mPackageDef, yyS[yypt-0].id)
			yyVAL.tPackageDef = mPackageDef
		}
	case 7:
		//line grammar.y:333
		{
			mImportList = new(ImportList)
			mImportList.List = []string{yyS[yypt-0].id}

			mImportDef = new(ImportDef)
			mImportDef.ImportList = mImportList
			yyVAL.tImportDef = mImportDef

			rlog2("IMPORT_ID LITERAL_STRING>>\n", mImportDef)
		}
	case 8:
		//line grammar.y:344
		{
			mImportDef = new(ImportDef)
			mImportDef.ImportList = yyS[yypt-1].tImportList
			yyVAL.tImportDef = mImportDef
			rlog2("IMPORT_ID '(' import_list ')' >>\n", mImportDef, mImportDef.ImportList)
		}
	case 9:
		//line grammar.y:354
		{
			mImportList = new(ImportList)
			mImportList.List = []string{}
			mImportList.List = append(mImportList.List, yyS[yypt-0].id)
			yyVAL.tImportList = mImportList
			rlog2("IMPORT_ID >> LITERAL_STRING", yyS[yypt-0].id)
		}
	case 10:
		//line grammar.y:362
		{
			yyVAL.tImportList.List = append(yyVAL.tImportList.List, yyS[yypt-0].id)
			rlog2("import_list >> LITERAL_STRING\n", yyS[yypt-0].id)
		}
	case 11:
		//line grammar.y:370
		{
			mFuncDef = new(FuncDef)
			mFuncDef.Declarator = yyS[yypt-1].tDeclarator
			mFuncDef.CmpndStmt = yyS[yypt-0].tCmpndStmt
			yyVAL.tFuncDef = mFuncDef
			rlog2("function_definition >> declarator compound_statement : \n", yyS[yypt-1].tDeclarator)
		}
	case 12:
		//line grammar.y:378
		{
			mFuncDef = new(FuncDef)
			mFuncDef.SpecDecl = yyS[yypt-2].tSpecDecl
			mFuncDef.Declarator = yyS[yypt-1].tDeclarator
			mFuncDef.CmpndStmt = yyS[yypt-0].tCmpndStmt
			yyVAL.tFuncDef = mFuncDef

			rlog2("function_definition >> declaration_specifiers declarator compound_statement : $$=%v CS=%v \n", yyVAL.tFuncDef, yyS[yypt-0].tCmpndStmt)
			//DumpFuncDef(mFuncDef,3)
		}
	case 13:
		//line grammar.y:389
		{

		}
	case 14:
		//line grammar.y:396
		{
			mDeclarator = new(Declarator)
			mDeclarator.DeclDirect = mDeclDirect
			yyVAL.tDeclarator = mDeclarator
			rlog2("declarator >> direct_declarator:\n", mDeclarator, mDeclarator.DeclDirect, mDeclarator.DeclDirect.Identifier)
		}
	case 15:
		//line grammar.y:406
		{
			mDeclDirect = new(DeclDirect)
			mIdentifier = new(Identifier)
			mIdentifier.Name = yyS[yypt-0].id
			mDeclDirect.Identifier = mIdentifier
			rlog2("direct_declarator IDENTIFIER-->\n", mDeclDirect, mIdentifier)
			yyVAL.tDeclDirect = mDeclDirect
		}
	case 16:
		//line grammar.y:415
		{
			//rlog2("direct_declarator '(' declarator ')'\n")
			mDeclDirect = new(DeclDirect)
			mDeclDirect.Declarator = yyS[yypt-1].tDeclarator
			mDeclDirect.Modifier = "()"
			yyVAL.tDeclDirect = mDeclDirect
		}
	case 17:
		//line grammar.y:423
		{
			mDeclDirect = new(DeclDirect)
			mDeclDirect.DeclDirect = yyS[yypt-3].tDeclDirect
			mDeclDirect.ExprConstant = yyS[yypt-1].tExprConstant
			mDeclDirect.Modifier = "[]"
			yyVAL.tDeclDirect = mDeclDirect
			rlog2("direct_declarator '[' constant_expression ']'\n", yyVAL.tDeclDirect)
		}
	case 18:
		//line grammar.y:432
		{
			mDeclDirect = new(DeclDirect)
			mDeclDirect.DeclDirect = yyS[yypt-2].tDeclDirect
			mDeclDirect.Modifier = "[]"
			yyVAL.tDeclDirect = mDeclDirect
			rlog2("direct_declarator '[' ']'")
		}
	case 19:
		//line grammar.y:440
		{
			mDeclDirect = new(DeclDirect)
			mDeclDirect.DeclDirect = yyS[yypt-3].tDeclDirect
			mDeclDirect.ParamTypeList = yyS[yypt-1].tParamTypeList
			mDeclDirect.Modifier = "()"
			yyVAL.tDeclDirect = mDeclDirect
			rlog2("direct_declarator '(' parameter_type_list ')'\n", mPackageDef, mFuncDef, mParamDeclList)
		}
	case 20:
		//line grammar.y:449
		{
			mDeclDirect = new(DeclDirect)
			mDeclDirect.DeclDirect = yyS[yypt-3].tDeclDirect
			mDeclDirect.IdentifierList = yyS[yypt-1].tIdentifierList
			mDeclDirect.Modifier = "()"
			yyVAL.tDeclDirect = mDeclDirect
			rlog2("direct_declarator '(' identifier_list ')'\n")
		}
	case 21:
		//line grammar.y:458
		{
			mDeclDirect = new(DeclDirect)
			mDeclDirect.DeclDirect = yyS[yypt-2].tDeclDirect
			mDeclDirect.Modifier = "()"
			yyVAL.tDeclDirect = mDeclDirect
			rlog2("direct_declarator '(' ')'\n")
		}
	case 22:
		//line grammar.y:469
		{
			mExprConstant = new(ExprConstant)
			mExprConstant.ExprConditional = yyS[yypt-0].tExprConditional
			yyVAL.tExprConstant = mExprConstant
			rlog2("constant_expression:conditional_expression\n", yyS[yypt-0].tExprConditional)
		}
	case 23:
		//line grammar.y:479
		{
			mIdentifier = new(Identifier)
			mIdentifier.Name = yyS[yypt-0].id

			mIdentifierList = new(IdentifierList)
			mIdentifierList.Identifier = mIdentifier
			yyVAL.tIdentifierList = mIdentifierList
		}
	case 24:
		//line grammar.y:488
		{
			mIdentifier = new(Identifier)
			mIdentifier.Name = yyS[yypt-0].id

			mIdentifierList = new(IdentifierList)
			mIdentifierList.Identifier = mIdentifier
			mIdentifierList.IdentifierList = yyS[yypt-2].tIdentifierList
			yyVAL.tIdentifierList = mIdentifierList
		}
	case 25:
		//line grammar.y:501
		{
			mParamTypeList = new(ParamTypeList)
			mParamTypeList.ParamList = yyS[yypt-0].tParamList
			yyVAL.tParamTypeList = mParamTypeList
			rlog2("parameter_type_list>>parameter_list: \n", mParamTypeList)
		}
	case 26:
		//line grammar.y:511
		{
			mParamList = new(ParamList)
			mParamList.ParamDecl = yyS[yypt-0].tParamDecl
			yyVAL.tParamList = mParamList

			rlog2("parameter_list>>parameter_declaration-->\n", yyVAL.tParamList)
		}
	case 27:
		//line grammar.y:519
		{
			mParamList = new(ParamList)
			mParamList.ParamDecl = yyS[yypt-0].tParamDecl
			mParamList.ParamList = yyS[yypt-2].tParamList
			yyVAL.tParamList = mParamList
			rlog2("parameter_list>>parameter_list ',' parameter_declaration -->\n", yyVAL.tParamList)
		}
	case 28:
		//line grammar.y:529
		{
			mParamDecl = new(ParamDecl)
			mParamDecl.SpecDecl = yyS[yypt-1].tSpecDecl
			mParamDecl.Declarator = yyS[yypt-0].tDeclarator
			yyVAL.tParamDecl = mParamDecl
			rlog2("parameter_declaration >> declaration_specifiers declarator:\n", yyVAL.tParamDecl, yyS[yypt-1].tSpecDecl, yyS[yypt-0].tDeclarator)
		}
	case 29:
		//line grammar.y:540
		{
			mCmpndStmt = new(CmpndStmt)
			yyVAL.tCmpndStmt = mCmpndStmt
			rlog2("'{' '}'\n", yyVAL.tCmpndStmt)
		}
	case 30:
		//line grammar.y:546
		{
			mCmpndStmt = new(CmpndStmt)
			mCmpndStmt.StmtList = yyS[yypt-1].tStmtList
			yyVAL.tCmpndStmt = mCmpndStmt
			rlog2("compound_statement >> '{' statement_list '}' :\n", yyVAL.tCmpndStmt)

		}
	case 31:
		//line grammar.y:554
		{
			mCmpndStmt = new(CmpndStmt)
			mCmpndStmt.DeclarationList = yyS[yypt-1].tDeclarationList
			yyVAL.tCmpndStmt = mCmpndStmt
			rlog2("compound_statement >> '{' declaration_list '}' :\n", yyVAL.tCmpndStmt, yyS[yypt-1].tDeclarationList, yyS[yypt-1].tDeclarationList.Declaration)
		}
	case 32:
		//line grammar.y:561
		{
			mCmpndStmt = new(CmpndStmt)
			mCmpndStmt.DeclarationList = yyS[yypt-2].tDeclarationList
			mCmpndStmt.StmtList = yyS[yypt-1].tStmtList
			yyVAL.tCmpndStmt = mCmpndStmt
			rlog2("compound_statement >> '{' declaration_list statement_list '}'\n", yyVAL.tCmpndStmt)
		}
	case 33:
		//line grammar.y:572
		{
			mStmtList = new(StmtList)
			mStmtList.Stmt = yyS[yypt-0].tStmt
			yyVAL.tStmtList = mStmtList
			/*
			   rlog2("statement\n")
			   mStmt = new(Stmt)
			   if mExprStmt != nil {
			       mStmt.ExprStmt = mExprStmt
			       mExprStmt = nil
			   }

			   if mStmtList == nil {
			       mStmtList = new(StmtList)
			       mStmtList.List = []Stmt{}
			   }
			   mStmtList.List  = append(mStmtList.List , *mStmt)
			*/
		}
	case 34:
		//line grammar.y:591
		{
			mStmtList = new(StmtList)
			mStmtList.Stmt = yyS[yypt-0].tStmt
			mStmtList.StmtList = yyS[yypt-1].tStmtList
			yyVAL.tStmtList = mStmtList
		}
	case 35:
		//line grammar.y:602
		{
			mStmt = new(Stmt)
			mStmt.CmpndStmt = yyS[yypt-0].tCmpndStmt
			yyVAL.tStmt = mStmt
			rlog2("statement:compound_statement:\n", mExprStmt, mExprAssign)
		}
	case 36:
		//line grammar.y:609
		{
			mStmt = new(Stmt)
			mStmt.ExprStmt = yyS[yypt-0].tExprStmt
			yyVAL.tStmt = mStmt
		}
	case 37:
		//line grammar.y:615
		{
			mStmt = new(Stmt)
			mStmt.StmtSelection = yyS[yypt-0].tStmtSelection
			yyVAL.tStmt = mStmt
			rlog2("selection_statement\n")
		}
	case 38:
		//line grammar.y:622
		{
			mStmt = new(Stmt)
			mStmt.StmtSelection = yyS[yypt-0].tStmtSelection
			yyVAL.tStmt = mStmt
			rlog2("selection_statement\n")
		}
	case 39:
		//line grammar.y:629
		{
			rlog2("jump_statement\n")
			mStmt = new(Stmt)
			mStmt.StmtJump = yyS[yypt-0].tStmtJump
			yyVAL.tStmt = mStmt
		}
	case 40:
		//line grammar.y:636
		{
			mStmt = new(Stmt)
			mStmt.StmtIteration = yyS[yypt-0].tStmtIteration
			yyVAL.tStmt = mStmt
			//DumpStmtIteration(mStmt.StmtIteration, 3)
		}
	case 41:
		//line grammar.y:646
		{
			mStmtJump = new(StmtJump)
			mStmtJump.Name = "return"
			yyVAL.tStmtJump = mStmtJump
		}
	case 42:
		//line grammar.y:652
		{
			mStmtJump = new(StmtJump)
			mStmtJump.Name = "return"
			mStmtJump.Expr = yyS[yypt-1].tExpr
			yyVAL.tStmtJump = mStmtJump
			rlog2("RETURN_ID expression ';'\n")
		}
	case 43:
		//line grammar.y:663
		{
			mStmtSelection = new(StmtSelection)
			mStmtSelection.Expr = yyS[yypt-2].tExpr
			mStmtSelection.Stmt = yyS[yypt-0].tStmt
			yyVAL.tStmtSelection = mStmtSelection
			rlog(1, "IF_ID '(' expression ')' statement\n")
			//DumpStmtSelection(mStmtSelection,1)
		}
	case 44:
		//line grammar.y:675
		{
			mStmtSelection = new(StmtSelection)
			mStmtSelection.Expr = yyS[yypt-4].tExpr
			mStmtSelection.Stmt = yyS[yypt-2].tStmt
			mStmtSelection.Stmt2 = yyS[yypt-0].tStmt
			yyVAL.tStmtSelection = mStmtSelection
			rlog(1, "IF_ID '(' expression ')' statement ELSE_ID  statement\n")
			//DumpStmtSelection(mStmtSelection,1)
		}
	case 45:
		//line grammar.y:692
		{
			mStmtWhile = new(StmtWhile)
			mStmtWhile.Expr = yyS[yypt-2].tExpr
			mStmtWhile.Stmt = yyS[yypt-0].tStmt

			mStmtIteration = new(StmtIteration)
			mStmtIteration.StmtWhile = mStmtWhile
			yyVAL.tStmtIteration = mStmtIteration
			//DumpStmtWhile(mStmtWhile,2)
		}
	case 46:
		//line grammar.y:703
		{
			mStmtDo = new(StmtDo)
			mStmtDo.Expr = yyS[yypt-2].tExpr
			mStmtDo.Stmt = yyS[yypt-5].tStmt

			mStmtIteration = new(StmtIteration)
			mStmtIteration.StmtDo = mStmtDo
			yyVAL.tStmtIteration = mStmtIteration
		}
	case 47:
		//line grammar.y:713
		{
			mStmtFor = new(StmtFor)
			mStmtFor.ExprStmt1 = yyS[yypt-3].tExprStmt
			mStmtFor.ExprStmt2 = yyS[yypt-2].tExprStmt
			mStmtFor.Stmt = yyS[yypt-0].tStmt

			mStmtIteration = new(StmtIteration)
			mStmtIteration.StmtFor = mStmtFor
			yyVAL.tStmtIteration = mStmtIteration
		}
	case 48:
		//line grammar.y:724
		{
			mStmtFor = new(StmtFor)
			mStmtFor.ExprStmt1 = yyS[yypt-4].tExprStmt
			mStmtFor.ExprStmt2 = yyS[yypt-3].tExprStmt
			mStmtFor.Expr = yyS[yypt-2].tExpr
			mStmtFor.Stmt = yyS[yypt-0].tStmt

			mStmtIteration = new(StmtIteration)
			mStmtIteration.StmtFor = mStmtFor
			yyVAL.tStmtIteration = mStmtIteration
		}
	case 49:
		//line grammar.y:736
		{
			mStmtStep = new(StmtStep)
			mStmtStep.Expr = yyS[yypt-6].tExpr
			mStmtStep.CmpndStmtStimulus = yyS[yypt-3].tCmpndStmt
			mStmtStep.CmpndStmtResponse = yyS[yypt-1].tCmpndStmt
			mIdentifier = new(Identifier)
			mIdentifier.Name = yyS[yypt-8].id
			mStmtStep.Identifier = mIdentifier

			mStmtIteration = new(StmtIteration)
			mStmtIteration.StmtStep = mStmtStep
			yyVAL.tStmtIteration = mStmtIteration
		}
	case 50:
		//line grammar.y:751
		{
			mStmtStep = new(StmtStep)
			mStmtStep.CmpndStmtStimulus = yyS[yypt-3].tCmpndStmt
			mStmtStep.CmpndStmtResponse = yyS[yypt-1].tCmpndStmt
			mIdentifier = new(Identifier)
			mIdentifier.Name = yyS[yypt-5].id
			mStmtStep.Identifier = mIdentifier

			mStmtIteration = new(StmtIteration)
			mStmtIteration.StmtStep = mStmtStep
			yyVAL.tStmtIteration = mStmtIteration
		}
	case 51:
		//line grammar.y:768
		{
			mDeclarationList = new(DeclarationList)
			mDeclarationList.Declaration = yyS[yypt-0].tDeclaration
			yyVAL.tDeclarationList = mDeclarationList
			rlog2("declaration_list>>declaration\n")
		}
	case 52:
		//line grammar.y:775
		{
			mDeclarationList = new(DeclarationList)
			mDeclarationList.Declaration = yyS[yypt-0].tDeclaration
			mDeclarationList.DeclarationList = yyS[yypt-1].tDeclarationList
			yyVAL.tDeclarationList = mDeclarationList
			rlog2("declaration_list>>declaration_list declaration\n")
		}
	case 53:
		//line grammar.y:786
		{
			mDeclaration = new(Declaration)
			mDeclaration.SpecDecl = yyS[yypt-1].tSpecDecl
			yyVAL.tDeclaration = mDeclaration
			rlog2("declaration_specifiers ';'\n")
			//DumpDeclaration(mDeclaration,3)
		}
	case 54:
		//line grammar.y:794
		{
			mDeclaration = new(Declaration)
			mDeclaration.SpecDecl = yyS[yypt-2].tSpecDecl
			mDeclaration.DeclInitList = yyS[yypt-1].tDeclInitList
			yyVAL.tDeclaration = mDeclaration
			rlog2("declaration_specifiers init_declarator_list ';'\n", yyVAL.tDeclaration, yyS[yypt-2].tSpecDecl, yyS[yypt-1].tDeclInitList)
			//DumpDeclaration(mDeclaration,3)
		}
	case 55:
		//line grammar.y:806
		{
			mSpecDecl = new(SpecDecl)
			mSpecDecl.SpecType = yyS[yypt-0].tSpecType
			yyVAL.tSpecDecl = mSpecDecl
			rlog2("type_specifier $1=%v $$=%v $$.SpecType=%v \n", yyS[yypt-0].tSpecType, yyVAL.tSpecDecl, yyVAL.tSpecDecl.SpecType)
		}
	case 56:
		//line grammar.y:816
		{
			mSpecType = new(SpecType)
			mSpecType.Id = TYPE_INT_ID
			mSpecType.Name = "int"
			yyVAL.tSpecType = mSpecType
		}
	case 57:
		//line grammar.y:823
		{
			mSpecType = new(SpecType)
			mSpecType.Id = LITERAL_STRING
			mSpecType.Name = "string"
			yyVAL.tSpecType = mSpecType
		}
	case 58:
		//line grammar.y:830
		{
			mSpecType = new(SpecType)
			mSpecType.Id = TYPE_FLOAT_ID
			mSpecType.Name = "float"
			yyVAL.tSpecType = mSpecType
		}
	case 59:
		//line grammar.y:837
		{
			mSpecType = new(SpecType)
			mSpecType.Id = TYPE_BOOL_ID
			mSpecType.Name = "bool"
			yyVAL.tSpecType = mSpecType
		}
	case 60:
		//line grammar.y:844
		{
			mSpecType = new(SpecType)
			mSpecType.Id = TYPE_TESTPLAN_ID
			mSpecType.Name = "testplan"
			yyVAL.tSpecType = mSpecType
		}
	case 61:
		//line grammar.y:855
		{
			mDeclInitList = new(DeclInitList)
			mDeclInitList.DeclInit = yyS[yypt-0].tDeclInit
			yyVAL.tDeclInitList = mDeclInitList
			rlog2("init_declarator_list>init_declarator\n", mDeclInitList, yyS[yypt-0].tDeclInit)

		}
	case 62:
		//line grammar.y:863
		{
			mDeclInitList = new(DeclInitList)
			mDeclInitList.DeclInit = yyS[yypt-0].tDeclInit
			mDeclInitList.DeclInitList = yyS[yypt-2].tDeclInitList
			yyVAL.tDeclInitList = mDeclInitList
			rlog2("init_declarator_list ',' init_declarator\n", mDeclInitList, yyS[yypt-2].tDeclInitList)
		}
	case 63:
		//line grammar.y:874
		{
			mDeclInit = new(DeclInit)
			mDeclInit.Declarator = yyS[yypt-0].tDeclarator
			yyVAL.tDeclInit = mDeclInit
			rlog2("declarator 2\n")
		}
	case 64:
		//line grammar.y:881
		{
			mDeclInit = new(DeclInit)
			mDeclInit.Declarator = yyS[yypt-2].tDeclarator
			mDeclInit.Initializer = yyS[yypt-0].tInitializer
			rlog2("declarator '=' initializer mDeclInit=%v mDeclarator=%v mInitializer=%v\n", mDeclInit, mDeclarator, mInitializer)
			yyVAL.tDeclInit = mDeclInit
			//DumpDeclInit(mDeclInit,3)
			//DumpInitializer($3, 3)
			//DumpDeclarator($1, 3)
		}
	case 65:
		//line grammar.y:895
		{
			mInitializer = new(Initializer)
			mInitializer.ExprAssign = mExprAssign
			mInitializer.ExprUnary = mExprUnary
			mInitializer.ExprCmd = mExprCmd
			rlog2("assignment_expression: mInitializer=%v mExprUnary=%v mExprAssign=%v mExprCmd=%v \n", mInitializer, mExprUnary, mExprAssign, mExprCmd)
			//DumpExprAssign(mExprAssign, 4)
			yyVAL.tInitializer = mInitializer
		}
	case 66:
		//line grammar.y:905
		{
			mInitializer = new(Initializer)
			mInitializer.InitializerList = yyS[yypt-1].tInitializerList
			yyVAL.tInitializer = mInitializer
			rlog2("'{' initializer_list '}'\n")
		}
	case 67:
		//line grammar.y:912
		{
			mInitializer = new(Initializer)
			mInitializer.InitializerList = yyS[yypt-2].tInitializerList
			yyVAL.tInitializer = mInitializer
			rlog2("'{' initializer_list ',' '}'\n")
		}
	case 68:
		//line grammar.y:922
		{
			mInitializerList = new(InitializerList)
			mInitializerList.Initializer = yyS[yypt-0].tInitializer
			yyVAL.tInitializerList = mInitializerList
			rlog2("initializer\n")
		}
	case 69:
		//line grammar.y:929
		{
			mInitializerList = new(InitializerList)
			mInitializerList.Initializer = yyS[yypt-0].tInitializer
			mInitializerList.InitializerList = yyS[yypt-2].tInitializerList
			yyVAL.tInitializerList = mInitializerList
			rlog2("initializer_list ',' initializer\n")
		}
	case 70:
		//line grammar.y:940
		{
			mExprStmt = new(ExprStmt)
			yyVAL.tExprStmt = mExprStmt
		}
	case 71:
		//line grammar.y:945
		{
			mExprStmt = new(ExprStmt)
			mExprStmt.Expr = yyS[yypt-1].tExpr
			yyVAL.tExprStmt = mExprStmt
			//rlog2("expression\n")
		}
	case 72:
		//line grammar.y:956
		{
			mExpr = new(Expr)
			mExpr.ExprAssign = yyS[yypt-0].tExprAssign
			yyVAL.tExpr = mExpr
		}
	case 73:
		//line grammar.y:962
		{
			mExpr = new(Expr)
			mExpr.ExprAssign = yyS[yypt-0].tExprAssign
			mExpr.Expr = yyS[yypt-2].tExpr
			yyVAL.tExpr = mExpr
			//rlog2("expression ',' assignment_expression\n")
		}
	case 74:
		//line grammar.y:973
		{
			mExprAssign = new(ExprAssign)
			mExprAssign.ExprConditional = yyS[yypt-0].tExprConditional
			yyVAL.tExprAssign = mExprAssign
			rlog2("conditional_expression mExprAssign=%v mExprCmd=%v mExprPostfix=%v mExprUnary=%v \n", mExprAssign, mExprCmd, mExprPostfix, mExprUnary)
			//DumpExprConditional($1, 3)
		}
	case 75:
		//line grammar.y:981
		{
			mExprAssign = new(ExprAssign)
			mExprAssign.ExprUnary = yyS[yypt-2].tExprUnary
			mExprAssign.ExprAssign = yyS[yypt-0].tExprAssign
			yyVAL.tExprAssign = mExprAssign
			rlog2("unary_expression assignment_operator assignment_expression >> mExprAssign=%v mExprCmd=%v mExprPostfix=%v mExprUnary=%v \n", mExprAssign, mExprCmd, mExprPostfix, mExprUnary)
		}
	case 76:
		//line grammar.y:989
		{
			mExprAssign = new(ExprAssign)
			mExprAssign.ExprCmd = yyS[yypt-0].tExprCmd
			//rlog2("command_expression 2 mExprCmd=%v mExprCmd.CmdListApp=%v \n", mExprCmd, mExprCmd.CmdListApp)
			yyVAL.tExprAssign = mExprAssign
		}
	case 86:
		//line grammar.y:1011
		{
			mExprConditional = new(ExprConditional)
			mExprConditional.ExprLogicOr = yyS[yypt-0].tExprLogicOr
			yyVAL.tExprConditional = mExprConditional
			rlog2("logical_or_expression")
		}
	case 87:
		//line grammar.y:1021
		{
			mExprLogicOr = new(ExprLogicOr)
			mExprLogicOr.ExprLogicAnd = yyS[yypt-0].tExprLogicAnd
			yyVAL.tExprLogicOr = mExprLogicOr
			rlog2("logical_and_expression\n")
		}
	case 88:
		//line grammar.y:1028
		{
			mExprLogicOr = new(ExprLogicOr)
			mExprLogicOr.ExprLogicAnd = yyS[yypt-0].tExprLogicAnd
			mExprLogicOr.ExprLogicOr = yyS[yypt-2].tExprLogicOr
			yyVAL.tExprLogicOr = mExprLogicOr
			rlog2("logical_or_expression OR_OP logical_and_expression\n")
		}
	case 89:
		//line grammar.y:1039
		{
			mExprLogicAnd = new(ExprLogicAnd)
			mExprLogicAnd.ExprInclOr = yyS[yypt-0].tExprInclOr
			yyVAL.tExprLogicAnd = mExprLogicAnd
			rlog2("inclusive_or_expression")
		}
	case 90:
		//line grammar.y:1046
		{
			mExprLogicAnd = new(ExprLogicAnd)
			mExprLogicAnd.ExprInclOr = yyS[yypt-0].tExprInclOr
			mExprLogicAnd.ExprLogicAnd = yyS[yypt-2].tExprLogicAnd
			yyVAL.tExprLogicAnd = mExprLogicAnd
			rlog2("logical_and_expression AND_OP inclusive_or_expression")
		}
	case 91:
		//line grammar.y:1057
		{
			mExprInclOr = new(ExprInclOr)
			mExprInclOr.ExprExclOr = yyS[yypt-0].tExprExclOr
			yyVAL.tExprInclOr = mExprInclOr
			rlog2("exclusive_or_expression\n")
		}
	case 92:
		//line grammar.y:1064
		{
			mExprInclOr = new(ExprInclOr)
			mExprInclOr.ExprExclOr = yyS[yypt-0].tExprExclOr
			mExprInclOr.ExprInclOr = yyS[yypt-2].tExprInclOr
			yyVAL.tExprInclOr = mExprInclOr
			rlog2("exclusive_or_expression\n")
		}
	case 93:
		//line grammar.y:1075
		{
			mExprExclOr = new(ExprExclOr)
			mExprExclOr.ExprAnd = yyS[yypt-0].tExprAnd
			yyVAL.tExprExclOr = mExprExclOr
			rlog2("and_expression\n")
		}
	case 94:
		//line grammar.y:1082
		{
			mExprExclOr = new(ExprExclOr)
			mExprExclOr.ExprAnd = yyS[yypt-0].tExprAnd
			mExprExclOr.ExprExclOr = yyS[yypt-2].tExprExclOr
			yyVAL.tExprExclOr = mExprExclOr
			rlog2("exclusive_or_expression '^' and_expression\n")
		}
	case 95:
		//line grammar.y:1093
		{
			mExprAnd = new(ExprAnd)
			mExprAnd.ExprEquality = yyS[yypt-0].tExprEquality
			yyVAL.tExprAnd = mExprAnd
			rlog2("equality_expression\n")
		}
	case 96:
		//line grammar.y:1100
		{
			mExprAnd = new(ExprAnd)
			mExprAnd.ExprEquality = yyS[yypt-0].tExprEquality
			mExprAnd.ExprAnd = yyS[yypt-2].tExprAnd
			yyVAL.tExprAnd = mExprAnd
			rlog2("and_expression '&' equality_expression\n")
		}
	case 97:
		//line grammar.y:1112
		{
			mExprEquality = new(ExprEquality)
			mExprEquality.ExprRelational = yyS[yypt-0].tExprRelational
			yyVAL.tExprEquality = mExprEquality
			rlog2("relational_expression\n")
		}
	case 98:
		//line grammar.y:1119
		{
			mExprEquality = new(ExprEquality)
			mExprEquality.ExprRelational = yyS[yypt-0].tExprRelational
			mExprEquality.ExprEquality = yyS[yypt-2].tExprEquality
			mExprEquality.Sign = "EQ"
			yyVAL.tExprEquality = mExprEquality
			rlog2("equality_expression EQ_OP relational_expression\n")
		}
	case 99:
		//line grammar.y:1128
		{
			mExprEquality = new(ExprEquality)
			mExprEquality.ExprRelational = yyS[yypt-0].tExprRelational
			mExprEquality.ExprEquality = yyS[yypt-2].tExprEquality
			mExprEquality.Sign = "NE"
			yyVAL.tExprEquality = mExprEquality
			rlog2("equality_expression NE_OP relational_expression\n")
		}
	case 100:
		//line grammar.y:1140
		{
			mExprRelational = new(ExprRelational)
			mExprRelational.ExprShift = yyS[yypt-0].tExprShift
			yyVAL.tExprRelational = mExprRelational
		}
	case 101:
		//line grammar.y:1146
		{
			mExprRelational = new(ExprRelational)
			mExprRelational.ExprShift = yyS[yypt-0].tExprShift
			mExprRelational.ExprRelational = yyS[yypt-2].tExprRelational
			mExprRelational.Sign = "<"
			yyVAL.tExprRelational = mExprRelational
			rlog2("relational_expression '<' shift_expression\n")
		}
	case 102:
		//line grammar.y:1155
		{
			mExprRelational = new(ExprRelational)
			mExprRelational.ExprShift = yyS[yypt-0].tExprShift
			mExprRelational.ExprRelational = yyS[yypt-2].tExprRelational
			mExprRelational.Sign = ">"
			yyVAL.tExprRelational = mExprRelational
			rlog2("relational_expression '>' shift_expression\n")
		}
	case 103:
		//line grammar.y:1164
		{
			mExprRelational = new(ExprRelational)
			mExprRelational.ExprShift = yyS[yypt-0].tExprShift
			mExprRelational.ExprRelational = yyS[yypt-2].tExprRelational
			mExprRelational.Sign = "LE"
			yyVAL.tExprRelational = mExprRelational
			rlog2("relational_expression LE_OP shift_expression\n")
		}
	case 104:
		//line grammar.y:1173
		{
			mExprRelational = new(ExprRelational)
			mExprRelational.ExprShift = yyS[yypt-0].tExprShift
			mExprRelational.ExprRelational = yyS[yypt-2].tExprRelational
			mExprRelational.Sign = "GE"
			yyVAL.tExprRelational = mExprRelational
			rlog2("relational_expression GE_OP shift_expression\n")
		}
	case 105:
		//line grammar.y:1185
		{
			mExprShift = new(ExprShift)
			mExprShift.ExprAdditive = yyS[yypt-0].tExprAdditive
			yyVAL.tExprShift = mExprShift
			rlog2("additive_expression\n")
		}
	case 106:
		//line grammar.y:1192
		{
			mExprShift = new(ExprShift)
			mExprShift.ExprAdditive = yyS[yypt-0].tExprAdditive
			mExprShift.ExprShift = yyS[yypt-2].tExprShift
			yyVAL.tExprShift = mExprShift
			rlog2("shift_expression LEFT_OP additive_expression\n")
		}
	case 107:
		//line grammar.y:1200
		{
			mExprShift = new(ExprShift)
			mExprShift.ExprAdditive = yyS[yypt-0].tExprAdditive
			mExprShift.ExprShift = yyS[yypt-2].tExprShift
			yyVAL.tExprShift = mExprShift
			rlog2("shift_expression RIGHT_OP additive_expression\n")
		}
	case 108:
		//line grammar.y:1211
		{
			mExprAdditive = new(ExprAdditive)
			mExprAdditive.ExprMultiplicative = mExprMultiplicative
			rlog2("additive_expression >> multiplicative_expression %v %v \n", mExprAdditive, mExprMultiplicative)
			yyVAL.tExprAdditive = mExprAdditive
		}
	case 109:
		//line grammar.y:1218
		{
			mExprAdditive = new(ExprAdditive)
			mExprAdditive.ExprAdditive = yyS[yypt-2].tExprAdditive
			mExprAdditive.ExprMultiplicative = yyS[yypt-0].tExprMultiplicative
			mExprAdditive.Sign = "+"
			yyVAL.tExprAdditive = mExprAdditive
			rlog2("additive_expression '+' multiplicative_expression\n")
		}
	case 110:
		//line grammar.y:1227
		{
			mExprAdditive = new(ExprAdditive)
			mExprAdditive.ExprAdditive = yyS[yypt-2].tExprAdditive
			mExprAdditive.ExprMultiplicative = yyS[yypt-0].tExprMultiplicative
			mExprAdditive.Sign = "-"
			yyVAL.tExprAdditive = mExprAdditive
			rlog2("additive_expression '-' multiplicative_expression\n")
		}
	case 111:
		//line grammar.y:1239
		{
			mExprMultiplicative = new(ExprMultiplicative)
			mExprMultiplicative.ExprCast = yyS[yypt-0].tExprCast
			rlog2("multiplicative_expression >> cast_expression mExprMultiplicative=%v mExprCast=%v \n", mExprMultiplicative, mExprCast)
			yyVAL.tExprMultiplicative = mExprMultiplicative
		}
	case 112:
		//line grammar.y:1246
		{
			mExprMultiplicative = new(ExprMultiplicative)
			mExprMultiplicative.ExprCast = yyS[yypt-0].tExprCast
			mExprMultiplicative.ExprMultiplicative = yyS[yypt-2].tExprMultiplicative
			mExprMultiplicative.Sign = "*"
			yyVAL.tExprMultiplicative = mExprMultiplicative
			rlog2("multiplicative_expression '*' cast_expression\n")
		}
	case 113:
		//line grammar.y:1255
		{
			mExprMultiplicative = new(ExprMultiplicative)
			mExprMultiplicative.ExprCast = yyS[yypt-0].tExprCast
			mExprMultiplicative.ExprMultiplicative = yyS[yypt-2].tExprMultiplicative
			mExprMultiplicative.Sign = "/"
			yyVAL.tExprMultiplicative = mExprMultiplicative
			rlog2("multiplicative_expression '/' cast_expression\n")
		}
	case 114:
		//line grammar.y:1264
		{
			mExprMultiplicative = new(ExprMultiplicative)
			mExprMultiplicative.ExprCast = yyS[yypt-0].tExprCast
			mExprMultiplicative.ExprMultiplicative = yyS[yypt-2].tExprMultiplicative
			mExprMultiplicative.Sign = "%"
			yyVAL.tExprMultiplicative = mExprMultiplicative
			rlog2("multiplicative_expression '%' cast_expression\n")
		}
	case 115:
		//line grammar.y:1276
		{
			mExprCast = new(ExprCast)
			mExprCast.ExprUnary = yyS[yypt-0].tExprUnary
			rlog2("cast_expression >> unary_expression mExprCast=%v mExprUnary=%v \n", mExprCast, mExprUnary)
			yyVAL.tExprCast = mExprCast
		}
	case 116:
		//line grammar.y:1286
		{
			mUnaryOperator = new(UnaryOperator)
			mUnaryOperator.Sign = "&"
			yyVAL.tUnaryOperator = mUnaryOperator
		}
	case 117:
		//line grammar.y:1292
		{
			mUnaryOperator = new(UnaryOperator)
			mUnaryOperator.Sign = "*"
			yyVAL.tUnaryOperator = mUnaryOperator
		}
	case 118:
		//line grammar.y:1298
		{
			mUnaryOperator = new(UnaryOperator)
			mUnaryOperator.Sign = "+"
			yyVAL.tUnaryOperator = mUnaryOperator
		}
	case 119:
		//line grammar.y:1304
		{
			mUnaryOperator = new(UnaryOperator)
			mUnaryOperator.Sign = "-"
			yyVAL.tUnaryOperator = mUnaryOperator
		}
	case 120:
		//line grammar.y:1310
		{
			mUnaryOperator = new(UnaryOperator)
			mUnaryOperator.Sign = "~"
			yyVAL.tUnaryOperator = mUnaryOperator
		}
	case 121:
		//line grammar.y:1316
		{
			mUnaryOperator = new(UnaryOperator)
			mUnaryOperator.Sign = "!"
			yyVAL.tUnaryOperator = mUnaryOperator
		}
	case 122:
		//line grammar.y:1325
		{
			mExprUnary = new(ExprUnary)
			mExprUnary.ExprPostfix = yyS[yypt-0].tExprPostfix
			mExprCmd = nil
			rlog2("unary_expression:postfix_expression\n", mExprPostfix, mExprUnary)
			yyVAL.tExprUnary = mExprUnary
		}
	case 123:
		//line grammar.y:1333
		{
			mUnaryOperator = new(UnaryOperator)
			mUnaryOperator.Sign = "++"

			mExprUnary = new(ExprUnary)
			mExprUnary.UnaryOperator = mUnaryOperator
			yyVAL.tExprUnary = mExprUnary
			rlog2("unary_expression:INC_OP postfix_expression\n")
		}
	case 124:
		//line grammar.y:1343
		{
			mUnaryOperator = new(UnaryOperator)
			mUnaryOperator.Sign = "--"

			mExprUnary = new(ExprUnary)
			mExprUnary.UnaryOperator = mUnaryOperator
			yyVAL.tExprUnary = mExprUnary
			rlog2("unary_expression:DEC_OP postfix_expression\n")
		}
	case 125:
		//line grammar.y:1353
		{
			mExprUnary = new(ExprUnary)
			mExprUnary.UnaryOperator = yyS[yypt-1].tUnaryOperator
			mExprUnary.ExprCast = yyS[yypt-0].tExprCast
			yyVAL.tExprUnary = mExprUnary
			rlog2("unary_expression cast_expression:postfix_expression\n")
		}
	case 126:
		//line grammar.y:1364
		{
			mExprPrimary = new(ExprPrimary)
			mExprPrimary.Value = yyS[yypt-0].id
			mExprPrimary.IdType = IDENTIFIER
			rlog2("primary_expression:IDENTIFIER>>\n", mExprPrimary, yyS[yypt-0].id)
			yyVAL.tExprPrimary = mExprPrimary
		}
	case 127:
		//line grammar.y:1372
		{
			mExprPrimary = new(ExprPrimary)
			//mExprPrimary.Value = nil
			mExprPrimary.IdType = EXPR_ID
			mExprPrimary.ExprAssign = mExprAssign
			rlog2("primary_expression:EXPR_ID>> mExprPrimary=%v ExprAssign=%v \n", mExprPrimary, mExprPrimary.ExprAssign)
			yyVAL.tExprPrimary = mExprPrimary
		}
	case 128:
		//line grammar.y:1381
		{
			mExprPrimary = new(ExprPrimary)
			mExprPrimary.Value = yyS[yypt-0].id
			mExprPrimary.IdType = TRUE_ID
			yyVAL.tExprPrimary = mExprPrimary
			rlog2("primary_expression:TRUE_ID>>\n", mExprPrimary, yyS[yypt-0].id)
		}
	case 129:
		//line grammar.y:1389
		{
			mExprPrimary = new(ExprPrimary)
			mExprPrimary.Value = yyS[yypt-0].id
			mExprPrimary.IdType = FALSE_ID
			yyVAL.tExprPrimary = mExprPrimary
			rlog2("primary_expression:FALSE_ID>>\n", mExprPrimary, yyS[yypt-0].id)
		}
	case 130:
		//line grammar.y:1397
		{
			mExprPrimary = new(ExprPrimary)
			mExprPrimary.Value = yyS[yypt-0].id
			mExprPrimary.IdType = FLOAT_ID
			rlog2("primary_expression:FLOAT_ID>>\n", mExprPrimary, yyS[yypt-0].id)
			yyVAL.tExprPrimary = mExprPrimary
		}
	case 131:
		//line grammar.y:1405
		{
			mExprPrimary = new(ExprPrimary)
			mExprPrimary.Value = yyS[yypt-0].id
			mExprPrimary.IdType = LITERAL_STRING
			rlog2("primary_expression:LITERAL_STRING>>\n", mExprPrimary, yyS[yypt-0].id)
			yyVAL.tExprPrimary = mExprPrimary
		}
	case 132:
		//line grammar.y:1413
		{
			mExprPrimary = new(ExprPrimary)
			mExprPrimary.Value = yyS[yypt-0].id
			mExprPrimary.IdType = INTEGER_ID
			rlog2("primary_expression:INTEGER_ID>>\n", mExprPrimary, yyS[yypt-0].id)
			yyVAL.tExprPrimary = mExprPrimary
		}
	case 133:
		//line grammar.y:1421
		{
			mExprPrimary = new(ExprPrimary)
			mExprPrimary.Value = yyS[yypt-0].id
			mExprPrimary.IdType = EMAIL_ID
			rlog2("primary_expression:EMAIL_ID>>\n", mExprPrimary, yyS[yypt-0].id)
			yyVAL.tExprPrimary = mExprPrimary
		}
	case 134:
		//line grammar.y:1429
		{
			mExprPrimary = new(ExprPrimary)
			mExprPrimary.Value = yyS[yypt-0].id
			mExprPrimary.IdType = UUID_ID
			rlog2("primary_expression:UUID_ID>>\n", mExprPrimary, yyS[yypt-0].id)
			yyVAL.tExprPrimary = mExprPrimary
		}
	case 135:
		//line grammar.y:1440
		{
			mExprPostfix = new(ExprPostfix)
			mExprPostfix.ExprPrimary = yyS[yypt-0].tExprPrimary
			mExprPostfix.Operator = ""
			rlog2("primary_expression: ", mExprPostfix, mExprPostfix.ExprPrimary)
			yyVAL.tExprPostfix = mExprPostfix
		}
	case 136:
		//line grammar.y:1448
		{
			mExprPostfix = new(ExprPostfix)
			mExprPostfix.ExprPostfix = yyS[yypt-3].tExprPostfix
			mExprPostfix.Expr = yyS[yypt-1].tExpr
			mExprPostfix.Operator = ""
			mExprPostfix.Modifier = "[]"
			yyVAL.tExprPostfix = mExprPostfix
			rlog2("postfix_expression '[' ']'\n")
		}
	case 137:
		//line grammar.y:1458
		{
			mExprPostfix = new(ExprPostfix)
			mExprPostfix.ExprPostfix = yyS[yypt-2].tExprPostfix
			mExprPostfix.Operator = ""
			mExprPostfix.Modifier = "()"
			yyVAL.tExprPostfix = mExprPostfix
			rlog2("postfix_expression '(' ')'\n")
		}
	case 138:
		//line grammar.y:1467
		{
			mExprPostfix = new(ExprPostfix)
			mExprPostfix.ExprPostfix = yyS[yypt-3].tExprPostfix
			mExprPostfix.ExprListArgument = yyS[yypt-1].tExprListArgument
			mExprPostfix.Operator = ""
			mExprPostfix.Modifier = "()"
			yyVAL.tExprPostfix = mExprPostfix
			rlog2("postfix_expression '(' argument_expression_list ')'\n")
		}
	case 139:
		//line grammar.y:1477
		{
			mExprPostfix = new(ExprPostfix)
			mExprPostfix.ExprPostfix = yyS[yypt-2].tExprPostfix
			mExprPostfix.Identifier = new(Identifier)
			mExprPostfix.Identifier.Name = yyS[yypt-0].id
			mExprPostfix.Operator = "."
			yyVAL.tExprPostfix = mExprPostfix
			rlog2("postfix_expression '.' IDENTIFIER\n", yyVAL.tExprPostfix, yyS[yypt-0].id)
		}
	case 140:
		//line grammar.y:1487
		{
			mExprPostfix = new(ExprPostfix)
			mExprPostfix.ExprPostfix = yyS[yypt-2].tExprPostfix
			mExprPostfix.Identifier = new(Identifier)
			mExprPostfix.Identifier.Name = yyS[yypt-0].id
			mExprPostfix.Operator = "PTR_OP"
			yyVAL.tExprPostfix = mExprPostfix
			rlog2("postfix_expression PTR_OP IDENTIFIER\n")
		}
	case 141:
		//line grammar.y:1497
		{
			mExprPostfix = new(ExprPostfix)
			mExprPostfix.ExprPostfix = yyS[yypt-1].tExprPostfix
			mExprPostfix.Identifier = new(Identifier)
			mExprPostfix.Operator = "INC_OP"
			yyVAL.tExprPostfix = mExprPostfix
			rlog2("postfix_expression INC_OP\n")
		}
	case 142:
		//line grammar.y:1506
		{
			mExprPostfix = new(ExprPostfix)
			mExprPostfix.ExprPostfix = yyS[yypt-1].tExprPostfix
			mExprPostfix.Identifier = new(Identifier)
			mExprPostfix.Operator = "DEC_OP"
			yyVAL.tExprPostfix = mExprPostfix
			rlog2("postfix_expression DEC_OP\n")
		}
	case 143:
		//line grammar.y:1518
		{
			mExprListArgument = new(ExprListArgument)
			mExprListArgument.ExprAssign = yyS[yypt-0].tExprAssign
			yyVAL.tExprListArgument = mExprListArgument
			rlog2("argument_expression_list>>assignment_expression\n")
		}
	case 144:
		//line grammar.y:1525
		{
			mExprListArgument = new(ExprListArgument)
			mExprListArgument.ExprListArgument = yyS[yypt-2].tExprListArgument
			mExprListArgument.ExprAssign = yyS[yypt-0].tExprAssign
			yyVAL.tExprListArgument = mExprListArgument
			rlog2("argument_expression_list ',' assignment_expression\n")
		}
	case 145:
		//line grammar.y:1537
		{
			mExprCmd = new(ExprCmd)
			mExprCmd.CmdListApp = yyS[yypt-0].tCmdListApp
			//rlog2("list_application_command:",mCmdListApp, mExprCmd)
			mCmdMakeAppUser = nil
			mCmdListApp = nil
			mCmdVerify = nil
			mExprUnary = nil
			yyVAL.tExprCmd = mExprCmd
		}
	case 146:
		//line grammar.y:1548
		{
			mExprCmd = new(ExprCmd)
			mExprCmd.CmdVerify = mCmdVerify
			//rlog2("verify_application_url_category_command:\n",mCmdVerify, mExprCmd)
			mCmdMakeAppUser = nil
			mCmdListApp = nil
			mCmdVerify = nil
			mExprUnary = nil
			yyVAL.tExprCmd = mExprCmd
		}
	case 147:
		//line grammar.y:1559
		{
			mExprCmd = new(ExprCmd)
			mExprCmd.CmdGenData = yyS[yypt-0].tCmdGenData
			//rlog2("generate_data_command\n",mCmdGenData,mExprCmd)
			mCmdMakeAppUser = nil
			mCmdListApp = nil
			mCmdVerify = nil
			mExprUnary = nil
			yyVAL.tExprCmd = mExprCmd
		}
	case 148:
		//line grammar.y:1570
		{
			mExprCmd = new(ExprCmd)
			mExprCmd.CmdMakeAppUser = mCmdMakeAppUser
			//rlog2("application_makeappuser_command\n",mCmdMakeAppUser,mExprCmd)
			mCmdMakeAppUser = nil
			mCmdListApp = nil
			mCmdVerify = nil
			mExprUnary = nil
			yyVAL.tExprCmd = mExprCmd
		}
	case 149:
		//line grammar.y:1581
		{
			mExprCmd = new(ExprCmd)
			mExprCmd.CmdMakeAppTid = yyS[yypt-0].tCmdMakeAppTid
			//rlog2("application_makeapptid_command\n",mCmdMakeAppTid,mExprCmd)
			mCmdMakeAppUser = nil
			mCmdListApp = nil
			mCmdVerify = nil
			mExprUnary = nil
			yyVAL.tExprCmd = mExprCmd
		}
	case 150:
		//line grammar.y:1592
		{
			mExprCmd = new(ExprCmd)
			mExprCmd.CmdTestRoute = yyS[yypt-0].tCmdTestRoute
			yyVAL.tExprCmd = mExprCmd
		}
	case 151:
		//line grammar.y:1598
		{
			mExprCmd = new(ExprCmd)
			mExprCmd.CmdCreateApp = yyS[yypt-0].tCmdCreateApp
			yyVAL.tExprCmd = mExprCmd
		}
	case 152:
		//line grammar.y:1604
		{
			mExprCmd = new(ExprCmd)
			mExprCmd.CmdDeleteApp = yyS[yypt-0].tCmdDeleteApp
			yyVAL.tExprCmd = mExprCmd
		}
	case 153:
		//line grammar.y:1610
		{
			mExprCmd = new(ExprCmd)
			mExprCmd.CmdDeleteAppTid = yyS[yypt-0].tCmdDeleteAppTid
			yyVAL.tExprCmd = mExprCmd
		}
	case 154:
		//line grammar.y:1616
		{
			mExprCmd = new(ExprCmd)
			mExprCmd.CmdRunSQL = yyS[yypt-0].tCmdRunSQL
			yyVAL.tExprCmd = mExprCmd
		}
	case 155:
		//line grammar.y:1622
		{
			mExprCmd = new(ExprCmd)
			mExprCmd.CmdDeleteAppUser = yyS[yypt-0].tCmdDeleteAppUser
			yyVAL.tExprCmd = mExprCmd
		}
	case 156:
		//line grammar.y:1631
		{
			mCmdMakeAppUser = new(CmdMakeAppUser)
			mCmdMakeAppUser.Name = "MakeAppUser"
			mCmdMakeAppUser.Expr1 = yyS[yypt-4].tExpr
			mCmdMakeAppUser.Expr2 = yyS[yypt-1].tExpr
			rlog2("MAKEAPPUSER_ID \n", mCmdMakeAppUser)
			yyVAL.tCmdMakeAppUser = mCmdMakeAppUser
		}
	case 157:
		//line grammar.y:1643
		{
			rlog2("GENDATA_ID expression ';'\n")
			mCmdGenData = new(CmdGenData)
			mCmdGenData.Name = "GENDATA"
			mCmdGenData.ExprApp = yyS[yypt-4].tExpr
			mCmdGenData.ExprLoop = yyS[yypt-1].tExpr
			yyVAL.tCmdGenData = mCmdGenData
		}
	case 158:
		//line grammar.y:1655
		{
			mCmdVerify = new(CmdVerify)
			mCmdVerify.Name = "verify"
			rlog2("VERIFY_ID \n", mCmdVerify)
			yyVAL.tCmdVerify = mCmdVerify
		}
	case 159:
		//line grammar.y:1665
		{
			mCmdListApp = new(CmdListApp)
			mCmdListApp.Name = "List Apps"
			rlog2("LIST_ID APPS_ID\n", mCmdListApp)
			yyVAL.tCmdListApp = mCmdListApp
		}
	case 160:
		//line grammar.y:1675
		{
			mCmdMakeAppTid = new(CmdMakeAppTid)
			mCmdMakeAppTid.Name = "MAKEAPPTID"
			mCmdMakeAppTid.Expr1 = yyS[yypt-4].tExpr
			mCmdMakeAppTid.Expr2 = yyS[yypt-1].tExpr
			rlog2("MAKEAPPTID_ID:\n", mCmdMakeAppTid)
			yyVAL.tCmdMakeAppTid = mCmdMakeAppTid
		}
	case 161:
		//line grammar.y:1687
		{
			mCmdTestRoute = new(CmdTestRoute)
			mCmdTestRoute.Name = "TESTROUTE"
			mCmdTestRoute.Expr1 = yyS[yypt-4].tExpr
			mCmdTestRoute.Expr2 = yyS[yypt-1].tExpr
			//rlog2("test_route_command:\n",mCmdTestRoute)
			yyVAL.tCmdTestRoute = mCmdTestRoute
		}
	case 162:
		//line grammar.y:1699
		{
			mCmdCreateApp = new(CmdCreateApp)
			mCmdCreateApp.Name = "CREATEAPP"
			mCmdCreateApp.Expr1 = yyS[yypt-1].tExpr
			//rlog2("create_application_command:\n",mCmdCreateApp)
			yyVAL.tCmdCreateApp = mCmdCreateApp
		}
	case 163:
		//line grammar.y:1710
		{
			mCmdDeleteApp = new(CmdDeleteApp)
			mCmdDeleteApp.Name = "DELETEAPP"
			mCmdDeleteApp.Expr1 = yyS[yypt-1].tExpr
			//rlog2("delete_application_command:\n",mCmdDeleteApp)
			yyVAL.tCmdDeleteApp = mCmdDeleteApp
		}
	case 164:
		//line grammar.y:1721
		{
			mCmdDeleteAppTid = new(CmdDeleteAppTid)
			mCmdDeleteAppTid.Name = "DELETEAPPTID"
			mCmdDeleteAppTid.Expr1 = yyS[yypt-1].tExpr
			mCmdDeleteAppTid.Expr2 = yyS[yypt-4].tExpr
			//rlog2("delete_application_command:\n",mCmdDeleteAppTid)
			yyVAL.tCmdDeleteAppTid = mCmdDeleteAppTid
		}
	case 165:
		//line grammar.y:1733
		{
			mCmdRunSQL = new(CmdRunSQL)
			mCmdRunSQL.Name = "RUNSQL"
			mCmdRunSQL.Expr1 = yyS[yypt-5].tExpr
			mCmdRunSQL.Expr2 = yyS[yypt-1].tExpr
			//rlog2("delete_application_command:\n",mCmdRunSQL)
			yyVAL.tCmdRunSQL = mCmdRunSQL
		}
	case 166:
		//line grammar.y:1745
		{
			mCmdDeleteAppUser = new(CmdDeleteAppUser)
			mCmdDeleteAppUser.Name = "DELETE_APPUSER_ID"
			mCmdDeleteAppUser.Expr1 = yyS[yypt-5].tExpr
			mCmdDeleteAppUser.Expr2 = yyS[yypt-1].tExpr
			//rlog2("delete_application_command:\n",mCmdDeleteAppUser)
			yyVAL.tCmdDeleteAppUser = mCmdDeleteAppUser
		}
	}
	goto yystack /* stack new state and value */
}
