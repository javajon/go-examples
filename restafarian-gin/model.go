package main

import (
	"fmt"
	"strconv"
)

type Idea struct {
	ID     string `json:"ID"`
	Topic  string `json:"Title"`
	Idea   string `json:"Description"`
}

type allIdeas []Idea

var ideas = allIdeas{
	{
		ID:		"1",
		Topic:  "Dinner",
		Idea:   "Fondue: Gruyere, Emmentaler, Kirsch, White wine, garlic, corn starch",
	},
}

func findIdea(id string) (*Idea, int, error) {
	for index, idea := range ideas {
		if idea.ID == id {
			return &idea, index, nil
		}
	}

    return nil, -1, fmt.Errorf("Idea ID %s not found.", id)
}

func addIdea(topic string, idea string) {
	var newIdea Idea
	
	newIdea.ID = strconv.Itoa(len(ideas) + 1) 
	newIdea.Topic = topic 
	newIdea.Idea = idea 
	
	ideas = append(ideas, newIdea)
}

func updateIdea(id string, topic string, idea string) (error) {
	_, index, err := findIdea(id)
	if err == nil {
	    ideas[index].Topic = topic
		ideas[index].Idea = idea
	}

	return err
}

func deleteIdea(id string) (error) {
	_, index, err := findIdea(id)
	if err == nil {
		ideas = append(ideas[:index], ideas[index+1:]...)
	}

	return err
}
