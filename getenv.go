package getenv

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var env *GetEnv

type GetEnv struct {
	filePath string
	valMap   map[string]string
}

func (t *GetEnv) Init() {
	if !t.checkFileExist() {
		fmt.Println("env file no exist")
		os.Exit(0)
	}
	err := t.readFile()
	if err != nil {
		fmt.Println("init env file fail: ", err.Error())
		os.Exit(0)
	}
	env = t
}

func (t *GetEnv) SetFilePath(filePath string) *GetEnv {
	t.filePath = filePath
	return t
}

func (t *GetEnv) getFilePath() string {
	return t.filePath
}

func (t *GetEnv) readFile() error {
	file, err := os.Open(t.getFilePath())
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	scanner := bufio.NewScanner(file)
	t.valMap = make(map[string]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		line, _ = t.filterLineNotes(line, "#")
		lineList := strings.Split(line, "=")
		if len(lineList) == 2 {
			t.valMap[strings.TrimSpace(lineList[0])] = strings.TrimSpace(lineList[1])
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func (t *GetEnv) filterLineNotes(content string, filter string) (str string, err error) {
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

func (t *GetEnv) GetVal(key string) *ValHandel {
	key = strings.TrimSpace(key)
	getVal := &ValHandel{}
	if val, isOk := t.valMap[key]; isOk {
		return getVal.setVal(val)
	}
	return getVal.setVal("")
}

func (t *GetEnv) checkFileExist() bool {
	_, err := os.Stat(t.getFilePath())
	if os.IsNotExist(err) {
		return false
	}
	return true
}

type ValHandel struct {
	val string
}

func (t *ValHandel) setVal(val string) *ValHandel {
	t.val = strings.TrimSpace(val)
	return t
}

func (t *ValHandel) String() string {
	return t.val
}

func (t *ValHandel) StrSlice() []string {
	valList := strings.Split(t.val, ",")
	return valList
}

func (t *ValHandel) Int() int {
	valInt, _ := strconv.Atoi(t.val)
	return valInt
}

func (t *ValHandel) IntSlice() []int {
	valList := strings.Split(t.val, ",")
	valIntList := make([]int, 0)
	for _, s := range valList {
		valInt, _ := strconv.Atoi(s)
		valIntList = append(valIntList, valInt)
	}
	return valIntList
}

func (t *ValHandel) Int64() int64 {
	valInt64, _ := strconv.ParseInt(t.val, 10, 64)
	return valInt64
}

func (t *ValHandel) Int64Slice() []int64 {
	valList := strings.Split(t.val, ",")
	valIntList := make([]int64, 0)
	for _, s := range valList {
		valInt64, _ := strconv.ParseInt(s, 10, 64)
		valIntList = append(valIntList, valInt64)
	}
	return valIntList
}

func GetVal(key string) *ValHandel {
	return env.GetVal(key)
}
