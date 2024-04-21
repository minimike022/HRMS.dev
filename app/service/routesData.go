package routesData

import (
	"fmt"
	Database "hrms-api/app/database"
	DataModels "hrms-api/app/model"

	"github.com/gofiber/fiber/v2"
)

func GetData(ctx *fiber.Ctx) error {
	applicantsData := make([]DataModels.ApplicantsData, 0)
	//Count Rows in Database Table
	rowsCount, _ := Database.Connect().Query("SELECT COUNT(*) FROM users")
	//Fetch Data from the Database
	rows, err := Database.Connect().Query("SELECT * FROM users")

	if err != nil {
		panic(err.Error())
	 }
	 defer rows.Close()

	var count int

	for rowsCount.Next() {
		rowsCount.Scan(&count)
		
	}

	for rows.Next() {
		datas := new(DataModels.ApplicantsData)
		rows.Scan(&datas.Name, &datas.Age, &datas.Address, &datas.Phone_Number)
		applicantsData = append(applicantsData, *datas)
   }

   fmt.Println(applicantsData)
	
	return ctx.Status(fiber.StatusOK).JSON(applicantsData)
}

func PostData(ctx *fiber.Ctx) error {
	applicationData := new(DataModels.ApplicantsData)	
	//var datas DataModels.ApplicantsData
	err := ctx.BodyParser(applicationData)
	if err != nil {
		panic(err.Error())
	}

	datas := DataModels.ApplicantsData{
		Name: applicationData.Name,
		Age: applicationData.Age,
		Address: applicationData.Address,
		Phone_Number: applicationData.Phone_Number,
	}
	  dbQuery, err := Database.Connect().Query("INSERT INTO users(name, age, address, phone_Number) VALUES (?,?,?,?)", datas.Name, datas.Age, datas.Address, datas.Phone_Number)
	  if err != nil {
	    	panic(err.Error())
	  }
	  defer dbQuery.Close()

	return ctx.Status(fiber.StatusOK).JSON(datas)
}

