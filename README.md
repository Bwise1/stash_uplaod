# File Upload Server

This is a simple file upload server written in Go. It allows users to upload files to the server via HTTP PUT requests, and it provides a download link for each uploaded file.

## Usage

To upload a file, make a PUT request to the server with the file in the request body. The URL of the request should be the desired file name. For example, to upload a file named "example.txt", you could use the following `curl` command:

```bash
curl localhost:8080 -T hello.txt
```

### Response

After the file is uploaded, the server will respond with a message like this:

```
File uploaded successfully
Download with: wget http://localhost:8080/downlod/hello-234343.txt
```
