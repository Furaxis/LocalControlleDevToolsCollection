package schemas

import (
	"gorm.io/gorm"
)

type ScPostagem struct {
	gorm.Model
	Role    string
	Company string
	link    string
}
