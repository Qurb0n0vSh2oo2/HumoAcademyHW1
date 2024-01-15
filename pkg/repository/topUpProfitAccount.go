package repository

import "fmt"

func (repo *Repository) TopUpProfitAccount(amount float64) {
  account, err := repo.GetAccount("profit")

  if err != nil {
    city, err := repo.GetCity("Dushanbe")

    if err != nil {
      fmt.Println("Ошибка, данный город не найден, звоните по номеру 544")
      return
    }

    err = repo.AddAccount("profit", 0, "544", "password", city)
    if err != nil {
      return
    }

    account, _ = repo.GetAccount("profit")
  }

  err = repo.ChangeRecipientBalance(account.Name,   account.Balance + amount)
  if err != nil {
    return
  }
}
