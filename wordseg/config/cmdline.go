/***************************************************************
 *
 * Copyright (c) 2014, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 *
 *
 * @file cmdline.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sun Nov 23 20:59:30 2014
 *
 **/

package config

import (
	"flag"
	"fmt"
	"os"
)

//===================================================================
// Public APIs
//===================================================================

var (
	Help           bool
	Verbose        bool
	Output         string
	DataPath       string
	Encoding       string
	LineDelimiter  string
	FieldDelimiter string
	ShowPOS        bool
)

func InitFlags() {
	flag.BoolVar(&Help, "help", false, "Print help message")
	flag.BoolVar(&Help, "h", false, "Print help message")
	flag.BoolVar(&Verbose, "verbose", false, "Use verbose output")
	flag.BoolVar(&Verbose, "v", false, "Use verbose output")
	flag.StringVar(&Output, "o", "", "Output file")
	flag.StringVar(&Output, "output", "", "Output file")
	flag.StringVar(&DataPath, "data", "", "ICTCLAS data path")
	flag.StringVar(&DataPath, "d", "", "ICTCLAS data path")
	flag.StringVar(&Encoding, "e", "UTF8", "Encoding")
	flag.StringVar(&Encoding, "encoding", "UTF8", "Encoding")
	flag.StringVar(&LineDelimiter, "ld", "\n", "Line delimiter")
	flag.StringVar(&LineDelimiter, "line-delimiter", "\n", "Line delimiter")
	flag.StringVar(&FieldDelimiter, "fd", "\t", "Field delimiter")
	flag.StringVar(&FieldDelimiter, "field-delimiter", "\t", "Field delimiter")
	flag.BoolVar(&ShowPOS, "pos", false, "Show POS (part of speech) info")
}

func Parse() {
	flag.Parse()
	if Help {
		showHelp(0)
	}
	if len(DataPath) == 0 {
		fmt.Println("invalid arguments: no data")
		showHelp(1)
	}
	if len(Output) == 0 {
		fmt.Println("invalid arguments: no output")
		showHelp(1)
	}
}

//===================================================================
// Private
//===================================================================

const (
	logoString = `                        __
.--.--.--.-----.----.--|  |.-----.-----.-----.
|  |  |  |  _  |   _|  _  ||__ --|  -__|  _  |
|________|_____|__| |_____||_____|_____|___  |
                                       |_____|
`
	helpString = `wordseg
Usage:
    wordseg [options]
Options:
    -h, --help               Print this message
    -v, --verbose            Use verbose output

    -o, --output             Output file

    -d, --data               Data path

    -e, --encoding           Specify encoding of input data
                             available: gbk, utf8, big5, gbk_fanti, utf8_fanti

    --ld, --line-delimiter   Specify line delimiter, default: \\n
    --fd, --field-delimiter  Specify field delimiter, default: \\t

    --pos                    Show POS (Part of speech) info
`
)

func showHelp(ret int) {
	fmt.Print(helpString)
	os.Exit(ret)
}

func LogoString() string {
	return logoString
}
