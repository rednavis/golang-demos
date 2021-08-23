package main

import (
	"encoding/json"
	"fmt"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"log"
	"net/http"
	"os"
	"strings"
)

var result map[string]interface{}

func makeRequest() {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=minsk" +
		"&units=metric&lang=en&appid=50b08e02174a79e2462e4c6fb1195769")

	if err != nil {
		log.Fatalln(err)
	}

	json.NewDecoder(resp.Body).Decode(&result)
}

func createTable() *widgets.QTableWidget {
	type keyValue struct {
		key   string
		value string
	}

	table := widgets.NewQTableWidget(nil)
	table.SetColumnCount(2)
	table.SetRowCount(7)

	table.VerticalHeader().SetVisible(true)

	header1 := widgets.NewQTableWidgetItem2("Parameter", 1)
	table.SetHorizontalHeaderItem(0, header1)

	header2 := widgets.NewQTableWidgetItem2("Value",1)
	table.SetHorizontalHeaderItem(1, header2)

	table.HorizontalHeader().SetStyleSheet("QHeaderView::section { border:1px solid #D8D8D8 }")

	var weatherParameters [] keyValue
	weatherParameters = append(weatherParameters, keyValue{"Temperature", fmt.Sprintf("%v", result["main"].(map[string]interface{})["temp"]) + " °C"})
	weatherParameters = append(weatherParameters, keyValue{"Feels like", fmt.Sprintf("%v", result["main"].(map[string]interface{})["feels_like"]) + " °C"})
	weatherParameters = append(weatherParameters, keyValue{"Pressure", fmt.Sprintf("%v", result["main"].(map[string]interface{})["pressure"]) + " hPa"})
	weatherParameters = append(weatherParameters, keyValue{"Humidity", fmt.Sprintf("%v", result["main"].(map[string]interface{})["humidity"]) + " %"})
	weatherParameters = append(weatherParameters, keyValue{"Wind speed", fmt.Sprintf("%v", result["wind"].(map[string]interface{})["speed"]) + " m/s"})
	weatherParameters = append(weatherParameters, keyValue{"Wind gust", fmt.Sprintf("%v", result["wind"].(map[string]interface{})["gust"]) + " m/s"})
	weatherParameters = append(weatherParameters, keyValue{"Visibility", fmt.Sprintf("%v", result["visibility"]) + " m"})

	for i, weatherParameter := range weatherParameters {
		parameter := widgets.NewQTableWidgetItem2(weatherParameter.key, 1)
		parameterValue := widgets.NewQTableWidgetItem2(weatherParameter.value, 1)

		table.SetItem(i,0, parameter)
		table.SetItem(i,1, parameterValue)
	}


	return table
}

func makeWidget () {
	application := widgets.NewQApplication(len(os.Args), os.Args)

	mainWindow := widgets.NewQMainWindow(nil, 0)
	mainWindow.SetMinimumSize2(350, 310)
	mainWindow.SetWindowTitle("Golang with Qt weather demo")
	mainWindow.SetWindowIcon(gui.NewQIcon5("./icons/icon32.png"))

	weatherWidget := widgets.NewQGroupBox2("Weather in Minsk", nil)
	weatherWidget.SetLayout(widgets.NewQVBoxLayout())
	weatherWidget.SetAlignment(4)
	mainWindow.SetCentralWidget(weatherWidget)

	weatherDescription := fmt.Sprintf("%v",result["weather"].([]interface{})[0].(map[string]interface{})["description"])
	weatherDescription = strings.Title(weatherDescription)

	weatherDescriptionLabel := widgets.NewQLabel2(weatherDescription, nil, 0)
	weatherWidget.Layout().AddWidget(weatherDescriptionLabel)

	cloudiness := "Cloudiness " + fmt.Sprintf("%v", result["clouds"].(map[string]interface{})["all"]) + " %"
	cloudinessLabel := widgets.NewQLabel2(cloudiness, nil, 0)
	weatherWidget.Layout().AddWidget(cloudinessLabel)

	weatherTable := createTable()
	weatherWidget.Layout().AddWidget(weatherTable)

	mainWindow.Show()

	application.Exec()
}


func main() {
	makeRequest()
	makeWidget()
}
