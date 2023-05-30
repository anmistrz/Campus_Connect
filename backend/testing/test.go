package testing

import (
	"first-app/models"

	"github.com/gin-gonic/gin"
)

type Signup struct {
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	NamaRektor string `json:"namaRektor" binding:"required"`
	UserType   string `json:"userType" binding:"required"`
	IsVerified bool   `json:"isVerified" binding:"required"`
	KtpRektor  string `json:"ktpRektor" binding:"required"`
	Alamat     string `json:"alamat" binding:"required"`
}

//signup user universitas
func SignupUser(c *gin.Context) {
	var register Signup
	err := c.BindJSON(&register)
	// if email same error
	if err := models.DB.Where("email = ?", register.Email).First(&models.User{}).Error; err == nil {
		c.JSON(400, gin.H{
			"status code": 400,
			"message":     "Email already exists",
		})
		return
	}
	if register.Name == "" || register.Email == "" || register.Password == "" || register.NamaRektor == "" || register.UserType == "" || register.IsVerified == false || register.KtpRektor == "" || register.Alamat == "" {
		c.JSON(400, gin.H{
			"status code": 400,
			"message":     "Please fill all required fields",
			"error":       err,
		})
		return
	}

	user := models.User{
		Name:       register.Name,
		ProfilePic: "",
		Email:      register.Email,
		Password:   register.Password,
		Bio:        "",
		Link:       "",
		Whatsapp:   "",
		UserType:   register.UserType,
	}
	models.DB.Create(&user)

	university := models.Universitas{
		NamaRektor: register.NamaRektor,
		KtpRektor:  register.KtpRektor,
		IsVerified: register.IsVerified,
		Alamat:     register.Alamat,
	}

	c.JSON(200, gin.H{
		"status code": 200,
		"message":     "Successfully created user",
		"user":        user,
		"university":  university,
	})
}
