package middleware

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Errors related to JWT authentication service (not all of them used by now)
var (
	// ErrEmptyAuthHeader can be thrown if authing with a HTTP header, the Auth header needs to be set
	ErrEmptyAuthHeader = errors.New("auth header is empty")
	ErrForbidden       = errors.New("you don't have permission to access this resource")

	// ErrInvalidAuthHeader indicates auth header is invalid, could for example have the wrong Realm name
	ErrInvalidAuthHeader = errors.New("auth header is invalid")
	jwt_msg              = "message"
	jwt_status           = "code"
	jwt_auth             = "Authorization"
	content_type         = "Content-Type"
	// AuthSigningKey auth sign key
	AuthSigningKey []byte = []byte("S@H4BAT-U7D")
	// AuthSigningMethod auth sign method
	AuthSigningMethod jwt.SigningMethod = jwt.SigningMethodHS256
	// AuthTokenValidityDuration token duration
	AuthTokenValidityDuration int64 = 60 * 60 // 1440 minutes in seconds (24hr)
	// AuthIssuer auth issuer
	AuthIssuer string = "SAHABAT-UTD"
	inArray           = func(str string, list []string) bool {
		str = strings.ToLower(str)
		for _, v := range list {
			if strings.ToLower(v) == str {
				return true
			}
		}
		return false
	}
)

// Structs used in auth system
type (
	// JWTClaims for JWT token
	JWTClaims struct {
		AuthUser
		StandardClaims
	}

	// StandardClaims standard claims
	StandardClaims struct {
		Audience  string `json:"aud,omitempty"`
		ExpiresAt int64  `json:"exp,omitempty"`
		Id        string `json:"jti,omitempty"`
		IssuedAt  int64  `json:"iat,omitempty"`
		Issuer    string `json:"iss,omitempty"`
		NotBefore int64  `json:"nbf,omitempty"`
		Subject   string `json:"sub,omitempty"`
	}

	// ValidateAccessRequest is used to request handler
	ValidateAccessRequest struct {
		Roles []string `json:"roles" gorm:"roles"`
	}

	// AuthUser is used as user model
	AuthUser struct {
		Email         string `json:"email"`
		DonorId       string `json:"donor_id"`
		Nama          string `json:"nama"`
		GolonganDarah string `json:"golongan_darah"`
		Role          string `json:"role"`
	}
	SimabaUser struct {
		Email         string `json:"email"`
		DonorId       string `json:"donor_id"`
		Nama          string `json:"nama"`
		GolonganDarah string `json:"golongan_darah"`
		Role          string `json:"role"`
	}
	SimabaClaims struct {
		SimabaUser
		jwt.StandardClaims
	}
)

// Function to get JWT token from header
// token must exists in "Authorization" header with "Bearer<space>" prefix
func JwtFromHeader(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get(jwt_auth)

	if authHeader == "" {
		return "", ErrEmptyAuthHeader
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", ErrInvalidAuthHeader
	}

	return parts[1], nil
}

func GetDefaultCorsJWT() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"}
	config.AllowHeaders = []string{"Origin", "Content-Length", content_type, jwt_auth, "X-Requested-With", "Accept",
		"Access-Control-Allow-Headers", "Accept-Encoding", "X-CSRF-Token"}
	return cors.New(config)
}

// JWTMiddleware to enforce JWT authorization to routes
func JWTMiddleware(roles ...string) gin.HandlerFunc {

	return func(c *gin.Context) {
		tokenString, err := JwtFromHeader(c)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				jwt_status: http.StatusUnauthorized,
				jwt_msg:    err.Error(),
			})
			return
		}

		claims, err := ValidateAccess(tokenString, roles)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				jwt_status: http.StatusUnauthorized,
				jwt_msg:    err.Error(),
			})
			return

		}
		// set parameter to be passed to handler
		c.Set("JWT_TOKEN", tokenString)
		c.Set("JWT_CLAIMS", claims)
		c.Next()
	}
}

func GenerateToken(user SimabaUser) (string, error) {
	claims := SimabaClaims{
		SimabaUser{
			Email:         user.Email,
			DonorId:       user.DonorId,
			Nama:          user.Nama,
			GolonganDarah: user.GolonganDarah,
			Role:          user.Role,
		},
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + AuthTokenValidityDuration,
			Issuer:    AuthIssuer,
		},
	}
	token := jwt.NewWithClaims(AuthSigningMethod, claims)
	// log.Println("Generating token string")
	// // Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(AuthSigningKey)

	return tokenString, err
}

// Function to parse claims (data) from JWT token
// claims will be returned in BRIBrainClaims format
func ParseJWT(tokenString string) (*SimabaClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &SimabaClaims{}, func(token *jwt.Token) (interface{}, error) {
		return AuthSigningKey, nil
	})
	if err == nil {
		if claims, ok := token.Claims.(*SimabaClaims); ok && token.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// parse JWT, then check
// if roles empty then pass the request
// otherwise, check the roles with current user roles (AccessLevel)
func ValidateAccess(tokenString string, roles []string) (*SimabaClaims, error) {
	claims, err := ParseJWT(tokenString)
	if err == nil {
		if len(roles) > 0 && inArray(claims.Role, roles) == false {
			err = ErrForbidden
		}
	}
	return claims, err
}
