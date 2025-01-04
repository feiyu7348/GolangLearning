// Package CreationalPattern
// author: zfy  date: 2024/12/6
package CreationalPattern

import (
	"fmt"
	"sync"
)

// 单例模式 https://developer.aliyun.com/article/1267689
// 1、饿汉模式

type ConfigManager struct {
	configData map[string]string
}

var instance *ConfigManager = &ConfigManager{
	configData: make(map[string]string),
}

func GetConfigManager() *ConfigManager {
	return instance
}

func (cm *ConfigManager) Set(key, value string) {
	cm.configData[key] = value
}

func (cm *ConfigManager) Get(key string) string {
	return cm.configData[key]
}

// 2、懒汉模式

type Logger struct {
	logs []string
}

var instance1 *Logger
var once sync.Once

func GetLogger() *Logger {
	once.Do(func() {
		instance1 = &Logger{
			logs: []string{}}
	})
	return instance1
}

func (l *Logger) AddLog(log string) {
	l.logs = append(l.logs, log)
}

func (l *Logger) PrintLogs() {
	for _, log := range l.logs {
		fmt.Println(log)
	}
}
