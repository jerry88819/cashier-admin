package handler

import (
	//"context"
	//"fmt"
	//"log"

	"fmt"
	"strconv"

	"cashier-admin/model"

	//"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
)

func ShowUserOrderForm(c *fiber.Ctx) error {
	// check if logged
	sess := getSession(c)

	if sess.Get("user") == nil {
		return c.Redirect("/user/login")
	}

	// render template here
	return c.Render("user_showorder", fiber.Map{
		"msg": getFlashMessage(c),
	})
} // ShowUserOrderForm()

func ShowUserOrder(c *fiber.Ctx) error {
	sess := getSession(c)
	if sess.Get("user") == nil {
		return c.Redirect("/user/login")
	} // if()

	str := fmt.Sprintf("%v", sess.Get("user"))
	temp, err := model.GetOneUserInfo(str)
	if err != nil {
		return err
	} // if()

	ans, err := model.GetUserOrder(int(temp.ID))
	if err != nil {
		return err
	} // if()

	fmt.Println(ans[0].MchId)

	count := len(ans)
	sess.Set("totalpages", count) // 把資料的頁數存到 session 做分頁

	if count%10 != 0 {
		count = (count / 10) + 1
	} else { // if()
		count = count / 10
	} // else()

	var counts []int
	i := 1
	for i <= count {
		fmt.Println(count)
		counts = append(counts, i)
		i = i + 1
	} // for()

	fmt.Println(counts)

	return c.Render("user_showorder", fiber.Map{
		"msg":        getFlashMessage(c),
		"order":      ans,
		"name":       temp.Name,
		"pagescount": counts,
	})
} // ShowUserOrder()

func ShowUserOrder1(c *fiber.Ctx) error {

	tempage := c.Params("pageId")

	fmt.Println("tempage: " + tempage)
	// string to int
	whichpage, err := strconv.Atoi(tempage)
	if err != nil {
		// handle error
		fmt.Println(err)
	} // if()

	// 確認 user 是在登入狀態
	sess := getSession(c)
	if sess.Get("user") == nil {
		return c.Redirect("/user/login")
	} // if()

	str := fmt.Sprintf("%v", sess.Get("user"))
	temp, err := model.GetOneUserInfo(str)
	if err != nil {
		return err
	} // if()

	ans, err := model.GetUserOrder(int(temp.ID))
	if err != nil {
		return err
	} // if()

	count := len(ans)

	if count%10 != 0 {
		count = (count / 10) + 1
	} else { // if()
		count = count / 10
	} // else()

	var counts []int
	i := 1
	for i <= count {
		counts = append(counts, i)
		i = i + 1
	} // for()

	first := 0
	last := 0
	// 先處理最後一頁狀況
	if whichpage == count { // 當前頁數 等於 總共頁數
		first = (count * 10) - 10
		last = len(ans) - 1
	} else { // if()
		first = (whichpage * 10) - 10
		last = first + 9
	} // else()

	fmt.Println("count : ", count)
	fmt.Println("whichpage : ", whichpage)
	fmt.Println("first : ", first)
	fmt.Println("last : ", last)

	var selectedAns []model.Order
	for i := first; i <= last; i++ {
		selectedAns = append(selectedAns, ans[i])
	} // for()

	fmt.Println(selectedAns)

	return c.Render("user_order", fiber.Map{
		"msg":        getFlashMessage(c),
		"order":      selectedAns,
		"name":       temp.Name,
		"pagesdata":  counts,
		"pagescount": count,
		"whichpage":  whichpage,
	})

} // ShowUserOrder()
