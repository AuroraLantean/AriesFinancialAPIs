#!/bin/bash
#chmod u+x this_script.sh
echo "inside addEnvToZip.sh ..."
echo "..."

zip env.zip .env config.yml
echo "confirm the zip file include the .env and config.yml file!"
