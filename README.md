# Golang Mongo

GoLang + MongoDB + /gorilla/mux 

step 1:
    sh build.sh

step 2:
    docker exec -it carscrudmongo_mongodb_1 mongo admin

step 3:
    use crudmongo
    db.carsdb.insert({model:"KIA", age: 10, price:40000})
    db.carsdb.insert({model:"Audi", age: 3, price:80000})
    db.carsdb.insert({model:"KIA", age: 3, price:45000})

    db.peopledb.insert({name:"Alina", age: 30})
    db.peopledb.insert({name:"Justin", age: 20})

step 4:
    localhost:8081/
