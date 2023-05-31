package main

import (
	"fmt"
	"github.com/bitfield/script"
	"github.com/briandowns/spinner"
	"github.com/spf13/pflag"
	"strings"
	"time"
)

type Spin *spinner.Spinner

var elapsed bool

func main() {
	var delay time.Duration
	var spinType int
	var demo bool
	pflag.DurationVarP(&delay, "delay", "d", 100*time.Millisecond, "Delay in milliseconds")
	pflag.IntVarP(&spinType, "type", "t", 9, "Type")
	pflag.BoolVarP(&elapsed, "elapsed", "e", false, "elapsed time.")
	pflag.BoolVarP(&demo, "demo", "D", false, "Demo -- shows all possible spin types.")
	pflag.Parse()
	s := spinner.New(spinner.CharSets[spinType], delay)
	if demo {
		runDemo(s)
	} else {
		cmd := strings.Join(pflag.Args(), " ")
		runCommand(s, cmd)
	}
}

func runCommand(s *spinner.Spinner, cmd string) {
	s.Start()
	ts := time.Now()
	script.Exec(cmd).Stdout()
	s.Stop()
	te := time.Now()
	d := te.Sub(ts)
	if elapsed {
		fmt.Println("Time", d)
	}
}

func runDemo(s *spinner.Spinner) {
	for i := 0; i <= 43; i++ {
		pre := fmt.Sprintf("%d ", i)
		s.Prefix = pre
		s.UpdateCharSet(spinner.CharSets[i])
		s.Start()
		time.Sleep(3 * time.Second)
		s.Stop()
	}
}
