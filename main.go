package main 

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type book struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Quantity int `json:"quantity"`
}



func main (){

}