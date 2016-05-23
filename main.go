package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/docopt/docopt-go"
	"github.com/kovetskiy/lorg"
	"github.com/seletskiy/hierr"
)

const version = "1.0"

const usage = `treetrunks - remove files from target dir that are not in source

treetrunks will search target directory for the files and directories that are
not exists in the source directory, and then will delete them recursively.

treetrunks can be run in dry-run mode optionally.

treetrunks will print removed files and directories on the standard output.

Every printed directory will end with single slash (/).

Usage:
    treetrunks -h | --help
    treetrunks [-n] <source> <target>

Arguments:
    <source>   Source directory to compare with.
    <target>   Target directory to delete from.

Options:
    -h --help     Show this help.
    -n --dry-run  Run in dry mode, do not delete anything, just print.
`

func main() {
	args, err := docopt.Parse(usage, nil, true, "treetrunks "+version, false)
	if err != nil {
		panic(err)
	}

	var (
		log = lorg.NewLog()

		sourceDir = args["<source>"].(string)
		targetDir = args["<target>"].(string)

		dryRun = args["--dry-run"].(bool)
	)

	filesListToRemove, err := collectFiles(sourceDir, targetDir)
	if err != nil {
		log.Fatal(err)
	}

	err = removeFiles(filesListToRemove, dryRun)
	if err != nil {
		log.Fatal(err)
	}
}

func collectFiles(sourceDir, targetDir string) ([]string, error) {
	filesListToRemove := []string{}

	err := filepath.Walk(
		targetDir,
		func(targetPath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			sourcePath := filepath.Join(
				sourceDir,
				strings.TrimPrefix(targetPath, targetDir),
			)

			_, err = os.Stat(sourcePath)
			if err == nil {
				return nil
			}

			if !os.IsNotExist(err) {
				return hierr.Errorf(
					err,
					`can't get file stat: '%s'`,
					sourcePath,
				)
			}

			absPath, err := filepath.Abs(targetPath)
			if err != nil {
				return hierr.Errorf(
					err,
					`can't get absolute path to: '%s'`,
					targetPath,
				)
			}

			if info.IsDir() {
				absPath += "/"
			}

			filesListToRemove = append(filesListToRemove, absPath)

			return nil
		},
	)

	if err != nil {
		return filesListToRemove, hierr.Errorf(
			err,
			`can't collect files to remove`,
		)
	}

	return filesListToRemove, nil
}

func removeFiles(filesList []string, dryRun bool) error {
	removedFiles := map[string]bool{}

	for i := len(filesList) - 1; i >= 0; i-- {
		absPath := filesList[i]

		remove := false

		if strings.HasSuffix(absPath, "/") {
			empty, err := isDirEmpty(absPath, removedFiles)
			if err != nil {
				return hierr.Errorf(
					err,
					`can't determine is dir empty or not: '%s'`,
					absPath,
				)
			}

			if empty {
				remove = true
			}
		} else {
			remove = true
		}

		if remove {
			err := removeAndLog(absPath, dryRun)
			if err != nil {
				return hierr.Errorf(
					err,
					`can't remove file: '%s'`,
					absPath,
				)
			}

			removedFiles[absPath] = true
		}
	}

	return nil
}

func removeAndLog(path string, dryRun bool) error {
	if !dryRun {
		err := os.Remove(path)
		if err != nil {
			return err
		}
	}

	fmt.Println(path)

	return nil
}

func isDirEmpty(name string, removedFiles map[string]bool) (bool, error) {
	dir, err := os.Open(name)
	if err != nil {
		return false, hierr.Errorf(
			err,
			`can't open directory for listing: '%s'`,
			name,
		)
	}

	filesInDir, err := dir.Readdirnames(-1)
	if err == io.EOF {
		return true, nil
	}

	if err != nil {
		return false, hierr.Errorf(
			err,
			`can't list directory entries in: '%s'`,
			name,
		)
	}

	for _, file := range filesInDir {
		if removedFiles[filepath.Join(name, file)] {
			continue
		}

		return false, nil
	}

	return true, nil
}
