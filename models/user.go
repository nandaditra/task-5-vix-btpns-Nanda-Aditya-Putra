package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       uint64   `gorm:"primary_key:auto_increment" json:"id"`
	Username string   `gorm:"type:varchar(255)" json:"name"`
	Email    string   `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string   `gorm:"->'<-; not null" json:"."`
	Token    string   `gorm:"-" json:"token,omitempty"`
	Photos   *[]Photo `json:"photos,omiempty"`
}

type UserRepo interface {
	InsertUser(user User) User
	UpdateUser(user User) User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (ab *gorm.DB)
	FindByEmail(email string) User
	GetUser(user User) User
	DeleteUser(user User)
	ProfileUser(userID string) User
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(user User) User {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

func (db *userConnection) UpdateUser(user User) User {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser User
		db.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}

	db.connection.Save(&user)
	return user
}

func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var users User
	res := db.connection.Where("email = ?", email).Take(&users)
	if res.Error == nil {
		return users
	}
	return nil
}

func (db *userConnection) IsDuplicateEmail(email string) (ab *gorm.DB) {
	var users User
	return db.connection.Where("email = ?", email).Take(&users)
}

func (db *userConnection) FindByEmail(email string) User {
	var users User
	db.connection.Where("email =?", email).Take(&users)
	return users
}

func (db *userConnection) GetUser(user User) User {
	var users User
	db.connection.Preload("User").Find(&users)
	return users
}

func (db *userConnection) DeleteUser(user User) {
	db.connection.Delete(&user)
}

func (db *userConnection) ProfileUser(userID string) User {
	var user User
	db.connection.Preload("Photos").Preload("Photos.User").Find("&user, userID")
	return user
}

func hashAndSalt(pass []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash the password")
	}
	return string(hash)
}
