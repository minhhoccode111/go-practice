package model

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type User struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	IsActive *bool  `json:"is_active"`
	Role     Role   `json:"role"`
	Password string
}

type UserDTO struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	IsActive *bool  `json:"is_active"`
	Role     Role   `json:"role"`
}

var JwtSecret = os.Getenv("JWT_SECRET")

func (u *User) ToUserDTO() UserDTO {
	return UserDTO{
		Id:       u.Id,
		Email:    u.Email,
		IsActive: u.IsActive,
		Role:     u.Role,
	}
}

func (u *User) GenerateJWT() (string, error) {
	secretKey := []byte(JwtSecret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    u.Id,
		"userEmail": u.Email,
		"userRole":  u.Role,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
		"iat":       time.Now().Unix(),
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Printf("error signing token: %v", err)
		return "", fmt.Errorf("error signing token: %v", err)
	}
	return tokenString, nil
}

func (u *User) ValidatePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

func (u *User) HashedPassword() (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (u *User) IsValidEmail() error {
	u.Email = strings.TrimSpace(u.Email)
	if u.Email == "" {
		return fmt.Errorf("email cannot be empty")
	}
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(u.Email) {
		return fmt.Errorf("invalid email: %v", u.Email)
	}
	return nil
}

func (u *User) IsValidPassword() error {
	u.Password = strings.TrimSpace(u.Password)
	if len(u.Password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}
	hasUpper, hasLower, hasDigit, hasSpecial := false, false, false, false
	for _, ch := range u.Password {
		if unicode.IsUpper(ch) {
			hasUpper = true
		} else if unicode.IsLower(ch) {
			hasLower = true
		} else if unicode.IsDigit(ch) {
			hasDigit = true
		} else if strings.ContainsAny(string(ch), "!@#$%^&*()_+{}|:\"<>?~") {
			hasSpecial = true
		}
	}
	if hasUpper && hasLower && hasDigit && hasSpecial {
		return nil
	}
	return fmt.Errorf("weak password: '%v'. Must contain at least: 1 uppercase, 1 lowercase, 1 digit, 1 special character", u.Password)
}
