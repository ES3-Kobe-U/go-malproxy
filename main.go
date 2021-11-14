package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	//ルート証明書を外部コマンドで実行させる。
	cmdMkdir := exec.Command("sudo", "mkdir", "/usr/local/share/ca-certificates/extra")
	cmdMkdir.Stderr = os.Stderr
	cmdMkdir.Stdin = os.Stdin

	out, err := cmdMkdir.Output()
	if err != nil {
		fmt.Println("Err", err)
	} else {
		fmt.Println("mkdir", string(out))
	}

	cmdCp := exec.Command("sudo", "cp", "foo.crt", "/usr/local/share/ca-certificates/extra/foo.crt")
	cmdCp.Stderr = os.Stderr
	cmdCp.Stdin = os.Stdin

	out, err = cmdCp.Output()
	if err != nil {
		fmt.Println("Err", err)
	} else {
		fmt.Println("cp", string(out))
	}

	cmdDpkg := exec.Command("sudo", "update-ca-certificates")
	cmdDpkg.Stderr = os.Stderr
	cmdDpkg.Stdin = os.Stdin

	out, err = cmdDpkg.Output()
	if err != nil {
		fmt.Println("Err", err)
	} else {
		fmt.Println("update", string(out))
	}
}
