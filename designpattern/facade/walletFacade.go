package main

type walletFacade struct {
	*account
	*securityCode
	*wallet
}

func NewWalletFacade(accountName, code string) *walletFacade {
	return &walletFacade{
		account:      &account{name: accountName},
		securityCode: &securityCode{password: code},
		wallet:       &wallet{balance: 0},
	}
}

func (w *walletFacade) check(name, code string) error {
	err := w.account.Check(name)
	if err != nil {
		return err
	}
	err = w.securityCode.checkSecurityCode(code)
	if err != nil {
		return err
	}
	return nil
}

func (w *walletFacade) AddAmount(name, code string, amount int) error {
	if err := w.check(name, code); err != nil {
		return err
	}
	if err := w.wallet.addBalance(amount); err != nil {
		return err
	}
	return nil
}
func (w *walletFacade) ReduceAmount(name, code string, amount int) error {
	if err := w.check(name, code); err != nil {
		return err
	}
	if err := w.wallet.reduceBalance(amount); err != nil {
		return err
	}
	return nil
}
