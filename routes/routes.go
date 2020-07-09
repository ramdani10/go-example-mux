package routes

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/ramdani10/mux-example/entity"
	"github.com/ramdani10/mux-example/repository"
)

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

//func init () {
//	posts = []Post{Post{Id:1, Title:"Title 1", Text: "Text 1"}}
//}

func GetPosts(resp http.ResponseWriter, req *http.Request) {

	resp.Header().Set("Content-type","application/json")
	posts, err := repo.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error marshlaing the posts array"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)

	//result, err := json.Marshal(posts)
	//
	//resp.Write(result)
}

func AddPost(resp http.ResponseWriter, req *http.Request)  {

	resp.Header().Set("Content-type","application/json")
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error" : "Error unmarshalling the request"}`))
		return
	}

	post.ID = rand.Int63()
	repo.Save(&post)

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(post)


}