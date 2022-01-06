package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Model struct {
	ID        uint  `gorm:"primaryKey"`
	CreateAt  int64 `gorm:"autoCreateTime"`
	UpdatedAT int64 `gorm:"autoCreateTime"`
}

type User struct {
	gorm.Model
	UserName     string
	Email        string
	PassWord     string
	Birthday     time.Time
	ImageUrl     string
	IsActive     bool
	ProjectUsers []ProjectUser
}
type Project struct {
	gorm.Model
	CategorieID       int
	CompanieID        int
	ProjectName       string
	ProjectedSpend    int
	ProjectedVariance int
	RevenueRecognised int
	ProjectUsers      []ProjectUser `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Categorie struct {
	gorm.Model
	CategoryName string
	Projects     []Project
}
type Companie struct {
	gorm.Model
	CompanyName string
	CompanyCode string
	Projects    []Project
}

type ProjectUser struct {
	gorm.Model
	ProjectID int
	UserID    int
}

func main() {
	//var ts Bangs
	// testt := Bangs{
	// 	Nam:    "Tran B",
	// 	Oldd:   18,
	// 	Adress: "abc",
	// }
	dsn := "Bangbb:TranBang11@@tcp(127.0.0.1:3306)/NewGoLang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Println("connect Fail!!")
		return
	} else {
		fmt.Println("connect Sucsessful!")
		// result := db.Migrator().CreateTable(&Hello{})
		// result := db.Migrator().DropTable(&Hello{})
		// fmt.Println(result)
		db.Migrator().CreateTable(&User{})
		// db.Migrator().DropTable(&User{})

		db.Migrator().CreateTable(&Companie{})
		// db.Migrator().DropTable(&Companie{})

		db.Migrator().CreateTable(&Categorie{})
		// db.Migrator().DropTable(&Categorie{})

		db.Migrator().CreateTable(&Project{})
		// db.Migrator().DropTable(&Project{})

		db.Migrator().CreateTable(&ProjectUser{})
		// db.Migrator().DropTable(&ProjectUser{})

		db.Migrator().CreateConstraint(&User{}, "ProjectUsers")
		//db.Migrator().CreateConstraint(&User{}, "fk_users_project_users")

		db.Migrator().CreateConstraint(&Project{}, "ProjectUsers")
		//db.Migrator().CreateConstraint(&Project{}, "fk_projects_project_users")

		db.Migrator().CreateConstraint(&Categorie{}, "Projects")
		//		db.Migrator().CreateConstraint(&Categorie{}, "fk_categories_projects")

		db.Migrator().CreateConstraint(&Companie{}, "Projects")
		//		db.Migrator().CreateConstraint(&Companie{}, "fk_companises_projects")

	}

}
