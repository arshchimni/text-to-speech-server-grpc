#!/bin/bash

go run main.go "$1"

afplay output.wav
 
