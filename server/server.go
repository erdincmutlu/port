/*

What goes into here:

import encoding/json

// Database is map in memory
// Port's code is the key of the map
var DB map[string]types.Port

// Reading the port file 
file:=os.Open()
dec:=json.NewDedoder(file) // Create the decoder
dec.Token() // the next JSON token in the input stream (To pass opening and closing delimiter)
for dec.More() { : While there is another element in the array
	dec.Decode(): Read next Json encoded value
}
dec.Token()

// Store the ports in database to access later
// In regular periods this service can look at the updated ports.json file and upserts the updated/news records to the DB

*/

func startGRPCServer() {
	lis, err := net.Listen("tcp", "localhost:"+serverPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterPortServer(grpcServer, &Port{})

	grpcServer.Serve(lis)
}