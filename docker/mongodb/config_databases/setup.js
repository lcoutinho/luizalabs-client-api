db = db.getSiblingDB('customer_db');

db.createUser({
    user: 'dev',
    pwd: 'dev',
    roles: [
        { role: 'dbOwner', db: 'customer_db' }
    ]
});

db.createCollection('products');
db.createCollection('customers');
db.createCollection('users');
db.customers.createIndex( { "email": 1 }, { unique: true } );
db.users.insert({"username":"admin", "password":"luizalabs"});

db = db.getSiblingDB('customer_db_test');

db.createUser({
    user: 'test',
    pwd: 'test',
    roles: [
        { role: 'dbOwner', db: 'customer_db_test' }
    ]
});

db.createCollection('products');
db.createCollection('customers');
db.createCollection('users');
db.customers.createIndex( { "email": 1 }, { unique: true } );
db.users.insert({"username":"admin", "password":"luizalabs"});

