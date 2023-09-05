package server

import (
	"model_generate/consts"
	"model_generate/utils"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

func ToCamel(s string) (string, error) {
	var words []string
	for _, word := range strings.Split(s, "_") {
		word, _ = UpperFirst(word)
		words = append(words, word)
	}
	return strings.Join(words, ""), nil
}

// UpperFirst 将首字母大写
func UpperFirst(s string) (string, error) {
	if len(s) == 0 {
		return s, nil
	}
	return strings.ToUpper(s[:1]) + s[1:], nil
}

func ConvertType(mysqlType string) (string, error) {
	switch mysqlType {
	case "tinyint", "smallint", "mediumint", "int", "integer":
		return "int", nil
	case "bigint":
		return "int64", nil
	case "float", "double", "decimal":
		return "float64", nil
	case "date", "datetime", "timestamp":
		return "time.Time", nil
	case "char", "varchar", "text", "tinytext", "mediumtext", "longtext":
		return "string", nil
	case "binary", "varbinary", "blob", "tinyblob", "mediumblob", "longblob":
		return "[]byte", nil
	default:
		return "interface{}", nil
	}
}

func JsonField(s string) (string, error) {
	s, _ = ToCamel(s)
	return LowerFirst(s), nil
}

// LowerFirst 将首字母小写
func LowerFirst(s string) string {
	if len(s) == 0 {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// GetXormFieldInfo 获取Xorm格式的字段信息
func GetXormFieldInfo(col ColumnInfo) string {
	res := make([]string, 0)
	if col.Null == "NO" {
		res = append(res, "notnull")
	}

	if col.Key != "" {
		switch col.Key {
		case "PRI":
			res = append(res, "pk")
			break
		}
	}

	if col.Extra == "auto_increment" {
		res = append(res, "autoincr")
	}

	result := strings.TrimSpace(strings.Join(res, " "))
	if result != "" {
		result = " " + result
	}
	return result
}

func init() {
	RegisterHelper("ToCamel", ToCamel)
	RegisterHelper("UpperFirst", UpperFirst)
	RegisterHelper("ConvertType", ConvertType)
	RegisterHelper("JsonField", JsonField)
	RegisterHelper("GetXormFieldInfo", GetXormFieldInfo)

	tmplFiles := utils.GetAllTmplFiles()
	for _, fileName := range tmplFiles {
		filePrefix := path.Base(fileName)[:len(fileName)-len(path.Ext(fileName))]
		t := template.New(fileName).Funcs(helpers)
		fileName = filepath.Join(consts.TmplDir, fileName)
		t = template.Must(t.ParseFiles(fileName))
		OrmMap[filePrefix] = t
	}
}
