#!/bin/bash

if [ "$1" == "-v" ]; then
	regex="(found [1-9][0-9]* occur|mutating \S* to|failed to break any tests)"
else
	regex="(failed to break any tests)"
fi

manbearpig -list | while read -r line; do
	manbearpig -import github.com/assertgo/assert -mutation "$line" | egrep "$regex"
done
