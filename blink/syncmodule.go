package blink

// Returns sync modules from manifest
func (account *Account) GetSyncModules() (*[]SyncModule, error) {
	manifest, err := account.GetManifest()

	if err != nil {
		return nil, err
	}

	return &manifest.SyncModules, nil
}
