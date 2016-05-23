#!/bin/bash

tests:make-tmp-dir "ice"
tests:make-tmp-dir "sweet"

tests:make-tmp-dir "sweet/kingdom"
tests:ensure touch "sweet/kingdom/princess"
tests:ensure touch "sweet/gunter"
tests:ensure touch "ice/gunter"

tests:ensure treetrunks -n "ice" "sweet"
tests:assert-no-diff "stdout" <<EXPECTED
$(tests:get-tmp-dir)/sweet/kingdom/princess
$(tests:get-tmp-dir)/sweet/kingdom/
EXPECTED

tests:assert-stderr ""

tests:assert-test -d "sweet/kingdom"
tests:assert-test -e "sweet/gunter"
