package stringutil

const AnExportBlankString = " "
const anotherBlankString = "\t"

func Reverse(src string) string {
	r := []rune(src)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
