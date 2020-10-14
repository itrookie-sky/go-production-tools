/**
pod 配置自定义文件
项目中配置文件 ==复制替换==> flutter 插件缓存路径配置
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

//搜索路径
const _defineSearch string = "./ios/podsconfigs/"

//插件配置路径
const _defineFile string = "./.flutter-plugins"

//Info 初始化信息
type Info struct {
	pwd string
	//目标插件路径映射
	dirMap map[string]string
	//项目内podspec文件
	localPods []string
}

//实例化数据
var info *Info = &Info{}

func main() {
	fmt.Printf("podspec searching...\n")
	pwd, _ := os.Getwd()
	info.pwd = pwd
	info.dirMap = make(map[string]string)
	info.localPods = []string{}

	fmt.Printf("执行路径:%+v\n", info.pwd)

	parsePluginFile()
	parseReadDir()
	replaceFile()
}

//遍历本地项目配置
func parseReadDir() {
	fList, e := ioutil.ReadDir(_defineSearch)

	if e != nil {
		fmt.Printf("资源路径不存在: %+v\n", e)
	}

	for _, f := range fList {
		if f.IsDir() {
		} else {
			if f.Name() != ".DS_Store" {
				info.localPods = append(info.localPods, f.Name())
				fmt.Printf("遍历配置文件: %+v\n", f.Name())
			}
		}
	}
}

//读取插件文件生成插件路径列表
func parsePluginFile() error {
	isExists, err := pathExists(_defineFile)
	if !isExists || err != nil {
		fmt.Println("插件配置文件不存在")
		return err
	}
	file, fErr := os.OpenFile(_defineFile, os.O_RDONLY, 0600)
	if fErr != nil {
		fmt.Printf("File read error %+v\n", fErr)
		return fErr
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		buf, err := reader.ReadBytes('\n')
		str := string(buf)
		plugin := strings.Split(str, "=")
		fmt.Printf("插件名称路径映射: %+v\n", plugin)
		if len(plugin) == 2 {
			info.dirMap[plugin[0]] = strings.Replace(plugin[1], "\n", "", 1)
		}

		if err == io.EOF {
			fmt.Printf("文件读取完毕: %+v\n", err)
			break
		}
	}

	return nil
}

func replaceFile() {
	if len(info.localPods) < 1 {
		fmt.Println("项目内配置文件为空")
		return
	}

	for _, v := range info.localPods {
		nameArr := strings.Split(v, ".")
		plugin := info.dirMap[nameArr[0]]
		if plugin != "" {
			fileName := _defineSearch + v
			targetName := plugin + "ios/" + v

			file, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0600)
			defer file.Close()

			var reader *bufio.Reader = bufio.NewReader(file)
			buf, _ := ioutil.ReadAll(reader)

			fmt.Printf("执行替换操作: \n%+v \n%+v\n", fileName, targetName)
			ioutil.WriteFile(targetName, buf, 0600)
		}
	}
}

//检测文件是否存在
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
