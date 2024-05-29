package controllers

import (
	"fmt"
	"github.com/go-ldap/ldap"
	"hitos_back/models"
	"hitos_back/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ldapServer   = "ldap://tu-servidor-ldap:389"
	ldapBindDN   = "cn=usuario,ou=Usuarios,dc=example,dc=com" // Reemplaza con tu información
	ldapBindPass = "contraseña"                               // Reemplaza con tu información
)

func authenticateUser(username, password string) bool {
	l, err := ldap.Dial("tcp", ldapServer)
	if err != nil {
		fmt.Println("Error al conectar al servidor LDAP:", err)
		return false
	}
	defer l.Close()

	err = l.Bind(ldapBindDN, ldapBindPass)
	if err != nil {
		fmt.Println("Error al hacer bind con el servidor LDAP:", err)
		return false
	}

	searchRequest := ldap.NewSearchRequest(
		"dc=example,dc=com", // Reemplaza con tu información
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(sAMAccountName=%s)", username),
		[]string{"dn"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		fmt.Println("Error al realizar la búsqueda en el servidor LDAP:", err)
		return false
	}

	if len(sr.Entries) != 1 {
		fmt.Println("Usuario no encontrado o múltiples coincidencias")
		return false
	}

	userDN := sr.Entries[0].DN

	err = l.Bind(userDN, password)
	if err != nil {
		fmt.Println("Autenticación fallida:", err)
		return false
	}

	fmt.Println("Autenticación exitosa para el usuario:", username)
	return true
}

func validateUser() {
	username := "nombre-de-usuario" // Reemplaza con tu información
	password := "contraseña"        // Reemplaza con tu información

	if authenticateUser(username, password) {
		fmt.Println("Bienvenido,", username)
	} else {
		fmt.Println("Autenticación fallida")
	}
}
func CurrentUser(c *gin.Context) {

	user_id, err := utils.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := GetUserByID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {

	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	//token, err := LoginCheck(u.Username, u.Password)

	if authenticateUser(u.Username, u.Password) {
		fmt.Println("Bienvenido,")
	} else {
		fmt.Println("Autenticación fallida")
	}

	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
	//	return
	//}
	var token = "sdf"
	c.JSON(http.StatusOK, gin.H{"token": token})
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {

	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	_, err := u.SaveUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})

}
