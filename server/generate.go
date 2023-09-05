package server

import (
	"io"
	"log"
	"model_generate/consts"
	"os"
	"os/exec"
	"text/template"
)

var OrmMap map[string]*template.Template

var helpers template.FuncMap

// GenStruct 使用策略模式生成代码
func GenStruct(ormName string) {
	t, ok := OrmMap[ormName]
	if !ok {
		t = OrmMap[consts.DefaultTmpl]
	}

	// write to file
	file, err := os.OpenFile("model.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)

	err = t.Execute(io.Writer(file), TmplContext)
	if err != nil {
		log.Fatal(err.Error())
	}

	// format go code
	goFormatFile("model.go")
}

func init() {
	OrmMap = make(map[string]*template.Template)
	helpers = make(template.FuncMap)
}

func RegisterHelper(funcName string, fc any) {
	helpers[funcName] = fc
}

func goFormatFile(fileName string) {
	cmd := exec.Command("go", "fmt", fileName)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
