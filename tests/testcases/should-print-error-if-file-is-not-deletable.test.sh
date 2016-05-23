#!/bin/bash

tests:make-tmp-dir "ice"
tests:make-tmp-dir "sweet"

tests:make-tmp-dir "sweet/kingdom"
tests:ensure touch "sweet/kingdom/princess"
tests:ensure chmod -rwx "sweet/kingdom"

tests:not tests:ensure treetrunks "ice" "sweet"

tests:assert-stderr-re "sweet/kingdom.*permission denied"

tests:ensure chmod +rwx "sweet/kingdom"

tests:assert-test -e "sweet/kingdom/princess"

