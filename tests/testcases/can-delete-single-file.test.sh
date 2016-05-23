#!/bin/bash

tests:make-tmp-dir "ice"
tests:make-tmp-dir "sweet"

tests:ensure touch "sweet/pancake"

tests:ensure treetrunks "ice" "sweet"
tests:assert-stdout "$(tests:get-tmp-dir)/sweet/pancake"
tests:assert-stderr ""

tests:assert-test ! -e "sweet/pancake"
