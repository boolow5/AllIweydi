package controllers

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/boolow5/AllIweydi/g"

	resty "gopkg.in/resty.v0"
)

var BOL_AUTH_URL string

func init() {
	BOL_AUTH_URL = os.Getenv("BOL_AUTH_URL")
	if BOL_AUTH_URL == "" {
		BOL_AUTH_URL = "http://localhost:8080"
	}
}

type AuthController struct {
	BaseController
}

func (c *AuthController) Get() {
	SetTemplate("pages/login.tpl", &c.Controller)
}

func (this *AuthController) Post() {
	// set the flash message system
	// flash := beego.NewFlash()
	// get username and password from the post data
	responseMessage := map[string]interface{}{}

	// send request to BolAuth to check if it returns a token
	requestBody := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	err := json.NewDecoder(this.Ctx.Request.Body).Decode(&requestBody)
	if err != nil {
		responseMessage["error"] = "iweydi-response-parsing-error"
		responseMessage["explation"] = err.Error()
		this.Data["json"] = responseMessage
		this.ServeJSON()
		return
	}

	username := requestBody.Username

	fmt.Println("Request Body:", requestBody)
	res, err := resty.R().
		SetQueryParams(map[string]string{
			"app_id": g.APP_ID,
		}).
		SetBody(requestBody).
		Post(BOL_AUTH_URL + "/login")

	if err != nil {
		responseMessage["error"] = "iweydi-auth-connection-error"
		responseMessage["explation"] = err.Error()
		this.Data["json"] = responseMessage
		this.ServeJSON()
		return
	}

	response := map[string]interface{}{}
	err = json.Unmarshal(res.Body(), &response)
	if err != nil {
		responseMessage["error"] = "iweydi-response-parsing-error"
		responseMessage["explation"] = err.Error()
		this.Data["json"] = responseMessage
		this.ServeJSON()
		return
	}

	if res.StatusCode() != 200 {
		responseMessage["error"] = "incorrect-username-or-password"
		responseMessage["explation"] = err.Error()
		this.Data["json"] = responseMessage
		this.ServeJSON()
		return
	}
	// if yes save the token and username to session
	if response["token"] == nil {
		responseMessage["error"] = "incorrect-username-or-password"
		responseMessage["explation"] = err.Error()
		this.Data["json"] = responseMessage
		this.ServeJSON()
		return
	}

	this.SetSession("token", response["token"])
	this.SetSession("username", username)

	responseMessage["success"] = "login-success"
	this.Data["json"] = responseMessage
	this.ServeJSON()
}

func (this *AuthController) Logout() {
	this.DestroySession()
	this.Data["json"] = map[string]string{"success": "logout-success"}
	this.ServeJSON()
}

func (this *AuthController) GetRegister() {
	SetTemplate("pages/register.tpl", &this.Controller)
}

func (this *AuthController) PostRegister() {
	this.Data["json"] = map[string]interface{}{
		"username":  this.GetString("username"),
		"password":  this.GetString("password"),
		"full_name": this.GetString("full_name"),
	}
	this.ServeJSON()
}
