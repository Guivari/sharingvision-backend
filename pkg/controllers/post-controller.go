package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/guivari/sharingvision-backend/pkg/models"
	"github.com/guivari/sharingvision-backend/pkg/utils"
)

var NewPost models.Post

func CreatePost(w http.ResponseWriter, req *http.Request) {
	createPost := &models.Post{}
	utils.ParseBody(req, createPost)

	p := createPost.CreatePost()
	res, _ := json.Marshal(p)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetPost(w http.ResponseWriter, req *http.Request) {
	newPosts := models.GetPosts
	res, _ := json.Marshal(newPosts)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetPostByID(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	postId := vars["postId"]
	ID, err := strconv.ParseInt(postId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	postContent, _ := models.GetPostByID(ID)
	res, _ := json.Marshal(postContent)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeletePost(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	postId := vars["postId"]
	ID, err := strconv.ParseInt(postId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	post := models.DeletePost(ID)
	res, _ := json.Marshal(post)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdatePost(w http.ResponseWriter, req *http.Request) {
	var UpdatePost = &models.Post{}
	utils.ParseBody(req, UpdatePost)

	vars := mux.Vars(req)
	postId := vars["postId"]
	ID, err := strconv.ParseInt(postId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	
	postContent, db := models.GetPostByID(ID)
	if UpdatePost.Title != "" {
		postContent.Title = UpdatePost.Title
	}
	if UpdatePost.Content != "" {
		postContent.Content = UpdatePost.Content
	}
	if UpdatePost.Category != "" {
		postContent.Category = UpdatePost.Category
	}
	postContent.Updated_date = UpdatePost.Updated_date
	postContent.Status = UpdatePost.Status

	db.Save(&postContent)
	res, _ := json.Marshal(postContent)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

