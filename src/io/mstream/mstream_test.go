package mstream

import "testing"

/*
To run tests run below command
$ go test -v ./src/io/mstream/.../
*/

func TestMinioClient(t *testing.T) {
	result := MinioTest()
	if result > 0 {
		t.Fatalf("Minio test failed")
	}
}

func TestCopyfile(t *testing.T) {
	result := CopyFiles()
	if result > 0 {
		t.Fatalf("Copyfile test failed")
	}
}
