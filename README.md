# Go Developer SAST Group - Home Assignment

This a simple application exposing CRUD endpoints through a REST API which creates, updates, deletes and retrieves items from mongodb table named ```people```  

To start the application please make sure you have Docker service running.

```sudo docker build . -t mend-crud-go```

followed by

```docker run -p 8080:8080 mend-crud-go```

If the application started without any issues the following message should appear:
```db connected```

The server should run on localhost on port 8080.

Now you can send requests to the application:

*```http://localhost:8080/people``` GET -  all items from the table.  
*```http://localhost:8080/people/{id}``` GET -  one person by id.  
*```http://localhost:8080/person``` POST - Create a new person.  
*```http://localhost:8080/people/{firstname}``` PUT - updates person's last name by first name
*```http://localhost:8080/people/{firstname}``` DELETE - deletes a person by first name

Please let me know if you have any questions or issues.

Thanks

Tamir Mayblat
