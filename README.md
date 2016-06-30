# bosh-target
This Go program will print out the currently targeted BOSH director name or alias.  

It will first retrieve your current BOSH target from your `$BOSH_TARGET` env var then `~/.bosh_config`.

The program was created with the intention of printing the BOSH director name or alias so that you may easily add
it to your terminal prompt.

## Usage:
```
$ bosh-target --help
Usage of bosh-target:
  -alias
    	alias instead of target name
```

## Installation
```
$ go install github.com/kkallday/bosh-target
```

### Screenshot
![bosh-target in action](https://raw.githubusercontent.com/kkallday/bosh-target/master/screenshot.png)
