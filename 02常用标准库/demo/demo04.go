package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "写入内容")
	/**os.OpenFIle是GO语言中用于打开文件的函数之一
	*func OpenFile(name string, flag int, perm FileMode) (*File, error)
	* name 文件路径
	* flag 文件标志，用于文件打开方式，如os.O_RDONLY（只读）、os.O_WRONLY（只写）、os.O_RDWR（读写），os.O_CREATE（创建文件，如果文件不存在）、os.O_APPEND（追加写入）等
	*perm：创建文件时的权限（文件模式），只在使用 os.O_CREATE 标志时才有效。可以使用 os.FileMode 类型的常量来表示文件的权限，如 0644
	*OpenFile 函数返回一个 *os.File 类型的文件对象和一个错误。如果打开文件成功，将返回一个文件对象，可以通过该对象进行读取或写入等操作。如果发生错误（如文件不存在或权限问题），将返回相应的错误信息
	*/
	fileObj, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 6044)
	if err != nil {
		fmt.Println("文件错误", err)
		return
	}
	name := "zss"
	fmt.Fprintf(fileObj, "写入内容：%s", name)
}
