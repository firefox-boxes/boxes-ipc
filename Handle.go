package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/firefox-boxes/boxes"
	"github.com/go-vgo/robotgo"
)

func getParts(str string) []string {
	partsRaw := strings.Split(str, "|")
	parts := make([]string, 0, 1)
	for _, p := range partsRaw {
		parts = append(parts, strings.TrimSpace(p))
	}
	return parts
}

const HELP = `boxes-ipc help
i:ls - List Installations of Firefox
box:ls - List Boxes
box:new <icon path>|<name>|<firefox executable> - Create a new Box
box:del <id> - Delete a Box
box:attrs set <id>|<icon path>|<name>|<firefox executable> - Set the attributes for a box
exec <id> - Start a Box
help - This thing you are looking at`

func Handle(cmd string, attrDB AttrDB, p boxes.ProbeResult) string {
	cmdSplit := strings.SplitN(cmd, " ", 2)
	switch cmdSplit[0] {
		case "i:ls":
			r := ""
			for _, I := range boxes.GetInstallations() {
				r += I.Exec
				r += " "
			}
			return strings.TrimSuffix(r, " ")
		case "box:ls":
			r := ""
			for _, attrs := range attrDB.GetAllBoxes() {
				r += attrs.Id + "|" + attrs.Icon + "|" + attrs.Name + "|" + attrs.Exec + "\n"
			}
			r += "<done>"
			return r
		case "box:new":
			parts := getParts(cmdSplit[1])
			profileID := boxes.NewProfileSetId(p)
			attrDB.AddBox(string(profileID), parts[0], parts[1], parts[2])
			return string(profileID)
		case "box:del":
			profileID := strings.TrimSpace(cmdSplit[1])
			boxes.DeleteProfile(p, boxes.ProfileID(profileID))
			attrDB.DeleteBox(profileID)
			return "<done>"
		case "box:attrs":
			cmdSplit = strings.SplitN(cmdSplit[1], " ", 2)
			parts := getParts(cmdSplit[1])
			switch cmdSplit[0] {
			case "set":
				attrDB.SetBoxAttrs(parts[0], parts[1], parts[2], parts[3])
				return "<done>"
			case "get":
				attrs, err := attrDB.GetBoxAttrs(parts[0])
				if err != nil {
					panic(err)
				}
				return attrs.Icon + "|" + attrs.Name + "|" + attrs.Exec
			default:
				return "boxes-shell: command not found"
			}
		case "whoami":
			pid, err := strconv.Atoi(strings.TrimSpace(cmdSplit[1]))
			if err != nil {
				panic(err)
			}
			targetProcessID := boxes.ProcessID(pid)
			for profileID, processID := range boxes.Exec {
				if targetProcessID == processID {
					return string(profileID)
				}
			}
			return "<nobody>"
		case "exec":
			profileID := boxes.ProfileID(strings.TrimSpace(cmdSplit[1]))
			running, processID := boxes.Exec.IsRunning(profileID)
			attrs, err := attrDB.GetBoxAttrs(string(profileID))
			if err != nil {
				log.Fatal(err)
			}
			if !running {
				installation := boxes.Installation{Exec:attrs.Exec}
				//TODO: Check if installation is legit
				if !installation.IExists() {
					return "<installation doesn't exist>"
				}
				processID = installation.ExecProfileAndDetach(profileID, p)
			}
			for int32(processID) != robotgo.GetPID() {
				robotgo.ActivePID(int32(processID))
			}
			return strconv.Itoa(int(processID))
		case "help":
			return HELP
	}
	return "boxes-shell: command not found"
}