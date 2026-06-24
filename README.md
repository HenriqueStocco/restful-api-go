# Frist RESTFUL API - Go *(zero AI)*

> Learning Go by building a API

## Routes

>[!NOTE]
>
> O Go não me permite utilizar a mesma rota /notes para diferentes metódos em HandleFunc diferentes, então há duas alternativas:
> 1. Usar uma função Handler para tratar os diferentes metodos [GET, POST, PUT, PATCH, DELETE] *(acho que o handler teria muitas responsabilidades)*
> 2. Criar varias rotas adicionando path param + query param para diferenciar *(opção escolhida)*

```sh
# BASE URL
/api/v1

# GET
/notes/all     # get all notes
/notes/{noteId} # get a note

# POST
/notes/create # create note

# PUT
/notes/update/{noteId} # update full note data

# PATCH
/notes/update/{noteId}?title=""       # update note title
/notes/update/{noteId}?description="" # update note description
/notes/update/{noteId}?color=""       # update note color

# DELETE
/notes/delete/{noteId} # delete a note
```