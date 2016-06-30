package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/go-yaml/yaml"
)

type aliases struct {
	Target map[string]string `yaml:"target"`
}

type config struct {
	TargetName string  `yaml:"target_name"`
	Target     string  `yaml:"target"`
	Aliases    aliases `yaml:"aliases"`
}

var alias bool

func main() {
	flag.BoolVar(&alias, "alias", false, "alias instead of target name")
	flag.Parse()

	boshTarget := os.Getenv("BOSH_TARGET")
	if boshTarget != "" {
		fmt.Print(boshTarget)
		os.Exit(0)
	}

	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		fmt.Println("$HOME not set")
		os.Exit(1)
	}

	buf, err := ioutil.ReadFile(fmt.Sprintf("%s/.bosh_config", homeDir))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	c := config{}

	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if alias {
		r := regexp.MustCompile("^[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[8|9|aA|bB][a-f0-9]{3}-[a-f0-9]{12}$")
		for key, value := range c.Aliases.Target {
			if value == c.Target {
				if !r.MatchString(key) {
					fmt.Print(key)
					os.Exit(0)
				}
			}
		}
	}

	fmt.Print(c.TargetName)
	os.Exit(0)
}
