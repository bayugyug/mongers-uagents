package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ua-parser/uap-go/uaparser"
)

func doIt() {

	if pUAShow {
		showUAMainList()
		return
	}
	if pUAList != "" {
		showUAList()
		return
	}

	if pUAGet != "" {
		showUADetails()
		return
	}
}

func showUAList() {
	ua := UserAgent{Agent: pUAList}
	if _, ok := UAHashMap["UA"+pUAList]; ok {
		ua.Data = UAHashMap["UA"+pUAList]
	}
	ua.Total = len(ua.Data)
	jsonStr, _ := json.MarshalIndent(ua, "", "\t")
	fmt.Println(string(jsonStr))
}

func showUAMainList() {
	ua := UserAgent{Agent: "UserAgents"}
	for k, _ := range UAHashMap {
		ua.Data = append(ua.Data, strings.TrimPrefix(k, "UA"))
	}
	ua.Total = len(ua.Data)
	jsonStr, _ := json.MarshalIndent(ua, "", "\t")
	fmt.Println(string(jsonStr))

}

func showUADetails() {

	parser := uaparser.NewFromSaved()
	client := parser.Parse(pUAGet)
	jsonStr, _ := json.MarshalIndent(client, "", "\t")
	fmt.Println(string(jsonStr))

}
