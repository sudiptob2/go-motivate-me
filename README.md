# GO-MOTIVATE-ME CLI

This cli will provide motivational advices in your command line. Built with
golang, cobra

## The examples

### Installation

    go get github.com/golang/example/hello

### Usage

    go-motivate-me [command]
    Available Commands:
    			  help        Help about any command
    			  motivation  Give you a single motivation
    Flags:
          --config string   config file (default is $HOME/.go-motivate-me.yaml)
    	  -h, --help            help for go-motivate-me
    	  -t, --toggle          Help message for toggle

    go-motivate-me motivation [flags]
    Flags:
      -h, --help   help for motivation
      -l, --loop   a boolean flag for starting motivation loop of 10 advices(default false)
