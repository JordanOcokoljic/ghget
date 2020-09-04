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

package main

import (
	"fmt"
	"ghget"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: ghget <username>")
	}

	username := os.Args[1]
	repos, err := ghget.ListRepos(username)
	if err != nil {
		fmt.Println("An error occured:\n%s\n")
	}

	err = ghget.GetRepos(repos)
	if err != nil {
		fmt.Println("An error occured:\n%s\n")
	}
}
