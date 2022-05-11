package users

import (
	"fmt"
	"io/ioutil"
	config "swimming-content-management/config"
	domainErrors "swimming-content-management/domain"

	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

const (
	TokenGeneratorError = "Error in creating new User"
	AdminAuthorized     = "This Route is only accessible by an admin"
	CoachAuthorized     = "This Route is only accessible by an coach"
	AdminUser           = "admin"
	CoachUser           = "coach"
	ParentUser          = "parent"
)

// Authenticate interface lists the methods that our authentication service should implement
type Authenticate interface {
	Authenticate(reqUser *User, user *User) bool
	GenerateAccessToken(user *User) (string error)
	ValidateAccessToken(token string) (string error)
	CoachRouteValidation(tokenString string) (*User, error)
	AdminRouteValidation(tokenString string) (*User, error)
}

type AccessTokenCustomClaim struct {
	UserId  uuid.UUID
	KeyType string
	jwt.StandardClaims
}

func GenerateAccessToken(user *User) (string, error) {
	configuration, err := config.NewConfig()

	if err != nil {
		panic(err)
	}

	UserId := user.ID
	KeyType := "access"

	claims := AccessTokenCustomClaim{
		UserId,
		KeyType,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(configuration.JwtExpiration)).Unix(),
			Issuer:    "swimming-cms.auth",
		},
	}
	signBytes, err := ioutil.ReadFile(configuration.AccessTokenPrivateKeyPath)
	if err != nil {
		appError := domainErrors.NewAppError(errors.Wrap(err, TokenGeneratorError), domainErrors.RepositoryError)
		return "", appError
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)

	if err != nil {
		appError := domainErrors.NewAppError(errors.Wrap(err, TokenGeneratorError), domainErrors.RepositoryError)
		return "", appError
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(signKey)
}

func ValidateAccessToken(tokenString string) (*uuid.UUID, error) {
	configuration, err := config.NewConfig()
	if err != nil {
		panic(err)

	}
	token, err := jwt.ParseWithClaims(tokenString, &AccessTokenCustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected signing method in auth Token")
		}
		verifiyBytes, err := ioutil.ReadFile(configuration.AccessTokenPublicKeyPath)

		if err != nil {
			return nil, err
		}

		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifiyBytes)
		if err != nil {
			return nil, err
		}

		return verifyKey, nil

	})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	claims, ok := token.Claims.(*AccessTokenCustomClaim)

	if !ok || !token.Valid || claims.KeyType != "access" {
		return nil, errors.New("invalid token: authentication failed")
	}

	return &claims.UserId, nil
}

//implement admin guards
//implement coach quards
//implement parent guards

func AdminRouteValidation(tokenString string) (*User, error) {
	userId, err := ValidateAccessToken(tokenString)
	if err != nil {
		return nil, errors.New("invalid authentication")
	}
	authUser, err := GetUserById(userId)
	if authUser.Role.Name != AdminUser {
		return nil, errors.New(AdminAuthorized)
	}
	return authUser, nil

}

func CoachRouteValidation(tokenString string) (*User, error) {
	userId, err := ValidateAccessToken(tokenString)
	if err != nil {
		return nil, errors.New("invalid authentication")
	}
	authUser, err := GetUserById(userId)
	if authUser.Role.Name != CoachUser {
		return nil, errors.New(CoachAuthorized)
	}
	return authUser, nil

}

func AdultRouteValidation(tokenString string) (*User, error) {

	userId, err := ValidateAccessToken(tokenString)
	if err != nil {
		return nil, errors.New("invalid authentication")
	}
	authUser, err := GetUserById(userId)
	if authUser.Role.Name != CoachUser {
		return nil, errors.New(CoachAuthorized)
	}
	return authUser, nil

}
