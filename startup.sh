#!/bin/sh
echo starting jackd in network mode
jackd -R -d net &

sleep 5

echo "starting supercollider"
scsynth -u 4555 -t 4555
