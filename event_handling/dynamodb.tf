resource "aws_dynamodb_table" "blog" {
  provider = aws

  hash_key          = "Tag"
  name              = "myBlog"
  stream_enabled    = true
  stream_view_type  = "NEW_AND_OLD_IMAGES"
  read_capacity     = 20
  write_capacity    = 20

  attribute {
    name = "Tag"
    type = "S"
  }
}

resource "aws_dynamodb_global_table" "myBlogTable" {
  depends_on = [aws_dynamodb_table.blog]
  provider = aws

  name = "myBlog"

  replica {
    region_name = var.region
  }
}
