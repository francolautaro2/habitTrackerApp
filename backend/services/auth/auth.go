package auth

import (
	"errors"
	"habitTrackerApi/services/domains"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("your-secret-key") // The secret key, change this in production

type CustomClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

type LoginRequest struct {
	UsernameOrEmail string `json:"username_or_email" binding:"required"`
	Password        string `json:"password" binding:"required"`
}

type AuthController struct {
	UserRepository domains.UserRepository
}

func NewAuthController(userRepo domains.UserRepository) *AuthController {
	return &AuthController{
		UserRepository: userRepo,
	}
}

func (ac *AuthController) Login(c *gin.Context) {
	var loginData LoginRequest
	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user, err := ac.UserRepository.GetUserByUsernameOrEmail(loginData.UsernameOrEmail)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	cookie := &http.Cookie{
		Name:  "auth_token",
		Value: token,
	}

	http.SetCookie(c.Writer, cookie)

	c.Header("Authorization", "Bearer "+token)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func GenerateToken(userID uint) (string, error) {
	// Crear los claims del token
	claims := CustomClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // El token expira en 24 horas
			IssuedAt:  time.Now().Unix(),
			Issuer:    "habitTracker",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
			c.Abort()
			return
		}

		if len(tokenString) < 7 || tokenString[:7] != "Bearer " {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token format"})
			c.Abort()
			return
		}
		tokenString = tokenString[7:]

		claims, err := VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			c.Abort()
			return
		}

		if claims.UserID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func VerifyToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}

// Get the ID from the JWT Token
func GetUserIDFromToken(c *gin.Context) (uint, error) {
	// Get token from the header authorization
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return 0, errors.New("missing authorization token")
	}

	// Parser and verify the jwt token
	claims, err := VerifyToken(tokenString)
	if err != nil {
		return 0, err
	}
	// Return id from the jwt token
	return claims.UserID, nil
}
