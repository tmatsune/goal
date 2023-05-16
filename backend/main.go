package main

import (
	"goal/backend/routes"
	"goal/backend/database"
	"github.com/gin-gonic/gin"
	"goal/backend/middlewares"
)

func mnRoute(c *gin.Context){
	c.IndentedJSON(200, gin.H{"msg":"goal web app"});
}


func main(){
	
	router := gin.Default();
	database.Connect(&database.PqConfig);
	router.Use(middlewares.CORSmiddlware())
	
	v1 := router.Group("v1");
	{
		routes.UserRoute(v1.Group("/user"))
		routes.GoalRoute(v1.Group("/goal"))
		routes.DocsRoute(v1.Group("/docs"))
		routes.HabitRoute(v1.Group("/habit"))  
	}
	
	router.GET("/goal", mnRoute);

	router.Run(":8080");
}