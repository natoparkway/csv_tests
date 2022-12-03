#!/usr/bin/env bash

fn=$1

output=$(/usr/bin/time -lp $fn 2>&1)
while IFS= read -r line; do
	if [[ "$line" == "real"* ]]; then
		time_elapsed=${line//[!0-9\.]/}
	elif [[ "$line" == *"peak memory footprint" ]]; then
		peak_memory_bytes=${line//[!0-9]/}
	fi
done <<< "${output}"

echo $output

echo "Time Elapsed: ${time_elapsed} second(s)"
peak_memory_mbs=$(echo "$peak_memory_bytes / (1024 * 1024)" | bc -l)
echo "Memory Used" ${peak_memory_mbs}"mbs"
