package main
import "fmt"
func main() {

	var a int =100
	var b int =200
	//var ret int
	swap2(&a,&b)
	//fmt.Printf(a,b)

	//ret = max(a,b)
	//fmt.Printf("最大值为：%d\n",ret)

    fmt.Printf("a值为：%d\n",a)
    fmt.Printf("b值为：%d\n",b)

	//var a,b string
    //a,b=swap("guoshuai","lin")
    //fmt.Printf(a,b)

	
}

/*func max(num1,num2 int) int {

	var result int

	if num1>num2 {
		result=num1
	}else{
		result=num2
	}
	return result
}

func swap(s1,s2 string) (string,string){
	return s2,s1
	
}*/

func swap2(x *int,y *int) {
	var temp int
	temp=*x
	*x=*y
	*y=temp
}