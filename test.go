package main

import (
	"time"
	"gorm.io/gorm"
)

type Test struct {
gorm.Model
id int
time time.Time
}