export CORE_PEER_TLS_ENABLED=true
export ORDERER_CA=${PWD}/artifacts/channel/crypto-config/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
export PEER0_VAUTHORITY1_CA=${PWD}/artifacts/channel/crypto-config/peerOrganizations/vauthority1.example.com/peers/peer0.vauthority1.example.com/tls/ca.crt
export PEER0_vauthority2_CA=${PWD}/artifacts/channel/crypto-config/peerOrganizations/vauthority2.example.com/peers/peer0.vauthority2.example.com/tls/ca.crt
export PEER0_vauthority3_CA=${PWD}/artifacts/channel/crypto-config/peerOrganizations/vauthority3.example.com/peers/peer0.vauthority3.example.com/tls/ca.crt
# export PEER0_voter_CA=${PWD}/artifacts/channel/crypto-config/peerOrganizations/voter.example.com/peers/peer0.voter.example.com/tls/ca.crt
# export PEER0_vcms_CA=${PWD}/artifacts/channel/crypto-config/peerOrganizations/vcms.example.com/peers/peer0.vcms.example.com/tls/ca.crt
export FABRIC_CFG_PATH=${PWD}/artifacts/channel/config/

export CHANNEL_NAME=mychannel

# set -x

setGlobalsForOrderer(){
    export CORE_PEER_LOCALMSPID="OrdererMSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/artifacts/channel/crypto-config/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
    export CORE_PEER_MSPCONFIGPATH=${PWD}/artifacts/channel/crypto-config/ordererOrganizations/example.com/users/Admin@example.com/msp
    
}

setGlobalsForPeer0vauthority1(){
    export CORE_PEER_LOCALMSPID="vauthority1MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_VAUTHORITY1_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/artifacts/channel/crypto-config/peerOrganizations/vauthority1.example.com/users/Admin@vauthority1.example.com/msp
    export CORE_PEER_ADDRESS=localhost:7051
}

setGlobalsForPeer1vauthority1(){
    export CORE_PEER_LOCALMSPID="vauthority1MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_VAUTHORITY1_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/artifacts/channel/crypto-config/peerOrganizations/vauthority1.example.com/users/Admin@vauthority1.example.com/msp
    export CORE_PEER_ADDRESS=localhost:8051
    
}

setGlobalsForPeer0vauthority2(){
    export CORE_PEER_LOCALMSPID="vauthority2MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_vauthority2_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/artifacts/channel/crypto-config/peerOrganizations/vauthority2.example.com/users/Admin@vauthority2.example.com/msp
    export CORE_PEER_ADDRESS=localhost:9051
    
}

setGlobalsForPeer1vauthority2(){
    export CORE_PEER_LOCALMSPID="vauthority2MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_vauthority2_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/artifacts/channel/crypto-config/peerOrganizations/vauthority2.example.com/users/Admin@vauthority2.example.com/msp
    export CORE_PEER_ADDRESS=localhost:10051
    
}

setGlobalsForPeer0vauthority3(){
    export CORE_PEER_LOCALMSPID="vauthority3MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_vauthority3_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/artifacts/channel/crypto-config/peerOrganizations/vauthority3.example.com/users/Admin@vauthority3.example.com/msp
    export CORE_PEER_ADDRESS=localhost:11051
    
}

setGlobalsForPeer1vauthority3(){
    export CORE_PEER_LOCALMSPID="vauthority3MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_vauthority3_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/artifacts/channel/crypto-config/peerOrganizations/vauthority3.example.com/users/Admin@vauthority3.example.com/msp
    export CORE_PEER_ADDRESS=localhost:12051
    
}

setGlobalsForPeer0voter(){
    export CORE_PEER_LOCALMSPID="voterMSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_voter_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/artifacts/channel/crypto-config/peerOrganizations/voter.example.com/users/Admin@voter.example.com/msp
    export CORE_PEER_ADDRESS=localhost:13051
    
}

setGlobalsForPeer1voter(){
    export CORE_PEER_LOCALMSPID="voterMSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_voter_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/artifacts/channel/crypto-config/peerOrganizations/voter.example.com/users/Admin@voter.example.com/msp
    export CORE_PEER_ADDRESS=localhost:14051
    
}

setGlobalsForPeer0vcms(){
    export CORE_PEER_LOCALMSPID="vcmsMSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_vcms_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/artifacts/channel/crypto-config/peerOrganizations/vcms.example.com/users/Admin@vcms.example.com/msp
    export CORE_PEER_ADDRESS=localhost:15051
    
}

setGlobalsForPeer1vcms(){
    export CORE_PEER_LOCALMSPID="vcmsMSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_vcms_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/artifacts/channel/crypto-config/peerOrganizations/vcms.example.com/users/Admin@vcms.example.com/msp
    export CORE_PEER_ADDRESS=localhost:16051
    
}

createChannel(){
    rm -rf ./channel-artifacts/*
    setGlobalsForPeer0vauthority1
    
    peer channel create -o localhost:7050 -c $CHANNEL_NAME \
    --ordererTLSHostnameOverride orderer.example.com \
    -f ./artifacts/channel/${CHANNEL_NAME}.tx --outputBlock ./channel-artifacts/${CHANNEL_NAME}.block \
    --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA
}

removeOldCrypto(){
    rm -rf ./api-1.4/crypto/*
    rm -rf ./api-1.4/fabric-client-kv-vauthority1/*
    rm -rf ./api-2.0/vauthority1-wallet/*
    rm -rf ./api-2.0/vauthority2-wallet/*
}


joinChannel(){
    setGlobalsForPeer0vauthority1
    peer channel join -b ./channel-artifacts/$CHANNEL_NAME.block
    
    setGlobalsForPeer0vauthority2
    peer channel join -b ./channel-artifacts/$CHANNEL_NAME.block
    
    setGlobalsForPeer0vauthority3
    peer channel join -b ./channel-artifacts/$CHANNEL_NAME.block
       
}

updateAnchorPeers(){
    setGlobalsForPeer0vauthority1
    peer channel update -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com -c $CHANNEL_NAME -f ./artifacts/channel/${CORE_PEER_LOCALMSPID}anchors.tx --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA
    
    setGlobalsForPeer0vauthority2
    peer channel update -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com -c $CHANNEL_NAME -f ./artifacts/channel/${CORE_PEER_LOCALMSPID}anchors.tx --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA
    
    setGlobalsForPeer0vauthority3
    peer channel update -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com -c $CHANNEL_NAME -f ./artifacts/channel/${CORE_PEER_LOCALMSPID}anchors.tx --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA
    
}

# removeOldCrypto

createChannel
joinChannel
updateAnchorPeers