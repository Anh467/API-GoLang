package dao

import (
	"github.com/jinzhu/gorm"
)

type ProductDao struct {
	Db *gorm.DB
}

/*func (dao *ProductDao) FindAll() ([]model.Customer, error) {
	query := "SELECT CustomerID, Account, Password, FullName, Mail, DateOfBirth, Gender, RoleCustomer, TimeEnjoy" + "\n" +
		"FROM dbo.CUSTOMER"

	rows, err := dao.Db.Query(query)
	if err != nil {
		fmt.Print("Error to get all customers")
		return nil, err
	}
	var customers []model.Customer
	for rows.Next() {
		var CustomerID string
		var Account string
		var Password string
		var FullName string
		var Mail string
		var DateOfBirth string
		var Gender string
		var RoleCustomer string
		var TimeEnjoy string
		err2 := rows.Scan(&CustomerID, &Account, &Password, &FullName, &Mail, &DateOfBirth, &Gender, &RoleCustomer, &TimeEnjoy)
		if err2 != nil {
			return customers, err2
		}
		customer := model.Customer{
			CustomerID:   CustomerID,
			Account:      Account,
			Password:     Password,
			FullName:     FullName,
			Mail:         Mail,
			DateOfBirth:  DateOfBirth,
			Gender:       Gender,
			RoleCustomer: RoleCustomer,
			TimeEnjoy:    TimeEnjoy,
		}
		customers = append(customers, customer)
	}
	return customers, nil
}
func PrintSqlResult(result sql.Result) {
	lastInsertID, _ := result.LastInsertId()
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Last Insert ID: %d, Rows affected: %d\n", lastInsertID, rowsAffected)
}
func (dao *ProductDao) CreateCustomer(customer *model.CustomerCreation) error {
	query := "INSERT INTO dbo.CUSTOMER" +
		" VALUES" +
		" (   ?,   -- Account - varchar(155)" +
		" ?,   -- Password - varchar(155)" +
		" ?,   -- FullName - nvarchar(255)" +
		" ?,   -- Mail - varchar(255)" +
		" ?,   -- DateOfBirth - date" +
		" ?,   -- Gender - varchar(20)" +
		" 'Customer',   -- RoleCustomer - varchar(20)" +
		" default -- TimeEnjoy - datetime" +
		"  )"
	resuult, err := dao.Db.Exec(query, customer.Account, customer.Password, customer.FullName, customer.Mail,
		customer.DateOfBirth, customer.Gender)
	PrintSqlResult(resuult)
	return err
}*/
