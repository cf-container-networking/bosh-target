package main

import (
	"errors"
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

var wantAlias bool

func main() {
	flag.BoolVar(&wantAlias, "alias", false, "prints BOSH director's alias instead of name")
	flag.Parse()

	boshTarget := os.Getenv("BOSH_TARGET")
	if boshTarget != "" {
		fmt.Print(boshTarget)
		os.Exit(0)
	}

	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		fail(errors.New("$HOME not set"))
	}

	buf, err := ioutil.ReadFile(fmt.Sprintf("%s/.bosh_config", homeDir))
	if err != nil {
		fail(err)
	}

	c := config{}

	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		fail(err)
	}

	if wantAlias {
		r := regexp.MustCompile("^[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[8|9|aA|bB][a-f0-9]{3}-[a-f0-9]{12}$")
		for alias, target := range c.Aliases.Target {
			if !r.MatchString(alias) && target == c.Target {
				fmt.Print(alias)
				os.Exit(0)
			}
		}
		fail(errors.New(fmt.Sprintf("No alias found for %q", c.Target)))
	} else {
		fmt.Print(c.TargetName)
	}
}

func fail(err error) {
	fmt.Fprint(os.Stderr, err)
	os.Exit(1)
}
