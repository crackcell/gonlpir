# gonlpir

## 简介

中科院分词（ICTCLAS/NLPIR）的Go语言封装，目前提供的接口:

- ParagraphProcess
- ParagraphProcessA
- ImportUserDict

## 安装

获得代码

    go get github.com/crackcell/gonlpir

示例代码

先将include和lib文件夹加到头文件和库文件搜索路径中去

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

## 文档

- 汉语词性对照表: https://github.com/crackcell/gonlpir/blob/master/doc/word_types.md
- 了解更多: http://blog.crackcell.com/posts/2015/03/10/gonlpir_golang_wrapper_for_nlpir_ictclas.html
