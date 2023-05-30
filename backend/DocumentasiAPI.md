### USERS

- `GET` : `users?page=1&limit=10&idUserUniversitas=<id_user_universitas>&userType=<user_type>&name=<name>` => get users (idUserUniversitas,userType,name)
- `GET` : `/users?page=1&limit=10&idUserUniversitas=<id_user_universitas>&userType=<use_type>&order=users.created_at` => get users (idUserUniversitas,userType)
- `GET` : `/users?page=1&limit=1&idUserOrganisasi=3&userType=mahasiswa&order=users.created_at` => get users (idUserOrganisasi,userType)
- `GET` : `/users?page=1&limit=10&userType=universitas&isVerified=true` => get users(isVerified)
- `GET` : `/users?page=1&limit=2&name=<user_name>` => get user by name
- `GET` : `/users/:id` => get single user by id
- `POST` : `/users` => create users
- `PUT` : `/users/:id/profile` => update profile user by id
- `PUT` : `/users/:id` => update user by id
- `DELETE` : `/users/:id` => delete user by id
- `DELETE` : `/users` => delete all users

### JABATAN

- `GET` : `/jabatan`
- `POST` : `/jabatan` => create jabatan
- `PUT` : `/jabatan/:id` => update jabatan by id
- `DELETE` : `/jabatan/:id` => delete jabatan by id
- `DELETE` : `/jabatan` => delete all jabatan

### FAKULTAS

- `GET` : `/fakultas`
- `GET` : `/fakultas/:id` => get single fakultas by id
- `POST` : `/fakultas` => create fakultas
- `PUT` : `/fakultas/:id` => update fakultas by id
- `DELETE` : `/fakultas/:id` => delete fakultas by id
- `DELETE` : `/fakultas` => delete all fakultas

### PRODI

- `GET` : `/prodi?page=1&limit=10&idUserUniversitas=<id_user_universitas>&namaProdi=<name_prodi>` => get prodi (namaProdi,idUserUniversitas)
- `GET` : `/prodi?page=1&limit=10&idUserUniversitas=<id_user_universitas>` => get prodi (idUserUniversitas)
- `GET` : `/prodi/:id` => get single prodi by id
- `POST` : `/prodi` => create prodi
- `PUT` : `/prodi/:id` => update prodi by id
- `DELETE` : `/prodi/:id` => delete prodi by id
- `DELETE` : `/prodi` => delete all prodi

### POST

- `GET` : `/posts?page=1&limit=30&idUser=<id_user>`
- `GET` : `/posts?page=1&limit=10&idUser=1&isSave=true` => get post (save)
- `GET` : `/posts?page=1&limit=10&idUser=1&idUserUniversitas=<id_user_universitas>`=> get post di univ
- `GET` : `/posts?page=1&limit=30&idUser=1&isNews=true` => get all post status isNews true
- `GET` : `/posts/:id` => get single post by id
- `POST` : `/posts` => create post
- `POST` : `/posts/:id/save` => save post
- `POST` : `/posts/:id/unsave` => unsave post
- `POST` : `/posts/:id/like` => like post
- `POST` : `/posts/:id/unlike` => unlike post
- `DELETE` : `/posts/:id` => delete post by id
- `DELETE` : `/posts` => delete all post
- `GET` : `/save` => get save
- `GET` : `/like` => get like

### COMMENT

- `GET` : `/comments?page=1&limit=10&idPost=1` => get comment by idpost
- `GET` : `/comments`
- `POST` : `/comments` => create comment
- `DELETE` : `/comments/:id` => delete comment by id
