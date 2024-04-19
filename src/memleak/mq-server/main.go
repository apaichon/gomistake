package main

import (
	"fmt"
	"encoding/json"
	_ "net"
	"time"
	"net/http"
	"mq-server/model"
	"sync"
	"runtime"
)

var (
	sliceMutex sync.Mutex
	messages []*model.MessageModel
	totalMessages int

)

func backgroundTask() {
	for {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Printf("Initial memory usage: %d KBytes\n", m.TotalAlloc /1024)

		// Perform your background task here
		fmt.Println("Background task executed at", time.Now())
		fmt.Println("Total Message", len(messages))

		// copy messages to run go routine for sub tasks
		totalMessages =len(messages)
		copiedMessages := make([]*model.MessageModel, totalMessages)
		
		copy(copiedMessages, messages)
		fmt.Printf("Copied Message memory usage: %d kBytes\n", m.TotalAlloc/1024)
		

		// bankWorkflow := workflows.NewBankingWorkflow()
		// messageWorkflow := workflows.NewMessageQueueWorkflow()

		// fmt.Printf("New Workflow memory usage: %d bytes\n", m.TotalAlloc)
		
		/*go messageWorkflow.CreateMessages(requests)
		fmt.Printf("Message Queue memory usage: %d bytes\n", m.TotalAlloc)
		
		go bankWorkflow.InsertDeposits(requests)
		fmt.Printf("Deposit usage: %d bytes\n", m.TotalAlloc)
		*/
		/*go bankWorkflow.InsertTransactionHistories(requests)
		fmt.Printf("Transaction History usage: %d bytes\n", m.TotalAlloc)
		*/

		copiedMessages = nil

	
		// Wait for 5 seconds
		if totalMessages > 0 {
			moveData()
		}
		fmt.Printf("After Slice usage: %d kBytes\n", m.TotalAlloc/1024)
		// Force garbage collection to reclaim memory
		runtime.GC()

		runtime.ReadMemStats(&m)
		fmt.Printf("Final memory usage: %d kBytes\n", m.TotalAlloc/1024)

		time.Sleep(5 * time.Second)
		
	}
}

func main() {
	// Specify the port to listen on
	port := "8888"
	host:="localhost"

	go backgroundTask()
	// Define HTTP route handlers
	http.HandleFunc("/messages", handleMessage)

	// Start HTTP server
	fmt.Printf("Server listening on port %s ...\n", port)
	http.ListenAndServe(host +":"+port, nil)
	
}

func moveData() {
	sliceMutex.Lock()
	defer sliceMutex.Unlock()
	messages = messages[totalMessages-1:len(messages)-1]
}

func appendData(value model.MessageModel) {
	sliceMutex.Lock()
	defer sliceMutex.Unlock()
	messages = append(messages, &value)
}

func handleMessage(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse JSON request body
	var message model.MessageModel
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Set CreatedAt to current time
	message.CreatedAt = time.Now()

	// Add message to the messages array
	appendData(message)
	
	// Send response
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Message received and stored successfully")
}

