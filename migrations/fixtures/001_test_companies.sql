-- ==========================================
-- SellerService - Тестовые фикстуры
-- ==========================================
-- Создаёт тестовые компании для интеграции с BookingService
-- Соответствует данным из BookingService фикстур

-- ==========================================
-- КОМПАНИЯ 1: Автомойка Премиум
-- ==========================================

-- Компания
INSERT INTO companies (id, name, logo, description, tags, manager_ids)
VALUES (
    1,
    'Автомойка Премиум',
    'https://storage.example.com/logos/premium-wash.png',
    'Профессиональная автомойка и детейлинг в центре Москвы. Более 10 лет на рынке, современное оборудование, опытные мастера.',
    ARRAY['#мойка', '#детейлинг', '#москва', '#премиум'],
    ARRAY[777777777]
)
ON CONFLICT (id) DO UPDATE SET
    name = EXCLUDED.name,
    logo = EXCLUDED.logo,
    description = EXCLUDED.description,
    tags = EXCLUDED.tags,
    manager_ids = EXCLUDED.manager_ids,
    updated_at = NOW();

-- Адрес 1: Тверская
INSERT INTO addresses (id, company_id, city, street, building, latitude, longitude)
VALUES (
    100,
    1,
    'Москва',
    'ул. Тверская',
    '10',
    55.755826,
    37.617299
)
ON CONFLICT (id) DO UPDATE SET
    city = EXCLUDED.city,
    street = EXCLUDED.street,
    building = EXCLUDED.building,
    latitude = EXCLUDED.latitude,
    longitude = EXCLUDED.longitude,
    updated_at = NOW();

-- Адрес 2: Ленина
INSERT INTO addresses (id, company_id, city, street, building, latitude, longitude)
VALUES (
    101,
    1,
    'Москва',
    'ул. Ленина',
    '5',
    55.751244,
    37.618423
)
ON CONFLICT (id) DO UPDATE SET
    city = EXCLUDED.city,
    street = EXCLUDED.street,
    building = EXCLUDED.building,
    latitude = EXCLUDED.latitude,
    longitude = EXCLUDED.longitude,
    updated_at = NOW();

-- Рабочие часы: Пн-Пт 09:00-21:00, Сб 10:00-20:00, Вс выходной
INSERT INTO working_hours (
    company_id,
    monday_is_open, monday_open_time, monday_close_time,
    tuesday_is_open, tuesday_open_time, tuesday_close_time,
    wednesday_is_open, wednesday_open_time, wednesday_close_time,
    thursday_is_open, thursday_open_time, thursday_close_time,
    friday_is_open, friday_open_time, friday_close_time,
    saturday_is_open, saturday_open_time, saturday_close_time,
    sunday_is_open
)
VALUES (
    1,
    true, '09:00', '21:00',
    true, '09:00', '21:00',
    true, '09:00', '21:00',
    true, '09:00', '21:00',
    true, '09:00', '21:00',
    true, '10:00', '20:00',
    false
)
ON CONFLICT (company_id) DO UPDATE SET
    monday_is_open = EXCLUDED.monday_is_open,
    monday_open_time = EXCLUDED.monday_open_time,
    monday_close_time = EXCLUDED.monday_close_time,
    tuesday_is_open = EXCLUDED.tuesday_is_open,
    tuesday_open_time = EXCLUDED.tuesday_open_time,
    tuesday_close_time = EXCLUDED.tuesday_close_time,
    wednesday_is_open = EXCLUDED.wednesday_is_open,
    wednesday_open_time = EXCLUDED.wednesday_open_time,
    wednesday_close_time = EXCLUDED.wednesday_close_time,
    thursday_is_open = EXCLUDED.thursday_is_open,
    thursday_open_time = EXCLUDED.thursday_open_time,
    thursday_close_time = EXCLUDED.thursday_close_time,
    friday_is_open = EXCLUDED.friday_is_open,
    friday_open_time = EXCLUDED.friday_open_time,
    friday_close_time = EXCLUDED.friday_close_time,
    saturday_is_open = EXCLUDED.saturday_is_open,
    saturday_open_time = EXCLUDED.saturday_open_time,
    saturday_close_time = EXCLUDED.saturday_close_time,
    sunday_is_open = EXCLUDED.sunday_is_open,
    updated_at = NOW();

-- Услуга 1: Комплексная мойка (доступна на обоих адресах)
INSERT INTO services (id, company_id, name, description, average_duration)
VALUES (
    1,
    1,
    'Комплексная мойка',
    'Полная мойка кузова, дисков, ковриков и салона пылесосом. Включает сушку и протирку насухо.',
    60
)
ON CONFLICT (id) DO UPDATE SET
    name = EXCLUDED.name,
    description = EXCLUDED.description,
    average_duration = EXCLUDED.average_duration,
    updated_at = NOW();

-- Связь услуги 1 с адресами
INSERT INTO service_addresses (service_id, address_id)
VALUES (1, 100), (1, 101)
ON CONFLICT (service_id, address_id) DO NOTHING;

-- Услуга 2: Экспресс-мойка (только на адресе 100 - Тверская)
INSERT INTO services (id, company_id, name, description, average_duration)
VALUES (
    2,
    1,
    'Экспресс-мойка',
    'Быстрая мойка кузова без салона. Идеально для регулярного поддержания чистоты.',
    30
)
ON CONFLICT (id) DO UPDATE SET
    name = EXCLUDED.name,
    description = EXCLUDED.description,
    average_duration = EXCLUDED.average_duration,
    updated_at = NOW();

-- Связь услуги 2 с адресом 100
INSERT INTO service_addresses (service_id, address_id)
VALUES (2, 100)
ON CONFLICT (service_id, address_id) DO NOTHING;

-- Услуга 3: Детейлинг (только на адресе 101 - Ленина)
INSERT INTO services (id, company_id, name, description, average_duration)
VALUES (
    3,
    1,
    'Детейлинг',
    'Полный детейлинг кузова: глубокая мойка, полировка, защитное покрытие. Премиальное качество.',
    120
)
ON CONFLICT (id) DO UPDATE SET
    name = EXCLUDED.name,
    description = EXCLUDED.description,
    average_duration = EXCLUDED.average_duration,
    updated_at = NOW();

-- Связь услуги 3 с адресом 101
INSERT INTO service_addresses (service_id, address_id)
VALUES (3, 101)
ON CONFLICT (service_id, address_id) DO NOTHING;

-- ==========================================
-- КОМПАНИЯ 2: СТО Профи
-- ==========================================

-- Компания
INSERT INTO companies (id, name, logo, description, tags, manager_ids)
VALUES (
    2,
    'СТО Профи',
    'https://storage.example.com/logos/sto-profi.png',
    'Профессиональное техническое обслуживание автомобилей. Работаем круглосуточно.',
    ARRAY['#сто', '#ремонт', '#москва', '#круглосуточно'],
    ARRAY[888888888]
)
ON CONFLICT (id) DO UPDATE SET
    name = EXCLUDED.name,
    logo = EXCLUDED.logo,
    description = EXCLUDED.description,
    tags = EXCLUDED.tags,
    manager_ids = EXCLUDED.manager_ids,
    updated_at = NOW();

-- Адрес: Новая
INSERT INTO addresses (id, company_id, city, street, building, latitude, longitude)
VALUES (
    200,
    2,
    'Москва',
    'ул. Новая',
    '15',
    55.745123,
    37.623456
)
ON CONFLICT (id) DO UPDATE SET
    city = EXCLUDED.city,
    street = EXCLUDED.street,
    building = EXCLUDED.building,
    latitude = EXCLUDED.latitude,
    longitude = EXCLUDED.longitude,
    updated_at = NOW();

-- Рабочие часы: Круглосуточно
INSERT INTO working_hours (
    company_id,
    monday_is_open, monday_open_time, monday_close_time,
    tuesday_is_open, tuesday_open_time, tuesday_close_time,
    wednesday_is_open, wednesday_open_time, wednesday_close_time,
    thursday_is_open, thursday_open_time, thursday_close_time,
    friday_is_open, friday_open_time, friday_close_time,
    saturday_is_open, saturday_open_time, saturday_close_time,
    sunday_is_open, sunday_open_time, sunday_close_time
)
VALUES (
    2,
    true, '00:00', '23:59',
    true, '00:00', '23:59',
    true, '00:00', '23:59',
    true, '00:00', '23:59',
    true, '00:00', '23:59',
    true, '00:00', '23:59',
    true, '00:00', '23:59'
)
ON CONFLICT (company_id) DO UPDATE SET
    monday_is_open = EXCLUDED.monday_is_open,
    monday_open_time = EXCLUDED.monday_open_time,
    monday_close_time = EXCLUDED.monday_close_time,
    tuesday_is_open = EXCLUDED.tuesday_is_open,
    tuesday_open_time = EXCLUDED.tuesday_open_time,
    tuesday_close_time = EXCLUDED.tuesday_close_time,
    wednesday_is_open = EXCLUDED.wednesday_is_open,
    wednesday_open_time = EXCLUDED.wednesday_open_time,
    wednesday_close_time = EXCLUDED.wednesday_close_time,
    thursday_is_open = EXCLUDED.thursday_is_open,
    thursday_open_time = EXCLUDED.thursday_open_time,
    thursday_close_time = EXCLUDED.thursday_close_time,
    friday_is_open = EXCLUDED.friday_is_open,
    friday_open_time = EXCLUDED.friday_open_time,
    friday_close_time = EXCLUDED.friday_close_time,
    saturday_is_open = EXCLUDED.saturday_is_open,
    saturday_open_time = EXCLUDED.saturday_open_time,
    saturday_close_time = EXCLUDED.saturday_close_time,
    sunday_is_open = EXCLUDED.sunday_is_open,
    sunday_open_time = EXCLUDED.sunday_open_time,
    sunday_close_time = EXCLUDED.sunday_close_time,
    updated_at = NOW();

-- Услуга 10: Замена масла
INSERT INTO services (id, company_id, name, description, average_duration)
VALUES (
    10,
    2,
    'Замена масла',
    'Замена моторного масла и масляного фильтра. Используем качественные масла известных брендов.',
    45
)
ON CONFLICT (id) DO UPDATE SET
    name = EXCLUDED.name,
    description = EXCLUDED.description,
    average_duration = EXCLUDED.average_duration,
    updated_at = NOW();

-- Связь услуги 10 с адресом 200
INSERT INTO service_addresses (service_id, address_id)
VALUES (10, 200)
ON CONFLICT (service_id, address_id) DO NOTHING;

-- ==========================================
-- КОМПАНИЯ 3: Детейлинг Центр
-- ==========================================

-- Компания
INSERT INTO companies (id, name, logo, description, tags, manager_ids)
VALUES (
    3,
    'Детейлинг Центр',
    'https://storage.example.com/logos/detailing-center.png',
    'Специализированный центр детейлинга в Санкт-Петербурге. Работаем только с премиальными автомобилями.',
    ARRAY['#детейлинг', '#санкт-петербург', '#премиум', '#полировка'],
    ARRAY[999999000]
)
ON CONFLICT (id) DO UPDATE SET
    name = EXCLUDED.name,
    logo = EXCLUDED.logo,
    description = EXCLUDED.description,
    tags = EXCLUDED.tags,
    manager_ids = EXCLUDED.manager_ids,
    updated_at = NOW();

-- Адрес: Невский проспект
INSERT INTO addresses (id, company_id, city, street, building, latitude, longitude)
VALUES (
    300,
    3,
    'Санкт-Петербург',
    'Невский пр.',
    '1',
    59.934280,
    30.335099
)
ON CONFLICT (id) DO UPDATE SET
    city = EXCLUDED.city,
    street = EXCLUDED.street,
    building = EXCLUDED.building,
    latitude = EXCLUDED.latitude,
    longitude = EXCLUDED.longitude,
    updated_at = NOW();

-- Рабочие часы: Пн-Вс 08:00-22:00
INSERT INTO working_hours (
    company_id,
    monday_is_open, monday_open_time, monday_close_time,
    tuesday_is_open, tuesday_open_time, tuesday_close_time,
    wednesday_is_open, wednesday_open_time, wednesday_close_time,
    thursday_is_open, thursday_open_time, thursday_close_time,
    friday_is_open, friday_open_time, friday_close_time,
    saturday_is_open, saturday_open_time, saturday_close_time,
    sunday_is_open, sunday_open_time, sunday_close_time
)
VALUES (
    3,
    true, '08:00', '22:00',
    true, '08:00', '22:00',
    true, '08:00', '22:00',
    true, '08:00', '22:00',
    true, '08:00', '22:00',
    true, '08:00', '22:00',
    true, '08:00', '22:00'
)
ON CONFLICT (company_id) DO UPDATE SET
    monday_is_open = EXCLUDED.monday_is_open,
    monday_open_time = EXCLUDED.monday_open_time,
    monday_close_time = EXCLUDED.monday_close_time,
    tuesday_is_open = EXCLUDED.tuesday_is_open,
    tuesday_open_time = EXCLUDED.tuesday_open_time,
    tuesday_close_time = EXCLUDED.tuesday_close_time,
    wednesday_is_open = EXCLUDED.wednesday_is_open,
    wednesday_open_time = EXCLUDED.wednesday_open_time,
    wednesday_close_time = EXCLUDED.wednesday_close_time,
    thursday_is_open = EXCLUDED.thursday_is_open,
    thursday_open_time = EXCLUDED.thursday_open_time,
    thursday_close_time = EXCLUDED.thursday_close_time,
    friday_is_open = EXCLUDED.friday_is_open,
    friday_open_time = EXCLUDED.friday_open_time,
    friday_close_time = EXCLUDED.friday_close_time,
    saturday_is_open = EXCLUDED.saturday_is_open,
    saturday_open_time = EXCLUDED.saturday_open_time,
    saturday_close_time = EXCLUDED.saturday_close_time,
    sunday_is_open = EXCLUDED.sunday_is_open,
    sunday_open_time = EXCLUDED.sunday_open_time,
    sunday_close_time = EXCLUDED.sunday_close_time,
    updated_at = NOW();

-- Услуга 20: Полировка кузова
INSERT INTO services (id, company_id, name, description, average_duration)
VALUES (
    20,
    3,
    'Полировка кузова',
    'Профессиональная многоэтапная полировка кузова с защитным покрытием. Удаление царапин, восстановление блеска.',
    180
)
ON CONFLICT (id) DO UPDATE SET
    name = EXCLUDED.name,
    description = EXCLUDED.description,
    average_duration = EXCLUDED.average_duration,
    updated_at = NOW();

-- Связь услуги 20 с адресом 300
INSERT INTO service_addresses (service_id, address_id)
VALUES (20, 300)
ON CONFLICT (service_id, address_id) DO NOTHING;

-- ==========================================
-- Сброс последовательностей ID (если нужно)
-- ==========================================
-- Это гарантирует, что следующие ID будут начинаться после тестовых данных
SELECT setval('companies_id_seq', (SELECT MAX(id) FROM companies));
SELECT setval('addresses_id_seq', (SELECT MAX(id) FROM addresses));
SELECT setval('services_id_seq', (SELECT MAX(id) FROM services));

-- ==========================================
-- ИТОГО: 3 компании, 4 адреса, 5 услуг
-- ==========================================
