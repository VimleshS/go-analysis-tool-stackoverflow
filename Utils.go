package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"time"
)

func GetDateTimeFormUnixTimeStamp(dt_tm int64) string {
	//tm := time.Unix(1440719944, 0)
	tm := time.Unix(dt_tm, 0)
	return tm.Format("02-01-2006")
}

func CheckAndCreateTempDirectory(DirName string) {
	file := DirName
	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			if e := os.Mkdir(file, 0777); e != nil {
				panic(e)
			}
		}
	}
}

func UniueIdStr() string {
	uuid := ""
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return uuid
	}
	uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}
