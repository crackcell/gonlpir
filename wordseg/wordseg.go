/***************************************************************
 *
 * Copyright (c) 2015, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 *
 *
 * @file wordseg.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Tue Mar 10 14:28:13 2015
 *
 **/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/crackcell/gonlpir"
	"os"
	"strings"
)

//===================================================================
// Public APIs
//===================================================================

//===================================================================
// Private
//===================================================================

var (
	help     bool
	verbose  bool
	dataPath string
	encoding string
)

const (
	LogoString = ` _______         _______ _____   ______ _______ ______
|     __|.-----.|    |  |     |_|   __ \_     _|   __ \
|    |  ||  _  ||       |       |    __/_|   |_|      <
|_______||_____||__|____|_______|___|  |_______|___|__|`

	HelpString = `Standalone GoNLPIR wordseg
Usage:
    wordseg [options]
Options:
    -h, --help     Print this message
    -v, --verbose  Use verbose output
    -d, --data     Root path of data
    -e, --encoding Encoding
`
)

func showHelp() {
	fmt.Println(LogoString)
	fmt.Println()
	fmt.Println(HelpString)
	os.Exit(0)
}

func parseEncodingString(s string) int {
	if strings.EqualFold(s, "gbk") {
		return gonlpir.GBK
	}
	if strings.EqualFold(s, "utf8") {
		return gonlpir.UTF8
	}
	if strings.EqualFold(s, "big5") {
		return gonlpir.BIG5
	}
	if strings.EqualFold(s, "gbk_fanti") {
		return gonlpir.GBK_FANTI
	}
	if strings.EqualFold(s, "utf8_fanti") {
		return gonlpir.UTF8_FANTI
	}

	return gonlpir.UTF8
}

func main() {
	flag.BoolVar(&help, "help", false, "Print help message")
	flag.BoolVar(&help, "h", false, "Print help message")
	flag.BoolVar(&verbose, "verbose", false, "Use verbose output")
	flag.BoolVar(&verbose, "v", false, "Use verbose output")
	flag.StringVar(&dataPath, "data", "./", "Path of data folder")
	flag.StringVar(&dataPath, "d", "./", "Path of data folder")
	flag.StringVar(&encoding, "encoding", "UTF8", "Encoding: UTF8")
	flag.StringVar(&encoding, "e", "UTF8", "Encoding: GBK, UTF8, BIG5, GBK_FANTI, UTF8_FANTI")

	flag.Parse()
	if help {
		showHelp()
	}

	seg, err := gonlpir.NewNLPIR(dataPath, parseEncodingString(encoding), "")
	if err != nil {
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()
		for _, result := range seg.ParagraphProcessA(text, true) {
			fmt.Println(result.Word)
		}
	}

	seg.Exit()
}
