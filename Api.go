package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type session struct {
}

func AuthURL(client_id, redirect_uri string, options map[string]string) (output string) {
	auth_url, _ := url.Parse(StackExchangeOauthUri)
	auth_query := auth_url.Query()
	auth_query.Add("client_id", client_id)
	auth_query.Add("redirect_uri", redirect_uri)

	for key, value := range options {
		auth_query.Add(key, value)
	}

	auth_url.RawQuery = auth_query.Encode()
	return auth_url.String()
}

func ObtainAccessToken(client_id, client_secret, code, key, redirect_uri string) (
	output map[string]string, error error) {

	client := &http.Client{Transport: http.DefaultTransport}

	auth_url := StackExchangeAccessTokenUri

	parsed_auth_url, _ := url.Parse(auth_url)
	auth_query := parsed_auth_url.Query()
	auth_query.Add("client_id", client_id)
	auth_query.Add("client_secret", client_secret)
	auth_query.Add("code", code)
	auth_query.Add("key", key)
	auth_query.Add("redirect_uri", redirect_uri)

	// make the request
	response, error := client.PostForm(auth_url, auth_query)
	if error != nil {
		return
	}

	//check whether the response is a bad request
	if response.StatusCode != 200 {
		collection := new(authError)
		error = parseResponse(response, collection)
		error = errors.New(collection.Error["type"] + ": " + collection.Error["message"])
	} else {
		// if not process the output
		bytes, err2 := ioutil.ReadAll(response.Body)

		if err2 != nil {
			return output, err2
		}

		collection, err3 := url.ParseQuery(string(bytes))
		output = map[string]string{"access_token": collection.Get("access_token"),
			"expires": collection.Get("expires")}
		error = err3
	}

	return
}
