# AWS Event Stream Processor

## What it does:

This Go program helps you process event streams generated from an AWS S3 `SelectObjectContent` operation. It provides a function named `awsEventStream` that takes the `SelectObjectContentOutput` and iterates through the received events, allowing you to process them as needed.

## Installation:

### Prerequisites:
Ensure you have Go installed on your system.

### Clone Repository:
Clone this repository using your preferred Git client:

```bash
git clone https://github.com/vishwaratna/awsEventStream_Test_coverage.git
```

Use code with caution.

## Usage:
### Import Packages: 
Import the required packages in your project:

```bash
import (
    "github.com/aws/aws-sdk-go/service/s3"
    "log"
) 
```


### Process Events:
Call the awsEventStream function with the SelectObjectContentOutput obtained from your S3 operation:

```bash
output, err := s3Client.SelectObjectContent(input)
if err != nil {
    log.Fatal(err)
}

err = awsEventStream(output)
if err != nil {
    log.Fatal(err)
}

```

### Handle Processed Events:
Within the awsEventStream function, implement your custom logic to handle each processed event.

## Testing:
This project includes unit tests to verify functionalities. To run the tests:

Navigate to the project directory containing the code.
Run the following command in your terminal:

```bash
go test
Use code with caution.
```

## Contributing:
We welcome contributions to this project! Feel free to fork the repository, make your changes, and submit a pull request.

## Additional Notes:
Remember to handle errors appropriately within your code using the provided error information.