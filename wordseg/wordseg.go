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
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/crackcell/gonlpir"
	"github.com/crackcell/gonlpir/wordseg/config"
	"os"
	"strings"
)

//===================================================================
// Public APIs
//===================================================================

//===================================================================
// Private
//===================================================================

type EncodingTable struct {
	ictclasOption string
	ictclasType   int
	dec           mahonia.Decoder
}

var encodingTable = []EncodingTable{
	{"gbk", gonlpir.GBK, mahonia.NewDecoder("gbk")},
	{"utf8", gonlpir.UTF8, mahonia.NewDecoder("utf8")},
	{"big5", gonlpir.BIG5, mahonia.NewDecoder("big5")},
	{"gbk_fanti", gonlpir.GBK_FANTI, mahonia.NewDecoder("gbk")},
	{"utf8_fanti", gonlpir.UTF8_FANTI, mahonia.NewDecoder("utf8")},
}

func getEncodingTable(s string) *EncodingTable {
	for _, e := range encodingTable {
		if strings.EqualFold(s, e.ictclasOption) {
			return &e
		}
	}
	return nil
}

func main() {
	config.InitFlags()
	config.Parse()

	encTable := getEncodingTable(config.InputEncoding)
	if encTable == nil {
		fmt.Printf("failed to get encoding table for: %s\n",
			config.InputEncoding)
		os.Exit(1)
	}
	enc := mahonia.NewEncoder(config.OutputEncoding)
	if enc == nil {
		fmt.Printf("failed to create encoder for: %s\n", config.OutputEncoding)
		os.Exit(1)
	}

	seg, err := gonlpir.NewNLPIR(config.DataPath, encTable.ictclasType, "")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()
		for _, result := range seg.ParagraphProcessA(text, true) {
			str := enc.ConvertString(encTable.dec.ConvertString(result.Word))
			if config.ShowPOS {
				str += config.FieldDelimiter + result.Spos
			}
			str += config.LineDelimiter
			fmt.Printf(str)
		}
	}

	seg.Exit()
}
