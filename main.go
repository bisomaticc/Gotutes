package main

import ("fmt")

func main(){

age:=45
fmt.Println(age<50)

name := []string{"ab","cd","ef","gh","mo"}

for index, values := range name{
 if index==1{
	fmt.Println("continuing ", index)
	continue
 }

 if index>2{
	fmt.Printf("ok %v",values)
	break
 }
 fmt.Printf("the value at %v postions is %v  \n", index, values)
}

}