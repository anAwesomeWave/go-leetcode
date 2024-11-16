package main

// import "fmt"

func compress(chars []byte) int {
	writePos := 0
	cnt := 1

	for i := 1; i < len(chars); i++ {
		// fmt.Println(i, chars[i], cnt)
		if chars[i] == chars[i-1] {
			cnt += 1
		} else {
			chars[writePos] = chars[i-1]
			writePos += 1
			if cnt > 1 {
				// идея. идти справа налево, потом отреверсить
				startingPos := writePos
				for cnt > 0 {
					var newChar byte = byte(cnt % 10)
					chars[writePos] = newChar + '0'
					cnt /= 10
					writePos++
				}
				for j := 0; j < (writePos-startingPos)/2; j++ {
					chars[startingPos+j], chars[writePos-j-1] = chars[writePos-j-1], chars[startingPos+j]
				}
			}
			cnt = 1
		}
	}
	chars[writePos] = chars[len(chars)-1]
	writePos += 1
	if cnt > 1 {
		// идея. идти справа налево, потом отреверсить
		startingPos := writePos
		for cnt > 0 {
			var newChar byte = byte(cnt % 10)
			chars[writePos] = newChar + '0'
			cnt /= 10
			writePos++
		}
		for j := 0; j < (writePos-startingPos)/2; j++ {
			chars[startingPos+j], chars[writePos-j-1] = chars[writePos-j-1], chars[startingPos+j]
		}
	}
	return writePos
}
