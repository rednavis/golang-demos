# Golang with Qt weather demo
An example of working Go with Qt when creating a desktop application. The application interacts with the [web service](https://openweathermap.org/ "Web service for obtaining weather data") to obtaining weather data using the API provided by the service (for more information on the [API](https://openweathermap.org/api/ "Information about the weather web service API")). The received data is displayed to the user in the desktop application window.  
Getting data about the weather in the city is carried out by the custom function **`MakeRequest()`** using the request:  

**`http://api.openweathermap.org/data/2.5/weather?q=minsk&units=metric&lang=en&appid={API key}`**

where **`API key`** is a special key provided to registered users of the service. The demo application uses a key generated to use the free functionality provided by the service. In the request, using the special variables **`units`** and **`lang`**, data are obtained in the metric system of measurements and in English. The request is sent using the function:  

**`resp, err := http.Get (url string)`**  

![Http GET request](https://raw.githubusercontent.com/rednavis/golang-demos/main/go_qt_weather/images/image1.png "Http GET request")
  
The service has the ability to provide responses in JSON, XML and HTML formats. The response contains information about the current, minimum and maximum temperature, pressure, humidity, wind speed and others. The basic response format used in the example is JSON. The response in JSON format is decoded by the **`NewDecoder(r io.Reader) *Decoder function`** and entered by the **`(dec *Decoder) Decode(v interface{})`** error function into the **`result map[string]interface{}`** variable with the string **`json.NewDecoder(resp.Body).Decode(&result)`**.  

![Decoding a JSON-received weather service response](https://raw.githubusercontent.com/rednavis/golang-demos/main/go_qt_weather/images/image2.png "Decoding a JSON-received weather service response")   

After receiving a response from the service with the **`makeWidget()`** function by calling **`Qt`** the **`QApplication`** class, designed to create applications with a graphical interface, a framework is formed to manage future widgets. Integration of Go and Qt is done using a special plug-in library that provides interoperability.  

![Creating a class for managing an application with a graphical interface](https://raw.githubusercontent.com/rednavis/golang-demos/main/go_qt_weather/images/image3.png "Creating a class for managing an application with a graphical interface")  

The main window widget is then created using the **`QMainWindow`** class, which will be the container for the rest of the widgets.

![Creating the main widget window](https://raw.githubusercontent.com/rednavis/golang-demos/main/go_qt_weather/images/image4.png "Creating the main widget window")  

The table containing the weather data is created using the **`createTable() * widgets.QTableWidget`** function. This function uses a special widget to display data - **`QTableWidget`**.  

![Creating a table with weather data](https://raw.githubusercontent.com/rednavis/golang-demos/main/go_qt_weather/images/image5.png "Creating a table with weather data")  

Adding widgets to the main application widget window is carried out using functions like **`weatherWidget.Layout().AddWidget(weatherTable)`**.  

![Adding widgets to the main application widget window](https://raw.githubusercontent.com/rednavis/golang-demos/main/go_qt_weather/images/image6.png "Adding widgets to the main application widget window")  

After that, the main window is displayed with the function **`(ptr *QWidget) Show()`**.  

![Main window display function](https://raw.githubusercontent.com/rednavis/golang-demos/main/go_qt_weather/images/image7.png "Main window display function")  

The application is started by the function **`(ptr *QApplication) Exec() int`**.  

![Launching the application for display](https://raw.githubusercontent.com/rednavis/golang-demos/main/go_qt_weather/images/image8.png "Launching the application for display")  

The entry point to the program is the main file, which contains function calls to obtain weather data **`makeRequest()`** and **`makeWidget()`** - create a widget to display weather information.  
When compiling the source code of the project, the output is the **`go_qt_weather.exe`** file.  
Additionally, you need to move the icon folder to the directory with the executable file.  
When you run the executable file **`go_qt_weather.exe`**, a **`qtbox`** folder is created in the directory and **`qtbox`** is downloaded. Qt is deployed in this folder and an **`qtbox.exe`** file is created, which is necessary for the application to work.  

![Downloading qtbox](https://raw.githubusercontent.com/rednavis/golang-demos/main/go_qt_weather/images/image9.png "Downloading qtbox")  

When the application is launched, a window is displayed with information about the weather at the current time in Minsk.  

![Golang with Qt weather demo application](https://raw.githubusercontent.com/rednavis/golang-demos/main/go_qt_weather/images/image10.png "Golang with Qt weather demo application")
