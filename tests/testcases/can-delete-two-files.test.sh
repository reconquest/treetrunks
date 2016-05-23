#!/bin/bash

tests:make-tmp-dir "ice"
tests:make-tmp-dir "sweet"

tests:ensure touch "sweet/pancake"
tests:ensure touch "sweet/rootbeer"

tests:ensure treetrunks "ice" "sweet"
tests:assert-stdout "$(tests:get-tmp-dir)/sweet/pancake"
tests:assert-stdout "$(tests:get-tmp-dir)/sweet/rootbeer"
tests:assert-stderr ""

tests:assert-test ! -e "sweet/pancake"
tests:assert-test ! -e "sweet/rootbeer"
