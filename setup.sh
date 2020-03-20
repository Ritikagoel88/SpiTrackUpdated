
echo "Setting up the network.."

echo "Creating channel genesis block.."

# Create the channel
docker exec -e "CORE_PEER_LOCALMSPID=ImporterMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Importer.spitrack.com/users/Admin@Importer.spitrack.com/msp" -e "CORE_PEER_ADDRESS=peer0.Importer.spitrack.com:7051" cli peer channel create -o orderer.spitrack.com:7050 -c spitrackchannel -f /etc/hyperledger/configtx/spitrackchannel.tx


sleep 5

echo "Channel genesis block created."

echo "peer0.Importer.spitrack.com joining the channel..."
# Join peer0.Importer.spitrack.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=ImporterMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Importer.spitrack.com/users/Admin@Importer.spitrack.com/msp" -e "CORE_PEER_ADDRESS=peer0.Importer.spitrack.com:7051" cli peer channel join -b spitrackchannel.block

echo "peer0.Importer.spitrack.com joined the channel"

echo "peer0.Exporter.spitrack.com joining the channel..."

# Join peer0.Exporter.spitrack.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=ExporterMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Exporter.spitrack.com/users/Admin@Exporter.spitrack.com/msp" -e "CORE_PEER_ADDRESS=peer0.Exporter.spitrack.com:7051" cli peer channel join -b spitrackchannel.block

echo "peer0.Exporter.spitrack.com joined the channel"

echo "peer0.Supplier.spitrack.com joining the channel..."
# Join peer0.Supplier.spitrack.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=SupplierMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Supplier.spitrack.com/users/Admin@Supplier.spitrack.com/msp" -e "CORE_PEER_ADDRESS=peer0.Supplier.spitrack.com:7051" cli peer channel join -b spitrackchannel.block
sleep 5

echo "peer0.Supplier.spitrack.com joined the channel"

echo "Installing spitrack chaincode to peer0.Importer.spitrack.com..."

# install chaincode
# Install code on Importer peer
docker exec -e "CORE_PEER_LOCALMSPID=ImporterMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Importer.spitrack.com/users/Admin@Importer.spitrack.com/msp" -e "CORE_PEER_ADDRESS=peer0.Importer.spitrack.com:7051" cli peer chaincode install -n spitrackcc -v 1.0 -p github.com/spitrack/go -l golang

echo "Installed spitrack chaincode to peer0.Importer.spitrack.com"

echo "Installing spitrack chaincode to peer0.Exporter.spitrack.com...."

# Install code on Exporter peer
docker exec -e "CORE_PEER_LOCALMSPID=ExporterMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Exporter.spitrack.com/users/Admin@Exporter.spitrack.com/msp" -e "CORE_PEER_ADDRESS=peer0.Exporter.spitrack.com:7051" cli peer chaincode install -n spitrackcc -v 1.0 -p github.com/spitrack/go -l golang

echo "Installed spitrack chaincode to peer0.Exporter.spitrack.com"

echo "Installing spitrack chaincode to peer0.Supplier.spitrack.com..."
# Install code on Supplier peer
docker exec -e "CORE_PEER_LOCALMSPID=SupplierMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Supplier.spitrack.com/users/Admin@Supplier.spitrack.com/msp" -e "CORE_PEER_ADDRESS=peer0.Supplier.spitrack.com:7051" cli peer chaincode install -n spitrackcc -v 1.0 -p github.com/spitrack/go -l golang

sleep 5

echo "Installed spitrack chaincode to peer0.Exporter.spitrack.com"

echo "Instantiating spitrack chaincode.."

docker exec -e "CORE_PEER_LOCALMSPID=ImporterMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Importer.spitrack.com/users/Admin@Importer.spitrack.com/msp" -e "CORE_PEER_ADDRESS=peer0.Importer.spitrack.com:7051" cli peer chaincode instantiate -o orderer.spitrack.com:7050 -C spitrackchannel -n spitrackcc -l golang -v 1.0 -c '{"Args":[""]}' -P "OR ('ImporterMSP.member','ExporterMSP.member','SupplierMSP.member')"

echo "Instantiated spitrack chaincode."

echo "Following is the docker network....."

docker ps
