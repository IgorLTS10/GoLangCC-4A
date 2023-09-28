package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"time"
)

type Repository struct {
	Name         string    `json:"name"`
	LastModified time.Time `json:"pushed_at"`
}

func main() {
	username := "IgorLTS10"
	token := "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"

	err := getAndPrintRecentRepositories(username, token)
	if err != nil {
		fmt.Println("Erreur:", err)
	}
	/* Clone all repositories
	repos, err := getRepositories(username, token)
	if err != nil {
		fmt.Println("Erreur lors de la récupération des référentiels:", err)
		return
	}

	for _, repo := range repos {
		cloneURL := fmt.Sprintf("https://github.com/%s/%s.git", username, repo.Name)
		err := exec.Command("git", "clone", cloneURL).Run()
		if err != nil {
			fmt.Println("Erreur lors du clonage du référentiel:", err)
		}
	}*/
}

func getAndPrintRecentRepositories(username, token string) error {
	repos, err := getRepositories(username, token)
	if err != nil {
		return err
	}

	sort.Slice(repos, func(i, j int) bool {
		return repos[i].LastModified.After(repos[j].LastModified)
	})

	if len(repos) > 100 {
		repos = repos[:100]
	}

	err = createCSV(username, repos)
	if err != nil {
		return err
	}

	for i, repo := range repos {
		fmt.Printf("%d. Nom du référentiel: %s\n", i+1, repo.Name)
		fmt.Printf("   Date de dernière modification: %s\n", repo.LastModified)
	}

	return nil
}

func getRepositories(username, token string) ([]Repository, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", username)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "token "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var repos []Repository
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&repos)
	if err != nil {
		return nil, err
	}

	return repos, nil
}

func createCSV(username string, repos []Repository) error {
	currentDate := time.Now().Format("2006-01-02")

	fileName := fmt.Sprintf("csv/%s_%s.csv", username, currentDate)
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)

	headers := []string{"Username", "Date de récupération"}
	csvWriter.Write(headers)

	data := []string{username, currentDate}
	csvWriter.Write(data)

	for _, repo := range repos {
		data = []string{repo.Name, repo.LastModified.String()}
		csvWriter.Write(data)
	}

	csvWriter.Flush()

	return csvWriter.Error()
}
