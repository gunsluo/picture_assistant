package main

import (
	"fmt"

	"github.com/gunsluo/picture_assistant"
)

func main() {
	demo()
}

func demo() {

	url := "https://www.jerrylou.me/images/avatar.jpg"
	read := assistant.NewSpiderPictureRead(url)
	write := assistant.NewFilePictureWrite("./")
	convert := assistant.NewResizeConvert().SetRatio(true).SetWidth(500)

	task := assistant.NewTask(read).Pipe(convert, write)
	//task := assistant.NewTask(read).Pipe(assistant.NullPictureConvert, write)
	infos, err := task.Exec()
	if err != nil {
		panic(err)
	}

	for _, info := range infos {
		fmt.Printf("save sucess. %#v\n", info)
	}
}
