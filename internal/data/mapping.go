/*
 * @FilePath: /workflow-server/internal/data/mapping.go
 * @Author: maggot-code
 * @Date: 2023-08-27 22:43:36
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-08-28 00:52:06
 * @Description:
 */
package data

import (
	"errors"
	"math"
	"strconv"
)

// 整数位和小数位对照表
var (
	IntMapping = []float64{
		0.03,
		0.10,
		0.20,
		0.41,
		0.62,
		0.83,
		1.04,
		1.25,
		1.46,
		1.67,
		1.88,
		2.09,
		2.30,
		2.51,
		2.72,
		2.93,
		3.14,
		3.35,
		3.56,
		3.77,
		3.98,
		4.19,
		4.40,
		4.61,
		4.82,
		5.03,
		5.24,
		5.45,
		5.66,
		5.87,
		6.07,
		6.28,
		6.50,
		6.71,
		6.92,
		7.13,
		7.34,
		7.55,
		7.76,
		7.97,
		8.18,
		8.39,
		8.60,
		8.81,
		9.02,
		9.23,
		9.44,
		9.65,
		9.86,
		10.00,
	}

	FloatMapping = []float64{
		0.02,
		0.04,
		0.06,
		0.08,
		0.10,
		0.12,
		0.14,
		0.16,
		0.18,
	}
)

func RecordToInt(record float64) (float64, error) {
	index := int(math.Floor(record))

	if index < 0 || index > len(IntMapping) {
		return 0, errors.New("record to int mapping error")
	}

	if index-1 >= 0 {
		return IntMapping[index-1], nil
	}

	return 0, nil
}

func RecordToFloat(record float64) (float64, error) {
	decimal := record - float64(int(math.Floor(record)))
	decimalStr := strconv.FormatFloat(float64(decimal), 'f', 2, 32)
	index, err := strconv.Atoi(decimalStr[2:3])
	if err != nil || index < 0 || index > len(FloatMapping) {
		return 0, errors.New("record to float mapping error")
	}

	if index-1 >= 0 {
		return FloatMapping[index-1], nil
	}

	return 0, nil
}
