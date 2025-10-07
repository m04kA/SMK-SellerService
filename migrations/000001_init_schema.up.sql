-- Таблица компаний
CREATE TABLE companies (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    logo TEXT,
    description TEXT,
    tags TEXT[] DEFAULT '{}',
    manager_ids BIGINT[] NOT NULL DEFAULT '{}',
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Индексы для поиска по тегам и фильтрации
CREATE INDEX idx_companies_tags ON companies USING GIN(tags);
CREATE INDEX idx_companies_created_at ON companies(created_at);

-- Таблица адресов компаний
CREATE TABLE addresses (
    id BIGSERIAL PRIMARY KEY,
    company_id BIGINT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
    city VARCHAR(100) NOT NULL,
    street VARCHAR(200) NOT NULL,
    building VARCHAR(50) NOT NULL,
    latitude DOUBLE PRECISION NOT NULL CHECK (latitude >= -90 AND latitude <= 90),
    longitude DOUBLE PRECISION NOT NULL CHECK (longitude >= -180 AND longitude <= 180),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Индексы для адресов
CREATE INDEX idx_addresses_company_id ON addresses(company_id);
CREATE INDEX idx_addresses_city ON addresses(city);
CREATE INDEX idx_addresses_coordinates ON addresses(latitude, longitude);

-- Таблица рабочих часов компаний
CREATE TABLE working_hours (
    id BIGSERIAL PRIMARY KEY,
    company_id BIGINT NOT NULL UNIQUE REFERENCES companies(id) ON DELETE CASCADE,

    -- Понедельник
    monday_is_open BOOLEAN NOT NULL DEFAULT false,
    monday_open_time TIME,
    monday_close_time TIME,

    -- Вторник
    tuesday_is_open BOOLEAN NOT NULL DEFAULT false,
    tuesday_open_time TIME,
    tuesday_close_time TIME,

    -- Среда
    wednesday_is_open BOOLEAN NOT NULL DEFAULT false,
    wednesday_open_time TIME,
    wednesday_close_time TIME,

    -- Четверг
    thursday_is_open BOOLEAN NOT NULL DEFAULT false,
    thursday_open_time TIME,
    thursday_close_time TIME,

    -- Пятница
    friday_is_open BOOLEAN NOT NULL DEFAULT false,
    friday_open_time TIME,
    friday_close_time TIME,

    -- Суббота
    saturday_is_open BOOLEAN NOT NULL DEFAULT false,
    saturday_open_time TIME,
    saturday_close_time TIME,

    -- Воскресенье
    sunday_is_open BOOLEAN NOT NULL DEFAULT false,
    sunday_open_time TIME,
    sunday_close_time TIME,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Индекс для быстрого поиска рабочих часов компании
CREATE INDEX idx_working_hours_company_id ON working_hours(company_id);

-- Таблица услуг
CREATE TABLE services (
    id BIGSERIAL PRIMARY KEY,
    company_id BIGINT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
    name VARCHAR(200) NOT NULL,
    description TEXT,
    average_duration INTEGER CHECK (average_duration IS NULL OR average_duration >= 1),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Индексы для услуг
CREATE INDEX idx_services_company_id ON services(company_id);
CREATE INDEX idx_services_name ON services(name);

-- Таблица связи услуг и адресов (many-to-many)
CREATE TABLE service_addresses (
    service_id BIGINT NOT NULL REFERENCES services(id) ON DELETE CASCADE,
    address_id BIGINT NOT NULL REFERENCES addresses(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

    PRIMARY KEY (service_id, address_id)
);

-- Индексы для связи
CREATE INDEX idx_service_addresses_service_id ON service_addresses(service_id);
CREATE INDEX idx_service_addresses_address_id ON service_addresses(address_id);

-- Функция для автоматического обновления updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Триггеры для автоматического обновления updated_at
CREATE TRIGGER update_companies_updated_at BEFORE UPDATE ON companies
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_addresses_updated_at BEFORE UPDATE ON addresses
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_working_hours_updated_at BEFORE UPDATE ON working_hours
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_services_updated_at BEFORE UPDATE ON services
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
