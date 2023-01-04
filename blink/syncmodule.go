package blink

func (account *Account) GetSyncModules() error {
	manifest, err := account.GetManifest()

	if err != nil {
		return err
	}

	account.SyncModules = &manifest.SyncModules
	return nil
}
