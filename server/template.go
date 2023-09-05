package server

import (
	"database/sql"
	"log"
	"model_generate/utils"
)

type (
	TemplateContext struct {
		DbName    string
		TableInfo []*TableInfo
	}
	ColumnInfo struct {
		Field     string `db:"Field"`
		Type      string `db:"Type"`
		ShortType string
		Null      string         `db:"Null"`
		Key       string         `db:"Key"`
		Default   sql.NullString `db:"Default"`
		Extra     string         `db:"Extra"`
	}
	TableInfo struct {
		TableName string
		Columns   []ColumnInfo
	}
)

var (
	db          *sql.DB
	TmplContext *TemplateContext
)

func Init(dns string) {
	TmplContext = new(TemplateContext)
	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatal(err)
		log.Printf("Please watch the datasource format: %s", "root:password@tcp(127.0.0.1:3306)/testdb?charset=utf8")
	}
	defer db.Close()

	// 获取库中所有表名
	tables, err := db.Query("SHOW TABLES")
	if err != nil {
		panic(err.Error())
	}

	var tableName string
	for tables.Next() {
		_ = tables.Scan(&tableName)
		columnInfos := buildSchema(db, tableName)
		TmplContext.TableInfo = append(TmplContext.TableInfo, &TableInfo{
			TableName: tableName,
			Columns:   columnInfos,
		})
	}
}

// 获取元数据
func buildSchema(db *sql.DB, tableName string) []ColumnInfo {
	rows, err := db.Query("SHOW COLUMNS FROM " + tableName)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var columns []ColumnInfo
	for rows.Next() {
		var column ColumnInfo
		err := rows.Scan(
			&column.Field,
			&column.Type,
			&column.Null,
			&column.Key,
			&column.Default,
			&column.Extra,
		)
		if err != nil {
			log.Fatal(err)
		}
		column.ShortType = utils.RemoveTypeLength(column.Type)
		columns = append(columns, column)
	}
	return columns
}
