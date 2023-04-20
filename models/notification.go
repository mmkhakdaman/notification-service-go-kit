package models

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	ID               uint64 `gorm:"primarykey"`
	UserId           uint64 `gorm:"column:user_id"`
	SenderId         uint64 `gorm:"column:sender_id"`
	Title            string
	Message          string
	Avatar           string
	Link             string
	LinkText         string
	Icon             string
	NotificationType string `gorm:"column:type"`
	ObjectId         uint   `gorm:"column:object_id"`
	ObjectType       string `gorm:"column:object_type"`
}
