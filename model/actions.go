package model

import (
	"fmt"
	"strings"
    "bufio"
    "database/sql"
    "io"
    "os"
    "bytes"
    "net/http"
    "net/url"
    "time"
    "io/ioutil" 
    "intel/test/datagen2"
    "intel/scriptgen2/model"
    //"path/filepath"
    //"unicode/utf8"
    //"intel/test/selectorproperty"
    //"intel/test/parser1"
)


var base int

// --------------------------------------------------------------------------------

type Jar struct {
    cookies map[string][]*http.Cookie
}

func NewJar() *Jar {
    jar := new(Jar)
    jar.cookies = make(map[string][]*http.Cookie)
    return jar
}

func (jar *Jar) SetCookies(u *url.URL, cookies []*http.Cookie) {
    jar.cookies[u.Host] = cookies
}

func (jar *Jar) Cookies(u *url.URL) []*http.Cookie {
    return jar.cookies[u.Host]
}

// --------------------------------------------------------------------------------

type CommandResult struct {
    Result     bool
    Message     string
}

func NewCommandResult() *CommandResult {
    v := new(CommandResult)
    return v
}

// --------------------------------------------------------------------------------

func readline(fi *bufio.Reader) (string, bool) {
    s, err := fi.ReadString('\n')
    if err != nil {
        return "", false
    }
    return s, true
}

func GetUserId(emailUser string) string {
    sql := fmt.Sprintf("SELECT user_id FROM users WHERE email = '%s'", emailUser )
    rows, err := DbHub.Query(sql)
    defer rows.Close()
    if err != nil {
        fmt.Println(err)
        return ""
    }
    for rows.Next() {
        var user_id string
        if err := rows.Scan(&user_id); err != nil {
            return ""
        }
        return user_id
    }
    return ""
}

func MakeAppTid(uuid string, app string) *CommandResult {
    rv := NewCommandResult()
    rv.Result = false
    var ret string
    appId := GetAppId(app)
    if len(appId) > 0 {
        sql := fmt.Sprintf("SELECT makeapptid_rest(%s, '%s')", appId, uuid)
        rows, err := DbHub.Query(sql)
        defer rows.Close()
        if err != nil {
            fmt.Println(err)
            return rv
        }
        for rows.Next() {
            //var txt string
            err = rows.Scan(&ret)
            if err != nil {
                panic(err)
            }
            if len(ret) > 0 {
                rv.Result = true
                rv.Message = ret
            }
            //fmt.Printf("MakeAppTid rows rv='%v' txt=%v \n", ret, txt) 
        }
        return rv
    }
    return rv
}

/*
func DeleteAppUser(appName string, userName string) *CommandResult {
    var appId string
    var userId string
    var result *CommandResult
    userId = GetUserId(userName)
    appId := GetAppId(appName)
    sql := fmt.Sprintf("SELECT deleteappuser_rest()",appId, userId)
    return result
}
*/

func RunSQL(sql string, appName string) *CommandResult {
    rv := NewCommandResult()
    //fmt.Println("sql=", sql) 
    mDb := GetDbRef(appName)
    defer mDb.Close()
    rows, err := mDb.Query(sql)
    if err != nil {
        panic(err)
        rv.Result = false
        rv.Message = fmt.Sprintf("ERROR: %v\n", err)
        return rv
    }
    for rows.Next() {
        var ret string
        //var txt string
        err = rows.Scan(&ret)
        if err != nil {
            panic(err)
        }
        if len(ret) > 0 {
            rv.Message = ret
        }
        //fmt.Printf("RunSQL rows rv='%v' txt=%v \n", ret, txt) 
    }
    defer rows.Close()
    rv.Result = true
    //rv.Message = "OK"
    //fmt.Printf("RunSQL rows rv='%v' \n", rv.Message) 
    return rv
}



func DeleteAppTid(appName string, uuid string) *CommandResult {
    rv := NewCommandResult()
    if !ApplicationExists(appName) {
        rv.Result = false
        rv.Message = fmt.Sprintf("Invalid Name: Application name %s not exist.\n", appName)
        //fmt.Printf("Invalid Name: Application name %s not exist.\n", appName)
        return rv
    }
    //fmt.Println("Create App Done") 
    appId := GetAppId(appName)
    sql := fmt.Sprintf("SELECT deleteapptid_rest('%v','%v')",appId,uuid)
    //fmt.Println("sql=", sql) 
    //fmt.Println("DbHub:=======>", DbHub)
    rows, err := DbHub.Query(sql)
    if err != nil {
        panic(err)
        rv.Result = false
        rv.Message = fmt.Sprintf("ERROR: %v\n", err)
        return rv
    }
    for rows.Next() {
        var ret string
        //var txt string
        err = rows.Scan(&ret)
        if err != nil {
            panic(err)
        }
        if len(ret) > 0 {
            rv.Message = ret
        }
        //fmt.Printf("DeleteAppTid rows rv='%v' txt=%v \n", ret, txt) 
    }
    defer rows.Close()
    rv.Result = true
    //rv.Message = "OK"
    return rv
}


func DeleteAppUser( emailUser string,  appName string) *CommandResult {
    rv := NewCommandResult()
    if !ApplicationExists(appName) {
        rv.Result = false
        rv.Message = fmt.Sprintf("Invalid Name: Application name %s not exist.\n", appName)
        //fmt.Printf("Invalid Name: Application name %s not exist.\n", appName)
        return rv
    }
    //fmt.Println("Create App Done") 
    appId := GetAppId(appName)
    userId := GetUserId(emailUser)
    sql := fmt.Sprintf("SELECT deleteappuser_rest('%v','%v')",appId,userId)
    //fmt.Println("sql=", sql) 
    //fmt.Println("DbHub:=======>", DbHub)
    rows, err := DbHub.Query(sql)
    if err != nil {
        panic(err)
        rv.Result = false
        rv.Message = fmt.Sprintf("ERROR: %v\n", err)
        return rv
    }
    defer rows.Close()
    for rows.Next() {
        var ret string
        //var txt string
        err = rows.Scan(&ret)
        if err != nil {
            panic(err)
        }
        if len(ret) > 0 {
            rv.Message = ret
            rv.Result = true
        }
        //fmt.Printf("DeleteAppTid rows rv='%v' txt=%v \n", ret, txt) 
    }
    return rv
}

func DeleteApp(appName string) *CommandResult {
    rv := NewCommandResult()
    if !ApplicationExists(appName) {
        rv.Result = false
        rv.Message = fmt.Sprintf("Invalid Name: Application name %s not exist.\n", appName)
        //fmt.Printf("Invalid Name: Application name %s not exist.\n", appName)
        return rv
    }
    //fmt.Println("Create App Done") 
    appId := GetAppId(appName)
    sql := fmt.Sprintf("SELECT deleteapp_rest('%s')",appId)
    //fmt.Println("sql=", sql) 
    //fmt.Println("DbHub:=======>", DbHub)
    rows, err := DbHub.Query(sql)
    if err != nil {
        panic(err)
        rv.Result = false
        rv.Message = fmt.Sprintf("ERROR: %v\n", err)
        return rv
    }
    defer rows.Close()
    for rows.Next() {
        var ret string
        //var txt string
        err = rows.Scan(&ret)
        if err != nil {
            panic(err)
        }
        if len(ret) > 0 {
            rv.Message = ret
            rv.Result = true
        }
    }
    return rv
}

func CreateApp(appName string) *CommandResult {
    rv := NewCommandResult()
    if ApplicationExists(appName) {
        rv.Result = false
        rv.Message = fmt.Sprintf("Invalid Name: Application name %s already taken.\n", appName)
        //fmt.Printf("Invalid Name: Application name %s already taken.\n", appName)
        return rv
    }
    //fmt.Println("Create App Done") 
    sql := fmt.Sprintf("SELECT makeapp_rest('%s')",appName)
    //fmt.Println("sql=", sql) 
    //fmt.Println("DbHub:=======>", DbHub)
    rows, err := DbHub.Query(sql)
    if err != nil {
        panic(err)
        rv.Result = false
        rv.Message = fmt.Sprintf("ERROR: %v\n", err)
        return rv
    }
    defer rows.Close()
    for rows.Next() {
        var res string
        if err := rows.Scan(&res); err != nil {
            return rv
        }
        if len(res) > 0 {
            rv.Message = res
        }
        rv.Result = true
        return rv
    }

    return rv
}

func DropApp() {
    currState=0 
    fmt.Println("Drop App Done") 
}

func RunCommandFile(filename string) {
    buf := bytes.NewBuffer(nil)
    f, _ := os.Open(filename) 
    io.Copy(buf, f) 
    f.Close()
    sbuf := string(buf.Bytes())
    //fmt.Println(sbuf)
    prvLexer.s = sbuf
    prvLexer.pos = 0
    currState = 0
    commandState = 0
    yyParse(prvLexer)
}

func ApplicationExists(app string) bool {
    var ok bool = false
    sql := fmt.Sprintf("SELECT app_id FROM apps WHERE appname = '%s'", app )
    rows, err := DbHub.Query(sql)
    defer rows.Close()
    if err != nil {
        fmt.Println(err)
    }
    for rows.Next() {
        ok = true
        var app_id string
        if err := rows.Scan(&app_id); err != nil {
        }
    }
    return ok
}

func UseApplication(app string) bool {
    var ok bool = false
    currentApp = app
    sql := fmt.Sprintf("SELECT app_id FROM apps WHERE appname = '%s'", app )
    rows, err := DbHub.Query(sql)
    defer rows.Close()
    if err != nil {
        fmt.Println(err)
    }
    for rows.Next() {
        ok = true
        var app_id string
        if err := rows.Scan(&app_id); err != nil {
        }
        currentAppId = app_id
    }
    if !ok {
        fmt.Printf("Application '%s' not found \n", app)
    }
    currState = 0
    commandState = 0
    return ok
}

func GetAppId(app string) string {
    sql := fmt.Sprintf("SELECT app_id FROM apps WHERE appname = '%s'", app )
    rows, err := DbHub.Query(sql)
    defer rows.Close()
    if err != nil {
        fmt.Println(err)
    }
    for rows.Next() {
        var app_id string
        if err := rows.Scan(&app_id); err != nil {
        }
        return app_id
    }
    return ""
}

func MakeAppuser(email string, app string) bool {
    appId := GetAppId(app)
    if len(appId) > 0 {
        sql := fmt.Sprintf("SELECT makeappuser_rest(%s, '%s')", appId, email)
        rows, err := DbHub.Query(sql)
        defer rows.Close()
        if err != nil {
            fmt.Println(err)
            return false
        }
        return true
    }
    return false
}

func MakeAppuser2(email string, app string) *CommandResult {
    rv := NewCommandResult()
    rv.Result = false
    appId := GetAppId(app)
    if len(appId) > 0 {
        sql := fmt.Sprintf("SELECT makeappuser_rest(%s, '%s')", appId, email)
        rows, err := DbHub.Query(sql)
        defer rows.Close()
        if err != nil {
            fmt.Println(err)
            return rv
        }
        for rows.Next() {
            var res string
            if err := rows.Scan(&res); err != nil {
                return rv
            }
            //fmt.Printf("Makeappuser result: %v \n",res)
            if len(res) > 0 {
                rv.Message = res
            }
            rv.Result = true
            return rv
        }
        return rv
    }
    return rv
}


func TestRoute(s string) string {
    var sOut string
    fmt.Sprintf("-------------------------------------------------------------------")
    var err error
    var cookies []*http.Cookie
    var cookie *http.Cookie
    for j:=0;j<len(mapCookie);j++ {
        cookie = new(http.Cookie)
        //fmt.Printf("\nSending...\nsession-name=%v\n",mapCookie["session-name"])
        cookie.Name = "session-name"
        cookie.Value = mapCookie["session-name"]
        cookies = append(cookies, cookie)
    }
    jar := NewJar()
    url, err := url.Parse(s)
    if err != nil {
        fmt.Println(err)
        panic(err)
    }
    jar.SetCookies(url, cookies)
    client := http.Client{nil, nil, jar, time.Minute*1000}
    //fmt.Printf("\nURL: %s\n",s)
    respRoot, err := client.Get(s)
    if err != nil {
        fmt.Println(err)
        panic(err)
    }
    defer respRoot.Body.Close()
    hdr := respRoot.Header
    ck := hdr["Set-Cookie"]
    if ck != nil {
        //fmt.Printf("\nReceive\nSet-Cookie\n%v\n\n",ck)
        arr := strings.Split(string(ck[0]),";")
        for i:=0;i<len(arr); i++ {
            itm := arr[i]
            indx := strings.Index(itm, "=")
            s2 := itm[indx+1:len(itm)]
            arr2 := strings.Split(itm,"=")
            //fmt.Printf("name:%s  value:%s\n", arr2[0], s2)
            if arr2[0]=="session-name" {
                mapCookie["session-name"] = s2
            }
        }
    }
    bodyRoot, err := ioutil.ReadAll(respRoot.Body)
    if err != nil {
        fmt.Println(err)
    }
    //fmt.Println("Response:")
    sOut = fmt.Sprintf("%v",string(bodyRoot))
    return sOut
}

func DoTest(appName string) bool {
    envHost := GetStyleObjectValue("env", "host")
    if len(envHost) == 0 {
        envHost = host
    }
    envEmail := GetStyleObjectValue("env", "email")
    if len(envEmail) == 0 {
        envEmail = "joshua.boelter@intel.com"
    }
    TestRoute(envHost)
    rval := TestRoute(fmt.Sprintf("%sadmin/prov/api/v1/app/%s/user?appname=%s",envHost, currentAppId, appName))
    fmt.Println(rval)
    rval = TestRoute(fmt.Sprintf("%sadmin/prov/api/v1/app/%s/tid?appname=%s",envHost, currentAppId, appName))
    fmt.Println(rval)
    list := GetAppTerminals()
    for i := range list {
        rval := TestRoute(fmt.Sprintf("%sapp/%s/api/v1/%s",envHost, currentAppId, list[i]))
        fmt.Println(rval)
    }
    rval = TestRoute(fmt.Sprintf("%smeta/api/v1/appsforuser?email=%s",envHost, envEmail))
    fmt.Println(rval)
    return true
}


func DoTestRoute1(appName string, category string) ( bool, string) {
    inURL := map[string]bool {
        "manufacturers": true,
        "eventnames": true,
        "newusers": true,
        "sessions": true,
        "resolutions": true,
        "memsize": true,
        "appversion": true,
        "locale": true,
        "cpurevision": true,
        "cpubrand": true,
        "usersbycountrycode": true,
        "usersbycountrysubdivisioncode": true,
        "newusersbydate": true,
        "totalusersbydate": true,
        "timeperuserbydate": true,
        "eventsbydate": true,
        "sessionsbydate": true,
        "allevents": true,
        "environmentenumerations": true,
        "environmenthistogram": true,
        "segmentnames": true,
        "manusku": true,
        "userloyalty": true,
    }
    var rval string
    UseApplication(appName)
    envHost := GetStyleObjectValue("env", "host")
    if len(envHost) == 0 {
        envHost = host
    }
    envEmail := GetStyleObjectValue("env", "email")
    if len(envEmail) == 0 {
        envEmail = "joshua.boelter@intel.com"
    }
    TestRoute(envHost)
    if category == "user" {
        s := fmt.Sprintf("%sadmin/prov/api/v1/app/%s/user?appname=%s",envHost, currentAppId, appName)
        rval = TestRoute(s)
        rlog7(1,"DoTestRoute1: rval=%v s=%v\n", rval, s)
    }
    if category == "tid" {
        s := fmt.Sprintf("%sadmin/prov/api/v1/app/%s/tid?appname=%s",envHost, currentAppId, appName)
        rval = TestRoute(s)
        rlog7(1,"DoTestRoute1: rval=%v s=%v\n", rval, s)
    }
    if category == "appsforuser" {
        s := fmt.Sprintf("%smeta/api/v1/appsforuser?email=%s",envHost, envEmail)
        rval = TestRoute(s)
        rlog7(1,"DoTestRoute1: rval=%v s=%v\n", rval, s)
    }
    if inURL[category] {
        s := fmt.Sprintf("%sapp/%s/api/v1/%s",envHost, currentAppId, category)
        rval = TestRoute(s)
        rlog7(1,"DoTestRoute1: rval=%v s=%v\n", rval, s)
    }

    return true, rval
}

func GetAppTerminals() []string {
    rv := []string {"locale", "manufacturers", "eventnames", "newusers", "sessions", "resolutions", "memsize",    "appversion", "locale", "cpurevision", "cpubrand", "usersbycountrycode",  "usersbycountrysubdivisioncode", "enumfilters", "enumfieldvalues",  "verifyfilter", "newusersbydate", "totalusersbydate", "timeperuserbydate", "eventsperuserbydate", "eventsbydate", "sessionsbydate", "allevents", "environmentenumerations", "environmenthistogram", "segmentnames", "manusku",  "userloyalty" }

    //"allusers",

    return rv
}

func GetStyleObjectValue(objName string, prop string) string {
    if len(symObjectList.List) > 0 {
        for i:=len(symObjectList.List)-1; i>=0; i-- {
            symObject := symObjectList.List[i]
            for j:=len(symObject.Selectors)-1; j>=0; j-- {
                sel := symObject.Selectors[j]
                if sel==objName {
                    for k:=0; k<len(symObject.Properties); k++ {
                        p := symObject.Properties[k]
                        if p.Name==prop  {
                            return (p.Value)
                        }
                    }
                }
            }
        }
    }
    return ""
}

func GenerateTestData(appName string, loop int) bool {
    valid := UseApplication(appName) 
    if !valid {
        return false
    }
    if loop==0 {
        return false
    }
    if loop > 1000 {
        fmt.Println("Loop over 1000 considered invalid.")
        return false
    }
    for i:=0; i<loop; i++ {
        var buf bytes.Buffer
        user := model.NewUser()
        buf = user.Simulate()
        //fmt.Println("Buffer:",buf)
        //buf.WriteTo(os.Stdout)
        SetClientDB(appName, GetIP())
        dg := datagen2.NewDataGen2(buf)
        dg.Generate(Db)
        //fmt.Println("LOOP COUNT: ", i+1)
    }
    Db.Close()
    return true
}

func SetClientDB(s string, h string) {
    clientDbConnectStr := GetStyleObjectValue("env", "ClientDbConnect")
    clientDbConnectStr = fmt.Sprintf("host=%s port=5432 user=docker password=docker dbname=%s sslmode=disable", GetIP(), s)
    //fmt.Println("SetClientDB:", clientDbConnectStr)
    Db, err = sql.Open("postgres", clientDbConnectStr)
    if err != nil {
        panic(err)    
    }
}

func PopulateAppFromFile(appName string, fileName string) {
    valid := UseApplication(appName) 
    if !valid {
        return
    }
    var b bytes.Buffer
    var arrByte []byte
    arrByte,_ = ioutil.ReadFile(fileName)
    b.Write([]byte(arrByte))
    SetClientDB(appName, GetIP())
    dg := datagen2.NewDataGen2(b)
    dg.Generate(Db)
    Db.Close()
    return
}


func VerifyCommand(appName string, urlId string, txt string) {
    fmt.Println("VerifyCommand")
    fmt.Println(txt)
    valid := UseApplication(appName) 
    if !valid {
        return
    }
}
/*
func DoImport(v string) {
    dir := fmt.Sprintf("%s%s",rootPath, v)
    fmt.Printf("File: %s \n", dir)
    files, _ := ioutil.ReadDir(dir)
    //_ = files
    //fmt.Println(files)
    for _, f := range files {
        if filepath.Ext(f.Name())==".ias" {
            fmt.Printf("%s %s \n",f.Name(), filepath.Ext(f.Name()))
            ParseFile(f.Name())
        }
    }    
}
*/

func DoListApps() string {
    var retval string
    currState = 0
    commandState = 0
    arrApps = []App{}
    sql := "SELECT appname FROM apps ORDER BY appname"
    rows, err := DbHub.Query(sql)
    defer rows.Close()
    if err != nil {
        fmt.Println(err)
    }
    retval = fmt.Sprintf("")
    for rows.Next() {
        var name string
        if err := rows.Scan(&name); err != nil {
        }
        retval = fmt.Sprintf("%v%v, ", retval, name)
        app := new(App)
        app.Appname = name
        arrApps = append(arrApps, *app)
    }
    //fmt.Println(arrApps)    
    return retval
}

func DoTestCommand(s string) {
    //fmt.Println("TEST ", $2)
    if UseApplication(s) {
        DoTest(s)
    }    
}
/*
func ParseFile(fileName string) {
    executionMode = 0xF
}

func DoLoadFile(p *LoadParameter) {
    loadpath := p.GetPath()
    fmt.Println(loadpath)
    dir := fmt.Sprintf("%s%s",rootPath, loadpath)
    fmt.Printf("File: %s \n", dir)
    files, _ := ioutil.ReadDir(dir)
    //_ = files
    //fmt.Println(files)
    parser := NewParser() 
    for _, f := range files {
        if filepath.Ext(f.Name())==".ias" {
            fmt.Printf("%s %s \n",f.Name(), filepath.Ext(f.Name()))
            fname := fmt.Sprintf("%s%s", dir, f.Name())
            parser.ParseFile(fname)
        }
    }    

}

func DoLoadPackages() {
    p := NewLoadParameter()
    p.Add("test1")
    DoLoadFile(p)
}
*/
