package main

import "net/http"
import "fmt"
import "io/ioutil"
import "log"
import "encoding/json"
import "strconv"

func main() {
	fmt.Printf("Starting Calculating EMI Server...\n")

	runWebServer()
	//emiTest()
}

func runWebServer() {
	fileServer := http.FileServer(http.Dir("static"))
	http.HandleFunc("/emiCalculator", compute)
	http.Handle("/", fileServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func compute(writer http.ResponseWriter, request *http.Request) {
	emiRequest := handleRequest(request)

	emiSlice := CalculateEmi(emiRequest)

	fmt.Println("Number of installments:", len(emiSlice))

	jsonData, err := json.Marshal(emiSlice)

	if err != nil {
		log.Println("Error marshalling to stream")
	}

	log.Println(jsonData)
	writer.Write(jsonData)
	log.Println("Successfully commited the http response")

}

func printEmi(emi EquatedMonthlyInstallment) {
	fmt.Printf("\n%d %f %f %f\n", emi.SerialNo, emi.PrincipalAmount, emi.InterestAmount, emi.TotalAmount)

}

func handleRequest(request *http.Request) EmiRequest {
	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		log.Println("Cannot read request body")
	}

	log.Println(string(body))

	var emiRequestString EmiRequestString

	err = json.Unmarshal(body, &emiRequestString)

	if err != nil {
		log.Println("Failed to parse json request")
		log.Println(err)
	}

	emiRequest := convert(emiRequestString)
	PrintEmiRequest(emiRequest)

	return emiRequest
}

func convert(emiRequestString EmiRequestString) EmiRequest {
	amount, _ := strconv.ParseFloat(emiRequestString.MortgageAmount, 64)
	cost, _ := strconv.ParseFloat(emiRequestString.AnnualCostOfMortgage, 64)
	term, _ := strconv.ParseInt(emiRequestString.MortgageTerm, 10, 0)
	emiRequest := EmiRequest{amount, cost, int(term)}

	return emiRequest
}

func emiTest() {

	emiRequest := EmiRequest{50000, 4.5, 36}
	emiSlice := CalculateEmi(emiRequest)

	for i := 0; i < len(emiSlice); i++ {
		emi := emiSlice[i]
		fmt.Printf("\n%d %f %f %f\n", emi.SerialNo, emi.PrincipalAmount, emi.InterestAmount, emi.TotalAmount)
	}
}
