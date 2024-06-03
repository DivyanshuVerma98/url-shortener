[![GO](https://skillicons.dev/icons?i=go)](https://go.dev/) [![AWS](https://skillicons.dev/icons?i=aws)]()

# URL Shortener
This repository contains the backend code for a URL shortening service - [www.chotalink.com](https://chotalink.com)

[![App Platform](https://s3.ap-south-1.amazonaws.com/www.chotalink.com/homepage-ss.png)](https://chotalink.com)

## Features
- Generates a short URL
- Allows setting a custom alias
- Shortened URLs expire after 24 hours

## Tech Stack
- **AWS Lambda**: Hosts the backend
- **AWS API Gateway**: Triggers the Lambda functions via APIs
- **AWS DynamoDB**: Database for the application
- **GoLang**: Programming language

## Prerequisites
Basic knowledge of the following:
- AWS Lambda
- AWS API Gateway
- AWS DynamoDB
- GoLang

## Setup

1. **Create APIs in AWS API Gateway**:
    - `/createurl`
    - `/{proxy+}`

    [![API Gateway](https://s3.ap-south-1.amazonaws.com/www.chotalink.com/apigateway-ss.png)]()

2. **Create a table in AWS DynamoDB**:
    - Table Name: `UrlShortenerMapper`

    [![DynamoDB](https://s3.ap-south-1.amazonaws.com/www.chotalink.com/dynamodb-ss.png)]()

3. **Upload the code to AWS Lambda**:
    - Ensure the Lambda function has the necessary execution role with permissions to interact with DynamoDB and API Gateway.

4. **Enable API Gateway triggers for the AWS Lambda functions**:
    - Connect the API Gateway methods to the respective Lambda functions.

> **Note**: You can change the URL and table name to anything you prefer. Just make sure to update the code accordingly.

