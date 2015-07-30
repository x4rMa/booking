#!/bin/sh


export DEBFULLNAME="Benjamin Borbe"
export EMAIL=bborbe@rocketnews.de

export DEB_SERVER=misc.rn.benjamin-borbe.de
export TARGET_DIR=opt/booking/bin

export NAME=booking
export BINS="booking_server"
export INSTALLS="github.com/bborbe/booking/bin/booking_server"
export SOURCEDIRECTORY="github.com/bborbe/booking"

export MAJOR=0
export MINOR=1
export BUGFIX=0

# exec
sh src/github.com/bborbe/jenkins/jenkins.sh