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
			fmt.Println(fmt.Errorf("ByID: error scanning gallery: %v", err))
			return nil, ErrNotFound
		}
	}
	return &gallery, nil
}

func (service *GalleryService) ByUserId(userID int) ([]Gallery, error){
	rows, err := service.DB.Query(`
		SELECT id, title
		FROM galleries
		WHERE user_id = $1;`, userID)
	if err != nil {
		return nil, fmt.Errorf("ByUserId: error querying galleries: %v", err)
	}

	var galleries []Gallery

	for rows.Next(){
		var gallery Gallery
		err := rows.Scan(&gallery.ID, &gallery.Title)
		if err != nil {
			return nil, fmt.Errorf("ByUserId: error scanning galleries: %v", err)
		}
		galleries = append(galleries, gallery)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("ByUserId: error iterating galleries: %v", rows.Err())
	}
	return galleries, nil
}
