package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var result map[string]interface{}

//Page url - http://localhost:60000/
func MakeRequest() {
		resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=minsk" +
									"&units=metric&lang=en&appid=50b08e02174a79e2462e4c6fb1195769")
		if err != nil {
			log.Fatalln(err)
		}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println("--------\nServer response:\n", result,"\n\nNow open http://localhost:60000")

	port := os.Getenv("PORT")
	if port == "" {
		port = "60000"
	}

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", htmlIndexPageTemplateHandler)
	http.ListenAndServe(":"+port, serveMux)
}

func htmlIndexPageTemplateHandler(w http.ResponseWriter, r *http.Request) {

	if result == nil {
		log.Fatalln("Request result not received or incorrect!")
		return
	}

	//Get weather parameters from request result
	var weatherDescription = fmt.Sprintf("%v",result["weather"].([]interface{})[0].(map[string]interface{})["description"])
	weatherDescription = "<h5>" + weatherDescription + "</h5>"

	var temperature = fmt.Sprintf("%v", result["main"].(map[string]interface{})["temp"])
	var pressure = fmt.Sprintf("%v", result["main"].(map[string]interface{})["pressure"])
	var humidity = fmt.Sprintf("%v", result["main"].(map[string]interface{})["humidity"])

	var parametersList ="<li class=\"list-group-item\">Temperature - " + temperature + " °C</li>"
	parametersList += "<li class=\"list-group-item\">Pressure - " + pressure + " mmHg</li>"
	parametersList += "<li class=\"list-group-item\">Humidity - " + humidity + " %</li>"
	parametersList = "<ul class=\"list-group\">" + parametersList + "</ul>"

	//Creation a template for a demo html page
	indexPageTemplate :=  `<html lang="ru">
					<head>
						<meta charset="UTF-8">
						<title>Title</title>
						<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-EVSTQN3/azprG1Anm3QDgpJLIm9Nao0Yz1ztcQTwFspd3yD65VohhpuuCOmLASjC" crossorigin="anonymous">
					</head>
					<body>
						<nav class="navbar navbar-dark bg-primary">
							<div class="container">
    							<span class="navbar-text fw-bold fs-4">
									<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" class="bi bi-cloud-sun" viewBox="0 0 16 16">
  										<path d="M7 8a3.5 3.5 0 0 1 3.5 3.555.5.5 0 0 0 .624.492A1.503 1.503 0 0 1 13 13.5a1.5 1.5 0 0 1-1.5 1.5H3a2 2 0 1 1 .1-3.998.5.5 0 0 0 .51-.375A3.502 3.502 0 0 1 7 8zm4.473 3a4.5 4.5 0 0 0-8.72-.99A3 3 0 0 0 3 16h8.5a2.5 2.5 0 0 0 0-5h-.027z"/>
  										<path d="M10.5 1.5a.5.5 0 0 0-1 0v1a.5.5 0 0 0 1 0v-1zm3.743 1.964a.5.5 0 1 0-.707-.707l-.708.707a.5.5 0 0 0 .708.708l.707-.708zm-7.779-.707a.5.5 0 0 0-.707.707l.707.708a.5.5 0 1 0 .708-.708l-.708-.707zm1.734 3.374a2 2 0 1 1 3.296 2.198c.199.281.372.582.516.898a3 3 0 1 0-4.84-3.225c.352.011.696.055 1.028.129zm4.484 4.074c.6.215 1.125.59 1.522 1.072a.5.5 0 0 0 .039-.742l-.707-.707a.5.5 0 0 0-.854.377zM14.5 6.5a.5.5 0 0 0 0 1h1a.5.5 0 0 0 0-1h-1z"/>
									</svg>
									<span class="align-middle">Weather demo</span>
    							</span>
 							 </div>
							</nav>
						<div class="container">
							<h2>Weather in Minsk</h2>` + "\n" +
							weatherDescription + "\n" +
							parametersList + "\n" +
					`	</div>
						<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js" integrity="sha384-MrcW6ZMFYlzcLA8Nl+NtUVF0sA7MsXsP1UyJoMp4YLEuNSfAP+JcXn/tWtIaxVXM" crossorigin="anonymous"></script>
					 </body>
				</html>`

	//Create html-template
	htmlTemplate, err := template.New("weather").Parse(indexPageTemplate)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	htmlTemplate.Execute(w, indexPageTemplate)
}

func main() {
	MakeRequest()
}