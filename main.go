package sysacademymoduleproject

import "fmt"

type MaxDate func(int64) bool

func mockDataUnits() [20]string {
	baseUnits := [20]string{
		"unit1", "unit2", "unit3", "unit4", "unit5",
		"unit6", "unit7", "unit8", "unit9", "unit10",
		"unit11", "unit12", "unit13", "unit14", "unit15",
		"unit16", "unit17", "unit18", "unit19", "unit20",
	}
	return baseUnits
}

func getStatus(date int64, maxDate MaxDate) bool {
	var result bool
	if maxDate(date) {
		result = true
	} else {
		result = false
	}

	return result
}

func getUnitNames(unitId ...int) []string {
	baseUnits := mockDataUnits()
	var unitNames []string
	for i := 0; i < len(unitId); i++ {
		var id int = unitId[i]
		unitNames = append(unitNames, baseUnits[id])
	}
	return unitNames
}

func GetDataBuyer(Name string, BornDate int64) (bool, int, []int) {
	var (
		Discrepancy int
		unitId      []int
	)

	if Name == "Patrick" {
		Discrepancy = 100000000
		unitId = []int{2, 3, 5}
	} else if Name == "Mas Septa" {
		Discrepancy = 250000000
		unitId = []int{6, 8, 10, 12}
	} else {
		Discrepancy = 10000
		unitId = []int{1, 9, 16, 19}
	}

	const maximumDate int64 = 1635875790373
	maxDate := func(date int64) bool {
		return date <= maximumDate
	}

	ActiveStatus := getStatus(BornDate, maxDate)
	return ActiveStatus, Discrepancy, unitId
}

type DataBuyer struct {
	Name         string
	ActiveStatus bool
	Discrepancy  int64
	UnitId       []int
}

func (buyer DataBuyer) PurchaseUnits() ([]string, int64, int) {
	unitNames := getUnitNames(buyer.UnitId...)

	basePrices := struct {
		unitGanjil int64
		unitGenap  int64
	}{
		unitGanjil: 15700000000000,
		unitGenap:  26800000000000,
	}

	var (
		totalPrice int64
		metaData   int
	)

	for _, unitId := range buyer.UnitId {
		if unitId%2 == 0 {
			totalPrice += basePrices.unitGenap
		} else {
			totalPrice += basePrices.unitGanjil
		}

		metaData++
	}

	totalPrice -= int64(buyer.Discrepancy)
	if totalPrice < 0 {
		totalPrice = 0
	}

	return unitNames, totalPrice, metaData
}

func (buyer DataBuyer) UnitEligibilityCheck() (bool, string) {
	unitEligible := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	var errMessage string
	var err bool

	var passCheck int = 0

	for _, unitValid := range unitEligible {
		for _, unit := range buyer.UnitId {
			if unit == unitValid-1 {
				passCheck++
			}
		}
	}

	if passCheck == len(buyer.UnitId) {
		err = false
		errMessage = "All unit are eligible"
	} else {
		err = true
		errMessage = "Check your units!"
	}

	return err, errMessage
}

type Error struct {
	error   bool
	code    int8
	message string
	data    []int
	meta    int8
}

func HandleError(error bool, code int8, message string, data []int, meta int8) string {
	if message == "" {
		return "Message must be filled"
	}

	if len(data) == 0 {
		data[0] = 0
	}

	newError := Error{
		error:   error,
		code:    code,
		message: message,
		data:    data,
		meta:    meta,
	}

	fmt.Printf(
		"code: %v, error: %v, message: %v, data: %v, meta: %v \n",
		newError.code, newError.error, newError.message, newError.data, newError.meta,
	)

	return "Success"
}

func PrintReceipt(metaData int32, totalPrice int64, unitNames ...string) {
	fmt.Printf("purchase units: %v (total: %v units) with a total price of %v successfully \n", unitNames, metaData, totalPrice)
}
