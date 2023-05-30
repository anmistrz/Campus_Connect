const user = {
  id: Number,
  name: String,
  profilePic: String,
  email: String,
  bio: String,
  link: String,
  whatsapp: String,
  password: String,
  userType: String, //mahasiswa ,organisasi, universitas, admin
  mahasiswa: "foreignKey",
  organisasi: "foreignKey",
  universitas: "foreignKey",
};

const mahasiswa = {
  id: String,
  semester: String,
  nim: String,
  // jabatan: String,
  jurusan: "foreignKey",
  prodi: "foreignKey",
  fakultas: "foreignKey",
  universitas: "foreignKey",
};

const organisasi = {
  id: String,
  //   pengurus: String,
  universitas: "foreignKey",
};

const universitas = {
  id: String,
  namaRektor: String,
  ktpRektor: String,
  isVerified: { type: Boolean, default: false },
  alamat: String,
};

const post = {
  id: String,
  idUser: "foreignKey",
  materi: String,
  caption: String,
  jumlahLike: String,
  jumlahComment: String,
  isNews: { type: Boolean, default: false },
  createdAt: Date,
};

const comment = {
  id: String,
  idPost: "foreignKey",
  idUser: "foreignKey",
  createdAt: date,
  comment: String,
};

const likes = {
  id: String,
  idPost: "foreignKey",
  idUser: "foreignKey",
};
