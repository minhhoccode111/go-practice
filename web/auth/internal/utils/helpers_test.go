package utils

import (
	"auth/internal/config"
	"auth/internal/model"
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		want    string
		wantErr bool
	}{
		{
			name:    "Valid Email",
			email:   "test@example.com",
			want:    "test@example.com",
			wantErr: false,
		},
		{
			name:    "Email with plus sign",
			email:   "test+alias@example.com",
			want:    "test+alias@example.com",
			wantErr: false,
		},
		{
			name:    "Email with dot in local part",
			email:   "first.last@example.com",
			want:    "first.last@example.com",
			wantErr: false,
		},
		{
			name:    "Email with subdomain",
			email:   "test@sub.example.com",
			want:    "test@sub.example.com",
			wantErr: false,
		},
		{
			name:    "Empty Email",
			email:   "",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Email without @",
			email:   "testexample.com",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Email without domain",
			email:   "test@",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Email without top-level domain",
			email:   "test@example",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Email with invalid characters",
			email:   "test!@example.com",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Email with leading/trailing spaces",
			email:   "  test@example.com  ",
			want:    "test@example.com",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsValidEmail(tt.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsValidEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsValidEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserToUserDTO(t *testing.T) {
	user := &model.User{
		Id:       "123",
		Email:    "test@example.com",
		IsActive: true,
		Role:     model.RoleUser,
		// Password and CreatedAt/UpdatedAt are not part of DTO, so no need to set them for this test
	}

	dto := UserToUserDTO(user)

	if dto.Id != user.Id {
		t.Errorf("Expected Id %v, got %v", user.Id, dto.Id)
	}
	if dto.Email != user.Email {
		t.Errorf("Expected Email %v, got %v", user.Email, dto.Email)
	}
	if dto.IsActive != user.IsActive {
		t.Errorf("Expected IsActive %v, got %v", user.IsActive, dto.IsActive)
	}
	if dto.Role != user.Role {
		t.Errorf("Expected Role %v, got %v", user.Role, dto.Role)
	}
}

func TestHashedPassword(t *testing.T) {
	password := "mysecretpassword"
	hashedPassword, err := HashedPassword(password)

	if err != nil {
		t.Fatalf("HashedPassword() error = %v", err)
	}

	if hashedPassword == "" {
		t.Errorf("HashedPassword() returned empty string")
	}

	// Verify that the hashed password can be validated
	if !ValidatePassword(hashedPassword, password) {
		t.Errorf("HashedPassword() generated a password that cannot be validated")
	}
}

func TestValidatePassword(t *testing.T) {
	password := "mysecretpassword"
	hashedPassword, _ := HashedPassword(password)

	tests := []struct {
		name   string
		hashed string
		plain  string
		want   bool
	}{
		{
			name:   "Correct Password",
			hashed: hashedPassword,
			plain:  password,
			want:   true,
		},
		{
			name:   "Incorrect Password",
			hashed: hashedPassword,
			plain:  "wrongpassword",
			want:   false,
		},
		{
			name:   "Empty Plain Password",
			hashed: hashedPassword,
			plain:  "",
			want:   false,
		},
		{
			name:   "Empty Hashed Password",
			hashed: "",
			plain:  password,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidatePassword(tt.hashed, tt.plain)
			if got != tt.want {
				t.Errorf("ValidatePassword() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     string
		wantErr  bool
	}{
		{
			name:     "Valid Password",
			password: "Password123!",
			want:     "Password123!",
			wantErr:  false,
		},
		{
			name:     "Password with leading/trailing spaces",
			password: "  Password123!  ",
			want:     "Password123!",
			wantErr:  false,
		},
		{
			name:     "Too Short",
			password: "Pass1!",
			want:     "",
			wantErr:  true,
		},
		{
			name:     "No Uppercase",
			password: "password123!",
			want:     "",
			wantErr:  true,
		},
		{
			name:     "No Lowercase",
			password: "PASSWORD123!",
			want:     "",
			wantErr:  true,
		},
		{
			name:     "No Digit",
			password: "Password!!",
			want:     "",
			wantErr:  true,
		},
		{
			name:     "No Special Character",
			password: "Password123",
			want:     "",
			wantErr:  true,
		},
		{
			name:     "Empty Password",
			password: "",
			want:     "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsValidPassword(tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsValidPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsValidPassword() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateJWT(t *testing.T) {
	jwtConfig := config.JWTConfig{
		Secret:     "supersecretjwtkey",
		Expiration: 3600, // 1 hour
		Issuer:     "test-issuer",
	}

	userDTO := &model.UserDTO{
		Id:       "user123",
		Email:    "user@example.com",
		IsActive: true,
		Role:     model.RoleUser,
	}

	tokenString, err := GenerateJWT(jwtConfig, userDTO)

	if err != nil {
		t.Fatalf("GenerateJWT() error = %v", err)
	}

	if tokenString == "" {
		t.Errorf("GenerateJWT() returned empty token string")
	}

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtConfig.Secret), nil
	}, jwt.WithLeeway(time.Minute))

	if err != nil {
		t.Fatalf("Error parsing token: %v", err)
	}

	if !token.Valid {
		t.Errorf("Generated token is not valid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		t.Errorf("Could not get claims from token")
	}

	if claims["userId"] != userDTO.Id {
		t.Errorf("Expected userId %v, got %v", userDTO.Id, claims["userId"])
	}

	if claims["iss"] != jwtConfig.Issuer {
		t.Errorf("Expected issuer %v, got %v", jwtConfig.Issuer, claims["iss"])
	}

	exp := int64(claims["exp"].(float64))
	if exp < time.Now().Unix() {
		t.Errorf("Token expiration time is not in the future")
	}
}
