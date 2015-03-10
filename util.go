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
 * @file util.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sun Mar  8 16:51:20 2015
 *
 **/

package gonlpir

//===================================================================
// Public APIs
//===================================================================

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

//===================================================================
// Private
//===================================================================
