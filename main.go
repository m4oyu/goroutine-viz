package visualizationtool

import (
	"os"
	"regexp"
	"runtime"
	"sort"
	"strings"

	"github.com/ddddddO/gtree"
)

var regexInt = regexp.MustCompile(`\d+`)
var regexBrackets = regexp.MustCompile(`\([^)]*\)`)

func WatchGoroutine() error {
	stackSlice := make([]byte, 2048)
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
		lines := strings.Split(block, "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "goroutine ") {
				oa.goroutineId = regexInt.FindString(line)
			} else if strings.HasSuffix(line, ")") {
				oa.funcStack = append(oa.funcStack, regexBrackets.ReplaceAllString(line, ""))
			} else if strings.HasPrefix(line, "created by ") {
				oa.createdBy, _ = strings.CutPrefix(line, "created by ")
			}
		}
		oaArray = append(oaArray, oa)
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

	if err := gtree.OutputProgrammably(os.Stdout, root); err != nil {
		return err
	}

	return nil

}
