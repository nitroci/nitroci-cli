/*
Copyright 2021 The NitroCI Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package terminal

import (
	"fmt"
	"strings"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

type TerminalMessageType int64

const (
	Info TerminalMessageType = iota
	Warning
	Error
)

func (t TerminalMessageType) color(isChange bool) string {
	switch t {
	case Info:
		if isChange {
			return string(colorGreen)
		}
		return string(colorReset)
	case Warning:
		return string(colorYellow)
	case Error:
		return string(colorRed)
	}
	return string(colorReset)
}

type TerminalItemsOutput struct {
	Messages    []string
	Suggestions []string
	Items       []string
	ItemsType   TerminalMessageType
}

type TerminalOutput struct {
	Messages    []string
	MessageType TerminalMessageType
	ItemsOutput []TerminalItemsOutput
	Output      string
}

type terminalActionOutputCursor struct {
	stepPrinted  bool
	outpustIndex int
}

type TerminalActionOutput struct {
	StepId  string
	Step    string
	Outputs []string
	cursor  *terminalActionOutputCursor
}

func printStrings(color string, strings []string) {
	printStringsWithTabs(color, 0, strings)
}

func printStringsWithTabs(color string, tabs int, strs []string) {
	for _, m := range strs {
		fmt.Print(string(color), strings.Repeat("  ", tabs)+m)
		fmt.Println()
	}
}

func Print(output *TerminalOutput) {
	if output.Messages != nil && len(output.Messages) > 0 {
		printStrings(output.MessageType.color(false), output.Messages)
		fmt.Print(string(colorReset))
	}
	if output.ItemsOutput != nil && len(output.ItemsOutput) > 0 {
		for _, m := range output.ItemsOutput {
			fmt.Println()
			printStrings(string(colorReset), m.Messages)
			printStringsWithTabs(string(colorReset), 1, m.Suggestions)
			printStringsWithTabs(m.ItemsType.color(true), 4, m.Items)
		}
		fmt.Println(string(colorReset))
	}
	if output != nil && len(output.Output) > 0 {
		fmt.Println(output.Output)
	}
}

func PrintActions(output *TerminalActionOutput) {
	if output.cursor == nil {
		output.cursor = &terminalActionOutputCursor{
			stepPrinted:  false,
			outpustIndex: -1,
		}
	}
	if (*output.cursor).stepPrinted == false {
		m := fmt.Sprintf("%s", output.Step)
		if len(output.StepId) > 0 {
			m = fmt.Sprintf("%s %s", output.StepId, m)
		}
		fmt.Print(string(colorGreen), m)
		fmt.Print(string(colorReset))
		(*output.cursor).stepPrinted = true
		fmt.Println()
	}
	if output.Outputs != nil && len(output.Outputs) > 0 {
		for i, m := range output.Outputs {
			if (*output.cursor).outpustIndex >= i {
				continue
			}
			fmt.Println(m)
			(*output.cursor).outpustIndex = i
		}
	}
}
