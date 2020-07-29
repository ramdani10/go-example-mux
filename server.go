package main

import (
	"fmt"
	"net/http"

	"github.com/ramdani10/mux-example/repository"
	"github.com/ramdani10/mux-example/controller"
	"github.com/ramdani10/mux-example/routes"
	"github.com/ramdani10/mux-example/services"
)

var (

	postRepository repository.PostRepository = repository.NewFirestoreRepository()
	postService services.PostService = services.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter routes.Router = routes.NewChiRouter()
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/", func(resp http.ResponseWriter, request *http.Request){
		fmt.Fprintln(resp, "Up and Running")
	})

	httpRouter.GET("/post", postController.GetPosts)
	httpRouter.POST("/post", postController.AddPost)

	httpRouter.SERVE(port)
}