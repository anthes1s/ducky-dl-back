#!/bin/sh

echo "starting cleanup..."

cd /app/uploads || exit

current_time=$(date +%s)

for file in *; do
    if [ -f "$file" ]; then
        file_time=$(stat -c %Z "$file")
        age=$((current_time - file_time))

        # '1800' is the amount of seconds that the file exists on the server
        if [ "$age" -gt 1800 ]; then
            rm "$file"
        fi
    fi
done
