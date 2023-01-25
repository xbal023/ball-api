package config

const (
	g = "*Access denied 📛*\n\n"
		// FOR Denied
	FOwner = g+ "```Fitur ini khusus pemilik bot ini```"
	FAdmin = g+ "```Fitur ini di khusus kan oleh admin group```"
	FBotAdmin = g+ "```Bot tidak admin! maaf ngga bisa```"
	FGroup = g+ "```Fitur ini di khusus kan untuk di group```"
)
// Using Func [ Custom Query ]
func CustomDenied (text string) string {
	return g + "```" + text + "```"
}

const (
	q = "*Using Query 🔑*\n\n"
	// FOR QUERY | TEXT
	TPath = q+ "```Masukan Path file nya dengan benar dengan benar contoh : %s %s```"
	TLink = q+ "```Masukan link nya setelah command, contoh : %s %s```"
	TDesc = q+ "```Masukan deskripsi group nya setelah command, contoh : %s %s```"
	TName = q+ "```Masukan nama / subject group nya setelah command, contoh : %s %s```"
	TQuery = q+ "```Masukan Query nya, contoh : %s %s```"
	)
// Using Func [ Custom Query ]
func CustomQuery (text string) string {
	return q + "```" + text + "```"
}
	
const (
	e = "*Using Quoted ⤴️*\n\n"
	// FOR QUOTED MESSAGE | MEDIA 
	QDel = e+ "```Reply pesan yg ingin kamu lenyapkan```"
	QMedia = e+ "```Reply / kirim media dengan caption command```"
	)
// Using Func [ Custom Quoted ]
func CustomQuoted (text string) string {
	return e + "```" + text + "```"
}
const (
	s = "*Succes ✔️*\n\n"
	// FOR SUCCES
	GCClose = s+ "```Sukses menutup chat Group, sekarang hanya admin yg dapat mengirim pesan```"
	GCOpen = s+ "```Sukses membuka chat Group, sekarang seluruh member dapat mengirim pesan```"
	InfoOpen = s+ "```Sukses membuka edit info, Sekarang semua member dapat mengedit info group```"
	InfoClose = s+ "```Sukses menutup edit info, Sekarang hanya admin yang dapat mengedit info group```"
	GCPP = s+ "```Sukses mengganti Photo group dengan yg baru```"
	GCName = s+ "```Sukses mengganti nama group dengan yg baru```"
	GCDesc = s+ "```Sukses mengganti Deskripsi dengan yg baru```"
	)
// Using Func [ Custom success ]
func CustomSuccess (text string) string {
	return s + "```" + text + "```"
}

const (
	x = "*Error ♨️*\n\n"
	)
// Using Func [ Custom error ]
func CustomError (text string) string {
	return x + "```" + text + "```"
}
	
