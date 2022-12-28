package schema

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func getNameFiles(directory string) []string {
	var file_names []string

	files, err := os.ReadDir("./database/schema/")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !f.IsDir() && strings.Contains(f.Name(), ".sql") {
			file_names = append(file_names, f.Name())
		}
	}

	return file_names
}

func dropDatabase(database string, user string) {
	var out bytes.Buffer

	command := exec.Command("psql", "-U", user, "-d", "postgres", "-c", fmt.Sprintf("DROP DATABASE IF EXISTS %s;", database))
	command.Stdout = &out

	err := command.Run()
	if err != nil {
		log.Println(err)
	}

	log.Print(out.String())
	out.Reset()
}

func createDatabase(database string, user string) {
	var out bytes.Buffer

	command := exec.Command("psql", "-U", user, "-d", "postgres", "-c", fmt.Sprintf("CREATE DATABASE %s;", database))
	command.Stdout = &out

	err := command.Run()
	if err != nil {
		log.Println(err)
	}
	log.Print(out.String())
	out.Reset()
}

func Generate(database string, user string) {
	log.Println("Generating database " + database + "... Postgres must be installed and running!")

	dropDatabase(database, user)

	createDatabase(database, user)

	files := getNameFiles("./database/schema/")

	for _, file := range files {
		var out bytes.Buffer

		command := exec.Command("psql", "-U", user, "-d", database, "-f", "./database/schema/"+file)
		command.Stdout = &out

		err := command.Run()
		if err != nil {
			log.Println(err)
		}
		log.Print("\n" + out.String())
		out.Reset()
	}

}
