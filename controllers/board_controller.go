package controllers

import (
	"github.com/Athorobban/Sistem-Project-Management-Golang/models"
	"github.com/Athorobban/Sistem-Project-Management-Golang/services"
	"github.com/Athorobban/Sistem-Project-Management-Golang/utils"
	"github.com/gofiber/fiber/v2"
)

type BoardController struct {
	service services.BoardService
}

func NewBoardController( s services.BoardService) *BoardController {
	return &BoardController{service: s}
}

func (c *BoardController) CreateBoard(ctx *fiber.Ctx) error {
	board := new(models.Board)

	if err := ctx.BodyParser(board); err != nil {
		return utils.BadRequest(ctx,"Gagal Membaca Request", err.Error())
	}
	if err := c.service.Create(board); err != nil {
		return utils.BadRequest(ctx, "Gagal Menyimpan Data", err.Error())
	}
	return utils.Success(ctx,"Board berhasil dibuat",board)
}