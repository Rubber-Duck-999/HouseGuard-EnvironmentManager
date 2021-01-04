package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	/*"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"*/
)

func ApiStatus(username string, password string, id string, endpoint string) {

	/*
	ses, err := session.NewSession(&aws.Config{Region: aws.String("eu-west-2")})
	if err != nil {
		fmt.Println("Error in session creation")
	}

	client_id := aws.String(id)

	_username := aws.String(username)
	_password := aws.String(password)

	params := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		AuthParameters: map[string]*string{
			"USERNAME": _username,
			"PASSWORD": _password,
		},
		ClientId: client_id,
	}
	cip := cognitoidentityprovider.New(ses)

	authResponse, authError := cip.InitiateAuth(params)
	if authError != nil {

		fmt.Println("Error = ", authError)
		//return nil, authError
	}

	//fmt.Println(*authResponse.AuthenticationResult.AccessToken)
	*/

	client := http.Client{}

	q := url.Values{}
	q.Add("user", "User")
	q.Add("state", "OFF")
	req, _ := http.NewRequest("POST", endpoint+"/alarmEvent", strings.NewReader(q.Encode()))

	resp, err := client.Do(req)
	if err != nil {
	   fmt.Printf("%s", err)
	}
	fmt.Println(resp.Header)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	   fmt.Printf("%s", err)
	}
	fmt.Println(string(body))
}
