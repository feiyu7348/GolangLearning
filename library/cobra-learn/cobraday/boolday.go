// author:zfy  date:2023/3/29

package cobraday

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var isSaveDb bool

func CobraBool() {
	rootCmd := &cobra.Command{
		Use:   "pnic [eth0]",
		Short: "pnic",
		Long:  "pnic",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("pnic %v\n", isSaveDb)
			fmt.Println("Echo: " + strings.Join(args, " "))
		},
	}

	rootCmd.Flags().BoolVarP(&isSaveDb, "DB", "", false, "123")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

/*
PS E:\GitHub\GolangLearning\library\cobra-learn> go run .\main.go eth0
pnic false
Echo: eth0
PS E:\GitHub\GolangLearning\library\cobra-learn> go run .\main.go eth0 --DB
pnic true
Echo: eth0
*/
