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
 * Writing Your First Blockchain Application
 */

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the bap structure, with 4 properties.  Structure tags are used by encoding/json library
type Bap struct {
	Carrier   string `json:"carrier"`
	Vessel    string `json:"vessel"`
	Voyage    string `json:"voyage"`
	Vsldate   string `json:"vsldate"`
	Snddate   string `json:"snddate"`
	Equiment  string `json:"equiment"`
}

/*
 * The Init method is called when the Smart Contract "baplie" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "baplie"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately

	if function == "queryBap" {
		return s.queryBap(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "createBap" {
		return s.createBap(APIstub, args)
	} else if function == "queryAllBaps" {
		return s.queryAllBaps(APIstub)
	} else if function == "changeBapOwner" {
		return s.changeBapOwner(APIstub, args)
	} else if function == "queryWithParam" {
		return s.queryWithParam(APIstub, args)
	} else if function == "queryByCarrier" {
		return s.queryByCarrier(APIstub, args)
	} else if function == "queryByVessel" {
		return s.queryByVessel(APIstub, args)
	} else if function == "queryByVoyage" {
		return s.queryByVoyage(APIstub, args)
	} else if function == "queryByVsldate" {
		return s.queryByVsldate(APIstub, args)
	} else if function == "queryBySnddate" {
		return s.queryBySnddate(APIstub, args)
	} else if function == "queryByEquiment" {
		return s.queryByEquiment(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name: " + function)
}

func (s *SmartContract) queryBap(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	bapAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(bapAsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	baps := []Bap{
/*
		Bap{Make: "Toyota", Model: "Prius", Colour: "blue", Owner: "Tomoko"},
		Bap{Make: "Ford", Model: "Mustang", Colour: "red", Owner: "Brad"},
		Bap{Make: "Hyundai", Model: "Tucson", Colour: "green", Owner: "Jin Soo"},
		Bap{Make: "Volkswagen", Model: "Passat", Colour: "yellow", Owner: "Max"},
		Bap{Make: "Tesla", Model: "S", Colour: "black", Owner: "Adriana"},
		Bap{Make: "Peugeot", Model: "205", Colour: "purple", Owner: "Michel"},
		Bap{Make: "Chery", Model: "S22L", Colour: "white", Owner: "Aarav"},
		Bap{Make: "Fiat", Model: "Punto", Colour: "violet", Owner: "Pari"},
		Bap{Make: "Tata", Model: "Nano", Colour: "indigo", Owner: "Valeria"},
		Bap{Make: "Holden", Model: "Barina", Colour: "brown", Owner: "Shotaro"},
		Bap{Make: "Kia", Model: "K5", Colour: "brown", Owner: "Hyunwoo"},
*/
	}

	i := 0
	for i < len(baps) {
		fmt.Println("i is ", i)
		bapAsBytes, _ := json.Marshal(baps[i])
		APIstub.PutState("BAP"+strconv.Itoa(i), bapAsBytes)
		fmt.Println("Added", baps[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) createBap(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}

	var bap = Bap{Carrier: args[1], Vessel: args[2], Voyage: args[3], Vsldate: args[4], Snddate: args[5], Equiment: args[6]}

	bapAsBytes, _ := json.Marshal(bap)
	APIstub.PutState(args[0], bapAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryAllBaps(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "BAP0"
	endKey := "BAP999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllBaps:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) changeBapOwner(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	bapAsBytes, _ := APIstub.GetState(args[0])
	bap := Bap{}

	json.Unmarshal(bapAsBytes, &bap)
	bap.Carrier = args[1]

	bapAsBytes, _ = json.Marshal(bap)
	APIstub.PutState(args[0], bapAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryWithParam(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
 		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	condition := args[0]
	fmt.Printf("args: %s", condition)
        //queryString := fmt.Sprintf("{\"selector\":{\"carrier\":\"%s\"}}", carrier)
        //queryString := fmt.Sprintf("{\"selector\":\"%s\"}", carrier)
	queryString := fmt.Sprintf("{\"selector\":{%s}}", condition)

	fmt.Printf("queryString: %s", queryString)

        queryResults, err := s.getQueryResultForQueryString(APIstub, queryString)
        if err != nil {
                return shim.Error(err.Error())
        }
        return shim.Success(queryResults)
}

func (s *SmartContract) queryByCarrier(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	carrier := args[0]

	//queryString := fmt.Sprintf("{\"selector\":{\"carrier\":\"%s\"}}", carrier)
        queryString := fmt.Sprintf("{\"selector\":\"%s\"}", carrier)

	queryResults, err := s.getQueryResultForQueryString(APIstub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (s *SmartContract) queryByVessel(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	vessel := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"vessel\":\"%s\"}}", vessel)

	queryResults, err := s.getQueryResultForQueryString(APIstub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (s *SmartContract) queryByVoyage(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	voyage := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"voyage\":\"%s\"}}", voyage)

	queryResults, err := s.getQueryResultForQueryString(APIstub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (s *SmartContract) queryByVsldate(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	vsldate := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"vsldate\":\"%s\"}}", vsldate)

	queryResults, err := s.getQueryResultForQueryString(APIstub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}


func (s *SmartContract) queryBySnddate(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	snddate := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"snddate\":\"%s\"}}", snddate)

	queryResults, err := s.getQueryResultForQueryString(APIstub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (s *SmartContract) queryByEquiment(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	equiment := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"equiment\":\"%s\"}}", equiment)
	//queryString := fmt.Sprintf("{\"selector\":%s}", carrier)

	queryResults, err := s.getQueryResultForQueryString(APIstub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (s *SmartContract) getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString: %s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult: %s\n", buffer.String())

	return buffer.Bytes(), nil
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {
	
	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
