package handler

import (
	"context"
	"fmt"
	"log"

	// "fmt"

	"cashier-admin/model"

	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
)

func ShowUserLoginForm(c *fiber.Ctx) error {
	// check if logged
	sess := getSession(c)

	if sess.Get("user") != nil {
		return c.Redirect("/")
	}

	// render template here
	return c.Render("user_login_formal", fiber.Map{
		"msg":  getFlashMessage(c),
		"msg1": "123",
	})
} // ShowUserLoginForm()

func ShowUserChangePasswordForm(c *fiber.Ctx) error {
	// check if logged
	sess := getSession(c)

	if sess.Get("user") == nil {
		return c.Redirect("/user/login")
	}

	// render template here
	return c.Render("user_changePassword", fiber.Map{
		"msg": getFlashMessage(c),
	})
} // ShowUserChangePasswordForm()

func HandleUserLogin(c *fiber.Ctx) error {
	// check if logged
	sess := getSession(c)

	if sess.Get("user") != nil { // 假如已經登入過 就不用再登入了
		return c.Redirect("/")
	}

	// parse user input
	credential := new(UserLoginRequest)

	if err := c.BodyParser(credential); err != nil {
		return err
	}

	// check the database if we found anything we returns
	user := new(model.Merchants)

	ctx := context.Background()
	db := model.GetDB()

	err := db.NewSelect().Model(user).Where("username = ?", credential.Username).Scan(ctx)

	if err != nil && err.Error() != "sql: no rows in result set" {
		return err
	}

	if user != nil {

		// ex : "123", hash
		match, _ := argon2id.ComparePasswordAndHash(credential.Password, user.Password)

		if match {

			sess.Set("user", credential.Username)

			// Save session
			if err := sess.Save(); err != nil {
				panic(err)
			}

			return c.Redirect("/")
		} // if()
	} // if()

	setFlashMessage(c, "Given credentials found no matched.")

	return c.Redirect("/user/login")
} // HandleUserLogin()

func HandleUserLogout(c *fiber.Ctx) error {
	sess := getSession(c)

	if err := sess.Destroy(); err != nil {
		panic(err)
	} // if()

	return c.Redirect("/user/login")
} // HandleUserLogout()

func CheckUserAuth() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		sess := getSession(c)

		if sess.Get("user") == nil {
			return c.Redirect("/user/login")
		}

		return c.Next()
	}
}

func ChangeUserPassword(c *fiber.Ctx) error {
	// check if logged
	sess := getSession(c)

	str := fmt.Sprintf("%v", sess.Get("user"))

	passworddata := new(UserChangePassword)
	if err := c.BodyParser(passworddata); err != nil { // 取得進來的新舊密碼
		return err
	}

	// check the database if we found anything we returns
	user := new(model.Merchants)

	ctx := context.Background()
	db := model.GetDB()

	err := db.NewSelect().Model(user).Where("username = ?", str).Scan(ctx) // 先去資料庫找尋目前 session 裡面所存放的 user 是哪一位

	if err != nil && err.Error() != "sql: no rows in result set" {
		return err
	} // if()

	// ex : "123", hash 去檢查使用者所輸入的 oldpassword 與資料庫中的 password 是否一致 是的話就可以進行帳密修改
	match, err := argon2id.ComparePasswordAndHash(passworddata.Previouspassword, user.Password)

	if !match || (err != nil) {
		return c.Redirect("/user/changePassword")
	} else {
		if passworddata.Password != passworddata.Newpassword {
			return c.Redirect("/user/changePassword")
		} // if()

		hash, err := argon2id.CreateHash(passworddata.Password, argon2id.DefaultParams)
		if err != nil {
			log.Fatal(err)
		}

		user.Password = hash

		err = model.ChangeUserPassword(user)
		if err != nil {
			return err
		} // if()
	} // else()

	fmt.Println("changed password success!!!")

	return c.Redirect("/user/dashboard")
} // ChangeUserPassword()

type UserLoginRequest struct {
	Username string
	Password string
} // UserLoginRequest()

type UserChangePassword struct {
	Previouspassword string
	Password         string
	Newpassword      string
} // UserChangePassword()
