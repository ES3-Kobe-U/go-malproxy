package logic

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//外部コマンドを実行する関数
func AssignLinuxCmdHandler(password string) error {
	//ルート権限で実行する。
	cmd := exec.Command("sudo", "-S", "command")
	cmd.Stdin = strings.NewReader(password + "\n")
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println("Err:", err)
	}

	//ルート証明書を外部コマンドでインストールさせる。
	cmdMkdir := exec.Command("sudo", "mkdir", "/usr/local/share/ca-certificates/extra")
	cmdMkdir.Stderr = os.Stderr
	cmdMkdir.Stdin = os.Stdin

	out, err := cmdMkdir.Output()
	if err != nil {
		fmt.Println("Err", err)
		return err
	} else {
		fmt.Println("mkdir success:", string(out))
	}

	cmdCp := exec.Command("sudo", "cp", "foo.crt", "/usr/local/share/ca-certificates/extra/foo.crt")
	cmdCp.Stderr = os.Stderr
	cmdCp.Stdin = os.Stdin

	out, err = cmdCp.Output()
	if err != nil {
		fmt.Println("Err", err)
		return err
	} else {
		fmt.Println("cp success:", string(out))
	}

	cmdDpkg := exec.Command("sudo", "update-ca-certificates")
	cmdDpkg.Stderr = os.Stderr
	cmdDpkg.Stdin = os.Stdin

	out, err = cmdDpkg.Output()
	if err != nil {
		fmt.Println("Err", err)
		return err
	} else {
		fmt.Println("update success:", string(out))
	}
	return nil
}
