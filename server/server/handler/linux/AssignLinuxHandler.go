package handler

import (
	"fmt"
	"net/http"

	"github.com/go-malproxy/server/handler"
	logic "github.com/go-malproxy/server/logic/linux"
	"github.com/labstack/echo"
)

func AssignLinuxAuthHander(c echo.Context) error {
	params := new(handler.AuthUserParams)
	if err := c.Bind(params); err != nil {
		return err
	}
	if params.Password != "" {
		fmt.Println("Get password: ", params.Password) //ここは後で削除する
		err := logic.AssignLinuxCmdHandler(params.Password)
		if err != nil {
			fmt.Println("Error :(")
			return err //外部コマンドが最後まで実行されなかった場合は、エラーを返す
		}
	}
	return c.JSON(http.StatusOK, params)
}
