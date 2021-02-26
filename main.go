package main

import "fmt"

func main(){
      
      var student [5] string
      student[0]="tahsin"
      student[1]="tithi"
      student[2]="tareq"
      student[3]="anonna"
      student[4]="ferdaus"      
      
      x:=student[0:3]

      fmt.Println(x)
    // fmt.Println(len(student))
     //fmt.Println(len(x))



}