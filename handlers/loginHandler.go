package handlers

import (
	"Login_And_SignUp_Example/helpers"
	"Login_And_SignUp_Example/modules"
	"Login_And_SignUp_Example/utils"
	"github.com/gofiber/fiber/v2"
	"log"
)

func LoginPage(ctx *fiber.Ctx) error {
	return ctx.Render("LoginPanel", nil)

}

func Login(ctx *fiber.Ctx) error {
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")

	emailCheck := helpers.IsEmpty(email)
	passCheck := helpers.IsEmpty(password)

	page := "LoginPanel"

	if emailCheck || passCheck {
		log.Printf("There is Empty data")
		return utils.UseTemplatePage(ctx, page, true, false, "Boş Veri Gönderemezsiniz")
	}

	user, err := modules.CallUser(email)
	if err != nil {
		//log.Printf("user not exist %w", err)
		return utils.UseTemplatePage(ctx, page, true, false, "Email Yada Şifre Hatalı. Lütfen Tekrar Deneyiniz")
	}

	err = utils.CheckPassword(password, user.Password)
	if err == nil && email == user.Email {
		return utils.UseTemplatePage(ctx, page, false, true, "Giriş Başarılı")
	} else {
		return utils.UseTemplatePage(ctx, page, true, false, "Email Yada Şifre Hatalı. Lütfen Tekrar Deneyiniz.")
	}
}
