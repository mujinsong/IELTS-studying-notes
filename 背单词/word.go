package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"math/rand"
	"time"
)
func main() {
	f, err := excelize.OpenFile("D:\\雅思\\单词本.xlsx")
	if err != nil {
		println(err.Error())
		return
	}
	rows, err := f.GetRows("Sheet1")
	var l =len(rows)
	//for cntrow,row := range rows {
	//	fmt.Println(cntrow,row[0],row[1])
	//	//for cntcol, colCell := range row {
	//	//	//fmt.Println(cntrow,cntcol,colCell)
	//	//}
	//}
	isok :=make([]bool,l)
	for i:=0 ; i<l ; i++ {
		rand.Seed(time.Now().Unix())
		x:=rand.Int63n(int64(l))
		if isok[x] {
			continue
		} else {
			for !isok[x] {
				fmt.Println(rows[x][0], rows[x][1])
				fmt.Print("写一遍:")
				var str1, str2 string
				fmt.Scanf("%s %s", &str1, &str2)
				if (str1 == rows[x][0] && str2 == rows[x][1]) || (str1 == "0") {
					isok[x]=true
				}
			}
		}
	}
}
