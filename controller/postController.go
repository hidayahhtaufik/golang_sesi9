package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sesi9/middleware"
	"sesi9/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	if !middleware.Auth(c) {
		return
	}

	var post []models.User
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	fmt.Println(res.Body)
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	defer res.Body.Close()

	c.JSON(http.StatusOK, gin.H{
		"data":   post,
		"status": "OK",
	})
}

func GetUserById(c *gin.Context) {
	if !middleware.Auth(c) {
		return
	}

	var post models.User
	var ID = c.Param("id")
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	fmt.Println(res.Body)
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	defer res.Body.Close()

	c.JSON(http.StatusOK, gin.H{
		"data":   post,
		"status": "OK",
	})
}

func AddUser(c *gin.Context) {
	if !middleware.Auth(c) {
		return
	}

	var post models.User
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	reqJson, err := json.Marshal(post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(reqJson))
	req.Header.Set("Contect-type", "application/json")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	res, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	json.Unmarshal(body, &post)

	c.JSON(http.StatusOK, gin.H{
		"data":   post,
		"status": "OK",
	})
}
