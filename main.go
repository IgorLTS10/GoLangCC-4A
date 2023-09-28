package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"sort"
)

type Repository struct {
	Name         string `json:"name"`
	LastModified string `json:"updated_at"`
}

func main() {
	username := "IgorLTS10"
	token := "sss"

	repos, err := getRepositories(username, token)
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}

	sort.Slice(repos, func(i, j int) bool {
		return repos[i].LastModified > repos[j].LastModified
	})

	for _, repo := range repos {
		cloneURL := fmt.Sprintf("https://github.com/%s/%s.git", username, repo.Name)
		err := exec.Command("git", "clone", cloneURL).Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func getRepositories(username string, token string) ([]Repository, error)) {
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
