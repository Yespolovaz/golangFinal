package users

import (
	"time"

	"github.com/Yespolovaz/golangFinal/pkg/bank/database"
	"github.com/Yespolovaz/golangFinal/pkg/bank/helpers"
	"github.com/Yespolovaz/golangFinal/pkg/bank/interfaces"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func prepareToken(user *interfaces.User) string {
	// Sign JWT token
	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry":  time.Now().Add(time.Minute ^ 60).Unix()}

	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)

	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	helpers.HandleErr(err)

	return token
}

func prepareResponse(user *interfaces.User, accounts []interfaces.ResponseAccount, withToken bool) map[string]interface{} {
	// Setup response
	responseUser := &interfaces.ResponseUser{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Name:     user.Name,
		Surname:  user.Surname,
		Accounts: accounts,
	}

	// Prepare response
	var response = map[string]interface{}{"message": "Success"}
	if withToken {
		var token = prepareToken(user)
		response["jwt"] = token
	}
	response["data"] = responseUser

	return response
}

func Login(username string, pass string) map[string]interface{} {
	// Validate
	valid := helpers.Validation(
		[]interfaces.Validation{
			{Value: username, Valid: "username"},
			{Value: pass, Valid: "password"}})

	if valid {
		// Connect db
		user := &interfaces.User{}
		if database.DB.Where("username = ?", username).First(&user).RecordNotFound() { // FIXME: - bug
			return map[string]interface{}{"message": "User not found"}
		}

		// Verify password
		passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
		if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
			return map[string]interface{}{"message": "Wrong password"}
		}

		// Find acc for the user
		accounts := []interfaces.ResponseAccount{}
		database.DB.Table("accounts").Select("id, name, balance").Where("user_id = ?", user.ID).Scan(&accounts)

		// Close connection to db
		var response = prepareResponse(user, accounts, true)

		return response
	} else {
		return map[string]interface{}{"message": "Not valid authentication values"}
	}
}

func Register(username string, email string, pass string) map[string]interface{} {
	// Validate
	valid := helpers.Validation(
		[]interfaces.Validation{
			{Value: username, Valid: "username"},
			{Value: email, Valid: "email"},
			{Value: pass, Valid: "password"}})

	if valid {

		generatedPassword := helpers.HashAndSalt([]byte(pass))

		user := &interfaces.User{
			Username: username,
			Email:    email,
			Password: generatedPassword}
		database.DB.Create(&user)

		account := &interfaces.Account{
			Type:    "Account",
			Name:    string(username + "'s" + " account"),
			Balance: 0,
			UserId:  user.ID}
		database.DB.Create(&account)

		accounts := []interfaces.ResponseAccount{}
		responseAccount := interfaces.ResponseAccount{
			Id:      account.ID,
			Name:    account.Name,
			Balance: int(account.Balance)}

		accounts = append(accounts, responseAccount)

		var response = prepareResponse(user, accounts, true)

		return response
	} else {
		return map[string]interface{}{"message": "Not valid authentication values"}
	}
}

func GetUser(id string, jwt string) map[string]interface{} {
	isValid := helpers.ValidateToken(id, jwt)

	if isValid {

		user := &interfaces.User{}
		if database.DB.Where("id = ?", id).First(&user).RecordNotFound() {
			return map[string]interface{}{"message": "User not found"}
		}
		accounts := []interfaces.ResponseAccount{}
		database.DB.Table("accounts").Select("id, name, balance").Where("user_id = ?", user.ID).Scan(&accounts)

		var response = prepareResponse(user, accounts, false)

		return response
	} else {
		return map[string]interface{}{"message": "Not valid JWT token"}
	}
}

func DeleteUSer(id string, jwt string) map[string]interface{} {
	isValid := helpers.ValidateToken(id, jwt)

	if isValid {

		user := &interfaces.User{}
		if database.DB.Where("id = ?", id).First(&user).RecordNotFound() {
			return map[string]interface{}{"message": "User not found"}
		}
		if err := database.DB.Where("id = ?", id).Delete(&interfaces.User{}).Error; err != nil {
			return map[string]interface{}{"error": err.Error()}
		}

		return map[string]interface{}{"message": "User deleted successfully"}
	} else {

		return map[string]interface{}{"message": "Not valid JWT token"}

	}
}

func ListUsers() ([]interfaces.User, error) {
	var users []interfaces.User
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
