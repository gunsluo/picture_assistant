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

	accountID := ""
	applicationKey := ""
	bucketName := ""
	write := assistant.NewB2PictureWrite(accountID, applicationKey, bucketName)
	write.SetNameGenerator(assistant.AutoNameGenerator)

	task := assistant.NewTask(read).Pipe(assistant.DefaultJPGToWEBPConvert, write)
	infos, err := task.Exec()
	if err != nil {
		panic(err)
	}

	for _, info := range infos {
		fmt.Printf("save sucess. %#v\n", info)
	}
}
