package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"strconv"
	"time"
)

type DataWrapper struct {
	Key  string `json:"key"`
	Html string
}

func GetUnAnsweredQuestion(page string, tag string, conn *connection) /*(output *Questions, error error)*/ {
	params := make(map[string]string)
	params["page"] = page
	params["tagged"] = tag
	params["order"] = "desc"
	params["sort"] = "creation"
	params["filter"] = "!9YdnSQVoS" //total question

	mysession := NewSession("stackoverflow")
	questions, err := mysession.UnansweredQuestions(params)
	if err != nil {
		return
	}

	for _, question := range questions.Items {

		pushData := &channelAndData{conn: conn}
		pageData := &UnAnsQuestionData{question.Owner.User_id,
			template.HTML(GetDateTimeFormUnixTimeStamp(question.Creation_date)),
			template.HTML(question.Owner.Display_name),
			template.HTML(question.Title),
			template.HTML(question.Link)}

		tmpl := template.New("questionresult")
		if tmpl, err = tmpl.ParseFiles("templates/unanswered.html"); err != nil {
			fmt.Println(err)
		}

		var buffer bytes.Buffer
		tmpl.ExecuteTemplate(&buffer, "content", pageData)

		sdata := buffer.String()
		dw := &DataWrapper{"LOADUNANSWEREDQUESTIONS", sdata}

		b, err := json.Marshal(dw)
		if err != nil {
			fmt.Println("error:", err)
		}

		pushData.data = b
		h.senddata <- pushData

	}
	//fmt.Println(questions.Has_more)

	if !questions.Has_more {
		dw := &DataWrapper{"LOADUNANSWEREDQUESTIONS", "<NOMOREDATA>"}
		b, err := json.Marshal(dw)
		if err != nil {
			fmt.Println("error:", err)
		}
		pushData := &channelAndData{conn: conn}
		pushData.data = b
		h.senddata <- pushData
	}

	dw := &DataWrapper{"LOADUNANSWEREDQUESTIONS", "<EOM>"}
	b, err := json.Marshal(dw)
	if err != nil {
		fmt.Println("error:", err)
	}
	pushData := &channelAndData{conn: conn}
	pushData.data = b
	h.senddata <- pushData
}

func GetRelatedTags(page string, tag string, conn *connection) /*(output *Questions, error error)*/ {
	params := make(map[string]string)
	params["sort"] = "votes"
	params["order"] = "desc"
	pageNo := 1
	params["page"] = strconv.Itoa(pageNo)
	tags := new(Tags)
	data := new(ChartData)
	jsonData := new(NCD)

	mysession := NewSession("stackoverflow")
	for params["page"] == "1" || tags.Has_more {
		_tags, err := mysession.RelatedTags([]string{tag}, params)
		if err != nil {
			fmt.Println("ERRROR in GetRelatedTags")
		}
		tags.Items = append(_tags.Items, tags.Items...)
		tags.Has_more = _tags.Has_more

		pageNo++
		params["page"] = strconv.Itoa(pageNo)
	}

	//Required for DiscreteBar Chart
	for _, item := range tags.Items {
		data.label = append(data.label, item.Name)
		data.value = append(data.value, item.Count)

		jsonData.Values = append(jsonData.Values, &NewChartData{item.Name, item.Count})
	}
	jsonData.Key = "PIECHARTDATA"

	b, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println("error:", err)
	}

	pushData := &channelAndData{conn: conn, data: b}
	h.senddata <- pushData

	//Default Page
	GetTotalOfQuestion(page, tag, conn)

	/* METHOD-2
	pieChartData := make([]*NewChartData, 0)
	for _, item := range tags.Items {
		//pieChartData.label = append(pieChartData.label, item.Name)
		//pieChartData.value = append(pieChartData.value, item.Count)

		pieChartData = append(pieChartData, &NewChartData{item.Name, item.Count})
	}

	fmt.Println("Out")
	b, err := json.Marshal(pieChartData)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(b))

	conn.write(websocket.TextMessage, []byte("PIECHARTDATASTART"))
	conn.ws.WriteJSON(pieChartData)
	conn.write(websocket.TextMessage, []byte("PIECHARTDATAENDS"))
	*/

	//return tags
}

func GetTotalOfQuestion(page string, tag string, conn *connection) /*(output *Questions, error error)*/ {
	params := make(map[string]string)
	params["page"] = "1"
	params["tagged"] = tag
	params["filter"] = "!9YdnSQVoS" //total question
	jsonData := new(QNCD)
	jsonData.Key = "INDIVIDUALTAGDATASTART"
	jsonData.Values.Key = "StatisticalData"

	mysession := NewSession("stackoverflow")
	questions, err := mysession.UnansweredQuestions(params)
	if err != nil {
		fmt.Printf("ERROR : %s \n", err)
		return
	}

	jsonleaf1 := NewChartData{"Unanswered", questions.Total}
	jsonData.Values.Values = append(jsonData.Values.Values, &jsonleaf1)
	allquestions, err := mysession.AllQuestions(params)
	if err != nil {
		return
	}
	jsonleaf := NewChartData{"TotalQuestion", allquestions.Total}
	jsonData.Values.Values = append(jsonData.Values.Values, &jsonleaf)

	b, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println("error:", err)
	}

	pushData := &channelAndData{conn: conn, data: b}
	h.senddata <- pushData
}

func GetUserDetails(userId int) *Users {

	params := make(map[string]string)
	params["ids"] = strconv.Itoa(userId)
	params["filter"] = "!9YdnSA07B"
	/*
		params["include"] = ".about_me"
		filter=!9YdnSA07B
	*/

	mysession := NewSession("stackoverflow")
	users, err := mysession.UserDetail(params)
	if err != nil {
		fmt.Println("ERRROR in USER DETAILS")
		return nil
	}
	return users
}

func DemmoRoutine(tag string, conn *connection) {
	ticker := time.NewTicker(1000 * time.Millisecond)
	for {
		select {
		case t := <-ticker.C:
			cd := &channelAndData{conn, []byte(tag + " " + t.String())}
			if _, ok := h.connections[cd.conn]; ok {
				h.senddata <- cd
			} else {
				fmt.Println("Not Found")
				return
			}
		}
	}

}
