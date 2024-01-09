package authentication

import (
	"golang.org/x/crypto/bcrypt"
	"isaacszf.antiqbrasblog.com/domain"
	"isaacszf.antiqbrasblog.com/domain/models"
)

func LoginCheck(username string, password string) (string, error) {
	var err error

	w := models.Writer{}
	err = domain.DB.Model(models.Writer{}).
		Where("username = ?", username).
		Take(&w).Error
	
	if err != nil {
		return "", err
	}

	err = verifyPassword(password, w.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := generateToken(w.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func verifyPassword(pw string, hashedPw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPw), []byte(pw))
}