package main

import "github.com/spf13/cobra"

func main() {
	rootCmd := &cobra.Command{
		Use:   "kubebenchctl",
		Short: "A CLI tool to diagnose Kubernetes clusters and provide CIS benchmark scores",
	}

	masterCmd := &cobra.Command{
		Use:   "master [node]",
		Short: "Diagnose the master node",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			node := args[0]
			// Your code here
			return nil
		},
	}

	workerCmd := &cobra.Command{
		Use:   "worker [node]",
		Short: "Diagnose the worker node",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			node := args[0]
			// Your code here
			return nil
		},
	}

	rootCmd.AddCommand(masterCmd, workerCmd)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
