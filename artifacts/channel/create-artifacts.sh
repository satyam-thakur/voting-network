
# chmod -R 0755 ./crypto-config
# # Delete existing artifacts
# rm -rf ./crypto-config
# rm genesis.block mychannel.tx
# rm -rf ../../channel-artifacts/*

# #Generate Crypto artifactes for organizations
# cryptogen generate --config=./crypto-config.yaml --output=./crypto-config/



# System channel
SYS_CHANNEL="sys-channel"

# channel name defaults to "mychannel"
CHANNEL_NAME_BALLOT="ballot-channel"
CHANNEL_NAME_MAP="map-channel"

echo $CHANNEL_NAME_BALLOT
echo $CHANNEL_NAME_MAP

# # Generate System Genesis block
# configtxgen -profile OrdererGenesis -configPath . -channelID $SYS_CHANNEL  -outputBlock ./genesis.block


# Generate channel configuration block
configtxgen -profile BasicChannel -configPath . -outputCreateChannelTx ./${CHANNEL_NAME_BALLOT}.tx -channelID $CHANNEL_NAME_BALLOT
configtxgen -profile BasicChannel -configPath . -outputCreateChannelTx ./${CHANNEL_NAME_MAP}.tx -channelID $CHANNEL_NAME_MAP

echo "#######    Generating anchor peer update for vauthority1MSP  ##########"
configtxgen -profile BasicChannel -configPath . -outputAnchorPeersUpdate ./vauthority1MSPanchors.tx -channelID $CHANNEL_NAME_BALLOT -asOrg vauthority1MSP
configtxgen -profile BasicChannel -configPath . -outputAnchorPeersUpdate ./vauthority1MSPanchors1.tx -channelID $CHANNEL_NAME_MAP -asOrg vauthority1MSP

echo "#######    Generating anchor peer update for vauthority2MSP  ##########"
configtxgen -profile BasicChannel -configPath . -outputAnchorPeersUpdate ./vauthority2MSPanchors.tx -channelID $CHANNEL_NAME_BALLOT -asOrg vauthority2MSP
configtxgen -profile BasicChannel -configPath . -outputAnchorPeersUpdate ./vauthority2MSPanchors1.tx -channelID $CHANNEL_NAME_MAP -asOrg vauthority2MSP

echo "#######    Generating anchor peer update for vauthority3MSP  ##########"
configtxgen -profile BasicChannel -configPath . -outputAnchorPeersUpdate ./vauthority3MSPanchors.tx -channelID $CHANNEL_NAME_BALLOT -asOrg vauthority3MSP
configtxgen -profile BasicChannel -configPath . -outputAnchorPeersUpdate ./vauthority3MSPanchors1.tx -channelID $CHANNEL_NAME_MAP -asOrg vauthority3MSP