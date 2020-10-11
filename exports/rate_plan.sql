CREATE TABLE IF NOT EXISTS rate_plans(
   hotel_id     VARCHAR(8) NOT NULL PRIMARY KEY
  ,rate_plan_id VARCHAR(3) NOT NULL
  ,name         VARCHAR(19) NOT NULL
  ,meal_plan    VARCHAR(9) NOT NULL
);
INSERT INTO rate_plans(hotel_id,rate_plan_id,name,meal_plan) VALUES ('BH~46455','BAR','BEST AVAILABLE RATE','Room only');
