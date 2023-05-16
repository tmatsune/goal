package routes

import (
	"fmt"
	"goal/backend/database"
	"goal/backend/models"
	"goal/backend/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

func getGoalTracker(c * gin.Context){
	var userGoalTracker models.GoalTracker;  // = models.GoalTracker{}
	var sqlStatement string = "SELECT * FROM goaltracker WHERE id = $1";
	id := c.Param("id");
	var months[12] pq.Int64Array = [12]pq.Int64Array{} //converts sql array to int64 arrays
	var ids[2] int = [2]int {};
	//var ret pq.Int64Array
	var dbErr error = database.Goaldb.DB.QueryRow(sqlStatement, id).Scan(
		&ids[0], &months[0], &months[1] , &months[2], &months[3], &months[4], &months[5], &months[6], &months[7], &months[8], &months[9], &months[10], &months[11], &ids[1],
	);
	userGoalTracker = models.GoalTracker{
		Id: ids[0],
		Jan: months[0],
		Feb: months[1],
		Mar: months[2],
		Apl: months[3],
		May: months[4],
		Jun: months[5],
		Jul: months[6],
		Aug: months[7],
		Sep: months[8],
		Oct: months[9],
		Nov: months[10],
		Dcm: months[11],
		User_id: ids[1],
	};
	if dbErr != nil {
		c.IndentedJSON(400, gin.H{"msg":"could not get data from postgres"});
		fmt.Println(dbErr)
		return;
	};
	fmt.Println(userGoalTracker);
	c.IndentedJSON(200, userGoalTracker);
}
func createGoalTracker(c *gin.Context){

}

type ChangeGoal struct {
	NwGoal int `json:"nwGoal"`
	UserId int `json:"user_id"`
}
func changeGoalTime(c *gin.Context){
	var userInput ChangeGoal;
	var err error = c.BindJSON(&userInput);
	if err != nil {
		c.IndentedJSON(500, gin.H{"msg":"could not get user input"});
		return;
	}
	var sqlStatement string = "UPDATE profile SET goaltime = $1 WHERE id = $2;";
	data, dbErr := database.Goaldb.DB.Exec(sqlStatement, userInput.NwGoal,userInput.UserId)
	if dbErr != nil {
		c.IndentedJSON(500, gin.H{"msg":"could not update postgres data"});
		fmt.Println(dbErr);
		return;
	}
	fmt.Println(data);
	c.IndentedJSON(200, gin.H{"msg":"successfully updated datat to postgres"})
}
type hoursHandler struct {
	Month string `json:"month"`
	Day int `json:"day"`
	Hours int `json:"hours"`
	User_id int `json:"user_id"`
}
func hoursStudied(c *gin.Context){
	var userInput hoursHandler;
	var err error = c.BindJSON(&userInput);
	if err != nil {
		panic(err)
	}
	var sqlStatment string = middlewares.MonthStatement(userInput.Month) //jan[2] = 3 WHERE user_id=1
	data, dbErr := database.Goaldb.DB.Exec(sqlStatment, userInput.Day,userInput.Hours, userInput.User_id)
	if dbErr != nil {
		panic(dbErr);
	}
	fmt.Print(data);
	c.IndentedJSON(200, gin.H{"msg":"data updted to postgres"});

}

func GoalRoute(g *gin.RouterGroup){
	g.GET("/getGoalTracker/:id", getGoalTracker);
	g.POST("/changeGoal", changeGoalTime);
	g.POST("/hoursStudied", hoursStudied)
}