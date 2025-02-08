curl -s -X GET 'http://localhost:8080/articles/feed' \
-H 'Content-Type: application/json' | jq -r

curl -s -X GET 'http://localhost:8080/articles' \
-H 'Content-Type: application/json' | jq -r

curl -s -X POST 'http://localhost:8080/articles' \
-H 'Content-Type: application/json' | jq -r

curl -s -X GET 'http://localhost:8080/articles/abc' \
-H 'Content-Type: application/json' | jq -r

curl -s -X PUT 'http://localhost:8080/articles/abc' \
-H 'Content-Type: application/json' | jq -r

curl -s -X DELETE 'http://localhost:8080/articles/abc' \
-H 'Content-Type: application/json' | jq -r