rm -R crypto-config/*

./bin/cryptogen generate --config=crypto-config.yaml

rm config/*

./bin/configtxgen -profile spitrackOrgOrdererGenesis -outputBlock ./config/genesis.block

./bin/configtxgen -profile spitrackOrgChannel -outputCreateChannelTx ./config/spitrackchannel.tx -channelID spitrackchannel
