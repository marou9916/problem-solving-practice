#!/usr/bin/bash
curl https://zone01normandie.org/assets/superhero/all.json | jq '.[] | select(.id == 70) | .name'
