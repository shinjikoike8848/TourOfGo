package main

// Pic return bitmap dx * dy
func Pic(dx, dy int) [][]uint8 {
	var bitMap [][]uint8
	for y := 0; y < dy; y++ {
		bitArray := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			bitArray[x] = uint8(x * y)
		}
		bitMap = append(bitMap, bitArray)
	}
	return bitMap
}
