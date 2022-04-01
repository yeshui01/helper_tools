package ormtools

import "strings"

const (
	FieldTypeUnknown int8 = 0
	FieldTypeInt     int8 = 1
	FieldTypeBigint  int8 = 2
	FieldTypeString  int8 = 3
	FieldTypeBlob    int8 = 4
)

// 数据表字段
type TableField struct {
	fieldName  string
	fieldType  int8
	primaryKey bool
	autoInc    bool
}

func (tf *TableField) GetFieldName() string {
	return tf.fieldName
}
func (tf *TableField) GetFieldTypeName() string {
	return getValueTypeName(tf.fieldType)
}
func transValueType(t string) int8 {
	if strings.Count(t, "bigint") > 0 {
		return FieldTypeBigint
	} else if strings.Count(t, "int") > 0 {
		return FieldTypeInt
	} else if strings.Count(t, "text") > 0 {
		return FieldTypeString
	} else if strings.Count(t, "char") > 0 {
		return FieldTypeString
	} else if strings.Count(t, "blob") > 0 {
		return FieldTypeBlob
	}
	return FieldTypeUnknown
}

func getValueTypeName(t int8) string {
	switch t {
	case FieldTypeInt:
		{
			return "int32"
		}
	case FieldTypeBigint:
		{
			return "int64"
		}
	case FieldTypeString:
		{
			return "string"
		}
	case FieldTypeBlob:
		{
			return "[]byte"
		}
	}
	return "unknown type"
}
func getValueTypePbName(t int8) string {
	switch t {
	case FieldTypeInt:
		{
			return "int32"
		}
	case FieldTypeBigint:
		{
			return "int64"
		}
	case FieldTypeString:
		{
			return "string"
		}
	case FieldTypeBlob:
		{
			return "bytes"
		}
	}
	return "unknown type"
}
