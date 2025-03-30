package main

import (
	"os"
	"crypto/rand"
)

const overwriteNum = 3

// Generates a random alphanumeric of given length n
func GenerateRandomData(size int) []byte {
	s := make([]byte, size)
	rand.Read(s)
	return s
}


// overwrite the given file 3 times withrandom data and delete the file afterwards.
func Shred(path string) error {
	fileinfo, err := os.Stat(path)
	size := fileinfo.Size()

	for i := 0; i < overwriteNum; i++ {
    	// overwrite the file content
		var randomData = GenerateRandomData(int(size))
		err := os.WriteFile(path, randomData, 0666)
		if (err != nil) {
			return err
		}
  	}

	// delete the file
	err = os.Remove(path)
	if (err != nil) {
		return err
	} 
	return nil
}


func main() {
	var path = os.Args[1]
	err := Shred(path)
	if (err != nil) {
		panic(err)
	}
}
