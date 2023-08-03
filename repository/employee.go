package repository

import (
	"database/sql"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
	"warungdana/models"
	"warungdana/utils/constant"
)

func Number2(db *sql.DB) ([]models.Employee, error) {
	var (
		result []models.Employee
	)

	rows, err := db.Query("SELECT Employee.FirstName, Employee.Lastname, Employee.ID FROM Employee WHERE Employee.LastName LIKE ? ORDER BY Employee.LastName,Employee.FirstName", "Smith%")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		data := models.Employee{}
		rows.Scan(&data.FirstName, &data.LastName, &data.ID)
		result = append(result, data)
	}
	return result, nil
}

func Number3(db *sql.DB) ([]models.Employee, error) {
	var (
		result []models.Employee
	)
	rows, err := db.Query("select Employee.FirstName,Employee.LastName,Employee.HireDate from Employee left join AnnualReviews AR on Employee.ID = AR.EmpID where AR.EmpID IS NULL ORDER BY Employee.HireDate")

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		data := models.Employee{}
		rows.Scan(&data.FirstName, &data.LastName, &data.HireDate)
		result = append(result, data)
	}
	return result, nil
}

func Number4(db *sql.DB) (string, error) {
	var (
		days string
	)
	rows, err := db.Query("select DATEDIFF(MAX(e.TerminationDate),MIN(e.HireDate)) as days from Employee e")

	if err != nil {
		return "", err
	}

	defer rows.Close()
	for rows.Next() {
		rows.Scan(&days)
	}

	return days, nil
}

func Number5(db *sql.DB) ([]models.NewEmployee, error) {
	var (
		result []models.NewEmployee
	)
	rows, err := db.Query("SELECT k.FirstName , k.LastName , k.Salary  AS gaji_awal, ROUND(k.salary * POWER(1.15, (YEAR('2016-01-01') - YEAR(k.HireDate))), 2) AS perkiraan_gaji_2016, COUNT(ar.EmpID) AS total_ulasan FROM Employee k LEFT JOIN AnnualReviews ar ON k.id = ar.EmpID GROUP BY k.ID , k.FirstName , k.LastName , k.Salary , k.HireDate ORDER BY perkiraan_gaji_2016 DESC, total_ulasan ASC")

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		data := models.NewEmployee{}
		datab := models.AnnualReviews{}
		rows.Scan(&data.FirstName, &data.LastName, &data.Salary, &data.PerkiraanGaji, &datab.ID)
		data.TotalUlasan = datab.ID
		result = append(result, data)
	}

	return result, nil
}

func Number6(db *sql.DB) ([]models.NewEmployee, error) {
	var (
		result []models.NewEmployee
	)
	rows, err := db.Query("SELECT k.FirstName , k.LastName , k.Salary  AS gaji_awal, ROUND(k.salary * POWER(1.15, (YEAR('2016-01-01') - YEAR(k.HireDate))), 2) AS perkiraan_gaji_2016, COUNT(ar.EmpID) AS total_ulasan FROM Employee k LEFT JOIN AnnualReviews ar ON k.id = ar.EmpID GROUP BY k.ID , k.FirstName , k.LastName , k.Salary , k.HireDate ORDER BY perkiraan_gaji_2016 DESC, total_ulasan ASC")

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		data := models.NewEmployee{}
		datab := models.AnnualReviews{}
		rows.Scan(&data.FirstName, &data.LastName, &data.Salary, &data.PerkiraanGaji, &datab.ID)
		data.TotalUlasan = datab.ID
		result = append(result, data)
	}

	return result, nil
}

func Number8(db *sql.DB, data *models.InputCity) (bool, string) {
	var (
		cities = []string{"Bandung", "Cimahi", "Ambon", "Jayapura", "Makasar"}
	)
	for _, city := range cities {
		if data.City == city {
			return true, ""
		}
	}
	suggestionByFirstLetter := []string{}
	suggestionByLastLetter := []string{}
	for _, city := range cities {
		if strings.HasPrefix(city, string(data.City[0])) {
			suggestionByFirstLetter = append(suggestionByFirstLetter, city)
		}
		if strings.HasSuffix(city, string(data.City[len(data.City)-1])) {
			suggestionByLastLetter = append(suggestionByLastLetter, city)
		}
	}
	sort.Strings(suggestionByFirstLetter)
	sort.Strings(suggestionByLastLetter)

	suggest := ""
	//suggest from first
	if len(suggestionByFirstLetter) > 0 {
		suggest = fmt.Sprintf("%s", strings.Join(suggestionByFirstLetter, ", "))
	}
	//suggest from last
	if len(suggestionByLastLetter) > 0 {
		if suggest != "" {
			suggest += " , "
		}
		suggest += fmt.Sprintf("%s", strings.Join(suggestionByLastLetter, ", "))
	}

	if suggest == "" {
		suggest = constant.DataNotFound
	}

	return false, suggest
}

func Number9(db *sql.DB) (models.TempDataArray, error) {
	var (
		result          models.TempDataArray
		resultSort      []int
		resultDuplicate []string
		arr             = []int{9, 1, 6, 4, 8, 6, 6, 3, 8, 2, 3, 3, 4, 1, 8, 2}
	)
	//sort data
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		if i == 0 || arr[i] != arr[i-1] {
			resultSort = append(resultSort, arr[i])
		}
	}

	//total duplicate
	countMap := make(map[int]int)
	for _, num := range arr {
		countMap[num]++
	}
	for num, count := range countMap {
		output := fmt.Sprintf("%d[%d] ", num, count)
		resultDuplicate = append(resultDuplicate, output)

	}
	result.SortArray = resultSort
	result.TotalDuolicate = resultDuplicate
	return result, nil
}

func Number9c(db *sql.DB, list []int) (models.TempDataArray, error) {
	var (
		arr        = []int{9, 1, 6, 4, 8, 6, 6, 3, 8, 2, 3, 3, 4, 1, 8, 2}
		tempresult = []int{}
		result     models.TempDataArray
	)
	for _, num := range arr {
		shouldRemove := false
		for _, value := range list {
			if num == value {
				shouldRemove = true
				break
			}
		}

		if !shouldRemove {
			tempresult = append(tempresult, num)
		}
	}
	result.DeleteValueFromInput = tempresult
	return result, nil
}

func Number9d(db *sql.DB, input int) ([]int, error) {
	var (
		arr = []int{9, 1, 6, 4, 8, 6, 6, 3, 8, 2, 3, 3, 4, 1, 8, 2}
	)
	result := make([]int, len(arr))

	remaining := input
	maxValue := 10

	for i, num := range arr {
		for i := 0; i < remaining; i++ {
			fmt.Print("A", num, remaining, maxValue)
			if num == maxValue {
				break
			} else {
				num += 1
				remaining--
				fmt.Print("B", num, remaining, maxValue)
			}
		}
		result[i] = num
	}
	return result, nil
}

func Number10a(db *sql.DB) (map[string]interface{}, error) {
	const (
		letters       = "abcdefghijklmnopqrstuvwxyz"
		vowels        = "aeiou"
		digits        = "0123456789"
		randomLength  = 100
		lettersLength = 50
		digitsLength  = 50
	)

	var (
		data map[string]interface{}
	)
	//generate letter
	rand.Seed(time.Now().UnixNano())
	tempLetter := make([]byte, lettersLength)
	for i := range tempLetter {
		tempLetter[i] = letters[rand.Intn(len(letters))]
	}
	//geenrate digit
	rand.Seed(time.Now().UnixNano())
	tempDigit := make([]byte, digitsLength)
	for i := range tempDigit {
		tempDigit[i] = digits[rand.Intn(len(digits))]
	}
	//hitung total vokal
	countTotalVowels := 0
	for _, char := range tempLetter {
		if contains(vowels, string(char)) {
			countTotalVowels++
		}
	}
	//hitung total digit
	countTotalDigit := 0
	for _, char := range tempDigit {
		if contains(digits, string(char)) {
			countTotalDigit++
		}
	}
	//hitung angka even
	countTotalEvenDigit := 0
	for _, char := range tempDigit {
		digit := string(char)
		if contains(digits, digit) {
			num, _ := strconv.Atoi(digit)
			if num%2 == 0 {
				countTotalEvenDigit++
			}
		}
	}
	sortAngkaTerbesarTerkecil := sortAndRemoveDuplicatesDigit(string(tempDigit))
	sortHurufTerkecilTerbesar := sortAndRemoveDuplicates(string(tempLetter))

	var sortedResult []string
	for i := 0; i < len(sortAngkaTerbesarTerkecil) && i < len(sortHurufTerkecilTerbesar); i++ {
		sortedResult = append(sortedResult, fmt.Sprintf("%s%s", sortAngkaTerbesarTerkecil[i], sortHurufTerkecilTerbesar[i]))
	}

	// Append remaining numbers or letters if any
	if len(sortAngkaTerbesarTerkecil) > len(sortHurufTerkecilTerbesar) {
		for i := len(sortHurufTerkecilTerbesar); i < len(sortAngkaTerbesarTerkecil); i++ {
			sortedResult = append(sortedResult, fmt.Sprintf("%s", sortAngkaTerbesarTerkecil[i]))
		}
	} else if len(sortHurufTerkecilTerbesar) > len(sortAngkaTerbesarTerkecil) {
		for i := len(sortAngkaTerbesarTerkecil); i < len(sortHurufTerkecilTerbesar); i++ {
			sortedResult = append(sortedResult, sortHurufTerkecilTerbesar[i])
		}
	}
	dataTemp := map[string]interface{}{
		"Total Letters":              len(tempLetter),
		"Total Vokal":                len(vowels),
		"Total Digit":                len(tempDigit),
		"Total Even Digit":           countTotalEvenDigit,
		"Huruf Terkecil Ke Terbesar": sortAndRemoveDuplicates(string(tempLetter)),
		"Angka terbesar ke terkecil": sortAndRemoveDuplicatesDigit(string(tempDigit)),
		"Hasil Generate Random":      sortedResult,
	}
	data = dataTemp
	return data, nil
}

func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

func sortAndRemoveDuplicates(str string) []string {
	splitted := strings.Split(str, "")
	sort.Strings(splitted)

	var result []string
	for i := 0; i < len(splitted)-1; i++ {
		if splitted[i] != splitted[i+1] {
			result = append(result, splitted[i])
		}
	}
	result = append(result, splitted[len(splitted)-1])
	return result
}
func sortAndRemoveDuplicatesDigit(str string) []string {
	splitAngka := strings.Split(str, "")
	intAngka := make([]int, len(splitAngka))
	for i, s := range splitAngka {
		intAngka[i], _ = strconv.Atoi(s)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(intAngka)))
	uniqueAngka := make(map[int]bool)
	var sortedAngka []string
	for _, num := range intAngka {
		if _, found := uniqueAngka[num]; !found {
			uniqueAngka[num] = true
			sortedAngka = append(sortedAngka, strconv.Itoa(num))
		}
	}
	return sortedAngka
}
