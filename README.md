# project-movies

Implemented a simple movies server that can perform all the basic CRUD operations
    
    CREATE: /add_movie adds a movie into the database; returns the created object after insertion.
    READ: /get_movies returns list of all available movies in the database
          /get_movies/{isbn} returns the desired record, uniquely identified by isbn
    UPDATE: /update_movie updates the desired record and returns it
    DELETE: /delete_movie deletes the desired record from the database and returns the status and isbn of the deleted record.
    

Implemented a GRPC gateway as well, to translate the Restful HTTP API calls into RPC calls twhich are then sent sent to the grPC server.

All DB and server details are in the app.env file
