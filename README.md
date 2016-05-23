<img
    src="http://pre01.deviantart.net/3e9d/th/pre/f/2013/143/3/3/tree_trunks___adventure_time_by_nanaruko-d66cx4s.jpg"
    align="right"
    width="250px"
/>

# treetrunks

Little tool, that will remove all files and directories from specified
directory, which are not exist in source directory.

It will not compare contents of the files, only existence matters.

treetrunks outputs list of deleted files and directories on the standard
output.

It will also provide dry run mode, in which treetrunks will print exactly the
same output how it will print in real delete mode, but will not delete any
files.

# Usage

```
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
```
