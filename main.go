package main

import (
    "fmt"
    //"intel/test/verify/parser"
    "os"
    "flag"
    //"io"
    //"bytes"
    _ "github.com/lib/pq"
    //"database/sql"
    //"intel/test/selectorproperty"
    "intel/qac/model"
    //"github.com/davecheney/profile"
    "runtime/pprof"
)
//go:generate -command yacc go tool yacc -o model/y.go model/grammar.y

func main() {
    //defer profile.Start(profile.CPUProfile).Stop()

    model.InitDB()

    var (
        f_compileFile      = flag.String("file", "file", "file to compile")
        f_cpuprofile       = flag.String("cpuprofile", "cpuprofile", "write cpu profile to file")
    )
    flag.Parse()

    if len(*f_compileFile) == 0 {
        fmt.Println("no file specified.")
        return
    }
    if *f_cpuprofile != "" {
        f, err := os.Create(*f_cpuprofile)
        if err != nil {
            fmt.Println("Error: ", err)
        }
        pprof.StartCPUProfile(f)
        defer pprof.StopCPUProfile()
    }

    indent := 1

    for i:=0; i<1; i++ {
        model.StartIt(*f_compileFile, indent)
        fmt.Printf("%v,",i)
    }

}


