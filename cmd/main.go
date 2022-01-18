package main

import (
	"fmt"
	"zabbix"

	"github.com/kav0484/go-zabbix-inventory/internal/zabbix"
)

func main() {
	server := "http://zabbix.ruadmin24.ru/api_jsonrpc.php"
	user := "Admin"
	password := "zabbix"

	token := zabbix.Login(server, user, password)

	fmt.Println(token)

	// LOGOUT
	/*
		//req := fmt.Sprintf(`{"jsonrpc": "2.0", "method": "host.get", "params": {"output": ["hostid", "host"], "selectInterfaces": ["interfaceid", "ip"]}, "id": 2, "auth": "%s"}`, auth.Token)
		req := fmt.Sprintf(`{"jsonrpc": "2.0", "method": "host.get", "params": {"output": ["hostid", "host","available","proxy_address"], "selectInterfaces": ["interfaceid","ip"], "selectGroups": ["groupid","name"]}, "id": 1, "auth": "%s"}`, auth.Token)

		inv, _ := http.NewRequest("GET", server, bytes.NewBuffer([]byte(req)))
		inv.Header.Set("Content-Type", "application/json")
		res1, _ := client.Do(inv)
		defer inv.Body.Close()

		body1, _ := ioutil.ReadAll(res1.Body)
		//fmt.Println(string(body1))

		json.Unmarshal(body1, &host)

		fmt.Println(host)
	*/
}
