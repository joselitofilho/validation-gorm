package gorm_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"validation-gorm/internal/models"
)

var _ = Describe("HasOne", func() {

	var _ = BeforeEach(func() {
		By("dropping collections", func() {
			err := gormDB.Migrator().DropTable(&models.Brand{}, &models.Category{}, &models.Product{})
			Expect(err).NotTo(HaveOccurred())
		})

		By("preparing collections", func() {
			err := gormDB.AutoMigrate(&models.Brand{}, &models.Category{}, &models.Product{})
			Expect(err).NotTo(HaveOccurred())
		})
	})

	var _ = AfterEach(func() {
		err := gormDB.Migrator().DropTable(&models.Brand{}, &models.Category{}, &models.Product{})
		Expect(err).NotTo(HaveOccurred())
	})

	It("inserts a record into the collection", func() {
		product := models.Product{
			Active:      true,
			Description: "product 1 desc",
			Brand:       models.Brand{Name: "brand1", Active: true},
			Category:    models.Category{Name: "category", Active: true},
		}
		tx := gormDB.Create(&product)
		Expect(tx).NotTo(BeNil())
		Expect(tx.Error).To(BeNil())

		firstProduct := models.Product{}
		gormDB.First(&firstProduct, product.ID)
		Expect(firstProduct.ID).To(Equal(product.ID))
	})

	When("brand and category is already associated with another product", func() {
		It("returns an 'duplicate key' error", func() {
			anotherProduct := models.Product{
				Active:      true,
				Description: "product 1 desc",
				Brand:       models.Brand{Name: "brand1", Active: true},
				Category:    models.Category{Name: "category", Active: true},
			}
			gormDB.Create(&anotherProduct)

			firstProduct := models.Product{}
			gormDB.First(&firstProduct, anotherProduct.ID)

			newProduct := models.Product{
				Active:      true,
				Description: "product 1 desc",
				Brand:       anotherProduct.Brand,
				Category:    anotherProduct.Category,
			}
			tx := gormDB.Create(&newProduct)
			Expect(tx).NotTo(BeNil())
			Expect(tx.Error).NotTo(BeNil())
			Expect(tx.Error.Error()).To(Equal("ERROR: duplicate key value violates unique constraint \"products_pkey\" (SQLSTATE 23505)"))
		})
	})
})
