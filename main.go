package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
)

var agent1 string
var options []string
var logoptions []string
var username string
var ipaddress string

func main() {
	lognumber := 1
	lognumberst := strconv.Itoa(lognumber)
	interfaceFlag := flag.Bool("i", false, "Print the target system's IP interfaces.")
	flag.Parse()

	if *interfaceFlag {
		getIPInterfacesFlag()
	}

	io.WriteString(os.Stdout, ` /\_/\
( o.o )
 > ^ <
Hello! Welcome to CAT, the Compact Application Tracker used to detect and log programs and files on other machines using ssh.
Please enter the username of the machine you'd like to connect to. Make sure it's on the same network as this one.
`)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	username = scanner.Text()

	fmt.Println("Okay, now enter the IP address of " + username + ".")

	scanner0 := bufio.NewScanner(os.Stdin)
	scanner0.Scan()
	ipaddress = scanner0.Text()

	agent1 = username + "@" + ipaddress

	options = []string{"1: View and log " + username + "'s running programs", "2: Open the logging menu", "3: View files within home (~)", "4: View files within root (/)", "5: View " + username + "'s IP interfaces", "6: Exit"}
	logoptions = []string{"1: View the list of logs", "2: View the contents of a log", "3: Delete a log", "4: Return to main menu"}

	fmt.Println("I'll be connecting to " + agent1 + ". Please enter the number of the command you'd like me to execute.")

	for i := 0; i < len(options); i++ {
		fmt.Println(options[i])
	}

	a := 0

	for {
		if a != 0 {
			break
		}
		scanner1 := bufio.NewScanner(os.Stdin)
		scanner1.Scan()
		input1 := scanner1.Text()

		switch input1 {
		case "1":
			cmd1, err1 := exec.Command("ssh", agent1, "cd ~/Logs/ && ps -aux", ">", "log"+lognumberst+".txt", "&&", "cat log"+lognumberst+".txt").Output()
			lognumber++
			lognumberst = strconv.Itoa(lognumber)
			if err1 != nil {
				fmt.Println("Uh-oh, looks like I couldn't connect to " + agent1 + ". Please make sure it is online.")
				log.Fatal(err1)
			}
			fmt.Println(string(cmd1))
			fmt.Println("The programs currently running on " + username + " were successfully logged and displayed! What would you like me to do now?")
			for i := 0; i < len(options); i++ {
				fmt.Println(options[i])
			}
		case "2":
			fmt.Println("Welcome to the logging menu! What would you like me to do?")

			for j := 0; j < len(logoptions); j++ {
				fmt.Println(logoptions[j])
			}

			b := 0

			for {
				if b != 0 {
					break
				}

				scanner2 := bufio.NewScanner(os.Stdin)
				scanner2.Scan()
				input2 := scanner2.Text()

				switch input2 {
				case "1":
					cmd11, err11 := exec.Command("ssh", agent1, "cd ~/Logs/ && ls").Output()
					if err11 != nil {
						fmt.Println("Uh-oh, looks like I couldn't connect to " + agent1 + ". Please make sure it is online.")
						log.Fatal(err11)
					}
					fmt.Println(string(cmd11))
					fmt.Println("The logs currently present in " + username + ", if any exist, were successfully displayed! What would you like me to do now?")
					for j := 0; j < len(logoptions); j++ {
						fmt.Println(logoptions[j])
					}
				case "2":
					fmt.Println("Please input the name of the log you would like to view (ex: log1).")
					scanner3 := bufio.NewScanner(os.Stdin)
					scanner3.Scan()
					input3 := scanner3.Text()
					cmd12, err12 := exec.Command("ssh", agent1, "cd ~/Logs/ && cat "+input3+".txt").Output()
					if err12 != nil {
						if input3 == "*" {
							fmt.Println("I wasn't able to print any logs; there are none currently on the system. What would you like me to do now?")
						} else {
							fmt.Println("Sorry, but I couldn't find that log. Try using command 1 to view available logs.")
						}
						for j := 0; j < len(logoptions); j++ {
							fmt.Println(logoptions[j])
						}
					} else {
						fmt.Println(string(cmd12))
						if input3 == "*" {
							fmt.Println("The contents of all existing logs were successfully displayed! What would you like me to do now?")
						} else {
							fmt.Println("The contents of " + input3 + " were successfully printed! What would you like me to do now?")
						}
						for j := 0; j < len(logoptions); j++ {
							fmt.Println(logoptions[j])
						}
					}
				case "3":
					fmt.Println("Please input the name of the log you would like to delete (ex: log1).")
					scanner4 := bufio.NewScanner(os.Stdin)
					scanner4.Scan()
					input4 := scanner4.Text()
					cmd13, err13 := exec.Command("ssh", agent1, "cd ~/Logs/ && rm "+input4+".txt").Output()
					if err13 != nil {
						if input4 == "*" {
							fmt.Println("I wasn't able to delete any logs; there are none currently on the system. What would you like me to do now?")
						} else {
							fmt.Println("Sorry, but I couldn't find that log. Try using command 1 to view available logs.")
						}
						for j := 0; j < len(logoptions); j++ {
							fmt.Println(logoptions[j])
						}

					} else {
						fmt.Println(string(cmd13))
						if input4 == "*" {
							fmt.Println("All logs were successfully deleted! What would you like me to do now?")
						} else {
							fmt.Println(input4 + " was successfully deleted! What would you like me to do now?")
						}
						for j := 0; j < len(logoptions); j++ {
							fmt.Println(logoptions[j])
						}
					}
				case "4":
					fmt.Println("Welcome back to the main menu! Please tell me what you would like me to execute.")
					for i := 0; i < len(options); i++ {
						fmt.Println(options[i])
					}
					b++
				default:
					fmt.Println("Please use one of the available options.")
					for j := 0; j < len(logoptions); j++ {
						fmt.Println(logoptions[j])
					}
				}

			}

		case "3":
			cmd3, err3 := exec.Command("ssh", agent1, "ls").Output()
			if err3 != nil {
				fmt.Println("Uh-oh, looks like I couldn't connect to " + agent1 + ". Please make sure it is online.")
				log.Fatal(err3)
			}
			fmt.Println(string(cmd3))
			fmt.Println("The files within " + username + "'s home folder were successfully displayed! What would you like me to do now?")
			for i := 0; i < len(options); i++ {
				fmt.Println(options[i])
			}
		case "4":
			cmd4, err4 := exec.Command("ssh", agent1, "cd / && ls").Output()
			if err4 != nil {
				fmt.Println("Uh-oh, looks like I couldn't connect to " + agent1 + ". Please make sure it is online.")
				log.Fatal(err4)
			}
			fmt.Println(string(cmd4))
			fmt.Println("The files within " + username + "'s root folder were successfully displayed! What would you like me to do now?")
			for i := 0; i < len(options); i++ {
				fmt.Println(options[i])
			}
		case "5":
			getIPInterfaces()
			fmt.Println(username + "'s IP interfaces were successfully displayed! What would you like me to do now?")
			for i := 0; i < len(options); i++ {
				fmt.Println(options[i])
			}
		case "6":
			io.WriteString(os.Stdout, ` /\_/\ zᶻᶻ
( ~.~ )
 > ^ <
Goodbye!
`)
			a++
		default:
			fmt.Println("Please use one of the available options.")
			for i := 0; i < len(options); i++ {
				fmt.Println(options[i])
			}
		}
	}
}

func getIPInterfaces() {
	cmd5, err5 := exec.Command("ssh", agent1, "go", "run", "interfaces.go").Output()
	if err5 != nil {
		fmt.Println("Uh-oh, looks like I couldn't connect to " + agent1 + ". Please make sure it is online.")
		log.Fatal(err5)
	}
	fmt.Println(string(cmd5))
}

func getIPInterfacesFlag() {
	fmt.Println("Please enter the handle of the target machine.")

	scanner10 := bufio.NewScanner(os.Stdin)
	scanner10.Scan()
	agent1 = scanner10.Text()

	cmd5, err5 := exec.Command("ssh", agent1, "go", "run", "interfaces.go").Output()
	if err5 != nil {
		fmt.Println("Uh-oh, looks like I couldn't connect to " + agent1 + ". Please make sure it is online.")
		log.Fatal(err5)
	}
	fmt.Println(string(cmd5))
	fmt.Println(agent1 + "'s IP interfaces were successfully displayed using the -i shortcut! Exiting now...")
	os.Exit(1)
}
