var mortgage = (function() {
	
	function displayRequest(emiRequest) {
		console.log("Mortgage Amount:"+ emiRequest.MortgageAmount);
		console.log("mortgage cost:"+emiRequest.AnnualCostOfMortgage);
		console.log("Term: "+emiRequest.MortgageTerm);
	}
	
	function getEmiRequest() {
		var emiRequest = {};
		
		emiRequest.MortgageAmount = $("#mortgageAmount").val();
		emiRequest.AnnualCostOfMortgage = $("#mortgageCost").val();
		emiRequest.MortgageTerm = $("#mortgageTerm").val();
		
		displayRequest(emiRequest);
		
		return emiRequest;
	}
	
	function getJsonEmiRequest() {
		var emiRequest = getEmiRequest();
		var jsonStr = JSON.stringify(emiRequest);
		console.log("Json for emrequest");
		
		return jsonStr;
	}
	
	function logPayment(payment) {
		console.log("SerialNo: "+payment.serialNo+" principalAmount: "+payment.principalAmount+" interestAmount: "+payment.interestAmount+" totalAmount: "+payment.totalAmount);
	}
	
    function logEmiResponse(paymentList) {
		
		for(var i=0; i<paymentList.length; i++) {
			logPayment(paymentList[i]);
		}
	}
    
    function createTableRow(payment) {
    	var row = "<tr>";
    	
    	for(var prop in payment) {
    		
    		if(payment.hasOwnProperty(prop)) {
	    		row += "<td>";
	    		row += payment[prop];
	    		row += "</td>";
    		}
    	}
    	
    	row += "</tr>";
    	return row;
    }
    
    function createHeaderCell(headerName) {
    	var header = "<th>"
        	header += headerName
        	header += "</th>"
        		
       return header;
    }
    
    function createTableHeader() {
    	var header = "<tr>"
    		
    	header += createHeaderCell("SlNo");  		
    	header += createHeaderCell("Principal Amt($)");
    	header += createHeaderCell("Interest Amt($)");
    	header += createHeaderCell("Total Amt($)");
    	
    	header += "</tr>";
    	
    	return header;
    }
    
    function createHtmlTable(paymentList) {
    	var htmlTable = "<table>";
    	htmlTable += createTableHeader();
    	
    	for(var i=1; i<paymentList.length; i++) {   			
    		htmlTable += createTableRow(paymentList[i]);
    	}
    	
    	htmlTable += "</table>";
    	
    	return htmlTable;
    }
    
    function getResponseHeader(payment) {
    	var header = "<div class='left'>";
    	
    	header += "Summary:&nbsp;&nbsp;";
    	header += "Total Principal Amt. ($): "+payment.principalAmount+"&nbsp;&nbsp;&nbsp;&nbsp;";
    	header += "Total Interest Amt. ($): "+payment.interestAmount+"&nbsp;&nbsp;&nbsp;&nbsp;";
    	header += "Total Paymet Amt. ($): "+payment.totalAmount+"&nbsp;&nbsp&nbsp;&nbsp;";
    	header += "</div>";
    	header += "<div><br/></div>"
    	
    	return header;
    }
    
    function displayEmiResponse(paymentList) {
    	var details = $("div#details");
    	var header = getResponseHeader(paymentList[0]);
    	var htmlTable = createHtmlTable(paymentList);
    	htmlTable = header + htmlTable;
    	console.log(htmlTable);
    	details.empty().append(htmlTable);
    }
    	
	function processResponse(paymentList) {
		logEmiResponse(paymentList);
		displayEmiResponse(paymentList);
	}
	
	var emiClient = {};
		
	emiClient.computeEmi = function() {
		$.post("emiCalculator",
				getJsonEmiRequest(),
				function(paymentList, status) {
						processResponse(paymentList);
					},
					
				"json");
	}
		
	return emiClient;
})();

$(document).ready(function(){
	
	$("#compute").click(function() {
		mortgage.computeEmi();
	});
	
	
});
