package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main(){

	file ,err := os.Open("/Users/lilin/dev/CODE/go/go_study/file/file.go")

	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	var buffer  [] byte
	var tempBuffer = make([]byte, 128)


	for {
		_,err := file.Read(tempBuffer)
		if err == io.EOF {
			break
		}
		buffer = append(buffer, tempBuffer...) //拼接缓存
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Printf("读取内容为%v", string(buffer))



	//write

	ff, err := os.OpenFile("./test", os.O_CREATE|os.O_RDWR, 0666)
	defer  ff.Close()

	if err != nil {
		fmt.Println("写文件失败",err)
	}

	ff.WriteString("hello world \n i'm ok")

	writer := bufio.NewWriter(ff)

	writer.WriteString("\nthis is from bufio writing ")
	writer.Flush()
}
