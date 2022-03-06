package envporter

import (
	"github.com/darkjinnee/go-err"
	"os"
)

func Overwrite(file *os.File, content string) {
	err := file.Truncate(0)
	goerr.Fatal(
		err,
		"[Error] envporter.WriteToFile: Failed to truncate file",
	)

	_, err = file.Seek(0, 0)
	goerr.Fatal(
		err,
		"[Error] envporter.WriteToFile: Failed to offset file",
	)

	_, err = file.WriteString("\n" + content)
	goerr.Fatal(
		err,
		"[Error] envporter.WriteToFile: Failed to write file",
	)
}
