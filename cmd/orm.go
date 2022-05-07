package cmd

import (
	"github.com/spf13/cobra"
	"helper/internal/orm"
	"log"
)

var ormCmd = &cobra.Command{
	Use:   "orm",
	Short: "SQL转换Golang",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := orm.DBInfo{
			DbType:   dbType,
			Host:     host,
			Username: username,
			Password: password,
			Charset:  charset,
		}
		db, err := orm.Open(dbInfo)
		if err != nil {
			log.Fatalf("Connect err: %s", err.Error())
		}
		data, err := db.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("GetColumns err: %s", err.Error())
		}
		err = orm.Generate(tableName, data)
		if err != nil {
			log.Fatalf("Generate err: %s", err.Error())
		}
	},
}

var (
	username  string
	password  string
	charset   string
	host      string
	dbType    string
	dbName    string
	tableName string
)

func init() {
	ormCmd.Flags().StringVarP(&username, "username", "", "root", "用户名")
	ormCmd.Flags().StringVarP(&password, "password", "", "", "密码")
	ormCmd.Flags().StringVarP(&charset, "charset", "", "utf8", "字符编码")
	ormCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "HOST")
	ormCmd.Flags().StringVarP(&dbType, "type", "", "mysql", "数据库类型")
	ormCmd.Flags().StringVarP(&dbName, "database", "", "", "数据库名称")
	ormCmd.Flags().StringVarP(&tableName, "table", "", "", "表名称")
}
