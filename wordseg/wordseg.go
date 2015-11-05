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

func parseEncodingString(s string) int {
	s = strings.ToLower(s)
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
	config.InitFlags()
	config.Parse()

	file, err := os.OpenFile(config.Output, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("open output file failed")
		os.Exit(1)
	}
	defer file.Close()

	seg, err := gonlpir.NewNLPIR(config.DataPath,
		parseEncodingString(config.Encoding), "")
	if err != nil {
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()
		for _, result := range seg.ParagraphProcessA(text, true) {
			str := result.Word
			if config.ShowPOS {
				str += config.FieldDelimiter + result.Spos
			}
			str += config.LineDelimiter
			if _, err := file.WriteString(str); err != nil {
				panic(err)
			}
		}
	}

	seg.Exit()
}
