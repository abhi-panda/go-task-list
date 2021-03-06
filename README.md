# go-task-list
 Simple Go Task list application , running with a local db (sqlite) and can be run by an adjoining shell script. Please read the [How To](https://github.com/abhi-panda/go-task-list/blob/master/README.md#how-to) Section to download and run the application.

# Architecture
The architecture is generally like the Model View Controller without the view or the Model Service Controller Pattern.

### Model
Model Contains the Blueprint of how the data element is stored , used, viewed and manipulated. Models can be found in the `model folder`.

### Controller
Controller here is in the main go file `gotasklistcontroller.go` . Controller controls the flow of every request that comes in to the server through the port.

### Service
Service is the part which actually perform the task . Here you can find service source codes in `handlers folder` which in turn call the `endpoints folder`. All the GET,POST,PUT,DELETE request get catered through these folders and packages.

### Misc parts of Architecture
```
middleware
utilities
```
The middleware folder contains source code which implements the middleware layer for the request . Panic and error conditions are handled through the middleware in this project.
utilities folder contain multiple codes which are reused in various places adn thus stored in a common package to be consumed.

# ENDPOINTS
Parent link - http://localhost:3000
This will get you the Homepage , which briefly lists all the endpoints.

add the tailing part if required

## Request 
In Head of the Request Add **Content-Type : application/json** .
In Body of the Request in json format provide :
```
{
	"TaskTitle":"",
	"DueDate":"",
	"TaskDone":false
}
```
**Note** :
1. DueDate should always be in YYYY-MM-DD formatted string.
2. TaskDone when first creating should always be false.

### GET
`/tasklist/{type}`

#### type
- all : This gets all the tasks in the Database
- alltodo : This gets all the tasks that have not been completed yet
- today : This gives you the list of tasks which are due today
- overdue : This gives you the list of tasks which are past their due date.
- <TaskTitle> : Provide a task title directly and this will return the task entity with that title

### POST (Create)
`/tasklist` : send it along with the body to create a task.

### PUT (Update)
`/tasklist` : send it along with the body to update a task.

### DELETE 
`/tasklist` : send it along with the body to delete the task.

# Unit Test
Unit test suite has been written for handlers and utilities package only as there lies the main functionality .
To run these test suite , after cloning / downloading please traverse to the required folder and run
```
go test -v
```
# Logging
Logging is an important part for every application , so keeping that in mind when the application and the CLI is run , logs are maintained for both seperately in the same folder location. 
```
server.log
shell-logs/tasklist-YYYYMMDD.logs
```
server.logs maintains the server side logs
and shell-logs folder contains logs accumulated date wise for the CLI.

# HOW TO
Please find all the information on how to download and run the application here.
If you want to look at the code please clone the repo and go :p ahead!

### Download the Executables
In the above Repository , please find the `go_task_list_executables` folder .This contains both windows and linux executables. Please get the required OS specific folder to a location of your comfort. 

### Run the Application
For now , Linux specific shellscript tasklist.sh has been created as a Command line Interface. In Future similar .psh will be created for Windows OS. 
Please make sure appropriate permissions are provided to the folder entirely and file manipulation is in purview.

-> Please make sure you do `go get github.com/sirupsen/logrus` before running the application.

To Start the Application , Please open two terminals and traverse to the folder where you have the downloaded executable folder

START SERVER :
Run `./go-task-list`

and that is it , you will see Verbose logging on this screen for requests , errors and handled panics.

USE Shell Script CLI :
Run `./tasklist.sh {Options}`

:: Options ::
1. add - To create a Task
2. list (Make choice when provided)
	- All Tasks
	- All Todo Tasks
	- Tasks Due Today
	- Tasks Overdue
	- Task By Title
3. done : Mark a Task Done
4. updateduedate : Update tasks due date
5. delete : Delete a Task

[Exit Code Documentation for CLI](go_task_list_executables/linux_amd64/exit_codes.md)

