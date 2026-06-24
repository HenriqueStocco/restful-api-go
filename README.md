# Frist RESTFUL API - Go *(zero AI)*

> Learning Go by building a API

## Routes

>[!NOTE]
>
> <del>O Go não me permite utilizar a mesma rota /notes para diferentes metódos em HandleFunc diferentes, então há duas alternativas:</del>
> 
> <del>1. Usar uma função Handler para tratar os diferentes metodos [GET, POST, PUT, PATCH, DELETE] *(acho que o handler teria muitas responsabilidades)*</del>
> 
> <del>2. Criar varias rotas adicionando path param + query param para diferenciar *(opção escolhida)*</del>
> 
> Permite desde que passemos o método no parametro pattern do handleFunc, e.g., POST /notes.

```sh
# BASE URL
/api/v1

# GET
/note          # get all notes
/note/{noteId} # get a note

# POST
/note # create note

# PUT
/note/{noteId} # update full note data

# PATCH
/note/:noteId/title       # update note title
/note/:noteId/description # update note description
/note/:noteId/color       # update note color

# DELETE
/note/{noteId} # delete a note
/note          # delete all notes
```