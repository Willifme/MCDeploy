package mcdeploy

import (
	"encoding/json"
	"fmt"
	"os"
)

// Base struct which bukget uses for it's JSON

type Plugin []struct {
	Description string `json:"description"`

	Plugin_Name string `json:"plugin_name"`

	Slug string `json:"slug"`
}

// Return three strings representing the plugin name, description and slug

func ParseJSON(body []uint8, data Plugin) ([]string, []string, []string) {

	json.Unmarshal(body, &data)

	var name, description, slug []string

	// If the plugin search returns nothing "[]", we inform the user

	if len(data) <= 0 {

		fmt.Println("No plugins were found")

		os.Exit(2)

	}

	for _, plugin := range data {

		if len(plugin.Description) <= 0 { // No description found

			plugin.Description = "No description available"

		}

		// Push the plugin info into their own splice

		name = append(name, plugin.Plugin_Name)

		description = append(description, plugin.Description)

		slug = append(slug, plugin.Slug)

	}

	return name, description, slug

}
