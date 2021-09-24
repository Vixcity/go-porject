pacage main // 当前程序的包名
/*
一行代码的引入方式
import "fmt"
import "time"
*/

import (
	"fmt"
	"time"
)

func main()  {//函数的左花括号一定要和函数在同一行
	
	fmt.Println("hello GO!")
	time.Sleep(1 * time.Second)

}
