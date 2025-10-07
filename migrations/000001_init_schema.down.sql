-- Удаляем триггеры
DROP TRIGGER IF EXISTS update_services_updated_at ON services;
DROP TRIGGER IF EXISTS update_working_hours_updated_at ON working_hours;
DROP TRIGGER IF EXISTS update_addresses_updated_at ON addresses;
DROP TRIGGER IF EXISTS update_companies_updated_at ON companies;

-- Удаляем функцию
DROP FUNCTION IF EXISTS update_updated_at_column();

-- Удаляем таблицы (в обратном порядке из-за внешних ключей)
DROP TABLE IF EXISTS service_addresses;
DROP TABLE IF EXISTS services;
DROP TABLE IF EXISTS working_hours;
DROP TABLE IF EXISTS addresses;
DROP TABLE IF EXISTS companies;
