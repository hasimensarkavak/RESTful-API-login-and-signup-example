package utils

import "github.com/gofiber/fiber/v2"

func UseTemplatePage(ctx *fiber.Ctx, page string, err bool, success bool, sonuc string) error {
	type Todo struct {
		Sonuc   string
		Err     bool
		Success bool
	}
	data := Todo{
		sonuc,
		err,
		success,
	}
	return ctx.Render(page, data)
}
