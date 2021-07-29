# Working with financial service demo
An example is working with a [web service](https://rapidapi.com/ "Web service for obtaining financial data") for obtaining financial data. In this case, getting the NASDAQ index using the API provided by the service (you can read more about the [API](https://docs.rapidapi.com/ "Information about the financial web service API")).
To start the program (entry point), the **`main()`** function is used (by analogy with other programming languages).  
To enter the input parameters, the **`GetArguments()`** function is used, which takes two parameters **`name`** and **`time`** from the command line - the name of the company and the time period for which the data is provided.  
Receiving data for plotting a chart is carried out by the custom function **`MakeRequest()`** using a request via the **`URL`**:

`url:=https://apidojo-yahoo-finance-v1.p.rapidapi.com/market/get-charts?symbol={company_name}&interval=1d&range={time_interval}&region=US&comparisons=NDAQ`  
`req, _ := http.NewRequest("GET", url, nil)`  

where **`company_name`** - is the name of the company in the service providing the data;  
      **`time_interval`** - is the time period (starting from the current day) for which data is provided.  
 Additionally, parameters are added to the request body: API key and hostname (to receive data from a service that provides information on financial indices)

`req.Header.Add("x-rapidapi-key", {API key}")`  
`req.Header.Add("x-rapidapi-host", "apidojo-yahoo-finance-v1.p.rapidapi.com")`

where **`API key`** is a special key provided to registered users of the service. The demo application uses a key generated to use the free functionality provided by the service.

![Http GET request](https://raw.githubusercontent.com/rednavis/golang-demos/main/chart/images/image1.png "Http GET request")

The request is sent using the function:  

` resp, err := http.DefaultClient.Do(req)`  

![Http request send](https://raw.githubusercontent.com/rednavis/golang-demos/main/chart/images/image2.png "Http request send")

The service provides a response in JSON format. The response contains information about the value of financial indices for the period specified in the request.
The response in JSON format is decoded by the **`NewDecoder(r io.Reader) *Decoder`** function and entered by the **`(dec *Decoder) Decode(v interface{})`** error function into the **`result map[string]interface{}`** variable with the string **`json.NewDecoder(resp.Body).Decode(&result).`**

![Decoding a JSON-received weather service response](https://raw.githubusercontent.com/rednavis/golang-demos/main/chart/images/image3.png "Decoding a JSON-received weather service response")

After that, the received response is checked for the absence of errors by the custom function **`ValidateRequestData()`**.

![Validating request data](https://raw.githubusercontent.com/rednavis/golang-demos/main/chart/images/image4.png "Validating request data")


If the response is successfully verified on port 60000 using the **`ServeMux`** HTTP request multiplexer, the html page template is generated using the custom function **`htmlIndexPageTemplateHandler(w http.ResponseWriter, r *http.Request)`**.

![Formation of the html-page template by the Http-request multiplexer](https://raw.githubusercontent.com/rednavis/golang-demos/main/chart/images/image5.png "Formation of the html-page template by the Http-request multiplexer")

The **`htmlIndexPageTemplateHandler(w http.ResponseWriter, r *http.Request)`** creates a **`chartScripts`** variable that receives the script body for filling and displaying a chart of financial index values using the **`getChartScripts()`** custom function.

![Variable storing the script of the chart](https://raw.githubusercontent.com/rednavis/golang-demos/main/chart/images/image6.png "Variable storing the script of the chart")

Based on this data, the text of the **`indexPageTemplate`** html-page template is generated using the received data. The template text is passed to the function **`htmlTemplate, err: =template.New("NDAQ index chart (closing price, USD)").Parse(indexPageTemplate)`** forming a new template using the **`New(name string) *Template`** function and the parse text of the **`indexPageTemplate`** template c using the **`(t *Template)Parse(text string)(*Template, error)`** function. 

![Generating a final html-page using predefined template](https://raw.githubusercontent.com/rednavis/golang-demos/main/chart/images/image7.png "Generating a final html-page using predefined template")

After that, this template is launched by **`htmlTemplate.Execute(w, indexPageTemplate)`** using the **`(t *Template) Execute(wr io.Writer, data interface{})`** error function, which applies the parsed template to the specified data object and writes to the **`wr`** variable to display data.

![Passing the html-page to display](https://raw.githubusercontent.com/rednavis/golang-demos/main/chart/images/image8.png "Passing the html-page to display")

When compiling the source code of the project, the output is the **`chart_go.exe`** file. 
After launching the application from the command line with input parameters, for example:

![Using the demo application](https://raw.githubusercontent.com/rednavis/golang-demos/main/chart/images/image9.png "Using the demo application")

when called in a browser at **`http://localhost:60000/`**, a page is displayed with a graph of the financial index change over time for a specified period.

![World stock indexes chart demo html-page](https://raw.githubusercontent.com/rednavis/golang-demos/main/chart/images/image10.png "World stock indexes chart demo html-page")

The resulting graph, when compared, repeats the shape and values of the graph obtained using the search service, for the company of interest to the user for the specified period.

![NASDAQ Index Chart obtained using search service](https://raw.githubusercontent.com/rednavis/golang-demos/main/chart/images/image11.png "NASDAQ Index Chart obtained using search service")
