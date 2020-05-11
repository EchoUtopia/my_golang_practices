package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

type operation uint

const (
	INSERT operation = 1
	DELETE           = 2
	MOVE             = 3
)

func generateDiff(src, dst []string) {
	operations := shortestOperations(src, dst)
	srcIndex, dstIndex := 0, 0

	for _, op := range operations {
		switch op {
		case INSERT:
			fmt.Println("+" + dst[dstIndex])
			dstIndex += 1

		case MOVE:
			fmt.Println(" " + src[srcIndex])
			srcIndex += 1
			dstIndex += 1

		case DELETE:
			fmt.Println("-" + src[srcIndex])
			srcIndex += 1
		}
	}
}

func shortestOperations(src, dst []string) []operation {
	sl := len(src)
	dl := len(dst)
	// 最大步数
	maxD := sl + dl
	var x, y int
	var traces []map[int]int

	// 处理初始的对角线
	xt := 0
	for xt < sl && xt < dl && src[xt] == dst[xt] {
		xt++
	}
	traces = append(traces, map[int]int{0: xt})

loop:
	for d := 1; d <= maxD; d++ {
		trace := map[int]int{}
		traces = append(traces, trace)
		lastT := traces[d-1]
		// 以0点为例，向右k = 1， 向下k = -1， 相差2
		for k := -d; k <= d; k += 2 {
			// k == -d, 说明是一直在往下走，k == d， 说明是一直在往右走
			// lastT[k-1] >= lastT[k+1]， 优先选择向右走, lastT[k-1] > lastT[k+1] 说明k已经最大了
			if k == d || lastT[k-1] >= lastT[k+1] {
				x = lastT[k-1] + 1
			} else {
				x = lastT[k+1]
			}
			y = x - k
			// 处理对角线
			for x < sl && y < dl && src[x] == dst[y] {
				x++
				y++
			}
			trace[k] = x
			if x == sl && y == dl {
				break loop
			}
		}
	}

	//printTrace(traces)
	// 从终点回溯
	var preK, preX, preY int
	operations := []operation{}
	for d := len(traces) - 1; d > 0; d-- {
		k := x - y
		lastT := traces[d-1]
		// 逻辑同上
		if k == d || lastT[k-1] >= lastT[k+1] {
			preK = k - 1
		} else {
			preK = k + 1
		}
		preX = lastT[preK]
		preY = preX - preK
		for x > preX && y > preY {
			operations = append(operations, MOVE)
			x--
			y--
		}
		if preX == x {
			operations = append(operations, INSERT)
		} else {
			operations = append(operations, DELETE)
		}
		x, y = preX, preY

	}
	// 处理初始对角线
	if traces[0][0] != 0 {
		for i := 0; i < traces[0][0]; i++ {
			operations = append(operations, MOVE)
		}
	}
	return reverse(operations)

}

func printTrace(trace []map[int]int) {
	for d := 0; d < len(trace); d++ {
		fmt.Printf("d = %d:\n", d)
		v := trace[d]
		for k := -d; k <= d; k += 2 {
			x := v[k]
			y := x - k
			fmt.Printf("  k = %2d: (%d, %d)\n", k, x, y)
		}
	}
}

func reverse(s []operation) []operation {
	result := make([]operation, len(s))

	for i, v := range s {
		result[len(s)-1-i] = v
	}

	return result
}

func main() {
	f1, err := ioutil.ReadFile(`/Users/echo/go/src/test/myer_diff/source.txt`)
	if err != nil {
		panic(err)
	}
	f2, err := ioutil.ReadFile(`/Users/echo/go/src/test/myer_diff/revision.txt`)
	if err != nil {
		panic(err)
	}
	source := []string{}
	revision := []string{}
	for idx := bytes.IndexByte(f1, '\n'); idx != -1; idx = bytes.IndexByte(f1, '\n') {
		source = append(source, string(f1[:idx]))
		f1 = f1[idx+1:]
	}
	source = append(source, string(f1))
	for idx := bytes.IndexByte(f2, '\n'); idx != -1; idx = bytes.IndexByte(f2, '\n') {
		revision = append(revision, string(f2[:idx]))
		f2 = f2[idx+1:]
	}
	revision = append(revision, string(f2))
	generateDiff(source, revision)
}
