package main

import (
	"context"
	"github.com/go-co-op/gocron"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	"tod/apikey"
	"tod/data"
	"tod/helper"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/qmgo"
)

func main() {
	port := os.Getenv("PORT")

	//Job Scheduler
	s := gocron.NewScheduler(time.UTC)
	s.Every(27).Minute().Do(func() {
		resp, err := http.Get("https://todapi-rafael.herokuapp.com/api")

		if err != nil {
			go helper.SendNotify("Error", "Error in scheduler")
			log.Fatal(err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			go helper.SendNotify("Error", "Error reading response body")
			log.Fatal(err)
		}

		go helper.SendNotify("Response Job", string(body))
	})

	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: "mongodb+srv://efal:michellecantik@efalapi.s7yfi.mongodb.net/?retryWrites=true&w=majority"})
	if err != nil {
		go helper.SendNotify("Error Connect To MongoDb", err.Error())
		panic(err)
	}
	db := client.Database("todapi")

	//collection
	apikeyColl := db.Collection("apikey")
	dataColl := db.Collection("data")

	repoApikey := apikey.NewRepo(apikeyColl)
	handlerApikey := apikey.NewHandler(repoApikey)

	repoData := data.NewRepository(dataColl)
	handlerData := data.NewHandler(repoData, repoApikey)

	//Gin
	g := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	//Router
	karine := g.Group("/api")

	karine.GET("/", HomeHandler)

	//Apikey
	karine.GET("/apikey/check", handlerApikey.CheckApikeyHandler)
	karine.GET("/apikey", handlerApikey.GetAllApikeyHandler)
	karine.POST("/apikey", handlerApikey.CreateApikeyHandler)
	karine.DELETE("/apikey", handlerApikey.DeleteApikeyHandler)

	//Data
	karine.GET("/data", handlerData.GetDataHandler)
	karine.POST("/data", handlerData.CreateDataHandler)

	//Run
	helper.SendNotify("Server Running", "Server Running On Port "+port)
	s.StartAsync()
	g.Run(port)
}

func HomeHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to TOD API",
	})
}
