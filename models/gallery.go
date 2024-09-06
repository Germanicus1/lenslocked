package models

import (
	"database/sql"
	"fmt"
)

type Gallery struct {
	ID     int
	UserID int
	Title  string
}

type GalleryService struct {
	DB *sql.DB
}

func (service *GalleryService) Create(title string, userId int) (*Gallery, error){
	gallery := Gallery{
		Title: title,
		UserID: userId,
	}
	row:= service.DB.QueryRow(
		`INSERT INTO galleries (title, user_id)
		VALUES ($1, $2) RETURNING id;`,
		gallery.Title, gallery.UserID)
	err := row.Scan(&gallery.ID)
	if err != nil {
		return nil, fmt.Errorf("create gallery: error creating gallery: %v", err)
	}
	return &gallery, nil
}