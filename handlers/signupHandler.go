package handlers

import (
	"Login_And_SignUp_Example/helpers"
	"Login_And_SignUp_Example/modules"
	"Login_And_SignUp_Example/utils"
	"github.com/gofiber/fiber/v2"
	"log"
)

func SignUpHandler(ctx *fiber.Ctx) error {
	var user modules.User
	user.FullName = ctx.FormValue("nameSurname")
	user.Email = ctx.FormValue("email")
	user.Password = ctx.FormValue("password")
	passwordAgain := ctx.FormValue("passwordAgain")

	fullNameCheck := helpers.IsEmpty(user.FullName)
	emailCheck := helpers.IsEmpty(user.Email)
	passCheck := helpers.IsEmpty(user.Password)
	passAgainCheck := helpers.IsEmpty(passwordAgain)

	page := "SignUp"

	if fullNameCheck || emailCheck || passCheck || passAgainCheck {
		log.Printf("There is Empty data")
		return utils.UseTemplatePage(ctx, page, true, false, "Boş Veri Gönderemezsiniz")
	}

	if user.Password != passwordAgain {
		log.Printf("passwords must be same")
		return utils.UseTemplatePage(ctx, page, true, false, "Girdiğiniz Şifreler Aynı Değil. Lütfen Tekrar Giriniz.")
	}

	if modules.UserIsExist(user.Email) == true {
		//log.Printf("User is exist")
		return utils.UseTemplatePage(ctx, page, true, false, "Girdiğiniz Email Adına Bir Kullanıcı Mevcut.\n Lütfen Başka Bir Email Kullanın.")
	}

	var err error
	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		log.Printf("failed to hash password %w", err)
		return utils.UseTemplatePage(ctx, page, true, false, "Kayıt Başarısız. Bir Hata Meydana Geldi")
	}

	if modules.CreateUser(user) == false {
		return utils.UseTemplatePage(ctx, page, true, false, "Kayıt Başarısız.")
	}

	return utils.UseTemplatePage(ctx, page, false, true, "Kayıt Başarılı")
}

func SignUpPage(ctx *fiber.Ctx) error {
	return ctx.Render("SignUp", nil)
}
