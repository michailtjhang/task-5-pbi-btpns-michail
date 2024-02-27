# go-photoUser

Projek Photo User Golang API

## Studi Kasus

Berdasarkan data yang telah diolah oleh tim Data Analysts, bahwa untuk meningkatkan engagement user pada aplikasi m-banking adalah meningkatkan aspek memiliki user pada aplikasi tersebut. Saran yang diberikan oleh tim data analysts adalah membentuk fitur personalize user, salah satunya adalah memungkinkan user dapat mengupload gambar untuk dijadikan foto profilnya. Tim developer bertanggung jawab untuk mengembangkan fitur ini, dan kalian diberikan tugas untuk merancang API pada fitur upload, dan menghapus gambar. Beberapa
ketentuannya antara lain :
- User dapat menambahkan foto profile
- Sistem dapat mengidentifikasi User ( log in / sign up)
- Hanya user yang telah login / sign up yang dapat melakukan delete / tambah foto profil
- User dapat menghapus gambar yang telah di post
- User yang berbeda tidak dapat menghapus / mengubah foto yang telah di buat oleh user lain

## Fitur
User Endpoint :
1. POST : /users/register, dan gunakan atribut berikut ini
- ID (primary key, required)
- Username (required)
- Email (unique & required)
- Password (required & minlength 6)
- Relasi dengan model Photo (Gunakan constraint cascade)
- Created At (timestamp)
- Updated At (timestamp)
2. POST: /users/login http://localhost/users/login
- Using email & password (required)
3. PUT : /users/:userId (Update User)
4. DELETE : /users/:userId (Delete User)

Photos Endpoint
1. POST : /photos
- ID (primary key, required)
- Title
- Caption
- PhotoUrl
- UserID
- Relasi dengan model User
2. GET : /photos
3. PUT : /photos/:photoId
4. DELETE : /photos/:photoId

## Tools 
- Gin Gonic Framework : https://github.com/gin-gonic/gin
- Gorm : https://gorm.io/index.html
- JWT Go : https://github.com/golang-jwt/jwt
- Go Validator : http://github.com/asaskevich/govalidator
- Database :  Server SQL open source (MySQL)

command run Application 
```bash
go run main.go
```