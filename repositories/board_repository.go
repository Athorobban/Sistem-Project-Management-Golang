package repositories

import (
	"github.com/Athorobban/Sistem-Project-Management-Golang/config"
	"github.com/Athorobban/Sistem-Project-Management-Golang/models"
)

type BoardRepository interface {
	Create(board *models.Board) error
}

type boardRepository struct {
}

func NewBoardRepository() BoardRepository {
	return &boardRepository{}
}

func (r *boardRepository) Create(board *models.Board) error {
	return config.DB.Create(board).Error
}