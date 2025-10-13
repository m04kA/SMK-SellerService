package priceservice

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client клиент для работы с PriceService
type Client struct {
	baseURL    string
	httpClient *http.Client
	log        Logger
}

// NewClient создает новый экземпляр клиента PriceService
func NewClient(baseURL string, log Logger) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		log: log,
	}
}

// CalculatePrices вызывает endpoint /api/v1/prices/calculate для расчёта цен
func (c *Client) CalculatePrices(ctx context.Context, req *CalculatePricesRequest) (*CalculatePricesResponse, error) {
	url := fmt.Sprintf("%s/api/v1/prices/calculate", c.baseURL)

	// Маршалим тело запроса
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("%w: failed to marshal request: %v", ErrInternal, err)
	}

	// Создаём HTTP запрос
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("%w: failed to create request: %v", ErrInternal, err)
	}

	httpReq.Header.Set("Content-Type", "application/json")

	// Выполняем запрос
	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("%w: failed to execute request: %v", ErrInternal, err)
	}
	defer resp.Body.Close()

	// Обработка статус-кодов
	switch resp.StatusCode {
	case http.StatusOK:
		// Продолжаем обработку
	case http.StatusBadRequest:
		return nil, fmt.Errorf("%w: bad request", ErrInvalidResponse)
	case http.StatusNotFound:
		return nil, ErrPricesNotFound
	default:
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("%w: unexpected status code %d: %s", ErrInvalidResponse, resp.StatusCode, string(body))
	}

	// Парсим ответ
	var pricesResp CalculatePricesResponse
	if err := json.NewDecoder(resp.Body).Decode(&pricesResp); err != nil {
		return nil, fmt.Errorf("%w: failed to decode response: %v", ErrInvalidResponse, err)
	}

	return &pricesResp, nil
}

// CalculatePricesWithGracefulDegradation вызывает расчёт цен с graceful degradation
// При недоступности PriceService возвращает ErrServiceDegraded, что позволяет сервису вернуть данные без цен
func (c *Client) CalculatePricesWithGracefulDegradation(ctx context.Context, req *CalculatePricesRequest) (*CalculatePricesResponse, error) {
	if req.UserID != nil {
		c.log.Info("Calculating prices for company_id=%d, user_id=%d, services=%v", req.CompanyID, *req.UserID, req.ServiceIDs)
	} else {
		c.log.Info("Calculating base prices for company_id=%d, services=%v", req.CompanyID, req.ServiceIDs)
	}

	prices, err := c.CalculatePrices(ctx, req)
	if err != nil {
		// Если это критичная бизнес-ошибка (не найдены цены),
		// пробрасываем её дальше
		if err == ErrPricesNotFound {
			c.log.Info("Prices not found for company_id=%d, services=%v", req.CompanyID, req.ServiceIDs)
			return nil, err
		}

		// Для всех остальных ошибок (недоступность сервиса, timeout, ошибки парсинга и т.д.)
		// применяем graceful degradation - возвращаем ErrServiceDegraded с контекстом
		// Повышаем уровень логирования до ERROR, чтобы быстрее заметить проблему
		c.log.Error("PriceService unavailable, applying graceful degradation for company_id=%d: %v", req.CompanyID, err)
		return nil, fmt.Errorf("%w: company_id=%d, error=%v", ErrServiceDegraded, req.CompanyID, err)
	}

	c.log.Info("Successfully calculated prices for company_id=%d, count=%d", req.CompanyID, len(prices.Prices))
	return prices, nil
}
