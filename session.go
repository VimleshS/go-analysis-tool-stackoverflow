package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var host string = "http://api.stackexchange.com"
var transport http.RoundTripper

type Session struct {
	Site string
}

func (session Session) AllQuestions(params map[string]string) (output *Questions, error error) {
	output = new(Questions)
	error = session.get("questions", params, output)
	return
}

func (session Session) RelatedTags(tags []string, params map[string]string) (output *Tags, error error) {
	request_path := strings.Join([]string{"tags", strings.Join(tags, ";"), "related"}, "/")

	output = new(Tags)
	error = session.get(request_path, params, output)
	return
}

func (session Session) UnansweredQuestions(params map[string]string) (output *Questions, error error) {
	output = new(Questions)
	error = session.get("questions/unanswered", params, output)
	return
}

func (session Session) UserDetail(params map[string]string) (output *Users, error error) {
	output = new(Users)
	userurl := fmt.Sprintf("users/%s", params["ids"])

	error = session.get(userurl, params, output)
	return
}

func NewSession(site string) *Session {
	return &Session{Site: site}
}

func (session Session) get(section string, params map[string]string, collection interface{}) (error error) {
	//set parameters for querystring
	params["site"] = session.Site
	return get(section, params, collection)
}

func parseResponse(response *http.Response, result interface{}) (error error) {
	defer response.Body.Close()
	bytes, error := ioutil.ReadAll(response.Body)
	if error != nil {
		return
	}
	error = json.Unmarshal(bytes, result)
	if error != nil {
		print(error.Error())
	}
	if response.StatusCode == 400 {
		error = errors.New(fmt.Sprintf("Bad Request: %s", string(bytes)))
	}
	return
}

func get(section string, params map[string]string, collection interface{}) (error error) {
	client := &http.Client{Transport: getTransport()}

	response, error := client.Get(setupEndpoint(section, params).String())

	if error != nil {
		return
	}
	//fmt.Println(response)
	error = parseResponse(response, collection)

	return

}

func getTransport() http.RoundTripper {
	if transport != nil {
		return transport
	}
	return http.DefaultTransport
}

func SetTransport(t http.RoundTripper) {
	transport = t
}

// construct the endpoint URL
func setupEndpoint(path string, params map[string]string) *url.URL {
	base_url, _ := url.Parse(host)
	base_url.Scheme = "https"
	endpoint, _ := base_url.Parse("/2.2/" + path)

	query := endpoint.Query()
	params["access_token"] = StackExchangeAccessToken.access_token
	params["key"] = Config[API_KEY_ID]
	for key, value := range params {
		query.Set(key, value)
	}

	endpoint.RawQuery = query.Encode()
	//fmt.Println(endpoint)
	return endpoint
}
