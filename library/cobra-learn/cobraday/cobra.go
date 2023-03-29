// author:zfy  date:2023/3/21

package cobraday

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var name string

func CobraName() {
	rootCmd := &cobra.Command{
		Use:   "pnic",
		Short: "pnic",
		Long:  "pnic",
		Run: func(cmd *cobra.Command, args []string) {
			for _, v := range args[1:] {
				_, err := strconv.ParseFloat(v, 64)
				fmt.Println(err == nil)
			}

			fmt.Printf("pnic %s", name)
		},
	}

	rootCmd.Flags().StringVarP(&name, "name", "n", "zfy", "123")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
