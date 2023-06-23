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

var regexInt = regexp.MustCompile(`\d+`)
var regexBrackets = regexp.MustCompile(`\([^)]*\)`)

func WatchGoroutine(header string) {
	stackSlice := make([]byte, 4096*2)
	s := runtime.Stack(stackSlice, true)

	type outputAnalysis struct {
		goroutineId string
		funcStack   []string
		createdBy   string
		node        *gtree.Node
	}

	var oaArray []outputAnalysis
	blocks := strings.Split(string(stackSlice[0:s]), "\n\n")
	for _, block := range blocks {
		var oa outputAnalysis
		target := true
		lines := strings.Split(block, "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "goroutine ") {
				oa.goroutineId = regexInt.FindString(line)
			} else if strings.HasSuffix(line, ")") {
				oa.funcStack = append(oa.funcStack, regexBrackets.ReplaceAllString(line, ""))
			} else if strings.HasPrefix(line, "created by ") {
				if strings.Contains(line, "gtree") {
					target = false
					break
				}
				oa.createdBy, _ = strings.CutPrefix(line, "created by ")
			}
		}
		if target {
			oaArray = append(oaArray, oa)
		}
	}

	sort.Slice(oaArray, func(i, j int) bool {
		return oaArray[i].goroutineId < oaArray[j].goroutineId
	})

	var root *gtree.Node
	root = gtree.NewRoot("goroutine 1 (main goroutine)")
	oaArray[0].node = root

	for i := 0; i < len(oaArray); i++ {
		for _, v := range oaArray[i].funcStack {
			for j := 0; j < len(oaArray); j++ {
				if v == oaArray[j].createdBy {
					oaArray[j].node = oaArray[i].node.Add("goroutine " + oaArray[j].goroutineId + " (created by " + oaArray[j].createdBy + ")")
				}
			}
		}
	}

	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	w.WriteString(header + "\n")

	err := gtree.OutputProgrammably(w, root)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b.String())
}
