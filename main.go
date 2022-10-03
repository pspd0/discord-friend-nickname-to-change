package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var id, nickname string
	fmt.Print("Enter Your Friend ID: ")
	fmt.Scan(&id)

	fmt.Print("Enter name to Change: ")
	fmt.Scan(&nickname)

	payload, err := json.Marshal(map[string]interface{}{
		"nickname": nickname,
	})
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPatch, "https://discord.com/api/v9/users/@me/relationships/"+id, bytes.NewBuffer(payload))
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9006 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36")
	req.Header.Set("content-type", "application/json")

	req.Header.Set("authorization", "★★★ [Your Token Here] ★★★")

	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))
}
