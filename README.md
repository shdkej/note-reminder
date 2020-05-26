- [ ] vpc
- [x] parsing tag
    - [x] parsing tag line
- [x] send telegram
- [x] send sqs
- [x] sqs to lambda
- [ ] sqs setting
- [ ] check sqs, lambda communication speed
- [ ] put elasticsearch
- [ ] search suggestion
- [x] golang elasticsearch
    - [ ] what is context?
- [x] make index file
- [ ] update score, weight
- [ ] get 5 list each other, in recent, keyword search, date search
    - keyword search: find recent search history data
- [ ] scaling weight by every day
- [x] cron every day
- [ ] data initial & update
- [ ] tag, tagline match conflict content check
- [ ] save tag with file
- [ ] date link cannot access link
- [ ] if search tag, show relate tag
- [ ] python recommend program, csv file to s3, 
- go-update -> csv update -> s3 upload -> lambda call(python) -> go web update and send telegram
- [x] mget err check
- [ ] make tagline parsing algorithm better
- [ ] go wasm <-> go server grpc

#### Goal
Remind My note tag list

#### Architecture
Parsing - Dynamodb - elasticsearch - lambda - telegram
and show remind list in web.
and remind content + suggestion content

#### Test
parsing - redis - elasticsearch - api - api

#### Parsing data
- Note file
- Search History

#### Crawling
- read file and store to redis
- elasticsearch get in redis

#### Mechanism
- state repetition
    - oldest update file send
    - 7-30-365 date 
- Data
    - name
    - last read
    - last update
    - weight
