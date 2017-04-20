package main

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (mr MyReader) Read(buffer []byte) (int, error) {
	asciiCharacterCodeOfA := byte('A')
	for i := range buffer {
		buffer[i] = asciiCharacterCodeOfA
	}
	return len(buffer), nil
}
