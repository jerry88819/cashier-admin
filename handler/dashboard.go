package handler

import (
	// "context"
	"fmt"
	// "os"
	// "path/filepath"
	// "strconv"
	// "time"

	// "image/jpeg"

	"cashier-admin/model"
	"github.com/gofiber/fiber/v2"
	// "github.com/gofrs/uuid"
	// "github.com/nfnt/resize"
	// "github.com/xuri/excelize/v2"
)

func ShowUserInfoList(c *fiber.Ctx) error {

	sess := getSession(c) // 找出剛剛登入的使用者
	str := fmt.Sprintf("%v", sess.Get("user"))
	fmt.Println(str)
	info, err := model.GetOneUserInfo(str)

	fmt.Println(info.Username)

	if err != nil {
		return err
	}

	return c.Render("user_dashboard", fiber.Map{
		"msg":  getFlashMessage(c),
		"info": info,
	})

} // ShowUserInfoList()
