package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	//ルート証明書を外部コマンドで実行させる。
	cmd := exec.Command("sudo", "mkdir", "/usr/local/share/ca-certificates/extra")
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Err", err)
	} else {
		fmt.Println("make", string(out))
	}
}
