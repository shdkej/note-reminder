#!/bin/bash

# Get Note File

docker-compose up -d

cd ~/workspace/note-reminder/recommender
cp ../tags.csv ./
docker build -t recommender .
docker run --rm recommender

# Update CSV File

# Get Content

# Get Tagline
    # parsing.go getTagline(recommend.txt) return tag, tagline, file

# Make index.html file
### Render {Title, Content, File} -> <a href="File">Title</a><p>Content</p>
    # parsing.go > index.html

### Render {Title, Content, File} -> <a href="static-site">Title</a>
# Send Telegram

# Send S3

# Telegram -> Web
