package main

import "flag"

type Cmd struct {
	urls string
	savePath string
	saveMode string
}


func parseCmd() Cmd {
	var cmd Cmd
	flag.StringVar(&cmd.urls, "urls", "https://easylist-downloads.adblockplus.org/easylist.txt|https://easylist-downloads.adblockplus.org/easyprivacy.txt|https://easylist-downloads.adblockplus.org/easylistchina.txt|https://raw.githubusercontent.com/cjx82630/cjxlist/master/cjx-annoyance.txt|https://raw.githubusercontent.com/cjx82630/cjxlist/master/cjxlist.txt", "urls")
	flag.StringVar(&cmd.savePath, "save.path", "./dnsmasq.adb.conf", "save path")
	flag.StringVar(&cmd.saveMode, "save.mode", "a", "save mode [a|t] append or trunc")
	flag.Parse()
	return cmd
}
