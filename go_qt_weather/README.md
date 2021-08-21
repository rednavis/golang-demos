# Golang with Qt weather demo
An example of working Go with Qt when creating a desktop application. The application interacts with the [web service](https://openweathermap.org/ "Web service for obtaining weather data") to obtaining weather data using the API provided by the service (for more information on the [API](https://openweathermap.org/api/ "Information about the weather web service API")). The received data is displayed to the user in the desktop application window.  
When compiling the source code of the project, the output is the **`go_qt_weather.exe`** file.  
Additionally, you need to move the icon folder to the directory with the executable file.
When you run the executable file **`go_qt_weather.exe`**, a **`qtbox`** folder is created in the directory and **`qtbox`** is downloaded. Qt is deployed in this folder and an **`qtbox.exe`** file is created, which is necessary for the application to work.
