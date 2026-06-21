# Frist RESTFUL API - Go *(zero AI)*

> Learning Go by building a API

## Routes

```sh
# BASE URL
/api/v1

# GET
/notes/all     # get all notes
/notes/:noteId # get a note

# POST
/notes/create # create note

# PUT
/notes/:noteId/update # update full note data

# PATCH
/notes/:noteId/updateTitle       # update note title
/notes/:noteId?description="" # update note description
/notes/:noteId?color=""       # update note color

# DELETE
/notes/:noteId # delete a note
```