package visualizationtool

import "runtime"

type traceNode struct {
	parent *traceNode
	child  []*traceNode
}

func WatchGoroutine() (output, debugOutput []byte) {
	stackSlice := make([]byte, 2048)
	s := runtime.Stack(stackSlice, true)

	return stackSlice[0:s], stackSlice[0:s]

	// 	goroutine 18 [running]:
	// main.stackExample.func1()
	// 	/home/m4oyu/go/src/github/m4oyu/visualizationTool/runtimeStack/main.go:28 +0x13e
	// created by main.stackExample
	// 	/home/m4oyu/go/src/github/m4oyu/visualizationTool/runtimeStack/main.go:14 +0x135

	// goroutine 1 [sleep]:
	// time.Sleep(0x3b9aca00)
	// 	/usr/local/go/src/runtime/time.go:195 +0x135
	// main.stackExample()
	// 	/home/m4oyu/go/src/github/m4oyu/visualizationTool/runtimeStack/main.go:33 +0x13f
	// main.main()
	// 	/home/m4oyu/go/src/github/m4oyu/visualizationTool/runtimeStack/main.go:55 +0x57

	// goroutine 19 [semacquire]:
	// runtime.Stack({0xc0000a8000, 0x800, 0x800}, 0x1)
	// 	/usr/local/go/src/runtime/mprof.go:1193 +0x4d
	// main.stackExample.func1.1()
	// 	/home/m4oyu/go/src/github/m4oyu/visualizationTool/runtimeStack/main.go:17 +0x45
	// created by main.stackExample.func1
	// 	/home/m4oyu/go/src/github/m4oyu/visualizationTool/runtimeStack/main.go:16 +0xad

}
