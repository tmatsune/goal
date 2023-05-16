package middlewares

func MonthStatement(month string)(string){
	switch month {
	case "jan":
		var sqlStatment string = "UPDATE goaltracker SET jan[$1] = $2 WHERE user_id = $3;"
		return sqlStatment
	case "feb":
		var sqlStatment string = "UPDATE goaltracker SET feb[$1] = $2 WHERE user_id = $3;"
		return sqlStatment
	case "mar":
		var sqlStatment string = "UPDATE goaltracker SET mar[$1] = $2 WHERE user_id = $3;"
		return sqlStatment
	case "apl":
		var sqlStatment string = "UPDATE goaltracker SET apl[$1] = $2 WHERE user_id = $3;"
		return sqlStatment
	case "may":
		var sqlStatment string = "UPDATE goaltracker SET may[$1] = $2 WHERE user_id = $3;"
		return sqlStatment
	case "jun":
		var sqlStatment string = "UPDATE goaltracker SET jun[$1] = $2 WHERE user_id = $3;"
		return sqlStatment
	case "jul":
		var sqlStatment string = "UPDATE goaltracker SET jul[$1] = $2 WHERE user_id = $3;"
		return sqlStatment
	case "aug":
		var sqlStatment string = "UPDATE goaltracker SET aug[$1] = $2 WHERE user_id = $3;"
		return sqlStatment
	case "sep":
		var sqlStatment string = "UPDATE goaltracker SET sep[$1] = $2 WHERE user_id = $3;"
		return sqlStatment
	case "oct":
		var sqlStatment string = "UPDATE goaltracker SET oct[$1] = $2 WHERE user_id = $3;"
		return sqlStatment
	case "nov":
		var sqlStatment string = "UPDATE goaltracker SET nov[$1] = $2 WHERE user_id = $3;"
		return sqlStatment
	case "dcm":
		var sqlStatment string = "UPDATE goaltracker SET dcm[$1] = $2 WHERE user_id = $3;"
		return sqlStatment
	default:
		return "none"
	}
}
