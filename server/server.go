package server


import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		initHeaders(w)
		next.ServeHTTP(w, r)
	})
}

func BuildManyItemsResourcePrefix(router *mux.Router, prefix string) {
	router.HandleFunc(prefix, GetAllItems).Methods("GET")
}

func BuildItemResource(router *mux.Router, prefix string) {
	router.HandleFunc(prefix+"/{id}", CreateItem).Methods("POST")
	router.HandleFunc(prefix+"/{id}", GetItemById).Methods("GET")
	router.HandleFunc(prefix+"/{id}", UpdateItemById).Methods("PUT")
	router.HandleFunc(prefix+"/{id}", DeleteItemById).Methods("DELETE")
}

func RegMiddlewares(router *mux.Router)  {
	router.Use(jsonMiddleware)
}
