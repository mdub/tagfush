package cmd

import (
	"os"
	"regexp"

	"github.com/docker/distribution/reference"
	rego "github.com/realestate-com-au/dfresh/registry"
	"github.com/realestate-com-au/dfresh/update"
	"github.com/spf13/cobra"
)

var refRegexp = regexp.MustCompile(reference.NameRegexp.String() + "(?::" + reference.TagRegexp.String() + ")?@" + reference.DigestRegexp.String() + "\\b")

func newUpdateCmd(client rego.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "update [flags] [FILE...]",
		Short: "Update image references",
		Args:  cobra.MinimumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return update.UpdateRefsInStream(client, os.Stdin, os.Stdout)
			} else {
				return update.UpdateRefsInFiles(client, args)
			}
		},
	}
}
