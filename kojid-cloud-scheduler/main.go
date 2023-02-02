package main

import (
    "os"
    "fmt"
    version_cmd "github.com/odra/cowji/kojid-cloud-scheduler/commands/version"
    "github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
    Use:   "kojid-cloud-scheduler",
    Short: "kojid-cloud-scheduler is something",
    Long:  `kojid-cloud-scheduler is something`,
    RunE: func(cmd *cobra.Command, args []string) error {
        return cmd.Help()
    },
}

func main() {
    RootCmd.AddCommand(version_cmd.Cmd)

    if err := RootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Error: %s\n", err)
        os.Exit(1)
    }
}

