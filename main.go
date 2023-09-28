package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Repository struct {
	Name         string `json:"name"`
	LastModified string `json:"updated_at"`
}

func main() {
	username := "IgorLTS10"
	token := ""

}

func getRepositories(username string, token string) {
	url := fmt.Sprintln("https://api.github.com/users/%s/repos", username)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Connexion", "token "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var repos []Repository
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&repos)
	if err != nil {
		fmt.Println(err)
	}
}
