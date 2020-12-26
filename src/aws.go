package main

import (
	/*"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"*/
)

func ApiStatus(username string, password string, id string, url string) {

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

	client := new(http.Client)
	req, _ := http.NewRequest("GET", url, nil)

	q := req.URL.Query()
	q.Add("status_request", "")
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Authorization", *authResponse.AuthenticationResult.AccessToken)

	resp, _ := client.Do(req)

	b, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("%s\n", b)*/
}
