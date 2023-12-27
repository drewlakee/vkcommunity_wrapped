#!/bin/bash

# Parameters
SOURCE="https://api.vk.com/method/likes.getList"
TOKEN="$(cat ~/.vk/token)"
COMMUNITY="-161290464"

# Input file with post ids
INPUT=$1

OUTPUT_DIR=$2
if [[ $OUTPUT_DIR == "" ]]; then
    OUTPUT_DIR="./likes"
fi

# Remove already existing one, start from the begining
rm -fr $OUTPUT_DIR
mkdir $OUTPUT_DIR

echo "Fetching process has started."
echo "Source: $SOURCE"
echo "Community: $COMMUNITY"
echo "Posts: $INPUT"
echo "Output: $OUTPUT_DIR"

COUNTER=0
PRINT_EVERY=10
while IFS= read -r postid
do
    PARAMS="owner_id=$COMMUNITY&type=post&item_id=$postid&extended=0&access_token=$TOKEN&v=5.199"
    ARRAY="$(curl -s -X POST $SOURCE --data $PARAMS | jq -c ".response.items")"
    FILENAME="$OUTPUT_DIR/$postid.json"
    
    # Remove if already exists
    rm -f $FILENAME

    if [[ $ARRAY != "[]" ]] && [[ $ARRAY != "" ]] && [[ $ARRAY != "null" ]]; then
        echo "$ARRAY" >> $FILENAME
        echo "Fetched post=$postid likes to $FILENAME"
    else
        echo "⚠️⚠️⚠️ Post $postid. Result $ARRAY"
    fi
    
    # Sleep before next call
    sleep .5
    
    COUNTER=$(($COUNTER + 1))
    if [ $(($COUNTER % $PRINT_EVERY)) -eq 0 ]; then
    echo "Already fetched $COUNTER posts"
    fi
done < "$INPUT"

echo "Fetching process has been completed."