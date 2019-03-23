# balance_manager
Balance manager microservices application using gRPC server with protocbuf communication

# Start mysql database & phpmyadmin:
  
 # docker-compose up

# Build project:

 # Generate protoc:
   # cd balance/
   # sudo chmod 755 protoc-gen.sh
   # ./protoc-gen.sh
  
 # Build balance:
   # cd balance/cmd/server/
   # go build .
   # ./server
  
 # Build transaction:
   # cd transaction/cmd/server/
   # go build .
   # ./server -id=<> -accountId=<> -description=<> -currency=<> -notes=<> -amount=<> -param=<crediter / debiter / getAmount>
   # ex:  ./server -id=1 -accountId=1 -description="paiement x" -currency=euro -notes="ceci est un test" -amount=20 -param=crediter
    
This is my first version more update are comming.
