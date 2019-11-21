//测试创建新的表
//数据库里面的数据是通过表的形式存储的。表有列（columns）和行（row）
//其中行是建立一个表头，而行是输入每一行信息 如
/*像这种
id	username	password	created_at
1	johndoe	secret	2019-08-10 12:30:00
*/
package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
    "fmt"
    "time"
)

func main() {
    //数据库是个很真实的东西，你改了就真的改了，程序结束了也不会对数据库造成什么影响，修改的东西仍然保留了下来。
    //Open方法，第一个参数固定为mysql，第二个首先是用户名，然后密码，然后是端口（这里是本地端口）然后是数据库名。会返回一个类型，该类型有各种方法可以对用户访问的这个数据库进行操作，什么创建table，删除之类的。
	// Configure the database connection (always check errors)
	db, err := sql.Open("mysql", "JacksieCheung:15811852133@(127.0.0.1:3306)/JacksieCheung?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	// Initialize the first connection to the database, to see if everything works correctly.
	// Make sure to check the error.
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

    //这里是创建新的table。其实相当于输入了mysql的命令，这里的table叫做users。
	/*{ // Create a new table
		query := `
        CREATE TABLE users (
            id INT AUTO_INCREMENT,
            username TEXT NOT NULL,
            password TEXT NOT NULL,
            created_at DATETIME,
            PRIMARY KEY (id)
        );`
//通过调用db类型的方法实现
		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}*/
//插入新的行
	/*{ // Insert a new user
		username := "JacksieCheung"
		password := "15811852133"
		createdAt := time.Now()
//还是通过调用db类型的方法进行操作
		result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		id, err := result.LastInsertId()
		fmt.Println(id)
	}*/

	{ // Query a single user
		var (//要访问时就建立一些列变量，通过方法scan导出
			id        int
			username  string
			password  string
			createdAt time.Time
		)
//注意table的名字叫做users
		query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
		if err := db.QueryRow(query, 6).Scan(&id, &username, &password, &createdAt); err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, username, password, createdAt)
	}

    /*{ // Query all users
        type user struct {
            id        int
            username  string
            password  string
            createdAt time.Time
        }

        rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
        if err != nil {
            log.Fatal(err)
        }
        defer rows.Close()

        var users []user
        for rows.Next() {//next是迭代查找，这里id从6开始故失败，迭代查找从一开始
            var u user

            err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
            if err != nil {
                log.Fatal(err)
            }
            users = append(users, u)
        }
        if err := rows.Err(); err != nil {
            log.Fatal(err)
        }

        fmt.Println( users)
    }*/
}
