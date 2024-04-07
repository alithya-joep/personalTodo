/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/alithya-joep/personalTodo/cmd"
	"github.com/alithya-joep/personalTodo/dict"
)

func main() {
	cmd.Execute()

	myTask := dict.Task{}
	myTask.Title = "My first task"
	myTask.Description = "This is my first task"

}
