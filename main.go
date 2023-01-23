package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hugolgst/rich-go/client"
	"github.com/mitchellh/go-ps"
)

func main() {
	appStatus := false
	rpcStatus := false

	for {
		processes, err := ps.Processes()
		appStatus = false

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, process := range processes {
			if process.Executable() == "Notion.exe" {
				appStatus = true
			}
		}

		if !appStatus {
			client.Logout()
			rpcStatus = false
		}

		if appStatus {
			if rpcStatus == false {
				err := client.Login("1067066169703542874")

				if err != nil {
					appStatus = false
				}

				currentTime := time.Now()

				client.SetActivity(client.Activity{
					State:   "Notion",
					Details: "Creating ✏",
					Timestamps: &client.Timestamps{
						Start: &currentTime,
					},
					LargeImage: "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e9/Notion-logo.svg/1024px-Notion-logo.svg.png",
				})

				rpcStatus = true
			}
		}

		time.Sleep(3 * time.Second)
	}
}
