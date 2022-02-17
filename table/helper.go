package table

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

type JsonTask struct {
	Priority int    `json:"priority"`
	Data     string `json:"task"`
}

func Print(res []JsonTask) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Priority", "Task", "ID"})

	for i, v := range res {
		index := strconv.Itoa(i+1)
		priority := strconv.Itoa(v.Priority)
		table.Append([]string{priority, v.Data, index})
	}
	table.Render()
}
