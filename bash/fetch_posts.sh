#!/bin/bash

# Parameters
SOURCE="https://api.vk.com/method/wall.get"
TOKEN="$(cat ~/.vk/token)"
COMMUNITY="-161290464"
# Date and time (GMT): Sunday, 1 January 2023 г., 0:00:00
UNTIL=1672531200
IS_UNTIL_REACHED=false

# File for results
OUTPUT=$1
if [[ $OUTPUT == "" ]]; then
    OUTPUT="./output_posts.txt"
fi

# Remove already existing one, start from the begining
rm -f $OUTPUT

echo "Fetching process has started."
echo "Source: $SOURCE"
echo "Community: $COMMUNITY"
echo "UNTIL: $UNTIL"

OFFSET=0
MAX_FETCH_COUNT=100
while [[ $IS_UNTIL_REACHED == false ]]
do
    PARAMS="owner_id=$COMMUNITY&offset=$OFFSET&count=$MAX_FETCH_COUNT&extended=1&access_token=$TOKEN&v=5.199"

    # Get array, filter by until, save only array without response wrapper
    ARRAY="$(curl -s -X POST $SOURCE --data $PARAMS | jq -c ".response.items" | jq -c "map(. |  select(.date > $UNTIL))")"

    # Sleep before next API call
    sleep .5

    echo "Fetched next queue of size $MAX_FETCH_COUNT with offset $OFFSET"

    # Sum up next queue
    OFFSET=$(( $OFFSET + 100 ))

    # Loop ending condition
    if [[ $ARRAY == "[]" ]] || [[ $ARRAY == "" ]]; then
        IS_UNTIL_REACHED=true
    else
        # Save intermediate result
        echo "$ARRAY" >> $OUTPUT
    fi
done

echo "Fetching process has been completed."