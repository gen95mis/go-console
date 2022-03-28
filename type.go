package dout

import (
	"fmt"
	"strings"
)

type Line struct {
	value string
}

func newLine() *Line {
	return &Line{}
}

// Set formats according to a format specifier and writes to standard output.
func (l *Line) Set(format string, a ...interface{}) {
	mutex.Lock()
	defer mutex.Unlock()

	l.value = fmt.Sprintf(format, a...)
}

// Get returns a Line value
func (l *Line) Get() string {
	return l.value
}

type ProgressBar struct {
	value string
}

func newProgressBar() *ProgressBar {
	return &ProgressBar{}
}

// Set formats according to a format specifier and writes to standard output.
func (p *ProgressBar) Set(description string, current, max int) {
	mutex.Lock()
	defer mutex.Unlock()

	p.value = genProgresBar(description, current, max)
}

// Get returns a ProgressBar value
func (p *ProgressBar) Get() string {
	return p.value
}

func genProgresBar(description string, carrent, max int) string {
	const countLineProcess int = 25

	persent := float32(carrent) / float32(max) * 100
	count := 0
	if persent != 0 {
		count = int(persent) * countLineProcess / (100)
	}

	str1 := strings.Repeat("#", count)
	str2 := strings.Repeat("*", countLineProcess-count)

	result := fmt.Sprintf("%s [%s%s] (%d/%d)", description, str1, str2, carrent, max)

	return result
}
