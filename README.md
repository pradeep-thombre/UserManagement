# UserManagement
 
## Task2 (Echo)

Create a simple application using golang (echo) with an API GET /user which returns the following json :
{"user": "pradeep"}

Initialize it as a go module (using go mod init)

Deploy the application on port 3000 and verify that you are able to access http://localhost:3000/user in browser and correct response is shown

Things covered in this assignment:
- Creating a simple echo application
- Running echo application on a specific port

## Additional APIS (Pending)

- Add a POST API /users
	On calling this API, a new user should be stored in-memory. 
- Modify the GET /users method to return a list of added users from memory
- Add a GET by id method to get a specific user by id


The application should now have the following APIS:

POST /users - adds a new user in the application, with a new id
	curl -iX POST -d '{"user": "bhborkar"}' https://localhost:3000/users 
	should return {"id":<number>} and status code 204


GET /users - should return all users added
	curl -iX GET https://localhost:3000/users
	should return in this format - [{"id":1, "user":"ashwin"}, ...] (list of all users added)
	
GET /users/<id> - should return user for a specific id
	curl -iX GET https://localhost:3000/users/<number>
	should return 
		- {"id":<number>, "user":<name>} (if user with id <nymber> is added) with status code 200
		- {"error": "Not found"} with status code 404 if user is not added with id <number>


Things covered in this assignment:
- Simple map operations
- CRUD apis

