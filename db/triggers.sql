DROP TRIGGER IF EXISTS update_balance_on_in_transaction;
DROP TRIGGER IF EXISTS update_balance_on_out_transaction;

CREATE TRIGGER IF NOT EXISTS update_balance_on_in_transaction
AFTER INSERT ON in_transactions
FOR EACH ROW
BEGIN
  UPDATE players
    SET balance = balance + NEW.amount
    WHERE NEW.is_arrived = 1
      AND username = NEW."to";

  UPDATE players
    SET fake_balance = fake_balance + NEW.amount
    WHERE NEW.is_arrived = 0
      AND username = NEW."to";
END;

CREATE TRIGGER update_balance_on_out_transaction
AFTER INSERT ON out_transactions
FOR EACH ROW
BEGIN
  UPDATE players
    SET balance = balance - NEW.amount
    WHERE NEW.is_arrived = 1
      AND username = NEW."from";

  UPDATE players
    SET fake_balance = fake_balance + NEW.amount
    WHERE NEW.is_arrived = 0
      AND username = NEW."from"
      AND NEW."from" <> NEW."to";

  UPDATE players
    SET fake_balance = fake_balance - NEW.amount
    WHERE NEW.is_arrived = 0
      AND username = NEW."to";
END;
