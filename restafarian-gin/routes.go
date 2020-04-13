package main

import (
	"fmt"
	"strings"
	"net/http"
	"github.com/gin-gonic/gin"
)


func GetAll(context *gin.Context) {
	context.JSON(http.StatusOK, ideas)
}

func Get(context *gin.Context) {
	idea, _, err := findIdea(context.Param("id"))
	if err == nil {
		context.JSON(http.StatusOK, idea)
	} else {
		statusNotFound(context, err)
    }
}

func Create(context *gin.Context) {
	topic, idea, err := parseForm(context)

    if err == nil {
        addIdea(topic, idea)
        statusOK(context, "Idea created")
	} else {
        statusUnprocessableEntity(context, err)
	}
}

func Update(context *gin.Context) {
	topic, ideaText, err := parseForm(context)
	if err == nil {
		err := updateIdea(context.Param("id"), topic, ideaText)
		if err == nil {
			statusOK(context, "Idea updated.")
		} else {
			statusNotFound(context, err)
		}
	} else {
        statusUnprocessableEntity(context, err)
	}
}

func Delete(context *gin.Context) {
	err := deleteIdea(context.Param("id"))
	if err == nil {
		statusOK(context, "Idea removed.")
	} else {	
		statusNotFound(context, err)
	}		
}

func datum(context *gin.Context, formKey string) (string) {
	return strings.TrimSpace(context.PostForm(formKey))
}

func parseForm(context *gin.Context) (string, string, error) {
	// Get idea parameters from the POST form
	
	topic := datum(context, "topic")
	idea := datum(context, "idea")
	if len(topic) == 0 || len(idea) == 0 {
		return "", "", fmt.Errorf("Topic and idea text are both required.")
	}

	return topic, idea, nil
}

func statusOK(context *gin.Context, report string) {
	context.JSON(http.StatusOK,
	    gin.H{"status": http.StatusOK, "message": report})
}

func statusNotFound(context *gin.Context, err error) {
    context.JSON(http.StatusNotFound,
	    gin.H{"status": http.StatusNotFound, "error": err})
}

func statusUnprocessableEntity(context *gin.Context, err error) {
	context.JSON(http.StatusUnprocessableEntity,
	    gin.H{"status": http.StatusUnprocessableEntity, "error": err})
}
