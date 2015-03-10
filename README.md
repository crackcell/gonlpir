# GoNLPIR

## 简介

中科院分词（ICTCLAS/NLPIR）的Go语言封装，目前提供的接口:

- ParagraphProcess
- ParagraphProcessA
- ImportUserDict

## 安装

获得代码

    go get github.com/crackcell/gonlpir

示例代码



## 示例

    go test

输出

    TestParagraphProcess: 我是中国人
    我/rr 是/vshi 中国/ns 人/n
    TestParagraphProcessA: 我是中国人
    &gonlpir.Result{word:"我", spos:[]string{"r", "r"}, ipos:42, wordId:11795, wordType:0, weight:6305}
    &gonlpir.Result{word:"是", spos:[]string{"v", "s", "h", "i"}, ipos:75, wordId:40459, wordType:0, weight:4710}
    &gonlpir.Result{word:"中国", spos:[]string{"n", "s"}, ipos:29, wordId:32696, wordType:0, weight:6097}
    &gonlpir.Result{word:"人", spos:[]string{"n"}, ipos:21, wordId:37533, wordType:0, weight:6055}
    PASS

## 文档

- 汉语词性对照表: https://github.com/crackcell/gonlpir/blob/master/doc/word_types.md
- 了解更多: http://blog.crackcell.com/posts/2015/03/10/gonlpir_golang_wrapper_for_nlpir_ictclas.html

