#!/bin/bash
# Requires lftp + aria2
FILENAME=$(lftp -c "open 5.44.137.84/ESStatistikListeModtag && ls" | tail -n 1 | awk '{print $NF}')
DOWNLOAD_PATH="ftp://5.44.137.84/ESStatistikListeModtag/${FILENAME}"

echo "Downloading from ${DOWNLOAD_PATH}"

aria2c -x 16 -s 16 "$DOWNLOAD_PATH" -o data.zip