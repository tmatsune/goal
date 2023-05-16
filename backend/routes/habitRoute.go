
package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goal/backend/database"
)

type Habit struct {
	Id int `json:"id"`;
	Habitname string `json:"habitname"`;
	Habitdata string `json:"habitdata"`;
	User_id int `json:"user_id"`
}
func createHabit(c *gin.Context){
	var userInput Habit;
	var err error = c.BindJSON(&userInput);
	if(err != nil){
		c.IndentedJSON(500, gin.H{"msg":"could not get user input"});
		fmt.Printf("%s \n", err)
		return;
	}
	var sqlinsert string = "INSERT INTO habit (habitname, habitdata, user_id) VALUES ($1, $2, $3)";
	data, dberr := database.Goaldb.DB.Exec(sqlinsert, userInput.Habitname, userInput.Habitdata, userInput.User_id)
	if(dberr != nil){
		c.IndentedJSON(500, gin.H{"msg":"could not insert data to postgres"});
		fmt.Printf("%s \n", err)
		return;
	}
	fmt.Println(data);
	c.IndentedJSON(200, gin.H{"msg":"habit added to postres"});
}
func getHabits(c *gin.Context){
	var id string = c.Param("id");
	rows, err := database.Goaldb.DB.Query("SELECT * FROM habit WHERE user_id = $1", id);
	if( err != nil){
		c.IndentedJSON(500, gin.H{"error":"error getting data from postgres"});
		fmt.Println(err);
		return;
	}
	defer rows.Close();
	var allUserHabits[] Habit = []Habit{};
	for (rows.Next()){
		var habit Habit;
		err := rows.Scan(&habit.Id, &habit.Habitname, &habit.Habitdata, &habit.User_id);
		if(err != nil){
			c.IndentedJSON(500, gin.H{"error":"error returning data"});
			fmt.Println(err);
			return;
		}
		allUserHabits = append(allUserHabits, habit);
	}
	c.IndentedJSON(200, allUserHabits);
}

func DeleteHabit(c *gin.Context){
	var userInput Habit = Habit{};
	var err error = c.BindJSON(&userInput);
	if(err != nil){
		c.IndentedJSON(500, gin.H{"msg":"could not get data from user"});
		fmt.Println(err);
		return;
	}
	var sqlStatement string = "DELETE FROM habit WHERE habitdata = $1 AND user_id = $2;";
	data, dbErr := database.Goaldb.DB.Exec(sqlStatement, userInput.Habitdata, userInput.User_id);
	if(dbErr != nil){
		c.IndentedJSON(500, gin.H{"msg":"could not delete data from postgres"});
		fmt.Println(dbErr);
		return;
	}
	fmt.Println(data);
	c.IndentedJSON(200, gin.H{"success":"habit data deleted"})
}
func HabitRoute(g *gin.RouterGroup){
	g.POST("/createHabit", createHabit);
	g.GET("/getAllHabits/:id", getHabits);
	g.POST("/deleteHabit", DeleteHabit)
}