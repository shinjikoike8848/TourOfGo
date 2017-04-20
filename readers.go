package main

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (mr MyReader) Read(buffer []byte) (int, error) {
	asciiCharacterCode := []rune("A")[0]
	for _, i := range buffer {
		buffer[i] = byte(asciiCharacterCode)
	}
	return len(buffer), nil
}
