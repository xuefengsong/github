package main
import (
    "fmt"
    "net/http"
    "net/url"
    //"encoding/json"
    "strconv"
)

func main() { 
http.HandleFunc("/api", entry_task)
http.ListenAndServe("127.0.0.1:9999", nil)
}
func entry_task(rw http.ResponseWriter, r *http.Request) {
  u, err := url.Parse(r.URL.String())
  if err != nil{
    fmt.Fprintln(rw,"{\"err_code\":11,\"err_msg\":\"system error\",\"reference\":\"\"}")
    return
    }

  values, err_query := url.ParseQuery(u.RawQuery)
  if err_query != nil{
    fmt.Fprintln(rw,"{\"err_code\":11,\"err_msg\":\"system error\",\"reference\":\"\"}")
    return
    }

  var i int =0
  for k, v := range values {
    i++
    fmt.Println(k,v)
    }
  if i !=2 {
    //return_str := fmt.Sprintf("\"err_code\":21,\"err_msg\":\"empry or wrong params\",\"reference\":\"\"")
    fmt.Fprintln(rw,"{\"err_code\":21,\"err_msg\":\"empry or wrong params\",\"reference\":\"\"}")
    }else {
        //判断必选参数是否为空
          _,ok_a := values["a"]
          _,ok_b := values["b"]
          
          if ok_a == false || ok_b == false{
                //return_str=:= fmt.Sprintf("\"err_code\":21,\"err_msg\":\"success\",\"reference\":\"NO.%s is %s\"",)
                fmt.Fprintln(rw,"{\"err_code\":21,\"err_msg\":\"empry or wrong params\",\"reference\":\"\"}")
            }else{
                _, err := strconv.Atoi(values["a"][0])
                if err != nil{
                    fmt.Fprintln(rw,"{\"err_code\":21,\"err_msg\":\"empry or wrong params\",\"reference\":\"\"}")
                    }else{
                        fmt.Fprintln(rw,"{\"err_code\":0,\"err_msg\":\"success\",\"reference\":\"NO."+values["a"][0]+" is "+values["b"][0]+"\"}")
                    }
                
                }
        }
}
