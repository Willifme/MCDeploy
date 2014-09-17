package mcdeploy

import (
	"fmt"
)

func DownloadPlugin(size int, plugin_name string) {

	slug := SearchBySlug(100, plugin_name)

	// Use slug[0] as it is the only variable in the array - Make the slug not be an array?

	url := fmt.Sprintf("http://api.bukget.org/3/plugins/bukkit/%s/latest/download?size=%d", slug[0], size)

	fmt.Printf("URL: %s\n", url)

	IterateOverPluginInfo(slug, func(i int) {

		fmt.Printf("Slug %s\n", slug[i])

		file_name := fmt.Sprintf("%s.jar", slug[i])

		GetDownload(url, file_name)

	})

}
