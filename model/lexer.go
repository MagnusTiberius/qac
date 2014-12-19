package model

import (
	"fmt"
	"strings"
    "unicode/utf8"
    //"bufio"
    //"database/sql"
    //"intel/test/selectorproperty"
    //"intel/test/parser1"
    //"os"
    //"io"
    //"bytes"
    //"net/http"
    //"net/url"
    //"time"
    //"io/ioutil" 
    //"intel/test/datagen2"
    //"intel/scriptgen2/model"
    //_ "github.com/lib/pq"
    //"strconv"
    //"intel/test/selectorproperty"
    //"net/http"
    //"net/url"
    //"bufio"
    //"os"
    //"io"
    //"bytes"
    //"unicode/utf8"
    //"strings"
    //"database/sql"
    //"io/ioutil"    
    //"time"
    //"intel/scriptgen2/model"
    //"intel/test/datagen2"
    //"intel/test/parser1"    
)




// ----------------------------------------------------------------------------

type PrvLexer struct {
    s       string
    pos     int
    width   int
    start   int
}

func (l *PrvLexer) Error(s string) {
    fmt.Printf("%s\nLine: %d\n", s, lineCounter)
    panic("Error:")
}

func (l *PrvLexer) peek() int32 {
    rune := l.next()
    l.backup()
    return rune
}

func (l *PrvLexer) backup() {
    l.pos -= l.width
}

func (l *PrvLexer) ignore() {
    l.start = l.pos
}

func (l *PrvLexer) next() (rune rune) {
    if l.pos >= len(l.s) {
        l.width = 0
        return 0
    }
    rune, l.width =
        utf8.DecodeRuneInString(l.s[l.pos:])
    l.pos += l.width
    return rune
}

func (l *PrvLexer) accept(valid string) bool {
    if strings.IndexRune(valid, l.next()) >= 0 {
        return true
    }
    l.backup()
    return false
}

func (l *PrvLexer) acceptRun(valid string) {
    for strings.IndexRune(valid, l.next()) >= 0 {
    }
    l.backup()
}

// ----------------------------------------------------------------------------

var yylex           yyLexer
var eqn             string
var currState       int
var commandState    int
var prvLexer        *PrvLexer


func (l *PrvLexer) Lex(lval *yySymType) int {
    //fmt.Printf("1. l.pos:%v, val:%v  \n", l.pos, string(l.s[l.pos]))
    var c rune = ' '
    for c == ' ' || c == '\r'  {
        if l.pos == len(l.s) {
            return 0
        }
        c = rune(l.s[l.pos])
        l.pos += 1
        /*
        if l.pos == len(l.s) && c == '\n' {
            return 0
        }
        if l.pos < len(l.s) && c == '\n' {
            c = rune(l.s[l.pos])
            l.pos += 1
        }
        */
    }    


    /*
    if strings.IndexRune("/",c) >= 0  && strings.IndexRune("*",rune(l.s[l.pos+0])) >= 0 {
        var ok bool = false
        beg := l.pos
        cur := l.pos+1
        d := rune(l.s[cur])
        for d != '*' && cur < len(l.s) {
            cur += 1
            d = rune(l.s[cur])
        }
        if d == '*' {
            cur += 1
            d = rune(l.s[cur])
            if d == '/' {
                l.pos = cur+1
                ok = true
            } 
        }
        if ok {
            currState = COMMENT_SECTION_ID
            //commandState = COMMENT_SECTION_ID
            fmt.Println("COMMENT_SECTION_ID")
            return COMMENT_SECTION_ID
        } else {
            l.pos = beg
        }
    }

    if strings.IndexRune("'",c) >= 0  {
        var ok bool = false
        beg := l.pos
        cur := l.pos+1
        d := rune(l.s[cur])
        for d != '\'' && cur < len(l.s) {
            cur += 1
            d = rune(l.s[cur])
        }
        if d == '\'' {
            l.pos = cur+1
            ok = true
        }
        if ok {
            lval.id = l.s[beg:cur]
            currState = STRING_CONST_ID
            //commandState = STRING_CONST_ID
            return STRING_CONST_ID
        } else {
            l.pos = beg
        }
    }

    if strings.IndexRune("/",c) >= 0  && strings.IndexRune("/",rune(l.s[l.pos+0])) >= 0 {
        var ok bool = false
        beg := l.pos
        cur := l.pos+1
        d := rune(l.s[cur])
        for (d != '\r' && d != '\n' ) && cur < len(l.s) {
            cur += 1
            d = rune(l.s[cur])
        }
        if d != '\r' || d != '\n'  {
            l.pos = cur+1
            ok = true
        }
        if ok {
            currState = COMMENT_SECTION_ID
            //commandState = COMMENT_SECTION_ID
            return COMMENT_SECTION_ID
        } else {
            l.pos = beg
        }
    }



    */

    //fmt.Println(".....00...", string(c), "c" )

    for strings.IndexRune("\r\n\t ",c) >= 0 {
        if strings.IndexRune("\n",c) >= 0 {
            lineCounter++
        }
        if l.pos+1 == len(l.s) {
            return 0
        }
        c = rune(l.s[l.pos])
        l.pos += 1
    }

    // comment section /*     */ can detect repeating * within range, not able to detect nested though.
    savept := l.pos
    if strings.IndexRune("/",c) >= 0 {
        //fmt.Println("......com1....",string(c))
        if strings.IndexRune("*",l.next()) >= 0 {
            //fmt.Println("......com2....", string(c))
            cur := l.pos
            beg := cur
            ok := false
            d := rune(l.s[cur])
            for d != '/' && cur < len(l.s) {
                cur += 1
                d = rune(l.s[cur])
                //fmt.Println("......com3....", string(d))
            }
            cur -= 1
            d = rune(l.s[cur])
            //fmt.Println("......com4....", string(d))
            if d == '*' {
                //fmt.Println("......com5....", string(d))
                cur += 2
                ok = true
            }
            if ok {
                lval.id = l.s[beg-2:cur]
                //fmt.Println("......com6....", lval.id )
                l.pos = cur
                //fmt.Println("COMMENT_SECTION_ID\n", lval.id)
                return COMMENT_SECTION_ID
            } else {
                //fmt.Println("......com7....")
                l.pos = savept
            }
        } else {
            //fmt.Println("......com8....")
            l.pos = savept
        }
    }

    if strings.IndexRune(".",c) >= 0 {
        s1 := l.pos-1
        l.acceptRun(".")
        s2 := l.pos
        lval.id = l.s[s1:s2]
        //fmt.Println(".....x...", lval.id )
        if strings.IndexRune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",l.peek()) >= 0 {
            return '.'
        }
        if strings.IndexRune("./\\",l.peek()) >= 0 {
            l.acceptRun("/\\._ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
            s2 := l.pos
            lval.id = l.s[s1:s2]
            //fmt.Println(".....x2...", lval.id, "FILE_SPEC_ID" )
            return FILE_SPEC_ID
        }
    }

    savept = l.pos
    if strings.IndexRune("\"",c) >= 0 {
        s1 := l.pos
        n := len(l.s)
        for i:=l.pos;i<n;i++ {
            v := l.s[i]
            //fmt.Printf("%s \n", string(v))
            if string(v)=="\"" {
                l.pos = i+1
                s2 := l.pos-1
                lval.id = l.s[s1:s2]
                //fmt.Println(".....z2...", lval.id, "\t\t-->LITERAL_STRING" )
                return LITERAL_STRING
            }
        }
    }

    if strings.IndexRune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",c) >= 0 {
        s1 := l.pos-1
        l.acceptRun("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
        s2 := l.pos
        lval.id = l.s[s1:s2]
        //fmt.Println(".....1...", lval.id )
        if l.peek()==' ' || l.peek()==';' || l.pos==len(l.s)-1  || l.peek()=='(' || l.peek()==')' || l.peek()=='{' || l.peek()=='}' || l.peek()=='\n' {
            lval.id = l.s[s1:s2]
            //fmt.Println(".....2...", lval.id )
            if lval.id=="string" {
                //fmt.Println(".....1a...", lval.id, "-->TYPE_STRING_ID" )
                return TYPE_STRING_ID 
            }
            if lval.id=="int" {
                //fmt.Println(".....1b...", lval.id, "-->TYPE_INT_ID" )
                return TYPE_INT_ID 
            }
            if lval.id=="float" {
                //fmt.Println(".....1c...", lval.id, "-->TYPE_FLOAT_ID" )
                return TYPE_FLOAT_ID 
            }
            if lval.id=="bool" {
                //fmt.Println(".....1c...", lval.id, "-->BOOL_ID" )
                return TYPE_BOOL_ID 
            }
            if lval.id=="true" {
                //fmt.Println(".....1c...", lval.id, "-->TRUE_ID" )
                return TRUE_ID 
            }
            if lval.id=="false" {
                //fmt.Println(".....1c...", lval.id, "-->FALSE_ID" )
                return FALSE_ID 
            }
            if lval.id=="package" {
                //fmt.Println(".....1d...", lval.id, "-->PACKAGE_ID" )
                return PACKAGE_ID 
            }
            if lval.id=="import" {
                return IMPORT_ID 
            }
            if lval.id=="create" {
                return CREATE_ID 
            }
            if lval.id=="delete" {
                return DELETE_ID 
            }
            if lval.id=="app" {
                return APP_ID 
            }
            if lval.id=="list" {
                return LIST_ID 
            }
            if lval.id=="apps" {
                return APPS_ID 
            }
            if lval.id=="test" {
                return TEST_ID 
            }
            if lval.id=="hub" {
                return HUB_ID 
            }
            if lval.id=="drop" {
                return DROP_ID 
            }
            if lval.id=="generatedata" {
                return GENDATA_ID 
            }
            if lval.id=="for" {
                //fmt.Println(".....1j...", lval.id, "-->FOR_ID" )
                return FOR_ID 
            }
            if lval.id=="loop" {
                return LOOP_ID 
            }
            if lval.id=="populate" {
                return POPULATE_ID 
            }
            if lval.id=="from" {
                return FROM_ID 
            }
            if lval.id=="with" {
                return WITH_ID 
            }
            if lval.id=="verify" {
                return VERIFY_ID 
            }
            if lval.id=="testroute" {
                return TESTROUTE_ID 
            }
            if lval.id=="use" {
                return USE_ID 
            }
            if lval.id=="run" {
                return RUN_ID 
            }
            if lval.id=="style" {
                //fmt.Println(".....1k...", lval.id, "-->STYLE_ID" )
                return STYLE_ID 
            }
            if lval.id=="urlid" {
                return URLID_ID 
            }
            if lval.id=="in" {
                return IN_ID 
            }
            if lval.id=="if" {
                return IF_ID 
            }
            if lval.id=="endif" {
                return ENDIF_ID 
            }
            if lval.id=="else" {
                return ELSE_ID 
            }
            if lval.id=="return" {
                return RETURN_ID 
            }
            if lval.id=="func" {
                //fmt.Println(".....1r...", lval.id, "-->FUNC_ID" )
                return FUNC_ID 
            }
            if lval.id=="makeappuser" {
                return MAKEAPPUSER_ID 
            }
            if lval.id=="deleteappuser" {
                return DELETE_APPUSER_ID 
            }
            if lval.id=="makeapptid" {
                return MAKEAPPTID_ID 
            }
            if lval.id=="tid" {
                return TID_ID 
            }
            if lval.id=="load" {
                return LOAD_ID 
            }
            if lval.id=="while" {
                return WHILE_ID 
            }
            if lval.id=="do" {
                return DO_ID 
            }
            if lval.id=="testplan" {
                return TYPE_TESTPLAN_ID 
            }
            if lval.id=="verification" {
                return TYPE_TESTPLAN_ID 
            }
            if lval.id=="step" {
                return STEP_ID 
            }
            if lval.id=="endstep" {
                return END_STEP_ID 
            }
            if lval.id=="stimulus" {
                return STIMULUS_ID 
            }
            if lval.id=="response" {
                return RESPONSE_ID 
            }
            if lval.id=="runsql" {
                return RUNSQL_ID 
            }



            //fmt.Println(".....1k...", lval.id, "-->IDENTIFIER" )
            return IDENTIFIER
        }
        //fmt.Println(".....3...", l.s[s1:s2], " peek:", string(l.peek()))
        if strings.IndexRune("0123456789",l.next()) >= 0 {
            l.acceptRun("0123456789")
            s2 := l.pos
            //fmt.Println(".....3.2...", l.s[s1:s2], string(l.peek()))
            if l.peek()==' ' || l.pos==len(l.s)-1 || l.pos==len(l.s) || l.peek()=='{' || l.peek()=='}' || l.peek()=='(' || l.peek()==')' || l.peek()=='\n' {
                lval.id = l.s[s1:s2]
                //fmt.Println(".....4...", lval.id, "-->IDENTIFIER" )
                return IDENTIFIER
            }
        } else {
            l.backup()
        }

        savept = l.pos
        lval.id = l.s[s1:l.pos]
        //fmt.Println(".....4.1...", lval.id, " peek:",  string(l.peek()) )
        if strings.IndexRune(":",l.peek()) >= 0 {
            //fmt.Println(".....4a1...", lval.id, string(l.peek()) )
            l.acceptRun(":")
            lval.id = l.s[s1:l.pos]
            //fmt.Println(".....4a2...", lval.id, string(l.peek()) )
            if l.peek()==' ' || l.pos==len(l.s)-1 {
                lval.id = l.s[s1:l.pos]
                //fmt.Println(".....4a...", lval.id, "-->FILE_SPEC_ID" )
                return FILE_SPEC_ID
            }
            if strings.IndexRune("\\",l.peek()) >= 0 {
                l.acceptRun("\\._ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
                lval.id = l.s[s1:l.pos]
                if l.peek()==' ' || l.pos==len(l.s)-1 {
                    //fmt.Println(".....4a...", lval.id, "FILE_SPEC_ID" )
                    return FILE_SPEC_ID
                } else {
                    //fmt.Println(".....4a...", lval.id, "BAD_ID" )
                    return BAD_ID
                }    
            }
            if strings.IndexRune("\"",l.peek()) >= 0 {
                l.pos = savept
                lval.id = l.s[s1:l.pos]   
                //fmt.Println(".....4b...", lval.id, "-->IDENTIFIER" )
                return IDENTIFIER             
            }
        }

        if strings.IndexRune(",/;",l.peek()) >= 0 {
            //fmt.Println(".....4c...", lval.id, "-->IDENTIFIER" )
            return IDENTIFIER  
        }

        savept = l.pos
        if strings.IndexRune(".",l.next()) >= 0 {
            l.acceptRun(".")
            if l.peek()==' ' || l.pos==len(l.s)-1 {
                s2 := l.pos
                lval.id = l.s[s1:s2]
                //fmt.Println(".....5. bad_id..", lval.id, "-->BAD_ID" )
                return BAD_ID
            }
            lval.id = l.s[s1:l.pos]
            //fmt.Println(".....5a...", lval.id )
            if strings.IndexRune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",l.next()) >= 0 {
                l.acceptRun("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
                lval.id = l.s[s1:l.pos]
                //fmt.Println(".....5b...", lval.id )
                if l.peek()==' ' || l.pos==len(l.s)-1 {
                    l.pos = savept
                    s2 := l.pos
                    lval.id = l.s[s1:s2]
                    //fmt.Println(".....6...", lval.id, "IDENTIFIER" )
                    return IDENTIFIER
                }
                lval.id = l.s[s1:l.pos]
                //fmt.Println(".....6a...", lval.id, "*" )
                if strings.IndexRune("(",l.next()) >= 0 {
                    l.pos = savept
                    s2 := l.pos
                    lval.id = l.s[s1:s2]
                    //fmt.Println(".....6b...", lval.id, "IDENTIFIER" )
                    return IDENTIFIER
                }
                if strings.IndexRune("@.ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",l.next()) >= 0 {
                    l.acceptRun("@.ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
                    lval.id = l.s[s1:l.pos]
                    //fmt.Println(".....6a1...", lval.id, "*" )
                    if strings.IndexRune(" ;)}]",l.peek()) >= 0 || l.pos==len(l.s)-1 {
                        s2 := l.pos
                        lval.id = l.s[s1:s2]
                        //fmt.Println(".....6b...", lval.id, "EMAIL_ID" )
                        return EMAIL_ID
                    }
                }
            }
        } else {
            l.backup()
        }

    }

    if strings.IndexRune("0123456789",c) >= 0 {
        s1 := l.pos-1
        l.acceptRun("0123456789")
        s2 := l.pos
        
        if l.peek()==' ' || l.pos==len(l.s)-1 {
            lval.id = l.s[s1:s2]
            //fmt.Println(".....w1...", lval.id, "\t\t-->INTEGER_ID", s1, s2 )
            return INTEGER_ID
        }
        if strings.IndexRune("{}[]|&:,;",l.peek()) >= 0 {
            lval.id = l.s[s1:s2]
            //fmt.Println(".....w2...", lval.id, "\t\t-->INTEGER_ID", s1, s2 )
            return INTEGER_ID
        }
        if l.peek()=='.' {
            l.acceptRun(".")
            if strings.IndexRune("0123456789",l.peek()) >= 0 {
                l.acceptRun("0123456789")
                s2 = l.pos
                lval.id = l.s[s1:s2]
                //fmt.Println(".....w3...", lval.id, "\t\t-->FLOAT_ID", s1, s2 )
                return FLOAT_ID
            } else {
                return BAD_ID
            }
        }
    }

    if strings.IndexRune("[",c) >= 0 {
    }

    if strings.IndexRune("|",c) >= 0 {
        if l.peek()=='|' {
            s1 := l.pos-1
            l.acceptRun("|")
            s2 := l.pos
            lval.id = l.s[s1:s2]
            return OR_OP
        }
    }

    if strings.IndexRune("&",c) >= 0 {
        if l.peek()=='&' {
            s1 := l.pos-1
            l.acceptRun("&")
            s2 := l.pos
            lval.id = l.s[s1:s2]
            return AND_OP
        }
    }

    if strings.IndexRune("=",c) >= 0 {
        if l.peek()=='=' {
            s1 := l.pos-1
            l.acceptRun("=")
            s2 := l.pos
            lval.id = l.s[s1:s2]
            return EQ_OP
        }
    }

    if strings.IndexRune("!",c) >= 0 {
        if l.peek()=='=' {
            s1 := l.pos-1
            l.acceptRun("=")
            s2 := l.pos
            lval.id = l.s[s1:s2]
            return NE_OP
        }
    }

    if strings.IndexRune("<",c) >= 0 {
        if l.peek()=='=' {
            s1 := l.pos-1
            l.acceptRun("=")
            s2 := l.pos
            lval.id = l.s[s1:s2]
            return LE_OP
        }
    }

    if strings.IndexRune(">",c) >= 0 {
        if l.peek()=='=' {
            s1 := l.pos-1
            l.acceptRun("=")
            s2 := l.pos
            lval.id = l.s[s1:s2]
            return GE_OP
        }
    }

    if strings.IndexRune("<",c) >= 0 {
        if l.peek()=='<' {
            s1 := l.pos-1
            l.acceptRun("<")
            s2 := l.pos
            lval.id = l.s[s1:s2]
            return LEFT_OP
        }
    }

    if strings.IndexRune(">",c) >= 0 {
        if l.peek()=='>' {
            s1 := l.pos-1
            l.acceptRun(">")
            s2 := l.pos
            lval.id = l.s[s1:s2]
            return RIGHT_OP
        }
    }

    if strings.IndexRune("+",c) >= 0 {
        if l.peek()=='+' {
            s1 := l.pos-1
            l.acceptRun("+")
            s2 := l.pos
            lval.id = l.s[s1:s2]
            return INC_OP
        }
    }

    if strings.IndexRune("-",c) >= 0 {
        if l.peek()=='-' {
            s1 := l.pos-1
            l.acceptRun("-")
            s2 := l.pos
            lval.id = l.s[s1:s2]
            return DEC_OP
        }
    }

    if strings.IndexRune("-",c) >= 0 {
        if l.peek()=='>' {
            s1 := l.pos-1
            l.acceptRun(">")
            s2 := l.pos
            lval.id = l.s[s1:s2]
            return PTR_OP
        }
    }


    if strings.IndexRune("+",c) >= 0 {
        if l.peek()=='=' {
            s1 := l.pos-1
            l.acceptRun("=")
            s2 := l.pos
            lval.id = l.s[s1:s2]
            return ADD_ASSIGN
        }
    }

    if strings.IndexRune("-",c) >= 0 {
        if l.peek()=='=' {
            s1 := l.pos-1
            l.acceptRun("=")
            s2 := l.pos
            lval.id = l.s[s1:s2]
            return SUB_ASSIGN
        }
    }

    if strings.IndexRune("*",c) >= 0 {
        if l.peek()=='=' {
            s1 := l.pos-1
            l.acceptRun("=")
            s2 := l.pos
            lval.id = l.s[s1:s2]
            return MUL_ASSIGN
        }
    }

    if strings.IndexRune("/",c) >= 0 {
        if l.peek()=='=' {
            s1 := l.pos-1
            l.acceptRun("=")
            s2 := l.pos
            lval.id = l.s[s1:s2]
            return DIV_ASSIGN
        }
    }

    if strings.IndexRune("*",c) >= 0 {
        if l.peek()=='=' {
            s1 := l.pos-1
            l.acceptRun("=")
            s2 := l.pos
            lval.id = l.s[s1:s2]
            return MOD_ASSIGN
        }
    }

    if strings.IndexRune("&",c) >= 0 {
        if l.peek()=='=' {
            s1 := l.pos-1
            l.acceptRun("=")
            s2 := l.pos
            lval.id = l.s[s1:s2]
            return AND_ASSIGN
        }
    }

    if strings.IndexRune("^",c) >= 0 {
        if l.peek()=='=' {
            s1 := l.pos-1
            l.acceptRun("=")
            s2 := l.pos
            lval.id = l.s[s1:s2]
            return XOR_ASSIGN
        }
    }

    if strings.IndexRune("|",c) >= 0 {
        if l.peek()=='=' {
            s1 := l.pos-1
            l.acceptRun("=")
            s2 := l.pos
            lval.id = l.s[s1:s2]
            return OR_ASSIGN
        }
    }

    //fmt.Println("Returning char with no lex match", c, string(c))

    return int(c)
}

