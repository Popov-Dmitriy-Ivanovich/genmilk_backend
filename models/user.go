package models

import (
	"errors"
	"gorm.io/gorm"
	"regexp"
)

type User struct {
	ID                    uint   `gorm:"primaryKey"`
	NameSurnamePatronimic string // ФИО
	Role                  Role   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	RoleId                int    // ID роли
	Email                 string `gorm:"uniqueIndex"` // Почта
	Phone                 string // телефон
	Password              []byte `json:"-"`
	Farm                  *Farm  `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	FarmId                *uint  // ID хозяйства
	Region                Region `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RegionId              uint   `example:"1"` // ID региона
}

func (u *User) Validate() error {
	matchedMail, err := regexp.MatchString(emailRegexp, u.Email)
	if err != nil {
		return err
	}
	if !matchedMail {
		return errors.New("invalid email address")
	}
	matchedPhone, err := regexp.MatchString(phoneRegexp, u.Phone)
	if err != nil {
		return err
	}
	if !matchedPhone {
		return errors.New("invalid phone number")
	}
	return nil
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	return u.Validate()
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	return u.Validate()
}

func GetUserById(id uint) (*User, error) {
	user := User{}
	db := GetDb()
	if err := db.First(&user, id).Error; err != nil {
		return &User{}, err
	}
	return &user, nil
}

func (u *User) AllowedRegions() ([]Region, error) {
	if u.RoleId == 3 || u.RoleId == 4 { // админ и федеральный чиновник могут смотреть любые регионы, региональный чиновник и фермер - только свой
		return []Region{}, nil
	}
	return []Region{u.Region}, nil
}

func (u *User) AllowedDistricts() ([]District, error) {
	districts := []District{}
	if u.RoleId == 3 || u.RoleId == 4 { // админ и федеральный чиновник могут смотреть любые районы, региональный чиновник и фермер - только свой
		return districts, nil
	}
	db := GetDb()
	if err := db.Find(&districts, map[string]any{"region_id": u.RegionId}).Error; err != nil {
		return nil, err
	}
	return districts, nil
}

func (u *User) AllowedFarms() ([]Farm, error) {
	if u.RoleId == 3 || u.RoleId == 4 { // админ и федеральный чиновник могут смотреть любые фермы
		return []Farm{}, nil
	}
	db := GetDb()
	farms := []Farm{}
	if u.RoleId == 2 { // региональный чиновник только фермы своего региона
		err := db.Where("EXISTS (SELECT 1 FROM districts WHERE (districts.id = farms.district_id AND districts.region_id = ?))", u.RegionId).Find(&farms).Error
		if err != nil {
			return nil, err
		}
	}
	if u.RoleId == 1 { // фермер видит только фермы своего холдинга/хозяйства
		err := db.Where("type in (1, 2) AND (parrent_id = ? or farm_id = ?)", u.FarmId, u.FarmId).Find(&farms).Error
		if err != nil {
			return nil, err
		}
	}
	return farms, nil
}
