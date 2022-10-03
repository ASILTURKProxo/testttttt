package models

type Ability struct {
	ID         uint    `json:"id"`
	Action     string  `json:"action"`
	Subject    string  `json:"subject"`
	ModuleID   uint    `json:"module_id"`
	Module     *Module `gorm:"foreignKey:ModuleID;" json:"-"`
	ModuleName string  `gorm:"column:module_name"`
	ModuleKey  string  `gorm:"column:module_key"`
}

type User struct {
	ID         uint                   `json:"id"`
	Name       string                 `json:"name"`
	Email      string                 `gorm:"unique" json:"email"`
	Password   []byte                 `json:"-"`
	UserRole   []Role                 `gorm:"many2many:user_roles;" json:"user_role"`
	Ability    []Ability              `gorm:"many2many:user_abilities;" json:"-"`
	AbilityMap map[string]interface{} `gorm:"-" json:"ability_map"`
}

type Role struct {
	ID      uint      `json:"id"`
	Name    string    `gorm:"type:character varying(50)" json:"name"`
	Ability []Ability `gorm:"many2many:roles_abilities;" json:"ability"`
}
