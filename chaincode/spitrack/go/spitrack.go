/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Trade Finance Use Case - WORK IN  PROGRESS
 */

package main


import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}


// Define the order attributes
type PurchaseOrder struct {
	PurchaseOrderId			string		`json:"purchaseOrderId"`
	ParentOrderId			string		`json:"parentOrderId"`
	ChildOrderId			string		`json:"childOrderId"`
	IssueDate		string		`json:"issueDate"`
	Sender    string   `json:"sender"`
	Receiver		string		`json:"receiver"`
	GradeType		string		`json:"gradeType"`
	Amount			int		`json:"amount,int"`
	Price			int		`json:"price,int"`
	Status			string		`json:"status"`
	MaterialName   string `json:"materialName"`
	MaterialType   string `json:"materialType"`
	Quantity   int `json:"quantity,int"`
	Units   string `json:"units"`

}


type ShipmentSpices struct {
	PurchaseOrderId			string		`json:"purchaseOrderId"`
	ParentOrderId			string		`json:"parentOrderId"`
	ChildOrderId			string		`json:"childOrderId"`
	IssueDate		string		`json:"issueDate"`
	Sender    string   `json:"sender"`
	Receiver		string		`json:"receiver"`
	GradeType		string		`json:"gradeType"`
	Amount			int		`json:"amount,int"`
	Price			int		`json:"price,int"`
	Status			string		`json:"status"`
	MaterialName   string `json:"materialName"`
	MaterialType   string `json:"materialType"`
	Quantity   int `json:"quantity,int"`
	Units   string `json:"units"`
	Qualityscore int `json:"qualityscore,int"`
	OriginCd string `json:"originCd"`

}


func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "requestPO" {
		return s.requestPO(APIstub, args)
	} else if function == "acceptPO" {
		return s.acceptPO(APIstub, args)
	} else if function == "rejectPO" {
		return s.shipSpices(APIstub, args)
	}	else if function == "shipSpices" {
		return s.shipSpices(APIstub, args)
	}else if function == "getPODetails" {
		return s.getPODetails(APIstub, args)
	}else if function == "getPOHistory" {
		return s.getPOHistory(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}


func (s *SmartContract) acceptPO(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	// purchaseOrderId := struct {
	// 	PurchaseOrderId  string `json:"purchaseOrderId"`
	  
	// }{}
	// err  := json.Unmarshal([]byte(args[0]),&purchaseOrderId)
	// if err != nil {
	// 	return shim.Error("Not able to parse args into PO")
	// }

	POAsBytes, err := APIstub.GetState(args[0])
	if err!=nil {
		return shim.Error("failed to get Purchase order"+err.Error())
	}
	po := PurchaseOrder{}

	err = json.Unmarshal(POAsBytes, &po)

	if err != nil {
		return shim.Error("Issue with PO json unmarshaling"+string(POAsBytes))
	}
   /*
  PurchaseOrderId			string		`json:"purchaseOrderId"`
	ParentOrderId			string		`json:"parentOrderId"`
	ChildOrderId			string		`json:"childOrderId"`
	IssueDate		string		`json:"issueDate"`
	Sender    string   `json:"sender"`
	Receiver		string		`json:"receiver"`
	GradeType		string		`json:"gradeType"`
	Amount			int		`json:"amount,int"`
	Price			int		`json:"price,int"`
	Status			string		`json:"status"`
	MaterialName   string `json:"materialName"`
	MaterialType   string `json:"materialType"`
	Quantity   int `json:"quantity,int"`
	Units   string `json:"units"`
   */

	po.Status="Accepted"
			
	POBytes, err := json.Marshal(po)

	if err != nil {
		return shim.Error("Issue with PO json marshaling")
	}

    APIstub.PutState(po.PurchaseOrderId,POBytes)
	fmt.Println("PO Accepted -> ", po)


	

	return shim.Success(nil)
}

func (s *SmartContract) rejectPO(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	purchaseOrderId := struct {
		PurchaseOrderId  string `json:"purchaseOrderId"`
	  
	}{}
	err  := json.Unmarshal([]byte(args[0]),&purchaseOrderId)
	if err != nil {
		return shim.Error("Not able to parse args into PO")
	}

	POAsBytes, _ := APIstub.GetState(purchaseOrderId.PurchaseOrderId)

	po := PurchaseOrder{}

	err = json.Unmarshal(POAsBytes, &po)

	if err != nil {
		return shim.Error("Issue with PO json unmarshaling")
	}
    /*
    PurchaseOrderId			string		`json:"purchaseOrderId"`
	ParentOrderId			string		`json:"parentOrderId"`
	ChildOrderId			string		`json:"childOrderId"`
	IssueDate		string		`json:"issueDate"`
	Sender    string   `json:"sender"`
	Receiver		string		`json:"receiver"`
	GradeType		string		`json:"gradeType"`
	Amount			int		`json:"amount,int"`
	Price			int		`json:"price,int"`
	Status			string		`json:"status"`
	MaterialName   string `json:"materialName"`
	MaterialType   string `json:"materialType"`
	Quantity   int `json:"quantity,int"`
	Units   string `json:"units"`
    */

	po.Status=  "Rejected"
			
	POBytes, err := json.Marshal(po)

	if err != nil {
		return shim.Error("Issue with PO json marshaling")
	}

    APIstub.PutState(po.PurchaseOrderId,POBytes)
	fmt.Println("PO Rejected -> ", po)


	

	return shim.Success(nil)
}



// This function is initiate by importer  to manufacturing/exporter OR manufacturing/exporter to Supplier
func (s *SmartContract) requestPO(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {


	PO := PurchaseOrder{}

	err  := json.Unmarshal([]byte(args[0]),&PO)
 if err != nil {
		return shim.Error("Not able to parse args into PO")
	}
	POBytes, err := json.Marshal(PO)
    APIstub.PutState(PO.PurchaseOrderId,POBytes)
	fmt.Println("PO Requested -> ", PO)

	

	return shim.Success(nil)
}

// This function is initiate by Supplier to Manufacturer/Exporter and manufacturing/exporter to importer
func (s *SmartContract) shipSpices(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	shipSpices := struct {
		PurchaseOrderId  string `json:"purchaseOrderId"`
	}{}
	err  := json.Unmarshal([]byte(args[0]),&shipSpices)
	if err != nil {
		return shim.Error("Not able to parse args to shipSpices")
	}
	
	// if err != nil {
	// 	return shim.Error("No Amount")
	// }

	SSBytes, _ := APIstub.GetState(shipSpices.PurchaseOrderId)

	ss:= ShipmentSpices{}
	err = json.Unmarshal(SSBytes, &ss)

	if err != nil {
		return shim.Error("Issue with SS json unmarshaling")
	}

	ss.Status = "Shipped spices"
	SSBytes, err = json.Marshal(ss)

	if err != nil {
		return shim.Error("Issue with SS json marshaling")
	}

    APIstub.PutState(ss.PurchaseOrderId,SSBytes)
	fmt.Println("Shipped Spices -> ", ss)


	return shim.Success(nil)
}

// This function is to retrieve the shipment details and status
func (s *SmartContract) getShippingDetails(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	
	SP := ShipmentSpices{}

	err  := json.Unmarshal([]byte(args[0]),&SP)
    if err != nil {
		return shim.Error("Not able to parse args into Shipping #")
	}
	
	SPAsBytes, _ := APIstub.GetState(SP.PurchaseOrderId)
	err  = json.Unmarshal(SPAsBytes,&SP)

	if err != nil{
		return shim.Error("Invalid Purchaseorder ID or Purchaseorder ID doesn't exist")
	}

	fmt.Println("Shipment Details  -> ", SP)

	return shim.Success(SPAsBytes)
}

// This function is to retrieve the PO details and status
func (s *SmartContract) getPODetails(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

/*
	PO := PurchaseOrder{}
	err  := json.Unmarshal([]byte(args[0]),&PO)
 if err != nil {
		return shim.Error("Not able to parse args into PO")
	}
	
	POAsBytes, _ := APIstub.GetState(PO.PurchaseOrderId)
	err  = json.Unmarshal(POAsBytes,&PO)

	if err != nil{
		return shim.Error("Invalid PO ID or PO ID doesn't exist")
	}

	fmt.Println("PO Details -> ", PO)

	return shim.Success(POAsBytes) */
	purchaseOrderId := args[0];
	
	// if err != nil {
	// 	return shim.Error("No Amount")
	// }

	POAsBytes, _ := APIstub.GetState(purchaseOrderId)

	return shim.Success(POAsBytes)
}

func (s *SmartContract) getPOHistory(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	purchaseOrderId := args[0];
	
	

	resultsIterator, err := APIstub.GetHistoryForKey(purchaseOrderId)
	if err != nil {
		return shim.Error("Error retrieving PO history.")
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error("Error retrieving PO history.")
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getPOHistory returning:\n%s\n", buffer.String())

	

	return shim.Success(buffer.Bytes())
}


// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
