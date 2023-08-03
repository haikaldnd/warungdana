package usecase

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"warungdana/models"
	"warungdana/presenter"
	"warungdana/repository"
	"warungdana/utils/constant"
)

func Number2(db *sql.DB) presenter.DefaultResponse {
	var (
		res presenter.DefaultResponse
	)
	existData, err := repository.Number2(db)
	if err != nil {
		res.Status.Message = constant.FailedGetData
		return res
	}
	res.Data = existData
	res.Status.Message = constant.SuccessGetData
	return res
}

func Number3(db *sql.DB) presenter.DefaultResponse {
	var (
		res presenter.DefaultResponse
	)
	existData, err := repository.Number3(db)
	if err != nil {
		res.Status.Message = constant.FailedGetData
		return res
	}
	res.Data = existData
	res.Status.Message = constant.SuccessGetData
	return res
}

func Number4(db *sql.DB) presenter.DefaultResponse {
	var (
		res presenter.DefaultResponse
	)
	result, err := repository.Number4(db)
	if err != nil {
		res.Status.Message = constant.FailedGetData
		return res
	}
	res.Data = fmt.Sprintf(`Selisih hari yaitu %s Hari`, result)
	res.Status.Message = constant.SuccessGetData
	return res
}

func Number5(db *sql.DB) presenter.DefaultResponse {
	var (
		res presenter.DefaultResponse
	)
	result, err := repository.Number5(db)
	if err != nil {
		res.Status.Message = constant.FailedGetData
		return res
	}
	res.Data = result
	res.Status.Message = constant.SuccessGetData
	return res
}

func Number6(db *sql.DB, data *models.File) presenter.DefaultResponse {
	var (
		res presenter.DefaultResponse
	)

	jsonData, err := json.MarshalIndent(data.Data, "", "    ")
	if err != nil {
		return res
	}
	result := os.WriteFile(data.FileName, jsonData, 0644)

	res.Status.Message = constant.SuccessWriteFile
	res.Data = result
	return res
}

func Number7(db *sql.DB, data *models.FileOpen) presenter.DefaultResponse {
	var (
		res  presenter.DefaultResponse
		temp models.File
	)
	fileData, err := os.ReadFile(data.FileName)
	if err != nil {
		res.Status.Message = constant.FailedOpenFile
		return res
	}
	err = json.Unmarshal(fileData, &temp.Data)
	if err != nil {
		res.Status.Message = constant.FailedOpenFile
		return res
	}
	res.Status.Message = constant.SuccessOpenFile
	res.Data = temp.Data
	return res
}

func Number8(db *sql.DB, data *models.InputCity) presenter.DefaultResponse {
	var (
		res presenter.DefaultResponse
	)
	exist, suggest := repository.Number8(db, data)
	if exist {
		res.Data = "true "
		res.Status.Message = constant.SuccessGetData
	} else {
		if suggest != constant.DataNotFound {
			res.Data = "false, Saran Kota " + suggest
			res.Status.Message = constant.SuccessGetData
		} else {
			res.Data = "false, " + constant.DataNotFound
			res.Status.Message = constant.DataNotFound
		}
		return res
	}
	return res
}

func Number9(db *sql.DB) presenter.DefaultResponse {
	var (
		res presenter.DefaultResponse
	)
	result, err := repository.Number9(db)
	if err != nil {
		res.Status.Message = constant.FailedGetData
		return res
	}
	res.Data = result
	return res
}

func Number9c(db *sql.DB, data *models.InputMulti) presenter.DefaultResponse {
	var (
		res presenter.DefaultResponse
	)
	temp := convertStringToArray(data.Data)
	result, err := repository.Number9c(db, temp)
	if err != nil {
		res.Status.Message = constant.FailedGetData
		return res
	}
	res.Data = result
	return res
}

func Number9d(db *sql.DB, data *models.InputMulti) presenter.DefaultResponse {
	var (
		res presenter.DefaultResponse
	)
	temp, err := strconv.Atoi(data.Data)

	if err != nil {
		res.Status.Message = constant.FailedConvertData
		return res
	}

	result, err := repository.Number9d(db, temp)
	if err != nil {
		res.Status.Message = constant.FailedGetData
		return res
	}
	res.Data = result
	return res
}

func Number10a(db *sql.DB) presenter.DefaultResponse {
	var (
		res presenter.DefaultResponse
	)

	result, err := repository.Number10a(db)
	if err != nil {
		res.Status.Message = constant.FailedGetData
		return res
	}
	res.Data = result
	res.Status.Message = constant.SuccessGetData
	return res
}

func convertStringToArray(input string) []int {
	stringArray := strings.Split(input, ",")
	intArray := make([]int, len(stringArray))

	for i, str := range stringArray {
		num, _ := strconv.Atoi(strings.TrimSpace(str))
		intArray[i] = num
	}

	return intArray
}
