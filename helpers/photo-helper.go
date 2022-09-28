package helpers

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/nandaditra/task-5-vix-btpns-Nanda-Aditya-Putra/database"
	"github.com/nandaditra/task-5-vix-btpns-Nanda-Aditya-Putra/models"
)

type PhotoHelper interface {
	Insert(p database.PhotoCreateData) models.Photo
	Update(p database.PhotoUpdateData) models.Photo
	Delete(p models.Photo)
	GetByID(photoID uint64) models.Photo
	IsAllowedToEdit(userID string, photoID uint64) bool
}

type photoHelper struct {
	photoRepo models.PhotoRepo
}

func NewPhotoHelper(photoRep models.PhotoRepo) PhotoHelper {
	return &photoHelper{
		photoRepo: photoRep,
	}
}

func (helper *photoHelper) Insert(p database.PhotoCreateData) models.Photo {
	photo := models.Photo{}
	err := smapping.FillStruct(&photo, smapping.MapFields(&p))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := helper.photoRepo.InsertPhoto(photo)
	return res
}

func (helper *photoHelper) Update(p database.PhotoUpdateData) models.Photo {
	photo := models.Photo{}
	err := smapping.FillStruct(&photo, smapping.MapFields(&p))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := helper.photoRepo.UpdatePhoto(photo)
	return res
}

func (helper *photoHelper) Delete(p models.Photo) {
	helper.photoRepo.DeletePhoto(p)
}

func (helper *photoHelper) GetByID(photoID uint64) models.Photo {
	return helper.photoRepo.GetPhotoByID(photoID)
}

func (helper *photoHelper) IsAllowedToEdit(userID string, photoID uint64) bool {
	pho := helper.photoRepo.GetPhotoByID(photoID)
	id := fmt.Sprintf("%v", pho.UserID)
	return userID == id
}
