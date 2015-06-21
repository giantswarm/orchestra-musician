#!/bin/sh
echo starting jackd in network mode
jackd -R -d net &

sleep 2 

echo starting orchestra-musician to wait for conductor
/musician/musician $@
