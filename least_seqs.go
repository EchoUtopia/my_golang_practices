package main
import (
	"fmt"
)

func LeastSeqs(a []int) []int{
	rec := make([]int, len(a))
	leng := 0
	out := make([]int,0, 10)
	for i:=0;i< len(a);i++{
		rec[i] = 1
		for j:=0;j<i;j++{
			if a[j] <= a[i] && rec[j] + 1>rec[i]{
				rec[i] = rec[j]+1
			}
		}
		if rec[i] > leng{
			leng = rec[i]
			out = append(out, a[i])
		}
	}
	println(leng)
	//fmt.Printf("%v\n", rec)
	return out
}

func main(){
	a := []int{'a', 'c', 'd', 'b','白', 'e', }
	out := LeastSeqs(a)
	fmt.Printf("%v\n", out)
}
