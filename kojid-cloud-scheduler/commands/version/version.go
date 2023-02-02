package version

import (
    "fmt"

    "github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
    Use:   "version",
    Short: "Shows the version of kojid-cloud-scheduler",
    Long:  `All software has versions. This is kojid-cloud-scheduler.`,
    Run: func(cmd *cobra.Command, args []string) {
        // TODO: move it someplace else
        version := "0.1.0"
        fmt.Printf("%s\n", version)
    },  
}
