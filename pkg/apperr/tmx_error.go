package apperr

import (
	"bbdate/pkg/logging"
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

type (
	TmxErrorCode       string // 内部エラーコード
	TmxErrorType       string // エラータイプ
	TmxErrorMessage    string // 内部的なエラーメッセージ
	TmxErrorAPIMessage string // api返却時用のメッセージ
)

type TmxError struct {
	ErrorType       TmxErrorType
	ErrorMessage    TmxErrorMessage
	ResponseMessage TmxErrorAPIMessage
	StatusCode      int
	ErrorDetail     error // もともとのエラー
	ErrorCaller     string
	ResponseBody    map[string]interface{}
}

func (e TmxError) Error() string {
	return fmt.Sprintf("%v %s:%v", string(e.ErrorMessage), e.ErrorCaller, e.ErrorDetail)
}

// public methods
func NewTmxError(code TmxErrorCode, orgErr error) TmxError {
	e := errorMap[code]
	e.ErrorCaller = errorCaller()
	e.ResponseBody = map[string]interface{}{
		"code":    string(e.ErrorType),
		"message": string(e.ResponseMessage),
	}
	e.ErrorDetail = orgErr
	// 既にTmxError型でラップ済みの場合はもともとのエラーを格納する
	if oErr, ok := orgErr.(TmxError); ok {
		e.ErrorDetail = oErr.ErrorDetail
	}
	return e
}

// （主にcontroller層での）TmxErrorへの型キャストのためのメソッド
func ConvertToTmxError(xrid string, err error) TmxError {
	te, ok := err.(TmxError)
	if ok {
		return te
	}
	// 通常発生しない想定なのでエラーログとして出力する
	logging.Error(xrid, fmt.Sprintf("ConverToTmxError unexpect error: %v", err))
	// unexpected error
	return NewTmxError(ErrCodeInternalServerError, err)

}

// private methods

func errorCaller() string {
	pc, file, line, _ := runtime.Caller(2)
	filename := file[strings.LastIndex(file, "/")+1:] + ":" + strconv.Itoa(line)
	funcname := runtime.FuncForPC(pc).Name()
	fn := funcname[strings.LastIndex(funcname, ".")+1:]
	fileName := filename
	functionName := fn
	return fmt.Sprintf("%s(%s)", fileName, functionName)
}
