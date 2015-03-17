# Go And Hack

This is just for fun! I realize that [Go](http://golang.org/) is a compiled language while [Hack](http://hacklang.org/) is a scripted language, but I consider myself to be fairly fluent in Go, and I want to learn Hack some more. I also understand that these are probably not the best languages to compare, but _why not_? I can do whatever I want, and I'm going to! The structure of this repo is going to be a folder for each example which contains both a Hack file and a Go file and a README which has results from these examples.

## HHVM Config

In order to run Hack scripts with [HHVM](http://hhvm.com/) I need to have a `.hhconfig` file somewhere in the directory. If HHVM can't find the file in the current working directory, [the directory tree will be traversed upward looking for that file until it finds it](http://docs.hhvm.com/manual/en/install.hack.bootstrapping.php). I'm leaving the configuration file empty as of right now.
