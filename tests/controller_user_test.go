package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/zineb-ada/cyrkl/api/models"
)

func TestCreateUser(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}
	samples := []struct {
		inputJSON      string
		statusCode     int
		name           string
		lastname       string
		email          string
		urlphoto       string
		telephone      string
		position       string
		positionsought pq.StringArray
		industry       string
		industrysought pq.StringArray
		errorMessage   string
	}{
		{
			inputJSON: `{
				"name": "laurie",
				"lastname": "clu",
				"email": "laurie-clu@gmail.com",
				"urlphoto": "thisisanurl",
				"telephone": "0876756788",
				"password": "password",
				"position": "work_position",
				"positionsought": ["pos1", "pos2", "pos3", "pos4"],
				"industry": "industry",
				"industrysought": ["ind1", "ind2", "ind3", "ind4"]
				}`,
			statusCode:     201,
			name:           "laurie",
			lastname:       "clu",
			email:          "laurie-clu@gmail.com",
			urlphoto:       "thisisanurl",
			telephone:      "0876756788",
			position:       "work_position",
			positionsought: []string{"pos1", "pos2", "pos3", "pos4"},
			industry:       "industry",
			industrysought: []string{"ind1", "ind2", "ind3", "ind4"},
			errorMessage:   "",
		},
		{
			inputJSON:    `{
				"name": "lolo",
				"lastname": "cluclu",
				"email": "laurie-clu@gmail.com",
				"urlphoto": "thisisanurl",
				"telephone": "0876756688",
				"password": "password",
				"position": "work_position",
				"positionsought": ["pos1", "pos2", "pos3", "pos4"],
				"industry": "industry",
				"industrysought": ["ind1", "ind2", "ind3", "ind4"]
				}`,
			statusCode:   500,
			errorMessage: "Email Already Taken",
		},
		{
			inputJSON:    `{
				"name": "ami",
				"lastname": "grand",
				"email": "grand-ami@gmail.com",
				"urlphoto": "thisisanurl",
				"telephone": "0876756788",
				"password": "password",
				"position": "work_position",
				"positionsought": ["pos1", "pos2", "pos3", "pos4"],
				"industry": "industry",
				"industrysought": ["ind1", "ind2", "ind3", "ind4"]
				}`,
			statusCode:   500,
			errorMessage: "Telephone Already Taken",
		},
		{
			inputJSON:    `{
				"name": "lala",
				"lastname": "marie",
				"email": "lala_mariegmail.com",
				"urlphoto": "thisisanurl",
				"telephone": "0576756788",
				"password": "password",
				"position": "work_position",
				"positionsought": ["pos1", "pos2", "pos3", "pos4"],
				"industry": "industry",
				"industrysought": ["ind1", "ind2", "ind3", "ind4"]
				}`,
			statusCode:   422,
			errorMessage: "Invalid Email",
		},
		{
			inputJSON:    `{
				"name": "",
				"lastname": "arnoud",
				"email": "fred_arnoud@gmail.com",
				"urlphoto": "thisisanurl",
				"telephone": "0576736758",
				"password": "password",
				"position": "CTO",
				"positionsought": ["pos1", "pos2", "pos3", "pos4"],
				"industry": "tech",
				"industrysought": ["ind1", "ind2", "ind3", "ind4"]
				}`,
			statusCode:   422,
			errorMessage: "Required Name",
		},
		{
			inputJSON:    `{
				"name": "fred",
				"lastname": "",
				"email": "fred@gmail.com",
				"urlphoto": "thisisanurl",
				"telephone": "0576756558",
				"password": "password",
				"position": "CTO",
				"positionsought": ["pos1", "pos2", "pos3", "pos4"],
				"industry": "tech",
				"industrysought": ["ind1", "ind2", "ind3", "ind4"]
				}`,
			statusCode:   422,
			errorMessage: "Required Lastnam",
		},
		{
			inputJSON:    `{
				"name": "mimi",
				"lastname": "lg",
				"email": "",
				"urlphoto": "thisisanurl",
				"telephone": "0572756558",
				"password": "password",
				"position": "CTO",
				"positionsought": ["pos1", "pos2", "pos3", "pos4"],
				"industry": "tech",
				"industrysought": ["ind1", "ind2", "ind3", "ind4"]
				}`,
				
			statusCode:   422,
			errorMessage: "Required Email",
		},
		{
			inputJSON:   `{
				"name": "fred",
				"lastname": "",
				"email": "fred@gmail.com",
				"urlphoto": "thisisanurl",
				"telephone": "0576756558",
				"password": "password",
				"position": "CTO",
				"positionsought": ["pos1", "pos2", "pos3", "pos4"],
				"industry": "tech",
				"industrysought": ["ind1", "ind2", "ind3", "ind4"]
				}`,
			statusCode:   422,
			errorMessage: "Required Password",
		},
	}

	for _, v := range samples {

		req, err := http.NewRequest("POST", "/users", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v", err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.CreateUser)
		handler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
		if err != nil {
			fmt.Printf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)
		if v.statusCode == 201 {
			assert.Equal(t, responseMap["name"], v.name)
			assert.Equal(t, responseMap["lastname"], v.lastname)
			assert.Equal(t, responseMap["email"], v.email)
			assert.Equal(t, responseMap["urlphoto"], v.urlphoto)
			assert.Equal(t, responseMap["telephone"], v.telephone)
			assert.Equal(t, responseMap["position"], v.position)
			assert.Equal(t, responseMap["urlphoto"], v.positionsought)
			assert.Equal(t, responseMap["telephone"], v.industry)
			assert.Equal(t, responseMap["position"], v.industrysought)
		}
		if v.statusCode == 422 || v.statusCode == 500 && v.errorMessage != "" {
			assert.Equal(t, responseMap["error"], v.errorMessage)
		}
	}
}

func TestGetUsers(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}
	_, err = seedUsers()
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetUsers)
	handler.ServeHTTP(rr, req)

	var users []models.User
	err = json.Unmarshal([]byte(rr.Body.String()), &users)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, len(users), 2)
}

func TestGetUserByID(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}
	user, err := seedOneUser()
	if err != nil {
		log.Fatal(err)
	}
	userSample := []struct {
		id           string
		statusCode   int
		name           string
		lastname       string
		email          string
		urlphoto       string
		telephone      string
		position       string
		positionsought pq.StringArray
		industry       string
		industrysought pq.StringArray
		errorMessage string
	}{
		{
			id:         strconv.Itoa(int(user.ID)),
			statusCode: 200,
			name:   user.Name,
			lastname: user.Lastname,
			email:      user.Email,
			urlphoto: user.Urlphoto,
			telephone: user.Telephone,
			position: user.Position,
			positionsought: user.Positionsought,
			industry: user.Industry,
			industrysought: user.Industrysought,
		},
		{
			id:         "unknwon",
			statusCode: 400,
		},
	}
	for _, v := range userSample {

		req, err := http.NewRequest("GET", "/users", nil)
		if err != nil {
			t.Errorf("This is the error: %v\n", err)
		}
		req = mux.SetURLVars(req, map[string]string{"id": v.id})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.GetUser)
		handler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
		if err != nil {
			log.Fatalf("Cannot convert to json: %v", err)
		}

		assert.Equal(t, rr.Code, v.statusCode)

		if v.statusCode == 200 {
			assert.Equal(t, user.Name, responseMap["name"])
			assert.Equal(t, user.Lastname, responseMap["lastname"])
			assert.Equal(t, user.Email, responseMap["email"])
			assert.Equal(t, user.Urlphoto, responseMap["urlphoto"])
			assert.Equal(t, user.Telephone, responseMap["telephone"])
			assert.Equal(t, user.Position, responseMap["position"])
			assert.Equal(t, user.Positionsought, responseMap["positionsought"])
			assert.Equal(t, user.Industry, responseMap["industry"])
			assert.Equal(t, user.Industrysought, responseMap["industrysought"])
		}
	}
}

func TestUpdateUser(t *testing.T) {

	var AuthEmail, AuthPassword string
	var AuthID uint32

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}
	users, err := seedUsers() //we need atleast two users to properly check the update
	if err != nil {
		log.Fatalf("Error seeding user: %v\n", err)
	}
	// Get only the first user
	for _, user := range users {
		if user.ID == 2 {
			continue
		}
		AuthID = user.ID
		AuthEmail = user.Email
		AuthPassword = "password" //Note the password in the database is already hashed, we want unhashed
	}
	//Login the user and get the authentication token
	token, err := server.SignIn(AuthEmail, AuthPassword)
	if err != nil {
		log.Fatalf("cannot login: %v\n", err)
	}
	tokenString := fmt.Sprintf("Bearer %v", token)

	samples := []struct {
		id             string
		updateJSON     string
		statusCode     int
		updateName     string
		updateLastname string
		updateEmail    string
		updateUrlphoto string
		updateTelephone string
		updatePosition string
		updatePositionsought pq.StringArray
		updateIndustry string
		updateIndustrysought pq.StringArray
		tokenGiven     string
		errorMessage   string
	}{
		{
			// Convert int32 to int first before converting to string
			id:             strconv.Itoa(int(AuthID)),
			updateJSON:     `{
				"name": "fred",
				"lastname": "arnoud",
				"email": "fred_arnoud@gmail.com",
				"urlphoto": "thisisanurl",
				"telephone": "0576756558",
				"password": "password",
				"position": "CTO",
				"positionsought": ["pos1", "pos2", "pos3", "pos4"],
				"industry": "tech",
				"industrysought": ["ind1", "ind2", "ind3", "ind4"]
				}`,
			statusCode:     200,
			updateName: "fred",
			updateLastname: "arnoud",
			updateEmail: "fred_arnoud@gmail.com",
			updateUrlphoto: "thisisanurl",
			updateTelephone: "0576756558",
			updatePosition: "CTO",
			updatePositionsought: []string{"pos1", "pos2", "pos3", "pos4"},
			updateIndustry: "tech",
			updateIndustrysought: []string{"pos1", "pos2", "pos3", "pos4"},
			tokenGiven:     tokenString,
			errorMessage:   "",
		},
		{
			// When password field is empty
			id:           strconv.Itoa(int(AuthID)),
			updateJSON:   `{
				"name": "toto",
				"lastname": "titi",
				"email": "toto_titi@gmail.com",
				"urlphoto": "thisisanurl",
				"telephone": "0546756558",
				"password": "",
				"position": "moka",
				"positionsought": ["pos1", "pos2", "pos3", "pos4"],
				"industry": "industry",
				"industrysought": ["ind1", "ind2", "ind3", "ind4"]
				}`,
			statusCode:   422,
			tokenGiven:   tokenString,
			errorMessage: "Required Password",
		},
		{
			// When no token was passed
			id:           strconv.Itoa(int(AuthID)),
			updateJSON:   `{
				"name": "lala",
				"lastname": "momo",
				"email": "lala_momo@gmail.com",
				"urlphoto": "thisisanurl",
				"telephone": "0516756558",
				"password": "password",
				"position": "marketeur",
				"positionsought": ["pos1", "pos2", "pos3", "pos4"],
				"industry": "marketing",
				"industrysought": ["ind1", "ind2", "ind3", "ind4"]
				}`,
			statusCode:   401,
			tokenGiven:   "",
			errorMessage: "Unauthorized",
		},
		{
			// When incorrect token was passed
			id:           strconv.Itoa(int(AuthID)),
			updateJSON:   `{
				"name": "romain",
				"lastname": "ledrogo",
				"email": "romain_ledrogo@gmail.com",
				"urlphoto": "thisisanurl",
				"telephone": "0616756558",
				"password": "password",
				"position": "directeur",
				"positionsought": ["pos1", "pos2", "pos3", "pos4"],
				"industry": "tech",
				"industrysought": ["ind1", "ind2", "ind3", "ind4"]
				}`,
			statusCode:   401,
			tokenGiven:   "This is incorrect token",
			errorMessage: "Unauthorized",
		},
		{
			id:           strconv.Itoa(int(AuthID)),
			updateJSON:   `{
				"name": "bab",
				"lastname": "jamb",
				"email": "babou_jam@gmail.com",
				"urlphoto": "thisisanurl",
				"telephone": "0576056558",
				"password": "password",
				"position": "archeologue",
				"positionsought": ["pos1", "pos2", "pos3", "pos4"],
				"industry": "archeologie",
				"industrysought": ["ind1", "ind2", "ind3", "ind4"]
				}`,
			statusCode:   500,
			tokenGiven:   tokenString,
			errorMessage: "Email Already Taken",
		},
		{
			// Remember "Kenny Morris" belongs to user 2
			id:           strconv.Itoa(int(AuthID)),
			updateJSON:   `{
				"name": "nil",
				"lastname": "armstrong",
				"email": "nil_armstrong@gmail.com",
				"urlphoto": "thisisanurl",
				"telephone": "0798456678",
				"password": "password",
				"position": "astronaute",
				"positionsought": ["pos1", "pos2", "pos3", "pos4"],
				"industry": "archeologie de l espace",
				"industrysought": ["ind1", "ind2", "ind3", "ind4"]
				}`,
			statusCode:   500,
			tokenGiven:   tokenString,
			errorMessage: "Telephone Already Taken",
		},
		{
			id:           strconv.Itoa(int(AuthID)),
			updateJSON:   `{
				"name": "agatha",
				"lastname": "christie",
				"email": "agatha_christiegmail.com",
				"urlphoto": "thisisanurl",
				"telephone": "0798156078",
				"password": "password",
				"position": "autrice",
				"positionsought": ["pos1", "pos2", "pos3", "pos4"],
				"industry": "litterature",
				"industrysought": ["ind1", "ind2", "ind3", "ind4"]
				}`,
			statusCode:   422,
			tokenGiven:   tokenString,
			errorMessage: "Invalid Email",
		},
		{
			id:           strconv.Itoa(int(AuthID)),
			updateJSON:   `{
				"name": "nil",
				"lastname": "armstrong",
				"email": "nil_armstrong@@gmail.com",
				"urlphoto": "thisisanurl",
				"telephone": "",
				"password": "password",
				"position": "astronaute",
				"positionsought": ["pos1", "pos2", "pos3", "pos4"],
				"industry": "archeologie de l espace",
				"industrysought": ["ind1", "ind2", "ind3", "ind4"]
				}`,
			statusCode:   422,
			tokenGiven:   tokenString,
			errorMessage: "Required Telephone",
		},
		{
			id:           strconv.Itoa(int(AuthID)),
			updateJSON:   `{
				"name": "nil",
				"lastname": "armstrong",
				"email": "",
				"urlphoto": "thisisanurl",
				"telephone": "0798456678",
				"password": "password",
				"position": "astronaute",
				"positionsought": ["pos1", "pos2", "pos3", "pos4"],
				"industry": "archeologie de l espace",
				"industrysought": ["ind1", "ind2", "ind3", "ind4"]
				}`,
			statusCode:   422,
			tokenGiven:   tokenString,
			errorMessage: "Required Email",
		},
		{
			id:         "unknwon",
			tokenGiven: tokenString,
			statusCode: 400,
		},
		{
			// When user 2 is using user 1 token
			id:           strconv.Itoa(int(2)),
			updateJSON:   `{
				"name": "nil",
				"lastname": "armstrong",
				"email": "nil_armstrong@gmail.com",
				"urlphoto": "thisisanurl",
				"telephone": "0798456678",
				"password": "password",
				"position": "astronaute",
				"positionsought": ["pos1", "pos2", "pos3", "pos4"],
				"industry": "archeologie de l espace",
				"industrysought": ["ind1", "ind2", "ind3", "ind4"]
				}`,
			tokenGiven:   tokenString,
			statusCode:   401,
			errorMessage: "Unauthorized",
		},
	}

	for _, v := range samples {

		req, err := http.NewRequest("POST", "/users", bytes.NewBufferString(v.updateJSON))
		if err != nil {
			t.Errorf("This is the error: %v\n", err)
		}
		req = mux.SetURLVars(req, map[string]string{"id": v.id})

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.UpdateUser)

		req.Header.Set("Authorization", v.tokenGiven)

		handler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
		if err != nil {
			t.Errorf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)
		if v.statusCode == 200 {
			assert.Equal(t, responseMap["name"], v.updateName)
			assert.Equal(t, responseMap["lastname"], v.updateLastname)
			assert.Equal(t, responseMap["email"], v.updateEmail)
			assert.Equal(t, responseMap["urlphoto"], v.updateUrlphoto)
			assert.Equal(t, responseMap["telephone"], v.updateTelephone)
			assert.Equal(t, responseMap["position"], v.updatePosition)
			assert.Equal(t, responseMap["positionsought"], v.updatePositionsought)
			assert.Equal(t, responseMap["industry"], v.updateIndustry)
			assert.Equal(t, responseMap["industrysought"], v.updateIndustrysought)
		}
		if v.statusCode == 401 || v.statusCode == 422 || v.statusCode == 500 && v.errorMessage != "" {
			assert.Equal(t, responseMap["error"], v.errorMessage)
		}
	}
}

func TestDeleteUser(t *testing.T) {

	var AuthEmail, AuthPassword string
	var AuthID uint32

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	users, err := seedUsers() //we need atleast two users to properly check the update
	if err != nil {
		log.Fatalf("Error seeding user: %v\n", err)
	}
	// Get only the first and log him in
	for _, user := range users {
		if user.ID == 2 {
			continue
		}
		AuthID = user.ID
		AuthEmail = user.Email
		AuthPassword = "password" ////Note the password in the database is already hashed, we want unhashed
	}

	//Login the user and get the authentication token
	token, err := server.SignIn(AuthEmail, AuthPassword)
	if err != nil {
		log.Fatalf("cannot login: %v\n", err)
	}
	tokenString := fmt.Sprintf("Bearer %v", token)

	userSample := []struct {
		id           string
		tokenGiven   string
		statusCode   int
		errorMessage string
	}{
		{
			// Convert int32 to int first before converting to string
			id:           strconv.Itoa(int(AuthID)),
			tokenGiven:   tokenString,
			statusCode:   204,
			errorMessage: "",
		},
		{
			// When no token is given
			id:           strconv.Itoa(int(AuthID)),
			tokenGiven:   "",
			statusCode:   401,
			errorMessage: "Unauthorized",
		},
		{
			// When incorrect token is given
			id:           strconv.Itoa(int(AuthID)),
			tokenGiven:   "This is an incorrect token",
			statusCode:   401,
			errorMessage: "Unauthorized",
		},
		{
			id:         "unknwon",
			tokenGiven: tokenString,
			statusCode: 400,
		},
		{
			// User 2 trying to use User 1 token
			id:           strconv.Itoa(int(2)),
			tokenGiven:   tokenString,
			statusCode:   401,
			errorMessage: "Unauthorized",
		},
	}

	for _, v := range userSample {

		req, err := http.NewRequest("GET", "/users", nil)
		if err != nil {
			t.Errorf("This is the error: %v\n", err)
		}
		req = mux.SetURLVars(req, map[string]string{"id": v.id})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.DeleteUser)

		req.Header.Set("Authorization", v.tokenGiven)

		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, v.statusCode)

		if v.statusCode == 401 && v.errorMessage != "" {
			responseMap := make(map[string]interface{})
			err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
			if err != nil {
				t.Errorf("Cannot convert to json: %v", err)
			}
			assert.Equal(t, responseMap["error"], v.errorMessage)
		}
	}
}
