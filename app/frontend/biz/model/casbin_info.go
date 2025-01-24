package model

type CasbinInfo struct {
	UserId uint32 `json:"user_id" gorm:"primaryKey"`
	Sub    string `json:"sub" gorm:"size=128"`
	Domain string `json:"domain" gorm:"size=128;default='default'"`
}
