package common

import "github.com/gin-gonic/gin"

func CheckErrorReturn(ctx *gin.Context,err error,msg string) {
	if err != nil {
		if msg != "" {
			Failed(ctx,msg)
		} else {
			Failed(ctx,err.Error())
		}
	}
	return
}

func CheckErrorReturnErr(err error) error {
	if err != nil {
		return err
	}
	return nil
}