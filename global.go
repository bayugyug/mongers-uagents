package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"time"
)

const (
	usageUAList = "Show all the user agent list"
	usageUAGet  = "Show the user agent details"
	usageUAShow = "Show the user agent main list"
)

var (
	//signal flag
	pStillRunning = true
	pBuildTime    = ""
	//stats
	pStats *StatsHelper
	//params
	pUAGet  = ""
	pUAList = ""
	pUAShow = false
	//time
	pAppStart time.Time
	//ver
	pVersion = ""
)

var parameters = map[string]*string{}
var UAHashMap = map[string][]string{
	"UAAndroid":       UAAndroid,
	"UABlackberry":    UABlackberry,
	"UAChromeOs":      UAChromeOs,
	"UAFireOs":        UAFireOs,
	"UAIos":           UAIos,
	"UALinuxOs":       UALinuxOs,
	"UAMac":           UAMac,
	"UAMacOs":         UAMacOs,
	"UAUnixOs":        UAUnixOs,
	"UAWindowsMobile": UAWindowsMobile,
	"UAWindowsOs":     UAWindowsOs,
	"UAWindowsPhone":  UAWindowsPhone,
	"UASymbianOs":     UASymbianOs,
}

type UserAgent struct {
	Agent string   `json:"agent"`
	Total int      `json:"total"`
	Data  []string `json:"data"`
}

func init() {

	pVersion = "Ver 0.01" + " - " + pBuildTime

	//uniqueness
	rand.Seed(time.Now().UnixNano())

	getPrametersFromEnv()
	initEnvParams()

	//more
	runtime.GOMAXPROCS(runtime.NumCPU())

	pStats = StatsHelperNew()
}

//initRecov is for dumpIng segv in
func initRecov() {
	//might help u
	defer func() {
		recvr := recover()
		if recvr != nil {
			log.Println("MAIN-RECOV-INIT: ", recvr)
		}
	}()
}

//getPrametersFromEnv parse envt vars
func getPrametersFromEnv() {
	for k, v := range parameters {
		if os.Getenv(k) != "" {
			*v = os.Getenv(k)
		}
	}

}

//initEnvParams enable all OS envt vars to reload internally
func initEnvParams() {
	//fmt
	flag.StringVar(&pUAList, "l", pUAList, usageUAList+" (short form) ")
	flag.StringVar(&pUAList, "list", pUAList, usageUAList)

	flag.StringVar(&pUAGet, "a", pUAGet, usageUAGet+" (short form) ")
	flag.StringVar(&pUAGet, "agent", pUAGet, usageUAGet)

	flag.BoolVar(&pUAShow, "s", pUAShow, usageUAShow+" (short form) ")
	flag.BoolVar(&pUAShow, "show", pUAShow, usageUAShow)
	flag.Parse()
	if pUAList == "" && pUAGet == "" && !pUAShow {
		showMessage()
	}

}

func showMessage() {

	msg := `
{VERSION}

Usage: {BIN} [options]

     Options are:

`
	msg = strings.Replace(msg, "{BIN}", os.Args[0], -1)
	msg = strings.Replace(msg, "{VERSION}", pVersion, -1)
	fmt.Println()
	fmt.Println()
	fmt.Println(msg)
	flag.PrintDefaults()
	fmt.Println()
	os.Exit(0)
}
