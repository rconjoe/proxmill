# shellcheck shell=bash


gh_token=$(curl -s -H "Authorization: Bearer $WM_TOKEN" \
  "$BASE_INTERNAL_URL/api/w/$WM_WORKSPACE/variables/get_value/f/github/access_token" | jq -r .)

git clone https://rconjoe:$token@github.com/rconjoe/proxmill.git
