package main

// MyReader implemets reader interface
type MyReader struct{}

// Read return byte array of ascii code of 'A'
func (mr MyReader) Read(buffer []byte) (int, error) {
	asciiCharacterCodeOfA := byte('A')
	for i := range buffer {
		buffer[i] = asciiCharacterCodeOfA
	}
	return len(buffer), nil
}
