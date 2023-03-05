package utils

import (
	flag "github.com/spf13/pflag"
)

func init() {
	flag.StringP("time", "t", "45m", "the duration of the session")
	flag.StringP("rest", "r", "15m", "the duration of the rest")
	flag.Parse()
}

func Get_flag(name string) string {
	ret, _ := flag.CommandLine.GetString(name)
	return ret
}
