package model

type User struct {
	Id               uint   `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Username         string `gorm:"not null" json:"username"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	DisplayName      string `json:"displayName"`
	Email            string `json:"email"`
	EmailVerified    bool   `json:"emailVerified"`
	PasswordHash     string `json:"passwordHash,omitempty"`
	Phone            string `json:"phone"`
	PhoneVerified    bool   `json:"phoneVerified"`
	TwoFactorEnabled bool   `json:"twoFactorEnabled"`
	Disabled         bool   `json:"disabled"`
	Role             string `json:"role"`
}
