package singularity

import (
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

const (
	//SlackAPI is the slack API url. Hit it with SlackAPI + apiCall
	SlackAPI = "slack.com/api/"
)

func (singularity *Singularity) ConnectToTeam(token string) (*SlackInstance, error) {
	helper := HTTPHelper{Client: &http.Client{}, Transport: "https://"}
	var response RTMResp
	_, err := helper.post("rtm.start", &response, "token", token)
	if err != nil {
		return nil, err
	}

	//Connect to websocket.
	conn, err := singularity.startWebsocketStream(response.URL)
	if err != nil {
		return nil, err
	}
	fmt.Println("29")
	return singularity.addTeam(conn, response)
}

func (singularity *Singularity) startWebsocketStream(url string) (*websocket.Conn, error) {
	conn, err := websocket.Dial(url, "", "http://api.slack.com") //TODO Don't hard code here.
	if err != nil {
		return nil, err
	}
	var hello struct {
		Hello string `json:"type"`
	}
	err = websocket.JSON.Receive(conn, &hello)
	if err != nil {
		return nil, err
	}
	if hello.Hello != "hello" {
		return nil, errors.New("Didn't recieve my hello message :C")
	}
	return conn, nil
}
