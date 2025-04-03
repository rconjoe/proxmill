# shellcheck shell=bash
set -e

token=$(curl -s -H "Authorization: Bearer $WM_TOKEN" \
  "$BASE_INTERNAL_URL/api/w/$WM_WORKSPACE/variables/get_value/f/github/access_token" | jq -r .)

npm i -g windmill-cli

git clone https://rconjoe:$token@github.com/rconjoe/proxmill.git
cd proxmill

git config user.name "rconjoe"
git config user.email "root@trog.codes"

wmill workspace add playground playground https://proxmill.trog.codes --token "$WM_TOKEN"

wmill sync pull --yes

date=$(date "+%Y-%m-%d %H:%M:%S")

git add .
git commit -m "sync $date"

# Push changes
git push https://rconjoe:$token@github.com/rconjoe/proxmill.git master

echo "Script finished."
