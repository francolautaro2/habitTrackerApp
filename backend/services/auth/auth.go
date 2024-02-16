package auth

import (
	"errors"
	"habitTrackerApi/services/users"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("your-secret-key") // Clave secreta para firmar el token (cambia esto por una clave segura en producción)

// CustomClaims define la estructura de los claims personalizados del token JWT
type CustomClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

// LoginRequest representa la estructura de la solicitud de inicio de sesión
type LoginRequest struct {
	UsernameOrEmail string `json:"username_or_email" binding:"required"`
	Password        string `json:"password" binding:"required"`
}

// AuthController representa el controlador de autenticación
type AuthController struct {
	UserRepository users.UserRepository
}

// NewAuthController crea una nueva instancia del controlador de autenticación
func NewAuthController(userRepo users.UserRepository) *AuthController {
	return &AuthController{
		UserRepository: userRepo,
	}
}

// Login realiza el proceso de inicio de sesión y genera un token JWT válido
func (ac *AuthController) Login(c *gin.Context) {
	var loginData LoginRequest
	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Buscar el usuario por nombre de usuario o correo electrónico
	user, err := ac.UserRepository.GetUserByUsernameOrEmail(loginData.UsernameOrEmail)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Verificar la contraseña
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generar token de autenticación
	token, err := GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	cookie := &http.Cookie{
		Name:  "auth_token", // Nombre de la cookie
		Value: token,
	}

	http.SetCookie(c.Writer, cookie)

	// Establecer el token en el encabezado de autorización
	c.Header("Authorization", "Bearer "+token)

	// Devolver el token al cliente
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// GenerateToken genera un token JWT para un usuario dado
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

	// Crear el token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firmar el token con la clave secreta
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// AuthMiddleware es un middleware para la autorización que verifica la validez del token JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el token de autorización del encabezado
		tokenString := c.GetHeader("Authorization")

		// Verificar si el token está presente
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
			c.Abort()
			return
		}

		// Verificar el formato del token y extraer el token
		if len(tokenString) < 7 || tokenString[:7] != "Bearer " {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token format"})
			c.Abort()
			return
		}
		tokenString = tokenString[7:]

		// Verificar la validez del token y extraer los claims
		claims, err := verifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			c.Abort()
			return
		}

		// Verificar si el usuario está logueado
		if claims.UserID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
			c.Abort()
			return
		}

		// Permitir que la solicitud continúe
		c.Next()
	}
}

// verifyToken verifica la validez del token JWT y extrae los claims
func verifyToken(tokenString string) (*CustomClaims, error) {
	// Parsear el token y verificar la firma
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Extraer los claims del token
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
