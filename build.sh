#!/bin/bash
OUTFILE=golang-for-developers.md
cd resources
echo "start building resources/golang-for-developers.md from examples/*/SLIDES.md"
rm -f $OUTFILE ; touch $OUTFILE ; find ../examples/**/SLIDES.md -exec cat {} >> $OUTFILE \; -exec echo >> $OUTFILE \;
echo "finished building resources/golang-for-developers.md"
