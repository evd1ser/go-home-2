package server

import (
	"encoding/json"
	"github.com/evd1ser/go-home-2/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func GetAllItems(writer http.ResponseWriter, request *http.Request) {
	log.Println("Get infos about all items in database")
	db := models.DB

	if len(db) == 0 {
		NewErrorResponse(writer, "No one items found in store back", 403)
		return
	}

	NewSuccessResponse(writer, models.DB, 200)
}

func CreateItem(writer http.ResponseWriter, request *http.Request) {
	log.Println("Creating new item...")

	//step 1 parse id
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("error while parsing happend:", err)
		NewErrorResponse(writer, "do not use parameter ID as uncasted to int type", 400)
		return
	}

	//step 2 parse data
	var item models.Item

	err = json.NewDecoder(request.Body).Decode(&item)

	if err != nil {
		NewErrorResponse(writer, "provideed json file is invalid", 400)
		return
	}

	//step 3 check data for exist
	_, exist := models.FindItemById(id)

	if exist {
		NewErrorResponse(writer, "Ityem with that id already exists", 400)
		return
	}

	item.ID = id
	models.DB = append(models.DB, item)
	NewSuccessResponseWithMessage(writer, item, 201, "Item created")
}

func GetItemById(writer http.ResponseWriter, request *http.Request) {

	id, err := strconv.Atoi(mux.Vars(request)["id"])

	if err != nil {
		log.Println("error while parsing happend:", err)
		NewErrorResponse(writer, "do not use parameter ID as uncasted to int type", 400)
		return
	}

	item, ok := models.FindItemById(id)
	log.Println("Get item with id:", id)
	if !ok {
		NewErrorResponse(writer, "Item with that id not found", 404)
		return
	} else {
		NewSuccessResponse(writer, item, 200)
		return
	}
}

func UpdateItemById(writer http.ResponseWriter, request *http.Request) {

	id, err := strconv.Atoi(mux.Vars(request)["id"])

	if err != nil {
		log.Println("error while parsing happend:", err)
		NewErrorResponse(writer, "do not use parameter ID as uncasted to int type", 400)
		return
	}

	itemPointer, ok := models.FindItemById(id)
	log.Println("Get item with id:", id)

	if !ok {
		log.Println("item not found in data base . id :", id)

		NewErrorResponse(writer, "Item with that id not found", 404)
		return
	}

	err = json.NewDecoder(request.Body).Decode(itemPointer)

	if err != nil {
		NewErrorResponse(writer, "provideed json file is invalid", 400)
		return
	}

	(*itemPointer).ID = id
	NewSuccessResponse(writer, *itemPointer, 200)
}

func DeleteItemById(writer http.ResponseWriter, request *http.Request) {

	id, err := strconv.Atoi(mux.Vars(request)["id"])

	if err != nil {
		log.Println("error while parsing happend:", err)
		NewErrorResponse(writer, "do not use parameter ID as uncasted to int type", 400)
		return
	}

	log.Println("Get item with id:", id)

	if _, ok := models.FindItemById(id); !ok {
		log.Println("item not found in data base . id :", id)
		NewErrorResponse(writer, "Item with that id not found", 404)
		return
	}

	item, _ := models.RemoveItemById(id)
	NewSuccessResponseWithMessage(writer, item, 202, "Item deleted")
}
