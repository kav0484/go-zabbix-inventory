package zabbix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Auth struct {
	Jsonrpc string `json:"jsonrpc"`
	Token   string `json:"result"`
	Id      int    `json:"id"`
}

type JsonResponse struct {
	//Jsonrpc string `json:"jsonrpc"`
	Result []struct {
		Host         string `json:"host"`
		Available    string `json:"available"`
		ProxyAddress string `json:"proxy_address"`
		Groups       []struct {
			Name string `json:"name"`
		} `json:"groups"`
		Interfaces []struct {
			IP string `json:"ip"`
		} `json:"interfaces"`
	} `json:"result"`
}

func Login(server, user, password string) string {

	var auth Auth

	client := &http.Client{}

	authJson := fmt.Sprintf(`{"jsonrpc": "2.0", "method": "user.login", "params": {"user": "%s", "password": "%s"}, "id": 1, "auth": null}`, user, password)
	authenticate, err := http.NewRequest("GET", server, bytes.NewBuffer([]byte(authJson)))
	authenticate.Header.Set("Content-Type", "application/json")

	res, err := client.Do(authenticate)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(body, &auth)

	return string(auth.Token)

}

func Logout(a *Auth) (server, token string) {
	client := &http.Client{}
	logoutJson := fmt.Sprintf(`{"jsonrpc": "2.0", "method": "user.logout", "params": [], "id": 1, "auth": "%s" }`, a.Token)
	logout, _ := http.NewRequest("GET", server, bytes.NewBuffer([]byte(logoutJson)))
	logout.Header.Set("Content-Type", "application/json")
	client.Do(logout)
	return
}
