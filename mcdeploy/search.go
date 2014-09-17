package mcdeploy

import (
	"fmt"
)

// The "i" variable is for the function to access the iterator
func IterateOverPluginInfo(iterator []string, output func(i int)) {

	for i := range iterator {

		output(i)

	}

}

func SearchBySlug(size int, plugin_name string) []string {

	url := fmt.Sprintf("http://api.bukget.org/3/search/slug/like/%s?size=%d", plugin_name, size)

	body := GetJSON(url)

	var data Plugin

	// We only want the slug so it can be used to in Bukget URL's

	_, _, slug := ParseJSON(body, data)

	return slug

}

func SearchPlugins(size int, plugin_name string) {

	url := fmt.Sprintf("http://api.bukget.org/3/search/slug/like/%s?size=%d", plugin_name, size)

	body := GetJSON(url)

	var data Plugin

	name, description, slug := ParseJSON(body, data)

	// Loop over the results of ParseJSON. Use the name as the "looper"?

	IterateOverPluginInfo(name, func(i int) {

		fmt.Printf("Plugin name: %s, description: %s, slug: %s\n",
			name[i], description[i], slug[i])

	})

}
