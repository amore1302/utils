//файл utils.go
//GO-03 03  курсы reabrain учебное задание по созданию пакета

package utils

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
