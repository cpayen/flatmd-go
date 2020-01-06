package path

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	path pathConfig
)

type pathConfig struct {
	Content     string
	PostData    string
	PostContent string
}

func getConfig() pathConfig {
	if (pathConfig{}) == path {
		configfile, err := ioutil.ReadFile("./config.path.json")
		if err == nil {
			path := pathConfig{}
			json.Unmarshal(configfile, &path)
		}
	}
	return path
}

// GetContentPath is..
func GetContentPath() string {
	return getConfig().Content
}

// GetPostDataPathFor is..
func GetPostDataPathFor(postSlug string) string {
	return fmt.Sprintf(getConfig().PostData, postSlug)
}

// GetPostContentPathFor is..
func GetPostContentPathFor(postSlug string) string {
	return fmt.Sprintf(getConfig().PostContent, postSlug)
}
