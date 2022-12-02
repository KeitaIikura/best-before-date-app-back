package apperr

import "net/http"

// code
const (
	ErrCodeInvalidRequest TmxErrorCode = "001" // リクエスト形式などが不正なエラー
	ErrCodeUnauthorized   TmxErrorCode = "002" // 認証失敗エラー
	ErrCodeTargetNotFound TmxErrorCode = "003" // 処理対象が存在しない（関連データが存在しない、かつ存在しないものの区別が不要な場合なども含む）

	ErrCodeDBConnection        TmxErrorCode = "101" // 区別不要なDBエラー全般
	ErrCodeInvalidArgs         TmxErrorCode = "201" // バッチ引数エラー
	ErrCodeInternalServerError TmxErrorCode = "500" // 一意に定義できないエラー全般
)

// type
const (
	ErrTypeBadRequest         TmxErrorType = "BAD_REQUEST"
	ErrTypeUnauthorized       TmxErrorType = "UNAUTHORIZED"
	ErrTypeTargetNotFound     TmxErrorType = "TARGET_NOT_FOUND"
	ErrTypeDBConnection       TmxErrorType = "DB_ERROR"
	ErrTypeInternaServerError TmxErrorType = "SYSTEM_ERROR"
)

// message
const (
	ErrMsgInvalidRequest      TmxErrorMessage = "不正なリクエストです。"
	ErrMsgUnauthorized        TmxErrorMessage = "ログインに失敗しました。"
	ErrMsgTargetNotFound      TmxErrorMessage = "処理対象が存在しませんでした。"
	ErrMsgDBConnetion         TmxErrorMessage = "DB処理でエラーが発生しました。"
	ErrMsgInternalServerError TmxErrorMessage = "何らかのエラーが発生しました。"
	ErrMsgInvalidArgs         TmxErrorMessage = "引数の指定が正しくありません。"
)

// response message
const (
	ErrResMsgApiInvalidRequest      TmxErrorAPIMessage = "不正な入力値です。"
	ErrResMsgApiUnauthorized        TmxErrorAPIMessage = "ログイン情報が異なります。"
	ErrResMsgApiTargetNotFound      TmxErrorAPIMessage = "処理対象が見つかりませんでした。"
	ErrResMsgApiInternalServerError TmxErrorAPIMessage = "何らかのエラーが発生しました。"
	ErrResMsgBatchMessage           TmxErrorAPIMessage = "" // バッチ用 表示されないのでブランチ
)

// Statusコード参考 (https://golang.org/pkg/net/http/#pkg-constants)
var errorMap = map[TmxErrorCode]TmxError{
	ErrCodeInvalidRequest:      {ErrorType: ErrTypeBadRequest, ErrorMessage: ErrMsgInvalidRequest, ResponseMessage: ErrResMsgApiInvalidRequest, StatusCode: http.StatusBadRequest},
	ErrCodeUnauthorized:        {ErrorType: ErrTypeUnauthorized, ErrorMessage: ErrMsgUnauthorized, ResponseMessage: ErrResMsgApiUnauthorized, StatusCode: http.StatusUnauthorized},
	ErrCodeTargetNotFound:      {ErrorType: ErrTypeTargetNotFound, ErrorMessage: ErrMsgTargetNotFound, ResponseMessage: ErrResMsgApiTargetNotFound, StatusCode: http.StatusNotFound},
	ErrCodeDBConnection:        {ErrorType: ErrTypeDBConnection, ErrorMessage: ErrMsgDBConnetion, ResponseMessage: ErrResMsgApiInternalServerError, StatusCode: http.StatusInternalServerError},
	ErrCodeInternalServerError: {ErrorType: ErrTypeInternaServerError, ErrorMessage: ErrMsgInternalServerError, ResponseMessage: ErrResMsgApiInternalServerError, StatusCode: http.StatusInternalServerError},
	// バッチ系 （responseとstatuscodeは固定
	ErrCodeInvalidArgs: {ErrorType: ErrTypeBadRequest, ErrorMessage: ErrMsgInvalidArgs, ResponseMessage: ErrResMsgBatchMessage, StatusCode: http.StatusInternalServerError},
}
