# go-bookstore
The codebase for a bookstore API built with Go

# Note
Make sure you have go installed on your system and it is available in your path. Also create the database before you start the server to avoid errors.

# Setup
Clone the repo and cd into the folder, cd into cmd/main and then type 

``
go run main.go
``

# Routes
The port for the application is 8002 and the available routes are

|   Route       |   Method      |     Description         |
| ------------- |:-------------:| :---------------------: |
|   /book       |   GET         | This route fetches all the books in the database |
|   /book       |   POST        | This route is for creating new books |
|   /book/{id}  |   GET         | This route fetches a single book's details |
|   /book/{id}  |   PUT         | This route is for updating a book's details |
|   /book/{id}  |   DELETE      | This route is for deleting a book from the database |

# Database Setup
The corresponding table for the database is go_bookstore. You can create it in your mysql database using your favourite GUI.
