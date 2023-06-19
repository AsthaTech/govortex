package govortex

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"nhooyr.io/websocket"
)

type SocketMessage struct {
	Type string            `json:"type"`
	Data SocketMessageData `json:"data"`
}
type SocketMessageData struct {
	OrderId                    string  `json:"order_id"`
	OrderNumber                string  `json:"order_number"`
	AmoOrderId                 string  `json:"amo_order_id"`
	PlacedBy                   string  `json:"placed_by"`
	ModifiedBy                 string  `json:"modified_by"`
	Status                     string  `json:"status"`
	StatusMessage              string  `json:"status_message"`
	Symbol                     string  `json:"symbol"`
	Series                     string  `json:"series"`
	InstrumentName             string  `json:"instrument_name"`
	Token                      int     `json:"token"`
	Exchange                   string  `json:"exchange"`
	ExpiryDate                 string  `json:"expiry_date"`
	StrikePrice                float32 `json:"strike_price"`
	OptionType                 string  `json:"option_type"`
	TransactionType            string  `json:"transaction_type"`
	Validity                   string  `json:"validity"`
	ValidityDays               int     `json:"validity_days"`
	Product                    string  `json:"product"`
	Variety                    string  `json:"variety"`
	DisclosedQuantity          int     `json:"disclosed_quantity"`
	DisclosedQuantityRemaining int     `json:"disclosed_quantity_remaining"`
	TotalQuantity              int     `json:"total_quantity"`
	PendingQuantity            int     `json:"pending_quantity"`
	TradedQuantity             int     `json:"traded_quantity"`
	MarketType                 string  `json:"market_type"`
	OrderPrice                 float32 `json:"order_price"`
	TriggerPrice               float32 `json:"trigger_price"`
	TradedPrice                float32 `json:"traded_price"`
	IsAmo                      bool    `json:"is_amo"`
	OrderIdentifier            string  `json:"order_identifier"`
	OrderCreatedAt             string  `json:"order_created_at"`
	OrderUpdatedAt             string  `json:"order_updated_at"`
	TradeNumber                string  `json:"trade_number,omitempty"`
	TradeTime                  string  `json:"trade_time,omitempty"`
	MarketSegmentId            int     `json:"market_segment_id"`
	GtdOrderStatus             string  `json:"gtd_order_status"`
}

type Packet struct {
	Exchange          [10]byte `json:"exchange"`
	Token             int32    `json:"token"`
	LastTradePrice    float64  `json:"last_trade_price"`
	LastTradeTime     int32    `json:"last_trade_time"`
	OpenPrice         float64  `json:"open_price"`
	HighPrice         float64  `json:"high_price"`
	LowPrice          float64  `json:"low_price"`
	ClosePrice        float64  `json:"close_price"`
	Volume            int32    `json:"volume"`
	LastUpdateTime    int32    `json:"last_update_time,omitempty"`
	LastTradeQuantity int32    `json:"last_trade_quantity,omitempty"`
	AverageTradePrice float64  `json:"average_trade_price,omitempty"`
	TotalBuyQuantity  int64    `json:"total_buy_quantity,omitempty"`
	TotalSellQuantity int64    `json:"total_sell_quantity,omitempty"`
	OpenInterest      int32    `json:"open_interest,omitempty"`
	Depth             struct {
		Buy [5]struct {
			Price    float64 `json:"price,omitempty"`
			Quantity int32   `json:"quantity,omitempty"`
			Orders   int32   `json:"orders,omitempty"`
		} `json:"buy,omitempty"`
		Sell [5]struct {
			Price    float64 `json:"price,omitempty"`
			Quantity int32   `json:"quantity,omitempty"`
			Orders   int32   `json:"orders,omitempty"`
		} `json:"sell,omitempty"`
	} `json:"depth,omitempty"`
	DPRHigh int32 `json:"dpr_high,omitempty"`
	DPRLow  int32 `json:"dpr_low,omitempty"`
}

func (p *Packet) transformToQuote() *FullQuoteData {
	q := FullQuoteData{
		Exchange:          strings.TrimRight(string(p.Exchange[:]), string([]rune{0})),
		Token:             int(p.Token),
		LastTradeTime:     int(p.LastTradeTime),
		LastUpdateTime:    int(p.LastUpdateTime),
		LastTradePrice:    p.LastTradePrice,
		Volume:            int(p.Volume),
		AverageTradePrice: p.AverageTradePrice,
		TotalBuyQuantity:  p.TotalBuyQuantity,
		TotalSellQuantity: p.TotalSellQuantity,
		OpenInterest:      int(p.OpenInterest),
		OpenPrice:         p.OpenPrice,
		HighPrice:         p.HighPrice,
		LowPrice:          p.LowPrice,
		ClosePrice:        p.ClosePrice,
		LastTradeQuantity: int(p.LastTradeQuantity),
		Depth: &QuoteDepth{
			Buy: []QuoteEntry{
				{
					Quantity: int(p.Depth.Buy[0].Quantity),
					Price:    p.Depth.Buy[0].Price,
					Orders:   int(p.Depth.Buy[0].Orders),
				},
				{
					Quantity: int(p.Depth.Buy[1].Quantity),
					Price:    p.Depth.Buy[1].Price,
					Orders:   int(p.Depth.Buy[1].Orders),
				},
				{
					Quantity: int(p.Depth.Buy[2].Quantity),
					Price:    p.Depth.Buy[2].Price,
					Orders:   int(p.Depth.Buy[2].Orders),
				},
				{
					Quantity: int(p.Depth.Buy[3].Quantity),
					Price:    p.Depth.Buy[3].Price,
					Orders:   int(p.Depth.Buy[3].Orders),
				},
				{
					Quantity: int(p.Depth.Buy[4].Quantity),
					Price:    p.Depth.Buy[4].Price,
					Orders:   int(p.Depth.Buy[4].Orders),
				},
			},
			Sell: []QuoteEntry{
				{
					Quantity: int(p.Depth.Sell[0].Quantity),
					Price:    p.Depth.Sell[0].Price,
					Orders:   int(p.Depth.Sell[0].Orders),
				},
				{
					Quantity: int(p.Depth.Sell[1].Quantity),
					Price:    p.Depth.Sell[1].Price,
					Orders:   int(p.Depth.Sell[1].Orders),
				},
				{
					Quantity: int(p.Depth.Sell[2].Quantity),
					Price:    p.Depth.Sell[2].Price,
					Orders:   int(p.Depth.Sell[2].Orders),
				},
				{
					Quantity: int(p.Depth.Sell[3].Quantity),
					Price:    p.Depth.Sell[3].Price,
					Orders:   int(p.Depth.Sell[3].Orders),
				},
				{
					Quantity: int(p.Depth.Sell[4].Quantity),
					Price:    p.Depth.Sell[4].Price,
					Orders:   int(p.Depth.Sell[4].Orders),
				},
			},
		},
		DPRHigh: float64(p.DPRHigh),
		DPRLow:  float64(p.DPRLow),
	}
	return &q
}

type Wire struct {
	Conn                *websocket.Conn
	accessToken         string
	url                 url.URL
	callbacks           callbacks
	autoreconnect       bool
	reconnectMaxDelay   time.Duration
	reconnectMaxRetries int
	connectTimeout      time.Duration

	reconnectAttempt int
	subscriptions    map[ExchangeTypes]map[int]QuoteModes
	serverContext    context.Context
	cancel           context.CancelFunc
}
type callbacks struct {
	onPriceUpdate func(*FullQuoteData)
	onMessage     func(int, []byte)
	onNoReconnect func(int)
	onReconnect   func(int, time.Duration)
	onConnect     func()
	onClose       func(int, string)
	onError       func(error)
	onOrderUpdate func(SocketMessage)
}

const (
	defaultReconnectMaxAttempts               = 300
	reconnectMinDelay           time.Duration = 5000 * time.Millisecond
	defaultReconnectMaxDelay    time.Duration = 60000 * time.Millisecond
	defaultConnectTimeout       time.Duration = 7000 * time.Millisecond
	connectionCheckInterval     time.Duration = 2000 * time.Millisecond
	dataTimeoutInterval         time.Duration = 60 * time.Second
)

var (
	websocketUrl = url.URL{Scheme: "wss", Host: "wire.asthatrade.com", Path: "ws"}
)

//Default method to create a new instance of Wire which can be used to get price updates and order updates
func NewWire(accessToken string) *Wire {
	wire := Wire{
		accessToken:         accessToken,
		url:                 websocketUrl,
		autoreconnect:       true,
		reconnectMaxDelay:   defaultReconnectMaxDelay,
		reconnectMaxRetries: defaultReconnectMaxAttempts,
		connectTimeout:      defaultConnectTimeout,
		subscriptions:       map[ExchangeTypes]map[int]QuoteModes{},
	}
	return &wire
}

// Use this function to set new url for websocket connection
func (t *Wire) SetRootURL(u url.URL) {
	t.url = u
}

// Use this function to change access token
func (t *Wire) SetAccessToken(accessToken string) {
	t.accessToken = accessToken
}

// Use this function to change connection timeout value. Dafault: 7 seconds
func (t *Wire) SetConnectTimeout(val time.Duration) {
	t.connectTimeout = val
}

// Use this function to change auto reconnection setting. Default: true
func (t *Wire) SetAutoReconnect(val bool) {
	t.autoreconnect = val
}

// Use this function to change max reconnection delay. Default: 60 seconds
func (t *Wire) SetReconnectMaxDelay(val time.Duration) error {
	if val > reconnectMinDelay {
		return fmt.Errorf("ReconnectMaxDelay can't be less than %fms", reconnectMinDelay.Seconds()*1000)
	}

	t.reconnectMaxDelay = val
	return nil
}

// Use this function to change max connection retries. Default: 300
func (t *Wire) SetReconnectMaxRetries(val int) {
	t.reconnectMaxRetries = val
}

// Set a function to receive update whenever the socket is connected
func (t *Wire) OnConnect(f func()) {
	t.callbacks.onConnect = f
}

// Set a function to receive update whenever there is an error
func (t *Wire) OnError(f func(err error)) {
	t.callbacks.onError = f
}

// Set a function to receive update whenever the socket closes
func (t *Wire) OnClose(f func(code int, reason string)) {
	t.callbacks.onClose = f
}

// Set a function to receive raw message
func (t *Wire) OnMessage(f func(messageType int, message []byte)) {
	t.callbacks.onMessage = f
}

// Set a function to receive update whenever the socket reconnects
func (t *Wire) OnReconnect(f func(attempt int, delay time.Duration)) {
	t.callbacks.onReconnect = f
}

func (t *Wire) OnNoReconnect(f func(attempt int)) {
	t.callbacks.onNoReconnect = f
}

// Set a function to receive Price Updates
func (t *Wire) OnPriceUpdate(f func(*FullQuoteData)) {
	t.callbacks.onPriceUpdate = f
}

// Set a function to receive Order Updates
func (t *Wire) OnOrderUpdate(f func(order SocketMessage)) {
	t.callbacks.onOrderUpdate = f
}
func (t *Wire) triggerNoReconnect(attempt int) {
	if t.callbacks.onNoReconnect != nil {
		t.callbacks.onNoReconnect(attempt)
	}
}
func (t *Wire) triggerMessage(messageType int, message []byte) {
	if t.callbacks.onMessage != nil {
		t.callbacks.onMessage(messageType, message)
	}
}
func (t *Wire) triggerReconnect(attempt int, delay time.Duration) {
	if t.callbacks.onReconnect != nil {
		t.callbacks.onReconnect(attempt, delay)
	}
}
func (t *Wire) triggerError(err error) {
	if t.callbacks.onError != nil {
		t.callbacks.onError(err)
	}
}
func (t *Wire) triggerConnect() {
	if t.callbacks.onConnect != nil {
		t.callbacks.onConnect()
	}
}
func (t *Wire) triggerOrderUpdate(msg SocketMessage) {
	if t.callbacks.onOrderUpdate != nil {
		t.callbacks.onOrderUpdate(msg)
	}
}

func (t *Wire) checkConnection(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			// Sleep before doing next check
			time.Sleep(connectionCheckInterval)

			pingContext, cancelFunc := context.WithTimeout(ctx, 1*time.Second)
			defer cancelFunc()
			if err := t.Conn.Ping(pingContext); err != nil {
				if t.Conn != nil {
					t.Conn.Close(websocket.StatusGoingAway, "ping not received in the time interval")
				}
				t.reconnectAttempt++
				return
			}

			// // If last ping time is greater then timeout interval then close the
			// // existing connection and reconnect
			// if time.Since(t.lastPingTime) > dataTimeoutInterval {
			// 	// Close the current connection without waiting for close frame
			// 	if t.Conn != nil {
			// 		t.Conn.Close(websocket.StatusGoingAway, "ping not received in the time interval")
			// 	}

			// 	// Increase reconnect attempt for next reconnection
			// 	t.reconnectAttempt++
			// 	// Mark it as done in wait group
			// 	return
			// }
		}
	}
}

func (t *Wire) Resubscribe() {
	// TODO: Write resubscribe logic
}

// Call this function to start the websocket server
func (t *Wire) Serve() {
	t.ServeWithContext(context.Background())
}

// Call this function to start the websocket server with a context
func (t *Wire) ServeWithContext(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	t.serverContext = ctx
	t.cancel = cancel

	for {
		select {
		case <-ctx.Done():
			return
		default:
			// If reconnect attempt exceeds max then close the loop
			if t.reconnectAttempt > t.reconnectMaxRetries {
				t.triggerNoReconnect(t.reconnectAttempt)
				return
			}

			// If its a reconnect then wait exponentially based on reconnect attempt
			if t.reconnectAttempt > 0 {
				nextDelay := time.Duration(math.Pow(2, float64(t.reconnectAttempt))) * time.Second
				if nextDelay > t.reconnectMaxDelay || nextDelay <= 0 {
					nextDelay = t.reconnectMaxDelay
				}

				t.triggerReconnect(t.reconnectAttempt, nextDelay)

				time.Sleep(nextDelay)

				// Close the previous connection if exists
				if t.Conn != nil {
					t.Conn.Close(websocket.StatusGoingAway, "reconnecting")
				}
			}

			// Prepare ticker URL with required params.
			q := t.url.Query()
			q.Set("auth_token", t.accessToken)
			t.url.RawQuery = q.Encode()
			conn, _, err := websocket.Dial(ctx, t.url.String(), &websocket.DialOptions{HTTPClient: &http.Client{
				Timeout: t.connectTimeout,
			}})
			if err != nil {
				t.triggerError(err)

				// If auto reconnect is enabled then try reconneting else return error
				if t.autoreconnect {
					t.reconnectAttempt++
					continue
				}
			}

			// Close the connection when its done.
			defer func() {
				if t.Conn != nil {
					t.Conn.Close(websocket.StatusNormalClosure, "bye")
				}
			}()

			// Assign the current connection to the instance.
			t.Conn = conn

			// Trigger connect callback.
			t.triggerConnect()

			// Resubscribe to stored tokens
			if t.reconnectAttempt > 0 {
				t.Resubscribe()
			}

			// Reset auto reconnect vars
			t.reconnectAttempt = 0

			var wg sync.WaitGroup

			// Receive ticker data in a go routine.
			wg.Add(1)
			go t.readMessage(ctx, &wg)

			// Run watcher to check last ping time and reconnect if required
			if t.autoreconnect {
				wg.Add(1)
				go t.checkConnection(ctx, &wg)
			}

			// Wait for go routines to finish before doing next reconnect
			wg.Wait()
		}
	}
}

// readMessage reads the data in a loop.
func (t *Wire) readMessage(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			// mType, msg, err := t.Conn.Read(ctx)

			mType, msgReader, err := t.Conn.Reader(ctx)
			if err != nil {
				t.triggerError(fmt.Errorf("Error reading data: %v", err))
				return
			}

			msg, err := io.ReadAll(msgReader)
			if err != nil {
				t.triggerError(fmt.Errorf("Error reading data from reader: %v", err))
				return
			}
			// Trigger message.

			t.triggerMessage(int(mType), msg)

			// If binary message then parse and send tick.
			if mType == websocket.MessageBinary {
				ticks, err := t.parseBinary(msg)
				if err != nil {
					t.triggerError(fmt.Errorf("Error parsing data received: %v", err))
				}

				// Trigger individual tick.
				for _, tick := range ticks {
					t.triggerPriceUpdate(tick)
				}
			} else if mType == websocket.MessageText {
				t.processTextMessage(msg)
			}
		}
	}
}
func (t *Wire) parseBinary(inp []byte) ([]*FullQuoteData, error) {
	pkts := t.splitPackets(inp)
	var ticks []*FullQuoteData

	for _, pkt := range pkts {
		tick, err := parsePacket(pkt)
		if err != nil {
			return nil, err
		}

		ticks = append(ticks, tick)
	}

	return ticks, nil
}

func (t *Wire) splitPackets(inp []byte) [][]byte {
	var pkts [][]byte
	if len(inp) < 2 {
		return pkts
	}

	pktLen := binary.LittleEndian.Uint16(inp[0:2])

	j := 2
	for i := 0; i < int(pktLen); i++ {
		pLen := binary.LittleEndian.Uint16(inp[j : j+2])
		pkts = append(pkts, inp[j+2:j+2+int(pLen)])
		j = j + 2 + int(pLen)
	}

	return pkts
}

func parsePacket(buf []byte) (*FullQuoteData, error) {
	// if len(buf) != binary.Size(Quote{}) {
	// 	return nil, errors.New("invalid packet length " + fmt.Sprintf("%d , %d", binary.Size(models.NatPacket{}), len(buf)))
	// }
	requiredLength := 266
	if len(buf) < requiredLength {
		// Calculate the number of zeros to add
		paddingLength := requiredLength - len(buf)
		padding := make([]byte, paddingLength)

		// Pad the slice with zeros
		buf = append(buf, padding...)
	}
	packet := &Packet{}

	err := binary.Read(bytes.NewReader(buf), binary.LittleEndian, packet)
	if err != nil {
		return nil, err
	}

	return packet.transformToQuote(), nil
}

func (t *Wire) triggerPriceUpdate(tick *FullQuoteData) {
	if t.callbacks.onPriceUpdate != nil {
		t.callbacks.onPriceUpdate(tick)
	}
}

func (t *Wire) processTextMessage(inp []byte) {
	var msg SocketMessage
	if err := json.Unmarshal(inp, &msg); err != nil {
		// May be error should be triggered
		return
	}
	t.triggerOrderUpdate(msg)
}
func (t *Wire) Close() error {
	return t.Conn.Close(websocket.StatusGoingAway, "close called")
}

// Stop the wire instance and all the goroutines it has spawned.
func (t *Wire) Stop() {
	if t.cancel != nil {
		t.cancel()
	}
}

// Call this function to subscribe to an instrument
func (t *Wire) Subscribe(exchange ExchangeTypes, token int, mode QuoteModes) {
	ctx, cancelFunc := context.WithTimeout(t.serverContext, 3*time.Second)
	defer cancelFunc()
	message := map[string]interface{}{
		"message_type": "subscribe",
		"exchange":     string(exchange),
		"token":        token,
		"mode":         string(mode),
	}
	aa, _ := json.Marshal(message)
	t.Conn.Write(ctx, websocket.MessageText, aa)
}

// Call this function to unsubscribe an instrument
func (t *Wire) Unsubscribe(exchange ExchangeTypes, token int, mode QuoteModes) {
	ctx, cancelFunc := context.WithTimeout(t.serverContext, 3*time.Second)
	defer cancelFunc()
	message := map[string]interface{}{
		"message_type": "unsubscribe",
		"exchange":     string(exchange),
		"token":        token,
	}
	aa, _ := json.Marshal(message)
	t.Conn.Write(ctx, websocket.MessageText, aa)
}
