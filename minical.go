package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func showv() {
	VERSIONSEM := "0.0.1"
	VERSIONBUI := "286"
	fmt.Println("\nminical v." + VERSIONSEM + " b." + VERSIONBUI + "\n")
}

func main() {
	showv()
	time.Sleep(2 * time.Second)
	t := genslots(7, 16, []int{0, 30})
	fmt.Println(renderout(t))
	do_go_on := true
	// inp := ""
	reader := bufio.NewReader(os.Stdin)
	for do_go_on {
		fmt.Print(">>> ")
		// fmt.Scanln(&inp)
		inp, _ := reader.ReadString('\n')
		inp = inp[:len(inp)-1]
		if inp == "help" || inp == "?" {
			printhelp()
		} else if inp == "quit" || inp == "exit" {
			do_go_on = false
		} else if inp == "show" {
			fmt.Println("\n" + renderout(t))
		} else if inp == "clear" {
			fmt.Print("Type `Y` in capital to confirm: ")
			// fmt.Scanln(&inp)
			inp, _ := reader.ReadString('\n')
			inp = inp[:len(inp)-1]
			if inp == "Y" {
				t = genslots(7, 16, []int{0, 30})
				fmt.Println("\n" + renderout(t))
			} else {
				fmt.Println("Aborted clear operation.\n")
			}
		} else if inp[0:int(math.Min(4, float64(len(inp))))] == "set " {
			s := strings.Split(inp, " ")
			if len(s) < 2 {
				fmt.Println("Expected `set x y`, got `" + inp + "`")
			} else {
				o := ""
				for i, j := range s {
					if i >= 2 {
						o += j
					}
				}
				c := 0
				for i, j := range t {
					if j.timecode == s[1] {
						c += 1
						if o != "" {
							t[i].desc = o
						} else {
							t[i].desc = "<<< EMPTY >>>"
						}
					}
				}
				if c != 1 {
					fmt.Println(strconv.Itoa(c) + " matching timecodes, not 1!")
				} else {
					fmt.Println(renderout(t))
				}
			}
		} else {
			fmt.Println("Unrecognized command `" + inp + "`\n")
		}
	}
}

func printhelp() {
	t := []string{
		"",
		"? or help    --> show this page",
		"set x y      --> set slot x to description y",
		"show         --> refresh the view",
		"clear        --> clear all slots",
		"exit or quit --> exit the program",
		"",
	}
	for _, i := range t {
		fmt.Println(i)
	}
}

func genslots(shour int, ehour int, mins []int) []Slot {
	if shour > ehour {
		panic("shour " + strconv.Itoa(shour) + " > ehour " + strconv.Itoa(ehour))
	}
	hrs := []string{}
	for i := shour; i < ehour; i++ {
		hrs = append(hrs, rpad(strconv.Itoa(i), 2, "0"))
	}
	mns := []string{}
	for _, i := range mins {
		mns = append(mns, rpad(strconv.Itoa(i), 2, "0"))
	}
	tcs := gentimecodes(hrs, mns)
	t := []Slot{}
	for _, i := range tcs {
		t = append(t, Slot{i, "<<< EMPTY >>>"})
	}
	return t
}

type Slot struct {
	timecode string
	desc     string
}

func (x Slot) srep() string {
	m := int(math.Min(76, float64(len(x.desc))))
	return x.timecode + " : " + x.desc[0:m]
}

func gentimecodes(hours []string, mins []string) []string {
	o := []string{}
	for _, h := range hours {
		for _, m := range mins {
			o = append(o, h+m)
		}
	}
	return o
}

func renderout(slots []Slot) string {
	s_len := 60
	s_begin := "| "
	s_end := " |\n"
	s_mlen := s_len - len(s_begin) - len(s_end) + 1
	o := s_begin + lpad("TIME : TASK", s_mlen, " ") + s_end + s_begin + lpad("---- : ", s_mlen, "-") + s_end
	for _, i := range slots {
		o += s_begin + lpad(i.srep(), s_mlen, " ") + s_end
	}
	return o
}

func lpad(s string, l int, pchar string) string {
	if len(s) == l {
		return s
	} else if len(s) > l {
		return s[0:l]
	} else {
		o := s
		for len(o) < l {
			o += pchar
		}
		return o
	}
}

func rpad(s string, l int, pchar string) string {
	if len(s) == l {
		return s
	} else if len(s) > l {
		return s[0:l]
	} else {
		o := s
		for len(o) < l {
			o = pchar + o
		}
		return o
	}
}
