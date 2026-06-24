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
/note/all     # get all notes
/note/{noteId} # get a note

# POST
/note/create # create note

# PUT
/note/update/{noteId} # update full note data

# PATCH
//note/update/title/{noteId}?value=""      # update note title
/note/update/description/{noteId}?value="" # update note description
/note/update/color/{noteId}?value=""       # update note color

# DELETE
/note/delete/{noteId} # delete a note
```