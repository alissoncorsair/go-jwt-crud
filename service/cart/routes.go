package cart

import (
	"fmt"
	"net/http"

	"github.com/alissoncorsair/goapi/service/auth"
	"github.com/alissoncorsair/goapi/types"
	"github.com/alissoncorsair/goapi/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	orderStore   types.OrderStore
	productStore types.ProductStore
	userStore    types.UserStore
}

func NewHandler(orderStore types.OrderStore, productStore types.ProductStore, userStore types.UserStore) *Handler {
	return &Handler{
		orderStore:   orderStore,
		productStore: productStore,
		userStore:    userStore,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", auth.WithJWTAuth(h.handleCheckout, h.userStore)).Methods(http.MethodPost)
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {
	userId := auth.GetUserIDFromContext(r.Context())
	var cart types.CartCheckoutPayload

	if err := utils.ParseJSON(r, &cart); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(cart); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	productIDS, err := getCartItemsIDS(cart.Items)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	products, err := h.productStore.GetProductsByID(productIDS)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	orderId, totalPrice, err := h.createOrder(products, cart.Items, userId)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	orderResponse := types.CheckoutResponse{
		OrderID:    orderId,
		TotalPrice: totalPrice,
	}

	utils.WriteJSON(w, http.StatusOK, orderResponse)

}
