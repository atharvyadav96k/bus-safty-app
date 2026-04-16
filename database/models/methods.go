package database_models

import (
	"time"

	"github.com/atharvyadav96k/gcp/common/entity"
	"golang.org/x/crypto/bcrypt"
)

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password, err = entity.NewPassword(string(hashedPassword))

	now := time.Now()
	if u.CreatedAt.IsZero() {
		u.CreatedAt = now
	}
	u.UpdatedAt = now

	return nil
}
