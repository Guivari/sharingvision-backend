package models

import (
	"time"
	"github.com/guivari/sharingvision-backend/pkg/config"
	"github.com/jinzhu/gorm"
)

type PostStatus string

const (
	Publish PostStatus = "Publish"
	Draft   PostStatus = "Draft"
	Thrash  PostStatus = "Thrash"
)

var db *gorm.DB

type Post struct {
	gorm.Model
	Title       	string    	`gorm:"size:200" json:"title"`
	Content     	string    	`gorm:"" json:"content"`
	Category    	string    	`gorm:"size:100" json:"category"`
	Created_date 	time.Time 	`gorm:"autoCreateTime" json:"created_date"`
	Updated_date 	time.Time 	`gorm:"autoUpdateTime" json:"updated_date"`
	Status      	PostStatus	`gorm:"type:enum('Publish','Draft','Thrash'); default:'Draft'" json:"status"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Post{})
}

func (p *Post) CreatePost() *Post {
	db.NewRecord(p)
	db.Create(&p)
	return p
}

func GetPosts() []Post {
	var Posts []Post
	db.Find(&Posts)
	return Posts
}

func GetPostByID(Id int64) (*Post, *gorm.DB) {
	var getPost Post
	db := db.Where("ID=?", Id)
	return &getPost, db
}

func DeletePost(Id int64) Post {
	var post Post
	db.Where("ID=?", Id).Delete(post)
	return post
}

// UpdatePost = get->delete->create