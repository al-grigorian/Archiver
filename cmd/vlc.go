package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var vlcCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Pack file using variable-length code",
	Run:   pack,
}

const packedExtension = "vlc" // имя расширения

var ErrEmptyPath = errors.New("path to file is not specified")

func pack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}
	filePath := args[0]         // путь к файлу получаем из списка аргументов, переданных в командой строке
	r, err := os.Open(filePath) // для открытия файла
	if err != nil {
		handleErr(err)
	}

	data, err := io.ReadAll(r)
	if err != nil {
		handleErr(err)
	}

	// packed := Encode(data)
	packed := "" // результат сжатия
	fmt.Println(string(data))

	// packedFileName - имя сжатого файла,  byte(packed) - содержимое файла, 0644- права
	err = os.WriteFile(packedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handleErr(err)
	}
}

// формирует имя сжатого файла c форматом vlc
func packedFileName(path string) string {

	// path = /path/to/file/myFile.txt -> myFile.vlc

	fileName := filepath.Base(path) // получаем имя файла 'myFile.txt'
	// получаем только расширение '.txt'
	// получаем имя файла без расширения 'myFile'

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packedExtension
}

func init() {
	packCmd.AddCommand(vlcCmd)
}
