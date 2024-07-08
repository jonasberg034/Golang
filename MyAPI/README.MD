````Bash

https://go.dev/doc/tutorial/web-service-gin



## getAlbums responds with the list of all albums as JSON.

func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}


In this code, you:

* Write a getAlbums function that takes a gin.Context parameter. Note that you could have given this function any name – neither Gin nor Go require a particular function name format.

* gin.Context is the most important part of Gin. It carries request details, validates and serializes JSON, and more. (Despite the similar name, this is different from Go’s built-in context package.)

* Call Context.IndentedJSON to serialize the struct into JSON and add it to the response.

* The function’s first argument is the HTTP status code you want to send to the client. Here, you’re passing the StatusOK constant from the net/http package to indicate 200 OK.

* Note that you can replace Context.IndentedJSON with a call to Context.JSON to send more compact JSON. In practice, the indented form is much easier to work with when debugging and the size difference is usually small.

````

````Bash

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)

    router.Run("localhost:8080")
}

In this code, you:

* Initialize a Gin router using Default.

* Use the GET function to associate the GET HTTP method and /albums path with a handler function.

* Note that you’re passing the name of the getAlbums function. This is different from passing the result of the function, which you would do by passing getAlbums() (note the parenthesis).

* Use the Run function to attach the router to an http.Server and start the server.

````

````Bash


func postAlbums(c *gin.Context) {
    var newAlbum album

    ## Call BindJSON to bind the received JSON to newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    ## Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

In this code, you:

* Use Context.BindJSON to bind the request body to newAlbum.
* Append the album struct initialized from the JSON to the albums slice.
* Add a 201 status code to the response, along with JSON representing the album you added.


````

````Bash
##getAlbumByID locates the album whose ID value matches the id
## parameter sent by the client, then returns that album as a response.

func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}


In this code, you:

* Use Context.Param to retrieve the id path parameter from the URL. When you map this handler to a path, you’ll include a placeholder for the parameter in the path.

* Loop over the album structs in the slice, looking for one whose ID field value matches the id parameter value. If it’s found, you serialize that album struct to JSON and return it as a response with a 200 OK HTTP code.

* As mentioned above, a real-world service would likely use a database query to perform this lookup.

* Return an HTTP 404 error with http.StatusNotFound if the album isn’t found.

````