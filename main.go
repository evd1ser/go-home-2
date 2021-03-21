package main
import (
	"github.com/evd1ser/go-home-2/server"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const (
	apiPrefix string = "/api/v1"
)

var (
	port                    string
	itemResourcePrefix      string = apiPrefix + "/item"  //api/v1/book/
	manyItemsResourcePrefix string = apiPrefix + "/items" //api/v1/books/
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not find .env file:", err)
	}
	port = os.Getenv("app_port")
}

func main() {
	log.Println("Starting REST API server on port:", port)
	router := mux.NewRouter()

	server.BuildItemResource(router, itemResourcePrefix)
	server.BuildManyItemsResourcePrefix(router, manyItemsResourcePrefix)

	server.RegMiddlewares(router)

	log.Println("Router initalizing successfully. Ready to go!")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
