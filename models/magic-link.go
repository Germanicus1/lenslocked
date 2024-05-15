package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"githubn.com/Germanicus1/lenslocked/rand"
)

const (
	DefaultMLDuration = 1 * time.Hour
)

type MagicLink struct {
	ID          int
	UserID      int
	MLToken     string // Only populated when we create a new magic link
	MLTokenHash string
	ExpiresAt   time.Time
}

type MagicLinkService struct {
	DB *sql.DB
	// BytesPerToken is used to determine how many bytes to use when generating
	// each token. If this value is not set or is less than the
	// MinBytesPerToken const it will be ignored and MinBytesPerToken will be
	// used.
	BytesPerToken int
	Duration      time.Duration
}

func (service *MagicLinkService) CreateMagicLink(email string) (*MagicLink, error) {
	// Verify we have a valid email address and get that users ID
	email = strings.ToLower(email)
	var userID int
	row := service.DB.QueryRow(`
		SELECT id FROM users WHERE email = $1;`, email)
	err := row.Scan(&userID)
	if err != nil {
		// TODO: Consider returning a specific error when the user does not exist.
		return nil, fmt.Errorf("create: %w", err)
	}
	// Build the magic link token reset
	bytesPerToken := service.BytesPerToken
	if bytesPerToken < MinBytesPerToken {
		bytesPerToken = MinBytesPerToken
	}
	token, err := rand.String(bytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("MagicLink.Create: %w", err)
	}
	fmt.Println("MLToken: ", token)
	duration := service.Duration
	if duration == 0 {
		duration = DefualtResetDuration
	}
	ml := MagicLink{
		UserID:      userID,
		MLToken:     token,
		MLTokenHash: service.hash(token),
		ExpiresAt:   time.Now().Add(duration),
	}
	row = service.DB.QueryRow(`
		INSERT INTO magic_links (user_id, ml_token_hash, expires_at)
		VALUES ($1, $2, $3) ON CONFLICT (user_id) DO
		UPDATE
		SET ml_token_hash = $2, expires_at =$3
		RETURNING id;`, ml.UserID, ml.MLTokenHash, ml.ExpiresAt)
	err = row.Scan(&ml.ID)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	return &ml, nil
}

func (service *MagicLinkService) ConsumeMagicLink(token string) (*User, error) {
	tokenHash := service.hash(token)
	fmt.Println("TH:", tokenHash)
	var user User
	var ml MagicLink
	tableName := "magic_links"
	query := fmt.Sprintf(`
		SELECT %s.id,
			%s.expires_at,
			users.id,
			users.email
		FROM %s
			JOIN users ON users.id = %s.user_id
		WHERE %s.ml_token_hash = $1`, tableName, tableName, tableName, tableName, tableName)
	row := service.DB.QueryRow(query, tokenHash)
	// row := service.DB.QueryRow(`
	// 	SELECT magic_links.id,
	// 		magic_links.expires_at,
	// 		users.id,
	// 		users.email
	// 	FROM magic_links
	// 		JOIN users ON users.id = magic_links.user_id
	// 	WHERE magic_links.ml_token_hash = $1`, tokenHash)
	err := row.Scan(
		&ml.ID, &ml.ExpiresAt,
		&user.ID, &user.Email)
	if err != nil {
		return nil, fmt.Errorf("ConsumeMagicLink: consume: %w", err)
	}
	if time.Now().After(ml.ExpiresAt) {
		return nil, fmt.Errorf("ConsumeMagicLink: token expired: %v", token)
	}
	err = service.delete(ml.ID)
	if err != nil {
		return nil, fmt.Errorf("ConsumeMagicLink: delete: %w", err)
	}
	return &user, nil
}

func (service *MagicLinkService) hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(tokenHash[:])
}

func (service *MagicLinkService) delete(id int) error {
	_, err := service.DB.Exec(`
		DELETE FROM magic_links
		WHERE id = $1;`, id)
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}
	return nil
}
