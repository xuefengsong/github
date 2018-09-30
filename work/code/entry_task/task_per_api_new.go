package main
import (
    "fmt"
    "net/http"
    "net/url"
    //"encoding/json"
    "strconv"
)
type BaseJsonBean struct {
    code    int         `json:"code"`
    err_msg  string      `json:"err_msg"`
    reference string      `json:"reference"`
}

func NewBaseJsonBean() *BaseJsonBean {
    return &BaseJsonBean{}
}
func main() { 
http.HandleFunc("/api", entry_task)
http.ListenAndServe("127.0.0.1:9999", nil)
}
func get_params_map(r *http.Request) (map,error) {
    u, err := url.Parse(r.URL.String())
    if err !=nil {
        return nil,err
    }
    values, err_query := url.ParseQuery(u.RawQuery)
    if err_query != nil {
        return nil , err_query
    }
    return values,nil
}
func entry_task(rw http.ResponseWriter, r *http.Request) {
  values, err := get_params_map(r)
  if err != nil {
        fmt.Fprintln(rw,"{\"err_code\":11,\"err_msg\":\"system error\",\"reference\":\"\"}")
        return
    }
  var i int =0
  for k, v := range values {
    i++
    fmt.Println(k,v)
    //k=k+v[0]
    }
  if i !=2 {
    //return_str := fmt.Sprintf("\"err_code\":21,\"err_msg\":\"empry or wrong params\",\"reference\":\"\"")
    fmt.Fprintln(rw,"{\"err_code\":21,\"err_msg\":\"empry or wrong params\",\"reference\":\"\"}")
    }else {
        //判断必选参数是否为空
          _,ok_a :=values["a"]
          _,ok_b :=values["b"]
          //var a_int int = 0
          //参数a或者参数b不存在的话返回系统错误提示
          if ok_a==false||ok_b==false{
                //return_str=:= fmt.Sprintf("\"err_code\":21,\"err_msg\":\"success\",\"reference\":\"NO.%s is %s\"",)
                fmt.Fprintln(rw,"{\"err_code\":21,\"err_msg\":\"empry or wrong params\",\"reference\":\"\"}")
            }else{
                //如果参数a的类型不对返回错误提示
                _, err := strconv.Atoi(values["a"][0])
                if err != nil{
                    fmt.Fprintln(rw,"{\"err_code\":21,\"err_msg\":\"empry or wrong params\",\"reference\":\"\"}")
                    }else{
                        fmt.Fprintln(rw,"{\"err_code\":0,\"err_msg\":\"success\",\"reference\":\"NO."+values["a"][0]+" is "+values["b"][0]+"\"}")
                    }
                
                }
        }
}
