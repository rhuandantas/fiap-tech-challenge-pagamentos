// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"fiap-tech-challenge-pagamentos/client"
	"fiap-tech-challenge-pagamentos/internal/adapters/http"
	"fiap-tech-challenge-pagamentos/internal/adapters/http/handlers"
	repository2 "fiap-tech-challenge-pagamentos/internal/adapters/repository"
	"fiap-tech-challenge-pagamentos/internal/core/usecase"
	"github.com/rhuandantas/fiap-tech-challenge-commons/pkg/db/mysql"
	"github.com/rhuandantas/fiap-tech-challenge-commons/pkg/middlewares/auth"
	"github.com/rhuandantas/fiap-tech-challenge-commons/pkg/util"
)

// Injectors from wire.go:

func InitializeWebServer() (*http.Server, error) {
	healthCheck := handlers.NewHealthCheck()
	dbConnector := repository.NewMySQLConnector()
	pagamentoRepo := repository2.NewPagamentoRepo(dbConnector)
	pesquisaPagamento := usecase.NewPesquisaPagamento(pagamentoRepo)
	validator := util.NewCustomValidator()
	token := auth.NewJwtToken()
	atualizaPagamento := usecase.NewAtualizaPagamento(pagamentoRepo)
	pedido := client.NewPedido()
	realizarCheckout := usecase.NewRealizaCheckout(pagamentoRepo, atualizaPagamento, pedido)
	pagamento := handlers.NewPagamento(pesquisaPagamento, validator, token, realizarCheckout)
	server := http.NewAPIServer(healthCheck, pagamento)
	return server, nil
}
