package testing_test

import (
	"first-app/models"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTesting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testing Suite")
}

var _ = Describe("universitas", func() {

	user := models.User{
		Name:       "STMIK AMIK BANDUNG",
		ProfilePic: "",
		Email:      "amikom@gmail.com",
		Password:   "1234",
		Bio:        "",
		Link:       "",
		Whatsapp:   "",
		UserType:   "university",
	}
	models.DB.Create(&user)

	university := models.Universitas{
		NamaRektor: "Pak Rudi",
		KtpRektor:  "lorem1.png",
		IsVerified: true,
		Alamat:     "Bandung",
	}

	models.DB.Create(&user.Universitas)

	Expect(university.NamaRektor).To(Equal("Pak Rudi"))
	Expect(university.KtpRektor).To(Equal("lorem1.png"))
	Expect(university.IsVerified).To(Equal(true))
	Expect(university.Alamat).To(Equal("Bandung"))
})
