package main

import "fmt"

var test string
type s struct{
    a int
    b string
}
func main() {
    test="aaa"
    sb :=s{1,"ddd"}
    fmt.Println(sb.b)
    //a :=sb.b
    abc,_ := fmt.Printf("this is a name:%s",sb.b)
    abc =2222
    fmt.Println(abc)
    /*var sheet_name string ="aa"
    excel.create(sheet_name)
    
    err :=SetCellValue(sheet_name ,"A2" ,"testkkk" )
    if err !=nil{
    	fmt.Print(err)
    }
    excel.save("./abc.xlsx" )
    */
	/*
	xlsx := excelize.CreateFile()
	 // Create a new sheet.
	 xlsx.NewSheet(2, "Sheet2")
	 // Set value of a cell.
	 xlsx.SetCellValue("Sheet2", "A2", "Hello world.")
	 xlsx.SetCellValue("Sheet1", "B2", 100)
	 // Set active sheet of the workbook.
	 xlsx.SetActiveSheet(2)
	 // Save xlsx file by the given path.
	 err := xlsx.WriteTo("./Workbook.xlsx")
	 if err != nil {
	  fmt.Println(err)
	  os.Exit(1)
	 }*/
	 
}
 