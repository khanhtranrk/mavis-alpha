package domain

import (
  "time"
)

type RequestMessage struct {
  RequestId string
  SenderId string
  ReceiverId string
  SafeKey string
  Message string
  ChainKey string
  ConKey string
  RequestAt time.Time
}

def NewRequestMessage(SenderId string, ReceiverId string Message)
