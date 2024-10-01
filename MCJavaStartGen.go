package main

import (
	"fmt"
	"strconv"
)

const (
	AIKAR string = "-XX:+AlwaysPreTouch -XX:+DisableExplicitGC -XX:+ParallelRefProcEnabled -XX:+PerfDisableSharedMem -XX:+UnlockExperimentalVMOptions -XX:+UseG1GC -XX:G1HeapRegionSize=8M -XX:G1HeapWastePercent=5 -XX:G1MaxNewSizePercent=40 -XX:G1MixedGCCountTarget=4 -XX:G1MixedGCLiveThresholdPercent=90 -XX:G1NewSizePercent=30 -XX:G1RSetUpdatingPauseTimePercent=5 -XX:G1ReservePercent=20 -XX:InitiatingHeapOccupancyPercent=15 -XX:MaxGCPauseMillis=200 -XX:MaxTenuringThreshold=1 -XX:SurvivorRatio=32 -Dusing.aikars.flags=https://mcflags.emc.gs -Daikars.new.flags=true"
	VALOCITY string = "-XX:+AlwaysPreTouch -XX:+ParallelRefProcEnabled -XX:+UnlockExperimentalVMOptions -XX:+UseG1GC -XX:G1HeapRegionSize=4M -XX:MaxInlineLevel=15"
)

var (
	mainFile string
	maxRam   int
	minRam   int
	gui      bool
	flags    string
)

func main() {
	var value string

	fmt.Println("Minecraft 用 JVM 變數生成器")
	
	fmt.Print("伺服器檔案 (server.jar) > ")
	fmt.Scanln(&mainFile)

	if mainFile == "" {
		mainFile = "server.jar"
	}

	fmt.Print("最大記憶體 (4096) MB > ")
	fmt.Scanln(&value)
	
	if value == "" {
		maxRam = 4096
	} else {
		var err error
		maxRam, err = strconv.Atoi(value)
		if err != nil {
			fmt.Println("無效的輸入，將使用預設值 4096 MB")
			maxRam = 4096
		}
	}

	fmt.Print("最小記憶體 (1024) MB > ")
	fmt.Scanln(&value)

	if value == "" {
		minRam = 1024
	} else {
		var err error
		var check int
		check, err = strconv.Atoi(value)
		if err != nil || check > maxRam {
			fmt.Println("無效的輸入，將使用預設值 1024 MB")
			minRam = 1024
		} else {
			minRam = check
		}
	}

	fmt.Print("是否彈出伺服器GUI? (N) [Yy/Nn] > ")
	fmt.Scanln(&value)

	if value == "Y" || value == "y" {
		gui = true
	} else if value == "N" || value == "n" || value == "" {
		gui = false
	} else {
		fmt.Println("無效的輸入，將使用預設值 N")
		gui = false
	}

	fmt.Println("請選擇伺服器優化?")
	fmt.Println("0) 無")
	fmt.Println("1) Aikar's")
	fmt.Println("2) Velocity")
	fmt.Print("(0) >")
	fmt.Scanln(&value)

	switch value {
		case "0":
			flags = ""
			break
		case "1":
			flags = AIKAR
			break
		case "2":
			flags = VALOCITY
			break
		case "":
			flags = ""
			break
		default:
			fmt.Println("無效的輸入，將使用預設值 0")
			flags = ""
	}

	output := "java -Xmx%dM -Xms%dM%s -jar %s%s\n"
	var guiOption string
	if gui {
		guiOption = " nogui"
	} else {
		guiOption = ""
	}

	var flagsOption string
	if flags != "" {
		flagsOption = " " + flags
	}

	fmt.Println("= = = = = = = = = = = = = = = = = = = =")
	fmt.Println()
	fmt.Printf(output, maxRam, minRam, flagsOption, mainFile, guiOption)
	fmt.Println()
	fmt.Println("= = = = = = = = = = = = = = = = = = = =")

	fmt.Println("按下 Enter 結束...")
	fmt.Scanln();

}
