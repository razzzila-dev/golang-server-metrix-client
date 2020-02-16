package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

type Config struct {
	UrlPrefix string `yaml:"url_prefix"`
	Port      string `yaml:"port"`
	Key       string `yaml:"key"`
}

func main() {
	config := readConfigFile()
	metrics := new(Metrics)

	http.HandleFunc(getUrl(&config, "/"), func(w http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(metrics.UpdateAll())

		if err != nil {
			_, err := fmt.Fprintf(w, string(data))

			if err != nil {
				fmt.Print(err)
				_, err := fmt.Fprintf(w, "null")

				if err != nil {
					fmt.Print(err)
				}

				return
			}
		}

		if err != nil {
			fmt.Print(err)
			_, err := fmt.Fprintf(w, "null")

			if err != nil {
				fmt.Print(err)
			}
		}
	})
	http.HandleFunc(getUrl(&config, "/cpu"), func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, metrics.GetCPU())
		if err != nil {
			fmt.Print(err)
			_, err := fmt.Fprintf(w, "null")

			if err != nil {
				fmt.Print(err)
			}
		}
	})
	http.HandleFunc(getUrl(&config, "/ram"), func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, metrics.GetRAM())
		if err != nil {
			fmt.Print(err)
			_, err := fmt.Fprintf(w, "null")

			if err != nil {
				fmt.Print(err)
			}
		}
	})
	http.HandleFunc(getUrl(&config, "/disk"), func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, metrics.GetDisk())
		if err != nil {
			fmt.Print(err)
			_, err := fmt.Fprintf(w, "null")

			if err != nil {
				fmt.Print(err)
			}
		}
	})

	err := http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		fmt.Print(err)
	}
}

func readConfigFile() Config {
	filename, _ := filepath.Abs("config.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}
	return config
}

func getUrl(config *Config, path string) string {
	return "/" + strings.Trim(config.UrlPrefix, "/") + "/" + config.Key + "/" + strings.Trim(path, "/")
}
