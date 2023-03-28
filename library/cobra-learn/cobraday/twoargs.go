// author:zfy  date:2023/3/22

package cobraday

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var cpu string

func TwoArgs() {
	rootCmd := &cobra.Command{
		Use:   "pnic [eth0]",
		Short: "pnic",
		Long:  "pnic",
		Args:  cobra.MinimumNArgs(1),
		RunE:  listPnic,
	}

	rootCmd.Flags().StringVarP(&cpu, "cpu", "c", "", "input cpu list")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func listPnic(_ *cobra.Command, args []string) error {
	if args == nil {
		log.Printf("args is nil: %s", args)
		return errors.New("args is nil")
	}

	log.Printf("args: %s\n", args)
	// 匹配网卡名称
	pnicRe, err := regexp.MatchString("^eth\\d{1,4}$", args[0])
	if err != nil {
		log.Fatal(err)
		return err
	}

	if pnicRe {
		pnic := args[0]
		log.Printf("pnic: %s\n", pnic)
	} else {
		log.Println("please input the correct nic name!")
		return errors.New("nic name error")
	}

	cpuList := make([]string, 0)
	re, err := regexp.Compile("\\d-\\d")
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Printf("arg[1:]: %s", args[1:])
	for _, v := range args[1:] {
		found := re.MatchString(v)
		if found {
			l := strings.Split(v, "-")
			startInt, _ := strconv.Atoi(l[0])
			endInt, _ := strconv.Atoi(l[1])
			for i := startInt; i <= endInt; i++ {
				cpu := strconv.Itoa(i)
				cpuList = append(cpuList, cpu)
			}
		} else {
			cpuList = append(cpuList, v)
		}
	}

	log.Printf("cpuList: %v\n", cpuList)

	return nil
}
