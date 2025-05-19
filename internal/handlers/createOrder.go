package handlers

import (
	"compress/gzip"
	"errors"
	"io"
	"net/http"

	"github.com/kirsh-nat/gophermart.git/internal/app"
	"github.com/kirsh-nat/gophermart.git/internal/models/order"
	orderservices "github.com/kirsh-nat/gophermart.git/internal/services/orderServices"
)

func (h *URLHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	if !h.checkMethod(w, r, http.MethodPost) {
		return
	}

	user, ok := h.getUserFromToken(w, r)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	var body io.Reader = r.Body
	if r.Header.Get("Content-Encoding") == "gzip" {
		gz, err := gzip.NewReader(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Can't create gzip reader"))
			return
		}
		defer gz.Close()
		body = gz
	}

	reqNumber, err := io.ReadAll(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Can't read request body"))
		return
	}

	errNumber := orderservices.CheckNumber(string(reqNumber))
	if errNumber != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(""))
		return
	}

	orderModel := order.NewOrderModel(h.db)
	_, err = orderModel.Create(r.Context(), &order.Order{userID: user.ID, Number: string(reqNumber)})
	if err != nil {
		var userErr *order.UserNumberExistsError
		if errors.As(err, &userErr) {
			w.WriteHeader(http.StatusOK)
			return
		}
		var dErr *order.NumberExists
		if errors.As(err, &dErr) {
			w.WriteHeader(http.StatusConflict)
			return
		}
		app.Sugar.Error(err)
		h.StatusServerError(w, r)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	_, _ = w.Write([]byte("Order is accepteed"))
}
