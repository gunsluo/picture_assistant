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
	write.SetNameGenerator(assistant.BulrNameGenerator)
	write2 := assistant.NewB2PictureWrite(accountID, applicationKey, bucketName)
	convert := assistant.NewResizeConvert().SetRatio(true).SetWidth(500)

	task := assistant.NewTask(read).Pipe(convert, write).Pipe(assistant.NullPictureConvert, write2)
	infos, err := task.Exec()
	if err != nil {
		panic(err)
	}

	for _, info := range infos {
		fmt.Printf("save sucess. %#v\n", info)
	}
}
