package config

import (
	"applyFK/holiday"
	"fmt"
	"os"
	"time"

	"github.com/go-yaml/yaml"
)

type Config struct {
	Url            string `yaml:"url"`
	Token          string `yaml:"token"`
	Ticket         string `yaml:"ticket"`
	ApplyPhone     string `yaml:"applyPhone"`
	QuickApplyType string `yaml:"quickApplyType"`
	QuickApplyNum  int    `yaml:"quickApplyNum"`
	ApplyDays      int    `yaml:"applyDays"`
}

var Conf *Config
var QuickApplyTypeMap map[string]bool = map[string]bool{
	"week":  true,
	"month": true,
	"year":  true,
}

func LoadConfig(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("配置文件不存在: %v", err)
	}
	config := &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return fmt.Errorf("配置文件解析失败: %v", err)
	}
	if len(config.Url) == 0 {
		return fmt.Errorf("配置文件中url不能为空")
	}
	if len(config.QuickApplyType) != 0 && !QuickApplyTypeMap[config.QuickApplyType] {
		return fmt.Errorf("配置文件中quickApplyType错误，只能为week, month, year")
	}
	if config.QuickApplyNum <= 0 {
		config.QuickApplyNum = 0
	}
	if config.QuickApplyType == "week" {
		config.ApplyDays = holiday.GetCustomWeekDays(config.QuickApplyNum)
	} else if config.QuickApplyType == "month" {
		config.ApplyDays = holiday.CalculateWorkDays(config.QuickApplyNum)
	} else if config.QuickApplyType == "year" {
		config.ApplyDays = holiday.CalculateWorkDaysToEndOfYear(time.Now())
	}
	Conf = config
	fmt.Println("配置文件加载成功～")
	return nil
}
