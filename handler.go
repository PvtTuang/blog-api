package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getPosts(w http.ResponseWriter, r *http.Request) {
	var posts []Post
	term := r.URL.Query().Get("term")
	query := DB

	if term != "" {
		like := "%" + term + "%"
		query = query.Where("title ILIKE ? OR content ILIKE ? OR category ILIKE ?", like, like, like)
	}

	if err := query.Find(&posts).Error; err != nil {
		http.Error(w, "ไม่สามารถดึงข้อมูลโพสต์ได้", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID ไม่ถูกต้อง", http.StatusBadRequest)
		return
	}

	var post Post
	if err := DB.First(&post, id).Error; err != nil {
		http.Error(w, "ไม่พบโพสต์", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	var post Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "คำขอไม่ถูกต้อง", http.StatusBadRequest)
		return
	}

	if err := DB.Create(&post).Error; err != nil {
		http.Error(w, "ไม่สามารถสร้างโพสต์ได้", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID ไม่ถูกต้อง", http.StatusBadRequest)
		return
	}

	var post Post
	if err := DB.First(&post, id).Error; err != nil {
		http.Error(w, "ไม่พบโพสต์", http.StatusNotFound)
		return
	}

	var input Post
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "ข้อมูล JSON ไม่ถูกต้อง", http.StatusBadRequest)
		return
	}

	post.Title = input.Title
	post.Content = input.Content
	post.Category = input.Category
	post.Tags = input.Tags

	if err := DB.Save(&post).Error; err != nil {
		http.Error(w, "ไม่สามารถอัปเดตโพสต์ได้", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID ไม่ถูกต้อง", http.StatusBadRequest)
		return
	}

	if err := DB.Delete(&Post{}, id).Error; err != nil {
		http.Error(w, "ไม่สามารถลบโพสต์ได้", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
