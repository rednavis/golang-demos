# Authorization demo
Golang demo project with registration / authorization of users and logging of visits via REST interface. When a user logs in, the statistics of the user's visits are displayed.  
User authentication is based on JWT tokens. PostgreSQL is used to store user data.  
When compiling the source code of the project, the output is the **`authorization_go.exe`** file.  
For the application to work, it is necessary:  
- to place folder with html-file templates;  
- .env file with database settings in the directory with the executable file;  
- installed and configured version of the PostgreSQL.
PostgreSQL should be with the following settings:  
- user: postgres  
- password: password  
- port: 5432
The user database must be previously created.  
The database tables are created by the application.
