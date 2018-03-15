// Copyright © 2018 Dhananjay Balan <mail@dbalan.in>
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice,
//    this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors
//    may be used to endorse or promote products derived from this software
//    without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package cmd

import (
	"fmt"

	"github.com/dbalan/pipet/pipetdata"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var body bool

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:     "show uid",
	Short:   "display the snippet",
	Args:    cobra.ExactArgs(1),
	PreRunE: ensureConfig,
	Run: func(cmd *cobra.Command, args []string) {
		dataStore := getDataStore()
		snip, err := dataStore.Read(args[0])
		errorGuard(err, "reading snippet failed")
		if body {
			fmt.Printf(snip.Data)
		} else {
			fmt.Printf(fancySnippet(snip))
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.PersistentFlags().BoolVarP(&body, "body-only", "b", false, "show only snippet content")
}

func fancySnippet(s *pipetdata.Snippet) string {
	green := color.New(color.FgGreen).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	sep := green("---\n")

	text := sep + green("Title: ") + fmt.Sprint(s.Meta.Title) + green("\nTags:\n")
	for _, t := range s.Meta.Tags {
		text += green("- ") + blue(t) + "\n"
	}
	text += sep
	text += s.Data
	return text
}
