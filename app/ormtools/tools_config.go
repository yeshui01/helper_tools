/* ====================================================================
 * Author           : tianyh(mknight)
 * Email            : 824338670@qq.com
 * Last modified    : 2022-03-30 11:33
 * Filename         : tools_config.go
 * Description      :
 * ====================================================================*/
package ormtools

type ToolsConfig struct {
	OrmGen struct {
		OrmPath string `yaml:"ormPath"`
		PbPath  string `yaml:"pbPath"`
		TbPath  string `yaml:"tbPath"`
		PkgPath string `yaml:"pkgPath"`
	} `yaml:"ormGen"`
	MysqlConnStr string `yaml:"mysqlConnStr"`
}

func NewToolsConfig() *ToolsConfig {
	return &ToolsConfig{}
}
