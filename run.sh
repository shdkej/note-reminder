#!/bin/bash

# Get Note File
cd ~/workspace/reminder/parsing
go test

# Update CSV File
    # parsing.go makeCSV()

# Get Content
cd ~/workspace/reminder
python3 recommender_system.py > recommend.txt

# Get Tagline
    # parsing.go getTagline(recommend.txt) return tag, tagline, file

# Make index.html file
### Render {Title, Content, File} -> <a href="File">Title</a><p>Content</p>
    # parsing.go > index.html

### Render {Title, Content, File} -> <a href="static-site">Title</a>
# Send Telegram
# sqs.go sendSQS(recommend.txt)
cd ~/workspace/reminder/note-aws-manager
go test

# Send S3
    # lambda.go uploadS3()

# Telegram -> Web
