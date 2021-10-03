CREATE TABLE IF NOT EXISTS banner (banner_id SERIAL PRIMARY KEY, banner_name TEXT, banner_url TEXT, banner_imgsrc TEXT, banner_startdate INT,banner_enddate INT, banner_maxtier INT, banner_mintier INT,banner_minbalance INT,banner_maxbalance INT,banner_mintokopoint INT,banner_maxtokopoint INT,banner_isactive BIT);

INSERT INTO banner ( banner_name, banner_url, banner_imgsrc, banner_startdate,banner_enddate, banner_maxtier, banner_mintier,banner_minbalance ,banner_maxbalance,
banner_mintokopoint,banner_maxtokopoint) VALUES('Waktu Indonesia Belanja', 'http://tokopedia.com', 'wib.jpg', 250120, 300120,'silver','gold',10,1000000,1,100000000);

INSERT INTO banner ( banner_name, banner_url, banner_imgsrc, banner_startdate,banner_enddate, banner_maxtier, banner_mintier,banner_minbalance ,banner_maxbalance,
banner_mintokopoint,banner_maxtokopoint) VALUES('Cashback Barang Elektronik', 'http://tokopediaelek.com', 'kom.jpg', 130120, 300120,'silver','gold',300,1000000,100,100000000);

INSERT INTO banner ( banner_name, banner_url, banner_imgsrc, banner_startdate,banner_enddate, banner_maxtier, banner_mintier,banner_minbalance ,banner_maxbalance,
banner_mintokopoint,banner_maxtokopoint) VALUES('Cashback OVO', 'http://tokopediaovo.com', 'ovo.jpg', 150120, 200120,'silver','gold',300,1000000,100,100000000);

INSERT INTO banner ( banner_name, banner_url, banner_imgsrc, banner_startdate,banner_enddate, banner_maxtier, banner_mintier,banner_minbalance ,banner_maxbalance,
banner_mintokopoint,banner_maxtokopoint) VALUES('Hari Skincare Indonesia', 'http://tokopediahsi.com', 'hsi.jpg', 150120, 250120,'silver','gold',300,1000000,100,100000000);


CREATE TABLE IF NOT EXISTS users (user_id SERIAL PRIMARY KEY, user_name TEXT, user_email TEXT, user_balance INT, user_tokopoint INT, user_tier TEXT, user_location TEXT, user_bannerlist JSON);

INSERT INTO user (user_name, user_email, user_balance, user_tokopoint, user_tier, user_location, user_bannerlist) 
    VALUES('Rizky', "rizkycintadamai@gmail.com", 50000, 130000, 'Gold', 'Jakarta', '[1]');

INSERT INTO user (user_name, user_email, user_balance, user_tokopoint, user_tier, user_location, user_bannerlist) 
    VALUES('Agung', "agungcintadamai@gmail.com", 3000, 109000, 'Gold', 'Balikpapan', '[3]');

INSERT INTO user (user_name, user_email, user_balance, user_tokopoint, user_tier, user_bannerlist) 
    VALUES('Bagas', "bagasanakmotor@gmail.com", 10, 70000, 'Silver', 'Surabaya', '[2]');

INSERT INTO user (user_name, user_email, user_balance, user_tokopoint, user_tier, user_bannerlist) 
    VALUES('Asep', "asepmanurung@gmail.com", 67100, 1090, 'Silver', 'Depok', '[2]');

INSERT INTO user (user_name, user_email, user_balance, user_tokopoint, user_tier, user_bannerlist) 
    VALUES('Udin', "udinbaskara@gmail.com", 54520, 115020, 'Gold', 'Bandung', '[4]');



