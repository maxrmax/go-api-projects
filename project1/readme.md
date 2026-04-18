#### Short mockup GO API for research and studying

run with
`go run cmd/api/main.go`

Test the get request with `curl`

`curl "http://localhost:8000/account/coins/?username=Jotaro" -H "Authorization: 123456"`

`curl "http://localhost:8000/account/coins/?username=Dio" -H "Authorization: 666FFF"`

`curl "http://localhost:8000/account/coins/?username=Zeppelin" -H "Authorization: ZEPZEP"`


and wrong request (400) with wrong username or wrong token:

`http://localhost:8000/account/coins/?username=any`

| Username      | Auth Token    |
| ------------- | ------------- |
| Jotaro		| 123456		|
| Dio			| 666FFF		|
| Zeppelin		| ZEPZEP		|