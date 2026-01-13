package file

import (
	"io"
	"os"

	"github.com/azuyamat/hermit/internal/types"
)

type Cat struct{}

func (c *Cat) Name() string {
	return "cat"
}

func (c *Cat) Execute(args []string, context *types.ExecutionContext, stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	if len(args) == 0 {
		_, err := io.Copy(stdout, stdin)
		return err
	}
	for _, filename := range args {
		file, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(stdout, file)
		if err != nil {
			return err
		}
	}
	return nil
}
