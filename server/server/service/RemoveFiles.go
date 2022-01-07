package service

import (
	"log"
	"os"
	"os/exec"
)

//作成したHTMLファイルを削除する。
func RemoveFile(fileName string) error {
	cmdRemove := exec.Command("rm", "test/"+fileName+".html")
	cmdRemove.Stderr = os.Stderr
	cmdRemove.Stdin = os.Stdin

	out, err := cmdRemove.Output()
	if err != nil {
		log.Println("Err", err)
		return err
	} else {
		log.Println("remove files --> success:", string(out))
	}
	return nil
}
