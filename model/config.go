package model

import (
    "fmt"
    "os"
    _ "github.com/lib/pq"
    "database/sql"
    "intel/test/selectorproperty"
 )

type ConfigQAC struct {
	IP 		string
	RootPath string
}

func NewConfigQAC() *ConfigQAC {
	v := new(ConfigQAC)
	v.RootPath = GetRootPath()
	v.IP = GetIP()
	return v
}

func GetRootPath() string {
    return "/home/bgonzax/Intel/Analytics/bbgonzax_ubuntu_9471/aws/gosrc/src/intel/tests/testcases/"

}

func GetIP() string {
    return "172.17.0.3"
}


func GetDbConnectionStringHub() string {
	return fmt.Sprintf("host=%s port=5432 user=docker password=docker dbname=iashub  sslmode=disable", GetIP())
}

var config 			*ConfigQAC
var symObjectList   *selectorproperty.SymObjectList
var mapCookie       map[string]string

func InitDB()  {
    config = NewConfigQAC()
    symObjectList = selectorproperty.NewSymObjectList()
    DbConnectionStringHub := GetDbConnectionStringHub()
    ClientDbConnectionString := os.Getenv("CONNECTION_STRING")
    ClientDbConnectionString = fmt.Sprintf("host=%s port=5432 user=docker password=docker dbname=test sslmode=disable", GetIP())
    Db, err = sql.Open("postgres", ClientDbConnectionString)
    if err != nil {
        panic(err)    
    }
    DbHub, err = sql.Open("postgres", DbConnectionStringHub)
    if err != nil {
        panic(err)    
    }
    mapCookie = make(map[string]string)
}

func GetDbRef(appName string) *sql.DB {
    connstr := fmt.Sprintf("host=%s port=5432 user=docker password=docker dbname=%v sslmode=disable", GetIP(),appName)
    refDb, err := sql.Open("postgres", connstr)
    if err != nil {
        panic(err)    
    }
    return refDb
}
