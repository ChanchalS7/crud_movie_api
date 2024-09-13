# Movie API

This is a simple RESTful API implemented in Go to manage a list of movies. It supports operations to create, read, update, and delete movies.

## Endpoints

### `GET /movies`

Fetches the list of all movies.

**Response:**
- Status Code: 200 OK
- Content-Type: application/json
- Body: Array of movie objects

### `GET /movies/{id}`

Fetches a single movie by its ID.

**Response:**
- Status Code: 200 OK
- Content-Type: application/json
- Body: Movie object

### `POST /movies`

Creates a new movie. The request body should be a JSON object with the following fields:
- `isbn` (string): The ISBN of the movie
- `title` (string): The title of the movie
- `director` (object): The director of the movie, which should include `firstname` and `lastname`

**Response:**
- Status Code: 200 OK
- Content-Type: application/json
- Body: The created movie object with a newly generated ID

### `PUT /movies/{id}`

Updates an existing movie by its ID. The request body should be a JSON object with the same structure as `POST /movies`.

**Response:**
- Status Code: 200 OK
- Content-Type: application/json
- Body: The updated movie object

### `DELETE /movies/{id}`

Deletes a movie by its ID.

**Response:**
- Status Code: 200 OK
- Content-Type: application/json
- Body: Array of remaining movies

## Running the API

1. Clone the repository:
   ```bash
   git clone <https://github.com/ChanchalS7/crud_movie_api.git>

2. Navigate to the project directory
   ```bash
   cd <project-directory>

3. Install the dependencies:
	```bash
	 go mod tidy

4. Run the application:
	```bash
	go run main.go

```The server will start on http://localhost:8080.```


## Dependencies

github.com/gorilla/mux - Router for handling HTTP requests.
encoding/json - For encoding and decoding JSON data.
math/rand - For generating random IDs.
net/http - For HTTP server and request handling.
strconv - For converting numbers to strings.
log - For logging errors and information.




## Endpoints

### 1. `GET /movies`

Fetches the list of all movies.

- **Method:** GET
- **URL:** `http://localhost:8080/movies`
- **Headers:** None
- **Body:** None

**Response:**
- **Status Code:** 200 OK
- **Content-Type:** application/json
- **Body:** JSON array of movie objects.

### 2. `GET /movies/{id}`

Fetches a single movie by its ID.

- **Method:** GET
- **URL:** `http://localhost:8080/movies/{id}` (replace `{id}` with the actual movie ID)
- **Headers:** None
- **Body:** None

**Response:**
- **Status Code:** 200 OK
- **Content-Type:** application/json
- **Body:** JSON object of the requested movie.

### 3. `POST /movies`

Creates a new movie.

- **Method:** POST
- **URL:** `http://localhost:8080/movies`
- **Headers:**
  - `Content-Type: application/json`
- **Body:** (Select `raw` and `JSON` format)
  ```json
  {
    "isbn": "1234567890",
    "title": "New Movie",
    "director": {
      "firstname": "Jane",
      "lastname": "Doe"
    }
  }


### 4. PUT /movies/{id}
Updates an existing movie by its ID.

Method: PUT
URL: http://localhost:8080/movies/{id} (replace {id} with the actual movie ID)
Headers:
Content-Type: application/json
Body: (Select raw and JSON format

```
{
  "isbn": "0987654321",
  "title": "Updated Movie",
  "director": {
    "firstname": "John",
    "lastname": "Smith"
  }
}
```

Response:

Status Code: 200 OK
Content-Type: application/json
Body: JSON object of the updated movie.
Example Use:

URL: http://localhost:8080/movies/1 (replace 1 with the ID of the movie you want to update)
Method: PUT
Body :
```
{
  "isbn": "0987654321",
  "title": "Updated Movie",
  "director": {
    "firstname": "John",
    "lastname": "Smith"
  }
}
```
Click "Send" to update the movie with the specified ID.


### 5. DELETE /movies/{id}
Deletes a movie by its ID.

Method: DELETE
URL: http://localhost:8080/movies/{id} (replace {id} with the actual movie ID)
Headers: None
Body: None
Response:

Status Code: 200 OK
Content-Type: application/json
Body: JSON array of remaining movies.