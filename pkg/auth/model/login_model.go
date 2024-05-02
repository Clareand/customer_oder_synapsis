package model

import "encoding/json"

type Auth struct {
	UserID string `json:"user_id"`
}

type ReqLogin struct {
	UserName           string          `json:"username"`
	Verified           bool            `json:"verified"`
	Email              string          `json:"email"`
	PartnerID          string          `json:"partner_id"`
	PartnerCredential  string          `json:"partner_credential"`
	MerchantID         int             `json:"merchant_id"`
	MerchantCode       string          `json:"merchant_code"`
	Roles              string          `json:"roles"`
	MerchantAdditional json.RawMessage `json:"merchant_additional"`
}

type ReqNewLogin struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RememberMe bool   `json:"remember_me"`
}

type IsTrue struct {
	IsTrue bool `json:"is_true" gorm:"column:is_true"`
}

type CheckUserIsTrue struct {
	IsTrue bool   `json:"is_true" gorm:"column:is_true"`
	UserId string `json:"user_id" gorm:"column:user_id2"`
}

type DataUserFromDB struct {
	UserCredential string `json:"user_credential" gorm:"column:user_credential"`
	Password       string `json:"password" gorm:"column:password"`
	UserID         int    `json:"user_id" gorm:"column:user_id"`
}

type Register struct {
	UserName     string `json:"username"`
	Email        string `json:"email"`
	UserID       int    `json:"user_id"`
	MerchantID   int    `json:"merchant_id"`
	MerchantCode string `json:"merchant_code"`
	Roles        string `json:"roles"`
	ParentID     int    `json:"parent_id"`
	Status       string `json:"status"`
	Verified     bool   `json:"verified"`
	Address      string `json:"address"`
	Occupation   string `json:"occupation"`
}

type ResultRefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}

type AccessToken struct {
	Type         string `json:"type"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
type CheckSession struct {
	IsNotExpired bool   `json:"is_not_expired"`
	UserId       string `json:"user_id"`
	RememberMe   bool   `json:"remember_me"`
}
type ResponseDeleteSessionRefreshToken struct {
	IsSuccess bool `json:"is_success"`
}

type User struct {
	UserID string `json:"user_id"`
	// PsgID          string `json:"psg_id"`
	UserName       string `json:"user_name"`
	UserEmail      string `json:"user_email"`
	UserActivation string `json:"user_activation"`
	CreatedAt      string `json:"created_at"`
	Session        string `json:"session_id"`
}

type ResultLogin struct {
	User        User        `json:"user"`
	Provider    string      `json:"provider"`
	AccessToken AccessToken `json:"access_token"`
}

// internal users
type GetterUserLogin struct {
	UserID         string `json:"user_id" gorm:"column:user_id2"`
	UserName       string `json:"user_name" gorm:"column:user_name"`
	UserEmail      string `json:"user_email" gorm:"column:user_email"`
	UserActivation string `json:"user_activation" gorm:"column:user_activation"`
	CreatedAt      string `json:"created_at" gorm:"column:created_at"`
	RefreshToken   string `json:"refresh_token" gorm:"column:refresh_token"`
	Session        string `json:"session_id" gorm:"column:session_id"`
}

type ResultLoginUserBE struct {
	User        User        `json:"user"`
	Provider    string      `json:"provider"`
	AccessToken AccessToken `json:"access_token"`
}

type UserBeStatus struct {
	Status string `json:"status" gorm:"column:status"`
}
type Session struct {
	ID             string          `json:"id"`
	UserId         string          `json:"user_id"`
	Key            string          `json:"key"`
	CreatedAt      string          `json:"created_at"`
	ExpiredAt      string          `json:"expired_at"`
	PartnerID      string          `json:"partner_id"`
	SessAdditional json.RawMessage `json:"sess_additional"`
	RememberMe     *bool           `json:"remember_me"`
	GroupApi       json.RawMessage `json:"group_api"`
}
