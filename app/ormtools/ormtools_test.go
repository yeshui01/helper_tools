package ormtools

import (
	"testing"

	"github.com/sirupsen/logrus"
)

var connStr string = "root:123456@tcp(192.168.13.26:3306)/logs_myth?charset=utf8mb4&parseTime=True&loc=Local"

func TestReadTableStruct(t *testing.T) {
	InitDB(connStr)
	logrus.Debug("TestReadTableStruct")
	ReadTableStruct("log_test")
}
