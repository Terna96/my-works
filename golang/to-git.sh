curl -s https://acad.learn2earn.org.ng/assets/superhero/all.json | jq -r '.[] | select(.id == 180) | "\(.name)\n\(.powerstats.power)\n\(.appearance.gender)"'
