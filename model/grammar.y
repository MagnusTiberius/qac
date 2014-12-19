/*

FILE:       grammar.y
AUTHOR:     Ben Gonzales
LANGUAGE:   Go
TOOLS:      yacc

TO RUN: 
    go tool yacc grammar.y && go run y.go lexer.go prv5main.go ast.go astdump.go runtime.go rtwalk.go eval.go Func.go -file=/test1/test4.ias

Test Data:

    package "test2"

    import (
        "test3"
        "test4"
        "test5"
        "category1/test1"
        "pdf"
    )

    verification TestFeature() {

        string compareData4;

        compareData4 = "[{'text':'8 64-bit','tally':116}, 
              {'text':'7 64-bit','tally':  3}, 
              {'text':'3 64-bit','tally':  4}, 
              {'text':'5 64-bit','tally':  2}, 
              {'text':'1 64-bit','tally':  1}]"; 

        step name1 ( param1 ) 
            stimulus {   
                description = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque faucibus tempor risus. Vivamus pretium volutpat arcu eu elementum. Aenean a augue ut diam feugiat vehicula sed non ipsum. Integer luctus sed elit ut bibendum. Pellentesque imperdiet lorem eu ligula molestie, a feugiat ligula feugiat. Nulla a mi tellus.";

                param1 = verify testapp23 urlid appversion with (compareData4) ;
            }
            response {   
                description = "Mauris luctus justo eget arcu condimentum, sed finibus augue sollicitudin. Morbi sed feugiat lacus. Aliquam tristique nec odio eget venenatis. Mauris volutpat molestie velit, nec egestas odio auctor in. Pellentesque neque eros, volutpat id iaculis in, mollis eu nisi. Phasellus sed semper dolor.";

                pdf.result = param1;
            };

        step name2 ( param2 )  
            stimulus {   
                description = "Maecenas nisl augue, dapibus eu dignissim eget, mollis at risus. Fusce varius velit a odio iaculis cursus. In blandit est vitae efficitur luctus. Donec accumsan congue eros, non tristique enim euismod ut. Curabitur mauris nibh, fringilla sit amet leo eu, efficitur lobortis magna.";

                param2 = verify testapp23 urlid appversion with (compareData4) ;
            }
            response {   
                description = "Vivamus orci elit, volutpat quis maximus nec, dignissim quis lorem. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Donec justo augue, dictum at urna sit amet, suscipit egestas velit. Integer bibendum sagittis ipsum sit amet mollis. Morbi erat purus, ullamcorper ut tempus et, vulputate vel diam.";

                param2 == compareData4;
            };

        step name3 ( param3 ) 
            stimulus {   
                param3 = verify testapp23 urlid appversion with (compareData4) ;
            }
            response {   
                param3 == compareData4;
            };

        return (bbbbbbbb && cccccccc && ddddddd);
    }


*/


%{

package model

import (
   // "fmt"
)


%}


// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.
%union{
    val                     int
    id                      string
    sel                     []string
    tExprAdditive           *ExprAdditive
    tExprMultiplicative     *ExprMultiplicative
    tExprCast               *ExprCast
    tExprUnary              *ExprUnary
    tExprPrimary            *ExprPrimary
    tExprCmd                *ExprCmd
    tExprPostfix            *ExprPostfix
    tExprAssign             *ExprAssign
    tExpr                   *Expr
    tExprStmt               *ExprStmt
    tStmt                   *Stmt
    tStmtList               *StmtList
    tCmpndStmt              *CmpndStmt
    tDeclarationList        *DeclarationList
    tDeclaration            *Declaration
    tExprShift              *ExprShift
    tExprRelational         *ExprRelational
    tExprEquality           *ExprEquality
    tExprAnd                *ExprAnd
    tExprExclOr             *ExprExclOr
    tExprInclOr             *ExprInclOr
    tExprLogicAnd           *ExprLogicAnd
    tExprLogicOr            *ExprLogicOr
    tExprConditional        *ExprConditional
    tSpecDecl               *SpecDecl
    tSpecType               *SpecType
    tDeclInitList           *DeclInitList
    tDeclInit               *DeclInit
    tDeclarator             *Declarator
    tInitializer            *Initializer
    tDeclDirect             *DeclDirect
    tFuncDef                *FuncDef
    tDeclExternal           *DeclExternal
    tTranslationUnit        *TranslationUnit
    tPackageDef             *PackageDef
    tImportDef              *ImportDef
    tImportList             *ImportList
    tExprConstant           *ExprConstant
    tCmdMakeAppUser         *CmdMakeAppUser
    tCmdGenData             *CmdGenData
    tCmdVerify              *CmdVerify
    tCmdListApp             *CmdListApp
    tCmdMakeAppTid          *CmdMakeAppTid
    tStmtJump               *StmtJump
    tExprListArgument       *ExprListArgument
    tStmtIteration          *StmtIteration
    tStmtWhile              *StmtWhile
    tStmtFor                *StmtFor
    tStmtDo                 *StmtDo
    tStmtSelection          *StmtSelection
    tInitializerList        *InitializerList
    tUnaryOperator          *UnaryOperator
    tParamDecl              *ParamDecl
    tParamList              *ParamList
    tParamTypeList          *ParamTypeList
    tIdentifierList         *IdentifierList
    tCmdTestRoute           *CmdTestRoute
    tCmdCreateApp           *CmdCreateApp
    tCmdDeleteApp           *CmdDeleteApp
    tCmdDeleteAppTid        *CmdDeleteAppTid
    tCmdRunSQL              *CmdRunSQL
    tCmdDeleteAppUser       *CmdDeleteAppUser
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
                



%type <tExprAdditive>           additive_expression
%type <tExprMultiplicative>     multiplicative_expression
%type <tExprCast>               cast_expression
%type <tExprUnary>              unary_expression
%type <tExprPrimary>            primary_expression
%type <tExprCmd>                command_expression
%type <tExprPostfix>            postfix_expression
%type <tExprAssign>             assignment_expression
%type <tExpr>                   expression
%type <tExprStmt>               expression_statement
%type <tStmt>                   statement
%type <tStmtList>               statement_list
%type <tCmpndStmt>              compound_statement
%type <tDeclarationList>        declaration_list
%type <tDeclaration>            declaration
%type <tExprShift>              shift_expression
%type <tExprRelational>         relational_expression
%type <tExprEquality>           equality_expression
%type <tExprAnd>                and_expression
%type <tExprExclOr>             exclusive_or_expression
%type <tExprInclOr>             inclusive_or_expression
%type <tExprLogicAnd>           logical_and_expression
%type <tExprLogicOr>            logical_or_expression
%type <tExprConditional>        conditional_expression
%type <tSpecDecl>               declaration_specifiers
%type <tSpecType>               type_specifier
%type <tDeclInitList>           init_declarator_list
%type <tDeclInit>               init_declarator
%type <tDeclarator>             declarator
%type <tInitializer>             initializer
%type <tDeclDirect>             direct_declarator
%type <tFuncDef>                function_definition
%type <tDeclExternal>           external_declaration
%type <tTranslationUnit>        translation_unit
%type <tPackageDef>             package_definition
%type <tImportDef>              import_definition
%type <tImportList>             import_list
%type <tExprConstant>           constant_expression
%type <tCmdMakeAppUser>         application_makeappuser_command
%type <tCmdGenData>             generate_data_command
%type <tCmdVerify>              verify_application_url_category_command
%type <tCmdListApp>             list_application_command
%type <tCmdMakeAppTid>          application_makeapptid_command
%type <tStmtJump>               jump_statement
%type <tExprListArgument>       argument_expression_list
%type <tStmtIteration>          iteration_statement
%type <tStmtSelection>          selection_statement
%type <tStmtSelection>          selection_statement_else

%type <tInitializerList>        initializer_list
%type <tUnaryOperator>          unary_operator
%type <tParamDecl>              parameter_declaration
%type <tParamList>              parameter_list
%type <tParamTypeList>          parameter_type_list
%type <tIdentifierList>         identifier_list
%type <tCmdTestRoute>           test_route_command
%type <tCmdCreateApp>           create_application_command
%type <tCmdDeleteApp>           delete_application_command
%type <tCmdDeleteAppTid>        delete_application_tid_command
%type <tCmdRunSQL>              run_sql_command
%type <tCmdDeleteAppUser>       delete_application_user_command








// same for terminals
%token <id> IDENTIFIER
%token <id> CREATE_ID HUB_ID APP_ID NAME_ID EQUAL_ID
            DROP_ID STYLE_ID LIST_ID APPS_ID RUN_ID
            FILE_SPEC_ID USE_ID TEST_ID
            MAKEAPPUSER_ID MAKEAPPTID_ID
            IN_ID EMAIL_ID UUID_ID
            COMMENT_SECTION_ID STRING_CONST_ID 
            GENDATA_ID LOOP_ID INTEGER_ID FLOAT_ID
            POPULATE_ID FROM_ID
            VERIFY_ID URLID_ID WITH_ID
            IMPORT_ID PACKAGE_ID
            FUNC_ID
            TYPE_STRING_ID TYPE_INT_ID TYPE_FLOAT_ID TYPE_BOOL_ID TYPE_TESTPLAN_ID
            TRUE_ID FALSE_ID
            BAD_ID
            LITERAL_STRING
            LOAD_ID
            NEWLINE
            OR_OP AND_OP EQ_OP NE_OP LE_OP GE_OP LEFT_OP RIGHT_OP
            INC_OP DEC_OP PTR_OP
            ADD_ASSIGN SUB_ASSIGN MUL_ASSIGN DIV_ASSIGN MOD_ASSIGN AND_ASSIGN XOR_ASSIGN OR_ASSIGN
            IF_ID ELSE_ID ENDIF_ID NO_ELSE_ID
            RETURN_ID
            EXPR_ID
            WHILE_ID FOR_ID DO_ID
            STEP_ID STIMULUS_ID RESPONSE_ID END_STEP_ID
            FUNC_CALL_ID OBJECT_FIELD_REF
            TESTROUTE_ID
            DELETE_ID
            TID_ID
            RUNSQL_ID
            DELETE_APPUSER_ID


%left '.'
%left '|'
%left '&'
%left '+'  '-'
%left '*'  '/'  '%'
%left UMINUS      /*  supplies  precedence  for  unary  minus  */

//
// http://efxa.org/2014/05/17/techniques-for-resolving-common-grammar-conflicts-in-parsers/
//
%nonassoc NO_ELSE_ID
%nonassoc ELSE_ID

%%

translation_unit 
    : external_declaration
        {
            mTranslationUnit = new(TranslationUnit)
            mTranslationUnit.DeclExternal = $1
            $$ = mTranslationUnit
            rlog2("external_declaration ET=%v \n", $1)
        }
    | translation_unit external_declaration
        {
            mTranslationUnit = new(TranslationUnit)
            mTranslationUnit.TranslationUnit = $1
            mTranslationUnit.DeclExternal = $2
            $$ = mTranslationUnit
            rlog2("translation_unit external_declaration $1=%v $1.ET=%v $2=%v \n", $1, $1.DeclExternal, $2)
        }
    ;

external_declaration
    : package_definition
        {
            mDeclExternal = new(DeclExternal)
            mDeclExternal.PackageDef = $1
            $$ = mDeclExternal
            rlog2("package_definition\n", $$, $1)
        }
    | function_definition
        {
            mDeclExternal = new(DeclExternal)
            mDeclExternal.FuncDef = $1
            $$ = mDeclExternal
            rlog2("function_definition\n", $$, $1)
        }
    | import_definition
        {
            mDeclExternal = new(DeclExternal)
            mDeclExternal.ImportDef = $1
            $$ = mDeclExternal
            rlog2("import_definition\n", $$, $1)
        }
    ;

package_definition
    : PACKAGE_ID LITERAL_STRING
        {
            mPackageDef = new(PackageDef)
            mPackageDef.Name = $2
            rlog2("package_definition > PACKAGE_ID LITERAL_STRING\n", mPackageDef, $2)
            $$ = mPackageDef
        }
    ;

import_definition
    : IMPORT_ID LITERAL_STRING
        {
            mImportList = new(ImportList)
            mImportList.List = []string{$2}

            mImportDef = new(ImportDef)
            mImportDef.ImportList = mImportList
            $$ = mImportDef

            rlog2("IMPORT_ID LITERAL_STRING>>\n", mImportDef)
        }
    | IMPORT_ID '(' import_list ')'
        {
            mImportDef = new(ImportDef)
            mImportDef.ImportList = $3
            $$ = mImportDef
            rlog2("IMPORT_ID '(' import_list ')' >>\n", mImportDef, mImportDef.ImportList)
        }
    ;

import_list
    : LITERAL_STRING
        {
            mImportList = new(ImportList)
            mImportList.List = []string{}
            mImportList.List = append(mImportList.List, $1)
            $$ = mImportList
            rlog2("IMPORT_ID >> LITERAL_STRING", $1)
        }
    | import_list LITERAL_STRING
        {
            $$.List = append($$.List, $2)
            rlog2("import_list >> LITERAL_STRING\n", $2)
        }
    ;

function_definition
    : declarator compound_statement
        {
            mFuncDef = new(FuncDef)
            mFuncDef.Declarator = $1
            mFuncDef.CmpndStmt = $2
            $$ = mFuncDef
            rlog2("function_definition >> declarator compound_statement : \n", $1)
        }
    | declaration_specifiers declarator compound_statement
        {
            mFuncDef = new(FuncDef)
            mFuncDef.SpecDecl = $1
            mFuncDef.Declarator = $2
            mFuncDef.CmpndStmt = $3
            $$ = mFuncDef

            rlog2("function_definition >> declaration_specifiers declarator compound_statement : $$=%v CS=%v \n", $$, $3)
            //DumpFuncDef(mFuncDef,3)
        }
    | COMMENT_SECTION_ID
        {

        }
    ;

declarator
    : direct_declarator
        {
            mDeclarator = new(Declarator)
            mDeclarator.DeclDirect = mDeclDirect
            $$ = mDeclarator
            rlog2("declarator >> direct_declarator:\n", mDeclarator, mDeclarator.DeclDirect, mDeclarator.DeclDirect.Identifier)
        }
    ;

direct_declarator
    : IDENTIFIER
        {
            mDeclDirect = new(DeclDirect)
            mIdentifier = new(Identifier)
            mIdentifier.Name = $1
            mDeclDirect.Identifier = mIdentifier
            rlog2("direct_declarator IDENTIFIER-->\n", mDeclDirect, mIdentifier)
            $$ = mDeclDirect
        }
    | '(' declarator ')'
        {
            //rlog2("direct_declarator '(' declarator ')'\n")
            mDeclDirect = new(DeclDirect)
            mDeclDirect.Declarator = $2
            mDeclDirect.Modifier = "()"
            $$ = mDeclDirect
        }
    | direct_declarator '[' constant_expression ']'
        {
            mDeclDirect = new(DeclDirect)
            mDeclDirect.DeclDirect = $1
            mDeclDirect.ExprConstant = $3
            mDeclDirect.Modifier = "[]"
            $$ = mDeclDirect
            rlog2("direct_declarator '[' constant_expression ']'\n", $$)
        }
    | direct_declarator '[' ']'
        {
            mDeclDirect = new(DeclDirect)
            mDeclDirect.DeclDirect = $1
            mDeclDirect.Modifier = "[]"
            $$ = mDeclDirect
            rlog2("direct_declarator '[' ']'")
        }
    | direct_declarator '(' parameter_type_list ')'
        {
            mDeclDirect = new(DeclDirect)
            mDeclDirect.DeclDirect = $1
            mDeclDirect.ParamTypeList = $3
            mDeclDirect.Modifier = "()"
            $$ = mDeclDirect
            rlog2("direct_declarator '(' parameter_type_list ')'\n", mPackageDef, mFuncDef, mParamDeclList)
        }
    | direct_declarator '(' identifier_list ')'    
        {    
            mDeclDirect = new(DeclDirect)
            mDeclDirect.DeclDirect = $1
            mDeclDirect.IdentifierList = $3
            mDeclDirect.Modifier = "()"
            $$ = mDeclDirect
            rlog2("direct_declarator '(' identifier_list ')'\n")
        }
    | direct_declarator '(' ')'
        {
            mDeclDirect = new(DeclDirect)
            mDeclDirect.DeclDirect = $1
            mDeclDirect.Modifier = "()"
            $$ = mDeclDirect
            rlog2("direct_declarator '(' ')'\n")
        }
    ;

constant_expression
    : conditional_expression
        {
            mExprConstant = new(ExprConstant)
            mExprConstant.ExprConditional =  $1
            $$ = mExprConstant
            rlog2("constant_expression:conditional_expression\n", $1)
        }
    ;

identifier_list
    : IDENTIFIER
        {
            mIdentifier = new(Identifier)
            mIdentifier.Name = $1

            mIdentifierList = new(IdentifierList)
            mIdentifierList.Identifier = mIdentifier
            $$ = mIdentifierList
        }
    | identifier_list ',' IDENTIFIER
        {
            mIdentifier = new(Identifier)
            mIdentifier.Name = $3

            mIdentifierList = new(IdentifierList)
            mIdentifierList.Identifier = mIdentifier
            mIdentifierList.IdentifierList = $1
            $$ = mIdentifierList
        }
    ;    

parameter_type_list
    : parameter_list
        {
            mParamTypeList = new(ParamTypeList)
            mParamTypeList.ParamList = $1
            $$ = mParamTypeList
            rlog2("parameter_type_list>>parameter_list: \n", mParamTypeList)
        }
    ;

parameter_list
    : parameter_declaration
        {
            mParamList = new(ParamList)
            mParamList.ParamDecl = $1 
            $$ = mParamList

            rlog2("parameter_list>>parameter_declaration-->\n", $$)
        }
    | parameter_list ',' parameter_declaration    
        {
            mParamList = new(ParamList)
            mParamList.ParamDecl = $3 
            mParamList.ParamList = $1 
            $$ = mParamList
            rlog2("parameter_list>>parameter_list ',' parameter_declaration -->\n", $$)
        }

parameter_declaration
    : declaration_specifiers declarator
        {
            mParamDecl = new(ParamDecl)
            mParamDecl.SpecDecl = $1
            mParamDecl.Declarator = $2
            $$ = mParamDecl
            rlog2("parameter_declaration >> declaration_specifiers declarator:\n", $$, $1, $2)
        }
    ;

compound_statement
    : '{' '}'
        {
            mCmpndStmt = new(CmpndStmt)
            $$ = mCmpndStmt
            rlog2("'{' '}'\n", $$)
        }
    | '{' statement_list '}'
        {
            mCmpndStmt = new(CmpndStmt)
            mCmpndStmt.StmtList = $2
            $$ = mCmpndStmt
            rlog2("compound_statement >> '{' statement_list '}' :\n", $$)

        }
    | '{' declaration_list '}'
        {
            mCmpndStmt = new(CmpndStmt)
            mCmpndStmt.DeclarationList = $2
            $$ = mCmpndStmt
            rlog2("compound_statement >> '{' declaration_list '}' :\n", $$, $2, $2.Declaration)
        }
    | '{' declaration_list statement_list '}'
        {
            mCmpndStmt = new(CmpndStmt)
            mCmpndStmt.DeclarationList = $2
            mCmpndStmt.StmtList = $3
            $$ = mCmpndStmt
            rlog2("compound_statement >> '{' declaration_list statement_list '}'\n", $$)
        }
    ;

statement_list
    : statement
        {
            mStmtList = new(StmtList)
            mStmtList.Stmt = $1
            $$ = mStmtList
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
    | statement_list statement
        {
            mStmtList = new(StmtList)
            mStmtList.Stmt = $2
            mStmtList.StmtList = $1
            $$ = mStmtList
        }

    ;

statement
    : compound_statement
        {
            mStmt = new(Stmt)
            mStmt.CmpndStmt = $1
            $$ = mStmt
            rlog2("statement:compound_statement:\n", mExprStmt, mExprAssign)
        }
    | expression_statement
        {
            mStmt = new(Stmt)
            mStmt.ExprStmt = $1
            $$ = mStmt
        }
    | selection_statement
        {
            mStmt = new(Stmt)
            mStmt.StmtSelection = $1
            $$ = mStmt
            rlog2("selection_statement\n")
        }
    | selection_statement_else
        {
            mStmt = new(Stmt)
            mStmt.StmtSelection = $1
            $$ = mStmt
            rlog2("selection_statement\n")
        }
    | jump_statement
        {
            rlog2("jump_statement\n")
            mStmt = new(Stmt)
            mStmt.StmtJump = $1
            $$ = mStmt
        }
    | iteration_statement
        {
            mStmt = new(Stmt)
            mStmt.StmtIteration = $1
            $$ = mStmt
            //DumpStmtIteration(mStmt.StmtIteration, 3)
        }
    ;    

jump_statement
    : RETURN_ID ';'
        {
            mStmtJump = new(StmtJump)
            mStmtJump.Name = "return"
            $$ = mStmtJump
        }
    | RETURN_ID expression ';'
        {
            mStmtJump = new(StmtJump)
            mStmtJump.Name = "return"
            mStmtJump.Expr = $2
            $$ = mStmtJump
            rlog2("RETURN_ID expression ';'\n")
        }
    ;

selection_statement
    : IF_ID '(' expression ')' statement %prec NO_ELSE_ID
        {
            mStmtSelection = new(StmtSelection)
            mStmtSelection.Expr = $3
            mStmtSelection.Stmt = $5
            $$ = mStmtSelection
            rlog(1, "IF_ID '(' expression ')' statement\n")
            //DumpStmtSelection(mStmtSelection,1)
        }
    ;
    
selection_statement_else        
    : IF_ID '(' expression ')' statement ELSE_ID  statement
        {
            mStmtSelection = new(StmtSelection)
            mStmtSelection.Expr = $3
            mStmtSelection.Stmt = $5
            mStmtSelection.Stmt2 = $7
            $$ = mStmtSelection
            rlog(1, "IF_ID '(' expression ')' statement ELSE_ID  statement\n")
            //DumpStmtSelection(mStmtSelection,1)
        }
    ;





iteration_statement
    : WHILE_ID '(' expression ')' statement
        {
            mStmtWhile = new(StmtWhile)
            mStmtWhile.Expr = $3
            mStmtWhile.Stmt = $5

            mStmtIteration = new(StmtIteration)
            mStmtIteration.StmtWhile = mStmtWhile
            $$ = mStmtIteration
            //DumpStmtWhile(mStmtWhile,2)
        }
    | DO_ID statement WHILE_ID '(' expression ')' ';'
        {
            mStmtDo = new(StmtDo)
            mStmtDo.Expr = $5
            mStmtDo.Stmt = $2

            mStmtIteration = new(StmtIteration)
            mStmtIteration.StmtDo = mStmtDo
            $$ = mStmtIteration
        }
    | FOR_ID '(' expression_statement expression_statement ')' statement
        {
            mStmtFor = new(StmtFor)
            mStmtFor.ExprStmt1 = $3
            mStmtFor.ExprStmt2 = $4
            mStmtFor.Stmt = $6

            mStmtIteration = new(StmtIteration)
            mStmtIteration.StmtFor = mStmtFor
            $$ = mStmtIteration
        }
    | FOR_ID '(' expression_statement expression_statement expression ')' statement    
        {
            mStmtFor = new(StmtFor)
            mStmtFor.ExprStmt1 = $3
            mStmtFor.ExprStmt2 = $4
            mStmtFor.Expr = $5
            mStmtFor.Stmt = $7

            mStmtIteration = new(StmtIteration)
            mStmtIteration.StmtFor = mStmtFor
            $$ = mStmtIteration
        }
    | STEP_ID IDENTIFIER '(' expression ')' STIMULUS_ID compound_statement RESPONSE_ID compound_statement END_STEP_ID
        {
            mStmtStep = new(StmtStep)
            mStmtStep.Expr = $4
            mStmtStep.CmpndStmtStimulus = $7
            mStmtStep.CmpndStmtResponse = $9
            mIdentifier = new(Identifier)
            mIdentifier.Name = $2
            mStmtStep.Identifier = mIdentifier

            mStmtIteration = new(StmtIteration)
            mStmtIteration.StmtStep = mStmtStep
            $$ = mStmtIteration
        }

    | STEP_ID IDENTIFIER  STIMULUS_ID compound_statement RESPONSE_ID compound_statement END_STEP_ID
        {
            mStmtStep = new(StmtStep)
            mStmtStep.CmpndStmtStimulus = $4
            mStmtStep.CmpndStmtResponse = $6
            mIdentifier = new(Identifier)
            mIdentifier.Name = $2
            mStmtStep.Identifier = mIdentifier

            mStmtIteration = new(StmtIteration)
            mStmtIteration.StmtStep = mStmtStep
            $$ = mStmtIteration
        }

    ;

declaration_list
    : declaration
        {
            mDeclarationList = new(DeclarationList)
            mDeclarationList.Declaration = $1
            $$ = mDeclarationList
            rlog2("declaration_list>>declaration\n")
        }
    | declaration_list declaration
        {
            mDeclarationList = new(DeclarationList)
            mDeclarationList.Declaration = $2
            mDeclarationList.DeclarationList = $1
            $$ = mDeclarationList
            rlog2("declaration_list>>declaration_list declaration\n")
        }
    ;

declaration
    : declaration_specifiers ';'
        {
            mDeclaration = new(Declaration)
            mDeclaration.SpecDecl = $1
            $$ = mDeclaration
            rlog2("declaration_specifiers ';'\n")
            //DumpDeclaration(mDeclaration,3)
        }
    | declaration_specifiers init_declarator_list ';'
        {
            mDeclaration = new(Declaration)
            mDeclaration.SpecDecl = $1
            mDeclaration.DeclInitList = $2
            $$ = mDeclaration
            rlog2("declaration_specifiers init_declarator_list ';'\n", $$, $1, $2)
            //DumpDeclaration(mDeclaration,3)
        }
    ;

declaration_specifiers
    : type_specifier
        {
            mSpecDecl = new(SpecDecl)
            mSpecDecl.SpecType = $1
            $$ = mSpecDecl
            rlog2("type_specifier $1=%v $$=%v $$.SpecType=%v \n", $1, $$, $$.SpecType)
        }
    ;

type_specifier
    : TYPE_INT_ID
        {
            mSpecType = new(SpecType)
            mSpecType.Id = TYPE_INT_ID
            mSpecType.Name = "int"
            $$ = mSpecType
        }
    | TYPE_STRING_ID
        {
            mSpecType = new(SpecType)
            mSpecType.Id = LITERAL_STRING
            mSpecType.Name = "string"
            $$ = mSpecType
        }
    | TYPE_FLOAT_ID
        {
            mSpecType = new(SpecType)
            mSpecType.Id = TYPE_FLOAT_ID
            mSpecType.Name = "float"
            $$ = mSpecType
        }
    | TYPE_BOOL_ID
        {
            mSpecType = new(SpecType)
            mSpecType.Id = TYPE_BOOL_ID
            mSpecType.Name = "bool"
            $$ = mSpecType
        }
    | TYPE_TESTPLAN_ID
        {
            mSpecType = new(SpecType)
            mSpecType.Id = TYPE_TESTPLAN_ID
            mSpecType.Name = "testplan"
            $$ = mSpecType
        }

    ;

init_declarator_list
    : init_declarator
        {
            mDeclInitList = new (DeclInitList)
            mDeclInitList.DeclInit = $1
            $$ = mDeclInitList
            rlog2("init_declarator_list>init_declarator\n", mDeclInitList, $1)
                    
        }
    | init_declarator_list ',' init_declarator
        {
            mDeclInitList = new (DeclInitList)
            mDeclInitList.DeclInit = $3
            mDeclInitList.DeclInitList = $1
            $$ = mDeclInitList
            rlog2("init_declarator_list ',' init_declarator\n", mDeclInitList, $1)
        }
    ;

init_declarator
    : declarator
        {
            mDeclInit = new(DeclInit)
            mDeclInit.Declarator = $1
            $$ = mDeclInit
            rlog2("declarator 2\n")
        }
    | declarator '=' initializer
        {
            mDeclInit = new(DeclInit)
            mDeclInit.Declarator = $1
            mDeclInit.Initializer = $3
            rlog2("declarator '=' initializer mDeclInit=%v mDeclarator=%v mInitializer=%v\n", mDeclInit, mDeclarator, mInitializer)
            $$ = mDeclInit
            //DumpDeclInit(mDeclInit,3)
            //DumpInitializer($3, 3)
            //DumpDeclarator($1, 3)
        }
    ;

initializer
    : assignment_expression
        {
            mInitializer = new(Initializer)
            mInitializer.ExprAssign = mExprAssign
            mInitializer.ExprUnary = mExprUnary
            mInitializer.ExprCmd = mExprCmd
            rlog2("assignment_expression: mInitializer=%v mExprUnary=%v mExprAssign=%v mExprCmd=%v \n",mInitializer,mExprUnary,mExprAssign, mExprCmd)
            //DumpExprAssign(mExprAssign, 4)
            $$ = mInitializer
        }
    | '{' initializer_list '}'
        {
            mInitializer = new(Initializer)
            mInitializer.InitializerList = $2
            $$ = mInitializer
            rlog2("'{' initializer_list '}'\n")
        }
    | '{' initializer_list ',' '}'
        {
            mInitializer = new(Initializer)
            mInitializer.InitializerList = $2
            $$ = mInitializer
            rlog2("'{' initializer_list ',' '}'\n")
        }
    ;

initializer_list
    : initializer
        {
            mInitializerList = new(InitializerList)
            mInitializerList.Initializer = $1
            $$ = mInitializerList
            rlog2("initializer\n")
        }
    | initializer_list ',' initializer
        {
            mInitializerList = new(InitializerList)
            mInitializerList.Initializer = $3
            mInitializerList.InitializerList = $1
            $$ = mInitializerList
            rlog2("initializer_list ',' initializer\n")
        }
    ;

expression_statement
    : ';'
        {
            mExprStmt = new(ExprStmt)
            $$ = mExprStmt
        }
    | expression ';'
        {
            mExprStmt = new(ExprStmt)
            mExprStmt.Expr = $1
            $$ = mExprStmt
            //rlog2("expression\n")
        }

    ;

expression
    : assignment_expression
        {
            mExpr = new(Expr)
            mExpr.ExprAssign = $1
            $$ = mExpr
        }
    | expression ',' assignment_expression
        {
            mExpr = new(Expr)
            mExpr.ExprAssign = $3
            mExpr.Expr = $1
            $$ = mExpr
            //rlog2("expression ',' assignment_expression\n")
        }
    ;

assignment_expression
    : conditional_expression
        {
           mExprAssign = new(ExprAssign)
           mExprAssign.ExprConditional = $1
           $$ = mExprAssign
           rlog2("conditional_expression mExprAssign=%v mExprCmd=%v mExprPostfix=%v mExprUnary=%v \n",mExprAssign, mExprCmd, mExprPostfix, mExprUnary)
           //DumpExprConditional($1, 3)
        }
    | unary_expression assignment_operator assignment_expression
        {
            mExprAssign = new(ExprAssign)
            mExprAssign.ExprUnary = $1
            mExprAssign.ExprAssign = $3
            $$ = mExprAssign
            rlog2("unary_expression assignment_operator assignment_expression >> mExprAssign=%v mExprCmd=%v mExprPostfix=%v mExprUnary=%v \n",mExprAssign, mExprCmd, mExprPostfix, mExprUnary)
        }
    | command_expression
        {
            mExprAssign = new(ExprAssign)
            mExprAssign.ExprCmd = $1
            //rlog2("command_expression 2 mExprCmd=%v mExprCmd.CmdListApp=%v \n", mExprCmd, mExprCmd.CmdListApp)
            $$ = mExprAssign
        }
    ;

assignment_operator
    : '='
    | MUL_ASSIGN
    | DIV_ASSIGN
    | MOD_ASSIGN
    | ADD_ASSIGN
    | SUB_ASSIGN
    | AND_ASSIGN
    | XOR_ASSIGN
    | OR_ASSIGN
    ;

conditional_expression
    : logical_or_expression
        {
            mExprConditional = new(ExprConditional)
            mExprConditional.ExprLogicOr = $1
            $$ = mExprConditional
            rlog2("logical_or_expression")
        }
    ;

logical_or_expression
    : logical_and_expression
        {
            mExprLogicOr = new(ExprLogicOr)
            mExprLogicOr.ExprLogicAnd = $1
            $$ = mExprLogicOr
            rlog2("logical_and_expression\n")
        }
     | logical_or_expression OR_OP logical_and_expression
         {
            mExprLogicOr = new(ExprLogicOr)
            mExprLogicOr.ExprLogicAnd = $3
            mExprLogicOr.ExprLogicOr = $1
            $$ = mExprLogicOr
            rlog2("logical_or_expression OR_OP logical_and_expression\n")
        }
    ;

logical_and_expression
    : inclusive_or_expression
        {
            mExprLogicAnd = new(ExprLogicAnd)
            mExprLogicAnd.ExprInclOr = $1
            $$ = mExprLogicAnd
            rlog2("inclusive_or_expression")
        }
    | logical_and_expression AND_OP inclusive_or_expression
        {
            mExprLogicAnd = new(ExprLogicAnd)
            mExprLogicAnd.ExprInclOr = $3
            mExprLogicAnd.ExprLogicAnd = $1
            $$ = mExprLogicAnd
            rlog2("logical_and_expression AND_OP inclusive_or_expression")
        }
    ;

inclusive_or_expression
    : exclusive_or_expression
        {
            mExprInclOr = new(ExprInclOr)
            mExprInclOr.ExprExclOr = $1
            $$ = mExprInclOr
            rlog2("exclusive_or_expression\n")
        }
    | inclusive_or_expression '|' exclusive_or_expression
        {
            mExprInclOr = new(ExprInclOr)
            mExprInclOr.ExprExclOr = $3
            mExprInclOr.ExprInclOr = $1
            $$ = mExprInclOr
            rlog2("exclusive_or_expression\n")
        }
    ;

exclusive_or_expression
    : and_expression
        {
            mExprExclOr = new(ExprExclOr)
            mExprExclOr.ExprAnd = $1
            $$ = mExprExclOr
            rlog2("and_expression\n")
        }
    | exclusive_or_expression '^' and_expression
        {
            mExprExclOr = new(ExprExclOr)
            mExprExclOr.ExprAnd = $3
            mExprExclOr.ExprExclOr = $1
            $$ = mExprExclOr
            rlog2("exclusive_or_expression '^' and_expression\n")
        }
    ;

and_expression
    : equality_expression
        {
            mExprAnd = new(ExprAnd)
            mExprAnd.ExprEquality = $1
            $$ = mExprAnd
            rlog2("equality_expression\n")
        }
    | and_expression '&' equality_expression
        {
            mExprAnd = new(ExprAnd)
            mExprAnd.ExprEquality = $3
            mExprAnd.ExprAnd = $1
            $$ = mExprAnd
            rlog2("and_expression '&' equality_expression\n")
        }
    ;


equality_expression
    : relational_expression
        {
            mExprEquality = new(ExprEquality)
            mExprEquality.ExprRelational = $1
            $$ = mExprEquality
            rlog2("relational_expression\n")
        }
    | equality_expression EQ_OP relational_expression
        {
            mExprEquality = new(ExprEquality)
            mExprEquality.ExprRelational = $3
            mExprEquality.ExprEquality = $1
            mExprEquality.Sign = "EQ"
            $$ = mExprEquality
            rlog2("equality_expression EQ_OP relational_expression\n")
        }
    | equality_expression NE_OP relational_expression
        {
            mExprEquality = new(ExprEquality)
            mExprEquality.ExprRelational = $3
            mExprEquality.ExprEquality = $1
            mExprEquality.Sign = "NE"
            $$ = mExprEquality
            rlog2("equality_expression NE_OP relational_expression\n")
        }
    ;

relational_expression
    : shift_expression
        {
            mExprRelational = new(ExprRelational)
            mExprRelational.ExprShift = $1
            $$ = mExprRelational
        }
    | relational_expression '<' shift_expression
        {
            mExprRelational = new(ExprRelational)
            mExprRelational.ExprShift = $3
            mExprRelational.ExprRelational = $1
            mExprRelational.Sign = "<"
            $$ = mExprRelational
            rlog2("relational_expression '<' shift_expression\n")
        }
    | relational_expression '>' shift_expression
        {
            mExprRelational = new(ExprRelational)
            mExprRelational.ExprShift = $3
            mExprRelational.ExprRelational = $1
            mExprRelational.Sign = ">"
            $$ = mExprRelational
            rlog2("relational_expression '>' shift_expression\n")
        }
    | relational_expression LE_OP shift_expression
        {
            mExprRelational = new(ExprRelational)
            mExprRelational.ExprShift = $3
            mExprRelational.ExprRelational = $1
            mExprRelational.Sign = "LE"
            $$ = mExprRelational
            rlog2("relational_expression LE_OP shift_expression\n")
        }
    | relational_expression GE_OP shift_expression
        {
            mExprRelational = new(ExprRelational)
            mExprRelational.ExprShift = $3
            mExprRelational.ExprRelational = $1
            mExprRelational.Sign = "GE"
            $$ = mExprRelational
            rlog2("relational_expression GE_OP shift_expression\n")
        }
    ;

shift_expression
    : additive_expression
        {
            mExprShift = new(ExprShift)
            mExprShift.ExprAdditive = $1
            $$ = mExprShift
            rlog2("additive_expression\n")
        }
    | shift_expression LEFT_OP additive_expression
        {
            mExprShift = new(ExprShift)
            mExprShift.ExprAdditive = $3
            mExprShift.ExprShift = $1
            $$ = mExprShift
            rlog2("shift_expression LEFT_OP additive_expression\n")
        }
    | shift_expression RIGHT_OP additive_expression
        {
            mExprShift = new(ExprShift)
            mExprShift.ExprAdditive = $3
            mExprShift.ExprShift = $1
            $$ = mExprShift
            rlog2("shift_expression RIGHT_OP additive_expression\n")
        }
    ;

additive_expression
    : multiplicative_expression
        {
            mExprAdditive = new(ExprAdditive)
            mExprAdditive.ExprMultiplicative = mExprMultiplicative
            rlog2("additive_expression >> multiplicative_expression %v %v \n", mExprAdditive, mExprMultiplicative)
            $$ = mExprAdditive
        }
    | additive_expression '+' multiplicative_expression
        {
            mExprAdditive = new(ExprAdditive)
            mExprAdditive.ExprAdditive = $1
            mExprAdditive.ExprMultiplicative = $3
            mExprAdditive.Sign = "+"
            $$ = mExprAdditive
            rlog2("additive_expression '+' multiplicative_expression\n")
        }
    | additive_expression '-' multiplicative_expression
        {
            mExprAdditive = new(ExprAdditive)
            mExprAdditive.ExprAdditive = $1
            mExprAdditive.ExprMultiplicative = $3
            mExprAdditive.Sign = "-"
            $$ = mExprAdditive
            rlog2("additive_expression '-' multiplicative_expression\n")
        }
    ;

multiplicative_expression
    : cast_expression
        {
            mExprMultiplicative = new(ExprMultiplicative)
            mExprMultiplicative.ExprCast = $1
            rlog2("multiplicative_expression >> cast_expression mExprMultiplicative=%v mExprCast=%v \n", mExprMultiplicative, mExprCast)
            $$ = mExprMultiplicative
        }
    | multiplicative_expression '*' cast_expression
        {
            mExprMultiplicative = new(ExprMultiplicative)
            mExprMultiplicative.ExprCast = $3
            mExprMultiplicative.ExprMultiplicative = $1
            mExprMultiplicative.Sign = "*"
            $$ = mExprMultiplicative
            rlog2("multiplicative_expression '*' cast_expression\n")
        }
    | multiplicative_expression '/' cast_expression
        {
            mExprMultiplicative = new(ExprMultiplicative)
            mExprMultiplicative.ExprCast = $3
            mExprMultiplicative.ExprMultiplicative = $1
            mExprMultiplicative.Sign = "/"
            $$ = mExprMultiplicative
            rlog2("multiplicative_expression '/' cast_expression\n")
        }
    | multiplicative_expression '%' cast_expression
        {
            mExprMultiplicative = new(ExprMultiplicative)
            mExprMultiplicative.ExprCast = $3
            mExprMultiplicative.ExprMultiplicative = $1
            mExprMultiplicative.Sign = "%"
            $$ = mExprMultiplicative
            rlog2("multiplicative_expression '%' cast_expression\n")
        }
    ;

cast_expression
    : unary_expression
        {
            mExprCast = new(ExprCast)
            mExprCast.ExprUnary = $1
            rlog2("cast_expression >> unary_expression mExprCast=%v mExprUnary=%v \n", mExprCast, mExprUnary)
            $$ = mExprCast
        }
    ;

unary_operator
    : '&'
        {
            mUnaryOperator = new(UnaryOperator)
            mUnaryOperator.Sign = "&"
            $$ = mUnaryOperator
        }
    | '*'
        {
            mUnaryOperator = new(UnaryOperator)
            mUnaryOperator.Sign = "*"
            $$ = mUnaryOperator
        }
    | '+'
        {
            mUnaryOperator = new(UnaryOperator)
            mUnaryOperator.Sign = "+"
            $$ = mUnaryOperator
        }
    | '-'
        {
            mUnaryOperator = new(UnaryOperator)
            mUnaryOperator.Sign = "-"
            $$ = mUnaryOperator
        }
    | '~'
        {
            mUnaryOperator = new(UnaryOperator)
            mUnaryOperator.Sign = "~"
            $$ = mUnaryOperator
        }
    | '!'
        {
            mUnaryOperator = new(UnaryOperator)
            mUnaryOperator.Sign = "!"
            $$ = mUnaryOperator
        }
    ;

unary_expression
    : postfix_expression
        {
            mExprUnary = new(ExprUnary)
            mExprUnary.ExprPostfix = $1
            mExprCmd = nil
            rlog2("unary_expression:postfix_expression\n", mExprPostfix, mExprUnary)
            $$ = mExprUnary
        }
    | INC_OP unary_expression
        {
            mUnaryOperator = new(UnaryOperator)
            mUnaryOperator.Sign = "++"

            mExprUnary = new(ExprUnary)
            mExprUnary.UnaryOperator =  mUnaryOperator
            $$ = mExprUnary
            rlog2("unary_expression:INC_OP postfix_expression\n")
        }
    | DEC_OP unary_expression
        {
            mUnaryOperator = new(UnaryOperator)
            mUnaryOperator.Sign = "--"

            mExprUnary = new(ExprUnary)
            mExprUnary.UnaryOperator =  mUnaryOperator
            $$ = mExprUnary
            rlog2("unary_expression:DEC_OP postfix_expression\n")
        }
    | unary_operator cast_expression
        {
            mExprUnary = new(ExprUnary)
            mExprUnary.UnaryOperator =  $1
            mExprUnary.ExprCast = $2
            $$ = mExprUnary
            rlog2("unary_expression cast_expression:postfix_expression\n")
        }
    ;

primary_expression
    : IDENTIFIER
        {
            mExprPrimary = new(ExprPrimary)
            mExprPrimary.Value = $1
            mExprPrimary.IdType = IDENTIFIER
            rlog2("primary_expression:IDENTIFIER>>\n",mExprPrimary, $1)
            $$ = mExprPrimary
        }
    | '(' expression ')'
        {
            mExprPrimary = new(ExprPrimary)
            //mExprPrimary.Value = nil
            mExprPrimary.IdType = EXPR_ID
            mExprPrimary.ExprAssign = mExprAssign
            rlog2("primary_expression:EXPR_ID>> mExprPrimary=%v ExprAssign=%v \n",mExprPrimary, mExprPrimary.ExprAssign)
            $$ = mExprPrimary
        }
    | TRUE_ID
        {
            mExprPrimary = new(ExprPrimary)
            mExprPrimary.Value = $1
            mExprPrimary.IdType = TRUE_ID
            $$ = mExprPrimary
            rlog2("primary_expression:TRUE_ID>>\n",mExprPrimary, $1)
        }
    | FALSE_ID
        {
            mExprPrimary = new(ExprPrimary)
            mExprPrimary.Value = $1
            mExprPrimary.IdType = FALSE_ID
            $$ = mExprPrimary
            rlog2("primary_expression:FALSE_ID>>\n",mExprPrimary, $1)
        }
    | FLOAT_ID
        {
            mExprPrimary = new(ExprPrimary)
            mExprPrimary.Value = $1
            mExprPrimary.IdType = FLOAT_ID
            rlog2("primary_expression:FLOAT_ID>>\n",mExprPrimary, $1)
            $$ = mExprPrimary
        }
    | LITERAL_STRING
        {
            mExprPrimary = new(ExprPrimary)
            mExprPrimary.Value = $1
            mExprPrimary.IdType = LITERAL_STRING
            rlog2("primary_expression:LITERAL_STRING>>\n",mExprPrimary, $1)
            $$ = mExprPrimary
        }
    | INTEGER_ID
        {
            mExprPrimary = new(ExprPrimary)
            mExprPrimary.Value = $1
            mExprPrimary.IdType = INTEGER_ID
            rlog2("primary_expression:INTEGER_ID>>\n",mExprPrimary, $1)
            $$ = mExprPrimary
        }
    | EMAIL_ID
        {
            mExprPrimary = new(ExprPrimary)
            mExprPrimary.Value = $1
            mExprPrimary.IdType = EMAIL_ID
            rlog2("primary_expression:EMAIL_ID>>\n",mExprPrimary, $1)
            $$ = mExprPrimary
        }
    | UUID_ID
        {
            mExprPrimary = new(ExprPrimary)
            mExprPrimary.Value = $1
            mExprPrimary.IdType = UUID_ID
            rlog2("primary_expression:UUID_ID>>\n",mExprPrimary, $1)
            $$ = mExprPrimary
        }
    ;

postfix_expression
    : primary_expression
        {
            mExprPostfix = new(ExprPostfix)
            mExprPostfix.ExprPrimary = $1
            mExprPostfix.Operator = ""
            rlog2("primary_expression: ", mExprPostfix, mExprPostfix.ExprPrimary)
            $$ = mExprPostfix
        }
    | postfix_expression '[' expression ']'
        {
            mExprPostfix = new(ExprPostfix)
            mExprPostfix.ExprPostfix = $1
            mExprPostfix.Expr = $3
            mExprPostfix.Operator = ""
            mExprPostfix.Modifier = "[]"
            $$ = mExprPostfix
            rlog2("postfix_expression '[' ']'\n")
        }
    | postfix_expression '(' ')'
        {
            mExprPostfix = new(ExprPostfix)
            mExprPostfix.ExprPostfix = $1
            mExprPostfix.Operator = ""
            mExprPostfix.Modifier = "()"
            $$ = mExprPostfix
            rlog2("postfix_expression '(' ')'\n")
        }
    | postfix_expression '(' argument_expression_list ')'
        {
            mExprPostfix = new(ExprPostfix)
            mExprPostfix.ExprPostfix = $1
            mExprPostfix.ExprListArgument = $3
            mExprPostfix.Operator = ""
            mExprPostfix.Modifier = "()"
            $$ = mExprPostfix
            rlog2("postfix_expression '(' argument_expression_list ')'\n")
        }
    | postfix_expression '.' IDENTIFIER
        {
            mExprPostfix = new(ExprPostfix)
            mExprPostfix.ExprPostfix = $1
            mExprPostfix.Identifier = new(Identifier)
            mExprPostfix.Identifier.Name = $3
            mExprPostfix.Operator = "."
            $$ = mExprPostfix
            rlog2("postfix_expression '.' IDENTIFIER\n", $$, $3)
        }
    | postfix_expression PTR_OP IDENTIFIER
        {
            mExprPostfix = new(ExprPostfix)
            mExprPostfix.ExprPostfix = $1
            mExprPostfix.Identifier = new(Identifier)
            mExprPostfix.Identifier.Name = $3
            mExprPostfix.Operator = "PTR_OP"
            $$ = mExprPostfix
            rlog2("postfix_expression PTR_OP IDENTIFIER\n")
        }
    | postfix_expression INC_OP
        {
            mExprPostfix = new(ExprPostfix)
            mExprPostfix.ExprPostfix = $1
            mExprPostfix.Identifier = new(Identifier)
            mExprPostfix.Operator = "INC_OP"
            $$ = mExprPostfix
            rlog2("postfix_expression INC_OP\n")
        }
    | postfix_expression DEC_OP
        {
            mExprPostfix = new(ExprPostfix)
            mExprPostfix.ExprPostfix = $1
            mExprPostfix.Identifier = new(Identifier)
            mExprPostfix.Operator = "DEC_OP"
            $$ = mExprPostfix
            rlog2("postfix_expression DEC_OP\n")
        }
    ;

argument_expression_list
    : assignment_expression
        {
            mExprListArgument = new(ExprListArgument)
            mExprListArgument.ExprAssign = $1
            $$ = mExprListArgument
            rlog2("argument_expression_list>>assignment_expression\n")
        }
    | argument_expression_list ',' assignment_expression
        {
            mExprListArgument = new(ExprListArgument)
            mExprListArgument.ExprListArgument = $1
            mExprListArgument.ExprAssign = $3
            $$ = mExprListArgument
            rlog2("argument_expression_list ',' assignment_expression\n")
        }
    ;


command_expression
    : list_application_command 
        {
            mExprCmd = new(ExprCmd)
            mExprCmd.CmdListApp = $1
            //rlog2("list_application_command:",mCmdListApp, mExprCmd)
            mCmdMakeAppUser = nil
            mCmdListApp = nil
            mCmdVerify = nil
            mExprUnary = nil
            $$ = mExprCmd
        }
    | verify_application_url_category_command
        {
            mExprCmd = new(ExprCmd)
            mExprCmd.CmdVerify = mCmdVerify
            //rlog2("verify_application_url_category_command:\n",mCmdVerify, mExprCmd)
            mCmdMakeAppUser = nil
            mCmdListApp = nil
            mCmdVerify = nil
            mExprUnary = nil
            $$ = mExprCmd
        }
    | generate_data_command
        {
            mExprCmd = new(ExprCmd)
            mExprCmd.CmdGenData = $1
            //rlog2("generate_data_command\n",mCmdGenData,mExprCmd)
            mCmdMakeAppUser = nil
            mCmdListApp = nil
            mCmdVerify = nil
            mExprUnary = nil
            $$ = mExprCmd
        }
    | application_makeappuser_command
        {
            mExprCmd = new(ExprCmd)
            mExprCmd.CmdMakeAppUser = mCmdMakeAppUser
            //rlog2("application_makeappuser_command\n",mCmdMakeAppUser,mExprCmd)
            mCmdMakeAppUser = nil
            mCmdListApp = nil
            mCmdVerify = nil
            mExprUnary = nil
            $$ = mExprCmd
        }
    | application_makeapptid_command
        {
            mExprCmd = new(ExprCmd)
            mExprCmd.CmdMakeAppTid = $1
            //rlog2("application_makeapptid_command\n",mCmdMakeAppTid,mExprCmd)
            mCmdMakeAppUser = nil
            mCmdListApp = nil
            mCmdVerify = nil
            mExprUnary = nil
            $$ = mExprCmd
        }
    | test_route_command
        {
            mExprCmd = new(ExprCmd)
            mExprCmd.CmdTestRoute = $1
            $$ = mExprCmd
        }
    | create_application_command
        {
            mExprCmd = new(ExprCmd)
            mExprCmd.CmdCreateApp = $1
            $$ = mExprCmd
        }
    | delete_application_command
        {
            mExprCmd = new(ExprCmd)
            mExprCmd.CmdDeleteApp = $1
            $$ = mExprCmd
        }
    | delete_application_tid_command
        {
            mExprCmd = new(ExprCmd)
            mExprCmd.CmdDeleteAppTid = $1
            $$ = mExprCmd
        }
    | run_sql_command
        {
            mExprCmd = new(ExprCmd)
            mExprCmd.CmdRunSQL = $1
            $$ = mExprCmd
        }
    | delete_application_user_command
        {
            mExprCmd = new(ExprCmd)
            mExprCmd.CmdDeleteAppUser = $1
            $$ = mExprCmd
        }
    ;

application_makeappuser_command
    : MAKEAPPUSER_ID expression IN_ID '(' expression ')'
        {
            mCmdMakeAppUser = new(CmdMakeAppUser)
            mCmdMakeAppUser.Name = "MakeAppUser"
            mCmdMakeAppUser.Expr1 = $2
            mCmdMakeAppUser.Expr2 = $5
            rlog2("MAKEAPPUSER_ID \n",mCmdMakeAppUser)
            $$ = mCmdMakeAppUser
        }
    ;

generate_data_command
    : GENDATA_ID expression LOOP_ID '(' expression ')'
        {
            rlog2("GENDATA_ID expression ';'\n")
            mCmdGenData = new(CmdGenData)
            mCmdGenData.Name = "GENDATA"
            mCmdGenData.ExprApp = $2
            mCmdGenData.ExprLoop = $5
            $$ = mCmdGenData
        }
    ;

verify_application_url_category_command
    : VERIFY_ID expression URLID_ID expression WITH_ID '('  expression ')'
        {
            mCmdVerify = new(CmdVerify)
            mCmdVerify.Name = "verify"
            rlog2("VERIFY_ID \n",mCmdVerify)
            $$ = mCmdVerify
        }
    ;

list_application_command
    : LIST_ID APPS_ID
        {
            mCmdListApp = new(CmdListApp) 
            mCmdListApp.Name = "List Apps"
            rlog2("LIST_ID APPS_ID\n",mCmdListApp)
            $$ = mCmdListApp
        }
    ;

application_makeapptid_command
    : MAKEAPPTID_ID expression IN_ID '(' expression ')'
        {
            mCmdMakeAppTid = new(CmdMakeAppTid) 
            mCmdMakeAppTid.Name = "MAKEAPPTID"
            mCmdMakeAppTid.Expr1 = $2
            mCmdMakeAppTid.Expr2 = $5
            rlog2("MAKEAPPTID_ID:\n",mCmdMakeAppTid)
            $$ = mCmdMakeAppTid
        }
    ;

test_route_command
    : TESTROUTE_ID expression IN_ID '(' expression ')'
        {
            mCmdTestRoute = new(CmdTestRoute) 
            mCmdTestRoute.Name = "TESTROUTE"
            mCmdTestRoute.Expr1 = $2
            mCmdTestRoute.Expr2 = $5
            //rlog2("test_route_command:\n",mCmdTestRoute)
            $$ = mCmdTestRoute
        }
    ;

create_application_command
    : CREATE_ID APP_ID '(' expression ')'
        {
            mCmdCreateApp = new(CmdCreateApp) 
            mCmdCreateApp.Name = "CREATEAPP"
            mCmdCreateApp.Expr1 = $4
            //rlog2("create_application_command:\n",mCmdCreateApp)
            $$ = mCmdCreateApp
        }
    ;

delete_application_command
    : DELETE_ID APP_ID '(' expression ')'
        {
            mCmdDeleteApp = new(CmdDeleteApp) 
            mCmdDeleteApp.Name = "DELETEAPP"
            mCmdDeleteApp.Expr1 = $4
            //rlog2("delete_application_command:\n",mCmdDeleteApp)
            $$ = mCmdDeleteApp
        }
    ;

delete_application_tid_command
    : DELETE_ID APP_ID TID_ID expression  IN_ID '(' expression ')'
        {
            mCmdDeleteAppTid = new(CmdDeleteAppTid) 
            mCmdDeleteAppTid.Name = "DELETEAPPTID"
            mCmdDeleteAppTid.Expr1 = $7
            mCmdDeleteAppTid.Expr2 = $4
            //rlog2("delete_application_command:\n",mCmdDeleteAppTid)
            $$ = mCmdDeleteAppTid
        }
    ;

run_sql_command
    : RUNSQL_ID '(' expression ')' IN_ID '(' expression ')'
        {
            mCmdRunSQL = new(CmdRunSQL) 
            mCmdRunSQL.Name = "RUNSQL"
            mCmdRunSQL.Expr1 = $3
            mCmdRunSQL.Expr2 = $7
            //rlog2("delete_application_command:\n",mCmdRunSQL)
            $$ = mCmdRunSQL
        }
    ;

delete_application_user_command
    : DELETE_APPUSER_ID '(' expression ')' IN_ID '(' expression ')'
        {
            mCmdDeleteAppUser = new(CmdDeleteAppUser) 
            mCmdDeleteAppUser.Name = "DELETE_APPUSER_ID"
            mCmdDeleteAppUser.Expr1 = $3
            mCmdDeleteAppUser.Expr2 = $7
            //rlog2("delete_application_command:\n",mCmdDeleteAppUser)
            $$ = mCmdDeleteAppUser
        }
    ;

%%      /*  start  of  programs  */


