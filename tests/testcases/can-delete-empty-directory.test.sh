#!/bin/bash

tests:make-tmp-dir "ice"
tests:make-tmp-dir "sweet"

tests:make-tmp-dir "sweet/kingdom"

tests:ensure treetrunks "ice" "sweet"
tests:assert-stdout "$(tests:get-tmp-dir)/sweet/kingdom"
tests:assert-stderr ""

tests:assert-test ! -d "sweet/kingdom"
