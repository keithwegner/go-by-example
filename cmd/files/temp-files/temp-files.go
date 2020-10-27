package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Throughout program execution, we often want to create data that isn't needed after the program exits. Temporary files
// and directories are useful for this purpose since they don't pollute the file system over time.

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// The easiest way to create a temporary file is by calling ioutil.TempFile. It creates a file and opens it for
	// reading and writing. We provide "" as the first argument, so it creates the file in the default location for
	// our OS.
	f, err := ioutil.TempFile("", "sample")
	check(err)

	// Displays the name of the temp file. On Unix based OS'es, the dir will likely be /tmp. The filename starts with
	// the prefix given as the second argument to ioutil.TempFile, and the rest is chosen automatically to ensure that
	// concurrent calls will always create different file names.
	fmt.Println("Temp file name:", f.Name())

	// Clean up the file after we're done. The OS is likely to clean up temp files by itself after some time, but it's
	// good practice to do this explicitly.
	defer os.Remove(f.Name())

	// We can write data to the file.
	_, err = f.Write([]byte{1, 2, 4, 5})
	check(err)

	// If we intend to wrtie many temp files, we may prefer to create a temp dir. ioutil.TempDir's arguments are the
	// same as TempFile's, but it returns a directory name rather than an open file.
	dname, err := ioutil.TempDir("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	// Now we can synthesize temp file names by prefixing them with our temp dir
	fname := filepath.Join(dname, "file1")
	err = ioutil.WriteFile(fname, []byte{1, 2, 3, 4}, 0666)
	check(err)

}
