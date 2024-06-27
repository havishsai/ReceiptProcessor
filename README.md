# ReceiptProcessor

This is a backend project built using Go Language. Follow the below steps to successfully run the project.

1) #### Clone the project with below command:
``` bash
git clone https://github.com/havishsai/ReceiptProcessor.git
```

2) #### Go to project space:
```bash
cd ReceiptProcess
```

3) #### Now here there is Dockerfile and we build the docker container using the below command:
``` bash
docker build -t receiptprocessor .
```

4) #### Now we will the docker container that is been just created with previous command with below command:
``` bash
docker run -p 8080:8080 receipt-processor
```

5) #### Now the server is up and running. We can test the API's with API Testing tools like postman. Lets see how it is going to be if run the api with below sample data.

- ##### POST /receipts/process
  we will test this api route with url "localhost:8080/receipts/process" with POST Method.
  The response will be as below:
  ``` JSON
      {"id":"267be270-c910-4e47-bd10-fee42c4e3f54"}
  ```
- ##### GET /receipts/{id}/points
  We will test this api route with url "localhost:8080/receipts/267be270-c910-4e47-bd10-fee42c4e3f54/points" with GET Method:
  The response will be as below:
  ``` JSON
      {"points":90}
  ```

  If there is no Receipt found with given "id" then the response will be:
  ``` JSON
  Receipt not found.
  ```
  
