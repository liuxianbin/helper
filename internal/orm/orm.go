package orm

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DBEngine struct {
	db *sql.DB
}

type Column struct {
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnComment string
}

type DBInfo struct {
	DbType   string
	Host     string
	Username string
	Password string
	Charset  string
}

func Open(d DBInfo) (*DBEngine, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/information_schema?charset=%s&parseTime=True&loc=Local",
		d.Username, d.Password, d.Host, d.Charset,
	)
	db, err := sql.Open(d.DbType, dsn)
	if err != nil {
		return nil, err
	}
	return &DBEngine{
		db,
	}, nil
}

func (e *DBEngine) GetColumns(dbName, tableName string) ([]*Column, error) {
	sql := `
		SELECT
			COLUMN_NAME,
			DATA_TYPE,
			IS_NULLABLE,
			COLUMN_COMMENT
		FROM COLUMNS
		WHERE TABLE_NAME = ?
		AND TABLE_SCHEMA = ?
		ORDER BY ORDINAL_POSITION
	`
	rows, err := e.db.Query(sql, tableName, dbName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var columns []*Column
	exists := false
	for rows.Next() {
		var column Column
		err = rows.Scan(&column.ColumnName, &column.DataType, &column.IsNullable, &column.ColumnComment)
		if err != nil {
			return nil, err
		}
		column.DataType = sqlTypeToGoType[column.DataType]
		columns = append(columns, &column)
		exists = true
	}
	if !exists {
		return nil, errors.New("无数据")
	}
	return columns, nil
}
