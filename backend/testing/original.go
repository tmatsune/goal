package testing
/*
type Todo struct {
	Todo_id int `json:"todo_id"`
	Description string  `json:"description"`
}
func getItem(c *gin.Context){

	rows, err := database.Goaldb.DB.Query("SELECT * FROM todo");
	if err != nil {
		c.IndentedJSON(400, gin.H{"msg":"could not get data"});
		return;
	}
	
	defer rows.Close();
	var allTodos[] Todo = []Todo{};
	for rows.Next() {
		var td Todo;
		err := rows.Scan(&td.Todo_id, &td.Description);
		if err != nil {
			panic(err);
		}
		allTodos = append(allTodos, td);//td
	}
	c.JSON(200, allTodos);
}
func addItem(c *gin.Context){
	sqlStatement := `INSERT INTO todo (description) values ($1);`
	data , err := database.Goaldb.DB.Exec(sqlStatement, "connect golang with postgres");
	if err != nil {
		panic(err);
	}
	fmt.Println(data);
	c.IndentedJSON(200, gin.H{"msg":"data inserted"});
}
func getOne(c *gin.Context){
	var td Todo = Todo{}
	err := database.Goaldb.DB.QueryRow("SELECT * FROM todo WHERE todo_id = 1").Scan(&td.Todo_id, &td.Description)
	if err != nil {
		panic(err);
	}
	c.IndentedJSON(200, td);
}
*/
	/*
	router.GET("/test", getItem );
	router.POST("/add", addItem );
	router.GET("/getOne", getOne );
	*/