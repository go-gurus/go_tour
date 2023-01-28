#!/bin/bash
OUTFILE=go-tour.md
cd resources
echo "start building resources/go-tour.md from slides/*/SLIDES.md"
rm -f $OUTFILE ; touch $OUTFILE ; find ../slides/**/SLIDES.md -exec cat {} >> $OUTFILE \; -exec echo >> $OUTFILE \;
echo "finished building resources/go-tour.md"
