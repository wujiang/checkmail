package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
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
	names := []string{}
	for k := range mailboxes {
		names = append(names, k)
	}
	sort.Strings(names)
	for _, name := range names {
		md := MailDir{
			name: name,
			path: mailboxes[name],
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
