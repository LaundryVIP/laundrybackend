package laundrybackend

import (
	"fmt"
	"testing"

	"github.com/aiteung/atdb"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCreateNewUserRole(t *testing.T) {
	var userdata User
	userdata.UsernameId = "04052003"
	userdata.Username = "farhanrizki"
	userdata.Nohp = "081213280892"
	userdata.Password = "olkiu567"
	userdata.PasswordHash = "olkiu567"
	userdata.Email = "farhanrizki101010@gmail.com"
	userdata.Role = "user"
	mconn := SetConnection("MONGOSTRING", "LaundryApp")
	CreateNewUserRole(mconn, "user", userdata)
}

func CreateNewUserToken(t *testing.T) {
	var userdata User
	userdata.UsernameId = "04052003"
	userdata.Username = "farhanrizki"
	userdata.Nohp = "081213280892"
	userdata.Password = "olkiu567"
	userdata.PasswordHash = "olkiu567"
	userdata.Email = "farhanrizki101010@gmail.com"
	userdata.Role = "user"

	// Create a MongoDB connection
	mconn := SetConnection("MONGOSTRING", "LaundryApp")

	// Call the function to create a admin and generate a token
	err := CreateUserAndAddToken("", mconn, "user", userdata)

	if err != nil {
		t.Errorf("Error creating user and token: %v", err)
	}
}

func TestGFCPostHandlerUser(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "LaundryApp")
	var userdata User
	userdata.UsernameId = "04052003"
	userdata.Username = "farhanrizki"
	userdata.Nohp = "081213280892"
	userdata.Password = "olkiu567"
	userdata.PasswordHash = "olkiu567"
	userdata.Email = "farhanrizki101010@gmail.com"
	userdata.Role = "user"
	CreateNewUserRole(mconn, "user", userdata)
}

func TestGeneratePasswordHash(t *testing.T) {
	passwordhash := "olkiu567"
	hash, _ := HashPass(passwordhash) // ignore error for the sake of simplicity

	fmt.Println("Password:", passwordhash)
	fmt.Println("Hash:    ", hash)
	match := CheckPasswordHash(passwordhash, hash)
	fmt.Println("Match:   ", match)
}
func TestGeneratePrivateKeyPaseto(t *testing.T) {
	privateKey, publicKey := watoken.GenerateKey()
	fmt.Println(privateKey)
	fmt.Println(publicKey)
	hasil, err := watoken.Encode("olkiu567", privateKey)
	fmt.Println(hasil, err)
}

func TestHashFunction(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "LaundryApp")
	var userdata User
	userdata.Email = "farhanrizki101010@gmail.com"
	userdata.PasswordHash = "olkiu567"

	filter := bson.M{"email": userdata.Email}
	res := atdb.GetOneDoc[User](mconn, "user", filter)
	fmt.Println("Mongo User Result: ", res)
	hash, _ := HashPass(userdata.PasswordHash)
	fmt.Println("Hash Password : ", hash)
	match := CheckPasswordHash(userdata.PasswordHash, res.PasswordHash)
	fmt.Println("Match:   ", match)

}

func TestIsPasswordValid(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "LaundryApp")
	var userdata User
	userdata.Email = "farhanrizki101010@gmail.com"
	userdata.PasswordHash = "olkiu567"

	anu := IsPasswordValidEmail(mconn, "user", userdata)
	fmt.Println(anu)
}

func TestUserFix(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "LaundryApp")
	var userdata User
	userdata.UsernameId = "04052003"
	userdata.Username = "farhanrizki"
	userdata.Nohp = "081213280892"
	userdata.Password = "olkiu567"
	userdata.PasswordHash = "olkiu567"
	userdata.Email = "farhanrizki101010@gmail.com"
	userdata.Role = "user"
	CreateUser(mconn, "user", userdata)
}

func TestAdminFix(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "LaundryApp")
	var admindata Admin
	admindata.UsernameId = "LaundryVIP2024"
	admindata.Username = "adminlaundry"
	admindata.Password = "adminlaundrypass"
	admindata.PasswordHash = "adminlaundrypass"
	admindata.Email = "laundryvip@gmail.com"
	admindata.Role = "admin"
	CreateAdmin(mconn, "admin", admindata)
}

func TestGeneratePrivateKeyPasetoAdmin(t *testing.T) {
	privateKey, publicKey := watoken.GenerateKey()
	fmt.Println(privateKey)
	fmt.Println(publicKey)
	hasil, err := watoken.Encode("adminlaundrypass", privateKey)
	fmt.Println(hasil, err)
}

func TestLoginn(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "LaundryApp")
	var userdata User
	userdata.Email = "farhanrizki101010@gmail.com"
	userdata.PasswordHash = "olkiu567"
	IsPasswordValidEmail(mconn, "user", userdata)
	fmt.Println(userdata)
}
