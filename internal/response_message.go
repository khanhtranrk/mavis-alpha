package domain

import "time"

type RespnoseMessage struct {
  ResponseId string
  RequestId string
  SenderId string
  ReceiverId string
  SafeKey string
  Message string
  ChainKey string
  ConKey string
  AcceptedAt time.Time
  InProgressAt time.Time
  CompletedAt time.Time
}
