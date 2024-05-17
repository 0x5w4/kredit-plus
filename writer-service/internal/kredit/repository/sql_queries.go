package repository

const (
	createKonsumenQuery = `
	INSERT INTO konsumens (
		id_konsumen,
		nik, 
		full_name,
		legal_name,
		gaji,
		tempat_lahir,
		tanggal_lahir,
		foto_ktp,
		foto_selfie,
		email,
		password,
		created_at, 
		updated_at
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8,
		$9,
		$10,
		$11,
		now(), 
		now()
	) RETURNING 
		id_konsumen,
		nik, 
		full_name,
		legal_name,
		gaji,
		tempat_lahir,
		tanggal_lahir,
		foto_ktp,
		foto_selfie,
		email,
		password,
		created_at, 
		updated_at
	`

	createLimitQuery = `
	INSERT INTO limits (
		id_limit,
		id_konsumen, 
		tenor,
		batas_kredit,
		created_at, 
		updated_at
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		now(), 
		now()
	) RETURNING 
		id_limit,
		id_konsumen, 
		tenor,
		batas_kredit,
		created_at, 
		updated_at
	`

	createTransaksiQuery = `
	INSERT INTO transaksis (
		id_transaksi,
		id_konsumen, 
		nomor_kontrak,
		tanggal_transaksi,
		otr,
		admin_fee,
		jumlah_cicilan,
		jumlah_bunga,
		nama_asset,
		jenis_transaksi,
		created_at, 
		updated_at
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8,
		$9,
		$10,
		now(), 
		now()
	) RETURNING 
		id_transaksi,
		id_konsumen, 
		nomor_kontrak,
		tanggal_transaksi,
		otr,
		admin_fee,
		jumlah_cicilan,
		jumlah_bunga,
		nama_asset,
		jenis_transaksi,
		created_at, 
		updated_at
	`

	getLimitQuery = `
	SELECT 
		l.id_limit,
		l.id_konsumen,
		l.tenor,
		l.batas_kredit
	FROM 
		limits l 
	WHERE 
		l.id_limit = $1
		AND l.id_konsumen = $2
	`

	getTransaksiQuery = `
	SELECT 
		t.id_transaksi,
		t.id_konsumen,
		t.nomor_kontrak,
		t.tanggal_transaksi,
		t.otr,
		t.admin_fee,
		t.jumlah_cicilan,
		t.jumlah_bunga,
		t.nama_asset,
		t.jenis_transaksi,
	FROM 
		transaksis t 
	WHERE 
		t.id_transaksi = $1
		AND t.id_konsumen = $2
	`
)
