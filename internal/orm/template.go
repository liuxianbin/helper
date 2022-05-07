package orm

import (
	"helper/internal/str"
	"os"
	"text/template"
)

const tmp = `
type {{.TableName|toCamelCase}} struct{
{{range .Columns}}    {{.ColumnName|toCamelCase}}    {{.DataType}} {{if gt (len .ColumnComment) 0}}    // {{.ColumnComment}}{{end}}
{{end}}}
`

// DTO 输出模板
type DTO struct {
	TableName string
	Columns   []*Column
}

func Generate(tableName string, data []*Column) error {
	t := template.Must(template.New("").Funcs(template.FuncMap{"toCamelCase": str.UnderscoreToUpperCamelCase}).Parse(tmp))
	return t.Execute(os.Stdout, DTO{
		TableName: tableName,
		Columns:   data,
	})
}
