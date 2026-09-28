package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type constructor func() Command

var registry = map[string]constructor{
	"list":   func() Command { return ListCommand{} },
	"add":    func() Command { return AddCommand{} },
	"update": func() Command { return UpdateCommand{} },
	"mark":   func() Command { return MarkCommand{} },
	"delete": func() Command { return DeleteCommand{} },
}

func main() {
	s := &Store{}
	s.InitilizeStore()

	arguments := os.Args

	cmd := registry[arguments[1]]

	Tasks, err := cmd().Run(arguments[2:], s)

	if err != nil {
		fmt.Printf("%v\n", err)
	} else if Tasks != nil {
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tSTATUS\tDescription")
		for _, t := range Tasks {
			fmt.Fprintf(w, "%d\t%s\t%s\n", t.Id, t.Status, t.Description)
		}

		w.Flush()
	} else {
		fmt.Println("Success!!")
	}
}
