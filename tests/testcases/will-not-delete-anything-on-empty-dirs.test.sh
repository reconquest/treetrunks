#!/bin/bash

tests:make-tmp-dir "ice"
tests:make-tmp-dir "sweet"

tests:ensure treetrunks "ice" "sweet"
tests:assert-stdout ""
tests:assert-stderr ""

tests:assert-test -d "ice"
tests:assert-test -d "sweet"
