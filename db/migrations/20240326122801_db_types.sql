-- Create transaction_type enum with values (payment, settlement)
CREATE TYPE transaction_type AS ENUM ('payment', 'settlement');

-- Create transaction_user_type enum with values (payer, payee)
CREATE TYPE transaction_user_type AS ENUM ('payer', 'payee');
