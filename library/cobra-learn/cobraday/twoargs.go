// author:zfy  date:2023/3/22

package cobraday

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var gro string
var mtu int

func TwoArgs() {
	rootCmd := &cobra.Command{
		Use:   "pnic [eth0]",
		Short: "pnic",
		Long:  "pnic",
		Args:  cobra.MinimumNArgs(2),
		RunE:  listPnic,
	}

	rootCmd.Flags().StringVarP(&gro, "gro", "g", "off", "set gro")
	rootCmd.Flags().IntVarP(&mtu, "mtu", "m", 0, "set mtu")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func listPnic(_ *cobra.Command, args []string) error {
	if args == nil {
		log.Printf("args is nil: %s", args)
		return nil
	}

	fmt.Printf("args: %s\n", args)
	//cpuList := make([]string, 0)
	for _, v := range args {
		fmt.Println(v)

	}

	log.Printf("mtu: %d", mtu)
	return nil
}
