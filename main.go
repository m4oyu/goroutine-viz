package visualizationtool

import (
	"bufio"
	"bytes"
	"fmt"
	"regexp"
	"runtime"
	"sort"
	"strings"

	"github.com/ddddddO/gtree"
)

type goroutineNode struct {
	goroutineId string
	funcStack   []string
	createdBy   string
	tree        *gtree.Node
}

var (
	regexInt      = regexp.MustCompile(`\d+`)
	regexBrackets = regexp.MustCompile(`\([^)]*\)`)
)

func parseText(blocks []string) []goroutineNode {
	var goroutineList []goroutineNode

	for _, block := range blocks {
		var gNode goroutineNode
		isTarget := true
		lines := strings.Split(block, "\n")

		for _, line := range lines {
			if strings.HasPrefix(line, "goroutine ") {
				gNode.goroutineId = regexInt.FindString(line)
			} else if strings.HasSuffix(line, ")") {
				gNode.funcStack = append(gNode.funcStack, regexBrackets.ReplaceAllString(line, ""))
			} else if strings.HasPrefix(line, "created by ") {
				if strings.Contains(line, "gtree") {
					isTarget = false
					break
				}
				gNode.createdBy, _ = strings.CutPrefix(line, "created by ")
			}
		}
		if isTarget {
			goroutineList = append(goroutineList, gNode)
		}
	}

	sort.Slice(goroutineList, func(i, j int) bool {
		return goroutineList[i].goroutineId < goroutineList[j].goroutineId
	})

	return goroutineList
}

func buildGoroutineTree(goroutineList []goroutineNode) *gtree.Node {
	var root *gtree.Node
	root = gtree.NewRoot("goroutine 1 (main goroutine)")
	goroutineList[0].tree = root

	for i := 0; i < len(goroutineList); i++ {
		for _, v := range goroutineList[i].funcStack {
			for j := 0; j < len(goroutineList); j++ {
				if v == goroutineList[j].createdBy {
					goroutineList[j].tree = goroutineList[i].tree.Add("goroutine " + goroutineList[j].goroutineId + " (created by " + goroutineList[j].createdBy + ")")
				}
			}
		}
	}
	return root
}

func WatchGoroutine(header string) {
	stackSlice := make([]byte, 4096*2)
	s := runtime.Stack(stackSlice, true)
	blocks := strings.Split(string(stackSlice[0:s]), "\n\n")

	goroutineList := parseText(blocks)

	root := buildGoroutineTree(goroutineList)

	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	w.WriteString(header + "\n")

	err := gtree.OutputProgrammably(w, root)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b.String())
}
