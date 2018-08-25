#!/bin/bash

ACTIVE_WINDOW=""

while true
do
    sleep 0.5

    NEW_ACTIVE_WINDOW=$(xprop -root 32x '\t$0' _NET_ACTIVE_WINDOW | cut -f 2)

    if [ "${NEW_ACTIVE_WINDOW}" != "${ACTIVE_WINDOW}" ]
    then
        ACTIVE_WINDOW=${NEW_ACTIVE_WINDOW}
        xprop -id ${ACTIVE_WINDOW} WM_CLASS WM_NAME _NET_WM_NAME

        echo
        echo "Press [CTRL+C] to stop..."
    fi

done
