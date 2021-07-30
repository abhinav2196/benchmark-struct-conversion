package main

import (
	"BenchmarkStructConversion/Approaches"
	"BenchmarkStructConversion/Models"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"sync"
	"time"
)

func main() {

	wg := sync.WaitGroup{}
	wg.Add(1)

	startTime := time.Now()

	go printStats()

	input := Models.Response1{
		Page:   1,
		Fruits: "x",
		Z:      &Models.Response2{X: 1, Y: "r"},
	}

	for i := 0; i < 10000000; i++ {
		var output Models.Response1
		//Uncomment any one of Approaches given below to test/benchmark them

		// err := Approaches.ConvertWithCopier(input,&output)
		//err := Approaches.ConvertWithNativeLib(input,&output)
		//err := Approaches.ConvertWithJsonIter(input,&output)
		//err := Approaches.ConvertWithEasyJson(input,&output)
		err := Approaches.ConvertWithManualMapping(input, &output)

		if err != nil {
			//fmt.Printf("Recieved an error")
			break
		}

		//fmt.Printf("Output printing= %v\n", output)

	}

	totalTimeTaken := time.Since(startTime)

	fmt.Println("Total time taken", totalTimeTaken)

	//runtime.GC()

	wg.Wait()

}

func printHeapInfo() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("env:%v, heapsys:%d,heapalloc:%d,heapidel:%d,heapreleased:%d,heapinuse:%d\n",
		os.Getenv("GODEBUG"), m.HeapSys, m.HeapAlloc, m.HeapIdle, m.HeapReleased, m.HeapInuse)
}

func printStats() {
	for {
		printHeapInfo()
		fmt.Printf("No of live Goroutines at time%d= %d\n", time.Now(), runtime.NumGoroutine())
		PrintMemUsage()
		time.Sleep(2 * time.Second)
	}
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func hiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}
