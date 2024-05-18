CREATE EXTENSION IF NOT EXISTS citext;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
DROP TABLE IF EXISTS konsumens CASCADE;
CREATE TABLE konsumens (
    id_konsumen UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nik VARCHAR(16) NOT NULL,
    full_name VARCHAR(250) NOT NULL,
    legal_name VARCHAR(250) NOT NULL,
    gaji NUMERIC NOT NULL,
    tempat lahir VARCHAR(250) NOT NULL,
    tanggal_lahir TIMESTAMP WITH TIME ZONE,
    foto_ktp VARCHAR(250) NOT NULL,
    foto_selfie VARCHAR(250) NOT NULL,
    email NUMERIC NOT NULL,
    password NUMERIC NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
DROP TABLE IF EXISTS limits CASCADE;
CREATE TABLE limits (
    id_limits UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    id_konsumen UUID NOT NULL,
    tenor NUMERIC NOT NULL,
    batas_kredit NUMERIC NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
DROP TABLE IF EXISTS transaksis CASCADE;
CREATE TABLE transaksi (
    id_transaksi UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    id_konsumen UUID NOT NULL,
    nomor_kontrak VARCHAR(250) NOT NULL,
    tanggal_transaksi TIMESTAMP NOT NULL,
    otr NUMERIC NOT NULL,
    admin_fee NUMERIC NOT NULL,
    jumlah_cicilan NUMERIC NOT NULL,
    jumlah_bunga NUMERIC NOT NULL,
    nama_asset VARCHAR(250) NOT NULL,
    jenis_transaksi VARCHAR(250) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);