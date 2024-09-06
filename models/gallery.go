package models

import (
	"database/sql"
	"errors"
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

// Create will create the provided gallery and backfill data
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

func (service *GalleryService) ByID(id int) (*Gallery, error){
	gallery := Gallery{
		ID: id,
	}
	row := service.DB.QueryRow(`
		SELECT title, user_id
		FROM galleries WHERE id = $1;`, gallery.ID)
	err := row.Scan(&gallery.Title, &gallery.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Errorf("ByID: error scanning gallery: %v", err)
			return nil, ErrNotFound
		}
	}
	return &gallery, nil
}