#! /bin/sh

cd $GOPATH/src/github.com/lcoutinho/luizalabs-client-api && dep ensure -v && cp -r $GOPATH/src/github.com/lcoutinho/luizalabs-client-api/schema/*json /;


until nc -z -v -w30 mongodb.lcoutinho.intranet 27017; do
 echo 'Waiting for MongoDB...'
 sleep 10
done
echo "MongoDB is up and running!"

tail -f /dev/null;
