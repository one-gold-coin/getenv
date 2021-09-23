package getenv

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var env *GetEnv

func GetVal(key string) *ValHandel {
	return env.GetVal(key)
}

type GetEnv struct {
	filePath string
	valMap   map[string]string
}

func (ge *GetEnv) Init() {
	if !ge.checkFileExist() {
		fmt.Println("env file no exist")
		os.Exit(0)
	}
	err := ge.readFile()
	if err != nil {
		fmt.Println("init env file fail: ", err.Error())
		os.Exit(0)
	}
	env = ge
}

func (ge *GetEnv) SetFilePath(filePath string) *GetEnv {
	ge.filePath = filePath
	return ge
}

func (ge *GetEnv) getFilePath() string {
	return ge.filePath
}

func (ge *GetEnv) readFile() error {
	file, err := os.Open(ge.getFilePath())
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	scanner := bufio.NewScanner(file)
	ge.valMap = make(map[string]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		line, _ = ge.filterLineNotes(line, "#")
		lineList := strings.Split(line, "=")
		if len(lineList) == 2 {
			ge.valMap[strings.TrimSpace(lineList[0])] = strings.TrimSpace(lineList[1])
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func (ge *GetEnv) filterLineNotes(content string, filter string) (str string, err error) {
	lineIndex := strings.Index(content, filter)
	if lineIndex < 0 {
		return content, nil
	}
	firstContent := content[0:lineIndex]
	firstContent = strings.Replace(firstContent, " ", "", -1)
	if len(firstContent) > 0 {
		str = content[0:lineIndex] + "\n"
	} else {
		str = ""
	}
	return str, nil
}

func (ge *GetEnv) GetVal(key string) *ValHandel {
	key = strings.TrimSpace(key)
	getVal := &ValHandel{}
	if val, isOk := ge.valMap[key]; isOk {
		return getVal.setVal(val)
	}
	return getVal.setVal("")
}

func (ge *GetEnv) checkFileExist() bool {
	_, err := os.Stat(ge.getFilePath())
	if os.IsNotExist(err) {
		return false
	}
	return true
}

type ValHandel struct {
	val string
}

func (ge *ValHandel) setVal(val string) *ValHandel {
	ge.val = strings.TrimSpace(val)
	return ge
}

func (ge *ValHandel) String() string {
	return ge.val
}

func (ge *ValHandel) StrSlice() []string {
	valList := strings.Split(ge.val, ",")
	return valList
}

func (ge *ValHandel) Int() int {
	if ge.val == "" {
		return 0
	}
	valInt, _ := strconv.Atoi(ge.val)
	return valInt
}

func (ge *ValHandel) IntSlice() []int {
	if ge.val == "" {
		return nil
	}
	valList := strings.Split(ge.val, ",")
	valIntList := make([]int, 0)
	for _, s := range valList {
		valInt, _ := strconv.Atoi(s)
		valIntList = append(valIntList, valInt)
	}
	return valIntList
}

func (ge *ValHandel) Int64() int64 {
	if ge.val == "" {
		return 0
	}
	valInt64, _ := strconv.ParseInt(ge.val, 10, 64)
	return valInt64
}

func (ge *ValHandel) Int64Slice() []int64 {
	if ge.val == "" {
		return nil
	}
	valList := strings.Split(ge.val, ",")
	valIntList := make([]int64, 0)
	for _, s := range valList {
		valInt64, _ := strconv.ParseInt(s, 10, 64)
		valIntList = append(valIntList, valInt64)
	}
	return valIntList
}
