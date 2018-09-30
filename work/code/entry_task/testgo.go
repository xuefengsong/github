package main

import (
 "fmt"
 "strconv"
)

func openFile(b string) (int,error) {
    a,err := strconv.Atoi(b)
    if err != nil{
        fmt.Println(err)
        fmt.Println("error")
    }
    return a,nil
}

func main() {
    a,err := openFile("example.go")
    if err !=nil{
        fmt.Println("类型错误")
    }
    fmt.Println(a)
    /*
    if err.Error() == "file does not exist" {
        // handle "file does not exist" error
        return
    }

    if err.Error() == "no priviledge" {
        // handle "no priviledge" error
        return
    }*/
}