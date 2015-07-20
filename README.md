# GoNLPIR
[![Build Status](https://travis-ci.org/crackcell/gonlpir.svg?branch=master)](https://travis-ci.org/crackcell/gonlpir)

## Introduction

GoNLPIR is a Golang wrapper for the famous Chinese word segmenter NLPIR(former ICTCLAS ) with the following interface implemented:

- ParagraphProcess
- ParagraphProcessA
- ImportUserDict

## Install

Get code and install dependences

    go get github.com/crackcell/gonlpir
    make install_deps
    go test -v

Example

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

## Documents

- Symbols of part of speech: https://github.com/crackcell/gonlpir/blob/master/doc/word_types.md
- Learn more: http://blog.crackcell.com/posts/2015/03/10/gonlpir_golang_wrapper_for_nlpir_ictclas.html

## Misc

- Official NLPIR site:
  - http://www.nlpir.org/
  - https://github.com/NLPIR-team/NLPIR