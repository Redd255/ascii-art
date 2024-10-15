#!/bin/bash

echo "go run . --color red \"banana\""
echo
go run . --color red "banana"
echo
echo "---------------------------------------------------------"

echo "go run . --color=red \"hello world\""
echo
go run . --color=red "hello world"
echo
echo "---------------------------------------------------------"

echo "go run . --color=green \"1 + 1 = 2\""
echo
go run . --color=green "1 + 1 = 2"
echo
echo "---------------------------------------------------------"