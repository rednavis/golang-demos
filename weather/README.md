# Working with a web service
An example is working with a [web service](https://openweathermap.org/ "Web service for obtaining weather data") for obtaining weather data using the API provided by the service (for more information about the [API](https://openweathermap.org/api/ "Information about the weather web service API")).
To start the program (entry point), the `main` function is used (by analogy with other programming languages).
Getting data about the weather in the city is carried out by the custom function **`MakeRequest ()`** using the request:
`http://api.openweathermap.org/data/2.5/weather?q=minsk&units=metric&lang=en&appid={API key}`
where `API key` is a special key provided to registered users of the service. The demo application uses a key generated to use the free functionality provided by the service.
In the request, using the special variables `units` and `lang`, data are obtained in the metric system of measurements and in English.
The request is sent using the function:  
`resp, err: = http.Get (url string)`  
![Http GET request](https://raw.githubusercontent.com/rednavis/golang-demos/main/weather/images/image1.png "Http GET request")


