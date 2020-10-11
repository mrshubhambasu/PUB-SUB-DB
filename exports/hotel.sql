CREATE TABLE IF NOT EXISTS hotels(
   hotel_id    VARCHAR(8) NOT NULL PRIMARY KEY
  ,name        VARCHAR(35) NOT NULL
  ,country     VARCHAR(2) NOT NULL
  ,address     VARCHAR(27) NOT NULL
  ,latitude    NUMERIC(9,6) NOT NULL
  ,longitude   NUMERIC(11,6) NOT NULL
  ,telephone   VARCHAR(14) NOT NULL
  ,description VARCHAR(250) NOT NULL
  ,room_count  BIT  NOT NULL
  ,currency    VARCHAR(3) NOT NULL
);
INSERT INTO hotels(hotel_id,name,country,address,latitude,longitude,telephone,description,room_count,currency) VALUES ('BH~46455','Hawthorn Suites by Wyndham Eagle CO','US','0315 Chambers Avenue, 81631',39.660193,-106.824123,'+1-970-3283000','Stay a while in beautiful mountain country at this Hawthorn Suites by Wyndham Eagle CO hotel, just off Interstate 70, only 6 miles from the Vail/Eagle Airport and close to skiing, golfing, Eagle River and great restaurants. Pets are welcome at this h',1,'USD');
