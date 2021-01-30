package main

import (
	"encoding/json"
	"flag"
	"ipcom1/parte2/model"
	"log"
	"os"
)

var (
	// output = make(map[string]model.Output)
	output         = make(map[string]map[string]map[string]bool)
	inputFileName  = flag.String("in", "example.in", "path of input file")
	outputFileName = flag.String("out", "", "set value if you want the output in a file")
	outWriter      = os.Stdout
	err            error
)

func init() {
	flag.Parse()
}

func main() {
	_, data := ReadData(*inputFileName)
	for _, d := range data {
		if _, ok := output[d[0]]; !ok {
			// output[d[0]] = model.NewOutput(d[0])
			output[d[0]] = make(map[string]map[string]bool)
		}
		if _, ok := output[d[0]][d[1]]; !ok {
			output[d[0]][d[1]] = make(map[string]bool, 0)
		}
		output[d[0]][d[1]][d[2]] = true
	}
	organizations := make([]model.Organization, 0)
	for organizationName, users := range output {
		organization := model.NewOrganization(organizationName)
		organization.Users = make([]model.User, 0)
		for userName, roles := range users {
			user := model.NewUser(userName)
			user.Roles = make([]string, 0)
			for role := range roles {
				user.Roles = append(user.Roles, role)
			}
			organization.Users = append(organization.Users, user)
		}
		organizations = append(organizations, organization)
	}
	if *outputFileName != "" {
		outWriter, err = os.Create(*outputFileName)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	err = json.NewEncoder(outWriter).Encode(organizations)
	if err != nil {
		log.Fatal(err.Error())
	}

}
