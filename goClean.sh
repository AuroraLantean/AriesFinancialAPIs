#!/bin/bash
#chmod u+x this_script.sh
echo "inside script ..."
echo "..."

FILE=./logs.txt
if [[ -f "$FILE" ]]; then
    echo "$FILE exists"
    rm logs.txt
    echo "$FILE has been deleted"
else 
    echo "$FILE does not exist."
fi
