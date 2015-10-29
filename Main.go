package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strconv"
)

var page Page

type Page struct {
	EndPoint    string
	WebSocketId string
	UserDetails string
	Labels      []string
}

func main() {
	go h.run()
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ws", serveWs)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/authenticated", authenticatedHandler)
	http.HandleFunc("/userdetail", getajaxHandler)

	http.ListenAndServe(GetPort(), nil)
}

func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4748"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	log.Println("\nServer Started On Port..")
	return ":" + port
}

func landingPage(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code != "" {
		authenticatedHandler(w, r)
	} else {
		auth_url := AuthURL(Config[APP_ID], Config[REDIRECT_URI], map[string]string{"scope": "read_inbox"})
		http.Redirect(w, r, auth_url, http.StatusFound)
	}
}

func getajaxHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}
	userId, err := strconv.Atoi(string(data))
	if err != nil {
		userId = 1
	}
	users := GetUserDetails(userId)

	ud := &UserData{template.HTML(users.Items[0].Display_name),
		template.HTML(users.Items[0].Website_url),
		template.HTML(users.Items[0].About_me),
		template.HTML(users.Items[0].Location),
		template.HTML(users.Items[0].Link),
		template.HTML(users.Items[0].Profile_image)}

	in := "templates/userdetail.html"
	tmpl, err := template.ParseFiles(in)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "content", ud /*users.Items[0]*/); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	auth_url := AuthURL(Config[APP_ID], Config[REDIRECT_URI], map[string]string{"scope": "read_inbox"})
	http.Redirect(w, r, auth_url, http.StatusFound)
}

func authenticatedHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	if StackExchangeAccessToken.access_token == "" {
		access_token, err := ObtainAccessToken(Config[APP_ID], Config[API_KEY_SECRET],
			code, Config[API_KEY_ID], Config[REDIRECT_URI])
		StackExchangeAccessToken = AccessToken{access_token["access_token"], access_token["expires"]}

		//fmt.Println(StackExchangeAccessToken)
		if err != nil {
			fmt.Fprintf(w, "%v", err.Error())
			http.Error(w, http.StatusText(500), 500)
			return
		}
	}

	lp := path.Join("templates", "layout.html")
	var fp string
	if r.URL.Path == "/" {
		fp = path.Join("templates", "/input.html")
	} else {
		fp = path.Join("templates", "/input.html")
	}
	p := Page{}
	p.EndPoint = r.Host
	RenderTemplates(w, lp, fp, p)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	lp := path.Join("templates", "layout.html")
	var fp string
	if r.URL.Path == "/" {
		fp = path.Join("templates", "/home.html")
	} else {
		fp = path.Join("templates", r.URL.Path+".html")
	}
	p := Page{}
	RenderTemplates(w, lp, fp, p)
}

func RenderTemplates(w http.ResponseWriter, t1 string, t2 string, page Page) {
	tmpl, err := template.ParseFiles(t1, t2)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}
	if err := tmpl.ExecuteTemplate(w, "layout", page); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}
