// This file is part of ghget, a tool for downloading repos in bulk from
// GitHub.
// Copyright (C) 2020 Jordan Ocokoljic.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package ghget

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
)

// Repo represents a GitHub repository owned by a user.
type Repo struct {
	Name string `json:"name"`
	URL  string `json:"html_url"`
}

// getApiUrl will return the url required to access the repo information of a
// specific user.
func getAPIURL(username string) string {
	return fmt.Sprintf("https://api.github.com/users/%s/repos", username)
}

// ListRepos will collect all the repos that a user has
func ListRepos(username string) ([]Repo, error) {
	url := getAPIURL(username)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	repos := []Repo{}
	json.Unmarshal(body, &repos)
	return repos, nil
}

// GetRepo prepares and executes the command to fetch a specific repo.
func GetRepo(url string) error {
	cmd := exec.Command("git", "clone", url)
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
