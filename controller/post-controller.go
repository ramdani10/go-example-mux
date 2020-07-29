package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ramdani10/mux-example/entity"
	"github.com/ramdani10/mux-example/services"
	"github.com/ramdani10/mux-example/errors"
)

type controller struct {

}

var (
	postService services.PostService
)

func NewPostController(service services.PostService) PostController  {
	postService = service
	return &controller{}
}

type PostController interface {
	GetPosts(resp http.ResponseWriter, req *http.Request)
	AddPost(resp http.ResponseWriter, req *http.Request)
}

func (*controller) GetPosts(resp http.ResponseWriter, req *http.Request) {

	resp.Header().Set("Content-type","application/json")
	posts, err := postService.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error marshlaing the posts array"})
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)

}

func (*controller) AddPost(resp http.ResponseWriter, req *http.Request)  {

	resp.Header().Set("Content-type","application/json")
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error unmarshalling the request"})
		return
	}

	errl := postService.Validate(&post)
	if errl != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: errl.Error()})
		return
	}

	result, err2 := postService.Create(&post)
	if err2 != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error saving post"})
		return
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result)


}