package web

///mysql範例
import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// db 設定
const (
	userName     = "root"
	rootpassword = "a4w941207!!"
	host         = "127.0.0.1"
	port         = "3306"
	dbName       = "userinfo"
)

var (
	Mysqlpath        string
	UserAccoutDouble bool
)

func InitDB() {
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	// "username:password@tcp(host:port)/數據庫?charset=utf8"
	path := strings.Join([]string{userName, ":", rootpassword, "@tcp(", host, ":", port, ")/", dbName, "?charset=utf8"}, "")
	temp := &Mysqlpath
	*temp = path
	fmt.Println(Mysqlpath)

	// 第一個是 driverName 第二個則是 database 的設定 path
	// 也可以用 var DB *sql.DB
	DB, _ := sql.Open("mysql", path)

	// 設定 database 最大連接數
	DB.SetConnMaxLifetime(100)

	//設定上 database 最大閒置連接數
	DB.SetMaxIdleConns(10)

	// 驗證是否連上 db
	if err := DB.Ping(); err != nil {
		fmt.Println("opon database fail:", err)
		return
	}
	fmt.Println("connnect success")

	//ReadFromMysql()

}

func ReadFromMysql() {

	DB, _ := sql.Open("mysql", Mysqlpath)

	rows, err := DB.Query("SELECT * FROM logininfo")

	if err != nil {
		fmt.Println("open fail:", err)
		return
	}
	fmt.Println("rows", rows)

	for rows.Next() {
		var id string
		var password string

		err = rows.Scan(&id, &password)
		fmt.Println(err)
		if err != nil {
			fmt.Println("open fail:", err)
			return
		}

		fmt.Println(id)
		fmt.Println(password)
	}

}

type UserRegisterInfo struct {
	ID           int
	Username     string
	UserPassword string
}

func (user UserRegisterInfo) InsertMysqlUserInfo(userget map[string]string) {
	var infodoublepoint = &UserAccoutDouble
	*infodoublepoint = false

	DB, _ := sql.Open("mysql", Mysqlpath)

	fmt.Println("receive info fromweb", user)

	rows, err := DB.Query("SELECT * FROM logininfo")
	if err != nil {
		fmt.Println("read from mysql failed", err)
	}
	var id int = 1

	for rows.Next() {
		var (
			IDinMysql          int
			UserAccountinMysql string
			Password           string
		)
		err := rows.Scan(&IDinMysql, &UserAccountinMysql, &Password)
		if err != nil {
			fmt.Println("read from mysql failed", err)
		}
		if IDinMysql != 0 {
			id++
		}
		if UserAccountinMysql == user.Username {
			*infodoublepoint = true
		}
		fmt.Println("eachID:", IDinMysql)
		fmt.Println("UserID:", id)
		fmt.Println()
	}

	if *infodoublepoint == false {
		user.ID = id
		fmt.Println("prepare insert info:", user)
		stmt, err := DB.Prepare("insert logininfo set id=?,useraccount=?,password=?")
		if err != nil {
			fmt.Println("insert fail", err)
		}

		res, err := stmt.Exec(user.ID, user.Username, user.UserPassword)
		if err != nil {
			fmt.Println("insert fail", err)
		}
		fmt.Println("insert mysql", res)
	}

}

func GetUserInfo(userget map[string]string) {
	temp := UserRegisterInfo{}
	temp.Username = userget["username"]
	temp.UserPassword = userget["password"]
	fmt.Println("map", userget)

	temp.InsertMysqlUserInfo(userget)

}
