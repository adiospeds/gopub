package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

//getRedisMaster - Gets the master IP reported by Redis command:
//	info replication.
func getRedisMaster(ip, rPort, pass string) (op string, err error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     ip + ":" + rPort,
		Password: pass,
		DB:       0, // use default DB
	})
	rInfo, err := rdb.Info(ctx, "Replication").Result()
	return rInfo, err
}

//getSentinelMaster - Gets the master IP reported by Sentinel command:
//	sentinel masters
func getSentinelMaster(ip, sPort string) (op string, err error) {
	srdb := redis.NewSentinelClient(&redis.Options{
		Addr: ip + ":" + sPort,
		DB:   0, // use default DB
	})
	sMasterInfo, err := srdb.Masters(ctx).Result()
	sMasterInfoInf := sMasterInfo[0]
	var sMasterInfoSlice = make([]string, len(sMasterInfoInf.([]interface{})))
	for i, item := range sMasterInfoInf.([]interface{}) {
		sMasterInfoSlice[i] = item.(string)
	}
	sIP := sMasterInfoSlice[3]
	return sIP, err
}

// parseRepInfo :
// Parses the 'info replication' output to find redis server master ip (rIP)
func parseRepInfo(ip, repInfo string) string {
	rIP := "NotFound"
	find := func(slice []string, val string) (int, bool) {
		for i, item := range slice {
			if strings.Contains(item, val) {
				return i, true
			}
		}
		return -1, false
	}

	repInfoArr := strings.Split(repInfo, "\r\n")
	_, slaveChk := find(repInfoArr, "role:slave")

	if slaveChk {
		idx, _ := find(repInfoArr, "master_host")
		rIP = strings.Split(repInfoArr[idx], ":")[1]
	}

	_, masterChk := find(repInfoArr, "role:master")
	if masterChk {
		rIP = ip
	}

	return rIP
}

func chkRequiredFields(ip string) {
	if ip == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main() {
	ip := flag.String("ip", "", "(Required field)\nTakes the ip of the redis-server.\nAvoid passing localhost or 127.0.0.1, unless sentinel/redis is only binded on it.")
	rPort := flag.String("redis-port", "6379", "Takes the port on which redis-server is listening.")
	sPort := flag.String("sentinel-port", "26379", "Takes the port on which sentinel server is listening.")
	pass := flag.String("pass", "", "Takes the auth password for redis server.")
	flag.Parse()
	chkRequiredFields(*ip)

	repInfo, err := getRedisMaster(*ip, *rPort, *pass)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	sIP, err := getSentinelMaster(*ip, *sPort)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	rIP := parseRepInfo(*ip, repInfo)

	if sIP != rIP {
		fmt.Printf("Redis reported master IP(%s) is not same as Sentinel reported master IP(%s)", rIP, sIP)
		os.Exit(2)
	}
	fmt.Printf("Redis reported master IP(%s) is same as Sentinel reported master IP(%s)", rIP, sIP)
	os.Exit(0)
}
