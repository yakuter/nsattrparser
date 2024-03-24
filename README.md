# nsattrparser: A Go Package for Parsing NSAttributedString Binary Data
The attrparser package provides a straightforward and efficient solution for extracting text from binary-encoded NSAttributedString data. Designed for Go developers, it simplifies the process of parsing complex binary formats that originate from Apple's Foundation framework, specifically those used in text styling and formatting within iOS and macOS applications.

## Key Features
Binary Pattern Recognition: Utilizes predefined start and end binary patterns to identify the relevant sections of data for extraction, making it highly effective at parsing structured binary text data.
UTF-8 Validation and Correction: Employs UTF-8 validation to ensure the extracted text adheres to UTF-8 standards, with built-in mechanisms for handling and correcting non-UTF-8 compliant data.
Simplicity and Efficiency: Designed with simplicity in mind, attrparser allows for easy integration into projects, providing a high-performance parsing solution without the overhead of external dependencies or complex configurations.

## Use Cases
**Data Extraction:** Ideal for developers needing to extract and manipulate text from binary-encoded NSAttributedString objects, especially when dealing with cross-platform data interchange between iOS/macOS applications and Go-based systems.
**Content Migration:** Facilitates the migration of rich text content from Apple ecosystem applications into other platforms or formats, preserving text content while omitting proprietary styling information.
**Debugging and Analysis:** Useful in debugging scenarios where binary NSAttributedString data needs to be inspected or analyzed, providing clear and immediate access to the encapsulated text.

## Getting Started
To use the attrparser package in your Go project, simply install and use it as follows:


``` go
go get github.com/yakuter/nsattrparser
```

```go
...

import "github.com/yakuter/nsattrparser"

...

func main() {
	hexData := "040b73747265616d747970656481e803840140848484194e534d757461626c6541747472696275746564537472696e67008484124e5341747472696275746564537472696e67008484084e534f626a6563740085928484840f4e534d757461626c65537472696e67018484084e53537472696e67019584012b29576861742061626f757420746865206c617374207461736b2077652074616c6b65642061626f75743f86840269490129928484840c4e5344696374696f6e617279009584016901928498981d5f5f6b494d4d657373616765506172744174747269627574654e616d658692848484084e534e756d626572008484074e5356616c7565009584012a849b9b00868686"

	// Convert the hexadecimal data to a byte slice
	binaryData, err := hex.DecodeString(hexData)
	if err != nil {
		fmt.Printf("Error decoding hex string: %v\n", err)
		return
	}

	// Parse the binary data
	parsedText, parseErr := parse(binaryData)
	if parseErr != nil {
		fmt.Printf("Error parsing binary data: %v\n", parseErr)
		return
	}

	fmt.Printf("Parsed text: %s\n", parsedText)
    // OUTPUT: Parsed text: What about the last task we talked about?
}

```



## Contributing
Contributions to **nsattrparser** are welcome! Whether it's extending functionality, improving efficiency, or enhancing documentation, your input can help make **nsattrparser** even better for the Go community.