# Lambda URL Shortener
This project is a serverless URL shortener built using AWS services. It allows users to shorten long URLs and retrieve them using a custom short link.

## Features
- Convert long URLs to short URLs.
- Redirect to long URLs by using custom short link.
- NoSQL database on AWS serverless.
- Support HTTPS (ACM - AWS Certificate Manager)
- Customer domains. (Configured via Route53 and ACM)
- Secure POST method when create short URLs. (Optional)

## AWS Services Used
- IAM – For access control and security roles.
- Route53 – For custom domain management. (Can use other DNS service)
- ACM (AWS Certificate Manager) – For HTTPS support.
- DynamoDB – NoSQL database to store short URLs.
- Lambda – Serverless compute function to handle URL processing.
- API Gateway – Exposes the API for URL shortening and redirection.

## API
- API `POST` for create short URL. (If setup API key must add `x-api-key` in Headers)
    - Body
    ```JSON
    {
        "long_url": "https://google.com"
    }
    ```
    -Response
    ```JSON
    {
        "short_url": "https://api.domain.com/{short_code}"
    }
    ```
- API `GET` to redirect short URL to long URL. `https://api.domain.com/{short_code}`

## Must have!
- Registered domain.

## Deployment Steps (Manual)
1. Setup ACM - AWS Certificate Manager for HTTPS.
    - Request public certificate.
    ![Image](https://github.com/user-attachments/assets/6ef61117-7406-4573-bb6a-1a9008b8ccab)
    - Get CNAME and CNAME value to verify owner.
    ![Image](https://github.com/user-attachments/assets/3981cb7e-37cb-4303-8b60-31002c6739a3)

2. Setup Route 53.
    - Setup domain record for verify owner.
    ![Image](https://github.com/user-attachments/assets/e6828346-a49d-4ccb-b4ed-a36f86ad6f29)

3. Wait for certificates status to be Issued.
![Image](https://github.com/user-attachments/assets/fd68004e-dcd7-4918-95f4-76fbe5fa0f8b)

4. Setup DynamoDB Table.
    - Table name: `ShortURLs` &rarr; Will use in `Lambda` environment .
    - Primary key: `short_code` (String).
![Image](https://github.com/user-attachments/assets/803fe171-2eac-4651-9d01-5d67948bb005)
![Image](https://github.com/user-attachments/assets/b4a74187-9096-4997-9107-d4ce3424d1bc)

5. Setup API Gateway.
    - Create API > REST API (Must use proxy for redirect to long URLs)
    ![Image](https://github.com/user-attachments/assets/150c0a2d-8192-460e-b0e5-8a90c5a98b28)

6. Setup code.
    - Ubuntu case.
        - Install go. (https://go.dev/doc/install)
        - Git clone project & cd to project.
        - `go mod tidy`
        - `CGO_ENABLED=0 go build -o bootstrap cmd/main.go`
        - `zip main.zip bootstrap`
        - You will get `main.zip` for upload in Lambda function.

7. Setup Lambda Function.
    - Creat Lambda function.
    ![Image](https://github.com/user-attachments/assets/2a7ee734-c9d1-4b0c-86b5-b3e2c262587a)
    - Upload `main.zip`.
    - Add Trigger `API Gateway`
    ![Image](https://github.com/user-attachments/assets/43224c7f-9f26-49ea-bbdd-469bb4b502de)
    - Select existing API.
    ![Image](https://github.com/user-attachments/assets/c07c5249-64c2-49a7-bcf8-02d1d7168de2)
    - Add `DynamoDB` table name `ShortURLs` in environment variables.
    ![Image](https://github.com/user-attachments/assets/2a2ef005-87ba-46eb-9c96-4bbe3b1c96aa)
    - Set permission in `IAM` to let `Lambda` can read and write `DynamoDB` by finding Lambda role name
    ![Image](https://github.com/user-attachments/assets/75ad401e-58f9-4c93-8279-f4ee493f47a3)
    ![Image](https://github.com/user-attachments/assets/cda17624-bb11-46ec-bb45-fd16c274a5fb)

8. Setup API path.
    - Create resource for `short_code` proxy path.
    ![Image](https://github.com/user-attachments/assets/ae76acef-7474-47fb-94d3-82413b22c696)
    - Create resource POST for `shorten` path.
    ![Image](https://github.com/user-attachments/assets/1d6d4c44-8dd4-484e-8af2-ecaf8ffc6c94)
    - Don't forget to Deploy API after finish setup path.
    ![Image](https://github.com/user-attachments/assets/4fdf6985-9900-4b54-93f3-9e11ce6562b1)
    - Create custom domain names.
    ![Image](https://github.com/user-attachments/assets/a13f4e73-4fe6-423c-bbca-6aab8228a855)
    - Select API mappings for use custom domain
    ![Image](https://github.com/user-attachments/assets/2e3ccf6a-56d0-4c8e-97b1-93658bbdf62e)

9. Setup custom domain for API
    - Get API Gateway domain name for `CNAME VALUE`
    ![Image](https://github.com/user-attachments/assets/3e395693-593f-4eed-b7b1-4fb6fce740e2)
    - Record name: `api.domain.com` and VALUE: `CNAME VALUE`
    ![Image](https://github.com/user-attachments/assets/e6828346-a49d-4ccb-b4ed-a36f86ad6f29)

10. (Optional) Create API key for POST method
    - Create usage plan for `customer domain name`
    ![Image](https://github.com/user-attachments/assets/b3a87bce-f17c-47ae-baff-ef7695625208)
    - Create API keys.
    ![Image](https://github.com/user-attachments/assets/fc1e91f6-08e9-45bf-811a-cf72d1a0b1bd)
    - Add usage plan to API keys.
    ![Image](https://github.com/user-attachments/assets/9f519209-bef8-4fb2-abcb-afa7e62ebd28)
    - Edit /shorten POST to use API key
    ![Image](https://github.com/user-attachments/assets/15274a9d-2a87-4bc7-9cce-3b95765adab3)
    ![Image](https://github.com/user-attachments/assets/2bf7832e-349b-4ea8-b137-01ed09deacb9)