package main

import (
	"bytes"
	"os"
	"os/exec"
	"strconv"
	"testing"
)

func TestFur(t *testing.T) {
	d := "test.db"
	wod := exec.Command("./fur", "-d", d, "-t", "1")
	wodq := exec.Command("./fur", "-d", d, "-q", "0.5")
	wodqw := exec.Command("./fur", "-d", d, "-q", "0.5", "-w", "150")
	wodqwt := exec.Command("./fur", "-d", d, "-q", "0.5", "-w", "150",
		"-t", "8")
	mbad := exec.Command("./fur", "-d", d, "-M")
	d = "masked.db"
	mgoodnom := exec.Command("./fur", "-d", d)
	mgood := exec.Command("./fur", "-d", d, "-M")
	var wants [][]byte
	for i := 0; i < 7; i++ {
		f := "r" + strconv.Itoa(i+1) + ".txt"
		content, err := os.ReadFile(f)
		if err != nil {
			t.Error(err)
		}
		wants = append(wants, content)
	}
	tests := []struct {
		name string
		cmd  *exec.Cmd
		want []byte
	}{
		// Without masking
		{"No Masking default", wod, wants[0]},
		{"No Masking -q 0.5", wodq, wants[1]},
		{"No Masking -q 0.5 -w 150", wodqw, wants[2]},
		{"No Masking -q 0.5 -w 150 -t 8", wodqwt, wants[3]},
		// With masking
		{"Masking no info", mbad, wants[4]},
		{"Masking no -M", mgoodnom, wants[5]},
		{"Masking -M", mgood, wants[6]},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testcmd := test.cmd
			get, err := testcmd.CombinedOutput()
			if err != nil {
				t.Error(err)
			}
			want := test.want
			if !bytes.Equal(get, want) {
				t.Errorf("get:\n%s\nwant:\n%s\n", get, want)
			}
		})
	}
}
