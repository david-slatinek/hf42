USE payment_db;

SET FOREIGN_KEY_CHECKS = 0;
DROP TABLE IF EXISTS payment;
SET FOREIGN_KEY_CHECKS = 1;

CREATE TABLE payment
(
    id_payment   INT            NOT NULL AUTO_INCREMENT,
    order_id     VARCHAR(255)   NOT NULL,
    user_id      VARCHAR(255)   NOT NULL,
    amount       DECIMAL(10, 2) NOT NULL,
    payment_date DATETIME       NOT NULL,
    PRIMARY KEY (id_payment)
);
