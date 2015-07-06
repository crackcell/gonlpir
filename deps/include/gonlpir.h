/* -*- encoding: utf-8; indent-tabs-mode: nil -*- */

/***************************************************************
 *
 * Copyright (c) 2015, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

#ifndef _GONLPIR_H_
#define _GONLPIR_H_

/**
 * 
 *
 * @file gonlpir.h
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sun Mar  8 16:27:57 2015
 *
 **/

#define POS_MAP_NUMBER 4 //add by qp 2008.11.25
#define ICT_POS_MAP_FIRST 1  //计算所一级标注集
#define ICT_POS_MAP_SECOND 0 //计算所二级标注集
#define PKU_POS_MAP_SECOND 2 //北大二级标注集
#define PKU_POS_MAP_FIRST 3//北大一级标注集
#define POS_SIZE 40

typedef struct Result {
  int start; //start position,词语在输入句子中的开始位置
  int length; //length,词语的长度
  char  sPOS[POS_SIZE];//word type，词性ID值，可以快速的获取词性表
  int iPOS;//词性标注的编号
  int word_ID; //该词的内部ID号，如果是未登录词，设成0或者-1
  int word_type; //区分用户词典;1，是用户词典中的词；0，非用户词典中的词
  int weight;//word weight,read weight
} result_t;

#define GBK_CODE 0//默认支持GBK编码
#define UTF8_CODE GBK_CODE+1//UTF8编码
#define BIG5_CODE GBK_CODE+2//BIG5编码
#define GBK_FANTI_CODE GBK_CODE+3//GBK编码，里面包含繁体字
#define UTF8_FANTI_CODE GBK_CODE+4//UTF8编码

int NLPIR_Init(const char* data_path, int encode, const char* lincence);
int NLPIR_Exit();
const char* NLPIR_ParagraphProcess(const char *para, int pos_tagged);
const result_t* NLPIR_ParagraphProcessA(const char *para, int *result_count,
                                        int use_user_dict);
unsigned int NLPIR_ImportUserDict(const char *filename, int overwrite);

#endif /* _GONLPIR_H_ */

/* vim: set expandtab shiftwidth=2 tabstop=2: */
