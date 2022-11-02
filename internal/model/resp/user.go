package resp

import "github.com/tmnhs/common-test/internal/model"

type (
	RspLogin struct {
		User  *model.User `json:"user"`
		Token string      `json:"token"`
	}
	RspUser struct {
		ID       int    `json:"id" `
		UserName string `json:"username"`
		Email    string `json:"email" `
		Role     int    `json:"role" `
		Status   int    `json:"status"`
		Created  int64  `json:"created"`
		Updated  int64  `json:"updated"`
	}
)
