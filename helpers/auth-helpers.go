package helpers

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/nandaditra/task-5-vix-btpns-Nanda-Aditya-Putra/app"
	"github.com/nandaditra/task-5-vix-btpns-Nanda-Aditya-Putra/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthHelper interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user app.RegisterApp) models.User
	FindByEmail(email string) models.User
	IsDuplicateEmail(email string) bool
}

type authHelper struct {
	userRep models.UserRepo
}

func NewAuthHelper(userAp models.UserRepo) AuthHelper {
	return &authHelper{
		userRep: userAp,
	}
}

func (helper *authHelper) VerifyCredential(email string, password string) interface{} {
	res := helper.userRep.VerifyCredential(email, password)
	if v, ok := res.(models.User); ok {
		comparedPassword := comparedPassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (helper *authHelper) CreateUser(user app.RegisterApp) models.User {
	userCreate := models.User{}
	err := smapping.FillStruct(&userCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := helper.userRep.InsertUser(userCreate)
	return res
}

func (helper *authHelper) FindByEmail(email string) models.User {
	return helper.userRep.FindByEmail(email)
}

func (helper *authHelper) IsDuplicateEmail(email string) bool {
	res := helper.userRep.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func comparedPassword(hashPassword string, plainPassword []byte) bool {
	byteHash := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
