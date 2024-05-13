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

func (service *MagicLinkService) Create(email string) (*MagicLink, error) {
	email = strings.ToLower(email)
	var userID int

	row := service.DB.QueryRow(`SELECT id FROM users WHERE email = $1`, email)
	err := row.Scan(&userID)
	if err != nil {
		return nil, fmt.Errorf("MagicLink.Create: %w", err)
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

	// INCOMPLETE
	// Write the data to the DB

	row = service.DB.QueryRow(`
		INSERT INTO magic_links (user_id, ml_token_hash, expires_at)
		VALUES ($1, $2, $3) ON CONFLICT (user_id) DO
		UPDATE
		SET token_hash = $2, expires_at =$3
		RETURNING id;`, ml.UserID, ml.MLTokenHash, ml.ExpiresAt)
	err = row.Scan(&ml.ID)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}

	return &ml, nil
}

func (service MagicLinkService) Consume(tokenHash string) (*User, error) {
	return nil, fmt.Errorf("INCOMPLETE: Implement Consume()")
}

func (service *MagicLinkService) hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(tokenHash[:])
}
