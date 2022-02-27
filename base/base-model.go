package model

import (
	"time"
)

type Model struct {
	createdBy    uint
	updatedBy    uint
	createdAt    time.Time
	updatedAt    time.Time 
}
