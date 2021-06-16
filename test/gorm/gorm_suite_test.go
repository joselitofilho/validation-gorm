package gorm_test

import (
	"fmt"
	"os"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var gormDB *gorm.DB

var _ = BeforeSuite(func() {
	By("Connecting to the ArangoDB", func() {
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBNAME"), os.Getenv("SSLMODE"))
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		Expect(err).NotTo(HaveOccurred())
		gormDB = db
	})
})

func TestGorm(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gorm Suite")
}
