package connections

import (
	"database/sql"
	"fmt"
	"log"
	"model"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func ConnectSqlServerTest() {
	/*server := "MSI\\SQLEXPRESS:1433"
	user := "viet080702"
	password := "nguyenanhviet"
	database := "QuizIrcnV"
	encrypt := "false"*/
	// Nếu bạn không sử dụng mã hóa, đặt thành "false"
	//db, err = sql.Open("mssql", "server=localhost; user id=sa; password=123456")
	// Tạo chuỗi kết nối
	connString := "server=localhost; user id=viet080702; password=nguyenanhviet"

	// Kết nối đến SQL Server
	db, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Lỗi khi kết nối đến SQL Server:", err)
	}
	defer db.Close()

	// Kiểm tra kết nối
	err = db.Ping()
	if err != nil {
		log.Fatal("Lỗi khi kiểm tra kết nối:", err)
	}

	fmt.Println("Kết nối thành công đến SQL Server!")
}

/*func ConnectSqlServer(Config model.SqlServer) (*sql.DB, error) {

	// connString := fmt.Sprintf("sqlserver://%s:%s@%s/instance?database=%s&TrustServerCertificate=true",
	// Config.User, Config.Password, Config.ServerName, Config.DatabaseName)
	// connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
	// Config.ServerName, Config.User, Config.Password, 1433, Config.DatabaseName)

	connString := fmt.Sprintf("server=localhost; user id=%s; password=%s; database=%s;",
		Config.User, Config.Password, Config.DatabaseName)
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	db, err := sql.Open("mssql", connString)
	if err != nil {
		fmt.Print("connection wrong " + "\n")
		log.Fatal(err)

	}

	// Kiểm tra kết nối
	err = db.Ping()
	if err != nil {
		fmt.Print("check connection wrong ")
		log.Fatal(err)

	}

	fmt.Println("Kết nối thành công!")
	return db, err
}*/

func ConnectSqlServerGorm(Config model.SqlServer) (*gorm.DB, error) {
	connString := fmt.Sprintf("server=localhost; user id=%s; password=%s; database=%s;",
		Config.User, Config.Password, Config.DatabaseName)
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		fmt.Print("connection wrong " + "\n")
		log.Fatal(err)
	}
	// Kiểm tra kết nối
	err = db.Error
	if err != nil {
		fmt.Print("check connection wrong ")
		log.Fatal(err)
	}
	fmt.Println("Kết nối thành công!")
	return db, err
}
