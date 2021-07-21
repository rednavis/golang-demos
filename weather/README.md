# Working with a web service demo
An example is working with a [web service](https://openweathermap.org/ "Web service for obtaining weather data") for obtaining weather data using the API provided by the service (read for more information about the [API](https://openweathermap.org/api/ "Information about the weather web service API")).
To start the program (entry point), the **`main()`** function is used (by analogy with other programming languages).
Getting data about the weather in the city is carried out by the custom function **`MakeRequest()`** using the request:  

`http://api.openweathermap.org/data/2.5/weather?q=minsk&units=metric&lang=en&appid={API key}`  

where **`API key`** is a special key provided to registered users of the service. The demo application uses a key generated to use the free functionality provided by the service.
In the request, using the special variables **`units`** and **`lang`**, data are obtained in the metric system of measurements and in English.
The request is sent using the function:  

`resp, err: = http.Get (url string)`  

![Http GET request](https://raw.githubusercontent.com/rednavis/golang-demos/main/weather/images/image1.png "Http GET request")

The service has the ability to provide responses in JSON, XML and HTML formats. The response contains information about the current, minimum and maximum temperature, pressure, humidity, wind speed and others. The basic response format used in the example is JSON.
The response in JSON format is decoded by the **`NewDecoder(r io.Reader) *Decoder`** function and entered by the **`(dec *Decoder) Decode(v interface{})`** error function into the **`result map[string]interface{}`** variable with the string **`json.NewDecoder(resp.Body).Decode(&result).`**

![Decoding a JSON-received weather service response](https://raw.githubusercontent.com/rednavis/golang-demos/main/weather/images/image2.png "Decoding a JSON-received weather service response")

After receiving a response on port 60000 using the **`ServeMux`** HTTP request multiplexer, the html page template is generated using the custom function **`htmlIndexPageTemplateHandler(w http.ResponseWriter, r *http.Request)`**.

![Formation of the html-page template by the Http-request multiplexer](https://raw.githubusercontent.com/rednavis/golang-demos/main/weather/images/image3.png "Formation of the html-page template by the Http-request multiplexer")

The **`htmlIndexPageTemplateHandler (w http.ResponseWriter, r *http.Request)`** function creates a **`weatherDescription`** variable containing a description of the weather in the language specified in the request (English by default), and variables **`temperature`**, **`pressure`**, **`humidity`** containing temperature, pressure and humidity data, respectively. Based on this data, the text of the **`indexPageTemplate`** html-page template is generated using the received data. The template text is passed to the **`template.New("weather").Parse(indexPageTemplate)`** function, which forms a new template using the **`New(name string) *Template`** function and the parse text of the **`indexPageTemplate`** template using the **`(t *Template) Parse(text string) (*Template, error)`**.  

![Generating a final html-page using predefined template](https://raw.githubusercontent.com/rednavis/golang-demos/main/weather/images/image4.png "Generating a final html-page using predefined template")

After that, this template is launched by **`htmlTemplate.Execute(w, indexPageTemplate)`** using the **`(t *Template) Execute(wr io.Writer, data interface{})`** error function, which applies the parsed template to the specified data object and writes to the **`wr`** variable to display data.

![Html-page template executing](https://raw.githubusercontent.com/rednavis/golang-demos/main/weather/images/image5.png "Html-page template executing")

When compiling the source code of the project, the output is the **`weather_go.exe`** file. After launching the application, when called in the browser at **`http://localhost:60000`**, a page with information about the weather should be displayed.

![Weather demo html-page](https://raw.githubusercontent.com/rednavis/golang-demos/main/weather/images/image6.png "Weather demo html-page")
