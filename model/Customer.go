package model

type Customer struct {
	ID           int    `json:"id" gorm:"column:ID"`
	CustomerID   string `json:"customerID" gorm:"column:CustomerID"`
	Account      string `json:"account" gorm:"column:Account"`
	Password     string `json:"password" gorm:"column:Password"`
	FullName     string `json:"fullName" gorm:"column:FullName"`
	Mail         string `json:"mail" gorm:"column:Mail"`
	DateOfBirth  string `json:"dateOfBirth" gorm:"column:DateOfBirth"`
	Gender       string `json:"gender" gorm:"column:Gender"`
	RoleCustomer string `json:"roleCustomer" gorm:"column:RoleCustomer"`
	TimeEnjoy    string `json:"timeEnjoy" gorm:"column:TimeEnjoy"`
}

type CustomerCreation struct {
	Account     string `json:"account" gorm:"column:Account"`
	Password    string `json:"password" gorm:"column:Password"`
	FullName    string `json:"fullName" gorm:"column:FullName"`
	Mail        string `json:"mail" gorm:"column:Mail"`
	DateOfBirth string `json:"dateOfBirth" gorm:"column:DateOfBirth"`
	Gender      string `json:"gender" gorm:"column:Gender"`
}
type CustomerUpdate struct {
	Password    *string `json:"password" gorm:"column:Password"`
	FullName    *string `json:"fullName" gorm:"column:FullName"`
	Mail        *string `json:"mail" gorm:"column:Mail"`
	DateOfBirth *string `json:"dateOfBirth" gorm:"column:DateOfBirth"`
	Gender      *string `json:"gender" gorm:"column:Gender"`
}

func (CustomerCreation) TableName() string { return "CUSTOMER" }

func (Customer) TableName() string { return CustomerCreation{}.TableName() }

func (CustomerUpdate) TableName() string { return CustomerCreation{}.TableName() }
