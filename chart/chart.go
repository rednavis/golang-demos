package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var companyName string
var timeInterval string
var requestResult map[string]interface{}

func printHelp() {
	help := "\n----- Help -----\n" +
		"To get a response, enter the parameters using the command line\n" +
		"Parameters insert  is realized in the next format:\n" +
		"-name={NDAQ company name}\n" +
		"-time={NDAQ index time interval} in format 1d | 5d | 1mo | 3mo | 6mo | ytd | 1y | 2y | 5y | 10y | max\n" +
		"example: -name=TSLA -time=3mo"
	fmt.Println(help)
}

func GetArguments () {
	defaultNdaqName := "Not valid NDAQ company name! Read help please."
	defaultTimeInterval := "Not valid time Interval! Read help please."

	//Get arguments from comand line
	flag.StringVar(&companyName, "name", defaultNdaqName, "NDAQ company name")
	flag.StringVar(&timeInterval, "time", defaultTimeInterval, "NDAQ index time interval")
	flag.Parse()

	if (companyName == defaultNdaqName || timeInterval == defaultTimeInterval) {
		fmt.Println("----- WARNING!-----")

		if (companyName == defaultNdaqName) {
			fmt.Println(defaultNdaqName)
		}
		if (timeInterval == defaultTimeInterval) {
			fmt.Println(defaultTimeInterval )
		}

		printHelp()
		os.Exit(0)
	}
}

func ValidateRequestData() bool {
	if requestResult == nil {
		fmt.Println("\nRequest result not received or incorrect!")
		return false
	}

	if requestResult["chart"].(map[string]interface{})["result"] == nil {
		errorDescription := fmt.Sprintf("%v",requestResult["chart"].(map[string]interface{})["error"].(map[string]interface{})["description"])
		fmt.Println("\n----- ERROR! -----\n" + errorDescription + "\nCheck the correctness of the entered data!")
		return false
	}

	_, ok := requestResult["chart"].(map[string]interface{})["result"].([]interface{})[0].(map[string]interface{})["timestamp"]
	if !ok {
		fmt.Println("\nNo date and time values found. Request result not received or incorrect!\nCheck that the company name matches the names provided by the service!")
		return false
	}

	_, ok = requestResult["chart"].(map[string]interface{})["result"].([]interface{})[0].(map[string]interface{})["indicators"].(map[string]interface{})["quote"].([]interface{})[0].(map[string]interface{})["close"]
	if !ok {
		fmt.Println("\nIndex values not found. Request result not received or incorrect!\nCheck that the company name matches the names provided by the service!")
		return false
	}

	return true
}

func  MakeRequest() {
	url := "https://apidojo-yahoo-finance-v1.p.rapidapi.com/market/get-charts?symbol="+
		   companyName + "&interval=1d&range=" + timeInterval + "&region=US&comparisons=NDAQ"
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", "abc740684cmsh40afd915f4617c6p11dba2jsn630cf879b404")
	req.Header.Add("x-rapidapi-host", "apidojo-yahoo-finance-v1.p.rapidapi.com")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	json.NewDecoder(resp.Body).Decode(&requestResult)
	defer resp.Body.Close()

	if ValidateRequestData() {
		fmt.Println("\nNow to see the chart open http://localhost:60000")
	} else {
		return
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "60000"
	}

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", htmlIndexPageTemplateHandler)
	http.ListenAndServe(":"+port, serveMux)
}

func getChartScripts () string{
	var dataPoints string
	var timeStamp []interface{}
	var indexValues []interface{}

	//Get chart data from request result
	timeStamp = requestResult["chart"].(map[string]interface{})["result"].([]interface{})[0].(map[string]interface{})["timestamp"].([]interface{})
	indexValues = requestResult["chart"].(map[string]interface{})["result"].([]interface{})[0].(map[string]interface{})["indicators"].(map[string]interface{})["quote"].([]interface{})[0].(map[string]interface{})["close"].([]interface{})

	if timeStamp == nil || indexValues == nil {
		fmt.Println("Request result not received or incorrect!")
		os.Exit(0)
	}

	for i:=0; i< len(timeStamp); i++ {
		dataPoints += `dataPoints.push(prepareBarValue(` + fmt.Sprintf("%f",timeStamp[i].(float64)*1000) +`,` + fmt.Sprintf("%v",indexValues[i]) + `));` + "\n"
	}

	//Create chart scrips for html-page template
	chartScripts := `<script type="text/javascript" src="https://canvasjs.com/assets/script/canvasjs.min.js"></script>` + "\n"
	chartScripts += `<script>
			function prepareBarValue(barDate, value) {
			return {
					x : barDate,
					y : value
				};
			}
		
			var y = 0;
			var data = [];
			var dataSeries = { type: "line", 
							   xValueType: "dateTime",
							   markerSize: 0};
			//var dataPoints = [];
			
			function fullChartData() {
				var dataPoints = [];` +
				"\n" + dataPoints +
				`return dataPoints;
			}
			
			dataSeries.dataPoints = fullChartData();
			data.push(dataSeries);
			
			window.onload = function () {
			var chart = new CanvasJS.Chart("chartContainer", {
				animationEnabled: true,
				zoomEnabled: true,
				title:{
					text: "` + companyName + `" 
				},
				data: data  // random generator below
			});
			chart.render();
			}
		</script>`

	return chartScripts
}

func htmlIndexPageTemplateHandler(w http.ResponseWriter, r *http.Request) {
	//Creation a template for a demo html page
	chartScripts := getChartScripts()
	indexPageTemplate :=  `<html lang="ru">
					<head>
						<meta charset="UTF-8">
						<title>Title</title>
						<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-EVSTQN3/azprG1Anm3QDgpJLIm9Nao0Yz1ztcQTwFspd3yD65VohhpuuCOmLASjC" crossorigin="anonymous">` + "\n" +
						chartScripts +
					`</head>
					<body>
						<nav class="navbar navbar-dark bg-success">
							<div class="container">
    							<span class="navbar-text fw-bold fs-4">
									<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" class="bi bi-graph-up" viewBox="0 0 16 16">
  										<path fill-rule="evenodd" d="M0 0h1v15h15v1H0V0zm10 3.5a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5v4a.5.5 0 0 1-1 0V4.9l-3.613 4.417a.5.5 0 0 1-.74.037L7.06 6.767l-3.656 5.027a.5.5 0 0 1-.808-.588l4-5.5a.5.5 0 0 1 .758-.06l2.609 2.61L13.445 4H10.5a.5.5 0 0 1-.5-.5z"/>
									</svg>
									<span class="align-middle">World stock indexes chart demo</span>
    							</span>
 							 </div>
							</nav>
						<div class="container">
							<h2>NDAQ index chart (closing price, USD)</h2>
							<div style="width:100%;">
								<div id="chartContainer" style="height: 300px; width: 100%;">
							</div>` + "\n" +
						`</div>
						<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js" integrity="sha384-MrcW6ZMFYlzcLA8Nl+NtUVF0sA7MsXsP1UyJoMp4YLEuNSfAP+JcXn/tWtIaxVXM" crossorigin="anonymous"></script>
					 </body>
				</html>`

	//Create html-template
	htmlTemplate, err := template.New("NDAQ index chart (closing price, USD)").Parse(indexPageTemplate)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	htmlTemplate.Execute(w, indexPageTemplate)
}

func main() {
	GetArguments()
	MakeRequest()
}