package user

import (
	"fmt"
	"test/dao"
	"test/models"
	"time"
)

// 用于 返回 用户CRUD时候的错误
// 进行用户的格式规范
type UserError struct {
	//  错误码
	ErrorCode int
	//	错误信息
	ErrorMessage string
}

func (e *UserError) Error() string {
	return fmt.Sprintf("error code: %d, message: %s", e.ErrorCode, e.ErrorMessage)
}
func AddUser(name, password string) (err error) {
	if len(name) > 20 {
		return &UserError{
			ErrorCode:    1,
			ErrorMessage: "用户名太长",
		}
	}
	user := models.User{}
	user.Name = name
	user.Password = password
	user.Status = "上线"
	user.CTime = time.Now()
	user.MTime = time.Now()
	user.IsDelete = false

	dao.CreateUser(user)
	return &UserError{
		ErrorCode:    0,
		ErrorMessage: "创建成功",
	}
}