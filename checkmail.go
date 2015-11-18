package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	cfg        = &config{}
	configFile = flag.String("c", "checkmailrc.json", "Configuration file")
)

func main() {
	flag.Parse()
	if flag.NArg() != 0 {
		flag.Usage()
	}
	err := cfg.parse(*configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	mailboxes := cfg.Mailboxes
	output := []string{}
	for name, path := range mailboxes {
		md := MailDir{
			name: name,
			path: path,
		}
		unseen, err := md.Unseen()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
		output = append(output, fmt.Sprintf("%s: %d", name, len(unseen)))
	}
	fmt.Fprintf(os.Stdout, strings.Join(output, ", "))
	os.Exit(0)
}
