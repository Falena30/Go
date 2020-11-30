package models

//BinarySearch adalah algoritma pencarian dasar
func BinarySearch(arr []int, l int, r int, x int) int {

	if r >= l {
		mid := l + (r-l)/2

		if arr[mid] == x {
			return mid
		} else if arr[mid] > x {
			return BinarySearch(arr, l, mid-1, x)
		}

		return BinarySearch(arr, mid+1, r, x)
	}
	return -1
}

//NormalFilter adalah fungsi untuk mengetahui true atau false untuk string
func NormalFilter(data []string, callback func(string) bool) []string {
	var result []string

	//ingat ini adalah deklarasi variable dan dimasukkan range data
	for _, each := range data {
		//untuk memisahkan kondisi gunakan ';'
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}
