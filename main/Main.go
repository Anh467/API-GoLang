package main

import (
	connections "connect"
	"context"
	"fmt"
	"log"
	"model"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/*
	func TestSqlServer() []model.Customer {
		// get config information
		config, err := model.GetConfig()
		if err != nil {
			fmt.Print("Get information of json wrong")
			return nil
		}
		fmt.Println(config)
		// connecto to db
		db, err := connections.ConnectSqlServer(config.Connections.Sqlserver)
		if err != nil {
			fmt.Print("connection wrong")
			return nil
		}
		customerDAO := dao.ProductDao{
			Db: db,
		}
		customers, err := customerDAO.FindAll()
		if err != nil {
			fmt.Print("find custom wrong")
			fmt.Print(err)
			return nil
		}
		return customers
	}
*/
func TestSqlServerGrom() error {
	config, err := model.GetConfig()
	if err != nil {
		fmt.Print("Get information of json wrong")
		return nil
	}
	fmt.Println(config)
	// connecto to db
	db, err := connections.ConnectSqlServerGorm(config.Connections.Sqlserver)
	if err != nil {
		fmt.Print("connection wrong")
		return nil
	}
	fmt.Print(db.Error)

	return nil
}
func GetCheck() bool {
	var check string
	fmt.Print("Bạn có muốn dừng server enter ? (Y/N): ")
	fmt.Scan(&check)
	return strings.ToUpper(check) == "Y"
}
func stopServer() {
	srv := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Simulate a long-running request
			time.Sleep(10 * time.Second)
			w.Write([]byte("Hello, Mấy cưng !"))
		}),
	}

	// Create a channel to receive signals
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a separate goroutine
	go func() {
		log.Printf("Server listening on %s\n", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for a signal to shutdown the server
	sig := <-signalCh
	log.Printf("Received signal: %v\n", sig)

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Shutdown the server gracefully
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v\n", err)
	}

	log.Println("Server shutdown gracefully")
}

// CRUD: Create, Read, Update, Delete
// POST /v1/customer (create new customer)
// GET /v1/customer (list customer) /v1/customer?page=1
// GET /v1/customer/:id (get dettail customer)
// (PUT||PATCH) /v1/customer/:id (update customer by id)
// DELETE /v1/customer/:id
func testGin() *gin.Engine {
	//ustomers := TestSqlServerGrom()
	config, err := model.GetConfig()
	if err != nil {
		fmt.Print("Get information of json wrong")
		return nil
	}

	db, err := connections.ConnectSqlServerGorm(config.Connections.Sqlserver)
	if err != nil {
		fmt.Print("Connection fail")
		return nil
	}
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		customer := v1.Group("/customer")
		{
			customer.POST("", CreateCustomer(db))
			customer.GET("", GetCustomers(db))
			customer.GET("/:id", GetCustomer(db))
			customer.PATCH("/:id", UpdateCustomer(db))
			customer.DELETE("/:id", DeleteCustomer(db))
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hehehe":  "hehehehehhe",
			"message": "nice",
		})
	})
	r.Run()
	return r
}
func CreateCustomer(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.CustomerCreation
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := db.Create(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
func GetCustomer(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var CID = c.Param("id")
		var data = model.Customer{CustomerID: CID}

		if err := db.Where("CustomerID = ?", CID).First(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
				"CID":   CID,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
func UpdateCustomer(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var CID = c.Param("id")
		var data model.CustomerUpdate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := db.Where("CustomerID = ?", CID).Updates(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
func DeleteCustomer(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var CID = c.Param("id")
		var data = model.Customer{CustomerID: CID}
		if err := db.Where("CustomerID = ?", CID).Delete(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
func GetCustomers(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data []model.Customer

		if err := db.Find(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
func main() {
	testGin()
	//connections.ConnectSqlServerTest()
}
