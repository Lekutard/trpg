package usecases

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

//UError Usecase Error
type UError struct {
	Code int
	Msg  string
}

const (
	//BadRequest リクエスト不正
	BadRequest = iota + 1
	//Unauthorized 認証不正
	Unauthorized
	//NotFound 存在しない
	NotFound
	//Conflict 競合
	Conflict
	//Internalexception その他内部エラー
	Internalexception
)

//Error error interfaceを実装
func (err *UError) Error() string {
	return fmt.Sprintf("ERROR: %s", err.Msg)
}

//GetUError Usecase Error
func GetUError(code int, msg string) *UError {
	uerr := &UError{
		Code: code,
		Msg:  msg,
	}

	return uerr
}

//GetUErrorByError Get Usecase Error By Error
func GetUErrorByError(err error) *UError {
	if err != nil {
		var code int
		if gorm.IsRecordNotFoundError(err) {
			code = NotFound
		} else if isDuplicatedUError(err.Error()) {
			code = Conflict
		} else {
			code = Internalexception
		}

		return GetUError(code, err.Error())
	}
	return nil
}

//IsRecordNotFoundUError Error is NotFound
func (err *UError) IsRecordNotFoundUError() bool {
	return err.Code == NotFound
}

func isDuplicatedUError(msg string) bool {
	return strings.Index(msg, "Duplicate") != -1
}
