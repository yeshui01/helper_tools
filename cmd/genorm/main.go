package main

import (
	"flag"
	"helpertools/app/ormtools"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var (
	tableName = flag.String("tableName", "none", "database table name")
)

func main() {
	flag.Parse()

	if *tableName == "none" {
		logrus.Error("tableName is none")
		return
	}

	logrus.SetFormatter(
		&logrus.TextFormatter{
			ForceColors:     false,
			TimestampFormat: "2006-01-02 15:04:05",
			// TimestampFormat: time.RFC3339,
		},
	)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Infof("gen orm table[%s] begin", *tableName)

	// 读取tools配置
	toolsConfig := ormtools.NewToolsConfig()
	content, err := ioutil.ReadFile("./ormtools.yaml")
	if err != nil {
		logrus.Errorf(err.Error())
		return
	}
	if err = yaml.Unmarshal(content, toolsConfig); err != nil {
		logrus.Fatalf("解析config.yaml出错: %v", err)
		return
	}
	if len(toolsConfig.OrmGen.OrmPath) < 1 {
		logrus.Error("orm gen path is none")
		return
	}
	if !ormtools.InitDB(toolsConfig.MysqlConnStr) {
		logrus.Error("connect mysql error!")
		return
	}
	fieldList, err := ormtools.ReadTableStruct(*tableName)
	if err != nil {
		logrus.Infof("fieldList Err:%s", err.Error())
		return
	} else {
		for _, v := range fieldList {
			logrus.Debugf("fieldName:%s, valueType:%s", v.GetFieldName(), v.GetFieldTypeName())
		}
	}

	if len(toolsConfig.OrmGen.OrmPath) > 0 {
		ormtools.GenGormObjectDefine(*tableName, fieldList, toolsConfig.OrmGen.OrmPath)
	} else {
		logrus.Error("OrmPath not find config")
	}
	if len(toolsConfig.OrmGen.PbPath) > 0 {
		ormtools.GenGormPbDefine(*tableName, fieldList, toolsConfig.OrmGen.PbPath)
	} else {
		logrus.Error("PbPath not find config")
	}
	if len(toolsConfig.OrmGen.TbPath) > 0 && len(toolsConfig.OrmGen.PkgPath) > 0 {
		ormtools.GenGormTableDefine(*tableName, fieldList, toolsConfig.OrmGen.TbPath, toolsConfig.OrmGen.PkgPath)
	} else {
		logrus.Error("TbPath or PkgPath not find config")
	}
	logrus.Infof("gen orm table[%s] end", *tableName)
}
