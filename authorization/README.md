# Authorization demo
Golang demo project with registration / authorization of users and logging of visits via REST interface. When a user logs in, the statistics of the user's visits are displayed.  
User authentication is based on JWT tokens. PostgreSQL is used to store user data.  
The structure of the project is shown in the figure below.  
![Project structure](https://raw.githubusercontent.com/rednavis/golang-demos/main/authorization/images/image1.png "Project structure")
The app directory contains the **`authorization.go`** file that implements Middleware for working with JWT. This Middleware provides for intercepting all requests, checking for the presence of an authentication token (JWT), checking whether it is valid and valid, and then sending an error message if deficiencies are found, or vice versa, processing the request if everything is in order.  
The **`controllers`** directory contains files:
- **`actionsControllers.go`**, which implements controller functions for the **`Action`** entity. Namely:
 validation, creation and receipt of user actions;
- **`authorizationControllers.go`**, which implements the controller functions for the **`Account`** entity required for authorization. Namely: account creation and authentication.  

The **`models`** directory contains files:
- **`accounts.go`**, which implements the structure and functions of the **`Account`** entity. Namely: validation, creating an account, logging in and getting a user;
- **`actions.go`**, which implements the structure and functions of the **`Action entity`**. Namely: validation, creating a record of user actions and getting a list of user actions;
- **`base.go`**, which implements interaction with the PosgreSQL database and the creation of database tables based on the **`Account`** and **`Action`** structures.  

The **`templates`** directory contains templates for html pages:
- **`index.html`** - the start page of the application;
- **`registration.html`** - new user registration page;
- **`visit.html`** - pages with information about user visits to the application;
- **`404.html`** - 404 error page informing the user about the absence of the requested page or resource.  

The **`utils`** directory contains the **`util.go`** file, which implements the helper functions required for the application to run.
The **`.env`** file contains settings for connecting to the database.
The entry point to the program is the **`main`** file, which contains functions for creating html pages from templates, the implementation of connecting access points to the application API, using the Middleware to authenticate the user using JWT and creating a connection to the port (port 60000 in this demo application).
When compiling the source code of the project, the output is the **`authorization_go.exe`** file.  
For the application to work, it is necessary:  
- place the **`templates`** folder with templates for html files and **`.env`** file with database settings in the directory with the executable file;  
- installed and configured version of the PostgreSQL.
  
PostgreSQL should be with the following settings:  
- user: postgres  
- password: password  
- port: 5432  

The **`users`** database must be previously created.  
The database tables are created by the application.
When you go to the application page at http: // localhost: 60000, the main application page should open.
![Authorization demo start html-page](https://raw.githubusercontent.com/rednavis/golang-demos/main/authorization/images/image2.png "Authorization demo start html-page")
