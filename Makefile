mongod: ./data
	mongod --dbpath ./data \
		--smallfiles --storageEngine wiredTiger

./data:
	mkdir -p ./data

distclean:
	rm -rf ./data
