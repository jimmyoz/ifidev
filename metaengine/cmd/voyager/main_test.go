package main

import (
	"fmt"
	//"fmt"
	//"github.com/klauspost/cpuid"
	//"github.com/yanhuangpai/voyager/cmd/voyager/cmd"
	//"os"
	"testing"
	"time"
)

func TestFn1(t *testing.T) {
	main()
}


type IntConv func(op int) int

func timeSpent(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

//func main() {
//	fmt.Println("logicalScore",float64(cpuid.CPU.LogicalCores)*float64(cpuid.CPU.Hz)/1000000000.00)
//
//	if err := cmd.Execute(); err != nil {
//		println(os.Stderr, "Error:", err.Error())
//		os.Exit(1)
//	}
//}