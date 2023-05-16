package routes

import (
	"fmt"
	"goal/backend/database"
	"github.com/gin-gonic/gin"
	//"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type AllDoc struct {
	DocId int `json:"docid"`
	DocName string `json:"docname"`;
	DocData string`json:"docdata"`;
	UserId int `json:"userid"`;
}
func getAllUserDocs(c *gin.Context){
	var id string = c.Param("id")
	rows, err := database.Goaldb.DB.Query("SELECT * FROM docsdata WHERE user_id = $1", id);
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(400, gin.H{"msg":"could not get data"});
		return;
	}
	defer rows.Close();
	var allUserDocs[] AllDoc = []AllDoc {};
	for rows.Next() {
		var doc AllDoc;
		err := rows.Scan(&doc.DocId, &doc.DocName, &doc.DocData, &doc.UserId);
		if err != nil {
			c.IndentedJSON(400, gin.H{"msg":"could not get data from postgres"});
			fmt.Println(err)
			return;
		}
		allUserDocs = append(allUserDocs, doc);
	}
	c.IndentedJSON(200, allUserDocs);
}
type Doc struct {
	DocId int `json:"docid"`
	UserId int `json:"userid"`;
	DocName string `json:"docname"`;
	DocData string `json:"docdata"`;
}
func getOneDoc(c *gin.Context){
	var userInput Doc = Doc{};
	var err error = c.BindJSON(&userInput);
	if err != nil {
		c.IndentedJSON(500, gin.H{"msg":"could not get user input"});
		return;
	}
	var currDoc Doc;
	fmt.Printf("%d, %d, %s, %s \n", userInput.DocId, userInput.UserId, userInput.DocName, userInput.DocData)
	var sqlStatement string = "SELECT * FROM docsdata WHERE docname = $1 AND user_id = $2"
	dbErr := database.Goaldb.DB.QueryRow(sqlStatement, userInput.DocName, userInput.UserId).Scan(&currDoc.DocId, &currDoc.DocName, &currDoc.DocData, &currDoc.UserId)
	if dbErr != nil {
		fmt.Println(dbErr)
		c.IndentedJSON(500, gin.H{"msg":"error getting data to postgres"});
		return;
	}
	c.IndentedJSON(200, currDoc);
}

type CreateDocHandler struct {
	NwTitle string `json:"nwtitle"`;
	CurrUser int `json:"curruser"`;
}
func createUserDoc(c *gin.Context){
	var userInput CreateDocHandler;
	var err error = c.BindJSON(&userInput);
	if err != nil {
		c.IndentedJSON(500, gin.H{"msg":"error getting user input"});
		return;
	}
	var sqlStatement string = "INSERT INTO docsdata (docname, user_id) VALUES ($1, $2)";
	data, err := database.Goaldb.DB.Exec(sqlStatement, userInput.NwTitle, userInput.CurrUser);
	if err != nil {
		c.IndentedJSON(500, gin.H{"msg":"error getting data into postgres"});
		return;
	}
	fmt.Println(data);
	c.IndentedJSON(200, gin.H{"msg":"data uploaded to postgres"});
}

func updateDocs(c *gin.Context){//UPDATE profile SET name = 'hiroshi' WHERE id = 2;
	var userInput Doc;
	var err error = c.BindJSON(&userInput);
	if err != nil {
		c.IndentedJSON(500, gin.H{"err":"error getting data user input"});
		return;
	}
	fmt.Printf("%s, ")
	var sqlStatement string = "UPDATE docsdata SET dacdata = $1 WHERE docname = $2 AND user_id = $3;"
	data, dbErr := database.Goaldb.DB.Exec(sqlStatement, userInput.DocData, userInput.DocName, userInput.UserId)
	if dbErr != nil {
		fmt.Printf("%s \n", dbErr);
		c.IndentedJSON(500, gin.H{"msg":"error getting data into postgres"});
		return;
	}
	fmt.Println(data);
	c.IndentedJSON(200, gin.H{"msg":"postgres updated"})
}
func deleteOneDoc(c *gin.Context){
	var userInput CreateDocHandler;
	var err error = c.BindJSON(&userInput);
	if(err != nil){
		c.IndentedJSON(500, gin.H{"msg":"could not get data from user"});
		fmt.Println(err);
		return;
	}
	var sqlStatement string = "DELETE FROM docsdata WHERE docname = $1 AND user_id = $2;";
	data, dbErr := database.Goaldb.DB.Exec(sqlStatement, userInput.NwTitle, userInput.NwTitle);
	if(dbErr != nil){
		c.IndentedJSON(500, gin.H{"msg":"could not delete data from postgres"});
		fmt.Println(dbErr);
		return;
	}
	fmt.Println(data);
	
	c.IndentedJSON(200, gin.H{"success":"doc data deleted"})
}

func DocsRoute(g *gin.RouterGroup){
	g.GET("/getAllUserDoc/:id", getAllUserDocs);
	g.POST("/getOneDoc", getOneDoc);
	g.POST("/createUserDoc", createUserDoc);
	g.POST("/updateUserDoc", updateDocs);
	g.DELETE("/deleteDoc", deleteOneDoc)
}
