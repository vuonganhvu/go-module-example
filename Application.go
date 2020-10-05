package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Article struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func main() {
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	fmt.Println("Begin Application")
	handleRequest()
}

func handleRequest() {
	appRouter := mux.NewRouter().StrictSlash(true)
	appRouter.HandleFunc("/", homePage)
	appRouter.HandleFunc("/articles", getAllArticles).Methods("GET")
	appRouter.HandleFunc("/articles/{id}", getArticleDetail).Methods("GET")
	appRouter.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")
	appRouter.HandleFunc("/articles", createArticle).Methods("POST")
	appRouter.HandleFunc("/articles/{id}", updateArticle).Methods("PUT")
	log.Fatal(http.ListenAndServe(":1000", appRouter))
}

func updateArticle(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	reqBody, _ := ioutil.ReadAll(request.Body)
	var article Article
	_ = json.Unmarshal(reqBody, &article)
	for  i := range Articles {
		if Articles[i].Id == id {
			pArticle := &Articles[i]
			pArticle.Title = article.Title
			pArticle.Content = article.Content
			pArticle.Desc = article.Desc
		}
	}
	log.Println("End update articles")
}

func createArticle(writer http.ResponseWriter, request *http.Request) {
	reqBody, _ := ioutil.ReadAll(request.Body)
	var article Article
	_ = json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	_ = json.NewEncoder(writer).Encode(article)
	writer.Header().Set("Content-Type","application/json")
}

func getArticleDetail(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			log.Println(json.NewEncoder(w).Encode(article))
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

func homePage(w http.ResponseWriter, request *http.Request)  {
	log.Println("Call home page ", &request)
	fmt.Fprint(w, "Hello home page")
	log.Println("End home page")
}

func getAllArticles(w http.ResponseWriter, r *http.Request){
	log.Println("Endpoint Hit: returnAllArticles")
	log.Println(json.NewEncoder(w).Encode(Articles))
	w.Header().Set("Content-Type","application/json")
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
	w.Header().Set("Content-Type","application/json")
}
