curl -s -X GET 'http://localhost:8080/articles/feed' \
-H 'Content-Type: application/json' | jq -r

curl -s -X GET 'http://localhost:8080/articles' \
-H 'Content-Type: application/json' | jq -r

curl -s -X POST 'http://localhost:8080/articles' \
-H 'Content-Type: application/json' \
-d '{
    "article": {
        "title": "Hello",
        "description": "World",
        "body": "Hello World!",
        "tagList": ["a", "b", "c"]
    }
}' | jq -r

curl -s -X GET 'http://localhost:8080/articles/abc' \
-H 'Content-Type: application/json' | jq -r

curl -s -X PUT 'http://localhost:8080/articles/abc' \
-H 'Content-Type: application/json' | jq -r

curl -s -X DELETE 'http://localhost:8080/articles/abc' \
-H 'Content-Type: application/json' | jq -r