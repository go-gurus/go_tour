#!/bin/bash
OUTFILE=golang-for-developers.md
cd resources
rm -f $OUTFILE ; touch $OUTFILE ; find ../examples/**/SLIDES.md -exec cat {} >> $OUTFILE \; -exec echo >> $OUTFILE \;
