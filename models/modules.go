package models

type Module struct {
	ID           uint    `json:"id"`
	ParentID     uint    `json:"parent_id"`
	ParentModule *Module `gorm:"foreignKey:ParentID"`
	// ChildModules []*Module `gorm:"-"`
	Name    string `json:"name"`
	Subject string `json:"subject"`
}
