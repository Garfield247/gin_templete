package models

import "golang.org/x/crypto/bcrypt"

const PasswordCyptLevel = 12

type Account struct {
	BaseModel
	Username string `gorm:"column:username;not null;unique_index;comment:'用户名'" json:"username" form:"username"`
	Password string `gorm:"column:password;comment:'密码'" form:"password" json:"-"`
	Name     string `form:"name" json:"name"`
	IsActive bool   `json:"-"`
}

//获取表名的方法
func (a *Account) TableName() string {
	return "user_accounts"
}

// 封装根据ID获取用户的方法
func (a *Account) GetUserByID(id uint) *Account {
	DB.Model(&Account{}).First(a, id)
	if a.ID > 0 {
		return a
	} else {
		return nil
	}
}

//加密密码
func (a *Account) SetPassword(password string) error {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCyptLevel)
	if err != nil {
		return err
	}
	a.Password = string(pwd)
	return nil
}

// 校验密码合法性
func (a *Account) CheckPassword() bool {
	pwd := a.Password
	DB.Where("username = ?", a.Username).First(&a)
	err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(pwd))
	return err == nil
}

// 验证登录密码是否正确
func (a *Account) IsPasswordRqual(pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(pwd))
	return err == nil
}

//验证用户名重复
func (a *Account) ChaeckDuplicateUsername() bool {
	var count int
	if DB.Model(&Account{}).Where("username = ?", a.Username).Count(&count); count > 0 {
		return false
	} else {
		return true
	}
}
