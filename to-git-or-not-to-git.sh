curl https://zone01normandie.org/assets/superhero/all.json | jq '.[] | select(.id == 170) | .name, .powerstats.power, .appearance.gender'
