package pkg

func StrToPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func StrArrayToPtr(s []string) []*string {
	if len(s) == 0 {
		return nil
	}
	ptrs := make([]*string, len(s))
	for i, v := range s {
		ptrs[i] = &v
	}
	return ptrs
}

func CreateImageURL(id string) *string {
	if id == "" {
		return nil
	}
	cfg, err := New()
	if err != nil {
		return nil
	}

	if cfg.CloudFlareR2 == "" {
		return nil
	}

	return StrToPtr(cfg.CloudFlareR2 + "/" + id + ".png")
}
