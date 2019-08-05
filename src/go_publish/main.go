package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const _defineInput string = "/images"
const _defineOutput string = "/pubspec.yaml"
const _begin string = "## <<assets begin>>"
const _end string = "## <<assets end>>"
const publishCli string = "go_publish"

//通用资源名头部
const _assetsTitle string = "- "

//应用初始数据
type Info struct {
	//命令执行路径
	pwd string
	//生成目标文件文件名
	output string
	//文件名列表
	files []string
	//资源路径
	input string
	//文件内容
	fileContent []string
}

//实例化初始数据
var info *Info = &Info{}

//获取输出文件路径
func (i *Info) GetOutput() (string, error) {
	path, err := filepath.Abs(i.pwd + i.output)
	return path, err
}

//获取输出文件路径
func (i *Info) GetOutputTest() (string, error) {
	path, err := filepath.Abs(i.pwd + "/pubspec1.yaml")
	return path, err
}

//获取列表第一个元素
func (i *Info) Shift() string {
	length := len(i.files)
	result := ""

	if length > 0 {
		result = i.files[0]
		i.files = i.files[1:]
	}

	return result
}

//文件缓冲
func (i *Info) Push(item string) {
	if i.fileContent == nil {
		i.fileContent = []string{}
	}
	i.fileContent = append(i.fileContent, item)
}

func main() {
	//初始数据初始化

	args := os.Args
	pwd, _ := os.Getwd()

	length := len(args)

	info.output = _defineOutput
	info.input = _defineInput
	info.pwd = pwd
	info.files = []string{"  assets:\n"}

	switch length {
	case 2:
		info.input = "/" + args[1]
	case 3:
		info.input = "/" + args[1]
		info.output = "/" + args[2]
	}

	fmt.Printf("初始化发布配置:\n%+v\n", info)
	fmt.Printf("命令参数:%+v\n", args)
	fmt.Printf("命令执行路径:%+v\n", pwd)

	dir, e := filepath.Abs(info.pwd + info.input)
	if e != nil {
		fmt.Printf("path error:%+v", e)
	}
	fmt.Printf("遍历文件根路径:%+v\n", dir)

	dirErr := readDir(filepath.Join(info.pwd, info.input))
	if dirErr != nil {
		return
	}

	fmt.Printf("结果文件列表:%+v\n", info.files)

	mErr := makeFileHandler()
	if mErr == nil {
		writeFileHandler()
	}
}

//遍历资源目录生成文件列表
func readDir(dirPath string) error {
	fList, e := ioutil.ReadDir(dirPath)

	if e != nil {
		fmt.Printf("资源路径不存在： %+v\n", e)
		return e
	}

	for _, f := range fList {
		if f.IsDir() {
			dir, e := filepath.Abs(dirPath + "/" + f.Name())
			if e != nil {
				fmt.Println("Dir read error")
			}
			fmt.Printf("PATH[%s]\n", dir)
			readDir(dir)
		} else {
			dir, e := filepath.Rel(info.pwd, dirPath+"/"+f.Name())
			//文件路径列表
			if e != nil {
				fmt.Println("File read error")
			}
			if f.Name() != ".DS_Store" {
				info.files = append(info.files, fileNameParse(dir))
				fmt.Printf("FILE[%+s]\n", dir)
			}
		}
	}

	return nil
}

func makeFileHandler() error {
	path, pErr := info.GetOutput()
	fmt.Printf("输出文件路径 %+v\n", path)
	if pErr != nil {
		fmt.Printf("File path error %+v\n", pErr)
		return pErr
	}

	file, fErr := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0600)
	if fErr != nil {
		fmt.Printf("File read or create error %+v\n", fErr)
		return fErr
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		buf, err := reader.ReadBytes('\n')

		str := string(buf)
		info.Push(str)
		// fmt.Printf("%+v", str)
		if err == io.EOF {
			fmt.Printf("文件读取完毕: %+v\n", err)
			break
		}
	}

	return nil
}

//写入文件处理
func writeFileHandler() error {
	// fmt.Printf("文件缓冲:\n%+v", info.fileContent)
	if info.fileContent == nil {
		return errors.New("目标文件为空，解析失败")
	}
	var start, end int
	for idx, va := range info.fileContent {
		if strings.Contains(va, _begin) {
			start = idx
		}
		if strings.Contains(va, _end) {
			end = idx
		}
	}

	t := append([]string{}, info.fileContent[end:]...)
	p := append(info.fileContent[0:start+1], info.files...)
	p = append(p, t...)

	path, pErr := info.GetOutput()
	if pErr != nil {
		fmt.Printf("File path error %+v\n", pErr)
		return pErr
	}
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Printf("os.OpenFile path error %+v\n", err)
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, str := range p {
		writer.WriteString(str)
	}
	writer.Flush()

	fmt.Printf("结果文件值\n %+v\n", p)
	return nil
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

//资源名处理
func fileNameParse(name string) string {
	return "    " + _assetsTitle + name + "\n"
}
