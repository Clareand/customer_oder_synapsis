package model

import "encoding/json"

type Auth struct {
	CustomerID string `json:"customer_id"`
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
	IsTrue     bool   `json:"is_true" gorm:"column:is_true"`
	CustomerID string `json:"customer_id" gorm:"column:customer_id"`
}

type DataUserFromDB struct {
	Credential string `json:"credential" gorm:"column:credential"`
	Password   string `json:"password" gorm:"column:password"`
	CustomerID int    `json:"customer_id" gorm:"column:customer_id"`
}

type Register struct {
	UserName     string `json:"username"`
	Email        string `json:"email"`
	CustomerID   int    `json:"customer_id"`
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
	CustomerID   string `json:"customer_id"`
	RememberMe   bool   `json:"remember_me"`
}
type ResponseDeleteSessionRefreshToken struct {
	IsSuccess bool `json:"is_success"`
}

type User struct {
	CustomerID string `json:"customer_id"`
	// PsgID          string `json:"psg_id"`
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
	CreatedAt string `json:"created_at"`
	Session   string `json:"session_id"`
}

type ResultLogin struct {
	User        User        `json:"user"`
	AccessToken AccessToken `json:"access_token"`
}

// internal users
type GetterUserLogin struct {
	CustomerID   string `json:"customer_id" gorm:"column:customer_id"`
	UserName     string `json:"username" gorm:"column:username"`
	UserEmail    string `json:"customer_email" gorm:"column:custome_email"`
	CreatedAt    string `json:"created_at" gorm:"column:created_at"`
	RefreshToken string `json:"refresh_token" gorm:"column:refresh_token"`
	Session      string `json:"session" gorm:"column:session"`
}

type ResultLoginUserBE struct {
	User        User        `json:"user"`
	AccessToken AccessToken `json:"access_token"`
}

type UserBeStatus struct {
	Status string `json:"status" gorm:"column:status"`
}
type Session struct {
	ID             string          `json:"id"`
	CustomerID     string          `json:"customer_id"`
	Key            string          `json:"key"`
	CreatedAt      string          `json:"created_at"`
	ExpiredAt      string          `json:"expired_at"`
	SessAdditional json.RawMessage `json:"sess_additional"`
	RememberMe     *bool           `json:"remember_me"`
}
