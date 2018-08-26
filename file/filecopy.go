package main

import (
	"flag"
	"os"
	"fmt"
	"bufio"
	"strings"
	"io"
)

//判断文件是否存在
func fileExists(filename string) bool{
	_, err :=os.Stat(filename)
	return err == nil || os.IsExist(err)
}

//拷贝文件
func copyFile(src,dst string)(w int64, err error){
	srcFile, err :=os.Open(src)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	defer srcFile.Close()
	dstFile, err :=os.Create(dst)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}

//拷贝文件的行为
func copyFileAction(src, dst string, showProgress, force bool){
	if !force{
		if fileExists(dst){ //如果目标文件存在，就确认是否覆盖
			fmt.Printf("%s exists override?y/n\n", dst)
			reader := bufio.NewReader(os.Stdin)//提取用户命令行输入
			data, _, _ := reader.ReadLine()
			if strings.ToLower(strings.TrimSpace(string(data))) != "y"{

			}
		}
	}

	copyFile(src, dst)
	if showProgress{
		fmt.Printf("'%s' -> '%s'\n", src, dst)
	}

}

//模拟实现文件拷贝函数
func main() {

	//接收两个bool参数
	var showProgress, force bool
	flag.BoolVar(&force, "f",false, "force copy when existing")
	flag.BoolVar(&showProgress, "v", false, "explain what is being done")

	//解析参数
	flag.Parse()

	//非法参数检测
	if flag.NArg() < 2{
		flag.Usage() //打印命令说明
		return
	}

	copyFileAction(flag.Arg(0), flag.Arg(1), showProgress, force)

}
