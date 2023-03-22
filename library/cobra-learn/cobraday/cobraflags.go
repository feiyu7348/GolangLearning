// author:zfy  date:2023/3/22

package cobraday

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var autoNeg string
var speed int

func CobraFlags() {
	rootCmd := &cobra.Command{
		Use:   "pnic [eth0]",
		Short: "pnic",
		Long:  "pnic",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("pnic %s\n", autoNeg)
			fmt.Println("Echo: " + strings.Join(args, " "))
		},
	}

	rootCmd.Flags().StringVarP(&autoNeg, "name", "n", "zfy", "123")
	rootCmd.Flags().IntVarP(&speed, "speed", "s", 1000, "set speed")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
