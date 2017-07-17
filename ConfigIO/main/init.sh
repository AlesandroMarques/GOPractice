#!/bin/bash
basedir="$(pwd)"
ssh-keygen -f "${basedir}/host_key" -P "mdmubu107.torolab.ibm.com"
ssh-keygen -f "${basedir}/user_key" -P "ws8admin"


