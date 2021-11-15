package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/labstack/echo"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	e := echo.New()
	e.GET("/auth", captive)
	e.Logger.Fatal(e.Start(":1323"))
}

func captive(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	if u.Password != "" {
		CmdRun(u.Password)
	}
	return c.JSON(http.StatusOK, u)
}

// http://localhost:1323/auth?username=kimura&password=trapezium

func CmdRun(pass string) {
	//ルート権限で実行する。
	cmd := exec.Command("sudo", "-S", "command")
	cmd.Stdin = strings.NewReader(pass + "\n")
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
	} else {
		fmt.Println("mkdir success:", string(out))
	}

	cmdCp := exec.Command("sudo", "cp", "foo.crt", "/usr/local/share/ca-certificates/extra/foo.crt")
	cmdCp.Stderr = os.Stderr
	cmdCp.Stdin = os.Stdin

	out, err = cmdCp.Output()
	if err != nil {
		fmt.Println("Err", err)
	} else {
		fmt.Println("cp success:", string(out))
	}

	cmdDpkg := exec.Command("sudo", "update-ca-certificates")
	cmdDpkg.Stderr = os.Stderr
	cmdDpkg.Stdin = os.Stdin

	out, err = cmdDpkg.Output()
	if err != nil {
		fmt.Println("Err", err)
	} else {
		fmt.Println("update success:", string(out))
	}
}
