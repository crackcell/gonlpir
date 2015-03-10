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
 * @file example1.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Wed Mar 11 00:01:09 2015
 *
 **/

package main

import (
	"fmt"
	"github.com/crackcell/gonlpir"
)

func main() {
	n, err := gonlpir.NewNLPIR("../", gonlpir.UTF8, "")
	if err != nil {
		panic(err)
	}
	fmt.Println(n.ParagraphProcess("我是中国人", true))
	fmt.Println(n.ParagraphProcess("我是中国人", false))
	n.Exit()
}
