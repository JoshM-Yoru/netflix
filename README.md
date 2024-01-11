# Netflix Movie and Show Watch List

## Tasks

### Task #1 – Database
- Utilize a relational database to store the data from the CSV. [x]
  + Bonus: How might you populate the database without manually inserting the records?

### Task #2 – Web Server

#### Develop a web server to interface with the database with the following capabilities

- Can retrieve a list of all movies and shows in the database. [x]
    + Bonus: How might you increase performance of this operation?
- Can add a new movie or show to the database. [x]
- Can modify information for a movie or show. [x]
- Can remove a movie or show from the database. [x]
 
### Task #3 – Web App

#### Develop a web application to communicate with the server that includes the following features

- Allows John to see all movies and shows in a table format. [x]
    + Bonus: How might you add a searching capability to this feature? [x]
- Allows John to add a new movie or show. [x]
- Allows John to update details for a movie or show. [x]
- Allows John to delete a movie or show. [x]

## Tools Used

- Golang: https://go.dev/ Used for the backend
- Templ: https://templ.guide/ HTML templating engine using Golang
- Echo: https://echo.labstack.com/ Minimal Web Framework for Go
- HTMX: https://htmx.org/ Used to make requests to the backend and swap html elements
- Tailwind: https://tailwindcss.com/ Used for styling

### Would Like To Do
- Implement result caching for faster results
- Implement proper handling of HX-Request and HX-Current-URL headers to build the proper route rather than defaulting to '/' on page refresh
- Implement a Tailwind build step to send necessary attributes to the client rather than using a cdn

### Bugs
- Defaults to '/' on refresh
- Use of pagination with search results is broken and will default to the main results page of the catalog or watch list
