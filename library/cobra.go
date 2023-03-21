// author:zfy  date:2023/3/21

package library

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var name string

func CobraName() {
	rootCmd := &cobra.Command{
		Use:   "pnic",
		Short: "pnic",
		Long:  "pnic",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("pnic %s", name)
		},
	}

	rootCmd.Flags().StringVarP(&name, "name", "n", "zfy", "123")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
