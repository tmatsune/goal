package routes

import (
	"fmt"
	"goal/backend/database"
	"goal/backend/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/lib/pq"
)

//INSERT INTO profile (name, username,email, password) VALUES ('dean', 'deanm','dean@gmail.com', 'laker4life', 60);
func createUser(c *gin.Context){
	var userInput models.User = models.User{};
	var inputPtr *models.User = &userInput;
	var err error = c.BindJSON(inputPtr);
	if err != nil {
		c.IndentedJSON(400, gin.H{"msg":"could not get user input"});
		return;
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost);
	if err != nil {
		c.IndentedJSON(400, gin.H{"msg":"error creating hash"});
		return;
	}
	var nwUser models.User = models.User{
		Name: userInput.Name,
		Username: userInput.Username,
		Email: userInput.Email,
		Password: string(hash),
	}
	var sqlInsert string = `INSERT INTO profile (name, username, email, password) values ($1, $2, $3, $4) RETURNING id;`;
	data,err := database.Goaldb.DB.Exec(sqlInsert, nwUser.Name, nwUser.Username, nwUser.Email, nwUser.Password);
	if err != nil {
		c.IndentedJSON(400, gin.H{"msg":"error with postgres"});
		return;
	}
	fmt.Println(data);
	var num int;
	var sqlStatement string = "SELECT id FROM profile WHERE email = $1";
	dbErr := database.Goaldb.DB.QueryRow(sqlStatement, nwUser.Email).Scan(&num);
	if dbErr != nil {
		c.IndentedJSON(400, gin.H{"msg":"errore getting data"});
		return;
	}
	fmt.Println(num)
	var calInsert string = `INSERT INTO goaltracker (user_id) VALUES ($1);`
	goalData, err := database.Goaldb.DB.Exec(calInsert, num);
	fmt.Println(goalData);

	c.IndentedJSON(200, gin.H{"msg":num});

}
func logInUser(c *gin.Context){
	var userInput models.User = models.User{};
	var inputPtr *models.User = &userInput;
	var err error = c.BindJSON(inputPtr);
	if err != nil {
		c.IndentedJSON(500, gin.H{"msg":"err getting user data"});
		return;
	}
	var rgb pq.Int64Array = pq.Int64Array{};
	var liUser models.User;
	var sqlStatement string = "SELECT * FROM profile WHERE email = $1";
	dbErr := database.Goaldb.DB.QueryRow(sqlStatement, userInput.Email).Scan(&liUser.Id, &liUser.Name, &liUser.Username, &liUser.Email, &liUser.Password, &liUser.Goaltime, &rgb, &liUser.Avatar);
	liUser.RGB = append(liUser.RGB, rgb...)
	if dbErr != nil {
		c.IndentedJSON(400, gin.H{"msg":"error with postgres"});
		fmt.Println(dbErr);
		return;
	};
	var hasErr error = bcrypt.CompareHashAndPassword([]byte(liUser.Password),[]byte(userInput.Password));
	if hasErr != nil {
		c.IndentedJSON(404, gin.H{"message":"wrong password"});
		return;
	}
	hidePass(&liUser);
	c.IndentedJSON(200, liUser);
}
func hidePass(u *models.User){
	(*u).Password = "";
}
type rgbHandler struct {
	Rcolors []int `json:"rcolors"`;
	User_id int `json:"user_id"`;
}
func updateRgb(c *gin.Context){
	var userInput rgbHandler = rgbHandler{};
	var err error = c.BindJSON(&userInput);
	if err != nil {
		c.IndentedJSON(400, gin.H{"msg":"could not get user input"});
		return;
	}
	var rgbSqlStatement string = "UPDATE profile SET rgb = $1 WHERE id = $2;";
	data, dbErr := database.Goaldb.DB.Exec(rgbSqlStatement, pq.Array(userInput.Rcolors), userInput.User_id)
	if dbErr != nil {
		c.IndentedJSON(500, gin.H{"msg":"could not update data in pstgres"});
		fmt.Println(dbErr)
		return;
	}
	fmt.Println(data);
	//fmt.Println(userInput.Rcolors)
	c.IndentedJSON(200, gin.H{"msg":"datat uploaded to postgres"})
}
type AvatarInput struct{
	Avatar int `json:"avatar"`;
	UserEmail string `json:"email"`;
}
func updateAvatar(c *gin.Context){
	var userInput AvatarInput;
	var err error = c.BindJSON(&userInput);
	if err != nil {
		c.IndentedJSON(400, gin.H{"msg":"could not get user input"});
		return;
	}
	var sqlStatement string = "UPDATE profile SET avatar = $1 WHERE email = $2;"
	data, dbErr := database.Goaldb.DB.Exec(sqlStatement, userInput.Avatar, userInput.UserEmail)
	if(dbErr != nil){
		c.IndentedJSON(500, gin.H{"err":"could not update data to postgres"})
		fmt.Println(dbErr);
		return;
	}
	fmt.Println(data)
	c.IndentedJSON(200, gin.H{"msg":"datat uploaded to postgres"})
}
func UserRoute(g *gin.RouterGroup){
	g.POST("/createUser", createUser);
	g.POST("/login", logInUser);
	g.POST("/updateRgb", updateRgb);
	g.POST("/updateAvatar", updateAvatar);
}
