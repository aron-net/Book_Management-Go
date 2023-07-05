# Book Management Go

> This is a simple book management system implemented in Go. It provides basic CRUD (Create, Read, Update, Delete) operations for books.

## Features

- Get a list of all book
- Get details of a specific book
- Add a new book
- Update an existing book
- Delete a book

## Getting Started

> To run the application locally, follow these steps:

1. Make sure you have Go installed on your system. You can download it from the official website:
```shell
 https://golang.org/
```

2. Clone this repository to your local machine or download the source code:
```shel
https://github.com/aron-helu/Book_Management-Go.git
```

3. Open a terminal or command prompt and navigate to the project's root directory:
```shell
cd Book_Management-Go
```


4. Set up the MySQL database and update the database connection details in the code if necessary.

5. Run the following command to start the API server:
```shell
go run main.go
```
The server will start running on http://localhost:8000.

## API Endpoints
## Get a list of all books
- URL: ``` /book ```
- Method:  ``` GET ```
- Description: ``` Retrieves a list of all books```.
- Response: ``` Returns a JSON array containing book objects.```

## Get a book

- URL: ```/book/:id ```
- Method: ``` GET ```
- Description: ``` Retrieves details of a specific book```
- Parameters:
  - ``` id (string) ``` The ID of the book.
- Response: ```Returns a JSON object representing the book```

## Add a book
- URL: ```/book```
- Method: ``` POST ```
- Description: ``` Adds a new book to the collection. ```
- Request Body: ``` Expects a JSON object representing the book.```
- Response: ``` Returns the newly created book as a JSON object.```

## Update a book
- URL: ```/book/:id```
- Method: ``` PUT ```
- Description: ``` Updates an existing book.```
- Parameters:
  - ```id (string): ``` The ID of the book.
- Request Body: ``` Expects a JSON object representing the updated book.```
- Response: ``` Returns the updated book as a JSON object. ``` 

## Delete a book
- URL: ``` /book/:id ```
- Method: ``` DELETE ```
- Description: ``` Deletes a book from the collection.```
- Parameters:
  - ```id (string): ``` The ID of the book.
- Response: ```Returns the updated book collection as a JSON array.```

## Dependencies

> This project uses the following external dependencies:

- httprouter - A lightweight HTTP request router.
- gorm - A powerful ORM library for Go.
- MySQL dialect for GORM - MySQL database driver for GORM.
## Example Usage
> Here's an example of how you can use cURL to interact with the book API:

1. Get all books:

```shell
curl http://localhost:8000/book
```
2. Get a specific book:

```shell
curl http://localhost:8000/book/1
```
3. Add a new book:

```shell
curl -X POST -H "Content-Type: application/json" -d '{
  "isbn": "123456",
  "title": "New book",
  "director": {
    "firstname": "Jane",
    "lastname": "Doe"
  }
}' http://localhost:8000/book
```

4. Update a book:

```shell
curl -X PUT -H "Content-Type: application/json" -d '{
  "isbn": "654321",
  "title": "Updated book",
  "director": {
    "firstname": "John",
    "lastname": "Smith"
  }
}' http://localhost:8000/book/1
```

5. Delete a book:

```shell
curl -X DELETE http://localhost:8000/book/1
```
> Feel free to explore and interact with the book API using different HTTP clients or tools.

## Authors

👤 **Aaron Abraham**

- GitHub: [@Aaron](https://github.com/aron-helu)

- LinkedIn: [@Aaron](https://www.linkedin.com/in/aron-abraham-90a4321b0/)


## 🤝 Contributing

Contributions, issues, and feature requests are welcome!

Feel free to check the [issues page](../../issues/).



## Show your support

Give a ⭐️ if you like this project!


## 📝 License

This project is [MIT](./LICENSE) licensed.
