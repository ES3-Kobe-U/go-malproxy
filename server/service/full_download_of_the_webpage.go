package service

import (
	"log"
	"os/exec"
)

func FullDownloadOfTheWebPage(URL string) error {
	log.Println("処理開始")
	cmd := exec.Command("goscrape", URL)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("実行中...")
	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("実行終了")
	return nil
}
