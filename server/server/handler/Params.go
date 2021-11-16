package handler

//AuthUserParams: 認証に用いるパラメータ
type AuthUserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
