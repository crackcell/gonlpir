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
 * @file ggnlpir_test.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sun Mar  8 15:40:16 2015
 *
 **/

package ggnlpir

import (
	"fmt"
	"testing"
)

//===================================================================
// Public APIs
//===================================================================

func TestParagraphProcess(t *testing.T) {
	nlpir, err := NewNLPIR("./", UTF8, "")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("TestParagraphProcess: 我是中国人")
	fmt.Println(nlpir.ParagraphProcess("我是中国人", true))
}

func TestParagraphProcessA(t *testing.T) {
	nlpir, err := NewNLPIR("./", UTF8, "")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("TestParagraphProcessA: 我是中国人")
	results := nlpir.ParagraphProcessA("我是中国人", true)
	for i := 0; i < len(results); i++ {
		fmt.Printf("%#v\n", results[i])
	}
}

//===================================================================
// Private
//===================================================================
