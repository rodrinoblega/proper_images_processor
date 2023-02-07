package main

import (
	"flag"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain_1000amount_Should_Not_Save_Files(t *testing.T) {
	//setup
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	flag.CommandLine = flag.NewFlagSet("flags set", flag.ExitOnError)
	os.Args = append([]string{"flags set"}, []string{"-amount", "1000"}...)

	//when
	main()

	//then
	files, _ := os.ReadDir(TestDirectory)
	var count int
	for _, _ = range files {
		count++
	}

	assert.Equal(t, 0, count)
}

func TestMain_1000threads_Should_Not_Save_Files(t *testing.T) {
	//setup
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	flag.CommandLine = flag.NewFlagSet("flags set", flag.ExitOnError)
	os.Args = append([]string{"flags set"}, []string{"-threads", "1000"}...)

	//when
	main()

	//then
	files, _ := os.ReadDir(TestDirectory)
	var count int
	for _, _ = range files {
		count++
	}

	assert.Equal(t, 0, count)
}
